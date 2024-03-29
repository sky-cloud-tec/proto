// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: src/grpc/pipeline.proto

package rpc

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

type ScheduleType int32

const (
	ScheduleType_IMMEDIATELY                ScheduleType = 0
	ScheduleType_TIMER                      ScheduleType = 1
	ScheduleType_NOT_AUTOMATICALLY_DISPATCH ScheduleType = 2
)

// Enum value maps for ScheduleType.
var (
	ScheduleType_name = map[int32]string{
		0: "IMMEDIATELY",
		1: "TIMER",
		2: "NOT_AUTOMATICALLY_DISPATCH",
	}
	ScheduleType_value = map[string]int32{
		"IMMEDIATELY":                0,
		"TIMER":                      1,
		"NOT_AUTOMATICALLY_DISPATCH": 2,
	}
)

func (x ScheduleType) Enum() *ScheduleType {
	p := new(ScheduleType)
	*p = x
	return p
}

func (x ScheduleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduleType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_grpc_pipeline_proto_enumTypes[0].Descriptor()
}

func (ScheduleType) Type() protoreflect.EnumType {
	return &file_src_grpc_pipeline_proto_enumTypes[0]
}

func (x ScheduleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduleType.Descriptor instead.
func (ScheduleType) EnumDescriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_API ActionType = 0
	ActionType_CLI ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "API",
		1: "CLI",
	}
	ActionType_value = map[string]int32{
		"API": 0,
		"CLI": 1,
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
	return file_src_grpc_pipeline_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_src_grpc_pipeline_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{1}
}

type PipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description     string            `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	ContextId       string            `protobuf:"bytes,15,opt,name=contextId,proto3" json:"contextId,omitempty"`
	ExecMode        string            `protobuf:"bytes,20,opt,name=execMode,proto3" json:"execMode,omitempty"`
	FailMode        string            `protobuf:"bytes,25,opt,name=failMode,proto3" json:"failMode,omitempty"`
	Schedule        *Schedule         `protobuf:"bytes,30,opt,name=schedule,proto3" json:"schedule,omitempty"`
	DevicePipelines []*DevicePipeline `protobuf:"bytes,35,rep,name=devicePipelines,proto3" json:"devicePipelines,omitempty"`
	Id              string            `protobuf:"bytes,40,opt,name=id,proto3" json:"id,omitempty"`
	Type            string            `protobuf:"bytes,45,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PipelineRequest) Reset() {
	*x = PipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineRequest) ProtoMessage() {}

func (x *PipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineRequest.ProtoReflect.Descriptor instead.
func (*PipelineRequest) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PipelineRequest) GetContextId() string {
	if x != nil {
		return x.ContextId
	}
	return ""
}

func (x *PipelineRequest) GetExecMode() string {
	if x != nil {
		return x.ExecMode
	}
	return ""
}

func (x *PipelineRequest) GetFailMode() string {
	if x != nil {
		return x.FailMode
	}
	return ""
}

func (x *PipelineRequest) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *PipelineRequest) GetDevicePipelines() []*DevicePipeline {
	if x != nil {
		return x.DevicePipelines
	}
	return nil
}

func (x *PipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PipelineRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PipelineIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId string `protobuf:"bytes,5,opt,name=pipelineId,proto3" json:"pipelineId,omitempty"`
}

func (x *PipelineIdRequest) Reset() {
	*x = PipelineIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineIdRequest) ProtoMessage() {}

func (x *PipelineIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineIdRequest.ProtoReflect.Descriptor instead.
func (*PipelineIdRequest) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineIdRequest) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     ScheduleType `protobuf:"varint,1,opt,name=type,proto3,enum=rpc.ScheduleType" json:"type,omitempty"`
	PlanTime int64        `protobuf:"varint,2,opt,name=planTime,proto3" json:"planTime,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *Schedule) GetType() ScheduleType {
	if x != nil {
		return x.Type
	}
	return ScheduleType_IMMEDIATELY
}

func (x *Schedule) GetPlanTime() int64 {
	if x != nil {
		return x.PlanTime
	}
	return 0
}

type DevicePipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId   string  `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	DeviceName string  `protobuf:"bytes,5,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceType string  `protobuf:"bytes,10,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	Steps      []*Step `protobuf:"bytes,15,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *DevicePipeline) Reset() {
	*x = DevicePipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicePipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicePipeline) ProtoMessage() {}

