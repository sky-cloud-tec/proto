// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: src/nap/nap.proto

package nap

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

type PolicyTicketStage int32

const (
	PolicyTicketStage_UNKNOWN_POLICY_TICKET_STAGE PolicyTicketStage = 0
	PolicyTicketStage_CREATED                     PolicyTicketStage = 1  // 已创建
	PolicyTicketStage_ROUTE_ANALYZED              PolicyTicketStage = 2  // 路径分析完成
	PolicyTicketStage_OBJECT_ANALYZED             PolicyTicketStage = 3  // 对象分析完成
	PolicyTicketStage_AUDITED                     PolicyTicketStage = 4  // 被批准
	PolicyTicketStage_DECLINED                    PolicyTicketStage = 5  // 被拒绝
	PolicyTicketStage_DISPATCH_PENDING            PolicyTicketStage = 6  // 等待下发(包括未设置下发时间和已经设置下发时间但是未开始)
	PolicyTicketStage_DISPATCHING                 PolicyTicketStage = 7  // 正在下发
	PolicyTicketStage_DISPATCHED                  PolicyTicketStage = 8  // 已下发完成
	PolicyTicketStage_ROLLBACK_PENDING            PolicyTicketStage = 9  // 等待回滚(已经设置回滚时间)
	PolicyTicketStage_ROLLING_BACK                PolicyTicketStage = 10 // 正在回滚
	PolicyTicketStage_ROLLED_BACK                 PolicyTicketStage = 11 // 回滚完成
	PolicyTicketStage_ROUTE_ANALYZED_ERROR        PolicyTicketStage = 12 // 路径分析出错
	PolicyTicketStage_OBJECT_ANALYZED_ERROR       PolicyTicketStage = 13 // 对象分析出错
	PolicyTicketStage_DISPATCHING_ERROR           PolicyTicketStage = 14 // 下发出错
	PolicyTicketStage_ROLLING_BACK_ERROR          PolicyTicketStage = 15 // 回滚出错
	PolicyTicketStage_ERROR                       PolicyTicketStage = 16 // 出错
)

// Enum value maps for PolicyTicketStage.
var (
	PolicyTicketStage_name = map[int32]string{
		0:  "UNKNOWN_POLICY_TICKET_STAGE",
		1:  "CREATED",
		2:  "ROUTE_ANALYZED",
		3:  "OBJECT_ANALYZED",
		4:  "AUDITED",
		5:  "DECLINED",
		6:  "DISPATCH_PENDING",
		7:  "DISPATCHING",
		8:  "DISPATCHED",
		9:  "ROLLBACK_PENDING",
		10: "ROLLING_BACK",
		11: "ROLLED_BACK",
		12: "ROUTE_ANALYZED_ERROR",
		13: "OBJECT_ANALYZED_ERROR",
		14: "DISPATCHING_ERROR",
		15: "ROLLING_BACK_ERROR",
		16: "ERROR",
	}
	PolicyTicketStage_value = map[string]int32{
		"UNKNOWN_POLICY_TICKET_STAGE": 0,
		"CREATED":                     1,
		"ROUTE_ANALYZED":              2,
		"OBJECT_ANALYZED":             3,
		"AUDITED":                     4,
		"DECLINED":                    5,
		"DISPATCH_PENDING":            6,
		"DISPATCHING":                 7,
		"DISPATCHED":                  8,
		"ROLLBACK_PENDING":            9,
		"ROLLING_BACK":                10,
		"ROLLED_BACK":                 11,
		"ROUTE_ANALYZED_ERROR":        12,
		"OBJECT_ANALYZED_ERROR":       13,
		"DISPATCHING_ERROR":           14,
		"ROLLING_BACK_ERROR":          15,
		"ERROR":                       16,
	}
)

func (x PolicyTicketStage) Enum() *PolicyTicketStage {
	p := new(PolicyTicketStage)
	*p = x
	return p
}

