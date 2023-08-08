// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/split_id.proto

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

type SplitID struct {
	AssetID *AssetID    `protobuf:"bytes,1,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
	OwnerID *IdentityID `protobuf:"bytes,2,opt,name=owner_i_d,json=ownerID,proto3" json:"owner_i_d,omitempty"`
}

func (m *SplitID) Reset()         { *m = SplitID{} }
func (m *SplitID) String() string { return proto.CompactTextString(m) }
func (*SplitID) ProtoMessage()    {}
func (*SplitID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1253af3276df6cff, []int{0}
}
func (m *SplitID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SplitID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SplitID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SplitID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitID.Merge(m, src)
}
func (m *SplitID) XXX_Size() int {
	return m.Size()
}
func (m *SplitID) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitID.DiscardUnknown(m)
}

var xxx_messageInfo_SplitID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SplitID)(nil), "assetmantle.schema.ids.base.SplitID")
}

func init() { proto.RegisterFile("ids/base/split_id.proto", fileDescriptor_1253af3276df6cff) }

var fileDescriptor_1253af3276df6cff = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x18, 0xc7, 0x2f, 0x47, 0x79, 0x4b, 0xef, 0x75, 0x2a, 0x82, 0x72, 0x85, 0x54, 0x44, 0xa8, 0x2e,
	0x09, 0xe8, 0x96, 0xc9, 0x9e, 0x05, 0xc9, 0x50, 0x28, 0x76, 0x93, 0x83, 0x92, 0x36, 0xe1, 0x1a,
	0xe8, 0x35, 0xa5, 0x89, 0x88, 0xab, 0x93, 0xa3, 0x93, 0xb3, 0x38, 0xfa, 0x49, 0xc4, 0xa9, 0xa3,
	0xa3, 0x5c, 0x37, 0x3f, 0x85, 0x24, 0xa9, 0x67, 0x5d, 0x6e, 0x7b, 0xe0, 0xf7, 0xff, 0xff, 0xee,
	0x79, 0x2e, 0xd1, 0x9e, 0xe4, 0x1a, 0x8f, 0x99, 0x16, 0x58, 0x2f, 0x66, 0xd2, 0x8c, 0x24, 0x47,
	0x8b, 0xa5, 0x32, 0xaa, 0xd9, 0x62, 0x5a, 0x0b, 0x93, 0xb3, 0xb9, 0x99, 0x09, 0xa4, 0x27, 0x53,
	0x91, 0x33, 0x24, 0xb9, 0x46, 0x36, 0x1b, 0xef, 0x66, 0x2a, 0x53, 0x2e, 0x87, 0xed, 0xe4, 0x2b,
	0xf1, 0xaf, 0xcb, 0x75, 0x4b, 0x57, 0x1c, 0x97, 0x40, 0x72, 0x31, 0x37, 0xd2, 0xdc, 0x95, 0xec,
	0xf0, 0x09, 0x44, 0xf5, 0xa1, 0xfd, 0x34, 0xed, 0x35, 0xcf, 0xa3, 0xc6, 0xa6, 0x39, 0xe2, 0xfb,
	0xe0, 0x00, 0x1c, 0xff, 0x3f, 0x3d, 0x42, 0x15, 0x7b, 0xa0, 0xae, 0x65, 0xb4, 0x77, 0x55, 0x67,
	0x7e, 0x68, 0x5e, 0x44, 0x0d, 0x75, 0x3b, 0x17, 0x4b, 0x67, 0x08, 0x9d, 0xa1, 0x53, 0x69, 0xa0,
	0x9b, 0x85, 0xac, 0xc4, 0x35, 0x69, 0x8f, 0xd4, 0x1e, 0x9e, 0xdb, 0x41, 0x72, 0x1f, 0xbe, 0x15,
	0x10, 0xac, 0x0a, 0x08, 0x3e, 0x0b, 0x08, 0x1e, 0xd7, 0x30, 0x58, 0xad, 0x61, 0xf0, 0xb1, 0x86,
	0x41, 0xd4, 0x9e, 0xa8, 0xbc, 0xca, 0x9a, 0xec, 0xf8, 0x8b, 0xf8, 0xc0, 0x9e, 0x38, 0x00, 0xd7,
	0x27, 0x99, 0x34, 0xd3, 0x9b, 0x31, 0x9a, 0xa8, 0x1c, 0xbb, 0x9d, 0xfb, 0xae, 0x87, 0x7d, 0x0f,
	0x67, 0x0a, 0xff, 0xfc, 0xa1, 0x97, 0xb0, 0xd6, 0x1d, 0xd2, 0xe4, 0x35, 0x6c, 0x75, 0xb7, 0xfc,
	0x43, 0xef, 0xa7, 0x5c, 0xa3, 0x84, 0x69, 0xf1, 0xfe, 0x87, 0xa6, 0x9e, 0xa6, 0x94, 0xeb, 0xd4,
	0xd2, 0x22, 0xec, 0x54, 0xd0, 0xf4, 0x72, 0x90, 0xf4, 0x85, 0x61, 0x9c, 0x19, 0xf6, 0x15, 0xc2,
	0xad, 0x24, 0x21, 0x3e, 0x4a, 0x08, 0xe5, 0x9a, 0x10, 0x1b, 0x1e, 0xff, 0x73, 0x8f, 0x74, 0xf6,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x44, 0x16, 0x78, 0x27, 0x02, 0x00, 0x00,
}

func (m *SplitID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SplitID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SplitID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OwnerID != nil {
		{
			size, err := m.OwnerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplitId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplitId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSplitId(dAtA []byte, offset int, v uint64) int {
	offset -= sovSplitId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SplitID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovSplitId(uint64(l))
	}
	if m.OwnerID != nil {
		l = m.OwnerID.Size()
		n += 1 + l + sovSplitId(uint64(l))
	}
	return n
}

func sovSplitId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSplitId(x uint64) (n int) {
	return sovSplitId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SplitID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSplitId
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
			return fmt.Errorf("proto: SplitID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SplitID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitId
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
				return ErrInvalidLengthSplitId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplitId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AssetID == nil {
				m.AssetID = &AssetID{}
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitId
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
				return ErrInvalidLengthSplitId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplitId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerID == nil {
				m.OwnerID = &IdentityID{}
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSplitId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSplitId
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
func skipSplitId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSplitId
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
					return 0, ErrIntOverflowSplitId
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
					return 0, ErrIntOverflowSplitId
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
				return 0, ErrInvalidLengthSplitId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSplitId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSplitId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSplitId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSplitId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSplitId = fmt.Errorf("proto: unexpected end of group")
)
