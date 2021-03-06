// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package platform

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/zokypesch/proto-lib/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Platform struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedBy            string               `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0}
}

func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (m *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(m, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Platform) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Platform) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type ResponseCreatePlatform struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResponseCreatePlatform) Reset()         { *m = ResponseCreatePlatform{} }
func (m *ResponseCreatePlatform) String() string { return proto.CompactTextString(m) }
func (*ResponseCreatePlatform) ProtoMessage()    {}
func (*ResponseCreatePlatform) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{1}
}

func (m *ResponseCreatePlatform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseCreatePlatform.Unmarshal(m, b)
}
func (m *ResponseCreatePlatform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseCreatePlatform.Marshal(b, m, deterministic)
}
func (m *ResponseCreatePlatform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCreatePlatform.Merge(m, src)
}
func (m *ResponseCreatePlatform) XXX_Size() int {
	return xxx_messageInfo_ResponseCreatePlatform.Size(m)
}
func (m *ResponseCreatePlatform) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCreatePlatform.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCreatePlatform proto.InternalMessageInfo

func (m *ResponseCreatePlatform) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResponseCreatePlatform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResponseCreatePlatform) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResponseCreatePlatform) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Platform)(nil), "platform.Platform")
	proto.RegisterType((*ResponseCreatePlatform)(nil), "platform.ResponseCreatePlatform")
}

func init() { proto.RegisterFile("platform.proto", fileDescriptor_918f3d50bfb447e4) }

var fileDescriptor_918f3d50bfb447e4 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x8f, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0x24, 0x77, 0xea, 0x19, 0x14, 0xc0, 0x3a, 0x55, 0x56, 0x0e, 0x44, 0x14, 0xdd,
	0x50, 0xa1, 0x23, 0x11, 0x61, 0xe2, 0x36, 0x82, 0x10, 0x2b, 0x8a, 0x98, 0x10, 0x12, 0x72, 0x1a,
	0x5f, 0x6a, 0x91, 0xc4, 0x26, 0x71, 0x91, 0xc2, 0xc8, 0xce, 0xc4, 0x3f, 0x93, 0xb9, 0x13, 0x3b,
	0x0b, 0x03, 0x03, 0x43, 0x27, 0x46, 0xfe, 0x02, 0x54, 0x3b, 0x4e, 0x5b, 0xca, 0xc0, 0x6d, 0xf6,
	0xfb, 0xbe, 0xf7, 0xde, 0xf7, 0x93, 0x1e, 0x74, 0x45, 0x49, 0xe4, 0x15, 0x6f, 0xaa, 0x50, 0x34,
	0x5c, 0x72, 0x34, 0x31, 0x7f, 0xef, 0xac, 0xe0, 0xbc, 0x28, 0x69, 0xa4, 0xea, 0xd9, 0xf2, 0x2a,
	0xa2, 0x95, 0x90, 0x9d, 0xb6, 0x79, 0x77, 0x07, 0x91, 0x08, 0x16, 0x91, 0xba, 0xe6, 0x92, 0x48,
	0xc6, 0xeb, 0x76, 0x50, 0xe3, 0x82, 0xc9, 0xc5, 0x32, 0x0b, 0xe7, 0xbc, 0x8a, 0x3e, 0xf2, 0x77,
	0x9d, 0xa0, 0xed, 0x7c, 0xa1, 0x07, 0x3d, 0x2c, 0x59, 0xa6, 0x5f, 0x11, 0x17, 0xbb, 0x3d, 0xf7,
	0xff, 0x5e, 0x27, 0x59, 0x45, 0x5b, 0x49, 0x2a, 0xa1, 0x0d, 0xc1, 0x1a, 0xc0, 0xc9, 0xcb, 0x21,
	0x1c, 0x3a, 0x87, 0x16, 0xcb, 0x31, 0xf0, 0xc1, 0xcc, 0x4e, 0x4e, 0xbf, 0xf7, 0x18, 0xfc, 0xec,
	0x31, 0xf8, 0xd5, 0xe3, 0x49, 0x43, 0xdf, 0x2f, 0x59, 0x43, 0xf3, 0xd4, 0x62, 0x39, 0x3a, 0x87,
	0x4e, 0x4d, 0x2a, 0x8a, 0x2d, 0x1f, 0xcc, 0x4e, 0x92, 0xdb, 0x07, 0x1e, 0xa5, 0xa2, 0x0b, 0x78,
	0x44, 0x2b, 0xc2, 0x4a, 0x6c, 0x2b, 0xdb, 0x74, 0xb0, 0xb9, 0xc6, 0x76, 0xa1, 0xd4, 0x54, 0x9b,
	0xd0, 0x3d, 0x08, 0xe7, 0x0d, 0x25, 0x92, 0xe6, 0x6f, 0xb3, 0x0e, 0x3b, 0x9b, 0x96, 0xf4, 0x64,
	0xa8, 0x24, 0x1d, 0x7a, 0xb2, 0x95, 0x89, 0xc4, 0x47, 0x3e, 0x98, 0xdd, 0x88, 0xbd, 0x50, 0xb3,
	0x85, 0x86, 0x2d, 0x7c, 0x65, 0xd8, 0xc6, 0xd6, 0xa7, 0xf2, 0xd2, 0xf9, 0xda, 0x63, 0x10, 0x7c,
	0x06, 0x70, 0x9a, 0xd2, 0x56, 0xf0, 0xba, 0xa5, 0xcf, 0x94, 0x36, 0x42, 0xbb, 0x5b, 0x68, 0x85,
	0x87, 0x76, 0xf1, 0x06, 0x98, 0xd3, 0x3d, 0x18, 0x13, 0x7a, 0x3f, 0x95, 0x73, 0x8d, 0x54, 0xf1,
	0x0f, 0x00, 0xc7, 0x9b, 0x40, 0x0d, 0x3c, 0xd6, 0x99, 0x10, 0x0a, 0xc7, 0xc3, 0x31, 0xf9, 0x3c,
	0x7f, 0x5b, 0xfb, 0x37, 0x41, 0x10, 0xaf, 0x7a, 0xec, 0x08, 0xde, 0xca, 0xdf, 0x3d, 0xbe, 0x65,
	0xaa, 0xa1, 0x36, 0x7d, 0xfa, 0xb6, 0xfe, 0x62, 0xdd, 0x09, 0x6e, 0x46, 0x1f, 0x1e, 0x45, 0x66,
	0xce, 0x25, 0x78, 0x80, 0xde, 0x40, 0xfb, 0x05, 0x95, 0x68, 0x7a, 0x10, 0xf7, 0xf9, 0xe6, 0x1e,
	0xff, 0x63, 0xe9, 0xd9, 0xaa, 0xc7, 0x76, 0x41, 0xa5, 0x5a, 0xe0, 0xa2, 0xbd, 0x05, 0x09, 0x7c,
	0x3d, 0xd2, 0x65, 0xc7, 0x6a, 0xf4, 0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xf7, 0xc0,
	0x5b, 0x14, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
	Create(ctx context.Context, in *Platform, opts ...grpc.CallOption) (*ResponseCreatePlatform, error)
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseCreatePlatform, error)
}

type platformClient struct {
	cc *grpc.ClientConn
}

func NewPlatformClient(cc *grpc.ClientConn) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) Create(ctx context.Context, in *Platform, opts ...grpc.CallOption) (*ResponseCreatePlatform, error) {
	out := new(ResponseCreatePlatform)
	err := c.cc.Invoke(ctx, "/platform.platform/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseCreatePlatform, error) {
	out := new(ResponseCreatePlatform)
	err := c.cc.Invoke(ctx, "/platform.platform/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
	Create(context.Context, *Platform) (*ResponseCreatePlatform, error)
	Get(context.Context, *empty.Empty) (*ResponseCreatePlatform, error)
}

// UnimplementedPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (*UnimplementedPlatformServer) Create(ctx context.Context, req *Platform) (*ResponseCreatePlatform, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPlatformServer) Get(ctx context.Context, req *empty.Empty) (*ResponseCreatePlatform, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Platform)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.platform/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Create(ctx, req.(*Platform))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.platform/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platform.platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Platform_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Platform_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}
