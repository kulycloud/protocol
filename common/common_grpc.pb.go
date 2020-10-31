// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ComponentClient is the client API for Component service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	RegisterStorageEndpoints(ctx context.Context, in *EndpointList, opts ...grpc.CallOption) (*Empty, error)
}

type componentClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentClient(cc grpc.ClientConnInterface) ComponentClient {
	return &componentClient{cc}
}

func (c *componentClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Component/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentClient) RegisterStorageEndpoints(ctx context.Context, in *EndpointList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Component/RegisterStorageEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentServer is the server API for Component service.
// All implementations must embed UnimplementedComponentServer
// for forward compatibility
type ComponentServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	RegisterStorageEndpoints(context.Context, *EndpointList) (*Empty, error)
	mustEmbedUnimplementedComponentServer()
}

// UnimplementedComponentServer must be embedded to have forward compatible implementations.
type UnimplementedComponentServer struct {
}

func (UnimplementedComponentServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedComponentServer) RegisterStorageEndpoints(context.Context, *EndpointList) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStorageEndpoints not implemented")
}
func (UnimplementedComponentServer) mustEmbedUnimplementedComponentServer() {}

// UnsafeComponentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentServer will
// result in compilation errors.
type UnsafeComponentServer interface {
	mustEmbedUnimplementedComponentServer()
}

func RegisterComponentServer(s grpc.ServiceRegistrar, srv ComponentServer) {
	s.RegisterService(&_Component_serviceDesc, srv)
}

func _Component_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Component/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Component_RegisterStorageEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentServer).RegisterStorageEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Component/RegisterStorageEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentServer).RegisterStorageEndpoints(ctx, req.(*EndpointList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Component_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Component",
	HandlerType: (*ComponentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Component_Ping_Handler,
		},
		{
			MethodName: "RegisterStorageEndpoints",
			Handler:    _Component_RegisterStorageEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}