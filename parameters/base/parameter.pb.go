// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/schema/parameters/base/parameter.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/properties/base"
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

type Parameter struct {
	MetaProperty *base.MetaProperty `protobuf:"bytes,1,opt,name=meta_property,json=metaProperty,proto3" json:"meta_property,omitempty"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_561716c784d0ccda, []int{0}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return m.Size()
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Parameter)(nil), "AssetMantle.schema.parameters.base.Parameter")
}

func init() {
	proto.RegisterFile("AssetMantle/schema/parameters/base/parameter.proto", fileDescriptor_561716c784d0ccda)
}

var fileDescriptor_561716c784d0ccda = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x72, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f,
	0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x2a, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0x45, 0xf0,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x94, 0x90, 0xf4, 0xe8, 0x41, 0xf4, 0xe8, 0x21, 0xf4,
	0xe8, 0x81, 0xf4, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x95, 0xeb, 0x83, 0x58, 0x10, 0x9d,
	0x52, 0x66, 0xd8, 0x6c, 0x2b, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x85, 0xda, 0x96, 0x9b,
	0x5a, 0x92, 0x18, 0x0f, 0x15, 0xac, 0x84, 0xe8, 0x53, 0xca, 0xe0, 0xe2, 0x0c, 0x80, 0x59, 0x20,
	0x14, 0xca, 0xc5, 0x8b, 0xa2, 0x46, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0xc8, 0x40, 0x0f, 0x9b,
	0xb3, 0xe0, 0x86, 0x83, 0x9d, 0xa5, 0xe7, 0x9b, 0x5a, 0x92, 0x18, 0x00, 0xd5, 0x17, 0xc4, 0x93,
	0x8b, 0xc4, 0xb3, 0x62, 0xe9, 0x58, 0x20, 0xcf, 0xe0, 0xb4, 0x9a, 0xe9, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0xd4, 0x92, 0xf3, 0x73, 0xf5, 0x08, 0x7b, 0xdd, 0x89, 0x0f,
	0xee, 0xd4, 0x00, 0x90, 0xe3, 0x03, 0x18, 0xa3, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x09, 0x87, 0xf7, 0x22, 0x26, 0x16, 0xc7, 0xe0, 0x00, 0xa7, 0x55, 0x4c,
	0x28, 0xe1, 0x1c, 0x0c, 0xb1, 0x2c, 0x00, 0x61, 0x99, 0x53, 0x62, 0x71, 0xea, 0x29, 0x14, 0x45,
	0x31, 0x10, 0x45, 0x31, 0x08, 0x45, 0x31, 0x20, 0x45, 0x8f, 0x98, 0xf4, 0x08, 0x2b, 0x8a, 0x71,
	0x0f, 0x70, 0x02, 0x85, 0x4e, 0x4a, 0x62, 0x49, 0xe2, 0x2b, 0x26, 0x55, 0x24, 0x0d, 0x56, 0x56,
	0x10, 0x1d, 0x56, 0x56, 0x08, 0x2d, 0x56, 0x56, 0x20, 0x3d, 0x49, 0x6c, 0xe0, 0xe8, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xee, 0x8d, 0x28, 0x46, 0x02, 0x00, 0x00,
}

func (m *Parameter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Parameter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Parameter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MetaProperty != nil {
		{
			size, err := m.MetaProperty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParameter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParameter(dAtA []byte, offset int, v uint64) int {
	offset -= sovParameter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Parameter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaProperty != nil {
		l = m.MetaProperty.Size()
		n += 1 + l + sovParameter(uint64(l))
	}
	return n
}

func sovParameter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParameter(x uint64) (n int) {
	return sovParameter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Parameter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParameter
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
			return fmt.Errorf("proto: Parameter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Parameter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaProperty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameter
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
				return ErrInvalidLengthParameter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaProperty == nil {
				m.MetaProperty = &base.MetaProperty{}
			}
			if err := m.MetaProperty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParameter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParameter
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
func skipParameter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
				return 0, ErrInvalidLengthParameter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParameter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParameter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParameter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParameter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParameter = fmt.Errorf("proto: unexpected end of group")
)
