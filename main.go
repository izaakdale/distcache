package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/hashicorp/serf/serf"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var spec specification

type specification struct {
	BindAddr      string `envconfig:"BIND_ADDR"`
	BindPort      string `envconfig:"BIND_PORT"`
	AdvertiseAddr string `envconfig:"ADVERTISE_ADDR"`
	AdvertisePort string `envconfig:"ADVERTISE_PORT"`
	ClusterAddr   string `envconfig:"CLUSTER_ADDR"`
	ClusterPort   string `envconfig:"CLUSTER_PORT"`
	Name          string `envconfig:"NAME"`
}

func main() {
	// app.Must(app.New()).Run()

	envconfig.Process("", &spec)
	log.Printf("%+v\n", spec)

	// since in k8s you will want to advertise the cluster ip service which changes,
	// we will enter the name in the format <svc-name>.<namespace>.svc.cluster.local to resolve the ip
	res, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", spec.AdvertiseAddr, spec.AdvertisePort))
	if err != nil {
		panic(err)
	}

	id := strings.Split(uuid.NewString(), "-")[0]
	name := fmt.Sprintf("%s-%s", spec.Name, id)

	cluster, evCh, err := setupCluster(
		spec.BindAddr,
		spec.BindPort, // BIND defines where the agent listen for incomming connection
		res.IP.String(),
		spec.AdvertisePort, // ADVERTISE defines where the agent is reachable, often it the same as BIND
		spec.ClusterAddr,
		spec.ClusterPort, // CLUSTER is the address of a first agent
		name)             // NAME must be unique in a cluster
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

Loop:
	for {
		select {
		case e := <-evCh:
			{
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
				}
			}
		case <-c:
			{
				cluster.Leave()
				break Loop
			}
		}
	}
	cluster.Leave()
	os.Exit(1)
}

func setupCluster(bindAddr, bindPort, advertiseAddr, advertisePort, clusterAddr, clusterPort, name string) (*serf.Serf, chan serf.Event, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr
	conf.MemberlistConfig.AdvertisePort, _ = strconv.Atoi(advertisePort)
	conf.MemberlistConfig.BindAddr = bindAddr
	conf.MemberlistConfig.BindPort, _ = strconv.Atoi(bindPort)
	conf.MemberlistConfig.ProtocolVersion = 3 // Version 3 enable the ability to bind different port for each agent
	conf.NodeName = name

	events := make(chan serf.Event)
	conf.EventCh = events

	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Couldn't create cluster")
	}

	if clusterAddr != "" && clusterPort != "" {
		_, err = cluster.Join([]string{clusterAddr + ":" + clusterPort}, true)
		if err != nil {
			log.Printf("Couldn't join cluster, starting own: %v\n", err)
		}
	}

	return cluster, events, nil
}

func handleJoin(m serf.Member) {
	log.Printf("member joined %s @ %s\n", m.Name, m.Addr)
}

func handleLeave(m serf.Member) {
	log.Printf("member leaving %s @ %s\n", m.Name, m.Addr)
}

func handleUpdate(m serf.Member) {
	log.Printf("receiving update from %s @ %s\n", m.Name, m.Addr)
}

func isLocal(c *serf.Serf, m serf.Member) bool {
	return c.LocalMember().Name == m.Name
}
