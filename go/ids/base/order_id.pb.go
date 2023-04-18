// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/order_id.proto

package base

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type OrderID struct {
	HashID *HashID `protobuf:"bytes,1,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d,omitempty"`
}

func (m *OrderID) Reset()         { *m = OrderID{} }
func (m *OrderID) String() string { return proto.CompactTextString(m) }
func (*OrderID) ProtoMessage()    {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8320e3936250525c, []int{0}
}
func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return m.Size()
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OrderID)(nil), "ids.OrderID")
}

func init() { proto.RegisterFile("ids/base/order_id.proto", fileDescriptor_8320e3936250525c) }

var fileDescriptor_8320e3936250525c = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x8a, 0xcf, 0x4c, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x12, 0x83, 0xcb, 0x66, 0x24, 0x16, 0x67,
	0xc0, 0x25, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0x64,
	0xc6, 0xc5, 0xee, 0x0f, 0x32, 0xc4, 0xd3, 0x45, 0x48, 0x95, 0x8b, 0x03, 0xa2, 0x23, 0x3e, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x5b, 0x2f, 0x33, 0xa5, 0x58, 0xcf, 0x23, 0xb1, 0x38,
	0xc3, 0xd3, 0x25, 0x88, 0x2d, 0x03, 0x4c, 0x5b, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0x6a, 0x63,
	0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xf6, 0xe4, 0xfc, 0x5c, 0x90,
	0x46, 0x27, 0x1e, 0x88, 0xc9, 0x29, 0x01, 0x20, 0x9b, 0x02, 0x18, 0xa3, 0x34, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x8b, 0x8b, 0x53, 0x4b, 0x7c, 0x13, 0xf3,
	0x4a, 0x72, 0x52, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0xf5, 0xd3, 0xf3, 0xf5, 0x61, 0x4e,
	0x5f, 0xc4, 0xc4, 0xec, 0x19, 0x11, 0xb1, 0x8a, 0x89, 0xd9, 0x33, 0xa5, 0xf8, 0x14, 0x98, 0x7c,
	0xc4, 0xc4, 0xef, 0x99, 0x52, 0x1c, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58,
	0x92, 0xf8, 0x0a, 0x2c, 0x9e, 0xc4, 0x06, 0xf6, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0xe2, 0x94, 0xd5, 0x15, 0x01, 0x00, 0x00,
}

func (m *OrderID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HashID != nil {
		{
			size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrderId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderId(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovOrderId(uint64(l))
	}
	return n
}

func sovOrderId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderId(x uint64) (n int) {
	return sovOrderId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderId
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
			return fmt.Errorf("proto: OrderID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrderId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrderId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HashID == nil {
				m.HashID = &HashID{}
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrderId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderId
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
func skipOrderId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderId
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
					return 0, ErrIntOverflowOrderId
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
					return 0, ErrIntOverflowOrderId
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
				return 0, ErrInvalidLengthOrderId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderId = fmt.Errorf("proto: unexpected end of group")
)