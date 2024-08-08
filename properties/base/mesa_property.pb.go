// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/schema/properties/base/mesa_property.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
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

type MesaProperty struct {
	ID     *base.PropertyID `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d,omitempty"`
	DataID *base.DataID     `protobuf:"bytes,2,opt,name=data_i_d,json=dataID,proto3" json:"data_i_d,omitempty"`
}

func (m *MesaProperty) Reset()         { *m = MesaProperty{} }
func (m *MesaProperty) String() string { return proto.CompactTextString(m) }
func (*MesaProperty) ProtoMessage()    {}
func (*MesaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd08199e90c8419, []int{0}
}
func (m *MesaProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MesaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MesaProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MesaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MesaProperty.Merge(m, src)
}
func (m *MesaProperty) XXX_Size() int {
	return m.Size()
}
func (m *MesaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_MesaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_MesaProperty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MesaProperty)(nil), "AssetMantle.schema.properties.base.MesaProperty")
}

func init() {
	proto.RegisterFile("AssetMantle/schema/properties/base/mesa_property.proto", fileDescriptor_6fd08199e90c8419)
}

var fileDescriptor_6fd08199e90c8419 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x73, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f,
	0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf,
	0x4d, 0x2d, 0x4e, 0x8c, 0x87, 0x0a, 0x56, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x29, 0x21,
	0xe9, 0xd3, 0x83, 0xe8, 0xd3, 0x43, 0xe8, 0xd3, 0x03, 0xe9, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x2b, 0xd7, 0x07, 0xb1, 0x20, 0x3a, 0xa5, 0x34, 0xb1, 0xd8, 0x98, 0x99, 0x02, 0xb5, 0x2a,
	0x25, 0xb1, 0x24, 0x31, 0x3e, 0x33, 0x05, 0xaa, 0x54, 0x17, 0x9f, 0x52, 0x98, 0x83, 0xe0, 0xca,
	0x95, 0x7a, 0x19, 0xb9, 0x78, 0x7c, 0x53, 0x8b, 0x13, 0x03, 0xa0, 0x32, 0x42, 0x16, 0x5c, 0xcc,
	0x99, 0xf1, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x58, 0x9c, 0x9c, 0x99,
	0x02, 0x71, 0xab, 0x1e, 0x4c, 0x8f, 0xa7, 0x4b, 0x10, 0x53, 0xa6, 0x8b, 0x90, 0x2d, 0x17, 0x07,
	0xc4, 0x29, 0xf1, 0x29, 0x12, 0x4c, 0x60, 0xed, 0xca, 0x78, 0xb5, 0xbb, 0x24, 0x96, 0x24, 0x7a,
	0xba, 0x04, 0xb1, 0xa5, 0x80, 0x69, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0xeb, 0x98, 0x4e,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x2d, 0x39, 0x3f, 0x57, 0x8f, 0x70,
	0x10, 0x3a, 0x09, 0x22, 0xfb, 0x27, 0x00, 0xe4, 0xcb, 0x00, 0xc6, 0x28, 0xdd, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc2, 0xd1, 0xb7, 0x88, 0x89, 0xc5, 0x31, 0x38,
	0xc0, 0x69, 0x15, 0x13, 0x4a, 0x94, 0x05, 0x43, 0xec, 0x0b, 0x40, 0xd8, 0xe7, 0x94, 0x58, 0x9c,
	0x7a, 0x0a, 0x45, 0x51, 0x0c, 0x44, 0x51, 0x0c, 0x42, 0x51, 0x0c, 0x48, 0xd1, 0x23, 0x26, 0x3d,
	0xc2, 0x8a, 0x62, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x41, 0x01, 0xf2, 0x8a, 0x49, 0x15,
	0x49, 0x83, 0x95, 0x15, 0x44, 0x87, 0x95, 0x15, 0x42, 0x8b, 0x95, 0x15, 0x48, 0x4f, 0x12, 0x1b,
	0x38, 0x1e, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x38, 0xc4, 0xd4, 0x95, 0x02, 0x00,
	0x00,
}

func (m *MesaProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MesaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MesaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DataID != nil {
		{
			size, err := m.DataID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMesaProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ID != nil {
		{
			size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMesaProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMesaProperty(dAtA []byte, offset int, v uint64) int {
	offset -= sovMesaProperty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MesaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != nil {
		l = m.ID.Size()
		n += 1 + l + sovMesaProperty(uint64(l))
	}
	if m.DataID != nil {
		l = m.DataID.Size()
		n += 1 + l + sovMesaProperty(uint64(l))
	}
	return n
}

func sovMesaProperty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMesaProperty(x uint64) (n int) {
	return sovMesaProperty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MesaProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesaProperty
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
			return fmt.Errorf("proto: MesaProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MesaProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesaProperty
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
				return ErrInvalidLengthMesaProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesaProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ID == nil {
				m.ID = &base.PropertyID{}
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesaProperty
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
				return ErrInvalidLengthMesaProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesaProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DataID == nil {
				m.DataID = &base.DataID{}
			}
			if err := m.DataID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesaProperty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMesaProperty
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
func skipMesaProperty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMesaProperty
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
					return 0, ErrIntOverflowMesaProperty
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
					return 0, ErrIntOverflowMesaProperty
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
				return 0, ErrInvalidLengthMesaProperty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMesaProperty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMesaProperty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMesaProperty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMesaProperty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMesaProperty = fmt.Errorf("proto: unexpected end of group")
)