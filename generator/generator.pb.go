// Code generated by protoc-gen-go.
// source: generator.proto
// DO NOT EDIT!

/*
Package generator is a generated protocol buffer package.

It is generated from these files:
	generator.proto

It has these top-level messages:
	RequireUuidCount
	UuidString
*/
package generator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the uuid
type RequireUuidCount struct {
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *RequireUuidCount) Reset()                    { *m = RequireUuidCount{} }
func (m *RequireUuidCount) String() string            { return proto.CompactTextString(m) }
func (*RequireUuidCount) ProtoMessage()               {}
func (*RequireUuidCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequireUuidCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The response message containing the uuid list
type UuidString struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UuidString) Reset()                    { *m = UuidString{} }
func (m *UuidString) String() string            { return proto.CompactTextString(m) }
func (*UuidString) ProtoMessage()               {}
func (*UuidString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UuidString) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RequireUuidCount)(nil), "generator.RequireUuidCount")
	proto.RegisterType((*UuidString)(nil), "generator.UuidString")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// generate UUIds
	GenerateUUIDs(ctx context.Context, in *RequireUuidCount, opts ...grpc.CallOption) (*UuidString, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GenerateUUIDs(ctx context.Context, in *RequireUuidCount, opts ...grpc.CallOption) (*UuidString, error) {
	out := new(UuidString)
	err := grpc.Invoke(ctx, "/generator.Greeter/GenerateUUIDs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// generate UUIds
	GenerateUUIDs(context.Context, *RequireUuidCount) (*UuidString, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GenerateUUIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequireUuidCount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GenerateUUIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generator.Greeter/GenerateUUIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GenerateUUIDs(ctx, req.(*RequireUuidCount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generator.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateUUIDs",
			Handler:    _Greeter_GenerateUUIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generator.proto",
}

func init() { proto.RegisterFile("generator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0xcd, 0x4b,
	0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x69, 0x70, 0x09, 0x04, 0xa5, 0x16, 0x96, 0x66, 0x16, 0xa5, 0x86, 0x96, 0x66, 0xa6, 0x38, 0xe7,
	0x97, 0xe6, 0x95, 0x08, 0x89, 0x70, 0xb1, 0x26, 0x83, 0x18, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac,
	0x41, 0x10, 0x8e, 0x92, 0x1a, 0x17, 0x17, 0x48, 0x49, 0x70, 0x49, 0x51, 0x66, 0x5e, 0xba, 0x90,
	0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0x2a, 0x58, 0x15, 0x67, 0x10, 0x8c, 0x6b,
	0x14, 0xc0, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0xe4, 0xca, 0xc5, 0xeb, 0x0e,
	0xb1, 0x29, 0x35, 0x34, 0xd4, 0xd3, 0xa5, 0x58, 0x48, 0x5a, 0x0f, 0xe1, 0x14, 0x74, 0x6b, 0xa5,
	0x44, 0x91, 0x24, 0x11, 0x36, 0x29, 0x31, 0x38, 0x99, 0x73, 0x29, 0x67, 0xe6, 0xeb, 0xa5, 0x17,
	0x15, 0x24, 0xeb, 0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x16, 0xeb, 0xa5, 0x65, 0xe6, 0x25,
	0xe6, 0x14, 0x14, 0xe5, 0x67, 0xa5, 0x26, 0x97, 0x38, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x38, 0xf1,
	0xb9, 0xc3, 0xb4, 0x07, 0x80, 0x7c, 0x19, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0xae, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x98, 0xcf, 0x35, 0x01, 0x01, 0x00, 0x00,
}
