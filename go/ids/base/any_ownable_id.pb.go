// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/any_ownable_id.proto

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

type AnyOwnableID struct {
	// Types that are valid to be assigned to Impl:
	//	*AnyOwnableID_AssetID
	//	*AnyOwnableID_CoinID
	Impl isAnyOwnableID_Impl `protobuf_oneof:"impl"`
}

func (m *AnyOwnableID) Reset()         { *m = AnyOwnableID{} }
func (m *AnyOwnableID) String() string { return proto.CompactTextString(m) }
func (*AnyOwnableID) ProtoMessage()    {}
func (*AnyOwnableID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6479d05b42e9587d, []int{0}
}
func (m *AnyOwnableID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyOwnableID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyOwnableID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyOwnableID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyOwnableID.Merge(m, src)
}
func (m *AnyOwnableID) XXX_Size() int {
	return m.Size()
}
func (m *AnyOwnableID) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyOwnableID.DiscardUnknown(m)
}

var xxx_messageInfo_AnyOwnableID proto.InternalMessageInfo

type isAnyOwnableID_Impl interface {
	isAnyOwnableID_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AnyOwnableID_AssetID struct {
	AssetID *AssetID `protobuf:"bytes,1,opt,name=asset_i_d,json=assetID,proto3,oneof" json:"asset_i_d,omitempty"`
}
type AnyOwnableID_CoinID struct {
	CoinID *CoinID `protobuf:"bytes,2,opt,name=coin_i_d,json=coinID,proto3,oneof" json:"coin_i_d,omitempty"`
}

func (*AnyOwnableID_AssetID) isAnyOwnableID_Impl() {}
func (*AnyOwnableID_CoinID) isAnyOwnableID_Impl()  {}

func (m *AnyOwnableID) GetImpl() isAnyOwnableID_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *AnyOwnableID) GetAssetID() *AssetID {
	if x, ok := m.GetImpl().(*AnyOwnableID_AssetID); ok {
		return x.AssetID
	}
	return nil
}

func (m *AnyOwnableID) GetCoinID() *CoinID {
	if x, ok := m.GetImpl().(*AnyOwnableID_CoinID); ok {
		return x.CoinID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnyOwnableID) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnyOwnableID_AssetID)(nil),
		(*AnyOwnableID_CoinID)(nil),
	}
}

func init() {
	proto.RegisterType((*AnyOwnableID)(nil), "ids.base.AnyOwnableID")
}

func init() { proto.RegisterFile("ids/base/any_ownable_id.proto", fileDescriptor_6479d05b42e9587d) }

var fileDescriptor_6479d05b42e9587d = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x4c, 0x29, 0xd6,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0xcc, 0xab, 0x8c, 0xcf, 0x2f, 0xcf, 0x4b, 0x4c, 0xca, 0x49,
	0x8d, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4c, 0x29, 0xd6, 0x03,
	0x49, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x05, 0xf5, 0x41, 0x2c, 0x88, 0xbc, 0x94, 0x38,
	0x42, 0x7b, 0x71, 0x71, 0x6a, 0x09, 0x5c, 0xa3, 0x94, 0x18, 0x5c, 0x22, 0x39, 0x3f, 0x33, 0x0f,
	0x2e, 0xae, 0x54, 0xcd, 0xc5, 0xe3, 0x98, 0x57, 0xe9, 0x0f, 0xb1, 0xc7, 0xd3, 0x45, 0x48, 0x9f,
	0x8b, 0x13, 0xaa, 0x33, 0x3e, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x50, 0x0f, 0x66,
	0xa9, 0x9e, 0x23, 0x48, 0xca, 0xd3, 0xc5, 0x83, 0x21, 0x88, 0x3d, 0x11, 0xc2, 0x14, 0xd2, 0xe1,
	0xe2, 0x80, 0x98, 0x18, 0x9f, 0x22, 0xc1, 0x04, 0x56, 0x2f, 0x80, 0x50, 0xef, 0x9c, 0x9f, 0x99,
	0x07, 0x56, 0xce, 0x96, 0x0c, 0x66, 0x59, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x70, 0x62, 0xe3, 0x62,
	0xc9, 0xcc, 0x2d, 0xc8, 0x71, 0x5a, 0xca, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x5c, 0x3c, 0xc9, 0xf9, 0xb9, 0x70, 0x73, 0x9c, 0x04, 0x91, 0xdc, 0x98, 0x12, 0x00, 0x72,
	0x78, 0x00, 0x63, 0x94, 0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e,
	0xd8, 0x61, 0xbe, 0x89, 0x79, 0x25, 0x39, 0xa9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0xfa,
	0xe9, 0xf9, 0xfa, 0x30, 0x3f, 0x2f, 0x62, 0x62, 0xf6, 0x74, 0x8a, 0x58, 0xc5, 0xc4, 0xe1, 0x99,
	0x52, 0xac, 0xe7, 0x94, 0x58, 0x9c, 0x7a, 0x0a, 0xcc, 0x8c, 0x01, 0x31, 0x1f, 0x31, 0x89, 0xc0,
	0x98, 0x31, 0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0xaf, 0x98, 0x38,
	0x3d, 0x53, 0x8a, 0xad, 0xac, 0x40, 0xe2, 0x49, 0x6c, 0xe0, 0xb0, 0x32, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0x69, 0x02, 0x8f, 0x9d, 0x01, 0x00, 0x00,
}

func (m *AnyOwnableID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyOwnableID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyOwnableID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Impl != nil {
		{
			size := m.Impl.Size()
			i -= size
			if _, err := m.Impl.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *AnyOwnableID_AssetID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyOwnableID_AssetID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyOwnableId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *AnyOwnableID_CoinID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyOwnableID_CoinID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CoinID != nil {
		{
			size, err := m.CoinID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyOwnableId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintAnyOwnableId(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnyOwnableId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnyOwnableID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Impl != nil {
		n += m.Impl.Size()
	}
	return n
}

func (m *AnyOwnableID_AssetID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovAnyOwnableId(uint64(l))
	}
	return n
}
func (m *AnyOwnableID_CoinID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoinID != nil {
		l = m.CoinID.Size()
		n += 1 + l + sovAnyOwnableId(uint64(l))
	}
	return n
}

func sovAnyOwnableId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnyOwnableId(x uint64) (n int) {
	return sovAnyOwnableId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnyOwnableID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnyOwnableId
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
			return fmt.Errorf("proto: AnyOwnableID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyOwnableID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyOwnableId
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
				return ErrInvalidLengthAnyOwnableId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyOwnableId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AssetID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyOwnableID_AssetID{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyOwnableId
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
				return ErrInvalidLengthAnyOwnableId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyOwnableId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CoinID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyOwnableID_CoinID{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnyOwnableId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnyOwnableId
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
func skipAnyOwnableId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnyOwnableId
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
					return 0, ErrIntOverflowAnyOwnableId
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
					return 0, ErrIntOverflowAnyOwnableId
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
				return 0, ErrInvalidLengthAnyOwnableId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnyOwnableId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnyOwnableId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnyOwnableId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnyOwnableId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnyOwnableId = fmt.Errorf("proto: unexpected end of group")
)
