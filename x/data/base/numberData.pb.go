// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/data/base/numberData.proto

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

type NumberData struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NumberData) Reset()         { *m = NumberData{} }
func (m *NumberData) String() string { return proto.CompactTextString(m) }
func (*NumberData) ProtoMessage()    {}
func (*NumberData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca06aa51be98fe0, []int{0}
}
func (m *NumberData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumberData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumberData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NumberData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberData.Merge(m, src)
}
func (m *NumberData) XXX_Size() int {
	return m.Size()
}
func (m *NumberData) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberData.DiscardUnknown(m)
}

var xxx_messageInfo_NumberData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NumberData)(nil), "data.NumberData")
}

func init() { proto.RegisterFile("x/data/base/numberData.proto", fileDescriptor_3ca06aa51be98fe0) }

var fileDescriptor_3ca06aa51be98fe0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xa9, 0xd0, 0x4f, 0x49,
	0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x72, 0x49,
	0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9, 0x49, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x05, 0xf4, 0x41, 0x2c, 0x88, 0x9c, 0x92, 0x06, 0x17, 0x97, 0x1f, 0x5c, 0xbd,
	0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10,
	0x84, 0x63, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0xa9, 0x8f, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x38, 0x92, 0xf3, 0x73, 0xf5, 0x40, 0x96, 0x38, 0xf1, 0x23, 0x0c,
	0x0b, 0x00, 0x99, 0x1f, 0xc0, 0x18, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57, 0x92, 0x93, 0xaa, 0x5f, 0x9c,
	0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x8f, 0xe4, 0xf2, 0x45, 0x4c, 0xcc, 0x2e, 0x11, 0x11, 0xab, 0x98,
	0x58, 0x40, 0xda, 0x4f, 0x41, 0xa8, 0x47, 0x4c, 0x02, 0x20, 0x2a, 0xc6, 0x3d, 0xc0, 0xc9, 0x37,
	0xb5, 0x24, 0x11, 0xa4, 0xf6, 0x15, 0x44, 0x26, 0x89, 0x0d, 0xec, 0x03, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x17, 0x87, 0x01, 0x77, 0xfd, 0x00, 0x00, 0x00,
}

func (m *NumberData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumberData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NumberData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintNumberData(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNumberData(dAtA []byte, offset int, v uint64) int {
	offset -= sovNumberData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NumberData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovNumberData(uint64(m.Value))
	}
	return n
}

func sovNumberData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNumberData(x uint64) (n int) {
	return sovNumberData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NumberData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNumberData
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
			return fmt.Errorf("proto: NumberData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumberData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNumberData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNumberData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNumberData
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
func skipNumberData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNumberData
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
					return 0, ErrIntOverflowNumberData
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
					return 0, ErrIntOverflowNumberData
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
				return 0, ErrInvalidLengthNumberData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNumberData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNumberData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNumberData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNumberData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNumberData = fmt.Errorf("proto: unexpected end of group")
)
