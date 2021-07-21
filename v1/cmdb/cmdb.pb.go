// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/cmdb/cmdb.proto

package cmdb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

var DataType_name = map[int32]string{
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

var DataType_value = map[string]int32{
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

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{0}
}

type Node struct {
	Label                string   `protobuf:"bytes,10,opt,name=label,proto3" json:"label,omitempty"`
	Properties           string   `protobuf:"bytes,20,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Node) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

type RelationProperty struct {
	Field                string   `protobuf:"bytes,10,opt,name=field,proto3" json:"field,omitempty"`
	RefModelId           string   `protobuf:"bytes,20,opt,name=ref_model_id,json=refModelId,proto3" json:"ref_model_id,omitempty"`
	RefModelName         string   `protobuf:"bytes,30,opt,name=ref_model_name,json=refModelName,proto3" json:"ref_model_name,omitempty"`
	RefField             string   `protobuf:"bytes,40,opt,name=ref_field,json=refField,proto3" json:"ref_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationProperty) Reset()         { *m = RelationProperty{} }
func (m *RelationProperty) String() string { return proto.CompactTextString(m) }
func (*RelationProperty) ProtoMessage()    {}
func (*RelationProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{1}
}

func (m *RelationProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationProperty.Unmarshal(m, b)
}
func (m *RelationProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationProperty.Marshal(b, m, deterministic)
}
func (m *RelationProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationProperty.Merge(m, src)
}
func (m *RelationProperty) XXX_Size() int {
	return xxx_messageInfo_RelationProperty.Size(m)
}
func (m *RelationProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationProperty.DiscardUnknown(m)
}

var xxx_messageInfo_RelationProperty proto.InternalMessageInfo

func (m *RelationProperty) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RelationProperty) GetRefModelId() string {
	if m != nil {
		return m.RefModelId
	}
	return ""
}

func (m *RelationProperty) GetRefModelName() string {
	if m != nil {
		return m.RefModelName
	}
	return ""
}

func (m *RelationProperty) GetRefField() string {
	if m != nil {
		return m.RefField
	}
	return ""
}

type FieldPublicOption struct {
	Editable             bool     `protobuf:"varint,10,opt,name=editable,proto3" json:"editable,omitempty"`
	Searchable           bool     `protobuf:"varint,20,opt,name=searchable,proto3" json:"searchable,omitempty"`
	Required             bool     `protobuf:"varint,30,opt,name=required,proto3" json:"required,omitempty"`
	Description          string   `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldPublicOption) Reset()         { *m = FieldPublicOption{} }
func (m *FieldPublicOption) String() string { return proto.CompactTextString(m) }
func (*FieldPublicOption) ProtoMessage()    {}
func (*FieldPublicOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{2}
}

func (m *FieldPublicOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldPublicOption.Unmarshal(m, b)
}
func (m *FieldPublicOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldPublicOption.Marshal(b, m, deterministic)
}
func (m *FieldPublicOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldPublicOption.Merge(m, src)
}
func (m *FieldPublicOption) XXX_Size() int {
	return xxx_messageInfo_FieldPublicOption.Size(m)
}
func (m *FieldPublicOption) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldPublicOption.DiscardUnknown(m)
}

var xxx_messageInfo_FieldPublicOption proto.InternalMessageInfo

func (m *FieldPublicOption) GetEditable() bool {
	if m != nil {
		return m.Editable
	}
	return false
}

func (m *FieldPublicOption) GetSearchable() bool {
	if m != nil {
		return m.Searchable
	}
	return false
}

func (m *FieldPublicOption) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *FieldPublicOption) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type FieldGroup struct {
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Index                int32    `protobuf:"varint,30,opt,name=index,proto3" json:"index,omitempty"`
	Removable            bool     `protobuf:"varint,40,opt,name=removable,proto3" json:"removable,omitempty"`
	Collapse             bool     `protobuf:"varint,50,opt,name=collapse,proto3" json:"collapse,omitempty"`
	Fields               []*Field `protobuf:"bytes,60,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldGroup) Reset()         { *m = FieldGroup{} }
func (m *FieldGroup) String() string { return proto.CompactTextString(m) }
func (*FieldGroup) ProtoMessage()    {}
func (*FieldGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{3}
}

func (m *FieldGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldGroup.Unmarshal(m, b)
}
func (m *FieldGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldGroup.Marshal(b, m, deterministic)
}
func (m *FieldGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldGroup.Merge(m, src)
}
func (m *FieldGroup) XXX_Size() int {
	return xxx_messageInfo_FieldGroup.Size(m)
}
func (m *FieldGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldGroup.DiscardUnknown(m)
}

var xxx_messageInfo_FieldGroup proto.InternalMessageInfo

func (m *FieldGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FieldGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldGroup) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *FieldGroup) GetRemovable() bool {
	if m != nil {
		return m.Removable
	}
	return false
}

func (m *FieldGroup) GetCollapse() bool {
	if m != nil {
		return m.Collapse
	}
	return false
}

func (m *FieldGroup) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	ModelId              string             `protobuf:"bytes,10,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PropertyGroupId      string             `protobuf:"bytes,20,opt,name=property_group_id,json=propertyGroupId,proto3" json:"property_group_id,omitempty"`
	PropertyId           string             `protobuf:"bytes,30,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	PropertyName         string             `protobuf:"bytes,40,opt,name=property_name,json=propertyName,proto3" json:"property_name,omitempty"`
	PropertyType         DataType           `protobuf:"varint,50,opt,name=property_type,json=propertyType,proto3,enum=cmdb.DataType" json:"property_type,omitempty"`
	Unit                 string             `protobuf:"bytes,60,opt,name=unit,proto3" json:"unit,omitempty"`
	PrimaryKey           bool               `protobuf:"varint,70,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	FieldPublicOption    *FieldPublicOption `protobuf:"bytes,80,opt,name=field_public_option,json=fieldPublicOption,proto3" json:"field_public_option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{4}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *Field) GetPropertyGroupId() string {
	if m != nil {
		return m.PropertyGroupId
	}
	return ""
}

