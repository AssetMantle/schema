// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ids/base/any_id.proto

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

type AnyID struct {
	// Types that are valid to be assigned to Impl:
	//	*AnyID_AssetID
	//	*AnyID_ClassificationID
	//	*AnyID_CoinID
	//	*AnyID_DataID
	//	*AnyID_HashID
	//	*AnyID_IdentityID
	//	*AnyID_MaintainerID
	//	*AnyID_OrderID
	//	*AnyID_OwnableID
	//	*AnyID_PropertyID
	//	*AnyID_SplitID
	//	*AnyID_StringID
	Impl isAnyID_Impl `protobuf_oneof:"impl"`
}

func (m *AnyID) Reset()         { *m = AnyID{} }
func (m *AnyID) String() string { return proto.CompactTextString(m) }
func (*AnyID) ProtoMessage()    {}
func (*AnyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8be662da6587db8, []int{0}
}
func (m *AnyID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyID.Merge(m, src)
}
func (m *AnyID) XXX_Size() int {
	return m.Size()
}
func (m *AnyID) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyID.DiscardUnknown(m)
}

var xxx_messageInfo_AnyID proto.InternalMessageInfo

type isAnyID_Impl interface {
	isAnyID_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AnyID_AssetID struct {
	AssetID *AssetID `protobuf:"bytes,1,opt,name=asset_i_d,json=assetID,proto3,oneof" json:"asset_i_d,omitempty"`
}
type AnyID_ClassificationID struct {
	ClassificationID *ClassificationID `protobuf:"bytes,2,opt,name=classification_i_d,json=classificationID,proto3,oneof" json:"classification_i_d,omitempty"`
}
type AnyID_CoinID struct {
	CoinID *CoinID `protobuf:"bytes,3,opt,name=coin_i_d,json=coinID,proto3,oneof" json:"coin_i_d,omitempty"`
}
type AnyID_DataID struct {
	DataID *DataID `protobuf:"bytes,4,opt,name=data_i_d,json=dataID,proto3,oneof" json:"data_i_d,omitempty"`
}
type AnyID_HashID struct {
	HashID *HashID `protobuf:"bytes,5,opt,name=hash_i_d,json=hashID,proto3,oneof" json:"hash_i_d,omitempty"`
}
type AnyID_IdentityID struct {
	IdentityID *IdentityID `protobuf:"bytes,6,opt,name=identity_i_d,json=identityID,proto3,oneof" json:"identity_i_d,omitempty"`
}
type AnyID_MaintainerID struct {
	MaintainerID *MaintainerID `protobuf:"bytes,7,opt,name=maintainer_i_d,json=maintainerID,proto3,oneof" json:"maintainer_i_d,omitempty"`
}
type AnyID_OrderID struct {
	OrderID *OrderID `protobuf:"bytes,8,opt,name=order_i_d,json=orderID,proto3,oneof" json:"order_i_d,omitempty"`
}
type AnyID_OwnableID struct {
	OwnableID *AnyOwnableID `protobuf:"bytes,9,opt,name=ownable_i_d,json=ownableID,proto3,oneof" json:"ownable_i_d,omitempty"`
}
type AnyID_PropertyID struct {
	PropertyID *PropertyID `protobuf:"bytes,10,opt,name=property_i_d,json=propertyID,proto3,oneof" json:"property_i_d,omitempty"`
}
type AnyID_SplitID struct {
	SplitID *SplitID `protobuf:"bytes,11,opt,name=split_i_d,json=splitID,proto3,oneof" json:"split_i_d,omitempty"`
}
type AnyID_StringID struct {
	StringID *StringID `protobuf:"bytes,12,opt,name=string_i_d,json=stringID,proto3,oneof" json:"string_i_d,omitempty"`
}

func (*AnyID_AssetID) isAnyID_Impl()          {}
func (*AnyID_ClassificationID) isAnyID_Impl() {}
func (*AnyID_CoinID) isAnyID_Impl()           {}
func (*AnyID_DataID) isAnyID_Impl()           {}
func (*AnyID_HashID) isAnyID_Impl()           {}
func (*AnyID_IdentityID) isAnyID_Impl()       {}
func (*AnyID_MaintainerID) isAnyID_Impl()     {}
func (*AnyID_OrderID) isAnyID_Impl()          {}
func (*AnyID_OwnableID) isAnyID_Impl()        {}
func (*AnyID_PropertyID) isAnyID_Impl()       {}
func (*AnyID_SplitID) isAnyID_Impl()          {}
func (*AnyID_StringID) isAnyID_Impl()         {}

func (m *AnyID) GetImpl() isAnyID_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *AnyID) GetAssetID() *AssetID {
	if x, ok := m.GetImpl().(*AnyID_AssetID); ok {
		return x.AssetID
	}
	return nil
}

