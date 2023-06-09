// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/pledge/v1/animal.proto

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

// AnimalClient is the client API for Animal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimalClient interface {
	CreateAnimal(ctx context.Context, in *CreateAnimalRequest, opts ...grpc.CallOption) (*CreateAnimalReply, error)
	UpdateAnimal(ctx context.Context, in *UpdateAnimalRequest, opts ...grpc.CallOption) (*UpdateAnimalReply, error)
	DeleteAnimal(ctx context.Context, in *DeleteAnimalRequest, opts ...grpc.CallOption) (*DeleteAnimalReply, error)
	GetAnimal(ctx context.Context, in *GetAnimalRequest, opts ...grpc.CallOption) (*GetAnimalReply, error)
	ListAnimal(ctx context.Context, in *ListAnimalRequest, opts ...grpc.CallOption) (*ListAnimalReply, error)
	KillAnimal(ctx context.Context, in *KillAnimalRequest, opts ...grpc.CallOption) (*KillAnimalReply, error)
	FeedingAnimal(ctx context.Context, in *FeedingAnimalRequest, opts ...grpc.CallOption) (*FeedingAnimalReply, error)
	ComposeAnimal(ctx context.Context, in *ComposeAnimalRequest, opts ...grpc.CallOption) (*ComposeAnimalReply, error)
}

type animalClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimalClient(cc grpc.ClientConnInterface) AnimalClient {
	return &animalClient{cc}
}

