package discovery

import (
	"github.com/hashicorp/serf/serf"
	"log"
)

var rpcAddrTag = "rpc_addr"

func (mship *Membership) HandleSerfEvent(e serf.Event, node *serf.Serf) {
	switch e.EventType() {
	case serf.EventMemberJoin:
		for _, member := range e.(serf.MemberEvent).Members {
			if mship.isLocal(node, member) {
				continue
			}
			err := mship.handleJoin(member)
			if err != nil {
				log.Printf("error handling member join: %e", err)
			}
		}
	case serf.EventMemberLeave, serf.EventMemberFailed:
		for _, member := range e.(serf.MemberEvent).Members {
			if mship.isLocal(node, member) {
				continue
			}
			err := mship.handleLeave(member)
			if err != nil {
				log.Printf("error while handling leave: %e", err)
			}
		}
	}
}

func (mship *Membership) handleJoin(m serf.Member) error {
	return mship.handler.Join(m.Name, m.Tags[rpcAddrTag])
}
func (mship *Membership) handleLeave(m serf.Member) error {
	return mship.handler.Leave(m.Name)
}
func (mship *Membership) isLocal(c *serf.Serf, m serf.Member) bool {
	// checks whether the event has come from itself
	return c.LocalMember().Name == m.Name
}
