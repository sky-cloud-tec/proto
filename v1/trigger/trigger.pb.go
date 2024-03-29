// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: src/trigger/trigger.proto

package trigger

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerType int32

const (
	TriggerType_TIMER TriggerType = 0
	TriggerType_HTTP  TriggerType = 1
)

// Enum value maps for TriggerType.
var (
	TriggerType_name = map[int32]string{
		0: "TIMER",
		1: "HTTP",
	}
	TriggerType_value = map[string]int32{
		"TIMER": 0,
		"HTTP":  1,
	}
)

func (x TriggerType) Enum() *TriggerType {
	p := new(TriggerType)
	*p = x
	return p
}

func (x TriggerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TriggerType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_trigger_trigger_proto_enumTypes[0].Descriptor()
}

func (TriggerType) Type() protoreflect.EnumType {
	return &file_src_trigger_trigger_proto_enumTypes[0]
}

func (x TriggerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TriggerType.Descriptor instead.
func (TriggerType) EnumDescriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_RPC   ActionType = 0
	ActionType_EVENT ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "RPC",
		1: "EVENT",
	}
	ActionType_value = map[string]int32{
		"RPC":   0,
		"EVENT": 1,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_trigger_trigger_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_src_trigger_trigger_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{1}
}

type TimerType int32

const (
	TimerType_CRON     TimerType = 0
	TimerType_REPEATED TimerType = 1
	TimerType_ONCE     TimerType = 2
)

// Enum value maps for TimerType.
var (
	TimerType_name = map[int32]string{
		0: "CRON",
		1: "REPEATED",
		2: "ONCE",
	}
	TimerType_value = map[string]int32{
		"CRON":     0,
		"REPEATED": 1,
		"ONCE":     2,
	}
)

func (x TimerType) Enum() *TimerType {
	p := new(TimerType)
	*p = x
	return p
}

func (x TimerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimerType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_trigger_trigger_proto_enumTypes[2].Descriptor()
}

func (TimerType) Type() protoreflect.EnumType {
	return &file_src_trigger_trigger_proto_enumTypes[2]
}

func (x TimerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimerType.Descriptor instead.
func (TimerType) EnumDescriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{2}
}

type UNIT int32

const (
	UNIT_SECOND UNIT = 0
	UNIT_MINUTE UNIT = 1
	UNIT_HOUR   UNIT = 2
	UNIT_DAY    UNIT = 3
)

// Enum value maps for UNIT.
var (
	UNIT_name = map[int32]string{
		0: "SECOND",
		1: "MINUTE",
		2: "HOUR",
		3: "DAY",
	}
	UNIT_value = map[string]int32{
		"SECOND": 0,
		"MINUTE": 1,
		"HOUR":   2,
		"DAY":    3,
	}
)

func (x UNIT) Enum() *UNIT {
	p := new(UNIT)
	*p = x
	return p
}

func (x UNIT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UNIT) Descriptor() protoreflect.EnumDescriptor {
	return file_src_trigger_trigger_proto_enumTypes[3].Descriptor()
}

func (UNIT) Type() protoreflect.EnumType {
	return &file_src_trigger_trigger_proto_enumTypes[3]
}

func (x UNIT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UNIT.Descriptor instead.
func (UNIT) EnumDescriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{3}
}

type TypeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TimerType `protobuf:"varint,10,opt,name=type,proto3,enum=trigger.TimerType" json:"type,omitempty"`
	// 当type 为 CRON/ONCE 填充下面这个值
	ParamValue string `protobuf:"bytes,20,opt,name=paramValue,proto3" json:"paramValue,omitempty"`
	// 当type 为 REPEATED 填充下面两个值
	Unit  UNIT  `protobuf:"varint,30,opt,name=unit,proto3,enum=trigger.UNIT" json:"unit,omitempty"`
	Value int32 `protobuf:"varint,40,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TypeValue) Reset() {
	*x = TypeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeValue) ProtoMessage() {}

func (x *TypeValue) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeValue.ProtoReflect.Descriptor instead.
func (*TypeValue) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{0}
}

func (x *TypeValue) GetType() TimerType {
	if x != nil {
		return x.Type
	}
	return TimerType_CRON
}

func (x *TypeValue) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *TypeValue) GetUnit() UNIT {
	if x != nil {
		return x.Unit
	}
	return UNIT_SECOND
}

func (x *TypeValue) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TypeValueOfTimer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       TimerType `protobuf:"varint,10,opt,name=type,proto3,enum=trigger.TimerType" json:"type,omitempty"`
	ParamValue string    `protobuf:"bytes,20,opt,name=paramValue,proto3" json:"paramValue,omitempty"`
}

func (x *TypeValueOfTimer) Reset() {
	*x = TypeValueOfTimer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeValueOfTimer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeValueOfTimer) ProtoMessage() {}

func (x *TypeValueOfTimer) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeValueOfTimer.ProtoReflect.Descriptor instead.
func (*TypeValueOfTimer) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{1}
}

func (x *TypeValueOfTimer) GetType() TimerType {
	if x != nil {
		return x.Type
	}
	return TimerType_CRON
}

func (x *TypeValueOfTimer) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

type PaasResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,10,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceId   string `protobuf:"bytes,20,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *PaasResource) Reset() {
	*x = PaasResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaasResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaasResource) ProtoMessage() {}