func (m *Field) GetPropertyId() string {
	if m != nil {
		return m.PropertyId
	}
	return ""
}

func (m *Field) GetPropertyName() string {
	if m != nil {
		return m.PropertyName
	}
	return ""
}

func (m *Field) GetPropertyType() DataType {
	if m != nil {
		return m.PropertyType
	}
	return DataType_UNKNOWN
}

func (m *Field) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Field) GetPrimaryKey() bool {
	if m != nil {
		return m.PrimaryKey
	}
	return false
}

func (m *Field) GetFieldPublicOption() *FieldPublicOption {
	if m != nil {
		return m.FieldPublicOption
	}
	return nil
}

type Table struct {
	Key                  string        `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string        `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	ModelIcon            string        `protobuf:"bytes,30,opt,name=model_icon,json=modelIcon,proto3" json:"model_icon,omitempty"`
	Description          string        `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
	GroupId              string        `protobuf:"bytes,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	FieldGroups          []*FieldGroup `protobuf:"bytes,60,rep,name=field_groups,json=fieldGroups,proto3" json:"field_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{5}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetModelIcon() string {
	if m != nil {
		return m.ModelIcon
	}
	return ""
}

func (m *Table) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Table) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Table) GetFieldGroups() []*FieldGroup {
	if m != nil {
		return m.FieldGroups
	}
	return nil
}

