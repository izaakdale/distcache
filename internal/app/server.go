package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	api "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ensure our server adheres to grpc cache server
var _ api.CacheServer = (*Server)(nil)

type Server struct {
	api.UnimplementedCacheServer
}

func (s *Server) Store(ctx context.Context, req *api.StoreRequest) (*api.StoreResponse, error) {
	err := store.Insert(req.Record.Key, req.Record.Value, int(req.Ttl))
	if err != nil {
		return nil, err
	}
	//bytes, err := proto.Marshal(req)
	//if err != nil {
	//	return nil, err
	//}
	//if err := clusterMembership.UserEvent("value-stored", bytes, false); err != nil {
	//	return nil, err
	//}
	return &api.StoreResponse{}, nil
}
func (s *Server) Fetch(ctx context.Context, req *api.FetchRequest) (*api.FetchResponse, error) {
	val, err := store.Fetch(req.Key)
	if err != nil {
		if err == redis.Nil {
			st := status.New(codes.NotFound, "no record for key")
			return nil, st.Err()
		}
		return nil, err
	}
	return &api.FetchResponse{Value: val}, nil
}
func (s *Server) AllKeys(ctx context.Context, req *api.AllKeysRequest) (*api.AllKeysResponse, error) {
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
	return &api.AllKeysResponse{Keys: keys}, nil
}
func (s *Server) AllRecords(req *api.AllRecordsRequest, srv api.Cache_AllRecordsServer) error {
	var notFoundKeys []string
	notFound := false
	for _, key := range req.Keys {
		val, err := store.Fetch(key)
		if err != nil {
			if err == redis.Nil {
				// don't want to stop here, retrieve as many keys as possible and then alert to missing ones.
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
		err = srv.Send(&api.AllRecordsResponse{
			Record: &api.KVRecord{
				Key:   key,
				Value: val,
			},
			Ttl: *ttl,
		})
		if err != nil {
			return err
		}
	}
	if notFound {
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
