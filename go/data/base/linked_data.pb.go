// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data/base/linked_data.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
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

type LinkedData struct {
	ResourceID      *base.HashID   `protobuf:"bytes,1,opt,name=resource_i_d,json=resourceID,proto3" json:"resource_i_d,omitempty"`
	ExtensionID     *base.StringID `protobuf:"bytes,2,opt,name=extension_i_d,json=extensionID,proto3" json:"extension_i_d,omitempty"`
	ServiceEndpoint string         `protobuf:"bytes,3,opt,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"`
}

func (m *LinkedData) Reset()         { *m = LinkedData{} }
func (m *LinkedData) String() string { return proto.CompactTextString(m) }
func (*LinkedData) ProtoMessage()    {}
func (*LinkedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d125fc542bbd49f1, []int{0}
}
func (m *LinkedData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkedData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedData.Merge(m, src)
}
func (m *LinkedData) XXX_Size() int {
	return m.Size()
}
func (m *LinkedData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedData.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LinkedData)(nil), "assetmantle.schema.data.base.LinkedData")
}

func init() { proto.RegisterFile("data/base/linked_data.proto", fileDescriptor_d125fc542bbd49f1) }

var fileDescriptor_d125fc542bbd49f1 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xe3, 0x40,
	0x1c, 0xc6, 0x33, 0xd9, 0xb2, 0xb0, 0xd3, 0x5d, 0xba, 0x84, 0x65, 0x29, 0xb5, 0xa4, 0x45, 0x11,
	0xaa, 0x87, 0x09, 0xe8, 0x2d, 0xb7, 0x86, 0x14, 0x0d, 0x58, 0x28, 0xed, 0x4d, 0x02, 0x61, 0x9a,
	0x0c, 0xc9, 0x60, 0x93, 0x29, 0x99, 0xa9, 0xf8, 0x08, 0x82, 0x17, 0x1f, 0x41, 0x3c, 0xfa, 0x24,
	0xe2, 0xc5, 0x1e, 0x3d, 0x4a, 0x7a, 0xf3, 0x29, 0x64, 0x26, 0x35, 0x56, 0x90, 0xde, 0x86, 0xef,
	0xf7, 0x7d, 0xdf, 0xff, 0xcf, 0x7f, 0xe0, 0x4e, 0x84, 0x05, 0xb6, 0xa6, 0x98, 0x13, 0x6b, 0x46,
	0xb3, 0x0b, 0x12, 0x05, 0x52, 0x40, 0xf3, 0x9c, 0x09, 0x66, 0xb4, 0x31, 0xe7, 0x44, 0xa4, 0x38,
	0x13, 0x33, 0x82, 0x78, 0x98, 0x90, 0x14, 0x23, 0x85, 0xa5, 0xbf, 0xf5, 0x2f, 0x66, 0x31, 0x53,
	0x46, 0x4b, 0xbe, 0xca, 0x4c, 0xeb, 0x3f, 0x8d, 0x78, 0xd9, 0x97, 0x60, 0x9e, 0x04, 0x34, 0x5a,
	0xeb, 0xcd, 0x4a, 0xe7, 0x22, 0xa7, 0x59, 0x5c, 0x91, 0xdd, 0x67, 0x00, 0xe1, 0x99, 0x9a, 0xed,
	0x62, 0x81, 0x8d, 0x01, 0xfc, 0x9d, 0x13, 0xce, 0x16, 0x79, 0x48, 0x02, 0x1a, 0x44, 0x4d, 0xd0,
	0x05, 0xbd, 0xfa, 0xd1, 0x1e, 0xfa, 0x66, 0x17, 0x1a, 0x71, 0xb5, 0x0a, 0x3a, 0xc5, 0x3c, 0xf1,
	0xdc, 0x31, 0xfc, 0x08, 0x7a, 0xae, 0xe1, 0xc1, 0x3f, 0xe4, 0x4a, 0x90, 0x8c, 0x53, 0x96, 0xa9,
	0x1e, 0x5d, 0xf5, 0xec, 0x6f, 0xed, 0x99, 0xa8, 0xd5, 0x3c, 0x77, 0x5c, 0xaf, 0xb2, 0x9e, 0x6b,
	0x1c, 0xc0, 0xbf, 0x9c, 0xe4, 0x97, 0x34, 0x24, 0x01, 0xc9, 0xa2, 0x39, 0xa3, 0x99, 0x68, 0xfe,
	0xe8, 0x82, 0xde, 0xaf, 0x71, 0x63, 0xad, 0x0f, 0xd6, 0xb2, 0x5d, 0xbb, 0xbe, 0xeb, 0x68, 0xce,
	0x8d, 0xfe, 0x58, 0x98, 0x60, 0x59, 0x98, 0xe0, 0xb5, 0x30, 0xc1, 0xed, 0xca, 0xd4, 0x96, 0x2b,
	0x53, 0x7b, 0x59, 0x99, 0x1a, 0xec, 0x86, 0x2c, 0x45, 0xdb, 0xce, 0xea, 0x34, 0x3e, 0x6f, 0x31,
	0x92, 0xf7, 0x19, 0x81, 0xf3, 0xc3, 0x98, 0x8a, 0x64, 0x31, 0x45, 0x21, 0x4b, 0xad, 0xbe, 0xcc,
	0x0e, 0x55, 0xd6, 0x2a, 0xb3, 0x56, 0xcc, 0xac, 0xea, 0x17, 0xef, 0xf5, 0x5a, 0x7f, 0xe2, 0x3a,
	0x0f, 0x7a, 0xbb, 0xbf, 0x31, 0x64, 0x52, 0x0e, 0x91, 0x9d, 0xc8, 0xc1, 0x9c, 0x3c, 0x7d, 0xc1,
	0x7e, 0x89, 0x7d, 0x89, 0x7d, 0x89, 0x0b, 0xbd, 0xb7, 0x0d, 0xfb, 0x27, 0x23, 0x67, 0x48, 0x04,
	0x96, 0x43, 0xdf, 0xf4, 0xce, 0x86, 0xd5, 0xb6, 0x4b, 0xaf, 0x6d, 0x4b, 0xb3, 0x6d, 0x4b, 0xf7,
	0xf4, 0xa7, 0xfa, 0xe6, 0xe3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x1e, 0x91, 0x7f, 0x6b,
	0x02, 0x00, 0x00,
}

