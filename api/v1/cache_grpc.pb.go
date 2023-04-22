// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/v1/cache.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type cacheClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheClient(cc grpc.ClientConnInterface) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Cache/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Cache/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
// All implementations must embed UnimplementedCacheServer
// for forward compatibility
type CacheServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	mustEmbedUnimplementedCacheServer()
}

// UnimplementedCacheServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (UnimplementedCacheServer) Store(context.Context, *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedCacheServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedCacheServer) mustEmbedUnimplementedCacheServer() {}

// UnsafeCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServer will
// result in compilation errors.
type UnsafeCacheServer interface {
	mustEmbedUnimplementedCacheServer()
}

func RegisterCacheServer(s grpc.ServiceRegistrar, srv CacheServer) {
	s.RegisterService(&Cache_ServiceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Cache/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cache_ServiceDesc is the grpc.ServiceDesc for Cache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Cache_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cache.proto",
}
