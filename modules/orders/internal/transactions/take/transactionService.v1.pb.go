// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/take/transactionService.v1.proto

package take

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("modules/orders/internal/transactions/take/transactionService.v1.proto", fileDescriptor_ec40104d843f54d3)
}

var fileDescriptor_ec40104d843f54d3 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x3d, 0x4b, 0xc3, 0x40,
	0x18, 0x6e, 0xa2, 0x38, 0xc4, 0x2d, 0x38, 0xd4, 0x50, 0x02, 0x16, 0x57, 0xef, 0xa8, 0x6e, 0xd9,
	0x2c, 0xf8, 0xb1, 0x94, 0x16, 0x3d, 0x1c, 0xa4, 0xcb, 0xdb, 0xf4, 0x25, 0x86, 0x26, 0x77, 0x25,
	0xf7, 0xb6, 0x3f, 0x40, 0xff, 0x80, 0xe0, 0xec, 0xe2, 0xe8, 0x2f, 0x11, 0xa7, 0x82, 0x8b, 0xa3,
	0xa4, 0x4e, 0xfe, 0x0a, 0x49, 0xaf, 0xd2, 0x40, 0x09, 0x14, 0xd7, 0x7b, 0x9e, 0xe7, 0x9e, 0x8f,
	0xd7, 0x39, 0x4b, 0xd5, 0x70, 0x92, 0xa0, 0xe6, 0x2a, 0x1b, 0x62, 0xa6, 0x79, 0x2c, 0x09, 0x33,
	0x09, 0x09, 0xa7, 0x0c, 0xa4, 0x86, 0x90, 0x62, 0x25, 0x35, 0x27, 0x18, 0x61, 0xf9, 0xe5, 0x1a,
	0xb3, 0x69, 0x1c, 0x22, 0x9b, 0xb6, 0xd8, 0x38, 0x53, 0xa4, 0xdc, 0xba, 0x91, 0xb3, 0xb2, 0x8a,
	0x15, 0x2a, 0xaf, 0x11, 0x29, 0x15, 0x25, 0xc8, 0x61, 0x1c, 0x73, 0x90, 0x52, 0x11, 0x18, 0x70,
	0xa1, 0xf3, 0x82, 0xcd, 0xed, 0x53, 0xd4, 0x1a, 0xa2, 0x95, 0xa7, 0x77, 0xfe, 0xaf, 0xe8, 0x57,
	0xa8, 0xc7, 0x4a, 0xea, 0xd5, 0x3f, 0xc7, 0x0f, 0x96, 0xb3, 0x2b, 0x56, 0x04, 0x97, 0x9c, 0x6d,
	0x01, 0x23, 0x74, 0x0f, 0x58, 0x55, 0x29, 0xd6, 0x31, 0x59, 0xbc, 0xa3, 0x6a, 0x8a, 0x58, 0xb7,
	0x6c, 0x7a, 0xf7, 0x1f, 0xdf, 0x4f, 0xf6, 0x5e, 0xd3, 0xe5, 0x29, 0x48, 0x4a, 0x8a, 0x4e, 0x04,
	0x26, 0x62, 0xfb, 0xd9, 0x7e, 0xcb, 0x7d, 0x6b, 0x96, 0xfb, 0xd6, 0x57, 0xee, 0x5b, 0x8f, 0x73,
	0xbf, 0x36, 0x9b, 0xfb, 0xb5, 0xcf, 0xb9, 0x5f, 0x73, 0x1a, 0xa1, 0x4a, 0x2b, 0x8d, 0xda, 0xfb,
	0x62, 0xed, 0x2e, 0x37, 0xad, 0x5e, 0xd1, 0xac, 0x67, 0xdd, 0x5e, 0x46, 0x31, 0xdd, 0x4d, 0x06,
	0x2c, 0x54, 0x29, 0x3f, 0xd5, 0x1a, 0xa9, 0xb3, 0x34, 0x5e, 0x4e, 0xb7, 0xf1, 0x84, 0x2f, 0xf6,
	0x56, 0x57, 0x88, 0x57, 0xbb, 0xde, 0x35, 0x39, 0x44, 0x39, 0x47, 0x31, 0xd7, 0xfb, 0x1f, 0xd4,
	0x2f, 0x43, 0xfd, 0x02, 0xca, 0xed, 0xc3, 0x2a, 0xa8, 0x7f, 0xd1, 0x6b, 0x77, 0x90, 0x60, 0x08,
	0x04, 0x3f, 0xb6, 0x67, 0x68, 0x41, 0x50, 0xe6, 0x05, 0x41, 0x41, 0x1c, 0xec, 0x2c, 0x8e, 0x75,
	0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x56, 0xc3, 0xaa, 0xa0, 0xb1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	Take(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionClient struct {
	cc grpc1.ClientConn
}

func NewTransactionClient(cc grpc1.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Take(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/orders.transactions.take.Transaction/Take", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	Take(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) Take(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Take not implemented")
}

func RegisterTransactionServer(s grpc1.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_Take_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Take(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.transactions.take.Transaction/Take",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Take(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.transactions.take.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Take",
			Handler:    _Transaction_Take_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/orders/internal/transactions/take/transactionService.v1.proto",
}