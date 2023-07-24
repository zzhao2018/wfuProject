// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sumFunc.proto

package generate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SumRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumReply struct {
	V                    int64    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReply) Reset()         { *m = SumReply{} }
func (m *SumReply) String() string { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()    {}
func (*SumReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{1}
}

func (m *SumReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReply.Unmarshal(m, b)
}
func (m *SumReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReply.Marshal(b, m, deterministic)
}
func (m *SumReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReply.Merge(m, src)
}
func (m *SumReply) XXX_Size() int {
	return xxx_messageInfo_SumReply.Size(m)
}
func (m *SumReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReply.DiscardUnknown(m)
}

var xxx_messageInfo_SumReply proto.InternalMessageInfo

func (m *SumReply) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *SumReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ConcatRequest struct {
	Data1                string   `protobuf:"bytes,1,opt,name=data1,proto3" json:"data1,omitempty"`
	Data2                string   `protobuf:"bytes,2,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatRequest) Reset()         { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string { return proto.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()    {}
func (*ConcatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{2}
}

func (m *ConcatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatRequest.Unmarshal(m, b)
}
func (m *ConcatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatRequest.Marshal(b, m, deterministic)
}
func (m *ConcatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatRequest.Merge(m, src)
}
func (m *ConcatRequest) XXX_Size() int {
	return xxx_messageInfo_ConcatRequest.Size(m)
}
func (m *ConcatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatRequest proto.InternalMessageInfo

func (m *ConcatRequest) GetData1() string {
	if m != nil {
		return m.Data1
	}
	return ""
}

func (m *ConcatRequest) GetData2() string {
	if m != nil {
		return m.Data2
	}
	return ""
}

type ConcatReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatReply) Reset()         { *m = ConcatReply{} }
func (m *ConcatReply) String() string { return proto.CompactTextString(m) }
func (*ConcatReply) ProtoMessage()    {}
func (*ConcatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{3}
}

func (m *ConcatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatReply.Unmarshal(m, b)
}
func (m *ConcatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatReply.Marshal(b, m, deterministic)
}
func (m *ConcatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatReply.Merge(m, src)
}
func (m *ConcatReply) XXX_Size() int {
	return xxx_messageInfo_ConcatReply.Size(m)
}
func (m *ConcatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatReply proto.InternalMessageInfo

func (m *ConcatReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ConcatReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "generate.SumRequest")
	proto.RegisterType((*SumReply)(nil), "generate.SumReply")
	proto.RegisterType((*ConcatRequest)(nil), "generate.ConcatRequest")
	proto.RegisterType((*ConcatReply)(nil), "generate.ConcatReply")
}

func init() { proto.RegisterFile("sumFunc.proto", fileDescriptor_7fdb574872ef000e) }

var fileDescriptor_7fdb574872ef000e = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xdf, 0x4a, 0x87, 0x30,
	0x1c, 0xc5, 0x5b, 0x96, 0xe8, 0xb7, 0x84, 0x18, 0x56, 0xe2, 0x55, 0xec, 0x4a, 0xba, 0x10, 0xb4,
	0x8b, 0xa0, 0x2e, 0x83, 0x1e, 0x60, 0xf6, 0x02, 0x9b, 0x8d, 0x6e, 0xfc, 0xd7, 0xdc, 0x04, 0xdf,
	0xa8, 0xc7, 0x8c, 0x6d, 0x0e, 0x29, 0xaf, 0x7e, 0x77, 0x9e, 0xef, 0x39, 0xc7, 0x0f, 0x87, 0x41,
	0x32, 0xeb, 0xfe, 0x5d, 0x0f, 0x6d, 0x39, 0xc9, 0x51, 0x8d, 0x38, 0xfa, 0x12, 0x83, 0x90, 0x4c,
	0x09, 0x52, 0x00, 0x34, 0xba, 0xa7, 0xe2, 0x5b, 0x8b, 0x59, 0xe1, 0x6b, 0x40, 0x2c, 0x43, 0x0f,
	0xa8, 0x08, 0x28, 0x62, 0x46, 0xf1, 0xec, 0xdc, 0x29, 0x4e, 0x1e, 0x21, 0xb2, 0xc9, 0xa9, 0x5b,
	0x8d, 0xb3, 0xf8, 0xdc, 0x82, 0x6f, 0x20, 0x10, 0x52, 0xda, 0x64, 0x4c, 0xcd, 0x27, 0x79, 0x85,
	0xe4, 0x6d, 0x1c, 0x5a, 0xa6, 0xfc, 0x8f, 0x53, 0xb8, 0xfc, 0x64, 0x8a, 0x55, 0xb6, 0x14, 0x53,
	0x27, 0xfc, 0xb5, 0xde, 0xaa, 0x4e, 0x90, 0x67, 0xb8, 0xf2, 0x65, 0xc3, 0xba, 0x83, 0x50, 0x8a,
	0x59, 0x77, 0x6a, 0xeb, 0x6e, 0xea, 0x48, 0xad, 0x7f, 0x10, 0x5c, 0x7c, 0x18, 0x5a, 0x05, 0x41,
	0xa3, 0x7b, 0x9c, 0x96, 0x7e, 0x66, 0xb9, 0x6f, 0xcc, 0xf1, 0xbf, 0xeb, 0xd4, 0xad, 0xe4, 0x0c,
	0xbf, 0x40, 0xe8, 0xa0, 0xf8, 0x7e, 0xf7, 0xff, 0x6c, 0xc8, 0x6f, 0x8f, 0x86, 0xeb, 0x5a, 0x1c,
	0x3f, 0x05, 0xc7, 0x43, 0xfb, 0x0e, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa7, 0x04,
	0xbf, 0x98, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error)
	Sub(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/generate.Test/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error) {
	out := new(ConcatReply)
	err := c.cc.Invoke(ctx, "/generate.Test/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Sub(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/generate.Test/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Sum(context.Context, *SumRequest) (*SumReply, error)
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
	Sub(context.Context, *SumRequest) (*SumReply, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) Sum(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedTestServer) Concat(ctx context.Context, req *ConcatRequest) (*ConcatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (*UnimplementedTestServer) Sub(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Test/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Test/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Test/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Sub(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generate.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Test_Sum_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _Test_Concat_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Test_Sub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sumFunc.proto",
}