func (c *animalClient) CreateAnimal(ctx context.Context, in *CreateAnimalRequest, opts ...grpc.CallOption) (*CreateAnimalReply, error) {
	out := new(CreateAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/CreateAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) UpdateAnimal(ctx context.Context, in *UpdateAnimalRequest, opts ...grpc.CallOption) (*UpdateAnimalReply, error) {
	out := new(UpdateAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/UpdateAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) DeleteAnimal(ctx context.Context, in *DeleteAnimalRequest, opts ...grpc.CallOption) (*DeleteAnimalReply, error) {
	out := new(DeleteAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/DeleteAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) GetAnimal(ctx context.Context, in *GetAnimalRequest, opts ...grpc.CallOption) (*GetAnimalReply, error) {
	out := new(GetAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/GetAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) ListAnimal(ctx context.Context, in *ListAnimalRequest, opts ...grpc.CallOption) (*ListAnimalReply, error) {
	out := new(ListAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/ListAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) KillAnimal(ctx context.Context, in *KillAnimalRequest, opts ...grpc.CallOption) (*KillAnimalReply, error) {
	out := new(KillAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/KillAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) FeedingAnimal(ctx context.Context, in *FeedingAnimalRequest, opts ...grpc.CallOption) (*FeedingAnimalReply, error) {
	out := new(FeedingAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/FeedingAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalClient) ComposeAnimal(ctx context.Context, in *ComposeAnimalRequest, opts ...grpc.CallOption) (*ComposeAnimalReply, error) {
	out := new(ComposeAnimalReply)
	err := c.cc.Invoke(ctx, "/api.pledge.v1.Animal/ComposeAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimalServer is the server API for Animal service.
// All implementations must embed UnimplementedAnimalServer
// for forward compatibility
type AnimalServer interface {
	CreateAnimal(context.Context, *CreateAnimalRequest) (*CreateAnimalReply, error)
	UpdateAnimal(context.Context, *UpdateAnimalRequest) (*UpdateAnimalReply, error)
	DeleteAnimal(context.Context, *DeleteAnimalRequest) (*DeleteAnimalReply, error)
	GetAnimal(context.Context, *GetAnimalRequest) (*GetAnimalReply, error)
	ListAnimal(context.Context, *ListAnimalRequest) (*ListAnimalReply, error)
	KillAnimal(context.Context, *KillAnimalRequest) (*KillAnimalReply, error)
	FeedingAnimal(context.Context, *FeedingAnimalRequest) (*FeedingAnimalReply, error)
	ComposeAnimal(context.Context, *ComposeAnimalRequest) (*ComposeAnimalReply, error)
	mustEmbedUnimplementedAnimalServer()
}

// UnimplementedAnimalServer must be embedded to have forward compatible implementations.
type UnimplementedAnimalServer struct {
}

func (UnimplementedAnimalServer) CreateAnimal(context.Context, *CreateAnimalRequest) (*CreateAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnimal not implemented")
}
func (UnimplementedAnimalServer) UpdateAnimal(context.Context, *UpdateAnimalRequest) (*UpdateAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnimal not implemented")
}
func (UnimplementedAnimalServer) DeleteAnimal(context.Context, *DeleteAnimalRequest) (*DeleteAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnimal not implemented")
}
func (UnimplementedAnimalServer) GetAnimal(context.Context, *GetAnimalRequest) (*GetAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimal not implemented")
}
func (UnimplementedAnimalServer) ListAnimal(context.Context, *ListAnimalRequest) (*ListAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnimal not implemented")
}
func (UnimplementedAnimalServer) KillAnimal(context.Context, *KillAnimalRequest) (*KillAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillAnimal not implemented")
}
func (UnimplementedAnimalServer) FeedingAnimal(context.Context, *FeedingAnimalRequest) (*FeedingAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedingAnimal not implemented")
}
func (UnimplementedAnimalServer) ComposeAnimal(context.Context, *ComposeAnimalRequest) (*ComposeAnimalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeAnimal not implemented")
}
func (UnimplementedAnimalServer) mustEmbedUnimplementedAnimalServer() {}

// UnsafeAnimalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimalServer will
// result in compilation errors.
type UnsafeAnimalServer interface {
	mustEmbedUnimplementedAnimalServer()
}

func RegisterAnimalServer(s grpc.ServiceRegistrar, srv AnimalServer) {
	s.RegisterService(&Animal_ServiceDesc, srv)
}

func _Animal_CreateAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).CreateAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/CreateAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).CreateAnimal(ctx, req.(*CreateAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_UpdateAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).UpdateAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/UpdateAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).UpdateAnimal(ctx, req.(*UpdateAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_DeleteAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).DeleteAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/DeleteAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).DeleteAnimal(ctx, req.(*DeleteAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_GetAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).GetAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/GetAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).GetAnimal(ctx, req.(*GetAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_ListAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).ListAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/ListAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).ListAnimal(ctx, req.(*ListAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_KillAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).KillAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/KillAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).KillAnimal(ctx, req.(*KillAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_FeedingAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedingAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).FeedingAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/FeedingAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).FeedingAnimal(ctx, req.(*FeedingAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Animal_ComposeAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServer).ComposeAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pledge.v1.Animal/ComposeAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServer).ComposeAnimal(ctx, req.(*ComposeAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Animal_ServiceDesc is the grpc.ServiceDesc for Animal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Animal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.pledge.v1.Animal",
	HandlerType: (*AnimalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnimal",
			Handler:    _Animal_CreateAnimal_Handler,
		},
		{
			MethodName: "UpdateAnimal",
			Handler:    _Animal_UpdateAnimal_Handler,
		},
		{
			MethodName: "DeleteAnimal",
			Handler:    _Animal_DeleteAnimal_Handler,
		},
		{
			MethodName: "GetAnimal",
			Handler:    _Animal_GetAnimal_Handler,
		},
		{
			MethodName: "ListAnimal",
			Handler:    _Animal_ListAnimal_Handler,
		},
		{
			MethodName: "KillAnimal",
			Handler:    _Animal_KillAnimal_Handler,
		},
		{
			MethodName: "FeedingAnimal",
			Handler:    _Animal_FeedingAnimal_Handler,
		},
		{
			MethodName: "ComposeAnimal",
			Handler:    _Animal_ComposeAnimal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pledge/v1/animal.proto",
}
