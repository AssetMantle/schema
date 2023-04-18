// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/base/height.proto

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

type Height struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Height) Reset()         { *m = Height{} }
func (m *Height) String() string { return proto.CompactTextString(m) }
func (*Height) ProtoMessage()    {}
func (*Height) Descriptor() ([]byte, []int) {
	return fileDescriptor_7588447ac7954677, []int{0}
}
func (m *Height) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Height) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Height.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Height) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Height.Merge(m, src)
}
func (m *Height) XXX_Size() int {
	return m.Size()
}
func (m *Height) XXX_DiscardUnknown() {
	xxx_messageInfo_Height.DiscardUnknown(m)
}

var xxx_messageInfo_Height proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Height)(nil), "types.Height")
}

func init() { proto.RegisterFile("types/base/height.proto", fileDescriptor_7588447ac7954677) }

var fileDescriptor_7588447ac7954677 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x45, 0xf4,
	0x41, 0x2c, 0x88, 0xa4, 0x92, 0x0a, 0x17, 0x9b, 0x07, 0x58, 0xb1, 0x90, 0x08, 0x17, 0x6b, 0x59,
	0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0x63, 0xc5, 0xd2, 0xb1,
	0x40, 0x9e, 0xc1, 0x69, 0x22, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70,
	0x71, 0x26, 0xe7, 0xe7, 0xea, 0x81, 0x6d, 0x70, 0xe2, 0x86, 0x98, 0x14, 0x00, 0x32, 0x38, 0x80,
	0x31, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8,
	0x38, 0xb5, 0xc4, 0x37, 0x31, 0xaf, 0x24, 0x27, 0x55, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x51,
	0x3f, 0x3d, 0x5f, 0x1f, 0xe1, 0xe2, 0x45, 0x4c, 0xcc, 0x21, 0x11, 0x11, 0xab, 0x98, 0x58, 0x43,
	0x40, 0x42, 0xa7, 0xa0, 0xf4, 0x23, 0x26, 0x41, 0x30, 0x1d, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a,
	0x92, 0x98, 0x92, 0x58, 0x92, 0xf8, 0x0a, 0x2a, 0x97, 0xc4, 0x06, 0xf6, 0x80, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x73, 0x1c, 0x08, 0x7a, 0xf8, 0x00, 0x00, 0x00,
}

func (m *Height) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Height) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Height) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintHeight(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeight(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeight(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Height) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovHeight(uint64(m.Value))
	}
	return n
}

func sovHeight(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeight(x uint64) (n int) {
	return sovHeight(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Height) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeight
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
			return fmt.Errorf("proto: Height: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Height: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeight
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
			skippy, err := skipHeight(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHeight
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
func skipHeight(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeight
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
					return 0, ErrIntOverflowHeight
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
					return 0, ErrIntOverflowHeight
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
				return 0, ErrInvalidLengthHeight
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeight
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeight
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeight        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeight          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeight = fmt.Errorf("proto: unexpected end of group")
)
