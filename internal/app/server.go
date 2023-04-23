package app

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis"
	msg "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ensure our server adheres to grpc cache server
var _ (msg.CacheServer) = (*Server)(nil)

type Server struct {
	msg.UnimplementedCacheServer
}

func (s *Server) Store(ctx context.Context, req *msg.StoreRequest) (*msg.StoreResponse, error) {
	err := store.Insert(req.Record.Key, req.Record.Value, int(req.Ttl))
	if err != nil {
		return nil, err
	}
	return &msg.StoreResponse{}, nil
}
func (s *Server) Fetch(ctx context.Context, req *msg.FetchRequest) (*msg.FetchResponse, error) {
	val, err := store.Fetch(req.Key)
	if err != nil {
		if err == redis.Nil {
			st := status.New(codes.NotFound, "no record for key")
			return nil, st.Err()
		}
		return nil, err
	}
	return &msg.FetchResponse{Value: val}, nil
}
func (s *Server) AllKeys(ctx context.Context, req *msg.AllKeysRequest) (*msg.AllKeysResponse, error) {
	keys, err := store.AllKeys(req.Pattern)
	if err != nil {
		if err != nil {
			if err == redis.Nil {
				st := status.New(codes.NotFound, "no keys")
				return nil, st.Err()
			}
			return nil, err
		}
		return nil, err
	}
	return &msg.AllKeysResponse{Keys: keys}, nil
}
func (c *Server) AllRecords(req *msg.AllRecordsRequest, srv msg.Cache_AllRecordsServer) error {
	for _, key := range req.Keys {
		val, err := store.Fetch(key)
		if err != nil {
			if err == redis.Nil {
				st := status.New(codes.NotFound, "a key provided was not found")
				return st.Err()
			}
			return err
		}
		ttl, err := store.GetTTL(key)
		if err != nil {
			return err
		}
		if ttl == nil {
			return errors.New("ttl returned as nil from store")
		}
		log.Printf("%+v:%s\n", key, val)
		srv.Send(&msg.AllRecordsResponse{
			Record: &msg.KVRecord{
				Key:   key,
				Value: val,
			},
			Ttl: *ttl,
		})
	}
	return nil
}
