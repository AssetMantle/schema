// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/deputize/transactionService.v1.proto

package deputize

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
	proto.RegisterFile("modules/assets/internal/transactions/deputize/transactionService.v1.proto", fileDescriptor_037742258bdbd1d6)
}

var fileDescriptor_037742258bdbd1d6 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xcc, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x4f, 0x49, 0x2d, 0x28, 0x2d, 0xc9, 0xac, 0x4a, 0x45, 0x16, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0x18, 0xa1,
	0x87, 0xac, 0x53, 0x0f, 0xa6, 0x53, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1,
	0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0x11, 0xa2, 0x00, 0xac, 0x57, 0xca, 0x8e, 0x34,
	0x67, 0xe4, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0x23, 0xec, 0x96, 0xf2, 0x22, 0xdb, 0x1b, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0x08, 0xb3, 0x8c, 0xa6, 0x32, 0x72, 0x71, 0x87, 0x20, 0x14, 0x08,
	0xb5, 0x32, 0x72, 0x71, 0xb8, 0x40, 0x75, 0x0a, 0xa9, 0xea, 0xe1, 0xf3, 0xa5, 0x9e, 0x2f, 0xc4,
	0x61, 0x52, 0x86, 0xf8, 0x95, 0x85, 0x60, 0xda, 0xaf, 0x24, 0xdf, 0x74, 0xf9, 0xc9, 0x64, 0x26,
	0x49, 0x25, 0x71, 0xfd, 0xdc, 0xc4, 0xbc, 0x12, 0x50, 0x48, 0x41, 0xbc, 0x02, 0xd3, 0xe4, 0xb4,
	0x95, 0xe9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x14, 0x92, 0xf3, 0x73,
	0xf1, 0xda, 0xe8, 0x24, 0x19, 0x82, 0x11, 0x73, 0x61, 0x86, 0x01, 0x20, 0xff, 0x06, 0x30, 0x46,
	0xf9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x82, 0x4c, 0xf1,
	0x85, 0xb8, 0x00, 0x16, 0xa8, 0x24, 0x05, 0xee, 0x22, 0x26, 0x66, 0xc7, 0x10, 0x97, 0x55, 0x4c,
	0x32, 0x8e, 0x10, 0xf7, 0x84, 0x20, 0xbb, 0x07, 0x16, 0x8e, 0xa7, 0x60, 0xd2, 0x31, 0xc8, 0xd2,
	0x31, 0x30, 0xe9, 0x47, 0x4c, 0x1a, 0xf8, 0xa4, 0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12,
	0x53, 0x12, 0x4b, 0x12, 0x5f, 0x31, 0xc9, 0x41, 0x94, 0x5a, 0x59, 0x21, 0xab, 0xb5, 0xb2, 0x82,
	0x29, 0x4e, 0x62, 0x03, 0x47, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x65, 0x8d, 0x3c, 0x08,
	0xeb, 0x02, 0x00, 0x00,
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
	Deputize(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionClient struct {
	cc grpc1.ClientConn
}

func NewTransactionClient(cc grpc1.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Deputize(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assets.transactions.deputize.Transaction/Deputize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	Deputize(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) Deputize(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deputize not implemented")
}

func RegisterTransactionServer(s grpc1.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_Deputize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Deputize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assets.transactions.deputize.Transaction/Deputize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Deputize(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.deputize.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deputize",
			Handler:    _Transaction_Deputize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/deputize/transactionService.v1.proto",
}