package app

import (
	"context"
	"fmt"
	"strings"

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
	var notFoundKeys []string
	var notFound bool = false
	for _, key := range req.Keys {
		val, err := store.Fetch(key)
		if err != nil {
			if err == redis.Nil {
				notFound = true
				notFoundKeys = append(notFoundKeys, key)
				continue
			} else {
				return err
			}
		}
		ttl, err := store.GetTTL(key)
		if err != nil {
			return err
		}
		if ttl == nil {
			st := status.New(codes.Internal, "ttl returned as nil from store")
			return st.Err()
		}
		srv.Send(&msg.AllRecordsResponse{
			Record: &msg.KVRecord{
				Key:   key,
				Value: val,
			},
			Ttl: *ttl,
		})
	}
	if notFound {
		// TODO logger
		// ultimately i do want to log here, but not just a printf
		// log.Printf("not found keys: %+v\n", notFoundKeys)

		// list will be in format "not_found_keys:key1/key2/key3"
		// clients can split on : and / to obtain keys
		b := strings.Builder{}
		b.WriteString("not_found_keys:")
		for _, k := range notFoundKeys {
			b.WriteString(fmt.Sprintf("%s/", k))
		}

		st := status.New(codes.NotFound, b.String())
		return st.Err()
	}
	return nil
}
