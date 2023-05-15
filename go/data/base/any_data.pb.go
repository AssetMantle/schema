// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data/base/any_data.proto

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

type AnyData struct {
	// Types that are valid to be assigned to Impl:
	//	*AnyData_AccAddressData
	//	*AnyData_BooleanData
	//	*AnyData_DecData
	//	*AnyData_HeightData
	//	*AnyData_IDData
	//	*AnyData_ListData
	//	*AnyData_NumberData
	//	*AnyData_StringData
	Impl isAnyData_Impl `protobuf_oneof:"impl"`
}

func (m *AnyData) Reset()         { *m = AnyData{} }
func (m *AnyData) String() string { return proto.CompactTextString(m) }
func (*AnyData) ProtoMessage()    {}
func (*AnyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b19532f46054c21, []int{0}
}
func (m *AnyData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyData.Merge(m, src)
}
func (m *AnyData) XXX_Size() int {
	return m.Size()
}
func (m *AnyData) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyData.DiscardUnknown(m)
}

var xxx_messageInfo_AnyData proto.InternalMessageInfo

type isAnyData_Impl interface {
	isAnyData_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AnyData_AccAddressData struct {
	AccAddressData *AccAddressData `protobuf:"bytes,1,opt,name=acc_address_data,json=accAddressData,proto3,oneof" json:"acc_address_data,omitempty"`
}
type AnyData_BooleanData struct {
	BooleanData *BooleanData `protobuf:"bytes,2,opt,name=boolean_data,json=booleanData,proto3,oneof" json:"boolean_data,omitempty"`
}
type AnyData_DecData struct {
	DecData *DecData `protobuf:"bytes,3,opt,name=dec_data,json=decData,proto3,oneof" json:"dec_data,omitempty"`
}
type AnyData_HeightData struct {
	HeightData *HeightData `protobuf:"bytes,4,opt,name=height_data,json=heightData,proto3,oneof" json:"height_data,omitempty"`
}
type AnyData_IDData struct {
	IDData *IDData `protobuf:"bytes,5,opt,name=i_d_data,json=iDData,proto3,oneof" json:"i_d_data,omitempty"`
}
type AnyData_ListData struct {
	ListData *ListData `protobuf:"bytes,6,opt,name=list_data,json=listData,proto3,oneof" json:"list_data,omitempty"`
}
type AnyData_NumberData struct {
	NumberData *NumberData `protobuf:"bytes,7,opt,name=number_data,json=numberData,proto3,oneof" json:"number_data,omitempty"`
}
type AnyData_StringData struct {
	StringData *StringData `protobuf:"bytes,8,opt,name=string_data,json=stringData,proto3,oneof" json:"string_data,omitempty"`
}

func (*AnyData_AccAddressData) isAnyData_Impl() {}
func (*AnyData_BooleanData) isAnyData_Impl()    {}
func (*AnyData_DecData) isAnyData_Impl()        {}
func (*AnyData_HeightData) isAnyData_Impl()     {}
func (*AnyData_IDData) isAnyData_Impl()         {}
func (*AnyData_ListData) isAnyData_Impl()       {}
func (*AnyData_NumberData) isAnyData_Impl()     {}
func (*AnyData_StringData) isAnyData_Impl()     {}

func (m *AnyData) GetImpl() isAnyData_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *AnyData) GetAccAddressData() *AccAddressData {
	if x, ok := m.GetImpl().(*AnyData_AccAddressData); ok {
		return x.AccAddressData
	}
	return nil
}

func (m *AnyData) GetBooleanData() *BooleanData {
	if x, ok := m.GetImpl().(*AnyData_BooleanData); ok {
		return x.BooleanData
	}
	return nil
}

func (m *AnyData) GetDecData() *DecData {
	if x, ok := m.GetImpl().(*AnyData_DecData); ok {
		return x.DecData
	}
	return nil
}

func (m *AnyData) GetHeightData() *HeightData {
	if x, ok := m.GetImpl().(*AnyData_HeightData); ok {
		return x.HeightData
	}
	return nil
}

func (m *AnyData) GetIDData() *IDData {
	if x, ok := m.GetImpl().(*AnyData_IDData); ok {
		return x.IDData
	}
	return nil
}

func (m *AnyData) GetListData() *ListData {
	if x, ok := m.GetImpl().(*AnyData_ListData); ok {
		return x.ListData
	}
	return nil
}

func (m *AnyData) GetNumberData() *NumberData {
	if x, ok := m.GetImpl().(*AnyData_NumberData); ok {
		return x.NumberData
	}
	return nil
}

