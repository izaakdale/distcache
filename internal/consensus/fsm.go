package consensus

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
	v1 "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"google.golang.org/protobuf/proto"
)

type RequestType uint8

const AppendRequestType RequestType = 0

type fsm struct {
	txer store.Transactioner
}

func (f *fsm) Apply(record *raft.Log) any {
	buf := record.Data
	reqType := RequestType(buf[0])
	switch reqType {
	case AppendRequestType:
		return f.applyAppend(buf[1:])
	}
	return nil
}

func (f *fsm) applyAppend(b []byte) any {
	var req v1.StoreRequest
	err := proto.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	err = f.txer.Insert(req.Record.Key, req.Record.Key, 0) //TODO remove 0
	if err != nil {
		return err
	}
	return &v1.StoreResponse{}
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	log.Printf("attempted snapshot in fsm")
	return &snapshot{
		//f.txer.Reader(),
	}, nil
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	log.Printf("attempted persist in fsm")
	//if _, err := io.Copy(sink, s.reader); err != nil {
	//	_ = sink.Cancel()
	//	return err
	//}
	//return sink.Close()
	return nil
}

func (s *snapshot) Release() {
	log.Printf("attempted release in fsm")
}

func (f *fsm) Restore(snapshot io.ReadCloser) error {
	log.Printf("attempted restore in fsm")
	//	var b []byte
	//loop:
	//	for {
	//		_, err := snapshot.Read(b)
	//		switch err {
	//		case io.EOF:
	//			break loop
	//		case nil:
	//			return err
	//		}
	//
	//		var req v1.StoreRequest
	//		if err = proto.Unmarshal(b, &req); err != nil {
	//			return err
	//		}
	//		if err = f.txer.Insert(req.Record.Key, req.Record.Value, 0); err != nil {
	//			return err
	//		}
	//	}
	return nil
}
