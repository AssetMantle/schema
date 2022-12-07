// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/data/base/dataI.proto

package base

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Impl:
	//
	//	*DataI_AccAddressData
	//	*DataI_BooleanData
	//	*DataI_DecData
	//	*DataI_HeightData
	//	*DataI_IdData
	//	*DataI_StringData
	Impl isDataI_Impl `protobuf_oneof:"impl"`
}

func (x *DataI) Reset() {
	*x = DataI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_data_base_dataI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataI) ProtoMessage() {}

func (x *DataI) ProtoReflect() protoreflect.Message {
	mi := &file_schema_data_base_dataI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataI.ProtoReflect.Descriptor instead.
func (*DataI) Descriptor() ([]byte, []int) {
	return file_schema_data_base_dataI_proto_rawDescGZIP(), []int{0}
}

func (m *DataI) GetImpl() isDataI_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (x *DataI) GetAccAddressData() *AccAddressData {
	if x, ok := x.GetImpl().(*DataI_AccAddressData); ok {
		return x.AccAddressData
	}
	return nil
}

func (x *DataI) GetBooleanData() *BooleanData {
	if x, ok := x.GetImpl().(*DataI_BooleanData); ok {
		return x.BooleanData
	}
	return nil
}

func (x *DataI) GetDecData() *DecData {
	if x, ok := x.GetImpl().(*DataI_DecData); ok {
		return x.DecData
	}
	return nil
}

func (x *DataI) GetHeightData() *HeightData {
	if x, ok := x.GetImpl().(*DataI_HeightData); ok {
		return x.HeightData
	}
	return nil
}

func (x *DataI) GetIdData() *IDData {
	if x, ok := x.GetImpl().(*DataI_IdData); ok {
		return x.IdData
	}
	return nil
}

func (x *DataI) GetStringData() *StringData {
	if x, ok := x.GetImpl().(*DataI_StringData); ok {
		return x.StringData
	}
	return nil
}

type isDataI_Impl interface {
	isDataI_Impl()
}

type DataI_AccAddressData struct {
	AccAddressData *AccAddressData `protobuf:"bytes,1,opt,name=accAddressData,proto3,oneof"`
}

type DataI_BooleanData struct {
	BooleanData *BooleanData `protobuf:"bytes,2,opt,name=booleanData,proto3,oneof"`
}

type DataI_DecData struct {
	DecData *DecData `protobuf:"bytes,3,opt,name=decData,proto3,oneof"`
}

type DataI_HeightData struct {
	HeightData *HeightData `protobuf:"bytes,4,opt,name=heightData,proto3,oneof"`
}

type DataI_IdData struct {
	IdData *IDData `protobuf:"bytes,5,opt,name=idData,proto3,oneof"`
}

type DataI_StringData struct {
	StringData *StringData `protobuf:"bytes,7,opt,name=stringData,proto3,oneof"`
}

func (*DataI_AccAddressData) isDataI_Impl() {}

func (*DataI_BooleanData) isDataI_Impl() {}

func (*DataI_DecData) isDataI_Impl() {}

func (*DataI_HeightData) isDataI_Impl() {}

func (*DataI_IdData) isDataI_Impl() {}

func (*DataI_StringData) isDataI_Impl() {}

var File_schema_data_base_dataI_proto protoreflect.FileDescriptor

var file_schema_data_base_dataI_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x1a, 0x25, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x64, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x61, 0x49, 0x12, 0x3e,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x32, 0x0a, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x44, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0a,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x06, 0x0a, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x42, 0x77, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x42, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65,
	0xca, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0xe2, 0x02, 0x10, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_data_base_dataI_proto_rawDescOnce sync.Once
	file_schema_data_base_dataI_proto_rawDescData = file_schema_data_base_dataI_proto_rawDesc
)

func file_schema_data_base_dataI_proto_rawDescGZIP() []byte {
	file_schema_data_base_dataI_proto_rawDescOnce.Do(func() {
		file_schema_data_base_dataI_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_data_base_dataI_proto_rawDescData)
	})
	return file_schema_data_base_dataI_proto_rawDescData
}

var file_schema_data_base_dataI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_data_base_dataI_proto_goTypes = []interface{}{
	(*DataI)(nil),          // 0: base.DataI
	(*AccAddressData)(nil), // 1: base.AccAddressData
	(*BooleanData)(nil),    // 2: base.BooleanData
	(*DecData)(nil),        // 3: base.DecData
	(*HeightData)(nil),     // 4: base.HeightData
	(*IDData)(nil),         // 5: base.IDData
	(*StringData)(nil),     // 6: base.StringData
}
var file_schema_data_base_dataI_proto_depIdxs = []int32{
	1, // 0: base.DataI.accAddressData:type_name -> base.AccAddressData
	2, // 1: base.DataI.booleanData:type_name -> base.BooleanData
	3, // 2: base.DataI.decData:type_name -> base.DecData
	4, // 3: base.DataI.heightData:type_name -> base.HeightData
	5, // 4: base.DataI.idData:type_name -> base.IDData
	6, // 5: base.DataI.stringData:type_name -> base.StringData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_schema_data_base_dataI_proto_init() }
func file_schema_data_base_dataI_proto_init() {
	if File_schema_data_base_dataI_proto != nil {
		return
	}
	file_schema_data_base_accAddressData_proto_init()
	file_schema_data_base_booleanData_proto_init()
	file_schema_data_base_decData_proto_init()
	file_schema_data_base_heightData_proto_init()
	file_schema_data_base_idData_proto_init()
	file_schema_data_base_stringData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_data_base_dataI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataI); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_schema_data_base_dataI_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataI_AccAddressData)(nil),
		(*DataI_BooleanData)(nil),
		(*DataI_DecData)(nil),
		(*DataI_HeightData)(nil),
		(*DataI_IdData)(nil),
		(*DataI_StringData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_data_base_dataI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_data_base_dataI_proto_goTypes,
		DependencyIndexes: file_schema_data_base_dataI_proto_depIdxs,
		MessageInfos:      file_schema_data_base_dataI_proto_msgTypes,
	}.Build()
	File_schema_data_base_dataI_proto = out.File
	file_schema_data_base_dataI_proto_rawDesc = nil
	file_schema_data_base_dataI_proto_goTypes = nil
	file_schema_data_base_dataI_proto_depIdxs = nil
}