func (m *AnyID) GetClassificationID() *ClassificationID {
	if x, ok := m.GetImpl().(*AnyID_ClassificationID); ok {
		return x.ClassificationID
	}
	return nil
}

func (m *AnyID) GetCoinID() *CoinID {
	if x, ok := m.GetImpl().(*AnyID_CoinID); ok {
		return x.CoinID
	}
	return nil
}

func (m *AnyID) GetDataID() *DataID {
	if x, ok := m.GetImpl().(*AnyID_DataID); ok {
		return x.DataID
	}
	return nil
}

func (m *AnyID) GetHashID() *HashID {
	if x, ok := m.GetImpl().(*AnyID_HashID); ok {
		return x.HashID
	}
	return nil
}

func (m *AnyID) GetIdentityID() *IdentityID {
	if x, ok := m.GetImpl().(*AnyID_IdentityID); ok {
		return x.IdentityID
	}
	return nil
}

func (m *AnyID) GetMaintainerID() *MaintainerID {
	if x, ok := m.GetImpl().(*AnyID_MaintainerID); ok {
		return x.MaintainerID
	}
	return nil
}

func (m *AnyID) GetOrderID() *OrderID {
	if x, ok := m.GetImpl().(*AnyID_OrderID); ok {
		return x.OrderID
	}
	return nil
}

func (m *AnyID) GetOwnableID() *AnyOwnableID {
	if x, ok := m.GetImpl().(*AnyID_OwnableID); ok {
		return x.OwnableID
	}
	return nil
}

func (m *AnyID) GetPropertyID() *PropertyID {
	if x, ok := m.GetImpl().(*AnyID_PropertyID); ok {
		return x.PropertyID
	}
	return nil
}

func (m *AnyID) GetSplitID() *SplitID {
	if x, ok := m.GetImpl().(*AnyID_SplitID); ok {
		return x.SplitID
	}
	return nil
}

func (m *AnyID) GetStringID() *StringID {
	if x, ok := m.GetImpl().(*AnyID_StringID); ok {
		return x.StringID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnyID) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnyID_AssetID)(nil),
		(*AnyID_ClassificationID)(nil),
		(*AnyID_CoinID)(nil),
		(*AnyID_DataID)(nil),
		(*AnyID_HashID)(nil),
		(*AnyID_IdentityID)(nil),
		(*AnyID_MaintainerID)(nil),
		(*AnyID_OrderID)(nil),
		(*AnyID_OwnableID)(nil),
		(*AnyID_PropertyID)(nil),
		(*AnyID_SplitID)(nil),
		(*AnyID_StringID)(nil),
	}
}

func init() {
	proto.RegisterType((*AnyID)(nil), "ids.AnyID")
}

func init() { proto.RegisterFile("ids/base/any_id.proto", fileDescriptor_d8be662da6587db8) }