func (x *PaasResource) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaasResource.ProtoReflect.Descriptor instead.
func (*PaasResource) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{2}
}

func (x *PaasResource) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *PaasResource) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type MapFieldEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,20,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MapFieldEntry) Reset() {
	*x = MapFieldEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapFieldEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapFieldEntry) ProtoMessage() {}

func (x *MapFieldEntry) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapFieldEntry.ProtoReflect.Descriptor instead.
func (*MapFieldEntry) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{3}
}

func (x *MapFieldEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapFieldEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ActionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string           `protobuf:"bytes,10,opt,name=service,proto3" json:"service,omitempty"`
	Params  []*MapFieldEntry `protobuf:"bytes,20,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *ActionParams) Reset() {
	*x = ActionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionParams) ProtoMessage() {}

func (x *ActionParams) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionParams.ProtoReflect.Descriptor instead.
func (*ActionParams) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{4}
}

func (x *ActionParams) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ActionParams) GetParams() []*MapFieldEntry {
	if x != nil {
		return x.Params
	}
	return nil
}

type ActionParamsOfRpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string                `protobuf:"bytes,10,opt,name=service,proto3" json:"service,omitempty"`
	Param   map[string]*anypb.Any `protobuf:"bytes,20,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ActionParamsOfRpc) Reset() {
	*x = ActionParamsOfRpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionParamsOfRpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionParamsOfRpc) ProtoMessage() {}

func (x *ActionParamsOfRpc) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionParamsOfRpc.ProtoReflect.Descriptor instead.
func (*ActionParamsOfRpc) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{5}
}

