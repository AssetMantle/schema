// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/data/base/accAddressData.proto

package base

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type AccAddressData struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AccAddressData) Reset()         { *m = AccAddressData{} }
func (m *AccAddressData) String() string { return proto.CompactTextString(m) }
func (*AccAddressData) ProtoMessage()    {}
func (*AccAddressData) Descriptor() ([]byte, []int) {
	return fileDescriptor_986f96c4a650b90f, []int{0}
}
func (m *AccAddressData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccAddressData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccAddressData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccAddressData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccAddressData.Merge(m, src)
}
func (m *AccAddressData) XXX_Size() int {
	return m.Size()
}
func (m *AccAddressData) XXX_DiscardUnknown() {
	xxx_messageInfo_AccAddressData.DiscardUnknown(m)
}

var xxx_messageInfo_AccAddressData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccAddressData)(nil), "data.AccAddressData")
}

func init() { proto.RegisterFile("x/data/base/accAddressData.proto", fileDescriptor_986f96c4a650b90f) }

var fileDescriptor_986f96c4a650b90f = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xa8, 0xd0, 0x4f, 0x49,
	0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0x4c, 0x4e, 0x76, 0x4c, 0x49, 0x29, 0x4a,
	0x2d, 0x2e, 0x76, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9,
	0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x05, 0xf4, 0x41, 0x2c, 0x88, 0x9c, 0x92, 0x0e, 0x17,
	0x9f, 0x23, 0x8a, 0x1e, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x08, 0xc7, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0xd3, 0x24, 0xc6, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0xe2, 0x48, 0xce, 0xcf, 0xd5, 0x03, 0x59,
	0xe4, 0x24, 0x8c, 0x6a, 0x60, 0x00, 0xc8, 0x9e, 0x00, 0xc6, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x92,
	0x9c, 0x54, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x7d, 0x24, 0x5f, 0x2c, 0x62, 0x62, 0x76,
	0x89, 0x88, 0x58, 0xc5, 0xc4, 0x02, 0xd2, 0x7e, 0x0a, 0x42, 0x3d, 0x62, 0x12, 0x00, 0x51, 0x31,
	0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x20, 0xb5, 0xaf, 0x20, 0x32, 0x49, 0x6c, 0x60, 0x9f,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x9b, 0x8d, 0xf6, 0x09, 0x01, 0x00, 0x00,
}

func (m *AccAddressData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccAddressData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccAddressData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintAccAddressData(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccAddressData(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccAddressData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccAddressData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovAccAddressData(uint64(l))
	}
	return n
}

func sovAccAddressData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccAddressData(x uint64) (n int) {
	return sovAccAddressData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccAddressData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccAddressData
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
			return fmt.Errorf("proto: AccAddressData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccAddressData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccAddressData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccAddressData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccAddressData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccAddressData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccAddressData
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
func skipAccAddressData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccAddressData
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
					return 0, ErrIntOverflowAccAddressData
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
					return 0, ErrIntOverflowAccAddressData
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
				return 0, ErrInvalidLengthAccAddressData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccAddressData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccAddressData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccAddressData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccAddressData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccAddressData = fmt.Errorf("proto: unexpected end of group")
)
