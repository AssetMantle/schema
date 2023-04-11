// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/data/base/heightData.proto

package base

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/x/types/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type HeightData struct {
	Value *base.Height `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *HeightData) Reset()         { *m = HeightData{} }
func (m *HeightData) String() string { return proto.CompactTextString(m) }
func (*HeightData) ProtoMessage()    {}
func (*HeightData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8536cb517979678f, []int{0}
}
func (m *HeightData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeightData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeightData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeightData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeightData.Merge(m, src)
}
func (m *HeightData) XXX_Size() int {
	return m.Size()
}
func (m *HeightData) XXX_DiscardUnknown() {
	xxx_messageInfo_HeightData.DiscardUnknown(m)
}

var xxx_messageInfo_HeightData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HeightData)(nil), "data.HeightData")
}

func init() { proto.RegisterFile("x/data/base/heightData.proto", fileDescriptor_8536cb517979678f) }

var fileDescriptor_8536cb517979678f = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xa9, 0xd0, 0x4f, 0x49,
	0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x71, 0x49,
	0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9, 0x49, 0x49, 0x56, 0xe8,
	0x97, 0x54, 0x16, 0xa4, 0x16, 0x23, 0x2b, 0x82, 0x28, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x92, 0x39, 0x17, 0x97, 0x07, 0xdc, 0x28, 0x21, 0x65, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x5e, 0x3d, 0xb0,
	0x61, 0x7a, 0x10, 0x15, 0x41, 0x10, 0x39, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0x7d, 0x8c,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0xc5, 0x91, 0x9c, 0x9f, 0xab, 0x07,
	0x72, 0x8e, 0x13, 0x3f, 0xc2, 0xec, 0x00, 0x90, 0x75, 0x01, 0x8c, 0x51, 0x9a, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x8e, 0xc5, 0xc5, 0xa9, 0x25, 0xbe, 0x89, 0x79,
	0x25, 0x39, 0xa9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0xfa, 0x48, 0x7e, 0x5c, 0xc4, 0xc4,
	0xec, 0x12, 0x11, 0xb1, 0x8a, 0x89, 0x05, 0xa4, 0xfd, 0x14, 0x84, 0x7a, 0xc4, 0x24, 0x00, 0xa2,
	0x62, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x41, 0x6a, 0x5f, 0x41, 0x64, 0x92, 0xd8, 0xc0,
	0x1e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xa8, 0xfc, 0x87, 0x27, 0x01, 0x00, 0x00,
}

func (m *HeightData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeightData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeightData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintHeightData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeightData(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeightData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeightData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovHeightData(uint64(l))
	}
	return n
}

func sovHeightData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeightData(x uint64) (n int) {
	return sovHeightData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeightData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeightData
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
			return fmt.Errorf("proto: HeightData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeightData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeightData
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
				return ErrInvalidLengthHeightData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeightData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &base.Height{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeightData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHeightData
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
func skipHeightData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeightData
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
					return 0, ErrIntOverflowHeightData
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
					return 0, ErrIntOverflowHeightData
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
				return 0, ErrInvalidLengthHeightData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeightData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeightData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeightData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeightData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeightData = fmt.Errorf("proto: unexpected end of group")
)
