// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/base/parameter.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/properties/base"
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

type Parameter struct {
	MetaProperty *base.MetaProperty `protobuf:"bytes,1,opt,name=meta_property,json=metaProperty,proto3" json:"meta_property,omitempty"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6471016db93f7d55, []int{0}
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
	proto.RegisterType((*Parameter)(nil), "assetmantle.schema.parameters.base.Parameter")
}

func init() { proto.RegisterFile("types/base/parameter.proto", fileDescriptor_6471016db93f7d55) }

var fileDescriptor_6471016db93f7d55 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4a, 0x2c, 0x2e, 0x4e, 0x2d, 0xc9, 0x4d, 0xcc,
	0x2b, 0xc9, 0x49, 0xd5, 0x2b, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x83, 0xab, 0x29, 0xd6, 0x03,
	0xe9, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd7, 0x07, 0xb1, 0x20, 0x3a, 0xa5, 0x94,
	0x0b, 0x8a, 0xf2, 0x0b, 0x52, 0x8b, 0x4a, 0x32, 0x61, 0x46, 0xe7, 0xa6, 0x96, 0x24, 0xc6, 0x43,
	0x05, 0x2b, 0x21, 0x8a, 0x94, 0x32, 0xb8, 0x38, 0x03, 0x60, 0xa6, 0x09, 0x85, 0x72, 0xf1, 0xa2,
	0xa8, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x32, 0xd0, 0xc3, 0xe6, 0x06, 0xb8, 0xe1, 0x60,
	0x37, 0xe8, 0xf9, 0xa6, 0x96, 0x24, 0x06, 0x40, 0xf5, 0x05, 0xf1, 0xe4, 0x22, 0xf1, 0xac, 0x58,
	0x3a, 0x16, 0xc8, 0x33, 0x38, 0xad, 0x64, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x06, 0x2e, 0xb5, 0xe4, 0xfc, 0x5c, 0x3d, 0xc2, 0xfe, 0x74, 0xe2, 0x83, 0x3b, 0x35, 0x00, 0xe4,
	0xf8, 0x00, 0xc6, 0x28, 0xed, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0x47, 0x90, 0x01, 0xbe, 0x60, 0x03, 0xf4, 0x21, 0x06, 0xe8, 0xa7, 0xe7, 0xeb, 0x23, 0x82, 0x76,
	0x11, 0x13, 0x8b, 0x63, 0x70, 0x80, 0xd3, 0x2a, 0x26, 0x25, 0x47, 0x24, 0xab, 0x82, 0x21, 0x56,
	0x05, 0x20, 0xac, 0x72, 0x4a, 0x2c, 0x4e, 0x3d, 0x85, 0xa2, 0x28, 0x06, 0xa2, 0x28, 0x06, 0xa1,
	0x28, 0x06, 0xa4, 0xe8, 0x11, 0x93, 0x1e, 0x61, 0x45, 0x31, 0xee, 0x01, 0x4e, 0xa0, 0xb0, 0x49,
	0x49, 0x2c, 0x49, 0x7c, 0xc5, 0xa4, 0x8a, 0xa4, 0xc1, 0xca, 0x0a, 0xa2, 0xc3, 0xca, 0x0a, 0xa1,
	0xc5, 0xca, 0x0a, 0xa4, 0x27, 0x89, 0x0d, 0x1c, 0x39, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x07, 0x47, 0xba, 0x35, 0x19, 0x02, 0x00, 0x00,
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
