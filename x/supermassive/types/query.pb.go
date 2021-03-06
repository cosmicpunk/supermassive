// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supermassive/query.proto

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

func init() { proto.RegisterFile("supermassive/query.proto", fileDescriptor_8c3353054aaa7115) }

var fileDescriptor_8c3353054aaa7115 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x86, 0x61, 0x50, 0x13, 0x46, 0x27, 0x43, 0x4c, 0x27, 0x27, 0x07, 0x2e, 0xe8, 0xe0, 0xee,
	0x1b, 0xb0, 0xba, 0xb5, 0xa4, 0xa9, 0x8d, 0xd2, 0xab, 0x5c, 0x4b, 0xe4, 0x2d, 0x7c, 0x2c, 0x47,
	0x46, 0x47, 0x03, 0x2f, 0x62, 0x00, 0x13, 0x65, 0xfc, 0x92, 0x2f, 0xdf, 0xdd, 0x1f, 0xad, 0xc8,
	0x5b, 0x59, 0x16, 0x9c, 0x48, 0x57, 0x12, 0x6e, 0x5e, 0x96, 0x75, 0x62, 0x4b, 0x74, 0xb8, 0xdc,
	0xe4, 0x48, 0x85, 0xce, 0xad, 0x37, 0x97, 0xe4, 0x5f, 0x9a, 0x40, 0xbc, 0x56, 0x88, 0xea, 0x2a,
	0x81, 0x5b, 0x0d, 0xdc, 0x18, 0x74, 0xdc, 0x69, 0x34, 0x34, 0x36, 0xe2, 0x6d, 0xdf, 0x40, 0x02,
	0xc1, 0xe9, 0x1b, 0x87, 0x2a, 0x15, 0xd2, 0xf1, 0x14, 0x2c, 0x57, 0xda, 0x0c, 0xf2, 0xe8, 0xee,
	0x16, 0xd1, 0x2c, 0xeb, 0x8d, 0x63, 0xf6, 0x6c, 0x59, 0xd8, 0xb4, 0x2c, 0x7c, 0xb7, 0x2c, 0x7c,
	0x74, 0x2c, 0x68, 0x3a, 0x16, 0xbc, 0x3a, 0x16, 0x9c, 0x0e, 0x4a, 0xbb, 0xb3, 0x17, 0x49, 0x8e,
	0x05, 0xfc, 0xbe, 0x83, 0xc9, 0x84, 0xfb, 0x14, 0x5d, 0x6d, 0x25, 0x89, 0xf9, 0x70, 0x62, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x15, 0xb7, 0xd6, 0xee, 0x00, 0x00, 0x00,
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
	ServiceName: "cosmicpunk.supermassive.supermassive.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "supermassive/query.proto",
}
