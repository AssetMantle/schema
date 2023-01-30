// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/parameters/base/parameter.v1.proto

package base

import (
	fmt "fmt"
	base1 "github.com/AssetMantle/modules/schema/data/base"
	base "github.com/AssetMantle/modules/schema/ids/base"
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

type Parameter struct {
	ID   *base.StringID `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d,omitempty"`
	Data *base1.AnyData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8190599c9feb459, []int{0}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return m.Size()
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Parameter)(nil), "parameters.Parameter")
}

func init() {
	proto.RegisterFile("schema/parameters/base/parameter.v1.proto", fileDescriptor_a8190599c9feb459)
}

var fileDescriptor_a8190599c9feb459 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x2a, 0xd6, 0x4f, 0x4a,
	0x2c, 0x4e, 0x45, 0xf0, 0xf5, 0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x10,
	0x6a, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xc2, 0xfa, 0x20, 0x16, 0x44, 0x85, 0x94, 0x22,
	0xd4, 0xb0, 0x94, 0xc4, 0x92, 0x44, 0x88, 0x31, 0x89, 0x79, 0x95, 0x2e, 0x89, 0x25, 0x89, 0x70,
	0x43, 0xe0, 0x4a, 0x32, 0x53, 0xa0, 0x16, 0x15, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x7b, 0xba, 0xc0,
	0x95, 0x28, 0x85, 0x70, 0x71, 0x06, 0xc0, 0x6c, 0x12, 0x92, 0xe3, 0x62, 0xce, 0x8c, 0x4f, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd5, 0xcb, 0x4c, 0x29, 0xd6, 0x0b, 0x86, 0xea, 0x08,
	0x62, 0xca, 0x74, 0x11, 0x52, 0xe4, 0x62, 0x01, 0xd9, 0x26, 0xc1, 0x04, 0x55, 0x00, 0xe2, 0xe8,
	0x39, 0x42, 0x6c, 0x0d, 0x02, 0x4b, 0x59, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0xda, 0xc9, 0x78,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0x7c, 0xc9, 0xf9, 0xb9, 0x7a, 0x08,
	0xcf, 0x39, 0x09, 0xc0, 0xad, 0x0f, 0x33, 0x0c, 0x00, 0x39, 0x29, 0x80, 0x31, 0xca, 0x34, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8, 0x38, 0xb5, 0xc4, 0x37,
	0x31, 0xaf, 0x24, 0x27, 0x55, 0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x58, 0x1f, 0x7b, 0x30,
	0x2e, 0x62, 0x62, 0x0e, 0x88, 0x88, 0x58, 0xc5, 0xc4, 0x05, 0x37, 0xb1, 0xf8, 0x14, 0x32, 0xe7,
	0x11, 0x93, 0x18, 0x82, 0x13, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x08, 0x72, 0xf4, 0x2b,
	0x64, 0x55, 0x49, 0x6c, 0xe0, 0x80, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x42, 0x56, 0x9c,
	0xca, 0xad, 0x01, 0x00, 0x00,
}

func (m *Parameter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Parameter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Parameter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParameterV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ID != nil {
		{
			size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParameterV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParameterV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovParameterV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Parameter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != nil {
		l = m.ID.Size()
		n += 1 + l + sovParameterV1(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovParameterV1(uint64(l))
	}
	return n
}

func sovParameterV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParameterV1(x uint64) (n int) {
	return sovParameterV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Parameter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParameterV1
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
			return fmt.Errorf("proto: Parameter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Parameter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameterV1
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
				return ErrInvalidLengthParameterV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameterV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ID == nil {
				m.ID = &base.StringID{}
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameterV1
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
				return ErrInvalidLengthParameterV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameterV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &base1.AnyData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParameterV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParameterV1
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
func skipParameterV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParameterV1
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
					return 0, ErrIntOverflowParameterV1
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
					return 0, ErrIntOverflowParameterV1
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
				return 0, ErrInvalidLengthParameterV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParameterV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParameterV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParameterV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParameterV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParameterV1 = fmt.Errorf("proto: unexpected end of group")
)