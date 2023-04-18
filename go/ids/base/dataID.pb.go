// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/dataID.proto

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

type DataID struct {
	TypeID *StringID `protobuf:"bytes,1,opt,name=type_i_d,json=typeID,proto3" json:"type_i_d,omitempty"`
	HashID *HashID   `protobuf:"bytes,2,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d,omitempty"`
}

func (m *DataID) Reset()         { *m = DataID{} }
func (m *DataID) String() string { return proto.CompactTextString(m) }
func (*DataID) ProtoMessage()    {}
func (*DataID) Descriptor() ([]byte, []int) {
	return fileDescriptor_35f708228d7fc875, []int{0}
}
func (m *DataID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataID.Merge(m, src)
}
func (m *DataID) XXX_Size() int {
	return m.Size()
}
func (m *DataID) XXX_DiscardUnknown() {
	xxx_messageInfo_DataID.DiscardUnknown(m)
}

var xxx_messageInfo_DataID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DataID)(nil), "ids.DataID")
}

func init() { proto.RegisterFile("ids/base/dataID.proto", fileDescriptor_35f708228d7fc875) }

var fileDescriptor_35f708228d7fc875 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xf4, 0x74, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x42, 0xc8, 0x65, 0x24, 0x16, 0x67, 0xc0, 0xe4, 0xa4,
	0xc4, 0xe1, 0xc2, 0xc5, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x70, 0x09, 0x91, 0xf4, 0xfc, 0xf4, 0x7c,
	0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0xc5, 0x71, 0xb1, 0xb9, 0x80, 0x8d, 0x16, 0x52, 0xe7,
	0xe2, 0x28, 0xa9, 0x2c, 0x48, 0x8d, 0xcf, 0x8c, 0x4f, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36,
	0xe2, 0xd5, 0xcb, 0x4c, 0x29, 0xd6, 0x0b, 0x86, 0x1a, 0x13, 0xc4, 0x06, 0x92, 0xf6, 0x74, 0x11,
	0x52, 0xe5, 0xe2, 0x00, 0xd9, 0x08, 0x56, 0xc8, 0x04, 0x56, 0xc8, 0x0d, 0x56, 0xe8, 0x01, 0x76,
	0x46, 0x10, 0x1b, 0xc4, 0x39, 0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0x9c, 0x5a, 0x19, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x3d, 0x39, 0x3f, 0x17, 0xa4, 0xd1, 0x89,
	0x1b, 0xe2, 0x82, 0x00, 0x90, 0x83, 0x02, 0x18, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x8b, 0x8b, 0x53, 0x4b, 0x7c, 0x13, 0xf3, 0x4a, 0x72, 0x52,
	0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0xf5, 0xd3, 0xf3, 0xf5, 0x61, 0x5e, 0x5c, 0xc4, 0xc4,
	0xec, 0x19, 0x11, 0xb1, 0x8a, 0x89, 0xd9, 0x33, 0xa5, 0xf8, 0x14, 0x98, 0x7c, 0xc4, 0xc4, 0xef,
	0x99, 0x52, 0x1c, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x08, 0x0a, 0xb4, 0x57, 0x60, 0xf1,
	0x24, 0x36, 0xb0, 0x77, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xb6, 0xe4, 0x8f, 0x52,
	0x01, 0x00, 0x00,
}

func (m *DataID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HashID != nil {
		{
			size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TypeID != nil {
		{
			size, err := m.TypeID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataID(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypeID != nil {
		l = m.TypeID.Size()
		n += 1 + l + sovDataID(uint64(l))
	}
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovDataID(uint64(l))
	}
	return n
}

func sovDataID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataID(x uint64) (n int) {
	return sovDataID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataID
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
			return fmt.Errorf("proto: DataID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataID
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
				return ErrInvalidLengthDataID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataID
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataID
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
				return ErrInvalidLengthDataID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HashID == nil {
				m.HashID = &HashID{}
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataID
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
func skipDataID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataID
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
					return 0, ErrIntOverflowDataID
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
					return 0, ErrIntOverflowDataID
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
				return 0, ErrInvalidLengthDataID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataID = fmt.Errorf("proto: unexpected end of group")
)
