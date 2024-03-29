// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: properties/base/any_property.proto

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
	return fileDescriptor_5bcb949bd5ee089b, []int{0}
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
	proto.RegisterType((*AnyProperty)(nil), "assetmantle.schema.properties.base.AnyProperty")
}

func init() {
	proto.RegisterFile("properties/base/any_property.proto", fileDescriptor_5bcb949bd5ee089b)
}

var fileDescriptor_5bcb949bd5ee089b = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xcc, 0xab, 0x8c,
	0x87, 0x8a, 0x55, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x29, 0x25, 0x16, 0x17, 0xa7, 0x96,
	0xe4, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0x15, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0x21, 0xb4,
	0xe9, 0x81, 0xb4, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x95, 0xeb, 0x83, 0x58, 0x10, 0x9d,
	0x52, 0xca, 0xe8, 0xa6, 0xe7, 0xa6, 0x16, 0x27, 0xa2, 0x19, 0x8f, 0x4d, 0x51, 0x09, 0xba, 0x22,
	0xa5, 0xb3, 0x8c, 0x5c, 0xdc, 0x8e, 0x79, 0x95, 0x01, 0x50, 0x51, 0xa1, 0x70, 0x2e, 0x5e, 0x14,
	0xb3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x0c, 0xf4, 0x08, 0xbb, 0x55, 0xcf, 0x37, 0xb5,
	0x38, 0x11, 0x66, 0x90, 0x07, 0x43, 0x10, 0x4f, 0x2e, 0x12, 0x1f, 0x62, 0x30, 0x92, 0xfd, 0x12,
	0x4c, 0xa4, 0x18, 0x5c, 0x82, 0x66, 0x30, 0x82, 0x6f, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0x89,
	0x8d, 0x8b, 0x25, 0x33, 0xb7, 0x20, 0xc7, 0x69, 0x03, 0xd3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x70, 0xa9, 0x25, 0xe7, 0xe7, 0x12, 0x61, 0x9b, 0x93, 0x00, 0x52, 0x78, 0x04,
	0x80, 0x02, 0x29, 0x80, 0x31, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0xdf, 0x11, 0x64, 0x84, 0x2f, 0xd8, 0x08, 0x7d, 0x88, 0x11, 0xfa, 0xe9, 0xf9, 0xfa, 0x68,
	0x81, 0xbd, 0x88, 0x89, 0xc5, 0x31, 0x38, 0xc0, 0x69, 0x15, 0x93, 0x92, 0x23, 0x92, 0x8d, 0xc1,
	0x10, 0x1b, 0x03, 0x10, 0x36, 0x3a, 0x25, 0x16, 0xa7, 0x9e, 0x42, 0x51, 0x14, 0x03, 0x51, 0x14,
	0x83, 0x50, 0x14, 0x03, 0x52, 0xf4, 0x88, 0x49, 0x8f, 0xb0, 0xa2, 0x18, 0xf7, 0x00, 0x27, 0x50,
	0x60, 0xa5, 0x24, 0x96, 0x24, 0xbe, 0x62, 0x52, 0x45, 0xd2, 0x60, 0x65, 0x05, 0xd1, 0x61, 0x65,
	0x85, 0xd0, 0x62, 0x65, 0x05, 0xd2, 0x93, 0xc4, 0x06, 0x4e, 0x09, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x80, 0x32, 0x83, 0xb3, 0x02, 0x00, 0x00,
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
