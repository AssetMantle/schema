// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/ids/base/coinID.proto

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

type CoinID struct {
	StringID *StringID `protobuf:"bytes,1,opt,name=string_i_d,json=stringID,proto3" json:"string_i_d,omitempty"`
}

func (m *CoinID) Reset()         { *m = CoinID{} }
func (m *CoinID) String() string { return proto.CompactTextString(m) }
func (*CoinID) ProtoMessage()    {}
func (*CoinID) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a1d4d96aedbdcd, []int{0}
}
func (m *CoinID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinID.Merge(m, src)
}
func (m *CoinID) XXX_Size() int {
	return m.Size()
}
func (m *CoinID) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinID.DiscardUnknown(m)
}

var xxx_messageInfo_CoinID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CoinID)(nil), "ids.CoinID")
}

func init() { proto.RegisterFile("x/ids/base/coinID.proto", fileDescriptor_45a1d4d96aedbdcd) }

var fileDescriptor_45a1d4d96aedbdcd = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xaf, 0xd0, 0xcf, 0x4c,
	0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xce, 0xcf, 0xcc, 0xf3, 0x74, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x92, 0x44, 0x92, 0x2d, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x87, 0xc9, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44,
	0x54, 0xc9, 0x9a, 0x8b, 0xcd, 0x19, 0x6c, 0x8a, 0x90, 0x36, 0x17, 0x17, 0x44, 0x47, 0x7c, 0x66,
	0x7c, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xaf, 0x5e, 0x66, 0x4a, 0xb1, 0x5e, 0x30,
	0xd4, 0xa0, 0x20, 0x0e, 0x98, 0x91, 0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0x9c, 0x5a, 0x18, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x3d, 0x39, 0x3f, 0x17, 0xa4, 0xd9,
	0x89, 0x1b, 0x62, 0x7c, 0x00, 0xc8, 0xb6, 0x00, 0xc6, 0x28, 0x8d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x92, 0x9c,
	0x54, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x7d, 0x84, 0xf3, 0x17, 0x31, 0x31, 0x7b, 0x46,
	0x44, 0xac, 0x62, 0x62, 0xf6, 0x4c, 0x29, 0x3e, 0x05, 0x26, 0x1f, 0x31, 0xf1, 0x7b, 0xa6, 0x14,
	0xc7, 0xb8, 0x07, 0x38, 0xf9, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0xbe, 0x02, 0x8b, 0x27,
	0xb1, 0x81, 0xbd, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x23, 0x38, 0x0e, 0x1b, 0x01,
	0x00, 0x00,
}

func (m *CoinID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StringID != nil {
		{
			size, err := m.StringID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCoinID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoinID(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoinID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringID != nil {
		l = m.StringID.Size()
		n += 1 + l + sovCoinID(uint64(l))
	}
	return n
}

func sovCoinID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinID(x uint64) (n int) {
	return sovCoinID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoinID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinID
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
			return fmt.Errorf("proto: CoinID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinID
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
				return ErrInvalidLengthCoinID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StringID == nil {
				m.StringID = &StringID{}
			}
			if err := m.StringID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinID
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
func skipCoinID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinID
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
					return 0, ErrIntOverflowCoinID
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
					return 0, ErrIntOverflowCoinID
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
				return 0, ErrInvalidLengthCoinID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinID = fmt.Errorf("proto: unexpected end of group")
)