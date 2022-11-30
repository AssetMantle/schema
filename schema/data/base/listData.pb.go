// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: schema/data/base/listData.proto

package base

import (
	base "github.com/AssetMantle/modules/schema/lists/base"
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

type ListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *base.List `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ListData) Reset() {
	*x = ListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_data_base_listData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListData) ProtoMessage() {}

func (x *ListData) ProtoReflect() protoreflect.Message {
	mi := &file_schema_data_base_listData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListData.ProtoReflect.Descriptor instead.
func (*ListData) Descriptor() ([]byte, []int) {
	return file_schema_data_base_listData_proto_rawDescGZIP(), []int{0}
}

func (x *ListData) GetValue() *base.List {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_schema_data_base_listData_proto protoreflect.FileDescriptor

var file_schema_data_base_listData_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_data_base_listData_proto_rawDescOnce sync.Once
	file_schema_data_base_listData_proto_rawDescData = file_schema_data_base_listData_proto_rawDesc
)

func file_schema_data_base_listData_proto_rawDescGZIP() []byte {
	file_schema_data_base_listData_proto_rawDescOnce.Do(func() {
		file_schema_data_base_listData_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_data_base_listData_proto_rawDescData)
	})
	return file_schema_data_base_listData_proto_rawDescData
}

var file_schema_data_base_listData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_data_base_listData_proto_goTypes = []interface{}{
	(*ListData)(nil),  // 0: base.ListData
	(*base.List)(nil), // 1: base.List
}
var file_schema_data_base_listData_proto_depIdxs = []int32{
	1, // 0: base.ListData.value:type_name -> base.List
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schema_data_base_listData_proto_init() }
func file_schema_data_base_listData_proto_init() {
	if File_schema_data_base_listData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_data_base_listData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListData); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_data_base_listData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_data_base_listData_proto_goTypes,
		DependencyIndexes: file_schema_data_base_listData_proto_depIdxs,
		MessageInfos:      file_schema_data_base_listData_proto_msgTypes,
	}.Build()
	File_schema_data_base_listData_proto = out.File
	file_schema_data_base_listData_proto_rawDesc = nil
	file_schema_data_base_listData_proto_goTypes = nil
	file_schema_data_base_listData_proto_depIdxs = nil
}
