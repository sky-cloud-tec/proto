// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: src/jrpc/eventti.proto

package jrpc

import (
	eventti "github.com/sky-cloud-tec/proto/v1/eventti"
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

type GetNotificationChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNotificationChannelRequest) Reset() {
	*x = GetNotificationChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_eventti_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationChannelRequest) ProtoMessage() {}

func (x *GetNotificationChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_eventti_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationChannelRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationChannelRequest) Descriptor() ([]byte, []int) {
	return file_src_jrpc_eventti_proto_rawDescGZIP(), []int{0}
}

type GetNotificationChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationChannels []*eventti.NotificationChannel `protobuf:"bytes,10,rep,name=notification_channels,json=notificationChannels,proto3" json:"notification_channels,omitempty"`
}

func (x *GetNotificationChannelResponse) Reset() {
	*x = GetNotificationChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_eventti_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationChannelResponse) ProtoMessage() {}

func (x *GetNotificationChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_eventti_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationChannelResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationChannelResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_eventti_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotificationChannelResponse) GetNotificationChannels() []*eventti.NotificationChannel {
	if x != nil {
		return x.NotificationChannels
	}
	return nil
}

type CreateNotificationChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateNotificationChannelRequest) Reset() {
	*x = CreateNotificationChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_eventti_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationChannelRequest) ProtoMessage() {}

func (x *CreateNotificationChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_eventti_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateNotificationChannelRequest) Descriptor() ([]byte, []int) {
	return file_src_jrpc_eventti_proto_rawDescGZIP(), []int{2}
}

type CreateNotificationChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,20,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateNotificationChannelResponse) Reset() {
	*x = CreateNotificationChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_eventti_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationChannelResponse) ProtoMessage() {}

func (x *CreateNotificationChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_eventti_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationChannelResponse.ProtoReflect.Descriptor instead.
func (*CreateNotificationChannelResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_eventti_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNotificationChannelResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_src_jrpc_eventti_proto protoreflect.FileDescriptor

var file_src_jrpc_eventti_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x72, 0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6a, 0x72, 0x70, 0x63, 0x1a, 0x19,
	0x73, 0x72, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x73, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x15,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x74, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x14, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22,
	0x22, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x42, 0x41, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x2e,
	0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a,
	0x72, 0x70, 0x63, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_jrpc_eventti_proto_rawDescOnce sync.Once
	file_src_jrpc_eventti_proto_rawDescData = file_src_jrpc_eventti_proto_rawDesc
)

func file_src_jrpc_eventti_proto_rawDescGZIP() []byte {
	file_src_jrpc_eventti_proto_rawDescOnce.Do(func() {
		file_src_jrpc_eventti_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_jrpc_eventti_proto_rawDescData)
	})
	return file_src_jrpc_eventti_proto_rawDescData
}

var file_src_jrpc_eventti_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_jrpc_eventti_proto_goTypes = []interface{}{
	(*GetNotificationChannelRequest)(nil),     // 0: jrpc.GetNotificationChannelRequest
	(*GetNotificationChannelResponse)(nil),    // 1: jrpc.GetNotificationChannelResponse
	(*CreateNotificationChannelRequest)(nil),  // 2: jrpc.CreateNotificationChannelRequest
	(*CreateNotificationChannelResponse)(nil), // 3: jrpc.CreateNotificationChannelResponse
	(*eventti.NotificationChannel)(nil),       // 4: eventti.NotificationChannel
}
var file_src_jrpc_eventti_proto_depIdxs = []int32{
	4, // 0: jrpc.GetNotificationChannelResponse.notification_channels:type_name -> eventti.NotificationChannel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_jrpc_eventti_proto_init() }
func file_src_jrpc_eventti_proto_init() {
	if File_src_jrpc_eventti_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_jrpc_eventti_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationChannelRequest); i {
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
		file_src_jrpc_eventti_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationChannelResponse); i {
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
		file_src_jrpc_eventti_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationChannelRequest); i {
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
		file_src_jrpc_eventti_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationChannelResponse); i {
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
			RawDescriptor: file_src_jrpc_eventti_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_jrpc_eventti_proto_goTypes,
		DependencyIndexes: file_src_jrpc_eventti_proto_depIdxs,
		MessageInfos:      file_src_jrpc_eventti_proto_msgTypes,
	}.Build()
	File_src_jrpc_eventti_proto = out.File
	file_src_jrpc_eventti_proto_rawDesc = nil
	file_src_jrpc_eventti_proto_goTypes = nil
	file_src_jrpc_eventti_proto_depIdxs = nil
}
