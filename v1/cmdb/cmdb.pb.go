// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: src/cmdb/cmdb.proto

package cmdb

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

type DataType int32

const (
	DataType_UNKNOWN     DataType = 0
	DataType_INT         DataType = 1
	DataType_LONG        DataType = 2
	DataType_NUMBER      DataType = 3
	DataType_ENUM        DataType = 4
	DataType_BOOLEAN     DataType = 5
	DataType_DOUBLE      DataType = 6
	DataType_JSONB       DataType = 7
	DataType_STRING      DataType = 8
	DataType_ARRAY       DataType = 9
	DataType_PASSWORD    DataType = 10
	DataType_UUID        DataType = 11
	DataType_DATE        DataType = 12
	DataType_TIME_ZONE   DataType = 13
	DataType_INET        DataType = 14
	DataType_CIDR        DataType = 15
	DataType_MAC_ADDR    DataType = 16
	DataType_OBJECT      DataType = 17
	DataType_LONG_STRING DataType = 18
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "INT",
		2:  "LONG",
		3:  "NUMBER",
		4:  "ENUM",
		5:  "BOOLEAN",
		6:  "DOUBLE",
		7:  "JSONB",
		8:  "STRING",
		9:  "ARRAY",
		10: "PASSWORD",
		11: "UUID",
		12: "DATE",
		13: "TIME_ZONE",
		14: "INET",
		15: "CIDR",
		16: "MAC_ADDR",
		17: "OBJECT",
		18: "LONG_STRING",
	}
	DataType_value = map[string]int32{
		"UNKNOWN":     0,
		"INT":         1,
		"LONG":        2,
		"NUMBER":      3,
		"ENUM":        4,
		"BOOLEAN":     5,
		"DOUBLE":      6,
		"JSONB":       7,
		"STRING":      8,
		"ARRAY":       9,
		"PASSWORD":    10,
		"UUID":        11,
		"DATE":        12,
		"TIME_ZONE":   13,
		"INET":        14,
		"CIDR":        15,
		"MAC_ADDR":    16,
		"OBJECT":      17,
		"LONG_STRING": 18,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_cmdb_cmdb_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_src_cmdb_cmdb_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{0}
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label      string `protobuf:"bytes,10,opt,name=label,proto3" json:"label,omitempty"`
	Properties string `protobuf:"bytes,20,opt,name=properties,proto3" json:"properties,omitempty"` // json string
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Node) GetProperties() string {
	if x != nil {
		return x.Properties
	}
	return ""
}

type RelationProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field        string `protobuf:"bytes,10,opt,name=field,proto3" json:"field,omitempty"`
	RefModelId   string `protobuf:"bytes,20,opt,name=ref_model_id,json=refModelId,proto3" json:"ref_model_id,omitempty"`
	RefModelName string `protobuf:"bytes,30,opt,name=ref_model_name,json=refModelName,proto3" json:"ref_model_name,omitempty"`
	RefField     string `protobuf:"bytes,40,opt,name=ref_field,json=refField,proto3" json:"ref_field,omitempty"`
}

func (x *RelationProperty) Reset() {
	*x = RelationProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationProperty) ProtoMessage() {}

func (x *RelationProperty) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationProperty.ProtoReflect.Descriptor instead.
func (*RelationProperty) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{1}
}

func (x *RelationProperty) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *RelationProperty) GetRefModelId() string {
	if x != nil {
		return x.RefModelId
	}
	return ""
}

func (x *RelationProperty) GetRefModelName() string {
	if x != nil {
		return x.RefModelName
	}
	return ""
}

func (x *RelationProperty) GetRefField() string {
	if x != nil {
		return x.RefField
	}
	return ""
}

type FieldPublicOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Editable    bool   `protobuf:"varint,10,opt,name=editable,proto3" json:"editable,omitempty"`
	Searchable  bool   `protobuf:"varint,20,opt,name=searchable,proto3" json:"searchable,omitempty"`
	Required    bool   `protobuf:"varint,30,opt,name=required,proto3" json:"required,omitempty"`
	Description string `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FieldPublicOption) Reset() {
	*x = FieldPublicOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldPublicOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldPublicOption) ProtoMessage() {}

func (x *FieldPublicOption) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldPublicOption.ProtoReflect.Descriptor instead.
func (*FieldPublicOption) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{2}
}

func (x *FieldPublicOption) GetEditable() bool {
	if x != nil {
		return x.Editable
	}
	return false
}

func (x *FieldPublicOption) GetSearchable() bool {
	if x != nil {
		return x.Searchable
	}
	return false
}

func (x *FieldPublicOption) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FieldPublicOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FieldGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Index     int32    `protobuf:"varint,30,opt,name=index,proto3" json:"index,omitempty"`
	Removable bool     `protobuf:"varint,40,opt,name=removable,proto3" json:"removable,omitempty"`
	Collapse  bool     `protobuf:"varint,50,opt,name=collapse,proto3" json:"collapse,omitempty"`
	Fields    []*Field `protobuf:"bytes,60,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *FieldGroup) Reset() {
	*x = FieldGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldGroup) ProtoMessage() {}

func (x *FieldGroup) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldGroup.ProtoReflect.Descriptor instead.
func (*FieldGroup) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{3}
}

func (x *FieldGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FieldGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FieldGroup) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *FieldGroup) GetRemovable() bool {
	if x != nil {
		return x.Removable
	}
	return false
}

func (x *FieldGroup) GetCollapse() bool {
	if x != nil {
		return x.Collapse
	}
	return false
}

func (x *FieldGroup) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId           string             `protobuf:"bytes,10,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PropertyGroupId   string             `protobuf:"bytes,20,opt,name=property_group_id,json=propertyGroupId,proto3" json:"property_group_id,omitempty"`
	PropertyId        string             `protobuf:"bytes,30,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	PropertyName      string             `protobuf:"bytes,40,opt,name=property_name,json=propertyName,proto3" json:"property_name,omitempty"`
	PropertyType      DataType           `protobuf:"varint,50,opt,name=property_type,json=propertyType,proto3,enum=cmdb.DataType" json:"property_type,omitempty"`
	Unit              string             `protobuf:"bytes,60,opt,name=unit,proto3" json:"unit,omitempty"`
	PrimaryKey        bool               `protobuf:"varint,70,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	FieldPublicOption *FieldPublicOption `protobuf:"bytes,80,opt,name=field_public_option,json=fieldPublicOption,proto3" json:"field_public_option,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{4}
}

func (x *Field) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *Field) GetPropertyGroupId() string {
	if x != nil {
		return x.PropertyGroupId
	}
	return ""
}

func (x *Field) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *Field) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *Field) GetPropertyType() DataType {
	if x != nil {
		return x.PropertyType
	}
	return DataType_UNKNOWN
}

func (x *Field) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Field) GetPrimaryKey() bool {
	if x != nil {
		return x.PrimaryKey
	}
	return false
}

func (x *Field) GetFieldPublicOption() *FieldPublicOption {
	if x != nil {
		return x.FieldPublicOption
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string        `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Name        string        `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	ModelIcon   string        `protobuf:"bytes,30,opt,name=model_icon,json=modelIcon,proto3" json:"model_icon,omitempty"`
	Description string        `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
	GroupId     string        `protobuf:"bytes,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	FieldGroups []*FieldGroup `protobuf:"bytes,60,rep,name=field_groups,json=fieldGroups,proto3" json:"field_groups,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{5}
}

func (x *Table) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetModelIcon() string {
	if x != nil {
		return x.ModelIcon
	}
	return ""
}

func (x *Table) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Table) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Table) GetFieldGroups() []*FieldGroup {
	if x != nil {
		return x.FieldGroups
	}
	return nil
}

type TableGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Name        string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,30,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TableGroup) Reset() {
	*x = TableGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableGroup) ProtoMessage() {}

func (x *TableGroup) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableGroup.ProtoReflect.Descriptor instead.
func (*TableGroup) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{6}
}

func (x *TableGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TableGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type NodeAndParent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Properties       string `protobuf:"bytes,10,opt,name=properties,proto3" json:"properties,omitempty"`
	ParentProperties string `protobuf:"bytes,20,opt,name=parent_properties,json=parentProperties,proto3" json:"parent_properties,omitempty"`
}