var fileDescriptor_d8be662da6587db8 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x6f, 0xda, 0x6e, 0x77, 0x77, 0x5a, 0x5d, 0x0d, 0xae, 0x0e, 0x45, 0xe3, 0xe2, 0xc5, 0x3f,
	0x60, 0x03, 0xf6, 0xa4, 0xb7, 0xd6, 0x8a, 0x9b, 0x43, 0xd9, 0xa2, 0x97, 0x45, 0x84, 0x61, 0x9a,
	0x17, 0x9b, 0x81, 0x26, 0x13, 0x32, 0x23, 0xd2, 0xbb, 0x07, 0x8f, 0x7e, 0x04, 0xf1, 0xe8, 0x27,
	0x11, 0x4f, 0x7b, 0xf4, 0x28, 0xad, 0x27, 0x3f, 0x85, 0xcc, 0x9b, 0x34, 0xc9, 0x2e, 0xbd, 0x84,
	0xf7, 0x7e, 0x7f, 0x26, 0xf3, 0xde, 0xfc, 0xc8, 0xb1, 0x00, 0xe5, 0xcf, 0xb9, 0x8a, 0x7c, 0x9e,
	0xae, 0x98, 0x80, 0x41, 0x96, 0x4b, 0x2d, 0xdd, 0x96, 0x00, 0xd5, 0xbf, 0xb5, 0x90, 0x0b, 0x89,
	0xbd, 0x6f, 0x2a, 0x4b, 0xf5, 0xef, 0x5d, 0x72, 0xc8, 0x4f, 0x29, 0x9f, 0x2f, 0xa3, 0xd2, 0xd9,
	0xbf, 0x53, 0xd1, 0x4a, 0x45, 0xba, 0x22, 0x4e, 0x4a, 0x22, 0x5c, 0x72, 0xa5, 0xc4, 0x07, 0x11,
	0x72, 0x2d, 0x64, 0x5a, 0x29, 0x6e, 0x57, 0x0a, 0x29, 0x76, 0xe2, 0xc0, 0x35, 0xdf, 0x85, 0xc7,
	0x5c, 0xc5, 0x15, 0xde, 0x2f, 0x71, 0x01, 0x51, 0xaa, 0x85, 0xae, 0x06, 0xeb, 0xdf, 0x2d, 0xb9,
	0x84, 0x8b, 0x54, 0x73, 0x91, 0x46, 0xf9, 0xae, 0xcb, 0xcb, 0x1c, 0xea, 0x44, 0x75, 0x64, 0x96,
	0xcb, 0x2c, 0xca, 0xeb, 0x47, 0x56, 0x26, 0x95, 0x2d, 0x45, 0x6d, 0x62, 0x5a, 0x11, 0x3a, 0x17,
	0xe9, 0xa2, 0x64, 0x1e, 0xfc, 0x6d, 0x93, 0xbd, 0x51, 0xba, 0x0a, 0x26, 0xee, 0x13, 0x72, 0x58,
	0xec, 0x89, 0x01, 0x75, 0x4e, 0x9c, 0x47, 0xdd, 0x67, 0xbd, 0x81, 0x00, 0x35, 0x18, 0x19, 0x34,
	0x98, 0x9c, 0x36, 0xde, 0xec, 0x73, 0x5b, 0xba, 0xaf, 0x88, 0x7b, 0x75, 0x75, 0x0c, 0x68, 0x13,
	0x4d, 0xc7, 0x68, 0x7a, 0x79, 0x89, 0x46, 0xf7, 0x8d, 0xf0, 0x0a, 0xe6, 0x3e, 0x24, 0x07, 0x76,
	0xbf, 0x0c, 0x68, 0x0b, 0xcd, 0x5d, 0x6b, 0x96, 0xc2, 0x5a, 0x3a, 0x21, 0x56, 0x46, 0x68, 0x17,
	0xce, 0x80, 0xb6, 0x6b, 0xc2, 0x09, 0xd7, 0xdc, 0x0a, 0x01, 0x2b, 0x23, 0xb4, 0x2f, 0xc0, 0x80,
	0xee, 0xd5, 0x84, 0xa7, 0x5c, 0xc5, 0x56, 0x18, 0x63, 0xe5, 0x0e, 0x49, 0xaf, 0x7a, 0x12, 0x06,
	0xb4, 0x83, 0xe2, 0x23, 0x14, 0x07, 0x05, 0x81, 0x06, 0x22, 0xca, 0xce, 0x7d, 0x4e, 0xae, 0xd7,
	0xdf, 0x8a, 0x01, 0xdd, 0x47, 0xdb, 0x4d, 0xb4, 0x4d, 0x4b, 0x0a, 0x8d, 0xbd, 0xa4, 0xd6, 0x9b,
	0xed, 0x16, 0x0f, 0xc9, 0x80, 0x1e, 0xd4, 0xb6, 0x7b, 0x66, 0x50, 0xbb, 0x5d, 0x69, 0x4b, 0x77,
	0x48, 0xba, 0x65, 0x98, 0x19, 0xd0, 0xc3, 0xda, 0x3f, 0x46, 0xe9, 0xea, 0xcc, 0x52, 0x68, 0x39,
	0x94, 0xdb, 0xc6, 0x0c, 0x54, 0x05, 0x82, 0x01, 0x25, 0xb5, 0x81, 0x66, 0x05, 0x61, 0x07, 0xca,
	0xca, 0xce, 0xdc, 0xaa, 0x48, 0x0a, 0x03, 0xda, 0xad, 0xdd, 0xea, 0xad, 0x41, 0xed, 0xad, 0x94,
	0x2d, 0xdd, 0xa7, 0x84, 0x6c, 0xc3, 0xc3, 0x80, 0xf6, 0x50, 0x7c, 0xcd, 0x8a, 0x11, 0x46, 0xf5,
	0x81, 0x2a, 0xea, 0x17, 0xed, 0x2f, 0xdf, 0xee, 0x37, 0xc6, 0x1d, 0xd2, 0x16, 0x49, 0xb6, 0x1c,
	0x7f, 0x76, 0x7e, 0xae, 0x3d, 0xe7, 0x62, 0xed, 0x39, 0x7f, 0xd6, 0x9e, 0xf3, 0x75, 0xe3, 0x35,
	0x2e, 0x36, 0x5e, 0xe3, 0xf7, 0xc6, 0x6b, 0x90, 0xfd, 0x50, 0x26, 0xe6, 0x98, 0x31, 0x31, 0x39,
	0x84, 0x99, 0x89, 0xe5, 0xcc, 0x79, 0xf7, 0x78, 0x21, 0x74, 0xfc, 0x71, 0x3e, 0x08, 0x65, 0xe2,
	0x63, 0x02, 0xa7, 0x3c, 0xd5, 0xcb, 0xc8, 0x57, 0x61, 0x1c, 0x25, 0xdc, 0x5f, 0x48, 0x7f, 0x9b,
	0xe9, 0xef, 0xcd, 0x56, 0x70, 0x7e, 0xfe, 0xa3, 0xd9, 0x0a, 0x40, 0xfd, 0xc2, 0xef, 0xba, 0x79,
	0x14, 0x80, 0x7a, 0xff, 0x7a, 0x36, 0x9e, 0x46, 0x9a, 0x9b, 0x68, 0xfc, 0x43, 0x7c, 0xde, 0xc1,
	0xd0, 0x0f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x1f, 0x2e, 0x02, 0x6c, 0x04, 0x00, 0x00,
}

