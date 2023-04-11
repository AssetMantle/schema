// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/types/base/split.proto

package base

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/x/ids/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Split struct {
	OwnerID   *base.IdentityID   `protobuf:"bytes,1,opt,name=owner_i_d,json=ownerID,proto3" json:"owner_i_d,omitempty"`
	OwnableID *base.AnyOwnableID `protobuf:"bytes,2,opt,name=ownable_i_d,json=ownableID,proto3" json:"ownable_i_d,omitempty"`
	Value     string             `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Split) Reset()         { *m = Split{} }
func (m *Split) String() string { return proto.CompactTextString(m) }
func (*Split) ProtoMessage()    {}
func (*Split) Descriptor() ([]byte, []int) {
	return fileDescriptor_361b67c1261f3310, []int{0}
}
func (m *Split) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Split) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Split.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Split) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Split.Merge(m, src)
}
func (m *Split) XXX_Size() int {
	return m.Size()
}
func (m *Split) XXX_DiscardUnknown() {
	xxx_messageInfo_Split.DiscardUnknown(m)
}

var xxx_messageInfo_Split proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Split)(nil), "types.Split")
}

func init() { proto.RegisterFile("x/types/base/split.proto", fileDescriptor_361b67c1261f3310) }

var fileDescriptor_361b67c1261f3310 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xa8, 0xd0, 0x2f, 0xa9,
	0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x8b, 0x4b, 0x49, 0x57, 0xe8, 0x67, 0xa6, 0x40, 0xa5,
	0x33, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0x4b, 0x2a, 0x3d, 0x5d, 0x20, 0x6a, 0xa4, 0x64, 0x91, 0x24,
	0x13, 0xf3, 0x2a, 0xfd, 0xcb, 0xf3, 0x12, 0x93, 0x72, 0x52, 0xe1, 0xd2, 0x22, 0xe9, 0xf9, 0xe9,
	0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0x6a, 0x66, 0xe4, 0x62, 0x0d, 0x06, 0x59, 0x24,
	0xa4, 0xcd, 0xc5, 0x99, 0x5f, 0x9e, 0x97, 0x5a, 0x14, 0x9f, 0x19, 0x9f, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x6d, 0xc4, 0xaf, 0x97, 0x99, 0x52, 0xac, 0xe7, 0x09, 0xb7, 0x28, 0x88, 0x1d, 0xac,
	0xc2, 0xd3, 0x45, 0xc8, 0x90, 0x8b, 0x3b, 0x1f, 0x62, 0x3e, 0x58, 0x39, 0x13, 0x58, 0xb9, 0x20,
	0x58, 0xb9, 0x23, 0x92, 0xd5, 0x41, 0x9c, 0xf9, 0x30, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62,
	0x4e, 0x69, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x63, 0xc5, 0xd2, 0xb1, 0x40,
	0x9e, 0xc1, 0xa9, 0x9f, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x38,
	0x93, 0xf3, 0x73, 0xf5, 0xc0, 0xbe, 0x77, 0xe2, 0x02, 0x3b, 0x34, 0x00, 0xe4, 0xee, 0x00, 0xc6,
	0x28, 0xad, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2,
	0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x92, 0x9c, 0x54, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x7d,
	0xe4, 0xa0, 0x5c, 0xc4, 0xc4, 0x1c, 0x12, 0x11, 0xb1, 0x8a, 0x89, 0x35, 0x04, 0x24, 0x74, 0x0a,
	0x4a, 0x3f, 0x62, 0x12, 0x04, 0xd3, 0x31, 0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x29, 0x89,
	0x25, 0x89, 0xaf, 0xa0, 0x72, 0x49, 0x6c, 0xe0, 0xe0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x99, 0x62, 0xd1, 0x93, 0x01, 0x00, 0x00,
}

func (m *Split) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Split) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Split) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSplit(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OwnableID != nil {
		{
			size, err := m.OwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.OwnerID != nil {
		{
			size, err := m.OwnerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSplit(dAtA []byte, offset int, v uint64) int {
	offset -= sovSplit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Split) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnerID != nil {
		l = m.OwnerID.Size()
		n += 1 + l + sovSplit(uint64(l))
	}
	if m.OwnableID != nil {
		l = m.OwnableID.Size()
		n += 1 + l + sovSplit(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSplit(uint64(l))
	}
	return n
}

func sovSplit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSplit(x uint64) (n int) {
	return sovSplit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Split) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSplit
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
			return fmt.Errorf("proto: Split: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Split: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplit
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
				return ErrInvalidLengthSplit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerID == nil {
				m.OwnerID = &base.IdentityID{}
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplit
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
				return ErrInvalidLengthSplit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnableID == nil {
				m.OwnableID = &base.AnyOwnableID{}
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplit
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
				return ErrInvalidLengthSplit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSplit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSplit
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
func skipSplit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSplit
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
					return 0, ErrIntOverflowSplit
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
					return 0, ErrIntOverflowSplit
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
				return 0, ErrInvalidLengthSplit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSplit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSplit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSplit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSplit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSplit = fmt.Errorf("proto: unexpected end of group")
)