func (m *LinkedData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceEndpoint) > 0 {
		i -= len(m.ServiceEndpoint)
		copy(dAtA[i:], m.ServiceEndpoint)
		i = encodeVarintLinkedData(dAtA, i, uint64(len(m.ServiceEndpoint)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ExtensionID != nil {
		{
			size, err := m.ExtensionID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLinkedData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ResourceID != nil {
		{
			size, err := m.ResourceID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLinkedData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLinkedData(dAtA []byte, offset int, v uint64) int {
	offset -= sovLinkedData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LinkedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResourceID != nil {
		l = m.ResourceID.Size()
		n += 1 + l + sovLinkedData(uint64(l))
	}
	if m.ExtensionID != nil {
		l = m.ExtensionID.Size()
		n += 1 + l + sovLinkedData(uint64(l))
	}
	l = len(m.ServiceEndpoint)
	if l > 0 {
		n += 1 + l + sovLinkedData(uint64(l))
	}
	return n
}

func sovLinkedData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLinkedData(x uint64) (n int) {
	return sovLinkedData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LinkedData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLinkedData
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
			return fmt.Errorf("proto: LinkedData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkedData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLinkedData
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
				return ErrInvalidLengthLinkedData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLinkedData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResourceID == nil {
				m.ResourceID = &base.HashID{}
			}
			if err := m.ResourceID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtensionID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLinkedData
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
				return ErrInvalidLengthLinkedData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLinkedData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExtensionID == nil {
				m.ExtensionID = &base.StringID{}
			}
			if err := m.ExtensionID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLinkedData
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
				return ErrInvalidLengthLinkedData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLinkedData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLinkedData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLinkedData
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
func skipLinkedData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLinkedData
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
					return 0, ErrIntOverflowLinkedData
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
					return 0, ErrIntOverflowLinkedData
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
				return 0, ErrInvalidLengthLinkedData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLinkedData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLinkedData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLinkedData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLinkedData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLinkedData = fmt.Errorf("proto: unexpected end of group")
)