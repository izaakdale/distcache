package app

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
)

type clusterSpec struct {
	BindAddr      string `envconfig:"BIND_ADDR"`
	BindPort      string `envconfig:"BIND_PORT"`
	AdvertiseAddr string `envconfig:"ADVERTISE_ADDR"`
	AdvertisePort string `envconfig:"ADVERTISE_PORT"`
	ClusterAddr   string `envconfig:"CLUSTER_ADDR"`
	ClusterPort   string `envconfig:"CLUSTER_PORT"`
	Name          string `envconfig:"NAME"`
}

func setupCluster(bindAddr, bindPort, advertiseAddr, advertisePort, name string) (*serf.Serf, chan serf.Event, error) {
	// allows separation of members with the same name from env
	id := strings.Split(uuid.NewString(), "-")[0]
	uname := fmt.Sprintf("%s-%s", spec.clutserCfg.Name, id)

	// since in k8s you will want to advertise the cluster ip service which changes,
	// we will enter the name in the format <svc-name>.<namespace>.svc.cluster.local to resolve the ip
	res, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", spec.clutserCfg.AdvertiseAddr, spec.clutserCfg.AdvertisePort))
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

	events := make(chan serf.Event)
	conf.EventCh = events

	cluster, err := serf.Create(conf)
	if err != nil {
		log.Printf("inside error: %e\n", err)
		return nil, nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{res.String()}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	return cluster, events, nil
}

func handleSerfEvent(e serf.Event, cluster *serf.Serf) {
	switch e.EventType() {
	case serf.EventMemberJoin:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(cluster, member) {
				continue
			}
			handleJoin(member)
		}
	case serf.EventMemberLeave, serf.EventMemberFailed:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(cluster, member) {
				continue
			}
			handleLeave(member)
		}
	case serf.EventMemberUpdate:
		for _, member := range e.(serf.MemberEvent).Members {
			if isLocal(cluster, member) {
				continue
			}
			handleUpdate(member)
		}
	case serf.EventUser:
		handleCustomEvent(e.(serf.UserEvent))
	}
}

func handleJoin(m serf.Member) {
	// 1. create a new grpc connection to the member
	// 2. keep conntection local to reuse
	// 3. request all its records
	// --- this may seem counter intuitive, no matter how new a member is it might have received a key store request,
	// --- also, you receive member joins from other members when you first join, so a new member will collect records

	log.Printf("member joined %s @ %s\n", m.Name, m.Addr)
}

func handleCustomEvent(e serf.UserEvent) error {
	// custom events will be used to relay information about a new store request
	// since the load balancer will send the request to one server, we need to broadcast what to store
	// 1. store information from event

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

type throwAway struct {
}

func (t *throwAway) Write(b []byte) (int, error) {
	return 0, nil
}
