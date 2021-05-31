// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: src/netai/netai.proto

package netai

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

type TrafficLoad int32

const (
	TrafficLoad_HIGH   TrafficLoad = 0 // 比预期偏高
	TrafficLoad_NORMAL TrafficLoad = 1 // 符合预期
	TrafficLoad_LOW    TrafficLoad = 2 // 比预期偏低
)

// Enum value maps for TrafficLoad.
var (
	TrafficLoad_name = map[int32]string{
		0: "HIGH",
		1: "NORMAL",
		2: "LOW",
	}
	TrafficLoad_value = map[string]int32{
		"HIGH":   0,
		"NORMAL": 1,
		"LOW":    2,
	}
)

func (x TrafficLoad) Enum() *TrafficLoad {
	p := new(TrafficLoad)
	*p = x
	return p
}

func (x TrafficLoad) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrafficLoad) Descriptor() protoreflect.EnumDescriptor {
	return file_src_netai_netai_proto_enumTypes[0].Descriptor()
}

func (TrafficLoad) Type() protoreflect.EnumType {
	return &file_src_netai_netai_proto_enumTypes[0]
}

func (x TrafficLoad) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrafficLoad.Descriptor instead.
func (TrafficLoad) EnumDescriptor() ([]byte, []int) {
	return file_src_netai_netai_proto_rawDescGZIP(), []int{0}
}

var File_src_netai_netai_proto protoreflect.FileDescriptor

var file_src_netai_netai_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x72, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x61,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x2a, 0x2c,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x42, 0x43, 0x0a, 0x18,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x74,
	0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x61,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_netai_netai_proto_rawDescOnce sync.Once
	file_src_netai_netai_proto_rawDescData = file_src_netai_netai_proto_rawDesc
)

func file_src_netai_netai_proto_rawDescGZIP() []byte {
	file_src_netai_netai_proto_rawDescOnce.Do(func() {
		file_src_netai_netai_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_netai_netai_proto_rawDescData)
	})
	return file_src_netai_netai_proto_rawDescData
}

var file_src_netai_netai_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_netai_netai_proto_goTypes = []interface{}{
	(TrafficLoad)(0), // 0: netai.TrafficLoad
}
var file_src_netai_netai_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_netai_netai_proto_init() }
func file_src_netai_netai_proto_init() {
	if File_src_netai_netai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_netai_netai_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_netai_netai_proto_goTypes,
		DependencyIndexes: file_src_netai_netai_proto_depIdxs,
		EnumInfos:         file_src_netai_netai_proto_enumTypes,
	}.Build()
	File_src_netai_netai_proto = out.File
	file_src_netai_netai_proto_rawDesc = nil
	file_src_netai_netai_proto_goTypes = nil
	file_src_netai_netai_proto_depIdxs = nil
}
