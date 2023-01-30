// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/revoke/service.v1.proto

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
	proto.RegisterFile("modules/identities/internal/transactions/revoke/service.v1.proto", fileDescriptor_af3caf0160ded92d)
}

var fileDescriptor_af3caf0160ded92d = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xc8, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x04, 0x31, 0xf3, 0x4a,
	0x52, 0x8b, 0xf2, 0x12, 0x73, 0xf4, 0x4b, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0xf5, 0x8b, 0x52, 0xcb, 0xf2, 0xb3, 0x53, 0xf5, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0xf5, 0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x10, 0x3a, 0xf5, 0x90,
	0x35, 0xe8, 0x41, 0x34, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64,
	0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x42, 0xa4, 0xc1, 0xba, 0xa5, 0x48, 0xb6, 0x3f, 0x37,
	0xb5, 0xb8, 0x38, 0x31, 0x1d, 0x61, 0xbf, 0x94, 0x23, 0xa9, 0x26, 0x14, 0xa5, 0x16, 0x17, 0xe4,
	0xe7, 0x15, 0x23, 0x8c, 0x30, 0xea, 0x64, 0xe4, 0x62, 0x0f, 0x86, 0xf8, 0x4b, 0xa8, 0x8e, 0x8b,
	0xcd, 0x23, 0x31, 0x2f, 0x25, 0x27, 0x55, 0x48, 0x5d, 0x0f, 0xbf, 0xcf, 0xf4, 0x7c, 0x21, 0x4e,
	0x91, 0xd2, 0x20, 0xa4, 0x30, 0x08, 0x6a, 0xa3, 0x92, 0x62, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xa4,
	0x95, 0x24, 0xf5, 0x73, 0x13, 0xf3, 0x4a, 0x72, 0x52, 0x91, 0xdd, 0x0c, 0x51, 0xeb, 0xb4, 0x8d,
	0xe9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x94, 0x92, 0xf3, 0x73, 0x09,
	0x58, 0xe5, 0xc4, 0x07, 0xf5, 0x47, 0x98, 0x61, 0x00, 0xc8, 0x6b, 0x01, 0x8c, 0x51, 0x7e, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x8e, 0xc5, 0xc5, 0xa9, 0x25, 0xbe,
	0x10, 0xab, 0x61, 0xc1, 0x46, 0x62, 0xf0, 0x2d, 0x62, 0x62, 0xf6, 0x0c, 0x09, 0x5a, 0xc5, 0x24,
	0xe7, 0x89, 0x70, 0x48, 0x08, 0xb2, 0x43, 0x82, 0xc0, 0xca, 0x4e, 0x21, 0x2b, 0x88, 0x41, 0x56,
	0x10, 0x03, 0x51, 0xf0, 0x88, 0x49, 0x0b, 0xbf, 0x82, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92,
	0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x4c, 0x0a, 0x08, 0xc5, 0x56, 0x56, 0xc8, 0xaa, 0xad, 0xac,
	0x20, 0xca, 0x93, 0xd8, 0xc0, 0x71, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x45, 0x47,
	0x3d, 0xd2, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.revoke.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.revoke.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.revoke.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/revoke/service.v1.proto",
}