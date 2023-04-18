// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data/base/boolean_data.proto

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
	return fileDescriptor_3b3cd0c49be158d4, []int{0}
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

func init() { proto.RegisterFile("data/base/boolean_data.proto", fileDescriptor_3b3cd0c49be158d4) }

var fileDescriptor_3b3cd0c49be158d4 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x49, 0x2c, 0x49,
	0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xca, 0xcf, 0xcf, 0x49, 0x4d, 0xcc, 0x8b, 0x07, 0x89,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x80, 0xd8, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9,
	0x60, 0x01, 0x7d, 0x10, 0x0b, 0x22, 0xa7, 0xa4, 0xc9, 0xc5, 0xed, 0x04, 0xd1, 0xe1, 0x92, 0x58,
	0x92, 0x28, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x11, 0x04, 0xe1, 0x58, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0x9a, 0xc0, 0x78, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0x1c, 0xc9, 0xf9, 0xb9, 0x7a, 0x20, 0x5b, 0x9c, 0x04,
	0x90, 0x4c, 0x0b, 0x00, 0xd9, 0x10, 0xc0, 0x18, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57, 0x92, 0x93, 0xaa,
	0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x9f, 0x9e, 0xaf, 0x0f, 0x77, 0xfe, 0x22, 0x26, 0x66,
	0x97, 0x88, 0x88, 0x55, 0x4c, 0x2c, 0x20, 0xfd, 0xa7, 0x20, 0xd4, 0x23, 0x26, 0x01, 0x10, 0x15,
	0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x08, 0x52, 0xfb, 0x0a, 0x22, 0x93, 0xc4, 0x06, 0xf6,
	0x84, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x33, 0x28, 0x62, 0x88, 0x00, 0x01, 0x00, 0x00,
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
