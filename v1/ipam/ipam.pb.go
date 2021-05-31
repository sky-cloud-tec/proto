// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: src/ipam/ipam.proto

package ipam

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

type IpamItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip     string `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty"`
	Status string `protobuf:"bytes,20,opt,name=status,proto3" json:"status,omitempty"`
	Os     string `protobuf:"bytes,30,opt,name=os,proto3" json:"os,omitempty"`
	Mac    string `protobuf:"bytes,40,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *IpamItem) Reset() {
	*x = IpamItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ipam_ipam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpamItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpamItem) ProtoMessage() {}

func (x *IpamItem) ProtoReflect() protoreflect.Message {
	mi := &file_src_ipam_ipam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpamItem.ProtoReflect.Descriptor instead.
func (*IpamItem) Descriptor() ([]byte, []int) {
	return file_src_ipam_ipam_proto_rawDescGZIP(), []int{0}
}

func (x *IpamItem) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IpamItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IpamItem) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *IpamItem) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

var File_src_ipam_ipam_proto protoreflect.FileDescriptor

var file_src_ipam_ipam_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x70, 0x61, 0x6d, 0x22, 0x54, 0x0a, 0x08, 0x49,
	0x70, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61,
	0x63, 0x42, 0x41, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x70, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_ipam_ipam_proto_rawDescOnce sync.Once
	file_src_ipam_ipam_proto_rawDescData = file_src_ipam_ipam_proto_rawDesc
)

func file_src_ipam_ipam_proto_rawDescGZIP() []byte {
	file_src_ipam_ipam_proto_rawDescOnce.Do(func() {
		file_src_ipam_ipam_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_ipam_ipam_proto_rawDescData)
	})
	return file_src_ipam_ipam_proto_rawDescData
}

var file_src_ipam_ipam_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_ipam_ipam_proto_goTypes = []interface{}{
	(*IpamItem)(nil), // 0: ipam.IpamItem
}
var file_src_ipam_ipam_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_ipam_ipam_proto_init() }
func file_src_ipam_ipam_proto_init() {
	if File_src_ipam_ipam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_ipam_ipam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpamItem); i {
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
			RawDescriptor: file_src_ipam_ipam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_ipam_ipam_proto_goTypes,
		DependencyIndexes: file_src_ipam_ipam_proto_depIdxs,
		MessageInfos:      file_src_ipam_ipam_proto_msgTypes,
	}.Build()
	File_src_ipam_ipam_proto = out.File
	file_src_ipam_ipam_proto_rawDesc = nil
	file_src_ipam_ipam_proto_goTypes = nil
	file_src_ipam_ipam_proto_depIdxs = nil
}
