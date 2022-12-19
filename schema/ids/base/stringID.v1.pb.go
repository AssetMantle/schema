// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/stringID.v1.proto

package base

import (
	fmt "fmt"
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
	IdString string `protobuf:"bytes,1,opt,name=id_string,json=idString,proto3" json:"id_string,omitempty"`
}

func (m *StringID) Reset()         { *m = StringID{} }
func (m *StringID) String() string { return proto.CompactTextString(m) }
func (*StringID) ProtoMessage()    {}
func (*StringID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f375577fcdb66a10, []int{0}
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

func (m *StringID) GetIdString() string {
	if m != nil {
		return m.IdString
	}
	return ""
}

func init() {
	proto.RegisterType((*StringID)(nil), "ids.StringID")
}

func init() { proto.RegisterFile("schema/ids/base/stringID.v1.proto", fileDescriptor_f375577fcdb66a10) }

var fileDescriptor_f375577fcdb66a10 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0x4c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0xf7, 0x74, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce,
	0x4c, 0x29, 0x56, 0x52, 0xe7, 0xe2, 0x08, 0x86, 0xca, 0x08, 0x49, 0x73, 0x71, 0x66, 0xa6, 0xc4,
	0x43, 0x14, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x64, 0xa6, 0x40, 0xa4, 0x9d, 0xfa,
	0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x3d, 0x39, 0x3f, 0x57,
	0x2f, 0x33, 0xa5, 0xd8, 0x89, 0x1f, 0x66, 0x54, 0x98, 0x61, 0x00, 0xc8, 0x8a, 0x00, 0xc6, 0x28,
	0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4,
	0x12, 0xdf, 0xc4, 0xbc, 0x92, 0x9c, 0x54, 0xfd, 0xdc, 0xfc, 0x94, 0xd2, 0x9c, 0xd4, 0x62, 0x7d,
	0x34, 0x67, 0x2e, 0x62, 0x62, 0xf6, 0x8c, 0x88, 0x58, 0xc5, 0xc4, 0xec, 0x99, 0x52, 0x7c, 0x0a,
	0x4c, 0x3e, 0x62, 0xe2, 0xf7, 0x4c, 0x29, 0x8e, 0x71, 0x0f, 0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x4c,
	0x49, 0x2c, 0x49, 0x7c, 0x05, 0x16, 0x4f, 0x62, 0x03, 0xfb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x04, 0xa8, 0x8d, 0x12, 0xea, 0x00, 0x00, 0x00,
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
	if len(m.IdString) > 0 {
		i -= len(m.IdString)
		copy(dAtA[i:], m.IdString)
		i = encodeVarintStringIDV1(dAtA, i, uint64(len(m.IdString)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStringIDV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovStringIDV1(v)
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
	l = len(m.IdString)
	if l > 0 {
		n += 1 + l + sovStringIDV1(uint64(l))
	}
	return n
}

func sovStringIDV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStringIDV1(x uint64) (n int) {
	return sovStringIDV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStringIDV1
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
				return fmt.Errorf("proto: wrong wireType = %d for field IdString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStringIDV1
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
				return ErrInvalidLengthStringIDV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStringIDV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStringIDV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStringIDV1
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
func skipStringIDV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStringIDV1
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
					return 0, ErrIntOverflowStringIDV1
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
					return 0, ErrIntOverflowStringIDV1
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
				return 0, ErrInvalidLengthStringIDV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStringIDV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStringIDV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStringIDV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStringIDV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStringIDV1 = fmt.Errorf("proto: unexpected end of group")
)
