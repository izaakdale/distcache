package consensus

import (
	"crypto/tls"
	"github.com/hashicorp/raft"
	"github.com/izaakdale/distcache/internal/discovery"
	"net"
)

var (
	// ensure that our structs implement these interfaces
	_ discovery.Handler = (*DistributedCache)(nil)
	_ raft.StreamLayer  = (*StreamLayer)(nil)
)

type DistributedCache struct {
	config Config
	raft   *raft.Raft
}

type Config struct {
	Raft struct {
		raft.Config
		BindAddr    string
		StreamLayer *StreamLayer
		Bootstrap   bool
	}
}

type StreamLayer struct {
	ln              net.Listener
	serverTLSConfig *tls.Config
	peerTLSConfig   *tls.Config
}
