// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

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

type Request struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_91b5335d5b15eb6c, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Response struct {
	Result               int64    `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_91b5335d5b15eb6c, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	GetMax(ctx context.Context, opts ...grpc.CallOption) (TestService_GetMaxClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetMax(ctx context.Context, opts ...grpc.CallOption) (TestService_GetMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/proto.TestService/GetMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceGetMaxClient{stream}
	return x, nil
}

type TestService_GetMaxClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceGetMaxClient struct {
	grpc.ClientStream
}

func (x *testServiceGetMaxClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceGetMaxClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	GetMax(TestService_GetMaxServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_GetMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).GetMax(&testServiceGetMaxServer{stream})
}

type TestService_GetMaxServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type testServiceGetMaxServer struct {
	grpc.ServerStream
}

func (x *testServiceGetMaxServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceGetMaxServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMax",
			Handler:       _TestService_GetMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_91b5335d5b15eb6c) }

var fileDescriptor_service_91b5335d5b15eb6c = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c,
	0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa5, 0xb9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0xa6, 0x92, 0x12, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x89, 0x04, 0x13,
	0x58, 0x01, 0x94, 0x67, 0x64, 0xc3, 0xc5, 0x1d, 0x92, 0x5a, 0x5c, 0x12, 0x0c, 0x31, 0x5c, 0x48,
	0x97, 0x8b, 0xcd, 0x3d, 0xb5, 0xc4, 0x37, 0xb1, 0x42, 0x88, 0x0f, 0x62, 0x91, 0x1e, 0xd4, 0x78,
	0x29, 0x7e, 0x38, 0x1f, 0x62, 0xa2, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0x5a, 0x59, 0x8c, 0x9d, 0x00, 0x00, 0x00,
}
