// Code generated by protoc-gen-go. DO NOT EDIT.
// source: function.proto

/*
Package function is a generated protocol buffer package.

It is generated from these files:
	function.proto

It has these top-level messages:
*/
package function

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fntypes "github.com/projectriff/function-sidecar/pkg/dispatcher/grpc/fntypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageFunction service

type MessageFunctionClient interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (MessageFunction_CallClient, error)
}

type messageFunctionClient struct {
	cc *grpc.ClientConn
}

func NewMessageFunctionClient(cc *grpc.ClientConn) MessageFunctionClient {
	return &messageFunctionClient{cc}
}

func (c *messageFunctionClient) Call(ctx context.Context, opts ...grpc.CallOption) (MessageFunction_CallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageFunction_serviceDesc.Streams[0], c.cc, "/function.MessageFunction/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageFunctionCallClient{stream}
	return x, nil
}

type MessageFunction_CallClient interface {
	Send(*fntypes.Message) error
	Recv() (*fntypes.Message, error)
	grpc.ClientStream
}

type messageFunctionCallClient struct {
	grpc.ClientStream
}

func (x *messageFunctionCallClient) Send(m *fntypes.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageFunctionCallClient) Recv() (*fntypes.Message, error) {
	m := new(fntypes.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MessageFunction service

type MessageFunctionServer interface {
	Call(MessageFunction_CallServer) error
}

func RegisterMessageFunctionServer(s *grpc.Server, srv MessageFunctionServer) {
	s.RegisterService(&_MessageFunction_serviceDesc, srv)
}

func _MessageFunction_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageFunctionServer).Call(&messageFunctionCallServer{stream})
}

type MessageFunction_CallServer interface {
	Send(*fntypes.Message) error
	Recv() (*fntypes.Message, error)
	grpc.ServerStream
}

type messageFunctionCallServer struct {
	grpc.ServerStream
}

func (x *messageFunctionCallServer) Send(m *fntypes.Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageFunctionCallServer) Recv() (*fntypes.Message, error) {
	m := new(fntypes.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageFunction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "function.MessageFunction",
	HandlerType: (*MessageFunctionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _MessageFunction_Call_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "function.proto",
}

func init() { proto.RegisterFile("function.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2b, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x78,
	0xd3, 0xf2, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x21, 0x12, 0x46, 0xce, 0x5c, 0xfc, 0xbe, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x6e, 0x50, 0x15, 0x42, 0x06, 0x5c, 0x2c, 0xce, 0x89, 0x39, 0x39, 0x42,
	0x02, 0x7a, 0x30, 0xa5, 0x50, 0x15, 0x52, 0x18, 0x22, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49,
	0x6c, 0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x92, 0x1a, 0x63, 0x76, 0x00,
	0x00, 0x00,
}
