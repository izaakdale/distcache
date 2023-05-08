package consensus

import (
	"crypto/tls"
	"github.com/hashicorp/raft"
	"github.com/izaakdale/distcache/internal/discovery"
	"github.com/izaakdale/distcache/internal/store"
	"io"
	"net"
)

var (
	// ensure that our structs implement these interfaces
	_ discovery.Handler = (*DistributedCache)(nil)
	_ raft.StreamLayer  = (*StreamLayer)(nil)
	_ raft.FSM          = (*fsm)(nil)
	_ raft.FSMSnapshot  = (*snapshot)(nil)
)

type (
	DistributedCache struct {
		config Config
		raft   *raft.Raft
	}
	Config struct {
		Txer store.Transactioner
		Raft *Raft
	}
	Raft struct {
		raft.Config
		BindAddr    string
		StreamLayer *StreamLayer
		Bootstrap   bool
	}
	StreamLayer struct {
		ln              net.Listener
		serverTLSConfig *tls.Config
		peerTLSConfig   *tls.Config
	}
	snapshot struct {
		reader io.Reader
	}
)