type TableGroup struct {
	Key                  string   `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,30,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableGroup) Reset()         { *m = TableGroup{} }
func (m *TableGroup) String() string { return proto.CompactTextString(m) }
func (*TableGroup) ProtoMessage()    {}
func (*TableGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_1255f1c24c5b70ca, []int{6}
}

func (m *TableGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableGroup.Unmarshal(m, b)
}
func (m *TableGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableGroup.Marshal(b, m, deterministic)
}
func (m *TableGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableGroup.Merge(m, src)
}
func (m *TableGroup) XXX_Size() int {
	return xxx_messageInfo_TableGroup.Size(m)
}
func (m *TableGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_TableGroup.DiscardUnknown(m)
}

var xxx_messageInfo_TableGroup proto.InternalMessageInfo

func (m *TableGroup) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TableGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("cmdb.DataType", DataType_name, DataType_value)
	proto.RegisterType((*Node)(nil), "cmdb.Node")
	proto.RegisterType((*RelationProperty)(nil), "cmdb.RelationProperty")
	proto.RegisterType((*FieldPublicOption)(nil), "cmdb.FieldPublicOption")
	proto.RegisterType((*FieldGroup)(nil), "cmdb.FieldGroup")
	proto.RegisterType((*Field)(nil), "cmdb.Field")
	proto.RegisterType((*Table)(nil), "cmdb.Table")
	proto.RegisterType((*TableGroup)(nil), "cmdb.TableGroup")
}

func init() { proto.RegisterFile("src/cmdb/cmdb.proto", fileDescriptor_1255f1c24c5b70ca) }

var fileDescriptor_1255f1c24c5b70ca = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x8e, 0xda, 0x56,
	0x10, 0x2e, 0xcb, 0x9f, 0x19, 0xb3, 0xec, 0xe1, 0x64, 0xa5, 0xd0, 0xbf, 0x2d, 0x22, 0x55, 0x85,
	0x22, 0x05, 0x54, 0xf6, 0x76, 0x6f, 0xcc, 0xda, 0x41, 0x4e, 0x16, 0x1b, 0x19, 0xa3, 0xa8, 0xb9,
	0xb1, 0x8c, 0x7d, 0xd8, 0x58, 0x6b, 0x63, 0xd7, 0x36, 0x51, 0xfd, 0x10, 0x7d, 0x8c, 0xbe, 0x46,
	0x6f, 0xfb, 0x30, 0x7d, 0x89, 0xea, 0xcc, 0xb1, 0x01, 0x6d, 0x2b, 0x35, 0x37, 0x68, 0xe6, 0x9b,
	0xf1, 0x99, 0x6f, 0xbe, 0x6f, 0x04, 0xbc, 0xc8, 0x52, 0x6f, 0xea, 0x45, 0xfe, 0x16, 0x7f, 0x26,
	0x49, 0x1a, 0xe7, 0x31, 0x6d, 0xf0, 0x78, 0x74, 0x07, 0x0d, 0x23, 0xf6, 0x19, 0xbd, 0x86, 0x66,
	0xe8, 0x6e, 0x59, 0x38, 0x80, 0x61, 0x6d, 0xdc, 0xb1, 0x44, 0x42, 0x6f, 0x00, 0x92, 0x34, 0x4e,
	0x58, 0x9a, 0x07, 0x2c, 0x1b, 0x5c, 0x63, 0xe9, 0x0c, 0x19, 0xfd, 0x5e, 0x03, 0x62, 0xb1, 0xd0,
	0xcd, 0x83, 0x78, 0xbf, 0x12, 0x70, 0xc1, 0x9f, 0xda, 0x05, 0x2c, 0xf4, 0xab, 0xa7, 0x30, 0xa1,
	0x43, 0xe8, 0xa6, 0x6c, 0xe7, 0x44, 0xb1, 0xcf, 0x42, 0x27, 0xf0, 0xab, 0xc7, 0x52, 0xb6, 0x5b,
	0x72, 0x48, 0xf7, 0xe9, 0x8f, 0xd0, 0x3b, 0x75, 0xec, 0xdd, 0x88, 0x0d, 0x6e, 0xb0, 0xa7, 0x5b,
	0xf5, 0x18, 0x6e, 0xc4, 0xe8, 0xb7, 0xd0, 0xe1, 0x5d, 0x62, 0xc2, 0x18, 0x1b, 0xa4, 0x94, 0xed,
	0xde, 0xf2, 0x9c, 0xf3, 0xe9, 0x63, 0xb4, 0x3a, 0x6c, 0xc3, 0xc0, 0x33, 0x13, 0x4e, 0x8c, 0x7e,
	0x03, 0x12, 0xf3, 0x83, 0xdc, 0xdd, 0x86, 0x0c, 0x39, 0x49, 0xd6, 0x31, 0xe7, 0x1b, 0x66, 0xcc,
	0x4d, 0xbd, 0x4f, 0x58, 0xbd, 0xc6, 0xea, 0x19, 0xc2, 0xbf, 0x4d, 0xd9, 0xaf, 0x87, 0x20, 0x65,
	0x3e, 0xd2, 0x91, 0xac, 0x63, 0x4e, 0x87, 0x20, 0xfb, 0x2c, 0xf3, 0xd2, 0x00, 0xc7, 0x94, 0x64,
	0xce, 0xa1, 0xd1, 0x1f, 0x35, 0x00, 0xe4, 0xb3, 0x48, 0xe3, 0x43, 0x42, 0x7b, 0x70, 0x11, 0x54,
	0xb2, 0x5c, 0x04, 0x3e, 0xa5, 0xd0, 0xc0, 0x3d, 0x85, 0x16, 0x18, 0x73, 0xf5, 0x82, 0xbd, 0xcf,
	0x7e, 0xc3, 0x69, 0x4d, 0x4b, 0x24, 0xf4, 0x3b, 0xbe, 0x75, 0x14, 0x7f, 0x46, 0x96, 0x63, 0xe4,
	0x71, 0x02, 0x38, 0x49, 0x2f, 0x0e, 0x43, 0x37, 0xc9, 0xd8, 0x60, 0x26, 0x48, 0x56, 0x39, 0x7d,
	0x05, 0x2d, 0xd4, 0x2a, 0x1b, 0xdc, 0x0d, 0xeb, 0x63, 0x79, 0x26, 0x4f, 0xf0, 0x06, 0x90, 0x95,
	0x55, 0x96, 0x46, 0x7f, 0x5d, 0x40, 0x13, 0x11, 0xfa, 0x35, 0x48, 0x47, 0x8b, 0x04, 0xd1, 0x76,
	0x54, 0xfa, 0xf3, 0x1a, 0xfa, 0xa5, 0xf5, 0x85, 0xf3, 0xc8, 0xf7, 0x39, 0xd9, 0x78, 0x55, 0x15,
	0x70, 0x4f, 0xdd, 0xa7, 0x3f, 0x80, 0x7c, 0xec, 0x0d, 0xfc, 0xd2, 0xc8, 0xea, 0x72, 0x0a, 0xdd,
	0xa7, 0xaf, 0xe0, 0xf2, 0xd8, 0x80, 0x1a, 0x08, 0xf5, 0xba, 0x15, 0x88, 0x5e, 0xdf, 0x9e, 0x35,
	0xe5, 0x45, 0x22, 0x96, 0xeb, 0xcd, 0x7a, 0x62, 0x05, 0xd5, 0xcd, 0x5d, 0xbb, 0x48, 0xd8, 0xe9,
	0x23, 0x9e, 0x71, 0x51, 0x0f, 0xfb, 0x20, 0x1f, 0xdc, 0x09, 0x51, 0x79, 0x2c, 0xe8, 0x04, 0x91,
	0x9b, 0x16, 0xce, 0x13, 0x2b, 0x06, 0x6f, 0x85, 0xcd, 0x25, 0xf4, 0x9e, 0x15, 0x74, 0x01, 0x2f,
	0x50, 0x0a, 0x27, 0xc1, 0xc3, 0x71, 0x62, 0x61, 0xe9, 0x6a, 0x58, 0x1b, 0xcb, 0xb3, 0x97, 0x67,
	0x92, 0x9d, 0x1f, 0x96, 0xd5, 0xdf, 0x3d, 0x87, 0x46, 0x7f, 0xd6, 0xa0, 0x69, 0xa3, 0x29, 0x04,
	0xea, 0x7c, 0x96, 0x10, 0x91, 0x87, 0xff, 0x69, 0xf7, 0xf7, 0x00, 0xa5, 0xde, 0x5e, 0xbc, 0x2f,
	0x75, 0xea, 0x08, 0xc5, 0xbd, 0x78, 0xff, 0xff, 0x27, 0xc6, 0x0d, 0x3b, 0x9a, 0x31, 0x13, 0x86,
	0x3d, 0x96, 0x26, 0xdc, 0x42, 0x57, 0x2c, 0x85, 0x40, 0x75, 0x00, 0xe4, 0x6c, 0x1b, 0xb4, 0xcb,
	0x92, 0x77, 0xc7, 0x38, 0x1b, 0xd9, 0x00, 0xc8, 0x5f, 0x5c, 0xec, 0x97, 0x2d, 0xf1, 0x8c, 0xe5,
	0xcd, 0xbf, 0x58, 0xbe, 0xfe, 0xbb, 0x06, 0x52, 0xe5, 0x17, 0x95, 0xa1, 0xbd, 0x31, 0xde, 0x1b,
	0xe6, 0x07, 0x83, 0x7c, 0x45, 0xdb, 0x50, 0xd7, 0x0d, 0x9b, 0xd4, 0xa8, 0x04, 0x8d, 0x07, 0xd3,
	0x58, 0x90, 0x0b, 0x0a, 0xd0, 0x32, 0x36, 0xcb, 0xb9, 0x66, 0x91, 0x3a, 0x47, 0x35, 0x63, 0xb3,
	0x24, 0x0d, 0xfe, 0xd5, 0xdc, 0x34, 0x1f, 0x34, 0xc5, 0x20, 0x4d, 0xde, 0xa2, 0x9a, 0x9b, 0xf9,
	0x83, 0x46, 0x5a, 0xb4, 0x03, 0xcd, 0x77, 0x6b, 0xd3, 0x98, 0x93, 0x36, 0x87, 0xd7, 0xb6, 0xa5,
	0x1b, 0x0b, 0x22, 0x71, 0x58, 0xb1, 0x2c, 0xe5, 0x17, 0xd2, 0xa1, 0x5d, 0x90, 0x56, 0xca, 0x7a,
	0xfd, 0xc1, 0xb4, 0x54, 0x02, 0xfc, 0xc9, 0xcd, 0x46, 0x57, 0x89, 0xcc, 0x23, 0x55, 0xb1, 0x35,
	0xd2, 0xa5, 0x97, 0xd0, 0xb1, 0xf5, 0xa5, 0xe6, 0x7c, 0x34, 0x0d, 0x8d, 0x5c, 0xf2, 0x82, 0x6e,
	0x68, 0x36, 0xe9, 0xf1, 0xe8, 0x5e, 0x57, 0x2d, 0x72, 0xc5, 0x1f, 0x59, 0x2a, 0xf7, 0x8e, 0xa2,
	0xaa, 0x16, 0x21, 0x7c, 0x92, 0x39, 0x7f, 0xa7, 0xdd, 0xdb, 0xa4, 0x4f, 0xaf, 0x40, 0xe6, 0xcc,
	0x9d, 0x72, 0x34, 0x9d, 0x2b, 0xf0, 0x72, 0xcf, 0xf2, 0x49, 0xf6, 0x54, 0x78, 0x61, 0x7c, 0xf0,
	0xc5, 0x1f, 0x2e, 0xca, 0xfe, 0xf1, 0xa7, 0xc7, 0x20, 0xff, 0x74, 0xd8, 0x4e, 0xbc, 0x38, 0x9a,
	0x66, 0x4f, 0xc5, 0x1b, 0x6c, 0x78, 0x93, 0x33, 0x6f, 0x8a, 0x4d, 0xd3, 0xcf, 0x3f, 0xe3, 0x7f,
	0xf4, 0xb6, 0x85, 0xe9, 0xed, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0xb1, 0x45, 0xe4, 0xbb,
	0x05, 0x00, 0x00,
}
