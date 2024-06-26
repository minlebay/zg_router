// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: message.proto

package message

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MessageRouter_ReceiveMessage_FullMethodName = "/router.MessageRouter/ReceiveMessage"
)

// MessageRouterClient is the client API for MessageRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageRouterClient interface {
	ReceiveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type messageRouterClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageRouterClient(cc grpc.ClientConnInterface) MessageRouterClient {
	return &messageRouterClient{cc}
}

func (c *messageRouterClient) ReceiveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MessageRouter_ReceiveMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageRouterServer is the server API for MessageRouter service.
// All implementations must embed UnimplementedMessageRouterServer
// for forward compatibility
type MessageRouterServer interface {
	ReceiveMessage(context.Context, *Message) (*Response, error)
	mustEmbedUnimplementedMessageRouterServer()
}

// UnimplementedMessageRouterServer must be embedded to have forward compatible implementations.
type UnimplementedMessageRouterServer struct {
}

func (UnimplementedMessageRouterServer) ReceiveMessage(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedMessageRouterServer) mustEmbedUnimplementedMessageRouterServer() {}

// UnsafeMessageRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageRouterServer will
// result in compilation errors.
type UnsafeMessageRouterServer interface {
	mustEmbedUnimplementedMessageRouterServer()
}

func RegisterMessageRouterServer(s grpc.ServiceRegistrar, srv MessageRouterServer) {
	s.RegisterService(&MessageRouter_ServiceDesc, srv)
}

func _MessageRouter_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageRouterServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageRouter_ReceiveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageRouterServer).ReceiveMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageRouter_ServiceDesc is the grpc.ServiceDesc for MessageRouter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageRouter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "router.MessageRouter",
	HandlerType: (*MessageRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveMessage",
			Handler:    _MessageRouter_ReceiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
