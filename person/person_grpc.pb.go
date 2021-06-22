// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: person/person.proto

package person

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

// GetPersonClient is the client API for GetPerson service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetPersonClient interface {
	// get data
	PersonData(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
}

type getPersonClient struct {
	cc grpc.ClientConnInterface
}

func NewGetPersonClient(cc grpc.ClientConnInterface) GetPersonClient {
	return &getPersonClient{cc}
}

func (c *getPersonClient) PersonData(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/main.GetPerson/PersonData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetPersonServer is the server API for GetPerson service.
// All implementations must embed UnimplementedGetPersonServer
// for forward compatibility
type GetPersonServer interface {
	// get data
	PersonData(context.Context, *Person) (*Person, error)
	mustEmbedUnimplementedGetPersonServer()
}

// UnimplementedGetPersonServer must be embedded to have forward compatible implementations.
type UnimplementedGetPersonServer struct {
}

func (UnimplementedGetPersonServer) PersonData(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PersonData not implemented")
}
func (UnimplementedGetPersonServer) mustEmbedUnimplementedGetPersonServer() {}

// UnsafeGetPersonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetPersonServer will
// result in compilation errors.
type UnsafeGetPersonServer interface {
	mustEmbedUnimplementedGetPersonServer()
}

func RegisterGetPersonServer(s grpc.ServiceRegistrar, srv GetPersonServer) {
	s.RegisterService(&GetPerson_ServiceDesc, srv)
}

func _GetPerson_PersonData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetPersonServer).PersonData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GetPerson/PersonData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetPersonServer).PersonData(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

// GetPerson_ServiceDesc is the grpc.ServiceDesc for GetPerson service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetPerson_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.GetPerson",
	HandlerType: (*GetPersonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PersonData",
			Handler:    _GetPerson_PersonData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person/person.proto",
}