func (x *NodeAndParent) Reset() {
	*x = NodeAndParent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_cmdb_cmdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAndParent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAndParent) ProtoMessage() {}

func (x *NodeAndParent) ProtoReflect() protoreflect.Message {
	mi := &file_src_cmdb_cmdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAndParent.ProtoReflect.Descriptor instead.
func (*NodeAndParent) Descriptor() ([]byte, []int) {
	return file_src_cmdb_cmdb_proto_rawDescGZIP(), []int{7}
}

func (x *NodeAndParent) GetProperties() string {
	if x != nil {
		return x.Properties
	}
	return ""
}

func (x *NodeAndParent) GetParentProperties() string {
	if x != nil {
		return x.ParentProperties
	}
	return ""
}

var File_src_cmdb_cmdb_proto protoreflect.FileDescriptor

var file_src_cmdb_cmdb_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6d, 0x64, 0x62, 0x22, 0x3c, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0xc7, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x46, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x47, 0x0a, 0x13, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x05,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x54, 0x0a, 0x0a,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x2a, 0xeb, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x4e, 0x55,
	0x4d, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x09, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x55, 0x49, 0x44, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0c,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x0d, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x4e, 0x45, 0x54, 0x10, 0x0e, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x49, 0x44,
	0x52, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x10,
	0x10, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x11, 0x12, 0x0f, 0x0a,
	0x0b, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x12, 0x42, 0x41,
	0x0a, 0x17, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6d, 0x64,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_cmdb_cmdb_proto_rawDescOnce sync.Once
	file_src_cmdb_cmdb_proto_rawDescData = file_src_cmdb_cmdb_proto_rawDesc
)

func file_src_cmdb_cmdb_proto_rawDescGZIP() []byte {
	file_src_cmdb_cmdb_proto_rawDescOnce.Do(func() {
		file_src_cmdb_cmdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_cmdb_cmdb_proto_rawDescData)
	})
	return file_src_cmdb_cmdb_proto_rawDescData
}

var file_src_cmdb_cmdb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_cmdb_cmdb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_src_cmdb_cmdb_proto_goTypes = []interface{}{
	(DataType)(0),             // 0: cmdb.DataType
	(*Node)(nil),              // 1: cmdb.Node
	(*RelationProperty)(nil),  // 2: cmdb.RelationProperty
	(*FieldPublicOption)(nil), // 3: cmdb.FieldPublicOption
	(*FieldGroup)(nil),        // 4: cmdb.FieldGroup
	(*Field)(nil),             // 5: cmdb.Field
	(*Table)(nil),             // 6: cmdb.Table
	(*TableGroup)(nil),        // 7: cmdb.TableGroup
	(*NodeAndParent)(nil),     // 8: cmdb.NodeAndParent
}
var file_src_cmdb_cmdb_proto_depIdxs = []int32{
	5, // 0: cmdb.FieldGroup.fields:type_name -> cmdb.Field
	0, // 1: cmdb.Field.property_type:type_name -> cmdb.DataType
	3, // 2: cmdb.Field.field_public_option:type_name -> cmdb.FieldPublicOption
	4, // 3: cmdb.Table.field_groups:type_name -> cmdb.FieldGroup
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_src_cmdb_cmdb_proto_init() }
func file_src_cmdb_cmdb_proto_init() {
	if File_src_cmdb_cmdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_cmdb_cmdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationProperty); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldPublicOption); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldGroup); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableGroup); i {
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
		file_src_cmdb_cmdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAndParent); i {
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
			RawDescriptor: file_src_cmdb_cmdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_cmdb_cmdb_proto_goTypes,
		DependencyIndexes: file_src_cmdb_cmdb_proto_depIdxs,
		EnumInfos:         file_src_cmdb_cmdb_proto_enumTypes,
		MessageInfos:      file_src_cmdb_cmdb_proto_msgTypes,
	}.Build()
	File_src_cmdb_cmdb_proto = out.File
	file_src_cmdb_cmdb_proto_rawDesc = nil
	file_src_cmdb_cmdb_proto_goTypes = nil
	file_src_cmdb_cmdb_proto_depIdxs = nil
}
