// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/properties/base/mesaProperty.proto

package base

import (
	base1 "github.com/AssetMantle/modules/schema/data/base"
	base "github.com/AssetMantle/modules/schema/ids/base"
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

type MesaProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId *base.PropertyID `protobuf:"bytes,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	Data       *base1.Data      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MesaProperty) Reset() {
	*x = MesaProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_properties_base_mesaProperty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MesaProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MesaProperty) ProtoMessage() {}

func (x *MesaProperty) ProtoReflect() protoreflect.Message {
	mi := &file_schema_properties_base_mesaProperty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MesaProperty.ProtoReflect.Descriptor instead.
func (*MesaProperty) Descriptor() ([]byte, []int) {
	return file_schema_properties_base_mesaProperty_proto_rawDescGZIP(), []int{0}
}

func (x *MesaProperty) GetPropertyId() *base.PropertyID {
	if x != nil {
		return x.PropertyId
	}
	return nil
}

func (x *MesaProperty) GetData() *base1.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_schema_properties_base_mesaProperty_proto protoreflect.FileDescriptor

var file_schema_properties_base_mesaProperty_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x61, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x1a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x61, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x31, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x44, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x84, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x42, 0x11, 0x4d, 0x65, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0xa2, 0x02, 0x03, 0x42,
	0x58, 0x58, 0xaa, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0xca, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65,
	0xe2, 0x02, 0x10, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schema_properties_base_mesaProperty_proto_rawDescOnce sync.Once
	file_schema_properties_base_mesaProperty_proto_rawDescData = file_schema_properties_base_mesaProperty_proto_rawDesc
)

func file_schema_properties_base_mesaProperty_proto_rawDescGZIP() []byte {
	file_schema_properties_base_mesaProperty_proto_rawDescOnce.Do(func() {
		file_schema_properties_base_mesaProperty_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_properties_base_mesaProperty_proto_rawDescData)
	})
	return file_schema_properties_base_mesaProperty_proto_rawDescData
}

var file_schema_properties_base_mesaProperty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_properties_base_mesaProperty_proto_goTypes = []interface{}{
	(*MesaProperty)(nil),    // 0: base.MesaProperty
	(*base.PropertyID)(nil), // 1: base.PropertyID
	(*base1.Data)(nil),      // 2: base.Data
}
var file_schema_properties_base_mesaProperty_proto_depIdxs = []int32{
	1, // 0: base.MesaProperty.property_id:type_name -> base.PropertyID
	2, // 1: base.MesaProperty.data:type_name -> base.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_properties_base_mesaProperty_proto_init() }
func file_schema_properties_base_mesaProperty_proto_init() {
	if File_schema_properties_base_mesaProperty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_properties_base_mesaProperty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MesaProperty); i {
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
			RawDescriptor: file_schema_properties_base_mesaProperty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_properties_base_mesaProperty_proto_goTypes,
		DependencyIndexes: file_schema_properties_base_mesaProperty_proto_depIdxs,
		MessageInfos:      file_schema_properties_base_mesaProperty_proto_msgTypes,
	}.Build()
	File_schema_properties_base_mesaProperty_proto = out.File
	file_schema_properties_base_mesaProperty_proto_rawDesc = nil
	file_schema_properties_base_mesaProperty_proto_goTypes = nil
	file_schema_properties_base_mesaProperty_proto_depIdxs = nil
}
