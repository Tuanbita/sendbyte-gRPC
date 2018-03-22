// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto.proto

It has these top-level messages:
	Request
	Chunk
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Req string `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto1.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto1.RegisterType((*Request)(nil), "Request")
	proto1.RegisterType((*Chunk)(nil), "Chunk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Chicken2(ctx context.Context, in *Request, opts ...grpc.CallOption) (Greeter_Chicken2Client, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Chicken2(ctx context.Context, in *Request, opts ...grpc.CallOption) (Greeter_Chicken2Client, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/Greeter/Chicken2", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterChicken2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_Chicken2Client interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type greeterChicken2Client struct {
	grpc.ClientStream
}

func (x *greeterChicken2Client) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Chicken2(*Request, Greeter_Chicken2Server) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Chicken2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).Chicken2(m, &greeterChicken2Server{stream})
}

type Greeter_Chicken2Server interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type greeterChicken2Server struct {
	grpc.ServerStream
}

func (x *greeterChicken2Server) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chicken2",
			Handler:       _Greeter_Chicken2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto.proto",
}

func init() { proto1.RegisterFile("proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x4a, 0xd2, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x02, 0x5c, 0xcc, 0x45, 0xa9, 0x85, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92,
	0x2c, 0x17, 0xab, 0x73, 0x46, 0x69, 0x5e, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x32, 0x88, 0x01, 0x96,
	0xe4, 0x09, 0x82, 0x70, 0x8c, 0xb4, 0xb9, 0xd8, 0xdd, 0x8b, 0x52, 0x53, 0x4b, 0x52, 0x8b, 0x84,
	0x14, 0xb8, 0x38, 0x9c, 0x33, 0x32, 0x93, 0xb3, 0x53, 0xf3, 0x8c, 0x84, 0x38, 0xf4, 0xa0, 0x26,
	0x4a, 0xb1, 0xe9, 0x81, 0xb5, 0x2b, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81, 0xed, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0x49, 0xb0, 0x49, 0x7e, 0x00, 0x00, 0x00,
}