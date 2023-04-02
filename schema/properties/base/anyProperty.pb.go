// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/properties/base/anyProperty.proto

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

type AnyProperty struct {
	// Types that are valid to be assigned to Impl:
	//	*AnyProperty_MesaProperty
	//	*AnyProperty_MetaProperty
	Impl isAnyProperty_Impl `protobuf_oneof:"impl"`
}

func (m *AnyProperty) Reset()         { *m = AnyProperty{} }
func (m *AnyProperty) String() string { return proto.CompactTextString(m) }
func (*AnyProperty) ProtoMessage()    {}
func (*AnyProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b749c7d5cc16ef5a, []int{0}
}
func (m *AnyProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyProperty.Merge(m, src)
}
func (m *AnyProperty) XXX_Size() int {
	return m.Size()
}
func (m *AnyProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyProperty.DiscardUnknown(m)
}

var xxx_messageInfo_AnyProperty proto.InternalMessageInfo

type isAnyProperty_Impl interface {
	isAnyProperty_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AnyProperty_MesaProperty struct {
	MesaProperty *MesaProperty `protobuf:"bytes,1,opt,name=mesa_property,json=mesaProperty,proto3,oneof" json:"mesa_property,omitempty"`
}
type AnyProperty_MetaProperty struct {
	MetaProperty *MetaProperty `protobuf:"bytes,2,opt,name=meta_property,json=metaProperty,proto3,oneof" json:"meta_property,omitempty"`
}

func (*AnyProperty_MesaProperty) isAnyProperty_Impl() {}
func (*AnyProperty_MetaProperty) isAnyProperty_Impl() {}

func (m *AnyProperty) GetImpl() isAnyProperty_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *AnyProperty) GetMesaProperty() *MesaProperty {
	if x, ok := m.GetImpl().(*AnyProperty_MesaProperty); ok {
		return x.MesaProperty
	}
	return nil
}

func (m *AnyProperty) GetMetaProperty() *MetaProperty {
	if x, ok := m.GetImpl().(*AnyProperty_MetaProperty); ok {
		return x.MetaProperty
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnyProperty) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnyProperty_MesaProperty)(nil),
		(*AnyProperty_MetaProperty)(nil),
	}
}

func init() {
	proto.RegisterType((*AnyProperty)(nil), "properties.AnyProperty")
}

func init() {
	proto.RegisterFile("schema/properties/base/anyProperty.proto", fileDescriptor_b749c7d5cc16ef5a)
}

var fileDescriptor_b749c7d5cc16ef5a = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0x4a,
	0x2c, 0x4e, 0xd5, 0x4f, 0xcc, 0xab, 0x0c, 0x80, 0x08, 0x55, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x71, 0x21, 0x94, 0x48, 0x69, 0xe2, 0xd0, 0x95, 0x9b, 0x5a, 0x9c, 0x88, 0xaa, 0x0d, 0x8f,
	0xd2, 0x12, 0x74, 0xa5, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55,
	0x9a, 0xcb, 0xc8, 0xc5, 0xed, 0x88, 0x70, 0x8d, 0x90, 0x3d, 0x17, 0x2f, 0xc8, 0x9a, 0x78, 0xa8,
	0x81, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x12, 0x7a, 0x08, 0x1b, 0xf4, 0x7c, 0x91,
	0xdc, 0xe1, 0xc1, 0x10, 0xc4, 0x83, 0xec, 0x2e, 0x88, 0x01, 0x25, 0x48, 0x06, 0x30, 0x61, 0x33,
	0xa0, 0x04, 0xcd, 0x00, 0x04, 0xdf, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0x13, 0x1b, 0x17, 0x4b,
	0x66, 0x6e, 0x41, 0x8e, 0xd3, 0x4e, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0xe0, 0xe2, 0x4b, 0xce, 0xcf, 0x45, 0x32, 0xd5, 0x49, 0x00, 0xc9, 0x1f, 0x01, 0x20, 0xcf, 0x05,
	0x30, 0x46, 0x99, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16,
	0x17, 0xa7, 0x96, 0xf8, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4,
	0x16, 0xeb, 0x63, 0x0f, 0xbe, 0x45, 0x4c, 0xcc, 0x01, 0x11, 0x11, 0xab, 0x98, 0xb8, 0x02, 0xe0,
	0xe2, 0xa7, 0x90, 0x39, 0x8f, 0x98, 0xc4, 0x10, 0x9c, 0x18, 0xf7, 0x00, 0x27, 0x90, 0x87, 0x52,
	0x12, 0x4b, 0x12, 0x5f, 0x21, 0xab, 0x4a, 0x62, 0x03, 0x07, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0x3d, 0x72, 0x5b, 0x06, 0x02, 0x00, 0x00,
}

func (m *AnyProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Impl != nil {
		{
			size := m.Impl.Size()
			i -= size
			if _, err := m.Impl.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *AnyProperty_MesaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyProperty_MesaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MesaProperty != nil {
		{
			size, err := m.MesaProperty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *AnyProperty_MetaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyProperty_MetaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MetaProperty != nil {
		{
			size, err := m.MetaProperty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintAnyProperty(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnyProperty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnyProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Impl != nil {
		n += m.Impl.Size()
	}
	return n
}

func (m *AnyProperty_MesaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MesaProperty != nil {
		l = m.MesaProperty.Size()
		n += 1 + l + sovAnyProperty(uint64(l))
	}
	return n
}
func (m *AnyProperty_MetaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaProperty != nil {
		l = m.MetaProperty.Size()
		n += 1 + l + sovAnyProperty(uint64(l))
	}
	return n
}

func sovAnyProperty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnyProperty(x uint64) (n int) {
	return sovAnyProperty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnyProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnyProperty
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
			return fmt.Errorf("proto: AnyProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MesaProperty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyProperty
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
				return ErrInvalidLengthAnyProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MesaProperty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyProperty_MesaProperty{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaProperty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyProperty
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
				return ErrInvalidLengthAnyProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MetaProperty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyProperty_MetaProperty{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnyProperty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnyProperty
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
func skipAnyProperty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnyProperty
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
					return 0, ErrIntOverflowAnyProperty
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
					return 0, ErrIntOverflowAnyProperty
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
				return 0, ErrInvalidLengthAnyProperty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnyProperty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnyProperty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnyProperty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnyProperty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnyProperty = fmt.Errorf("proto: unexpected end of group")
)
