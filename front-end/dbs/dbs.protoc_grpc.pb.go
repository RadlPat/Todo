// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: dbs.protoc

package dbs

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

// DbServiceClient is the client API for DbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbServiceClient interface {
	GetAllToDos(ctx context.Context, in *TodoListRequest, opts ...grpc.CallOption) (*TodoListResponse, error)
}

type dbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDbServiceClient(cc grpc.ClientConnInterface) DbServiceClient {
	return &dbServiceClient{cc}
}

func (c *dbServiceClient) GetAllToDos(ctx context.Context, in *TodoListRequest, opts ...grpc.CallOption) (*TodoListResponse, error) {
	out := new(TodoListResponse)
	err := c.cc.Invoke(ctx, "/dbs.DbService/GetAllToDos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServiceServer is the server API for DbService service.
// All implementations must embed UnimplementedDbServiceServer
// for forward compatibility
type DbServiceServer interface {
	GetAllToDos(context.Context, *TodoListRequest) (*TodoListResponse, error)
	mustEmbedUnimplementedDbServiceServer()
}

// UnimplementedDbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDbServiceServer struct {
}

func (UnimplementedDbServiceServer) GetAllToDos(context.Context, *TodoListRequest) (*TodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllToDos not implemented")
}
func (UnimplementedDbServiceServer) mustEmbedUnimplementedDbServiceServer() {}

// UnsafeDbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServiceServer will
// result in compilation errors.
type UnsafeDbServiceServer interface {
	mustEmbedUnimplementedDbServiceServer()
}

func RegisterDbServiceServer(s grpc.ServiceRegistrar, srv DbServiceServer) {
	s.RegisterService(&DbService_ServiceDesc, srv)
}

func _DbService_GetAllToDos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetAllToDos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbs.DbService/GetAllToDos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetAllToDos(ctx, req.(*TodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DbService_ServiceDesc is the grpc.ServiceDesc for DbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbs.DbService",
	HandlerType: (*DbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllToDos",
			Handler:    _DbService_GetAllToDos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbs.protoc",
}
