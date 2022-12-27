// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/renumerate/service.v1.proto

package renumerate

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
	proto.RegisterFile("modules/assets/internal/transactions/renumerate/service.v1.proto", fileDescriptor_385bcbff35193a43)
}

var fileDescriptor_385bcbff35193a43 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x69, 0x4d, 0x30, 0xe9, 0xe0, 0xc0, 0x66, 0x35, 0x17, 0x65, 0xd1, 0x38, 0xdc, 0x05,
	0xdd, 0x3a, 0x59, 0x16, 0x5d, 0x30, 0x04, 0x89, 0x83, 0x61, 0x79, 0x29, 0x6f, 0x6a, 0x93, 0xde,
	0x1d, 0xb9, 0x3b, 0x18, 0x1d, 0xdc, 0xdc, 0x4c, 0xfc, 0x06, 0x8e, 0x7e, 0x00, 0x3f, 0x83, 0x71,
	0x22, 0x71, 0x71, 0x34, 0xc5, 0xc9, 0x4f, 0x61, 0xe0, 0xc0, 0xeb, 0x04, 0x61, 0xff, 0x3d, 0xef,
	0xf3, 0xe7, 0x0d, 0xce, 0xb9, 0x1c, 0x8c, 0x72, 0xd4, 0x0c, 0xb4, 0x46, 0xa3, 0x59, 0x26, 0x0c,
	0x2a, 0x01, 0x39, 0x33, 0x0a, 0x84, 0x86, 0xc4, 0x64, 0x52, 0x68, 0xa6, 0x50, 0x8c, 0x38, 0x2a,
	0x30, 0xc8, 0x34, 0xaa, 0x71, 0x96, 0x20, 0x1d, 0x37, 0xe8, 0x50, 0x49, 0x23, 0x6b, 0xc4, 0x2a,
	0x69, 0x59, 0x40, 0x9d, 0x20, 0xdc, 0x4f, 0xa5, 0x4c, 0x73, 0x64, 0x30, 0xcc, 0x18, 0x08, 0x21,
	0x0d, 0x58, 0x64, 0xae, 0x0e, 0x37, 0xf6, 0xe7, 0xa8, 0x35, 0xa4, 0xce, 0x3f, 0x8c, 0x37, 0xbd,
	0xa0, 0x50, 0x0f, 0xa5, 0xd0, 0xee, 0xc4, 0xe9, 0xa3, 0x17, 0x6c, 0x5f, 0xdb, 0x5e, 0xb5, 0xfb,
	0xa0, 0x7a, 0x09, 0x62, 0x90, 0x63, 0xed, 0x88, 0xae, 0x6e, 0x46, 0x5b, 0x36, 0x4a, 0x78, 0xbc,
	0x0e, 0xec, 0x2c, 0x1c, 0xeb, 0x87, 0x0f, 0x9f, 0x3f, 0xcf, 0xfe, 0x5e, 0x7d, 0x97, 0x71, 0x10,
	0x66, 0x36, 0x8a, 0xcd, 0xec, 0xd8, 0xe6, 0x9b, 0xff, 0x5e, 0x10, 0x6f, 0x52, 0x10, 0xef, 0xbb,
	0x20, 0xde, 0xd3, 0x94, 0x54, 0x26, 0x53, 0x52, 0xf9, 0x9a, 0x92, 0x4a, 0x50, 0x4f, 0x24, 0x5f,
	0x63, 0xd5, 0xdc, 0x59, 0xf4, 0xb8, 0x69, 0xb4, 0x67, 0xd5, 0xda, 0xde, 0xed, 0x55, 0x9a, 0x99,
	0xbb, 0x51, 0x9f, 0x26, 0x92, 0xb3, 0x78, 0x26, 0x6e, 0x59, 0xeb, 0xe5, 0x6c, 0x1b, 0xce, 0xf7,
	0xe2, 0x6f, 0xc5, 0xdd, 0xce, 0xab, 0x4f, 0x62, 0x1b, 0xa4, 0x5b, 0x0e, 0xd2, 0xf9, 0xc7, 0x3e,
	0x96, 0x40, 0xaf, 0x0c, 0xf4, 0x1c, 0x50, 0xf8, 0x27, 0xab, 0x81, 0xde, 0x45, 0xbb, 0xd9, 0x42,
	0x03, 0x03, 0x30, 0xf0, 0xeb, 0x1f, 0x58, 0x38, 0x8a, 0xca, 0x74, 0x14, 0x39, 0xbc, 0x5f, 0x9d,
	0xff, 0xf2, 0xec, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x95, 0x71, 0x4f, 0x1c, 0xd2, 0x02, 0x00, 0x00,
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
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/assets.transactions.renumerate.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*Response, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*Response, error) {
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
		FullMethod: "/assets.transactions.renumerate.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.renumerate.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/renumerate/service.v1.proto",
}