func (m *AnyID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AnyID_AssetID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_AssetID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_ClassificationID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_ClassificationID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_CoinID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_CoinID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CoinID != nil {
		{
			size, err := m.CoinID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_DataID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_DataID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DataID != nil {
		{
			size, err := m.DataID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_HashID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_HashID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HashID != nil {
		{
			size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_IdentityID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_IdentityID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IdentityID != nil {
		{
			size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_MaintainerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_MaintainerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MaintainerID != nil {
		{
			size, err := m.MaintainerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderID != nil {
		{
			size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_OwnableID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_OwnableID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OwnableID != nil {
		{
			size, err := m.OwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_PropertyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_PropertyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PropertyID != nil {
		{
			size, err := m.PropertyID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_SplitID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_SplitID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SplitID != nil {
		{
			size, err := m.SplitID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *AnyID_StringID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyID_StringID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringID != nil {
		{
			size, err := m.StringID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyId(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func encodeVarintAnyId(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnyId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnyID) Size() (n int) {
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

func (m *AnyID_AssetID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_ClassificationID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_CoinID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoinID != nil {
		l = m.CoinID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_DataID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DataID != nil {
		l = m.DataID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_HashID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_IdentityID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_MaintainerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaintainerID != nil {
		l = m.MaintainerID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderID != nil {
		l = m.OrderID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_OwnableID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnableID != nil {
		l = m.OwnableID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_PropertyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropertyID != nil {
		l = m.PropertyID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_SplitID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SplitID != nil {
		l = m.SplitID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}
func (m *AnyID_StringID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringID != nil {
		l = m.StringID.Size()
		n += 1 + l + sovAnyId(uint64(l))
	}
	return n
}

func sovAnyId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnyId(x uint64) (n int) {
	return sovAnyId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnyID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnyId
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
			return fmt.Errorf("proto: AnyID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AssetID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_AssetID{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ClassificationID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_ClassificationID{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CoinID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_CoinID{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DataID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_DataID{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HashID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_HashID{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IdentityID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_IdentityID{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainerID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MaintainerID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_MaintainerID{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OrderID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_OrderID{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AnyOwnableID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_OwnableID{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PropertyID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_PropertyID{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SplitID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_SplitID{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyId
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
				return ErrInvalidLengthAnyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyID_StringID{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnyId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnyId
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
func skipAnyId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnyId
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
					return 0, ErrIntOverflowAnyId
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
					return 0, ErrIntOverflowAnyId
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
				return 0, ErrInvalidLengthAnyId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnyId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnyId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnyId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnyId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnyId = fmt.Errorf("proto: unexpected end of group")
)