func (x PolicyTicketStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyTicketStage) Descriptor() protoreflect.EnumDescriptor {
	return file_src_nap_nap_proto_enumTypes[0].Descriptor()
}

func (PolicyTicketStage) Type() protoreflect.EnumType {
	return &file_src_nap_nap_proto_enumTypes[0]
}

func (x PolicyTicketStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyTicketStage.Descriptor instead.
func (PolicyTicketStage) EnumDescriptor() ([]byte, []int) {
	return file_src_nap_nap_proto_rawDescGZIP(), []int{0}
}

type CapType int32

const (
	CapType_UNKNOWN_CAP_TYPE         CapType = 0
	CapType_POLICY_SEARCH            CapType = 1
	CapType_POLICY_TICKET            CapType = 2
	CapType_NSXT_POLICY_TICKET       CapType = 3 // NSX-T防火墙策略开通
	CapType_OBJECTS_MANAGER          CapType = 4
	CapType_NAT_MANAGER              CapType = 5
	CapType_DEVICE_DISCOVERER        CapType = 6
	CapType_NAT_AND_POLICY           CapType = 7
	CapType_POLICY_TRACE             CapType = 8
	CapType_SWITCH_INTERFACE_CONFIG  CapType = 9  //接口配置
	CapType_POLICY_CHANGE_REPORT     CapType = 10 //策略变更查询
	CapType_POLICY_COMPLIANCE_RULE   CapType = 11 //策略合规规则
	CapType_DEVICE_QUOTA_MAX_DEVICES CapType = 12
)

// Enum value maps for CapType.
var (
	CapType_name = map[int32]string{
		0:  "UNKNOWN_CAP_TYPE",
		1:  "POLICY_SEARCH",
		2:  "POLICY_TICKET",
		3:  "NSXT_POLICY_TICKET",
		4:  "OBJECTS_MANAGER",
		5:  "NAT_MANAGER",
		6:  "DEVICE_DISCOVERER",
		7:  "NAT_AND_POLICY",
		8:  "POLICY_TRACE",
		9:  "SWITCH_INTERFACE_CONFIG",
		10: "POLICY_CHANGE_REPORT",
		11: "POLICY_COMPLIANCE_RULE",
		12: "DEVICE_QUOTA_MAX_DEVICES",
	}
	CapType_value = map[string]int32{
		"UNKNOWN_CAP_TYPE":         0,
		"POLICY_SEARCH":            1,
		"POLICY_TICKET":            2,
		"NSXT_POLICY_TICKET":       3,
		"OBJECTS_MANAGER":          4,
		"NAT_MANAGER":              5,
		"DEVICE_DISCOVERER":        6,
		"NAT_AND_POLICY":           7,
		"POLICY_TRACE":             8,
		"SWITCH_INTERFACE_CONFIG":  9,
		"POLICY_CHANGE_REPORT":     10,
		"POLICY_COMPLIANCE_RULE":   11,
		"DEVICE_QUOTA_MAX_DEVICES": 12,
	}
)

func (x CapType) Enum() *CapType {
	p := new(CapType)
	*p = x
	return p
}

func (x CapType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_nap_nap_proto_enumTypes[1].Descriptor()
}

func (CapType) Type() protoreflect.EnumType {
	return &file_src_nap_nap_proto_enumTypes[1]
}

func (x CapType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapType.Descriptor instead.
func (CapType) EnumDescriptor() ([]byte, []int) {
	return file_src_nap_nap_proto_rawDescGZIP(), []int{1}
}

type DeviceState int32

const (
	DeviceState_UNKNOWN DeviceState = 0
	DeviceState_RUNNING DeviceState = 1
	DeviceState_DOWN    DeviceState = 2
)

// Enum value maps for DeviceState.
var (
	DeviceState_name = map[int32]string{
		0: "UNKNOWN",
		1: "RUNNING",
		2: "DOWN",
	}
	DeviceState_value = map[string]int32{
		"UNKNOWN": 0,
		"RUNNING": 1,
		"DOWN":    2,
	}
)

func (x DeviceState) Enum() *DeviceState {
	p := new(DeviceState)
	*p = x
	return p
}

func (x DeviceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceState) Descriptor() protoreflect.EnumDescriptor {
	return file_src_nap_nap_proto_enumTypes[2].Descriptor()
}

func (DeviceState) Type() protoreflect.EnumType {
	return &file_src_nap_nap_proto_enumTypes[2]
}

func (x DeviceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceState.Descriptor instead.
func (DeviceState) EnumDescriptor() ([]byte, []int) {
	return file_src_nap_nap_proto_rawDescGZIP(), []int{2}
}

type ResourceObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ResourceObject) Reset() {
	*x = ResourceObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_nap_nap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceObject) ProtoMessage() {}

func (x *ResourceObject) ProtoReflect() protoreflect.Message {
	mi := &file_src_nap_nap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceObject.ProtoReflect.Descriptor instead.
func (*ResourceObject) Descriptor() ([]byte, []int) {
	return file_src_nap_nap_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceObject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_src_nap_nap_proto protoreflect.FileDescriptor

var file_src_nap_nap_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x72, 0x63, 0x2f, 0x6e, 0x61, 0x70, 0x2f, 0x6e, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x61, 0x70, 0x22, 0x48, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x2a, 0xe4, 0x02, 0x0a, 0x11, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x55, 0x44, 0x49, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x49,
	0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x10,
	0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10,
	0x08, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4c, 0x4c, 0x49,
	0x4e, 0x47, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f,
	0x55, 0x54, 0x45, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41,
	0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0d, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x4c, 0x4c, 0x49, 0x4e,
	0x47, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0f, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x10, 0x2a, 0xb1, 0x02, 0x0a, 0x07, 0x43, 0x61,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x43, 0x41, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x53, 0x58, 0x54, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x41, 0x54, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x05, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x45, 0x52, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x41, 0x54, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f,
	0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17,
	0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x46, 0x41, 0x43, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x10, 0x0b, 0x12,
	0x1c, 0x0a, 0x18, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x5f,
	0x4d, 0x41, 0x58, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x53, 0x10, 0x0c, 0x2a, 0x31, 0x0a,
	0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02,
	0x42, 0x3f, 0x0a, 0x16, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x61, 0x70, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_nap_nap_proto_rawDescOnce sync.Once
	file_src_nap_nap_proto_rawDescData = file_src_nap_nap_proto_rawDesc
)

func file_src_nap_nap_proto_rawDescGZIP() []byte {
	file_src_nap_nap_proto_rawDescOnce.Do(func() {
		file_src_nap_nap_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_nap_nap_proto_rawDescData)
	})
	return file_src_nap_nap_proto_rawDescData
}

var file_src_nap_nap_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_src_nap_nap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_nap_nap_proto_goTypes = []interface{}{
	(PolicyTicketStage)(0), // 0: nap.PolicyTicketStage
	(CapType)(0),           // 1: nap.CapType
	(DeviceState)(0),       // 2: nap.DeviceState
	(*ResourceObject)(nil), // 3: nap.ResourceObject
}
var file_src_nap_nap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_nap_nap_proto_init() }
func file_src_nap_nap_proto_init() {
	if File_src_nap_nap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_nap_nap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceObject); i {
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
			RawDescriptor: file_src_nap_nap_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_nap_nap_proto_goTypes,
		DependencyIndexes: file_src_nap_nap_proto_depIdxs,
		EnumInfos:         file_src_nap_nap_proto_enumTypes,
		MessageInfos:      file_src_nap_nap_proto_msgTypes,
	}.Build()
	File_src_nap_nap_proto = out.File
	file_src_nap_nap_proto_rawDesc = nil
	file_src_nap_nap_proto_goTypes = nil
	file_src_nap_nap_proto_depIdxs = nil
}
