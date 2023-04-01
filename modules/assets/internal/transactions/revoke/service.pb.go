// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/revoke/service.proto

package revoke

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
	proto.RegisterFile("modules/assets/internal/transactions/revoke/service.proto", fileDescriptor_523c2dda42f1e873)
}

var fileDescriptor_523c2dda42f1e873 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x13, 0x2a, 0x04, 0xa7, 0x82, 0x08, 0x47, 0xbd, 0xa1, 0x82, 0xe3, 0x1d, 0xe8,
	0x64, 0xb6, 0x16, 0x44, 0x11, 0x0a, 0xa5, 0x76, 0x92, 0x2e, 0xaf, 0xe9, 0x23, 0x06, 0x93, 0xbb,
	0x92, 0xbb, 0x76, 0x73, 0x71, 0x72, 0x14, 0x5c, 0x9d, 0x1c, 0xfd, 0x24, 0xe2, 0x54, 0x70, 0x71,
	0x94, 0xd4, 0xc9, 0x4f, 0x21, 0xed, 0x4b, 0x21, 0x20, 0x2d, 0x64, 0xcd, 0xff, 0xf7, 0xcf, 0xef,
	0xbd, 0x77, 0xfe, 0x59, 0x6a, 0xc6, 0xd3, 0x04, 0xad, 0x02, 0x6b, 0xd1, 0x59, 0x15, 0x6b, 0x87,
	0x99, 0x86, 0x44, 0xb9, 0x0c, 0xb4, 0x85, 0xd0, 0xc5, 0x46, 0x5b, 0x95, 0xe1, 0xcc, 0xdc, 0xa1,
	0xb2, 0x98, 0xcd, 0xe2, 0x10, 0xe5, 0x24, 0x33, 0xce, 0x34, 0x38, 0x55, 0x64, 0x99, 0x94, 0x44,
	0xf2, 0x66, 0x64, 0x4c, 0x94, 0xa0, 0x82, 0x49, 0xac, 0x40, 0x6b, 0xe3, 0x80, 0xe2, 0x55, 0x93,
	0x57, 0x92, 0xa6, 0x68, 0x2d, 0x44, 0x85, 0x94, 0x9f, 0x57, 0xa9, 0x96, 0xbe, 0xf5, 0xd1, 0x4e,
	0x8c, 0xb6, 0xc5, 0x6f, 0x4e, 0x1e, 0x3d, 0x7f, 0xf7, 0x9a, 0xb6, 0x69, 0xdc, 0xfb, 0xf5, 0x4b,
	0xd0, 0xe3, 0x04, 0x1b, 0x47, 0x72, 0xf3, 0x4a, 0xb2, 0x4b, 0x73, 0x70, 0xb5, 0x0d, 0x1a, 0xfc,
	0x37, 0xb6, 0x0e, 0x1f, 0x3e, 0x7f, 0x9e, 0xd9, 0x41, 0x6b, 0x5f, 0xa5, 0xa0, 0xdd, 0xf2, 0x28,
	0x34, 0x3a, 0x55, 0x3a, 0x2f, 0xec, 0x3d, 0x17, 0xde, 0x3c, 0x17, 0xde, 0x77, 0x2e, 0xbc, 0xa7,
	0x85, 0xa8, 0xcd, 0x17, 0xa2, 0xf6, 0xb5, 0x10, 0x35, 0x5f, 0x84, 0x26, 0xdd, 0x62, 0xeb, 0xec,
	0x15, 0x2b, 0xf4, 0x96, 0x3b, 0xf5, 0xbc, 0x9b, 0xab, 0x28, 0x76, 0xb7, 0xd3, 0x91, 0x0c, 0x4d,
	0xaa, 0xda, 0xcb, 0x5a, 0x97, 0x84, 0xeb, 0x9b, 0x55, 0xb8, 0xdd, 0x2b, 0xdb, 0x69, 0x0f, 0xfa,
	0x6f, 0x8c, 0xb7, 0x49, 0x3f, 0x28, 0xeb, 0xfb, 0x2b, 0xe4, 0x63, 0x1d, 0x0e, 0xcb, 0xe1, 0x90,
	0xc2, 0x9c, 0x1d, 0x6f, 0x0e, 0x87, 0x17, 0xbd, 0x4e, 0x17, 0x1d, 0x8c, 0xc1, 0xc1, 0x2f, 0x6b,
	0x12, 0x18, 0x04, 0x65, 0x32, 0x08, 0x08, 0x1d, 0xd5, 0x57, 0x0f, 0x76, 0xfa, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x16, 0x36, 0xb6, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assets.transactions.revoke.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assets.transactions.revoke.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.revoke.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/revoke/service.proto",
}