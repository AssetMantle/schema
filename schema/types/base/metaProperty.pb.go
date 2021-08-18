// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/metaProperty.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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

type MetaProperty struct {
	Id       github_com_persistenceOne_persistenceSDK_schema_types.ID       `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"id"`
	MetaFact github_com_persistenceOne_persistenceSDK_schema_types.MetaFact `protobuf:"bytes,2,opt,name=metaFact,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaFact" json:"metaFact"`
}

func (m *MetaProperty) Reset()         { *m = MetaProperty{} }
func (m *MetaProperty) String() string { return proto.CompactTextString(m) }
func (*MetaProperty) ProtoMessage()    {}
func (*MetaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00c620db06edd3e, []int{0}
}
func (m *MetaProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetaProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaProperty.Merge(m, src)
}
func (m *MetaProperty) XXX_Size() int {
	return m.Size()
}
func (m *MetaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_MetaProperty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetaProperty)(nil), "persistence_sdk.schema.types.base.MetaProperty")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/metaProperty.proto", fileDescriptor_e00c620db06edd3e)
}

var fileDescriptor_e00c620db06edd3e = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4d,
	0x2d, 0x49, 0x0c, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x44, 0xd3, 0xa5, 0x07, 0xd1, 0xa5, 0x07, 0xd6, 0xa5, 0x07, 0xd2, 0x25, 0x25, 0x92,
	0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xad, 0x0f, 0x62, 0x41, 0x34, 0x2a, 0x1d, 0x61, 0xe4, 0xe2, 0xf1,
	0x45, 0x32, 0x4f, 0x28, 0x80, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0xc9,
	0xe1, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x2d, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x91, 0x2c, 0xf2, 0xcf, 0x4b, 0x45, 0xe6, 0x06, 0xbb, 0x78, 0xa3, 0x38,
	0x56, 0xcf, 0xd3, 0x25, 0x88, 0x29, 0x33, 0x45, 0x28, 0x89, 0x8b, 0x03, 0xe4, 0x62, 0xb7, 0xc4,
	0xe4, 0x12, 0x09, 0x26, 0xb0, 0xb9, 0x6e, 0x50, 0x73, 0xed, 0xc8, 0x33, 0xd7, 0x17, 0x6a, 0x5a,
	0x10, 0xdc, 0x5c, 0xa7, 0x90, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2,
	0x22, 0xcb, 0x0e, 0x70, 0x40, 0x27, 0xb1, 0x81, 0xc3, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0xc3, 0x20, 0xe6, 0x94, 0x01, 0x00, 0x00,
}

func (m *MetaProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MetaFact.Size()
		i -= size
		if _, err := m.MetaFact.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMetaProperty(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMetaProperty(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMetaProperty(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetaProperty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovMetaProperty(uint64(l))
	l = m.MetaFact.Size()
	n += 1 + l + sovMetaProperty(uint64(l))
	return n
}

func sovMetaProperty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetaProperty(x uint64) (n int) {
	return sovMetaProperty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetaProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetaProperty
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
			return fmt.Errorf("proto: MetaProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetaProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaProperty
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
				return ErrInvalidLengthMetaProperty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetaProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaFact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaProperty
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
				return ErrInvalidLengthMetaProperty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetaProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MetaFact.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetaProperty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetaProperty
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
func skipMetaProperty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetaProperty
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
					return 0, ErrIntOverflowMetaProperty
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
					return 0, ErrIntOverflowMetaProperty
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
				return 0, ErrInvalidLengthMetaProperty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetaProperty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetaProperty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetaProperty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetaProperty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetaProperty = fmt.Errorf("proto: unexpected end of group")
)