func (m *AnyData) GetStringData() *StringData {
	if x, ok := m.GetImpl().(*AnyData_StringData); ok {
		return x.StringData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnyData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnyData_AccAddressData)(nil),
		(*AnyData_BooleanData)(nil),
		(*AnyData_DecData)(nil),
		(*AnyData_HeightData)(nil),
		(*AnyData_IDData)(nil),
		(*AnyData_ListData)(nil),
		(*AnyData_NumberData)(nil),
		(*AnyData_StringData)(nil),
	}
}

func init() {
	proto.RegisterType((*AnyData)(nil), "assetmantle.schema.data.base.AnyData")
}

func init() { proto.RegisterFile("data/base/any_data.proto", fileDescriptor_1b19532f46054c21) }

var fileDescriptor_1b19532f46054c21 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x18, 0xdb, 0x3a, 0x5d, 0x44, 0x82, 0xe0, 0xb8, 0x96, 0x6c, 0x11, 0x95, 0x2a,
	0x92, 0x80, 0xde, 0x72, 0x32, 0xa1, 0xe2, 0x8a, 0xee, 0x52, 0xec, 0x45, 0xa4, 0x50, 0x26, 0x93,
	0x21, 0x19, 0x48, 0x32, 0x4b, 0x67, 0xf6, 0xb0, 0xdf, 0x60, 0x8f, 0x7e, 0x04, 0xf1, 0xe8, 0x27,
	0x11, 0x4f, 0x7b, 0xf4, 0x28, 0xed, 0xcd, 0xbb, 0x77, 0x99, 0x99, 0x34, 0x99, 0x5d, 0xd8, 0xb0,
	0xb7, 0xf7, 0xf2, 0x7f, 0xff, 0xdf, 0xcb, 0xbc, 0x79, 0x03, 0x60, 0x8a, 0x04, 0x0a, 0x12, 0xc4,
	0x49, 0x80, 0xaa, 0xb3, 0x95, 0xcc, 0xfc, 0x93, 0x35, 0x13, 0xcc, 0x1d, 0x23, 0xce, 0x89, 0x28,
	0x51, 0x25, 0x0a, 0xe2, 0x73, 0x9c, 0x93, 0x12, 0xf9, 0x4a, 0x96, 0xc5, 0xfb, 0x13, 0xc3, 0x87,
	0xf1, 0x0a, 0xa5, 0xe9, 0x9a, 0x70, 0x6e, 0xf8, 0xf7, 0xc7, 0x6d, 0x45, 0xc2, 0x58, 0x41, 0x50,
	0x65, 0xaa, 0x46, 0xdf, 0x94, 0x60, 0x53, 0x79, 0xd4, 0x2a, 0x39, 0xa1, 0x59, 0x2e, 0x4c, 0xf1,
	0x41, 0x2b, 0xd2, 0xd4, 0x14, 0x1e, 0xb6, 0x42, 0x41, 0xb9, 0xb8, 0x06, 0x58, 0x9d, 0x96, 0x09,
	0x59, 0x5f, 0x23, 0x72, 0xb1, 0xa6, 0x55, 0x66, 0x8a, 0xf7, 0x33, 0x96, 0x31, 0x15, 0x06, 0x32,
	0xd2, 0x5f, 0x1f, 0xff, 0x73, 0xc0, 0x20, 0xaa, 0xce, 0x66, 0x48, 0x20, 0xf7, 0x33, 0xb8, 0x77,
	0xf5, 0xf8, 0xd0, 0x9a, 0x58, 0xd3, 0xd1, 0xab, 0x97, 0x7e, 0xd7, 0xfc, 0xfc, 0x08, 0xe3, 0x48,
	0x9b, 0x24, 0xe7, 0xb0, 0xf7, 0xe9, 0x2e, 0xba, 0xf4, 0xc5, 0x3d, 0x06, 0x7b, 0xe6, 0xd8, 0xa0,
	0xad, 0xa8, 0xcf, 0xbb, 0xa9, 0xb1, 0x76, 0xd4, 0xc8, 0x51, 0xd2, 0xa6, 0x6e, 0x0c, 0x86, 0xbb,
	0x41, 0xc3, 0x5b, 0x8a, 0xf5, 0xb4, 0x9b, 0x35, 0x23, 0xb8, 0xe6, 0x0c, 0x52, 0x1d, 0xba, 0x1f,
	0xc0, 0xc8, 0xb8, 0x12, 0xe8, 0x28, 0xcc, 0xb4, 0x1b, 0x73, 0xa8, 0x0c, 0x35, 0x09, 0xe4, 0x4d,
	0xe6, 0xbe, 0x01, 0x43, 0xba, 0xd2, 0x77, 0x08, 0x6f, 0x2b, 0xd2, 0x93, 0x6e, 0xd2, 0xfb, 0x59,
	0x4d, 0xe9, 0x53, 0x15, 0xb9, 0x6f, 0xc1, 0x9d, 0xe6, 0xae, 0x61, 0x5f, 0x21, 0x9e, 0x75, 0x23,
	0x3e, 0x52, 0xbe, 0xfb, 0x95, 0x61, 0x51, 0xc7, 0xf2, 0x54, 0xc6, 0x5e, 0xc0, 0xc1, 0x4d, 0x4e,
	0x75, 0xac, 0x0c, 0xbb, 0x53, 0x55, 0x4d, 0x26, 0x61, 0xc6, 0x1e, 0xc1, 0xe1, 0x4d, 0x60, 0x0b,
	0x65, 0xd8, 0xc1, 0x78, 0x93, 0x85, 0xce, 0xf9, 0xb7, 0x83, 0x5e, 0xdc, 0x07, 0x0e, 0x2d, 0x4f,
	0x8a, 0xf8, 0xdc, 0xfe, 0xb9, 0xf1, 0xac, 0x8b, 0x8d, 0x67, 0xfd, 0xd9, 0x78, 0xd6, 0xd7, 0xad,
	0xd7, 0xbb, 0xd8, 0x7a, 0xbd, 0xdf, 0x5b, 0xaf, 0x07, 0x26, 0x98, 0x95, 0x9d, 0x3d, 0xe2, 0xbd,
	0x7a, 0x63, 0xe7, 0x72, 0x85, 0xe7, 0xd6, 0x97, 0x17, 0x19, 0x15, 0xf9, 0x69, 0xe2, 0x63, 0x56,
	0x06, 0x91, 0x34, 0x1e, 0x29, 0x63, 0xa0, 0x8d, 0x41, 0xc6, 0x82, 0xe6, 0x61, 0x7c, 0xb7, 0x9d,
	0x68, 0x31, 0x8b, 0x7f, 0xd8, 0xe3, 0xc8, 0xe8, 0xb0, 0xd0, 0x1d, 0x24, 0xd3, 0x8f, 0x11, 0x27,
	0xbf, 0x2e, 0xc9, 0x4b, 0x2d, 0x2f, 0xa5, 0xbc, 0x94, 0xf2, 0xc6, 0x9e, 0x76, 0xc9, 0xcb, 0x77,
	0xf3, 0xf8, 0x88, 0x08, 0x24, 0x9b, 0xfe, 0xb5, 0x0f, 0x8c, 0xd2, 0x30, 0xd4, 0xb5, 0x61, 0xa8,
	0x46, 0x12, 0xca, 0xea, 0xa4, 0xaf, 0x5e, 0xe2, 0xeb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0xd9, 0x99, 0x46, 0xbe, 0x04, 0x00, 0x00,
}

