// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: schema/ids/base/id.proto

package base

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	base "schema/ids/base"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Impl:
	//	*ID_AssetId
	//	*ID_ClassificationId
	//	*ID_DataId
	//	*ID_HashId
	//	*ID_IdentityId
	//	*ID_MaintainerId
	//	*ID_OrderId
	//	*ID_OwnableId
	//	*ID_PropertyId
	//	*ID_SplitId
	//	*ID_StringId
	Impl isID_Impl `protobuf_oneof:"impl"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_ids_base_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_schema_ids_base_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_schema_ids_base_id_proto_rawDescGZIP(), []int{0}
}

func (m *ID) GetImpl() isID_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (x *ID) GetAssetId() *AssetID {
	if x, ok := x.GetImpl().(*ID_AssetId); ok {
		return x.AssetId
	}
	return nil
}

func (x *ID) GetClassificationId() *ClassificationID {
	if x, ok := x.GetImpl().(*ID_ClassificationId); ok {
		return x.ClassificationId
	}
	return nil
}

func (x *ID) GetDataId() *DataID {
	if x, ok := x.GetImpl().(*ID_DataId); ok {
		return x.DataId
	}
	return nil
}

func (x *ID) GetHashId() *base.HashID {
	if x, ok := x.GetImpl().(*ID_HashId); ok {
		return x.HashId
	}
	return nil
}

func (x *ID) GetIdentityId() *IdentityID {
	if x, ok := x.GetImpl().(*ID_IdentityId); ok {
		return x.IdentityId
	}
	return nil
}

func (x *ID) GetMaintainerId() *MaintainerID {
	if x, ok := x.GetImpl().(*ID_MaintainerId); ok {
		return x.MaintainerId
	}
	return nil
}

func (x *ID) GetOrderId() *OrderID {
	if x, ok := x.GetImpl().(*ID_OrderId); ok {
		return x.OrderId
	}
	return nil
}

func (x *ID) GetOwnableId() *OwnableID {
	if x, ok := x.GetImpl().(*ID_OwnableId); ok {
		return x.OwnableId
	}
	return nil
}

func (x *ID) GetPropertyId() *PropertyID {
	if x, ok := x.GetImpl().(*ID_PropertyId); ok {
		return x.PropertyId
	}
	return nil
}

func (x *ID) GetSplitId() *SplitID {
	if x, ok := x.GetImpl().(*ID_SplitId); ok {
		return x.SplitId
	}
	return nil
}

func (x *ID) GetStringId() *base.StringID {
	if x, ok := x.GetImpl().(*ID_StringId); ok {
		return x.StringId
	}
	return nil
}

type isID_Impl interface {
	isID_Impl()
}

type ID_AssetId struct {
	AssetId *AssetID `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3,oneof"`
}

type ID_ClassificationId struct {
	ClassificationId *ClassificationID `protobuf:"bytes,2,opt,name=classification_id,json=classificationId,proto3,oneof"`
}

type ID_DataId struct {
	DataId *DataID `protobuf:"bytes,3,opt,name=data_id,json=dataId,proto3,oneof"`
}

type ID_HashId struct {
	HashId *base.HashID `protobuf:"bytes,4,opt,name=hash_id,json=hashId,proto3,oneof"`
}

type ID_IdentityId struct {
	IdentityId *IdentityID `protobuf:"bytes,5,opt,name=identity_id,json=identityId,proto3,oneof"`
}

type ID_MaintainerId struct {
	MaintainerId *MaintainerID `protobuf:"bytes,6,opt,name=maintainer_id,json=maintainerId,proto3,oneof"`
}

type ID_OrderId struct {
	OrderId *OrderID `protobuf:"bytes,7,opt,name=order_id,json=orderId,proto3,oneof"`
}

type ID_OwnableId struct {
	OwnableId *OwnableID `protobuf:"bytes,8,opt,name=ownable_id,json=ownableId,proto3,oneof"`
}

type ID_PropertyId struct {
	PropertyId *PropertyID `protobuf:"bytes,9,opt,name=property_id,json=propertyId,proto3,oneof"`
}

type ID_SplitId struct {
	SplitId *SplitID `protobuf:"bytes,10,opt,name=split_id,json=splitId,proto3,oneof"`
}

type ID_StringId struct {
	StringId *base.StringID `protobuf:"bytes,11,opt,name=string_id,json=stringId,proto3,oneof"`
}

func (*ID_AssetId) isID_Impl() {}

func (*ID_ClassificationId) isID_Impl() {}

func (*ID_DataId) isID_Impl() {}

func (*ID_HashId) isID_Impl() {}