func (x *ActionParamsOfRpc) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ActionParamsOfRpc) GetParam() map[string]*anypb.Any {
	if x != nil {
		return x.Param
	}
	return nil
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64           `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	TriggerName  string          `protobuf:"bytes,20,opt,name=triggerName,proto3" json:"triggerName,omitempty"`
	Disabled     bool            `protobuf:"varint,30,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Description  string          `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
	Bindings     []*PaasResource `protobuf:"bytes,50,rep,name=bindings,proto3" json:"bindings,omitempty"`
	CreatedTime  string          `protobuf:"bytes,60,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime  string          `protobuf:"bytes,70,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	TriggerType  TriggerType     `protobuf:"varint,80,opt,name=triggerType,proto3,enum=trigger.TriggerType" json:"triggerType,omitempty"`
	TypeValue    *TypeValue      `protobuf:"bytes,90,opt,name=typeValue,proto3" json:"typeValue,omitempty"`
	ActionType   ActionType      `protobuf:"varint,100,opt,name=actionType,proto3,enum=trigger.ActionType" json:"actionType,omitempty"`
	ActionParams *ActionParams   `protobuf:"bytes,110,opt,name=actionParams,proto3" json:"actionParams,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_trigger_trigger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_src_trigger_trigger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_src_trigger_trigger_proto_rawDescGZIP(), []int{6}
}

func (x *Trigger) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Trigger) GetTriggerName() string {
	if x != nil {
		return x.TriggerName
	}
	return ""
}

func (x *Trigger) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Trigger) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Trigger) GetBindings() []*PaasResource {
	if x != nil {
		return x.Bindings
	}
	return nil
}

func (x *Trigger) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *Trigger) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *Trigger) GetTriggerType() TriggerType {
	if x != nil {
		return x.TriggerType
	}
	return TriggerType_TIMER
}

func (x *Trigger) GetTypeValue() *TypeValue {
	if x != nil {
		return x.TypeValue
	}
	return nil
}

func (x *Trigger) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_RPC
}

func (x *Trigger) GetActionParams() *ActionParams {
	if x != nil {
		return x.ActionParams
	}
	return nil
}

var File_src_trigger_trigger_proto protoreflect.FileDescriptor

var file_src_trigger_trigger_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x01, 0x0a, 0x09, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x4e,
	0x49, 0x54, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a,
	0x0a, 0x10, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x66, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52, 0x0a, 0x0c, 0x50, 0x61,
	0x61, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x37,
	0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x4f, 0x66, 0x52, 0x70, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x66, 0x52, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x4e,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xca,
	0x03, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x61, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x74, 0x79, 0x70,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0c, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a, 0x22, 0x0a, 0x0b, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x49,
	0x4d, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x2a,
	0x20, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x2a, 0x2d, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x43, 0x52, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x50, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4e, 0x43, 0x45, 0x10, 0x02,
	0x2a, 0x31, 0x0a, 0x04, 0x55, 0x4e, 0x49, 0x54, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41,
	0x59, 0x10, 0x03, 0x42, 0x47, 0x0a, 0x1a, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b,
	0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_trigger_trigger_proto_rawDescOnce sync.Once
	file_src_trigger_trigger_proto_rawDescData = file_src_trigger_trigger_proto_rawDesc
)

func file_src_trigger_trigger_proto_rawDescGZIP() []byte {
	file_src_trigger_trigger_proto_rawDescOnce.Do(func() {
		file_src_trigger_trigger_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_trigger_trigger_proto_rawDescData)
	})
	return file_src_trigger_trigger_proto_rawDescData
}

var file_src_trigger_trigger_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_src_trigger_trigger_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_src_trigger_trigger_proto_goTypes = []interface{}{
	(TriggerType)(0),          // 0: trigger.TriggerType
	(ActionType)(0),           // 1: trigger.ActionType
	(TimerType)(0),            // 2: trigger.TimerType
	(UNIT)(0),                 // 3: trigger.UNIT
	(*TypeValue)(nil),         // 4: trigger.TypeValue
	(*TypeValueOfTimer)(nil),  // 5: trigger.TypeValueOfTimer
	(*PaasResource)(nil),      // 6: trigger.PaasResource
	(*MapFieldEntry)(nil),     // 7: trigger.MapFieldEntry
	(*ActionParams)(nil),      // 8: trigger.ActionParams
	(*ActionParamsOfRpc)(nil), // 9: trigger.ActionParamsOfRpc
	(*Trigger)(nil),           // 10: trigger.Trigger
	nil,                       // 11: trigger.ActionParamsOfRpc.ParamEntry
	(*anypb.Any)(nil),         // 12: google.protobuf.Any
}
var file_src_trigger_trigger_proto_depIdxs = []int32{
	2,  // 0: trigger.TypeValue.type:type_name -> trigger.TimerType
	3,  // 1: trigger.TypeValue.unit:type_name -> trigger.UNIT
	2,  // 2: trigger.TypeValueOfTimer.type:type_name -> trigger.TimerType
	7,  // 3: trigger.ActionParams.params:type_name -> trigger.MapFieldEntry
	11, // 4: trigger.ActionParamsOfRpc.param:type_name -> trigger.ActionParamsOfRpc.ParamEntry
	6,  // 5: trigger.Trigger.bindings:type_name -> trigger.PaasResource
	0,  // 6: trigger.Trigger.triggerType:type_name -> trigger.TriggerType
	4,  // 7: trigger.Trigger.typeValue:type_name -> trigger.TypeValue
	1,  // 8: trigger.Trigger.actionType:type_name -> trigger.ActionType
	8,  // 9: trigger.Trigger.actionParams:type_name -> trigger.ActionParams
	12, // 10: trigger.ActionParamsOfRpc.ParamEntry.value:type_name -> google.protobuf.Any
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_src_trigger_trigger_proto_init() }
func file_src_trigger_trigger_proto_init() {
	if File_src_trigger_trigger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_trigger_trigger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeValue); i {
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
		file_src_trigger_trigger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeValueOfTimer); i {
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
		file_src_trigger_trigger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaasResource); i {
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
		file_src_trigger_trigger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapFieldEntry); i {
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
		file_src_trigger_trigger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionParams); i {
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
		file_src_trigger_trigger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionParamsOfRpc); i {
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
		file_src_trigger_trigger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
			RawDescriptor: file_src_trigger_trigger_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_trigger_trigger_proto_goTypes,
		DependencyIndexes: file_src_trigger_trigger_proto_depIdxs,
		EnumInfos:         file_src_trigger_trigger_proto_enumTypes,
		MessageInfos:      file_src_trigger_trigger_proto_msgTypes,
	}.Build()
	File_src_trigger_trigger_proto = out.File
	file_src_trigger_trigger_proto_rawDesc = nil
	file_src_trigger_trigger_proto_goTypes = nil
	file_src_trigger_trigger_proto_depIdxs = nil
}
