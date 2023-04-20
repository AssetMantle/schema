// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/property_id.proto

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

type PropertyID struct {
	KeyID  *StringID `protobuf:"bytes,1,opt,name=key_i_d,json=keyID,proto3" json:"key_i_d,omitempty"`
	TypeID *StringID `protobuf:"bytes,2,opt,name=type_i_d,json=typeID,proto3" json:"type_i_d,omitempty"`
}

func (m *PropertyID) Reset()         { *m = PropertyID{} }
func (m *PropertyID) String() string { return proto.CompactTextString(m) }
func (*PropertyID) ProtoMessage()    {}
func (*PropertyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ba70a7d0ce7006, []int{0}
}
func (m *PropertyID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PropertyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PropertyID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PropertyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropertyID.Merge(m, src)
}
func (m *PropertyID) XXX_Size() int {
	return m.Size()
}
func (m *PropertyID) XXX_DiscardUnknown() {
	xxx_messageInfo_PropertyID.DiscardUnknown(m)
}

var xxx_messageInfo_PropertyID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PropertyID)(nil), "ids.base.PropertyID")
}

func init() { proto.RegisterFile("ids/base/property_id.proto", fileDescriptor_06ba70a7d0ce7006) }

var fileDescriptor_06ba70a7d0ce7006 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0x8c, 0xcf, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4c, 0x29, 0xd6, 0x03, 0xc9, 0x49, 0x49,
	0xc0, 0x55, 0x15, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0xc3, 0xd5, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x29, 0x87, 0x8b, 0x2b, 0x00, 0x6a, 0x9c, 0xa7, 0x8b,
	0x90, 0x16, 0x17, 0x7b, 0x76, 0x6a, 0x65, 0x7c, 0x66, 0x7c, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0x90, 0x1e, 0xcc, 0x64, 0xbd, 0x60, 0xb0, 0x79, 0x9e, 0x2e, 0x41, 0xac, 0xd9, 0xa9,
	0x20, 0xb5, 0x3a, 0x5c, 0x1c, 0x25, 0x95, 0x05, 0xa9, 0x60, 0xc5, 0x4c, 0x38, 0x15, 0xb3, 0x81,
	0xd4, 0x78, 0xba, 0x58, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0x5a, 0xcc, 0x78, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0x3c, 0xc9, 0xf9, 0xb9, 0x70, 0xfd, 0x4e, 0xfc, 0x70,
	0x47, 0xa5, 0x04, 0x80, 0xdc, 0x19, 0xc0, 0x18, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57, 0x92, 0x93, 0xaa,
	0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x9f, 0x9e, 0xaf, 0x0f, 0xf3, 0xfa, 0x22, 0x26, 0x66,
	0x4f, 0xa7, 0x88, 0x55, 0x4c, 0x1c, 0x9e, 0x29, 0xc5, 0x7a, 0x4e, 0x89, 0xc5, 0xa9, 0xa7, 0xc0,
	0xcc, 0x18, 0x10, 0xf3, 0x11, 0x93, 0x08, 0x8c, 0x19, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92,
	0x98, 0x92, 0x58, 0x92, 0xf8, 0x8a, 0x89, 0xd3, 0x33, 0xa5, 0xd8, 0xca, 0x0a, 0x24, 0x9e, 0xc4,
	0x06, 0x0e, 0x1a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x3f, 0x77, 0x89, 0x72, 0x01,
	0x00, 0x00,
}

func (m *PropertyID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PropertyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PropertyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TypeID != nil {
		{
			size, err := m.TypeID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPropertyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.KeyID != nil {
		{
			size, err := m.KeyID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPropertyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPropertyId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPropertyId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PropertyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeyID != nil {
		l = m.KeyID.Size()
		n += 1 + l + sovPropertyId(uint64(l))
	}
	if m.TypeID != nil {
		l = m.TypeID.Size()
		n += 1 + l + sovPropertyId(uint64(l))
	}
	return n
}

func sovPropertyId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPropertyId(x uint64) (n int) {
	return sovPropertyId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PropertyID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPropertyId
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
			return fmt.Errorf("proto: PropertyID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PropertyID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPropertyId
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
				return ErrInvalidLengthPropertyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPropertyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.KeyID == nil {
				m.KeyID = &StringID{}
			}
			if err := m.KeyID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPropertyId
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
				return ErrInvalidLengthPropertyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPropertyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TypeID == nil {
				m.TypeID = &StringID{}
			}
			if err := m.TypeID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPropertyId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPropertyId
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
func skipPropertyId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPropertyId
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
					return 0, ErrIntOverflowPropertyId
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
					return 0, ErrIntOverflowPropertyId
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
				return 0, ErrInvalidLengthPropertyId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPropertyId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPropertyId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPropertyId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPropertyId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPropertyId = fmt.Errorf("proto: unexpected end of group")
)