func (*ID_IdentityId) isID_Impl() {}

func (*ID_MaintainerId) isID_Impl() {}

func (*ID_OrderId) isID_Impl() {}

func (*ID_OwnableId) isID_Impl() {}

func (*ID_PropertyId) isID_Impl() {}

func (*ID_SplitId) isID_Impl() {}

func (*ID_StringId) isID_Impl() {}

var File_schema_ids_base_id_proto protoreflect.FileDescriptor

var file_schema_ids_base_id_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x1a, 0x1d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x49, 0x44, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6f, 0x77, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x02, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x49, 0x44, 0x48, 0x00, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a,
	0x11, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x48, 0x00, 0x52, 0x10, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x07, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x48, 0x00, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0d, 0x6d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x77,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x07, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_ids_base_id_proto_rawDescOnce sync.Once
	file_schema_ids_base_id_proto_rawDescData = file_schema_ids_base_id_proto_rawDesc
)

func file_schema_ids_base_id_proto_rawDescGZIP() []byte {
	file_schema_ids_base_id_proto_rawDescOnce.Do(func() {
		file_schema_ids_base_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_ids_base_id_proto_rawDescData)
	})
	return file_schema_ids_base_id_proto_rawDescData
}

var file_schema_ids_base_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_ids_base_id_proto_goTypes = []interface{}{
	(*ID)(nil),               // 0: base.ID
	(*AssetID)(nil),          // 1: base.AssetID
	(*ClassificationID)(nil), // 2: base.ClassificationID
	(*DataID)(nil),           // 3: base.DataID
	(*base.HashID)(nil),      // 4: base.HashID
	(*IdentityID)(nil),       // 5: base.IdentityID
	(*MaintainerID)(nil),     // 6: base.MaintainerID
	(*OrderID)(nil),          // 7: base.OrderID
	(*OwnableID)(nil),        // 8: base.OwnableID
	(*PropertyID)(nil),       // 9: base.PropertyID
	(*SplitID)(nil),          // 10: base.SplitID
	(*base.StringID)(nil),    // 11: base.StringID
}
var file_schema_ids_base_id_proto_depIdxs = []int32{
	1,  // 0: base.ID.asset_id:type_name -> base.AssetID
	2,  // 1: base.ID.classification_id:type_name -> base.ClassificationID
	3,  // 2: base.ID.data_id:type_name -> base.DataID
	4,  // 3: base.ID.hash_id:type_name -> base.HashID
	5,  // 4: base.ID.identity_id:type_name -> base.IdentityID
	6,  // 5: base.ID.maintainer_id:type_name -> base.MaintainerID
	7,  // 6: base.ID.order_id:type_name -> base.OrderID
	8,  // 7: base.ID.ownable_id:type_name -> base.OwnableID
	9,  // 8: base.ID.property_id:type_name -> base.PropertyID
	10, // 9: base.ID.split_id:type_name -> base.SplitID
	11, // 10: base.ID.string_id:type_name -> base.StringID
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_schema_ids_base_id_proto_init() }
func file_schema_ids_base_id_proto_init() {
	if File_schema_ids_base_id_proto != nil {
		return
	}
	file_schema_ids_base_assetID_proto_init()
	file_schema_ids_base_classificationID_proto_init()
	file_schema_ids_base_dataID_proto_init()
	file_schema_ids_base_identityID_proto_init()
	file_schema_ids_base_maintainerID_proto_init()
	file_schema_ids_base_orderID_proto_init()
	file_schema_ids_base_ownableID_proto_init()
	file_schema_ids_base_propertyID_proto_init()
	file_schema_ids_base_splitID_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_ids_base_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
	file_schema_ids_base_id_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ID_AssetId)(nil),
		(*ID_ClassificationId)(nil),
		(*ID_DataId)(nil),
		(*ID_HashId)(nil),
		(*ID_IdentityId)(nil),
		(*ID_MaintainerId)(nil),
		(*ID_OrderId)(nil),
		(*ID_OwnableId)(nil),
		(*ID_PropertyId)(nil),
		(*ID_SplitId)(nil),
		(*ID_StringId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_ids_base_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_ids_base_id_proto_goTypes,
		DependencyIndexes: file_schema_ids_base_id_proto_depIdxs,
		MessageInfos:      file_schema_ids_base_id_proto_msgTypes,
	}.Build()
	File_schema_ids_base_id_proto = out.File
	file_schema_ids_base_id_proto_rawDesc = nil
	file_schema_ids_base_id_proto_goTypes = nil
	file_schema_ids_base_id_proto_depIdxs = nil
}
