// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/data/base/booleanData.proto

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

type BooleanData struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BooleanData) Reset()         { *m = BooleanData{} }
func (m *BooleanData) String() string { return proto.CompactTextString(m) }
func (*BooleanData) ProtoMessage()    {}
func (*BooleanData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ca566743c419af, []int{0}
}
func (m *BooleanData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BooleanData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BooleanData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BooleanData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanData.Merge(m, src)
}
func (m *BooleanData) XXX_Size() int {
	return m.Size()
}
func (m *BooleanData) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanData.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BooleanData)(nil), "data.BooleanData")
}

func init() { proto.RegisterFile("x/data/base/booleanData.proto", fileDescriptor_c6ca566743c419af) }

var fileDescriptor_c6ca566743c419af = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xad, 0xd0, 0x4f, 0x49,
	0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xca, 0xcf, 0xcf, 0x49, 0x4d, 0xcc, 0x73,
	0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x49, 0x4a, 0x89, 0xa4,
	0xe7, 0xa7, 0xe7, 0x83, 0x05, 0xf4, 0x41, 0x2c, 0x88, 0x9c, 0x92, 0x26, 0x17, 0xb7, 0x13, 0x42,
	0x83, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x84, 0x63, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0xa9, 0x9f, 0xf1, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x38, 0x92, 0xf3, 0x73, 0xf5, 0x40, 0xb6, 0x38, 0x09, 0x20,
	0x99, 0x16, 0x00, 0xb2, 0x21, 0x80, 0x31, 0x4a, 0x33, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8, 0x38, 0xb5, 0xc4, 0x37, 0x31, 0xaf, 0x24, 0x27, 0x55, 0xbf,
	0x38, 0x39, 0x23, 0x35, 0x37, 0x51, 0x1f, 0xc9, 0xf1, 0x8b, 0x98, 0x98, 0x5d, 0x22, 0x22, 0x56,
	0x31, 0xb1, 0x80, 0xb4, 0x9f, 0x82, 0x50, 0x8f, 0x98, 0x04, 0x40, 0x54, 0x8c, 0x7b, 0x80, 0x93,
	0x6f, 0x6a, 0x49, 0x22, 0x48, 0xed, 0x2b, 0x88, 0x4c, 0x12, 0x1b, 0xd8, 0x0f, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x08, 0x45, 0xa4, 0xa8, 0x00, 0x01, 0x00, 0x00,
}

func (m *BooleanData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BooleanData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BooleanData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBooleanData(dAtA []byte, offset int, v uint64) int {
	offset -= sovBooleanData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BooleanData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func sovBooleanData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBooleanData(x uint64) (n int) {
	return sovBooleanData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BooleanData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBooleanData
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
			return fmt.Errorf("proto: BooleanData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BooleanData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooleanData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBooleanData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBooleanData
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
func skipBooleanData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBooleanData
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
					return 0, ErrIntOverflowBooleanData
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
					return 0, ErrIntOverflowBooleanData
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
				return 0, ErrInvalidLengthBooleanData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBooleanData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBooleanData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBooleanData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBooleanData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBooleanData = fmt.Errorf("proto: unexpected end of group")
)
