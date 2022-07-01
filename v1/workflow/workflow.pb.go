// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: src/workflow/workflow.proto

package workflow

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

type WorkflowState int32

const (
	WorkflowState_UNKNOWN_WORKFLOW_STATE WorkflowState = 0
	WorkflowState_RUNNING                WorkflowState = 1 //运行中
	WorkflowState_SUCCEEDED              WorkflowState = 2 //成功
	WorkflowState_FAILED                 WorkflowState = 3 //失败
)

// Enum value maps for WorkflowState.
var (
	WorkflowState_name = map[int32]string{
		0: "UNKNOWN_WORKFLOW_STATE",
		1: "RUNNING",
		2: "SUCCEEDED",
		3: "FAILED",
	}
	WorkflowState_value = map[string]int32{
		"UNKNOWN_WORKFLOW_STATE": 0,
		"RUNNING":                1,
		"SUCCEEDED":              2,
		"FAILED":                 3,
	}
)

func (x WorkflowState) Enum() *WorkflowState {
	p := new(WorkflowState)
	*p = x
	return p
}

func (x WorkflowState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowState) Descriptor() protoreflect.EnumDescriptor {
	return file_src_workflow_workflow_proto_enumTypes[0].Descriptor()
}

func (WorkflowState) Type() protoreflect.EnumType {
	return &file_src_workflow_workflow_proto_enumTypes[0]
}

func (x WorkflowState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowState.Descriptor instead.
func (WorkflowState) EnumDescriptor() ([]byte, []int) {
	return file_src_workflow_workflow_proto_rawDescGZIP(), []int{0}
}

var File_src_workflow_workflow_proto protoreflect.FileDescriptor

var file_src_workflow_workflow_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x72, 0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2a, 0x53, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x49, 0x0a, 0x1b,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_workflow_workflow_proto_rawDescOnce sync.Once
	file_src_workflow_workflow_proto_rawDescData = file_src_workflow_workflow_proto_rawDesc
)

func file_src_workflow_workflow_proto_rawDescGZIP() []byte {
	file_src_workflow_workflow_proto_rawDescOnce.Do(func() {
		file_src_workflow_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_workflow_workflow_proto_rawDescData)
	})
	return file_src_workflow_workflow_proto_rawDescData
}

var file_src_workflow_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_workflow_workflow_proto_goTypes = []interface{}{
	(WorkflowState)(0), // 0: workflow.WorkflowState
}
var file_src_workflow_workflow_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_workflow_workflow_proto_init() }
func file_src_workflow_workflow_proto_init() {
	if File_src_workflow_workflow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_workflow_workflow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_workflow_workflow_proto_goTypes,
		DependencyIndexes: file_src_workflow_workflow_proto_depIdxs,
		EnumInfos:         file_src_workflow_workflow_proto_enumTypes,
	}.Build()
	File_src_workflow_workflow_proto = out.File
	file_src_workflow_workflow_proto_rawDesc = nil
	file_src_workflow_workflow_proto_goTypes = nil
	file_src_workflow_workflow_proto_depIdxs = nil
}
