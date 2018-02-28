// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compress.proto

/*
Package compress is a generated protocol buffer package.

It is generated from these files:
	compress.proto

It has these top-level messages:
	CompressRequest
	CompressReply
*/
package compress

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

// 压缩类型
type Type int32

const (
	// GZIP BROTLI用于文本的压缩
	Type_GZIP    Type = 0
	Type_BROTLI  Type = 1
	Type_JPEG    Type = 2
	Type_PNG     Type = 3
	Type_WEBP    Type = 4
	Type_GUETZLI Type = 5
)

var Type_name = map[int32]string{
	0: "GZIP",
	1: "BROTLI",
	2: "JPEG",
	3: "PNG",
	4: "WEBP",
	5: "GUETZLI",
}
var Type_value = map[string]int32{
	"GZIP":    0,
	"BROTLI":  1,
	"JPEG":    2,
	"PNG":     3,
	"WEBP":    4,
	"GUETZLI": 5,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The request message containing the user's name.
type CompressRequest struct {
	// 压缩类型
	Type Type `protobuf:"varint,1,opt,name=type,enum=compress.Type" json:"type,omitempty"`
	// 数据
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// 压缩质量
	Quality uint32 `protobuf:"varint,7,opt,name=quality" json:"quality,omitempty"`
	// 图片类型
	ImageType Type `protobuf:"varint,8,opt,name=imageType,enum=compress.Type" json:"imageType,omitempty"`
	// 图片宽度
	Width uint32 `protobuf:"varint,9,opt,name=width" json:"width,omitempty"`
	// 图片高度
	Height uint32 `protobuf:"varint,10,opt,name=height" json:"height,omitempty"`
}

func (m *CompressRequest) Reset()                    { *m = CompressRequest{} }
func (m *CompressRequest) String() string            { return proto.CompactTextString(m) }
func (*CompressRequest) ProtoMessage()               {}
func (*CompressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompressRequest) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_GZIP
}

func (m *CompressRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CompressRequest) GetQuality() uint32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *CompressRequest) GetImageType() Type {
	if m != nil {
		return m.ImageType
	}
	return Type_GZIP
}

func (m *CompressRequest) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *CompressRequest) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

// The response message containing the greetings
type CompressReply struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CompressReply) Reset()                    { *m = CompressReply{} }
func (m *CompressReply) String() string            { return proto.CompactTextString(m) }
func (*CompressReply) ProtoMessage()               {}
func (*CompressReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CompressReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CompressRequest)(nil), "compress.CompressRequest")
	proto.RegisterType((*CompressReply)(nil), "compress.CompressReply")
	proto.RegisterEnum("compress.Type", Type_name, Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Compress service

type CompressClient interface {
	// do compress
	Do(ctx context.Context, in *CompressRequest, opts ...grpc.CallOption) (*CompressReply, error)
}

type compressClient struct {
	cc *grpc.ClientConn
}

func NewCompressClient(cc *grpc.ClientConn) CompressClient {
	return &compressClient{cc}
}

func (c *compressClient) Do(ctx context.Context, in *CompressRequest, opts ...grpc.CallOption) (*CompressReply, error) {
	out := new(CompressReply)
	err := grpc.Invoke(ctx, "/compress.Compress/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Compress service

type CompressServer interface {
	// do compress
	Do(context.Context, *CompressRequest) (*CompressReply, error)
}

func RegisterCompressServer(s *grpc.Server, srv CompressServer) {
	s.RegisterService(&_Compress_serviceDesc, srv)
}

func _Compress_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompressServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compress.Compress/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompressServer).Do(ctx, req.(*CompressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "compress.Compress",
	HandlerType: (*CompressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Compress_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compress.proto",
}

func init() { proto.RegisterFile("compress.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0x59, 0x28, 0xb4, 0xcc, 0xff, 0x0f, 0x6e, 0x26, 0x46, 0x57, 0x4f, 0xa4, 0x5c, 0x88,
	0x31, 0x1c, 0xf0, 0xe6, 0xb1, 0x5a, 0x9a, 0x1a, 0xa2, 0x9b, 0xa6, 0xc6, 0x84, 0x5b, 0x85, 0x0d,
	0x6d, 0x52, 0xec, 0x42, 0x17, 0x4d, 0x3f, 0x9d, 0x5f, 0xcd, 0x74, 0xa5, 0x6d, 0x62, 0xb8, 0xed,
	0xef, 0xcd, 0xec, 0xdb, 0x79, 0x3b, 0x30, 0x5c, 0x65, 0x5b, 0xb9, 0x17, 0x79, 0x3e, 0x95, 0xfb,
	0x4c, 0x65, 0x68, 0x55, 0x6c, 0x7f, 0x13, 0x38, 0x7b, 0x38, 0x42, 0x20, 0x76, 0x07, 0x91, 0x2b,
	0xb4, 0xc1, 0x50, 0x85, 0x14, 0x8c, 0x8c, 0xc8, 0x64, 0x38, 0x1b, 0x4e, 0xeb, 0xcb, 0x61, 0x21,
	0x45, 0xa0, 0x6b, 0x88, 0x60, 0xac, 0x23, 0x15, 0xb1, 0xf6, 0x88, 0x4c, 0xfe, 0x07, 0xfa, 0x8c,
	0x0c, 0xcc, 0xdd, 0x21, 0x4a, 0x13, 0x55, 0x30, 0x73, 0x44, 0x26, 0x83, 0xa0, 0x42, 0xbc, 0x85,
	0x7e, 0xb2, 0x8d, 0x36, 0xa2, 0x34, 0x60, 0xd6, 0x49, 0xdb, 0xa6, 0x01, 0xcf, 0xa1, 0xfb, 0x95,
	0xac, 0x55, 0xcc, 0xfa, 0xda, 0xe5, 0x17, 0xf0, 0x02, 0x7a, 0xb1, 0x48, 0x36, 0xb1, 0x62, 0xa0,
	0xe5, 0x23, 0xd9, 0x63, 0x18, 0x34, 0x01, 0x64, 0x5a, 0xd4, 0xa3, 0x91, 0x66, 0xb4, 0x9b, 0x39,
	0x18, 0xda, 0xda, 0x02, 0xc3, 0x5b, 0xfa, 0x9c, 0xb6, 0x10, 0xa0, 0xe7, 0x04, 0x2f, 0xe1, 0xc2,
	0xa7, 0xa4, 0x54, 0x9f, 0xb8, 0xeb, 0xd1, 0x36, 0x9a, 0xd0, 0xe1, 0xcf, 0x1e, 0xed, 0x94, 0xd2,
	0x9b, 0xeb, 0x70, 0x6a, 0xe0, 0x3f, 0x30, 0xbd, 0x57, 0x37, 0x5c, 0x2e, 0x7c, 0xda, 0x9d, 0xcd,
	0xc1, 0xaa, 0x1e, 0xc3, 0x7b, 0x68, 0x3f, 0x66, 0x78, 0xd5, 0xe4, 0xf8, 0xf3, 0x8f, 0xd7, 0x97,
	0xa7, 0x4a, 0x32, 0x2d, 0xec, 0x96, 0x33, 0x06, 0xfa, 0x99, 0xac, 0xa2, 0x8f, 0x3c, 0xab, 0x7b,
	0x9c, 0x3a, 0x06, 0x2f, 0x77, 0xc4, 0xc9, 0x7b, 0x4f, 0x2f, 0xeb, 0xee, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x62, 0x8c, 0x16, 0x1e, 0xbe, 0x01, 0x00, 0x00,
}