func (m *AnyData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AnyData_AccAddressData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_AccAddressData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AccAddressData != nil {
		{
			size, err := m.AccAddressData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_BooleanData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_BooleanData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BooleanData != nil {
		{
			size, err := m.BooleanData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_DecData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_DecData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DecData != nil {
		{
			size, err := m.DecData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_HeightData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_HeightData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeightData != nil {
		{
			size, err := m.HeightData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_IDData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_IDData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IDData != nil {
		{
			size, err := m.IDData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_ListData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_ListData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ListData != nil {
		{
			size, err := m.ListData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_NumberData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_NumberData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NumberData != nil {
		{
			size, err := m.NumberData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *AnyData_StringData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyData_StringData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringData != nil {
		{
			size, err := m.StringData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAnyData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func encodeVarintAnyData(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnyData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnyData) Size() (n int) {
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

func (m *AnyData_AccAddressData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccAddressData != nil {
		l = m.AccAddressData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_BooleanData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BooleanData != nil {
		l = m.BooleanData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_DecData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DecData != nil {
		l = m.DecData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_HeightData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeightData != nil {
		l = m.HeightData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_IDData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IDData != nil {
		l = m.IDData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_ListData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListData != nil {
		l = m.ListData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_NumberData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberData != nil {
		l = m.NumberData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}
func (m *AnyData_StringData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringData != nil {
		l = m.StringData.Size()
		n += 1 + l + sovAnyData(uint64(l))
	}
	return n
}

func sovAnyData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnyData(x uint64) (n int) {
	return sovAnyData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnyData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnyData
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
			return fmt.Errorf("proto: AnyData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccAddressData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AccAddressData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_AccAddressData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BooleanData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_BooleanData{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DecData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_DecData{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeightData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HeightData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_HeightData{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IDData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_IDData{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ListData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_ListData{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NumberData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_NumberData{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnyData
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
				return ErrInvalidLengthAnyData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &AnyData_StringData{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnyData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnyData
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
func skipAnyData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnyData
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
					return 0, ErrIntOverflowAnyData
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
					return 0, ErrIntOverflowAnyData
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
				return 0, ErrInvalidLengthAnyData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnyData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnyData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnyData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnyData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnyData = fmt.Errorf("proto: unexpected end of group")
)
