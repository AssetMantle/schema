// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/schema/qualified/base/mutables.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/lists/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Mutables struct {
	PropertyList *base.PropertyList `protobuf:"bytes,1,opt,name=property_list,json=propertyList,proto3" json:"property_list,omitempty"`
}

func (m *Mutables) Reset()         { *m = Mutables{} }
func (m *Mutables) String() string { return proto.CompactTextString(m) }
func (*Mutables) ProtoMessage()    {}
func (*Mutables) Descriptor() ([]byte, []int) {
	return fileDescriptor_68db4e6cfc1258ed, []int{0}
}
func (m *Mutables) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mutables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mutables.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mutables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutables.Merge(m, src)
}
func (m *Mutables) XXX_Size() int {
	return m.Size()
}
func (m *Mutables) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutables.DiscardUnknown(m)
}

var xxx_messageInfo_Mutables proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Mutables)(nil), "AssetMantle.schema.qualified.base.Mutables")
}

func init() {
	proto.RegisterFile("AssetMantle/schema/qualified/base/mutables.proto", fileDescriptor_68db4e6cfc1258ed)
}

var fileDescriptor_68db4e6cfc1258ed = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x70, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f,
	0x2c, 0x4d, 0xcc, 0xc9, 0x4c, 0xcb, 0x4c, 0x4d, 0xd1, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x2d,
	0x2d, 0x49, 0x4c, 0xca, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x44, 0xd2,
	0xa1, 0x07, 0xd1, 0xa1, 0x07, 0xd7, 0xa1, 0x07, 0xd2, 0x21, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xad, 0x0f, 0x62, 0x41, 0x34, 0x4a, 0x19, 0x62, 0xb1, 0x2a, 0x27, 0xb3, 0xb8, 0xa4, 0x18,
	0x62, 0x4d, 0x41, 0x51, 0x7e, 0x41, 0x6a, 0x51, 0x49, 0x65, 0x3c, 0x48, 0x0c, 0xa2, 0x45, 0x29,
	0x89, 0x8b, 0xc3, 0x17, 0x6a, 0xbb, 0x50, 0x00, 0x17, 0x2f, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x6d, 0x3d, 0x2c, 0xee, 0x01, 0x1b, 0x0b, 0x76, 0x8b, 0x5e, 0x00, 0x54, 0x8f,
	0x4f, 0x66, 0x71, 0x49, 0x10, 0x4f, 0x01, 0x12, 0xcf, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0xd3,
	0x12, 0xa6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x52, 0x4d, 0xce, 0xcf,
	0xd5, 0x23, 0xe8, 0x5d, 0x27, 0x5e, 0x98, 0x1b, 0x03, 0x40, 0x8e, 0x0e, 0x60, 0x8c, 0xd2, 0x49,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x27, 0x18, 0xbe, 0x8b, 0x98, 0x58,
	0x1c, 0x83, 0x03, 0x9d, 0x56, 0x31, 0xa1, 0x04, 0x6c, 0x30, 0xc4, 0xa6, 0x40, 0xb8, 0x4d, 0x4e,
	0x89, 0xc5, 0xa9, 0xa7, 0x50, 0xd4, 0xc4, 0x40, 0xd4, 0xc4, 0xc0, 0xd5, 0xc4, 0x80, 0xd4, 0x3c,
	0x62, 0xd2, 0x25, 0xa8, 0x26, 0xc6, 0x3d, 0xc0, 0xc9, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24,
	0xf1, 0x15, 0x93, 0x0a, 0x92, 0x7a, 0x2b, 0x2b, 0x88, 0x06, 0x2b, 0x2b, 0xb8, 0x0e, 0x2b, 0x2b,
	0x90, 0x96, 0x24, 0x36, 0x70, 0x8c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x8c, 0x3a,
	0x46, 0x31, 0x02, 0x00, 0x00,
}

func (m *Mutables) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mutables) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mutables) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PropertyList != nil {
		{
			size, err := m.PropertyList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMutables(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMutables(dAtA []byte, offset int, v uint64) int {
	offset -= sovMutables(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mutables) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropertyList != nil {
		l = m.PropertyList.Size()
		n += 1 + l + sovMutables(uint64(l))
	}
	return n
}

func sovMutables(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMutables(x uint64) (n int) {
	return sovMutables(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mutables) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMutables
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
			return fmt.Errorf("proto: Mutables: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mutables: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutables
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
				return ErrInvalidLengthMutables
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMutables
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PropertyList == nil {
				m.PropertyList = &base.PropertyList{}
			}
			if err := m.PropertyList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMutables(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMutables
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
func skipMutables(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMutables
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
					return 0, ErrIntOverflowMutables
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
					return 0, ErrIntOverflowMutables
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
				return 0, ErrInvalidLengthMutables
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMutables
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMutables
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMutables        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMutables          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMutables = fmt.Errorf("proto: unexpected end of group")
)
