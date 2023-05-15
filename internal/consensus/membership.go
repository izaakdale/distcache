package consensus

import (
	"github.com/hashicorp/raft"
	"log"
)

func (d *DistributedCache) Join(name, addr string) error {
	log.Printf("-------- member joining --------")
	configFuture := d.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return err
	}

	serverID := raft.ServerID(name)
	serverAddr := raft.ServerAddress(addr)
	log.Printf("%s @ %s", serverID, serverAddr)
	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == serverID || srv.Address == serverAddr {
			// server already join
			return nil
		}
		// remove exisiting server
		removeFuture := d.raft.RemoveServer(serverID, 0, 0)
		if err := removeFuture.Error(); err != nil {
			return err
		}
	}
	log.Printf("------!!!!-----%s-----!!!!!------", serverID)
	addFuture := d.raft.AddVoter(serverID, serverAddr, 0, 0)
	if err := addFuture.Error(); err != nil {
		return err
	}
	return nil
}

func (d *DistributedCache) Leave(name string) error {
	log.Printf("--------- member leaving --------")
	removeFuture := d.raft.RemoveServer(raft.ServerID(name), 0, 0)
	return removeFuture.Error()
}
