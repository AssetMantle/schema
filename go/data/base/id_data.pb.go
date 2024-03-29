// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data/base/id_data.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
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

type IDData struct {
	Value *base.AnyID `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *IDData) Reset()         { *m = IDData{} }
func (m *IDData) String() string { return proto.CompactTextString(m) }
func (*IDData) ProtoMessage()    {}
func (*IDData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc4f319ae578334, []int{0}
}
func (m *IDData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDData.Merge(m, src)
}
func (m *IDData) XXX_Size() int {
	return m.Size()
}
func (m *IDData) XXX_DiscardUnknown() {
	xxx_messageInfo_IDData.DiscardUnknown(m)
}

var xxx_messageInfo_IDData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IDData)(nil), "assetmantle.schema.data.base.IDData")
}

func init() { proto.RegisterFile("data/base/id_data.proto", fileDescriptor_8dc4f319ae578334) }

var fileDescriptor_8dc4f319ae578334 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0x2c, 0x49,
	0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4c, 0x89, 0x07, 0x71, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x64, 0x12, 0x8b, 0x8b, 0x53, 0x4b, 0x72, 0x13, 0xf3, 0x4a, 0x72, 0x52, 0xf5, 0x8a,
	0x93, 0x33, 0x52, 0x73, 0x13, 0xf5, 0xc0, 0xd2, 0x20, 0xb5, 0x52, 0xa2, 0x99, 0x29, 0xc5, 0x10,
	0x5d, 0x89, 0x79, 0x95, 0xf1, 0x99, 0x29, 0x10, 0x4d, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60,
	0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0xf2, 0xe0, 0x62, 0xf3, 0x74, 0x71, 0x49, 0x2c, 0x49, 0x14,
	0xb2, 0xe0, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52,
	0xd2, 0xc3, 0x62, 0x49, 0x66, 0x4a, 0x31, 0xd8, 0x0e, 0x3d, 0xc7, 0xbc, 0x4a, 0x4f, 0x97, 0x20,
	0x88, 0x06, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0xed, 0x4c, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0xa5, 0x90, 0x9c, 0x9f, 0xab, 0x87, 0xcf, 0xcd, 0x4e, 0xdc, 0x9e,
	0x29, 0x20, 0x47, 0x04, 0x80, 0xdc, 0x14, 0xc0, 0x18, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x08, 0xd2, 0xe7, 0x0b, 0xd6, 0xa7, 0x0f, 0xd1, 0xa7, 0x9f,
	0x9e, 0xaf, 0x0f, 0x0f, 0x9a, 0x45, 0x4c, 0x2c, 0x8e, 0xc1, 0x2e, 0x4e, 0xab, 0x98, 0x64, 0x1c,
	0x91, 0x2c, 0x08, 0x86, 0x58, 0x00, 0x32, 0x53, 0xcf, 0x29, 0xb1, 0x38, 0xf5, 0x14, 0x8a, 0x74,
	0x0c, 0x44, 0x3a, 0x06, 0x24, 0x1d, 0x03, 0x92, 0x7e, 0xc4, 0xa4, 0x81, 0x4f, 0x3a, 0xc6, 0x3d,
	0xc0, 0xc9, 0x37, 0xb5, 0x24, 0x11, 0x64, 0xe9, 0x2b, 0x26, 0x79, 0x24, 0xa5, 0x56, 0x56, 0x10,
	0xb5, 0x56, 0x56, 0x20, 0xc5, 0x56, 0x56, 0x20, 0xd5, 0x49, 0x6c, 0xe0, 0xa0, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x22, 0xd0, 0x79, 0xc0, 0x01, 0x00, 0x00,
}

func (m *IDData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdData(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IDData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovIdData(uint64(l))
	}
	return n
}

func sovIdData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdData(x uint64) (n int) {
	return sovIdData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IDData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdData
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
			return fmt.Errorf("proto: IDData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdData
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
				return ErrInvalidLengthIdData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &base.AnyID{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdData
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
func skipIdData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdData
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
					return 0, ErrIntOverflowIdData
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
					return 0, ErrIntOverflowIdData
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
				return 0, ErrInvalidLengthIdData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdData = fmt.Errorf("proto: unexpected end of group")
)