func (x *DevicePipeline) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicePipeline.ProtoReflect.Descriptor instead.
func (*DevicePipeline) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *DevicePipeline) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DevicePipeline) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DevicePipeline) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *DevicePipeline) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message         string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	RollbackCommand *Action `protobuf:"bytes,5,opt,name=rollbackCommand,proto3" json:"rollbackCommand,omitempty"`
	DoCommand       *Action `protobuf:"bytes,10,opt,name=doCommand,proto3" json:"doCommand,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *Step) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Step) GetRollbackCommand() *Action {
	if x != nil {
		return x.RollbackCommand
	}
	return nil
}

func (x *Step) GetDoCommand() *Action {
	if x != nil {
		return x.DoCommand
	}
	return nil
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ActionType   `protobuf:"varint,1,opt,name=type,proto3,enum=rpc.ActionType" json:"type,omitempty"`
	NetdAction  *NetDAction  `protobuf:"bytes,2,opt,name=netdAction,proto3" json:"netdAction,omitempty"`
	HttpdAction *HttpDAction `protobuf:"bytes,3,opt,name=httpdAction,proto3" json:"httpdAction,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{5}
}

func (x *Action) GetType() ActionType {
	if x != nil {
		return x.Type
	}
	return ActionType_API
}

func (x *Action) GetNetdAction() *NetDAction {
	if x != nil {
		return x.NetdAction
	}
	return nil
}

func (x *Action) GetHttpdAction() *HttpDAction {
	if x != nil {
		return x.HttpdAction
	}
	return nil
}

type NetDAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Mode    string `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *NetDAction) Reset() {
	*x = NetDAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetDAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetDAction) ProtoMessage() {}

func (x *NetDAction) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetDAction.ProtoReflect.Descriptor instead.
func (*NetDAction) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{6}
}

func (x *NetDAction) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NetDAction) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type HttpDAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *IRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *HttpDAction) Reset() {
	*x = HttpDAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpDAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpDAction) ProtoMessage() {}

func (x *HttpDAction) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpDAction.ProtoReflect.Descriptor instead.
func (*HttpDAction) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{7}
}

func (x *HttpDAction) GetRequest() *IRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type PipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PipelineResponse) Reset() {
	*x = PipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineResponse) ProtoMessage() {}

func (x *PipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineResponse.ProtoReflect.Descriptor instead.
func (*PipelineResponse) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{8}
}

func (x *PipelineResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PipelineResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PipelineGetTriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PipelineGetTriggerResponse) Reset() {
	*x = PipelineGetTriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_pipeline_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineGetTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineGetTriggerResponse) ProtoMessage() {}

func (x *PipelineGetTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_pipeline_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineGetTriggerResponse.ProtoReflect.Descriptor instead.
func (*PipelineGetTriggerResponse) Descriptor() ([]byte, []int) {
	return file_src_grpc_pipeline_proto_rawDescGZIP(), []int{9}
}

func (x *PipelineGetTriggerResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_src_grpc_pipeline_proto protoreflect.FileDescriptor

var file_src_grpc_pipeline_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x17,
	0x73, 0x72, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x08, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70,
	0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x0f,
	0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x92,
	0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x74, 0x44, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x44,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x44, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x36, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x44, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x2c, 0x0a, 0x1a, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x2a,
	0x4a, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x4c, 0x59, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59,
	0x5f, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x2a, 0x1e, 0x0a, 0x0a, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x49,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x4c, 0x49, 0x10, 0x01, 0x32, 0xe7, 0x02, 0x0a, 0x0f,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x0a, 0x16, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_grpc_pipeline_proto_rawDescOnce sync.Once
	file_src_grpc_pipeline_proto_rawDescData = file_src_grpc_pipeline_proto_rawDesc
)

func file_src_grpc_pipeline_proto_rawDescGZIP() []byte {
	file_src_grpc_pipeline_proto_rawDescOnce.Do(func() {
		file_src_grpc_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_grpc_pipeline_proto_rawDescData)
	})
	return file_src_grpc_pipeline_proto_rawDescData
}

var file_src_grpc_pipeline_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_src_grpc_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_src_grpc_pipeline_proto_goTypes = []interface{}{
	(ScheduleType)(0),                  // 0: rpc.ScheduleType
	(ActionType)(0),                    // 1: rpc.ActionType
	(*PipelineRequest)(nil),            // 2: rpc.PipelineRequest
	(*PipelineIdRequest)(nil),          // 3: rpc.PipelineIdRequest
	(*Schedule)(nil),                   // 4: rpc.Schedule
	(*DevicePipeline)(nil),             // 5: rpc.DevicePipeline
	(*Step)(nil),                       // 6: rpc.Step
	(*Action)(nil),                     // 7: rpc.Action
	(*NetDAction)(nil),                 // 8: rpc.NetDAction
	(*HttpDAction)(nil),                // 9: rpc.HttpDAction
	(*PipelineResponse)(nil),           // 10: rpc.PipelineResponse
	(*PipelineGetTriggerResponse)(nil), // 11: rpc.PipelineGetTriggerResponse
	(*IRequest)(nil),                   // 12: rpc.IRequest
}
var file_src_grpc_pipeline_proto_depIdxs = []int32{
	4,  // 0: rpc.PipelineRequest.schedule:type_name -> rpc.Schedule
	5,  // 1: rpc.PipelineRequest.devicePipelines:type_name -> rpc.DevicePipeline
	0,  // 2: rpc.Schedule.type:type_name -> rpc.ScheduleType
	6,  // 3: rpc.DevicePipeline.steps:type_name -> rpc.Step
	7,  // 4: rpc.Step.rollbackCommand:type_name -> rpc.Action
	7,  // 5: rpc.Step.doCommand:type_name -> rpc.Action
	1,  // 6: rpc.Action.type:type_name -> rpc.ActionType
	8,  // 7: rpc.Action.netdAction:type_name -> rpc.NetDAction
	9,  // 8: rpc.Action.httpdAction:type_name -> rpc.HttpDAction
	12, // 9: rpc.HttpDAction.request:type_name -> rpc.IRequest
	2,  // 10: rpc.PipelineService.createPipeline:input_type -> rpc.PipelineRequest
	3,  // 11: rpc.PipelineService.ExecutePipeline:input_type -> rpc.PipelineIdRequest
	3,  // 12: rpc.PipelineService.UpdatePipeline:input_type -> rpc.PipelineIdRequest
	3,  // 13: rpc.PipelineService.DeletePipeline:input_type -> rpc.PipelineIdRequest
	3,  // 14: rpc.PipelineService.GetTriggerId:input_type -> rpc.PipelineIdRequest
	10, // 15: rpc.PipelineService.createPipeline:output_type -> rpc.PipelineResponse
	10, // 16: rpc.PipelineService.ExecutePipeline:output_type -> rpc.PipelineResponse
	10, // 17: rpc.PipelineService.UpdatePipeline:output_type -> rpc.PipelineResponse
	10, // 18: rpc.PipelineService.DeletePipeline:output_type -> rpc.PipelineResponse
	11, // 19: rpc.PipelineService.GetTriggerId:output_type -> rpc.PipelineGetTriggerResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_src_grpc_pipeline_proto_init() }
func file_src_grpc_pipeline_proto_init() {
	if File_src_grpc_pipeline_proto != nil {
		return
	}
	file_src_grpc_irequest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_grpc_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineRequest); i {
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
		file_src_grpc_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineIdRequest); i {
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
		file_src_grpc_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_src_grpc_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicePipeline); i {
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
		file_src_grpc_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_src_grpc_pipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_src_grpc_pipeline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetDAction); i {
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
		file_src_grpc_pipeline_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpDAction); i {
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
		file_src_grpc_pipeline_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineResponse); i {
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
		file_src_grpc_pipeline_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineGetTriggerResponse); i {
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
			RawDescriptor: file_src_grpc_pipeline_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_grpc_pipeline_proto_goTypes,
		DependencyIndexes: file_src_grpc_pipeline_proto_depIdxs,
		EnumInfos:         file_src_grpc_pipeline_proto_enumTypes,
		MessageInfos:      file_src_grpc_pipeline_proto_msgTypes,
	}.Build()
	File_src_grpc_pipeline_proto = out.File
	file_src_grpc_pipeline_proto_rawDesc = nil
	file_src_grpc_pipeline_proto_goTypes = nil
	file_src_grpc_pipeline_proto_depIdxs = nil
}
