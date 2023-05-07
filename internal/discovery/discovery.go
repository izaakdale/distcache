package discovery

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
	"log"
	"net"
	"strconv"
	"strings"
)

type (
	Membership struct {
		handler   Handler
		EventChan chan serf.Event
		Node      *serf.Serf
	}
	Handler interface {
		Join(name, addr string) error
		Leave(name string) error
	}
	Specification struct {
		BindAddr      string `envconfig:"BIND_ADDR"`
		BindPort      string `envconfig:"BIND_PORT"`
		AdvertiseAddr string `envconfig:"ADVERTISE_ADDR"`
		AdvertisePort string `envconfig:"ADVERTISE_PORT"`
		ClusterAddr   string `envconfig:"CLUSTER_ADDR"`
		ClusterPort   string `envconfig:"CLUSTER_PORT"`
		Name          string `envconfig:"NAME"`
	}
)

func NewMembership(handler Handler, spec Specification, grpcPort int) (*Membership, error) {
	// allows separation of members with the same name from env
	id := strings.Split(uuid.NewString(), "-")[0]
	uname := fmt.Sprintf("%s-%s", spec.Name, id)

	// since in k8s you will want to advertise the cluster ip service which changes,
	// we will enter the name in the format <svc-name>.<namespace>.svc.cluster.local to resolve the ip
	res, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", spec.AdvertiseAddr, spec.AdvertisePort))
	if err != nil {
		return nil, err
	}

	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = res.IP.String()
	conf.MemberlistConfig.AdvertisePort = res.Port
	conf.MemberlistConfig.BindAddr = spec.BindAddr
	conf.MemberlistConfig.BindPort, _ = strconv.Atoi(spec.BindPort)
	conf.MemberlistConfig.ProtocolVersion = 3 // Version 3 enable the ability to bind different port for each agent
	conf.NodeName = uname

	conf.Tags = map[string]string{
		rpcAddrTag: fmt.Sprintf("%s:%d", spec.BindAddr, grpcPort),
	}

	events := make(chan serf.Event)
	conf.EventCh = events

	node, err := serf.Create(conf)
	if err != nil {
		log.Printf("inside error: %e\n", err)
		return nil, errors.Wrap(err, "couldn't create serf instance")
	}

	_, err = node.Join([]string{res.String()}, true)
	if err != nil {
		log.Printf("couldn't join cluster, starting own: %v\n", err)
	}

	//return node, events, nil
	return &Membership{
		handler,
		events,
		node,
	}, nil
}
