// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/transactions/unwrap/transactionRequest.v1.proto

package unwrap

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

type TransactionRequest struct {
	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID    string `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	OwnableID string `protobuf:"bytes,3,opt,name=ownable_i_d,json=ownableID,proto3" json:"ownable_i_d,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dc7b322ac641be4, []int{0}
}
func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "splits.transactions.unwrap.TransactionRequest")
}

func init() {
	proto.RegisterFile("modules/splits/internal/transactions/unwrap/transactionRequest.v1.proto", fileDescriptor_6dc7b322ac641be4)
}

var fileDescriptor_6dc7b322ac641be4 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x2f, 0xcd, 0x2b, 0x2f, 0x4a, 0x2c, 0x40, 0x16, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0x18, 0xa0, 0x87, 0xac,
	0x4f, 0x0f, 0xa2, 0x4f, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x4c, 0x1f, 0xc4, 0x82, 0xe8,
	0x50, 0xaa, 0xe3, 0x12, 0x0a, 0xc1, 0x30, 0x50, 0x48, 0x88, 0x8b, 0x25, 0xad, 0x28, 0x3f, 0x57,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0xe2, 0x00, 0xd1, 0xf1, 0x99,
	0xf1, 0x29, 0x12, 0x4c, 0x60, 0x71, 0x36, 0x10, 0xdf, 0xd3, 0x45, 0x48, 0x8e, 0x8b, 0x3b, 0xbf,
	0x3c, 0x2f, 0x31, 0x29, 0x27, 0x15, 0x2c, 0xc9, 0x0c, 0x96, 0xe4, 0x84, 0x0a, 0x79, 0xba, 0x08,
	0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0xb0, 0x80, 0x65, 0x20, 0x1c, 0x2b, 0x96,
	0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0x2b, 0x99, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x81, 0x4b, 0x2e, 0x39, 0x3f, 0x57, 0x0f, 0xb7, 0x87, 0x9c, 0x24, 0x31, 0x1d, 0x1e, 0x66, 0x18,
	0x00, 0xf2, 0x55, 0x00, 0x63, 0x94, 0x57, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x49, 0x4e, 0xaa, 0x3e, 0x2c, 0xa4,
	0x49, 0x08, 0xf1, 0x45, 0x4c, 0xcc, 0xc1, 0x21, 0xa1, 0xab, 0x98, 0xa4, 0x82, 0x21, 0x6e, 0x09,
	0x41, 0x76, 0x4b, 0x28, 0x58, 0xc9, 0x29, 0x98, 0x64, 0x0c, 0xb2, 0x64, 0x0c, 0x44, 0xf2, 0x11,
	0x93, 0x1a, 0x6e, 0xc9, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4,
	0x57, 0x4c, 0x32, 0x10, 0x85, 0x56, 0x56, 0xc8, 0x2a, 0xad, 0xac, 0x20, 0x4a, 0x93, 0xd8, 0xc0,
	0x51, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xde, 0x3d, 0x8c, 0x27, 0x2f, 0x02, 0x00, 0x00,
}

func (m *TransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OwnableID) > 0 {
		i -= len(m.OwnableID)
		copy(dAtA[i:], m.OwnableID)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.OwnableID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromID) > 0 {
		i -= len(m.FromID)
		copy(dAtA[i:], m.FromID)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.FromID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionRequestV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionRequestV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.FromID)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.OwnableID)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	return n
}

func sovTransactionRequestV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionRequestV1(x uint64) (n int) {
	return sovTransactionRequestV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionRequestV1
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
			return fmt.Errorf("proto: TransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnableID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionRequestV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionRequestV1
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
func skipTransactionRequestV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionRequestV1
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
					return 0, ErrIntOverflowTransactionRequestV1
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
					return 0, ErrIntOverflowTransactionRequestV1
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
				return 0, ErrInvalidLengthTransactionRequestV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionRequestV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionRequestV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionRequestV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionRequestV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionRequestV1 = fmt.Errorf("proto: unexpected end of group")
)