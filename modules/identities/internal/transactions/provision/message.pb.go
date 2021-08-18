// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/transactions/provision/message.proto

package provision

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Message struct {
	From       github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required field From missing, matches(^[a-z0-9]*$)~field To is invalid"`
	To         github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,2,opt,name=to,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"to" valid:"required~required field To missing, matches(^[a-z0-9]*$)~field To is invalid"`
	IdentityID github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"identity_i_d" valid:"required~required IdentityID missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_0604abea0b2258dd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0604abea0b2258dd, []int{1}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "modules.identities.internal.transactions.provision.Message")
	proto.RegisterType((*TransactionResponse)(nil), "modules.identities.internal.transactions.provision.TransactionResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/transactions/provision/message.proto", fileDescriptor_0604abea0b2258dd)
}

var fileDescriptor_0604abea0b2258dd = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x69, 0x05, 0xf4, 0xc4, 0x64, 0x40, 0x8a, 0x3a, 0xd8, 0x96, 0xc5, 0x50, 0x50,
	0x7b, 0x87, 0xca, 0x02, 0x65, 0x6a, 0x15, 0x40, 0x11, 0xaa, 0xa8, 0x42, 0xc4, 0x40, 0x05, 0xd1,
	0xc5, 0xf7, 0xea, 0x9c, 0x88, 0xef, 0xcc, 0xbd, 0x4b, 0xa5, 0x32, 0x74, 0x66, 0x04, 0x36, 0xb6,
	0x2c, 0x88, 0x6f, 0xc0, 0xc4, 0x07, 0xe8, 0xd8, 0x11, 0x31, 0x44, 0x28, 0x59, 0x98, 0xfb, 0x09,
	0x50, 0x1d, 0xa7, 0xcd, 0x00, 0x82, 0x44, 0xdd, 0x9e, 0x2d, 0xdf, 0xcf, 0xbf, 0xbf, 0xfd, 0x7f,
	0x74, 0x27, 0x07, 0x8b, 0x0a, 0x1d, 0xe8, 0x04, 0x5a, 0x28, 0x5f, 0xf3, 0xcc, 0xc8, 0x5e, 0x17,
	0x90, 0x2b, 0x09, 0xda, 0x29, 0xa7, 0x4e, 0x47, 0xed, 0xc0, 0x6a, 0xd1, 0xe5, 0xce, 0x0a, 0x8d,
	0x22, 0x71, 0xca, 0x68, 0xe4, 0xb9, 0x35, 0xfb, 0x0a, 0x95, 0xd1, 0x3c, 0x03, 0x44, 0x91, 0x02,
	0xcb, 0xad, 0x71, 0xc6, 0x5f, 0x2f, 0x09, 0xec, 0x9c, 0xc0, 0x26, 0x04, 0x36, 0x4d, 0x60, 0x67,
	0x84, 0xe5, 0xeb, 0xa9, 0x49, 0x4d, 0x71, 0x9c, 0x9f, 0x4e, 0x63, 0x52, 0xdc, 0x5f, 0xa4, 0x97,
	0xb7, 0xc7, 0x6c, 0xff, 0x1b, 0xa1, 0x8b, 0x7b, 0xd6, 0x64, 0x55, 0x12, 0x91, 0x95, 0xa5, 0xad,
	0x2f, 0xe4, 0x68, 0x10, 0x7a, 0x3f, 0x06, 0xe1, 0xc3, 0x54, 0xb9, 0x4e, 0xaf, 0xcd, 0x12, 0x93,
	0xf1, 0xa9, 0x28, 0x4f, 0x35, 0x4c, 0x5f, 0x3e, 0xab, 0x3d, 0xe1, 0x98, 0x74, 0x20, 0x13, 0xdc,
	0x1d, 0xe4, 0x80, 0xbc, 0x2d, 0x10, 0xd8, 0x66, 0x92, 0x6c, 0x4a, 0x69, 0x01, 0xf1, 0x64, 0x10,
	0x3e, 0xdf, 0x17, 0x5d, 0x25, 0x37, 0x62, 0x0b, 0x6f, 0x7a, 0xca, 0x82, 0x3c, 0x9c, 0x0c, 0xd1,
	0x9e, 0x82, 0xae, 0x8c, 0x1e, 0x59, 0x93, 0x45, 0x99, 0x42, 0x54, 0x3a, 0x5d, 0x8d, 0x32, 0xe1,
	0x92, 0x0e, 0xe0, 0xca, 0xab, 0x5d, 0xb1, 0xf6, 0xf6, 0xce, 0xda, 0xfd, 0x97, 0xb7, 0x6f, 0xde,
	0x3a, 0x1c, 0x3f, 0xd8, 0x34, 0x91, 0xc2, 0x48, 0xe9, 0x82, 0x17, 0x37, 0x0a, 0x6b, 0xff, 0x2b,
	0xa1, 0x15, 0x67, 0xaa, 0x95, 0x42, 0xfe, 0xf3, 0x85, 0xcb, 0x37, 0xff, 0x21, 0xdf, 0x34, 0xf3,
	0xa8, 0x57, 0x9c, 0xf1, 0x3f, 0x12, 0x7a, 0xb5, 0xfc, 0x91, 0x07, 0x2d, 0xd5, 0x92, 0xd5, 0x85,
	0x22, 0x42, 0x5e, 0x26, 0xb8, 0x37, 0x57, 0x02, 0x56, 0xaf, 0x9d, 0x0c, 0xc2, 0xd5, 0xbf, 0x4a,
	0xd7, 0xcb, 0x17, 0xd6, 0x6b, 0x13, 0xed, 0xb8, 0x41, 0xd5, 0xd9, 0xcd, 0x8d, 0x2b, 0xef, 0xfa,
	0xa1, 0xf7, 0xab, 0x1f, 0x7a, 0xf1, 0x0d, 0x7a, 0xad, 0x79, 0x5e, 0xa9, 0x06, 0x60, 0x6e, 0x34,
	0xc2, 0xfa, 0x27, 0x42, 0x17, 0xb6, 0x31, 0xf5, 0x3f, 0x10, 0xba, 0xb4, 0x33, 0x69, 0x99, 0xff,
	0x80, 0xcd, 0x5e, 0x4d, 0x56, 0x16, 0x70, 0xf9, 0xf1, 0x3c, 0x87, 0xff, 0xe0, 0xb6, 0xd5, 0x3b,
	0x1a, 0x06, 0xe4, 0x78, 0x18, 0x90, 0x9f, 0xc3, 0x80, 0xbc, 0x1f, 0x05, 0xde, 0xf1, 0x28, 0xf0,
	0xbe, 0x8f, 0x02, 0xef, 0xc5, 0xee, 0x7f, 0x7f, 0xcc, 0xd9, 0x97, 0xb4, 0x7d, 0xa9, 0xd8, 0xa9,
	0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x23, 0x4e, 0xa3, 0x96, 0xf1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Provision(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Provision(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/modules.identities.internal.transactions.provision.Msg/Provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Provision(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Provision(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provision not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modules.identities.internal.transactions.provision.Msg/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Provision(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modules.identities.internal.transactions.provision.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provision",
			Handler:    _Msg_Provision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/identities/internal/transactions/provision/message.proto",
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.IdentityID.Size()
		i -= size
		if _, err := m.IdentityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.To.Size()
		i -= size
		if _, err := m.To.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.To.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.IdentityID.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
