// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/classifications/internal/genesis/genesis.v1.proto

package genesis

import (
	fmt "fmt"
	mappable "github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	base "github.com/AssetMantle/modules/schema/parameters/base"
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

type Genesis struct {
	Mappables     []*mappable.Mappable `protobuf:"bytes,1,rep,name=mappables,proto3" json:"mappables,omitempty"`
	ParameterList *base.ParameterList  `protobuf:"bytes,2,opt,name=parameterList,proto3" json:"parameterList,omitempty"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6a48d48ffacd76, []int{0}
}
func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return m.Size()
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Genesis)(nil), "classifications.Genesis")
}

func init() {
	proto.RegisterFile("modules/classifications/internal/genesis/genesis.v1.proto", fileDescriptor_5c6a48d48ffacd76)
}

var fileDescriptor_5c6a48d48ffacd76 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcc, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x4f, 0xcd,
	0x4b, 0x2d, 0xce, 0x2c, 0x86, 0xd1, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xfc, 0x68, 0x5a, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x72, 0xfa, 0x20, 0x16, 0x44, 0x99,
	0x94, 0x35, 0x41, 0x1b, 0x72, 0x13, 0x0b, 0x0a, 0x12, 0x93, 0x72, 0x52, 0xe1, 0x0c, 0xb8, 0x1d,
	0x52, 0xba, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0xfa, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xa9, 0x25,
	0xa9, 0x45, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0x08, 0xbe, 0x4f, 0x66, 0x71, 0x09, 0x5c, 0xb9,
	0x52, 0x27, 0x23, 0x17, 0xbb, 0x3b, 0xc4, 0x9d, 0x42, 0xe6, 0x5c, 0x9c, 0x30, 0xf3, 0x8a, 0x25,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0xd0, 0xdc, 0xa0, 0xe7, 0x0b, 0x55, 0x11, 0x84,
	0x50, 0x2b, 0x64, 0xcf, 0xc5, 0x8b, 0x62, 0xbc, 0x04, 0x93, 0x02, 0x23, 0x58, 0x33, 0xc2, 0x11,
	0x7a, 0x01, 0xc8, 0x0a, 0x82, 0x50, 0xd5, 0x5b, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0x7a, 0xc4,
	0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0xe8,
	0x4e, 0x71, 0xe2, 0x83, 0x3a, 0x3c, 0xcc, 0x30, 0x00, 0xe4, 0x97, 0x00, 0xc6, 0x28, 0xf7, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4, 0x12, 0xdf,
	0xc4, 0xbc, 0x12, 0x50, 0x30, 0x41, 0x03, 0x94, 0xd8, 0xa8, 0x5b, 0xc4, 0xc4, 0xec, 0x1c, 0x11,
	0xb1, 0x8a, 0x89, 0xdf, 0x19, 0x55, 0xe1, 0x29, 0x0c, 0x91, 0x47, 0x4c, 0xd2, 0x68, 0x22, 0x31,
	0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0xaf, 0x30, 0xd4, 0x27, 0xb1,
	0x81, 0xc3, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x55, 0xc5, 0xd0, 0x47, 0x02, 0x00,
	0x00,
}

func (m *Genesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Genesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Genesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParameterList != nil {
		{
			size, err := m.ParameterList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesisV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Mappables) > 0 {
		for iNdEx := len(m.Mappables) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Mappables[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesisV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesisV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesisV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Genesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Mappables) > 0 {
		for _, e := range m.Mappables {
			l = e.Size()
			n += 1 + l + sovGenesisV1(uint64(l))
		}
	}
	if m.ParameterList != nil {
		l = m.ParameterList.Size()
		n += 1 + l + sovGenesisV1(uint64(l))
	}
	return n
}

func sovGenesisV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesisV1(x uint64) (n int) {
	return sovGenesisV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Genesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisV1
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
			return fmt.Errorf("proto: Genesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Genesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mappables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisV1
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
				return ErrInvalidLengthGenesisV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mappables = append(m.Mappables, &mappable.Mappable{})
			if err := m.Mappables[len(m.Mappables)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisV1
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
				return ErrInvalidLengthGenesisV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParameterList == nil {
				m.ParameterList = &base.ParameterList{}
			}
			if err := m.ParameterList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisV1
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
func skipGenesisV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesisV1
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
					return 0, ErrIntOverflowGenesisV1
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
					return 0, ErrIntOverflowGenesisV1
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
				return 0, ErrInvalidLengthGenesisV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesisV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesisV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesisV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesisV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesisV1 = fmt.Errorf("proto: unexpected end of group")
)
