// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/nub/message.proto

package nub

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
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

type Message struct {
	From  string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	NubID *base.StringID `protobuf:"bytes,2,opt,name=nub_i_d,json=nubID,proto3" json:"nub_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fac5b654e790f95, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetNubID() *base.StringID {
	if m != nil {
		return m.NubID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "identities.transactions.nub.Message")
}

func init() {
	proto.RegisterFile("modules/identities/internal/transactions/nub/message.proto", fileDescriptor_3fac5b654e790f95)
}

var fileDescriptor_3fac5b654e790f95 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0x33, 0xfb, 0x7d, 0x1a, 0x5c, 0xb5, 0xd9, 0x2a, 0x24, 0x38, 0x06, 0x41, 0x4c, 0x35,
	0x03, 0xda, 0x6d, 0x67, 0x08, 0xc8, 0x22, 0x59, 0x42, 0x4c, 0x25, 0x81, 0x30, 0xb3, 0x3b, 0x26,
	0x03, 0xd9, 0x19, 0xd9, 0x7b, 0xa7, 0xf0, 0x2d, 0x7c, 0x86, 0x94, 0x3e, 0x89, 0x58, 0xa5, 0xb4,
	0x94, 0x4d, 0xe7, 0x53, 0x48, 0xfe, 0x91, 0xc5, 0x22, 0x60, 0x77, 0xe1, 0xfc, 0xce, 0x9c, 0x33,
	0xc7, 0x0f, 0x33, 0x9b, 0xba, 0xa9, 0x02, 0xae, 0x53, 0x65, 0x50, 0xa3, 0x5e, 0x9e, 0x06, 0x55,
	0x6e, 0xc4, 0x94, 0x63, 0x2e, 0x0c, 0x88, 0x04, 0xb5, 0x35, 0xc0, 0x8d, 0x93, 0x3c, 0x53, 0x00,
	0x62, 0xac, 0xd8, 0x73, 0x6e, 0xd1, 0x06, 0x8d, 0x9d, 0x87, 0x95, 0x51, 0x66, 0x9c, 0xac, 0x37,
	0x21, 0x99, 0xa8, 0x4c, 0x70, 0x9d, 0x02, 0x97, 0x02, 0xd4, 0x36, 0xe0, 0x25, 0xea, 0xac, 0xed,
	0x75, 0xfa, 0x9b, 0x00, 0xcc, 0xb5, 0x19, 0x6f, 0xf5, 0x8b, 0x8e, 0x5f, 0xed, 0xae, 0xf3, 0x82,
	0xc0, 0xff, 0xff, 0x94, 0xdb, 0xac, 0x46, 0x9a, 0xa4, 0x75, 0xd4, 0x5f, 0xdd, 0xc1, 0xa5, 0x5f,
	0x35, 0x4e, 0x8e, 0xf4, 0x28, 0xad, 0x79, 0x4d, 0xd2, 0x3a, 0xbe, 0x3e, 0x65, 0x3a, 0x05, 0xf6,
	0xb0, 0x79, 0xa4, 0x7f, 0x60, 0x9c, 0x8c, 0x3a, 0xed, 0x99, 0xf7, 0x5e, 0x50, 0x32, 0x2f, 0x28,
	0xf9, 0x2a, 0x28, 0x79, 0x5d, 0xd0, 0xca, 0x7c, 0x41, 0x2b, 0x9f, 0x0b, 0x5a, 0xf1, 0xcf, 0x13,
	0x9b, 0xb1, 0x3d, 0x7f, 0x68, 0x9f, 0x6c, 0xf2, 0x7b, 0xcb, 0x3e, 0x3d, 0xf2, 0x78, 0x3f, 0xd6,
	0x38, 0x71, 0x92, 0x25, 0x36, 0xe3, 0xb7, 0x00, 0x0a, 0xbb, 0xc2, 0xe0, 0x54, 0xf1, 0xed, 0x86,
	0x7f, 0xd9, 0x72, 0xe6, 0xfd, 0x8b, 0x06, 0xf1, 0x9b, 0xd7, 0x88, 0x76, 0x05, 0x06, 0xe5, 0x02,
	0xb1, 0x93, 0x1f, 0x65, 0x75, 0x58, 0x56, 0x87, 0xb1, 0x93, 0x85, 0x77, 0xb5, 0x47, 0x1d, 0xde,
	0xf5, 0xda, 0x5d, 0x85, 0x22, 0x15, 0x28, 0xbe, 0xbd, 0xb3, 0x1d, 0x19, 0x86, 0x65, 0x34, 0x0c,
	0x63, 0x27, 0xe5, 0xe1, 0x6a, 0xf1, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x3a, 0x85,
	0xa0, 0x0e, 0x02, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NubID != nil {
		{
			size, err := m.NubID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.NubID != nil {
		l = m.NubID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NubID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NubID == nil {
				m.NubID = &base.StringID{}
			}
			if err := m.NubID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)