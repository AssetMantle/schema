// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/documents/base/document.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/x/ids/base"
	base1 "github.com/AssetMantle/schema/x/qualified/base"
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

type Document struct {
	ClassificationID *base.ClassificationID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	Immutables       *base1.Immutables      `protobuf:"bytes,2,opt,name=immutables,proto3" json:"immutables,omitempty"`
	Mutables         *base1.Mutables        `protobuf:"bytes,3,opt,name=mutables,proto3" json:"mutables,omitempty"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca9fd9cb06ce5a9, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Document.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return m.Size()
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Document)(nil), "documents.Document")
}

func init() { proto.RegisterFile("x/documents/base/document.proto", fileDescriptor_9ca9fd9cb06ce5a9) }

var fileDescriptor_9ca9fd9cb06ce5a9 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xaf, 0xd0, 0x4f, 0xc9,
	0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0x85, 0x73, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xe1, 0xd2, 0x52, 0x8a, 0x15, 0xfa, 0x99, 0x29, 0x50, 0x55,
	0xc9, 0x39, 0x89, 0xc5, 0xc5, 0x99, 0x69, 0x99, 0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x9e, 0x2e,
	0x10, 0xd5, 0x20, 0x25, 0x85, 0xa5, 0x89, 0x39, 0x99, 0x69, 0x99, 0xa9, 0x29, 0x10, 0x85, 0x99,
	0xb9, 0xb9, 0xa5, 0x25, 0x89, 0x49, 0x39, 0xa9, 0xc5, 0x50, 0x25, 0xf2, 0x18, 0x4a, 0xd0, 0x14,
	0x88, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x69, 0x2f, 0x23, 0x17,
	0x87, 0x0b, 0xd4, 0x29, 0x42, 0xce, 0x5c, 0x42, 0xa8, 0x0e, 0x88, 0xcf, 0x8c, 0x4f, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd5, 0xcb, 0x4c, 0x29, 0xd6, 0x73, 0x46, 0x73, 0x5f, 0x90,
	0x00, 0xba, 0x8b, 0x85, 0x4c, 0xb9, 0xb8, 0x10, 0x8e, 0x93, 0x60, 0x82, 0x6a, 0x86, 0xbb, 0x4d,
	0xcf, 0x13, 0x2e, 0x19, 0x84, 0xa4, 0x50, 0x48, 0x9f, 0x8b, 0x03, 0xae, 0x89, 0x19, 0xac, 0x49,
	0x18, 0x49, 0x93, 0x2f, 0x4c, 0x0b, 0x5c, 0x91, 0x15, 0x4b, 0xc7, 0x02, 0x79, 0x06, 0xa7, 0x55,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0xc5, 0x9b, 0x9c, 0x9f, 0xab,
	0x07, 0x0f, 0x66, 0x27, 0x5e, 0x98, 0x37, 0x03, 0x40, 0x1e, 0x0f, 0x60, 0x8c, 0xd2, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x2c, 0x2e, 0x4e, 0x2d, 0xf1, 0x4d,
	0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x47, 0x8f, 0xc1, 0x45,
	0x4c, 0xcc, 0x2e, 0x11, 0x11, 0xab, 0x98, 0x38, 0x61, 0xe6, 0x14, 0x9f, 0x42, 0x62, 0x3f, 0x62,
	0x12, 0x85, 0xb3, 0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f,
	0x21, 0xa9, 0x49, 0x62, 0x03, 0x87, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x60, 0x28, 0xec,
	0x65, 0x1e, 0x02, 0x00, 0x00,
}

func (m *Document) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Document) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Document) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mutables != nil {
		{
			size, err := m.Mutables.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocument(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Immutables != nil {
		{
			size, err := m.Immutables.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocument(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocument(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocument(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocument(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Document) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.Immutables != nil {
		l = m.Immutables.Size()
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.Mutables != nil {
		l = m.Mutables.Size()
		n += 1 + l + sovDocument(uint64(l))
	}
	return n
}

func sovDocument(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocument(x uint64) (n int) {
	return sovDocument(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Document) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocument
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
			return fmt.Errorf("proto: Document: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Document: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Immutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Immutables == nil {
				m.Immutables = &base1.Immutables{}
			}
			if err := m.Immutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mutables == nil {
				m.Mutables = &base1.Mutables{}
			}
			if err := m.Mutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocument(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocument
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
func skipDocument(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
				return 0, ErrInvalidLengthDocument
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocument
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocument
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocument        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocument          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocument = fmt.Errorf("proto: unexpected end of group")
)
