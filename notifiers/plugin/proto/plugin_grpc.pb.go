// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifierClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NameResponse, error)
	Id(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IdResponse, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	NotifySubscriber(ctx context.Context, in *NotifySubscriberRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	MetadataFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMetadataField, error)
	PreCheck(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/types.Notifier/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/types.Notifier/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) Id(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/types.Notifier/Id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/types.Notifier/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) NotifySubscriber(ctx context.Context, in *NotifySubscriberRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/types.Notifier/NotifySubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) MetadataFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMetadataField, error) {
	out := new(ListMetadataField)
	err := c.cc.Invoke(ctx, "/types.Notifier/MetadataFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) PreCheck(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/types.Notifier/PreCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
// All implementations must embed UnimplementedNotifierServer
// for forward compatibility
type NotifierServer interface {
	Init(context.Context, *InitRequest) (*emptypb.Empty, error)
	Name(context.Context, *emptypb.Empty) (*NameResponse, error)
	Id(context.Context, *emptypb.Empty) (*IdResponse, error)
	Notify(context.Context, *NotifyRequest) (*ErrorResponse, error)
	NotifySubscriber(context.Context, *NotifySubscriberRequest) (*ErrorResponse, error)
	MetadataFields(context.Context, *emptypb.Empty) (*ListMetadataField, error)
	PreCheck(context.Context, *NotifyRequest) (*ErrorResponse, error)
	mustEmbedUnimplementedNotifierServer()
}

// UnimplementedNotifierServer must be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (UnimplementedNotifierServer) Init(context.Context, *InitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedNotifierServer) Name(context.Context, *emptypb.Empty) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedNotifierServer) Id(context.Context, *emptypb.Empty) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Id not implemented")
}
func (UnimplementedNotifierServer) Notify(context.Context, *NotifyRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNotifierServer) NotifySubscriber(context.Context, *NotifySubscriberRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifySubscriber not implemented")
}
func (UnimplementedNotifierServer) MetadataFields(context.Context, *emptypb.Empty) (*ListMetadataField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MetadataFields not implemented")
}
func (UnimplementedNotifierServer) PreCheck(context.Context, *NotifyRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreCheck not implemented")
}
func (UnimplementedNotifierServer) mustEmbedUnimplementedNotifierServer() {}

// UnsafeNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifierServer will
// result in compilation errors.
type UnsafeNotifierServer interface {
	mustEmbedUnimplementedNotifierServer()
}

func RegisterNotifierServer(s grpc.ServiceRegistrar, srv NotifierServer) {
	s.RegisterService(&Notifier_ServiceDesc, srv)
}

func _Notifier_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Name(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_Id_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Id(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/Id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Id(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_NotifySubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifySubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifySubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/NotifySubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifySubscriber(ctx, req.(*NotifySubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_MetadataFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).MetadataFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/MetadataFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).MetadataFields(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_PreCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).PreCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Notifier/PreCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).PreCheck(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifier_ServiceDesc is the grpc.ServiceDesc for Notifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "types.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Notifier_Init_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _Notifier_Name_Handler,
		},
		{
			MethodName: "Id",
			Handler:    _Notifier_Id_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Notifier_Notify_Handler,
		},
		{
			MethodName: "NotifySubscriber",
			Handler:    _Notifier_NotifySubscriber_Handler,
		},
		{
			MethodName: "MetadataFields",
			Handler:    _Notifier_MetadataFields_Handler,
		},
		{
			MethodName: "PreCheck",
			Handler:    _Notifier_PreCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}
