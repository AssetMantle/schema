// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/ids/base/classificationID.proto

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

type ClassificationID struct {
	HashID *HashID `protobuf:"bytes,1,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d,omitempty"`
}

func (m *ClassificationID) Reset()         { *m = ClassificationID{} }
func (m *ClassificationID) String() string { return proto.CompactTextString(m) }
func (*ClassificationID) ProtoMessage()    {}
func (*ClassificationID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d456d6663ba5c580, []int{0}
}
func (m *ClassificationID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassificationID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassificationID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassificationID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationID.Merge(m, src)
}
func (m *ClassificationID) XXX_Size() int {
	return m.Size()
}
func (m *ClassificationID) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationID.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClassificationID)(nil), "ids.ClassificationID")
}

func init() { proto.RegisterFile("x/ids/base/classificationID.proto", fileDescriptor_d456d6663ba5c580) }

var fileDescriptor_d456d6663ba5c580 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xac, 0xd0, 0xcf, 0x4c,
	0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c,
	0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xf3, 0x74, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce,
	0x4c, 0x29, 0x96, 0x12, 0x47, 0x52, 0x97, 0x91, 0x58, 0x9c, 0x01, 0x93, 0x95, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x92, 0x3d, 0x97, 0x80, 0x33, 0x9a, 0x69,
	0x42, 0xaa, 0x5c, 0x1c, 0x20, 0x9d, 0xf1, 0x99, 0xf1, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xdc, 0x7a, 0x99, 0x29, 0xc5, 0x7a, 0x1e, 0x60, 0xe3, 0x82, 0xd8, 0x20, 0xc6, 0x5a, 0xb1,
	0x74, 0x2c, 0x90, 0x67, 0x70, 0xea, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x06, 0x2e, 0xf6, 0xe4, 0xfc, 0x5c, 0x90, 0x46, 0x27, 0x51, 0x74, 0x2b, 0x02, 0x40, 0x76, 0x07,
	0x30, 0x46, 0x69, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16,
	0x17, 0xa7, 0x96, 0xf8, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26,
	0xea, 0x23, 0xbc, 0xb2, 0x88, 0x89, 0xd9, 0x33, 0x22, 0x62, 0x15, 0x13, 0xb3, 0x67, 0x4a, 0xf1,
	0x29, 0x30, 0xf9, 0x88, 0x89, 0xdf, 0x33, 0xa5, 0x38, 0xc6, 0x3d, 0xc0, 0xc9, 0x37, 0xb5, 0x24,
	0x31, 0x25, 0xb1, 0x24, 0xf1, 0x15, 0x58, 0x3c, 0x89, 0x0d, 0xec, 0x31, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xa6, 0x71, 0x14, 0x31, 0x01, 0x00, 0x00,
}

func (m *ClassificationID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassificationID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassificationID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintClassificationID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClassificationID(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassificationID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassificationID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovClassificationID(uint64(l))
	}
	return n
}

func sovClassificationID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassificationID(x uint64) (n int) {
	return sovClassificationID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassificationID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassificationID
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
			return fmt.Errorf("proto: ClassificationID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassificationID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassificationID
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
				return ErrInvalidLengthClassificationID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassificationID
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
			skippy, err := skipClassificationID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassificationID
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
func skipClassificationID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassificationID
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
					return 0, ErrIntOverflowClassificationID
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
					return 0, ErrIntOverflowClassificationID
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
				return 0, ErrInvalidLengthClassificationID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassificationID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassificationID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassificationID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassificationID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassificationID = fmt.Errorf("proto: unexpected end of group")
)