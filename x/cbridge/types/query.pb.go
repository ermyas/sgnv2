// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cbridge/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cbridge/query.proto", fileDescriptor_810cf729730a14ee) }

var fileDescriptor_810cf729730a14ee = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x31, 0x6e, 0x83, 0x30,
	0x18, 0x05, 0x60, 0x18, 0x5a, 0x24, 0xc6, 0x76, 0x43, 0x95, 0x0f, 0x50, 0x09, 0x7e, 0x41, 0x6f,
	0xd0, 0x2d, 0x63, 0xd6, 0x6c, 0xb6, 0xf3, 0xcb, 0xb1, 0x02, 0xfe, 0x1d, 0xdb, 0x90, 0x70, 0x8b,
	0x1c, 0x2b, 0x23, 0x63, 0xc6, 0x08, 0x2e, 0x12, 0x05, 0xc8, 0xfa, 0xf4, 0xbd, 0xa7, 0x97, 0x7e,
	0x4b, 0xe1, 0xf4, 0x5e, 0x21, 0x9c, 0x5a, 0x74, 0x7d, 0x61, 0x1d, 0x05, 0xfa, 0x4a, 0xd6, 0x30,
	0xfb, 0x51, 0x44, 0xaa, 0x46, 0xe0, 0x56, 0x03, 0x37, 0x86, 0x02, 0x0f, 0x9a, 0x8c, 0x5f, 0x58,
	0xf6, 0x2b, 0xc9, 0x37, 0xe4, 0x41, 0x70, 0xbf, 0xf6, 0xa1, 0x2b, 0x05, 0x06, 0x5e, 0x82, 0xe5,
	0x4a, 0x9b, 0x19, 0x2f, 0xb6, 0x4a, 0xd2, 0x8f, 0xed, 0x4b, 0xfc, 0x6f, 0x6e, 0x23, 0x8b, 0x87,
	0x91, 0xc5, 0x8f, 0x91, 0xc5, 0xd7, 0x89, 0x45, 0xc3, 0xc4, 0xa2, 0xfb, 0xc4, 0xa2, 0x1d, 0x28,
	0x1d, 0x0e, 0xad, 0x28, 0x24, 0x35, 0x20, 0xb1, 0x46, 0x97, 0x1b, 0x0c, 0x67, 0x72, 0x47, 0xf0,
	0xca, 0xe4, 0x5d, 0x05, 0x17, 0x78, 0x9f, 0x0d, 0xbd, 0x45, 0x2f, 0x3e, 0xe7, 0xe9, 0xbf, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xa4, 0xd7, 0xf8, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cbridge.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cbridge/query.proto",
}
