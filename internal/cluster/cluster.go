package cluster

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/serf/serf"
	api "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type Specification struct {
	BindAddr      string `envconfig:"BIND_ADDR"`
	BindPort      string `envconfig:"BIND_PORT"`
	AdvertiseAddr string `envconfig:"ADVERTISE_ADDR"`
	AdvertisePort string `envconfig:"ADVERTISE_PORT"`
	ClusterAddr   string `envconfig:"CLUSTER_ADDR"`
	ClusterPort   string `envconfig:"CLUSTER_PORT"`
	Name          string `envconfig:"NAME"`
}

func SetupNode(bindAddr, bindPort, advertiseAddr, advertisePort, name string, grpcPort int) (*serf.Serf, chan serf.Event, error) {
	// allows separation of members with the same name from env
	id := strings.Split(uuid.NewString(), "-")[0]
	uname := fmt.Sprintf("%s-%s", name, id)

	// since in k8s you will want to advertise the cluster ip service which changes,
	// we will enter the name in the format <svc-name>.<namespace>.svc.cluster.local to resolve the ip
	res, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", advertiseAddr, advertisePort))
	if err != nil {
		return nil, nil, err
	}

	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = res.IP.String()
	conf.MemberlistConfig.AdvertisePort = res.Port
	conf.MemberlistConfig.BindAddr = bindAddr
	conf.MemberlistConfig.BindPort, _ = strconv.Atoi(bindPort)
	conf.MemberlistConfig.ProtocolVersion = 3 // Version 3 enable the ability to bind different port for each agent
	conf.NodeName = uname

	conf.Tags = map[string]string{
		"grpc_addr": fmt.Sprintf("%s:%d", bindAddr, grpcPort),
	}

	events := make(chan serf.Event)
	conf.EventCh = events

	node, err := serf.Create(conf)
	if err != nil {
		log.Printf("inside error: %e\n", err)
		return nil, nil, errors.Wrap(err, "couldn't create serf instance")
	}

	_, err = node.Join([]string{res.String()}, true)
	if err != nil {
		log.Printf("couldn't join cluster, starting own: %v\n", err)
	}

	return node, events, nil
}

func HandleSerfEvent(e serf.Event, node *serf.Serf) {
	switch e.EventType() {
	case serf.EventMemberJoin:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(node, member) {
				continue
			}
			err := handleJoin(member)
			if err != nil {
				log.Printf("error handling member join: %e", err)
			}
		}
	case serf.EventMemberLeave, serf.EventMemberFailed:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(node, member) {
				continue
			}
			handleLeave(member)
		}
	case serf.EventMemberUpdate:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(node, member) {
				continue
			}
			handleUpdate(member)
		}
	case serf.EventUser:
		err := handleCustomEvent(e.(serf.UserEvent))
		if err != nil {
			log.Printf("error handling custom event: %e", err)
		}
	}
}

func handleJoin(m serf.Member) error {
	log.Printf("member joined %s @ %s\n", m.Name, m.Addr)
	grpcSocket, ok := m.Tags["grpc_addr"]
	if !ok {
		return fmt.Errorf("grpc address tag was not included in the event")
	}

	var err error
	conn, err := grpc.Dial(grpcSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to %s", grpcSocket)
	}
	client := api.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	keyResp, err := client.AllKeys(ctx, &api.AllKeysRequest{
		Pattern: "",
	})
	if err != nil {
		return err
	}
	recStream, err := client.AllRecords(ctx, &api.AllRecordsRequest{
		Keys: keyResp.Keys,
	})
	if err != nil {
		return err
	}
	for {
		record, err := recStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err = store.Insert(record.Record.Key, record.Record.Value, int(record.Ttl)); err != nil {
			return err
		}
	}
	return nil
}

func handleCustomEvent(e serf.UserEvent) error {
	// custom events will be used to relay information about a new store request
	// since the load balancer will send the request to one server, we need to broadcast what to store
	// 1. store information from event
	var req api.StoreRequest
	if err := proto.Unmarshal(e.Payload, &req); err != nil {
		return err
	}
	if err := store.Insert(req.Record.Key, req.Record.Value, int(req.Ttl)); err != nil {
		return err
	}
	return nil
}

func isLocal(c *serf.Serf, m serf.Member) bool {
	// checks whether the event has come from itself
	return c.LocalMember().Name == m.Name
}

// TODO do we need these, basically not doing anything and could be ignored
func handleLeave(m serf.Member) {
	// no action needed, just say goodbye
	log.Printf("member leaving %s @ %s\n", m.Name, m.Addr)
}
func handleUpdate(m serf.Member) {
	// again, no action needed
	log.Printf("receiving update from %s @ %s\n", m.Name, m.Addr)
}
