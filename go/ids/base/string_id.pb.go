// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/string_id.proto

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

type StringID struct {
	IDString string `protobuf:"bytes,1,opt,name=i_d_string,json=iDString,proto3" json:"i_d_string,omitempty"`
}

func (m *StringID) Reset()         { *m = StringID{} }
func (m *StringID) String() string { return proto.CompactTextString(m) }
func (*StringID) ProtoMessage()    {}
func (*StringID) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b6d40ef8933e6b, []int{0}
}
func (m *StringID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringID.Merge(m, src)
}
func (m *StringID) XXX_Size() int {
	return m.Size()
}
func (m *StringID) XXX_DiscardUnknown() {
	xxx_messageInfo_StringID.DiscardUnknown(m)
}

var xxx_messageInfo_StringID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StringID)(nil), "ids.StringID")
}

func init() { proto.RegisterFile("ids/base/string_id.proto", fileDescriptor_12b6d40ef8933e6b) }

var fileDescriptor_12b6d40ef8933e6b = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x8f, 0xcf, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x92, 0x1e, 0x17, 0x47, 0x30, 0x58, 0xb5, 0xa7, 0x8b, 0x90,
	0x0c, 0x17, 0x57, 0x66, 0x7c, 0x4a, 0x3c, 0x44, 0xb7, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x47, 0xa6, 0x0b, 0x44, 0xde, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0x53, 0x3b, 0xe3, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0xea, 0x65, 0xa6, 0x14,
	0x3b, 0xf1, 0x42, 0x4d, 0x4c, 0x09, 0x00, 0x59, 0x11, 0xc0, 0x18, 0xa5, 0x99, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57,
	0x92, 0x93, 0xaa, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x9f, 0x9e, 0xaf, 0x0f, 0x73, 0xfa,
	0x22, 0x26, 0x66, 0xcf, 0x88, 0x88, 0x55, 0x4c, 0xcc, 0x9e, 0x29, 0xc5, 0xa7, 0xc0, 0xe4, 0x23,
	0x26, 0x7e, 0xcf, 0x94, 0xe2, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92,
	0xc4, 0x57, 0x60, 0xf1, 0x24, 0x36, 0xb0, 0x07, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0x6e, 0x49, 0x72, 0xf7, 0x00, 0x00, 0x00,
}

func (m *StringID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IDString) > 0 {
		i -= len(m.IDString)
		copy(dAtA[i:], m.IDString)
		i = encodeVarintStringId(dAtA, i, uint64(len(m.IDString)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStringId(dAtA []byte, offset int, v uint64) int {
	offset -= sovStringId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StringID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IDString)
	if l > 0 {
		n += 1 + l + sovStringId(uint64(l))
	}
	return n
}

func sovStringId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStringId(x uint64) (n int) {
	return sovStringId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStringId
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
			return fmt.Errorf("proto: StringID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStringId
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
				return ErrInvalidLengthStringId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStringId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStringId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStringId
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
func skipStringId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStringId
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
					return 0, ErrIntOverflowStringId
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
					return 0, ErrIntOverflowStringId
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
				return 0, ErrInvalidLengthStringId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStringId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStringId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStringId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStringId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStringId = fmt.Errorf("proto: unexpected end of group")
)