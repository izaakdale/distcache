package app

import (
	"context"

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
	err := store.Insert(req.Key, req.Value)
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
