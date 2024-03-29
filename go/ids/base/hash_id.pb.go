// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/hash_id.proto

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

type HashID struct {
	IDBytes []byte `protobuf:"bytes,1,opt,name=i_d_bytes,json=iDBytes,proto3" json:"i_d_bytes,omitempty"`
}

func (m *HashID) Reset()         { *m = HashID{} }
func (m *HashID) String() string { return proto.CompactTextString(m) }
func (*HashID) ProtoMessage()    {}
func (*HashID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6241413c9a775af3, []int{0}
}
func (m *HashID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashID.Merge(m, src)
}
func (m *HashID) XXX_Size() int {
	return m.Size()
}
func (m *HashID) XXX_DiscardUnknown() {
	xxx_messageInfo_HashID.DiscardUnknown(m)
}

var xxx_messageInfo_HashID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HashID)(nil), "assetmantle.schema.ids.base.HashID")
}

func init() { proto.RegisterFile("ids/base/hash_id.proto", fileDescriptor_6241413c9a775af3) }

var fileDescriptor_6241413c9a775af3 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x48, 0x2c, 0xce, 0x88, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x4e, 0x2c, 0x2e, 0x4e, 0x2d, 0xc9, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5,
	0x2b, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0xcb, 0x4c, 0x29, 0xd6, 0x03, 0x29, 0x95, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd3, 0x07, 0xb1, 0x20, 0x5a, 0x94, 0xb4, 0xb8, 0xd8, 0x3c, 0x12,
	0x8b, 0x33, 0x3c, 0x5d, 0x84, 0xa4, 0xb8, 0x38, 0x33, 0xe3, 0x53, 0xe2, 0x93, 0x2a, 0x4b, 0x52,
	0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xd8, 0x33, 0x5d, 0x9c, 0x40, 0x5c, 0x2b, 0x96,
	0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0x8d, 0x4c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0xc0, 0x25, 0x9f, 0x9c, 0x9f, 0xab, 0x87, 0xc7, 0x76, 0x27, 0x6e, 0xb0, 0x2d, 0x29, 0x01, 0x20,
	0x4b, 0x03, 0x18, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x1d, 0x41, 0xda, 0x7c, 0xc1, 0xda, 0xf4, 0x21, 0xda, 0xf4, 0xd3, 0xf3, 0xf5, 0x61, 0x5e, 0x5c,
	0xc4, 0xc4, 0xe2, 0x18, 0xec, 0xe9, 0xb4, 0x8a, 0x49, 0xda, 0x11, 0xc9, 0xf8, 0x60, 0x88, 0xf1,
	0x9e, 0x29, 0xc5, 0x7a, 0x4e, 0x89, 0xc5, 0xa9, 0xa7, 0x50, 0x64, 0x63, 0x20, 0xb2, 0x31, 0x9e,
	0x29, 0xc5, 0x31, 0x20, 0xd9, 0x47, 0x4c, 0xea, 0x78, 0x64, 0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53,
	0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0x31, 0xc9, 0x21, 0xa9, 0xb4, 0xb2, 0x82, 0x28, 0xb5,
	0xb2, 0xf2, 0x4c, 0x29, 0xb6, 0xb2, 0x02, 0x29, 0x4e, 0x62, 0x03, 0x07, 0x9b, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xda, 0xa0, 0x0a, 0x20, 0x83, 0x01, 0x00, 0x00,
}

func (m *HashID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IDBytes) > 0 {
		i -= len(m.IDBytes)
		copy(dAtA[i:], m.IDBytes)
		i = encodeVarintHashId(dAtA, i, uint64(len(m.IDBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHashId(dAtA []byte, offset int, v uint64) int {
	offset -= sovHashId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HashID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IDBytes)
	if l > 0 {
		n += 1 + l + sovHashId(uint64(l))
	}
	return n
}

func sovHashId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHashId(x uint64) (n int) {
	return sovHashId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HashID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashId
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
			return fmt.Errorf("proto: HashID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashId
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
				return ErrInvalidLengthHashId
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHashId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDBytes = append(m.IDBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.IDBytes == nil {
				m.IDBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHashId
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
func skipHashId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHashId
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
					return 0, ErrIntOverflowHashId
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
					return 0, ErrIntOverflowHashId
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
				return 0, ErrInvalidLengthHashId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHashId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHashId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHashId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHashId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHashId = fmt.Errorf("proto: unexpected end of group")
)
