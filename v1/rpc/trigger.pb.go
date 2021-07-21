// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/grpc/trigger.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	trigger "github.com/sky-cloud-tec/proto/v1/trigger"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ByKeyAndPageRequest struct {
	Key                  string   `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Page                 int32    `protobuf:"varint,20,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,30,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByKeyAndPageRequest) Reset()         { *m = ByKeyAndPageRequest{} }
func (m *ByKeyAndPageRequest) String() string { return proto.CompactTextString(m) }
func (*ByKeyAndPageRequest) ProtoMessage()    {}
func (*ByKeyAndPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{0}
}

func (m *ByKeyAndPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByKeyAndPageRequest.Unmarshal(m, b)
}
func (m *ByKeyAndPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByKeyAndPageRequest.Marshal(b, m, deterministic)
}
func (m *ByKeyAndPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByKeyAndPageRequest.Merge(m, src)
}
func (m *ByKeyAndPageRequest) XXX_Size() int {
	return xxx_messageInfo_ByKeyAndPageRequest.Size(m)
}
func (m *ByKeyAndPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByKeyAndPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByKeyAndPageRequest proto.InternalMessageInfo

func (m *ByKeyAndPageRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ByKeyAndPageRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ByKeyAndPageRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ByTriggerIdsRequest struct {
	Ids                  []int64  `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByTriggerIdsRequest) Reset()         { *m = ByTriggerIdsRequest{} }
func (m *ByTriggerIdsRequest) String() string { return proto.CompactTextString(m) }
func (*ByTriggerIdsRequest) ProtoMessage()    {}
func (*ByTriggerIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{1}
}

func (m *ByTriggerIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByTriggerIdsRequest.Unmarshal(m, b)
}
func (m *ByTriggerIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByTriggerIdsRequest.Marshal(b, m, deterministic)
}
func (m *ByTriggerIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByTriggerIdsRequest.Merge(m, src)
}
func (m *ByTriggerIdsRequest) XXX_Size() int {
	return xxx_messageInfo_ByTriggerIdsRequest.Size(m)
}
func (m *ByTriggerIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByTriggerIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByTriggerIdsRequest proto.InternalMessageInfo

func (m *ByTriggerIdsRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ByTriggerIdRequest struct {
	Id                   int64    `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByTriggerIdRequest) Reset()         { *m = ByTriggerIdRequest{} }
func (m *ByTriggerIdRequest) String() string { return proto.CompactTextString(m) }
func (*ByTriggerIdRequest) ProtoMessage()    {}
func (*ByTriggerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{2}
}

func (m *ByTriggerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByTriggerIdRequest.Unmarshal(m, b)
}
func (m *ByTriggerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByTriggerIdRequest.Marshal(b, m, deterministic)
}
func (m *ByTriggerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByTriggerIdRequest.Merge(m, src)
}
func (m *ByTriggerIdRequest) XXX_Size() int {
	return xxx_messageInfo_ByTriggerIdRequest.Size(m)
}
func (m *ByTriggerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByTriggerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByTriggerIdRequest proto.InternalMessageInfo

func (m *ByTriggerIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ByResourceRequest struct {
	Reousrce             *trigger.PaasResource `protobuf:"bytes,10,opt,name=reousrce,proto3" json:"reousrce,omitempty"`
	Page                 int32                 `protobuf:"varint,20,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32                 `protobuf:"varint,30,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ByResourceRequest) Reset()         { *m = ByResourceRequest{} }
func (m *ByResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ByResourceRequest) ProtoMessage()    {}
func (*ByResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{3}
}

func (m *ByResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByResourceRequest.Unmarshal(m, b)
}
func (m *ByResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByResourceRequest.Marshal(b, m, deterministic)
}
func (m *ByResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByResourceRequest.Merge(m, src)
}
func (m *ByResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ByResourceRequest.Size(m)
}
func (m *ByResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByResourceRequest proto.InternalMessageInfo

func (m *ByResourceRequest) GetReousrce() *trigger.PaasResource {
	if m != nil {
		return m.Reousrce
	}
	return nil
}

func (m *ByResourceRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ByResourceRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type TriggerRequest struct {
	Trigger              *trigger.Trigger `protobuf:"bytes,10,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{4}
}

func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (m *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(m, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

func (m *TriggerRequest) GetTrigger() *trigger.Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type ListTriggerRequest struct {
	Triggers             []*trigger.Trigger `protobuf:"bytes,10,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListTriggerRequest) Reset()         { *m = ListTriggerRequest{} }
func (m *ListTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*ListTriggerRequest) ProtoMessage()    {}
func (*ListTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{5}
}

func (m *ListTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggerRequest.Unmarshal(m, b)
}
func (m *ListTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggerRequest.Marshal(b, m, deterministic)
}
func (m *ListTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggerRequest.Merge(m, src)
}
func (m *ListTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_ListTriggerRequest.Size(m)
}
func (m *ListTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggerRequest proto.InternalMessageInfo

func (m *ListTriggerRequest) GetTriggers() []*trigger.Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{6}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type TriggerResponse struct {
	Trigger              *trigger.Trigger `protobuf:"bytes,10,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TriggerResponse) Reset()         { *m = TriggerResponse{} }
func (m *TriggerResponse) String() string { return proto.CompactTextString(m) }
func (*TriggerResponse) ProtoMessage()    {}
func (*TriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{7}
}

func (m *TriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerResponse.Unmarshal(m, b)
}
func (m *TriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerResponse.Marshal(b, m, deterministic)
}
func (m *TriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerResponse.Merge(m, src)
}
func (m *TriggerResponse) XXX_Size() int {
	return xxx_messageInfo_TriggerResponse.Size(m)
}
func (m *TriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerResponse proto.InternalMessageInfo

func (m *TriggerResponse) GetTrigger() *trigger.Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type DeleteTriggerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTriggerResponse) Reset()         { *m = DeleteTriggerResponse{} }
func (m *DeleteTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTriggerResponse) ProtoMessage()    {}
func (*DeleteTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{8}
}

func (m *DeleteTriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTriggerResponse.Unmarshal(m, b)
}
func (m *DeleteTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTriggerResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTriggerResponse.Merge(m, src)
}
func (m *DeleteTriggerResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTriggerResponse.Size(m)
}
func (m *DeleteTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTriggerResponse proto.InternalMessageInfo

type ListTriggerResponse struct {
	Count                int32              `protobuf:"varint,10,opt,name=count,proto3" json:"count,omitempty"`
	PageSize             int32              `protobuf:"varint,20,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	AllPage              int32              `protobuf:"varint,30,opt,name=allPage,proto3" json:"allPage,omitempty"`
	Triggers             []*trigger.Trigger `protobuf:"bytes,40,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListTriggerResponse) Reset()         { *m = ListTriggerResponse{} }
func (m *ListTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*ListTriggerResponse) ProtoMessage()    {}
func (*ListTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22f8bbc941bf2b09, []int{9}
}

func (m *ListTriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggerResponse.Unmarshal(m, b)
}
func (m *ListTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggerResponse.Marshal(b, m, deterministic)
}
func (m *ListTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggerResponse.Merge(m, src)
}
func (m *ListTriggerResponse) XXX_Size() int {
	return xxx_messageInfo_ListTriggerResponse.Size(m)
}
func (m *ListTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggerResponse proto.InternalMessageInfo

func (m *ListTriggerResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListTriggerResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTriggerResponse) GetAllPage() int32 {
	if m != nil {
		return m.AllPage
	}
	return 0
}

func (m *ListTriggerResponse) GetTriggers() []*trigger.Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func init() {
	proto.RegisterType((*ByKeyAndPageRequest)(nil), "rpc.ByKeyAndPageRequest")
	proto.RegisterType((*ByTriggerIdsRequest)(nil), "rpc.ByTriggerIdsRequest")
	proto.RegisterType((*ByTriggerIdRequest)(nil), "rpc.ByTriggerIdRequest")
	proto.RegisterType((*ByResourceRequest)(nil), "rpc.ByResourceRequest")
	proto.RegisterType((*TriggerRequest)(nil), "rpc.TriggerRequest")
	proto.RegisterType((*ListTriggerRequest)(nil), "rpc.ListTriggerRequest")
	proto.RegisterType((*EmptyResponse)(nil), "rpc.EmptyResponse")
	proto.RegisterType((*TriggerResponse)(nil), "rpc.TriggerResponse")
	proto.RegisterType((*DeleteTriggerResponse)(nil), "rpc.DeleteTriggerResponse")
	proto.RegisterType((*ListTriggerResponse)(nil), "rpc.ListTriggerResponse")
}

func init() { proto.RegisterFile("src/grpc/trigger.proto", fileDescriptor_22f8bbc941bf2b09) }

var fileDescriptor_22f8bbc941bf2b09 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x1b, 0x42, 0xc3, 0x44, 0x49, 0xcb, 0x26, 0x4d, 0x8d, 0x0f, 0x28, 0xb2, 0x40, 0x44,
	0x88, 0xda, 0x6a, 0xb9, 0xb6, 0x14, 0x12, 0x3e, 0x05, 0x12, 0x95, 0x03, 0x17, 0x6e, 0xee, 0x7a,
	0xe4, 0x5a, 0xf9, 0xb0, 0xd9, 0x5d, 0x57, 0x32, 0x7f, 0x82, 0x1f, 0xca, 0x9f, 0x40, 0xbb, 0xb6,
	0x57, 0xd9, 0x24, 0x10, 0x5a, 0x6e, 0xbb, 0x33, 0x6f, 0xde, 0x9b, 0x79, 0x3b, 0x36, 0xf4, 0x38,
	0xa3, 0x5e, 0xc4, 0x52, 0xea, 0x09, 0x16, 0x47, 0x11, 0x32, 0x37, 0x65, 0x89, 0x48, 0x48, 0x8d,
	0xa5, 0xd4, 0x7e, 0x20, 0x93, 0x65, 0xdc, 0xcc, 0x3b, 0x9f, 0xa1, 0x33, 0xcc, 0x3f, 0x62, 0xfe,
	0x6a, 0x1e, 0x5e, 0x04, 0x11, 0xfa, 0xf8, 0x3d, 0x43, 0x2e, 0xc8, 0x3e, 0xd4, 0x26, 0x98, 0x5b,
	0xd0, 0xdf, 0x1e, 0xdc, 0xf3, 0xe5, 0x91, 0x10, 0xb8, 0x93, 0x06, 0x11, 0x5a, 0xdd, 0xfe, 0xf6,
	0xa0, 0xee, 0xab, 0xb3, 0x8c, 0xf1, 0xf8, 0x07, 0x5a, 0x0f, 0x8b, 0x98, 0x3c, 0x3b, 0x4f, 0x24,
	0xe1, 0x97, 0x42, 0xe3, 0x43, 0xc8, 0x17, 0x08, 0xe3, 0x90, 0x5b, 0xd0, 0xaf, 0x0d, 0x6a, 0xbe,
	0x3c, 0x3a, 0x8f, 0x80, 0x2c, 0x00, 0x2b, 0x5c, 0x1b, 0x76, 0xe2, 0x50, 0xe9, 0xd6, 0xfc, 0x9d,
	0x38, 0x74, 0xe6, 0x70, 0x7f, 0x98, 0xfb, 0xc8, 0x93, 0x8c, 0x51, 0xdd, 0xdd, 0x31, 0x34, 0x18,
	0x26, 0x19, 0x67, 0x14, 0x15, 0xb4, 0x79, 0x72, 0xe0, 0x56, 0x63, 0x5d, 0x04, 0x01, 0xd7, 0x78,
	0x0d, 0xfb, 0xe7, 0xf6, 0x4f, 0xa1, 0x5d, 0xf6, 0x54, 0x89, 0x3d, 0x85, 0xdd, 0x92, 0xbb, 0xd4,
	0xda, 0xd7, 0x5a, 0x15, 0xb2, 0x02, 0x38, 0x43, 0x20, 0x9f, 0x62, 0x2e, 0x96, 0x18, 0x9e, 0x41,
	0xa3, 0x04, 0x14, 0x06, 0xac, 0xa3, 0xd0, 0x08, 0x67, 0x0f, 0x5a, 0x6f, 0x66, 0xa9, 0x90, 0x43,
	0xa7, 0xc9, 0x9c, 0xa3, 0x73, 0x06, 0x7b, 0x9a, 0xb0, 0x08, 0xdd, 0xa8, 0xa7, 0x43, 0x38, 0x78,
	0x8d, 0x53, 0x14, 0xb8, 0x44, 0xe2, 0xfc, 0xdc, 0x86, 0x8e, 0xd1, 0x6d, 0x49, 0xde, 0x85, 0x3a,
	0x4d, 0xb2, 0xb9, 0x50, 0xd4, 0x75, 0xbf, 0xb8, 0x10, 0x1b, 0x1a, 0xd2, 0xb4, 0xb1, 0x34, 0xac,
	0x30, 0x51, 0xdf, 0x89, 0x05, 0xbb, 0xc1, 0x74, 0x2a, 0xf7, 0xa7, 0xf4, 0xb2, 0xba, 0x1a, 0xa3,
	0x0f, 0x36, 0x8d, 0x7e, 0xf2, 0xab, 0xae, 0xdd, 0x1f, 0x23, 0xbb, 0x8e, 0x29, 0x92, 0x11, 0x34,
	0x17, 0x7a, 0x24, 0x96, 0xcb, 0x52, 0xea, 0xae, 0xd9, 0x58, 0xbb, 0xc8, 0xac, 0x99, 0xc7, 0xd9,
	0x22, 0xa7, 0xd0, 0xfa, 0x9a, 0x86, 0x81, 0xb6, 0x80, 0x74, 0x14, 0xd8, 0x7c, 0x26, 0xbb, 0x6b,
	0x06, 0x75, 0xf5, 0x39, 0xb4, 0x8d, 0x6a, 0x4e, 0x0e, 0x57, 0xb5, 0x0a, 0x0a, 0xa2, 0x12, 0xe6,
	0xf3, 0x29, 0xf9, 0x11, 0xc3, 0xff, 0x90, 0x37, 0xaa, 0x6f, 0x2c, 0xff, 0x16, 0x5a, 0xc6, 0x02,
	0x94, 0xf5, 0xab, 0x1f, 0x9f, 0x6d, 0xab, 0xc4, 0xfa, 0x6d, 0xd9, 0x22, 0xef, 0xa1, 0x6d, 0xa4,
	0xb8, 0x7e, 0x8d, 0x95, 0xcf, 0x7d, 0x03, 0xd3, 0x19, 0xc0, 0x3b, 0x14, 0x1b, 0xdb, 0xf9, 0x93,
	0x23, 0x2f, 0xa0, 0x39, 0x16, 0x49, 0x7a, 0xeb, 0xfa, 0x97, 0xd0, 0xf2, 0x91, 0x67, 0x33, 0xbc,
	0x35, 0xc3, 0x08, 0x5a, 0x63, 0x0c, 0x18, 0xbd, 0xaa, 0x18, 0x7a, 0x25, 0xc3, 0xd2, 0x9f, 0xea,
	0x6f, 0x5b, 0x39, 0x3c, 0x87, 0xde, 0x1c, 0x85, 0xcb, 0x27, 0x39, 0x9d, 0x26, 0x59, 0x58, 0xfc,
	0x90, 0x25, 0xfe, 0xdb, 0xe3, 0x28, 0x16, 0x57, 0xd9, 0xa5, 0x4b, 0x93, 0x99, 0xc7, 0x27, 0xf9,
	0x91, 0xca, 0x1f, 0x09, 0xa4, 0x9e, 0xc2, 0x78, 0xd7, 0xc7, 0x1e, 0x4b, 0xe9, 0xe5, 0x5d, 0x75,
	0x7b, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x4f, 0xf3, 0x62, 0xfc, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TriggerServiceClient is the client API for TriggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TriggerServiceClient interface {
	// 根据关键字分页查找触发器
	ListTrigger(ctx context.Context, in *ByKeyAndPageRequest, opts ...grpc.CallOption) (*ListTriggerResponse, error)
	// 更新触发器信息
	UpdateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// 批量更新触发器信息
	UpdateTriggers(ctx context.Context, in *ListTriggerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// 创建触发器信息
	CreateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// 批量创建触发器信息
	CreateTriggers(ctx context.Context, in *ListTriggerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// 删除触发器信息
	DeleteTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*DeleteTriggerResponse, error)
	// 批量删除触发器信息
	DeleteTriggers(ctx context.Context, in *ByTriggerIdsRequest, opts ...grpc.CallOption) (*DeleteTriggerResponse, error)
	// 获取触发器信息
	GetTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// 暂停触发器
	StopTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// 恢复触发器
	ResumeTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// 根据资源查找对应的触发器
	SearchTrigger(ctx context.Context, in *ByResourceRequest, opts ...grpc.CallOption) (*ListTriggerResponse, error)
}

type triggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerServiceClient(cc grpc.ClientConnInterface) TriggerServiceClient {
	return &triggerServiceClient{cc}
}

func (c *triggerServiceClient) ListTrigger(ctx context.Context, in *ByKeyAndPageRequest, opts ...grpc.CallOption) (*ListTriggerResponse, error) {
	out := new(ListTriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/ListTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) UpdateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/UpdateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) UpdateTriggers(ctx context.Context, in *ListTriggerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/UpdateTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) CreateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/CreateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) CreateTriggers(ctx context.Context, in *ListTriggerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/CreateTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) DeleteTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*DeleteTriggerResponse, error) {
	out := new(DeleteTriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/DeleteTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) DeleteTriggers(ctx context.Context, in *ByTriggerIdsRequest, opts ...grpc.CallOption) (*DeleteTriggerResponse, error) {
	out := new(DeleteTriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/DeleteTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) GetTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/GetTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) StopTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/StopTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) ResumeTrigger(ctx context.Context, in *ByTriggerIdRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/ResumeTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SearchTrigger(ctx context.Context, in *ByResourceRequest, opts ...grpc.CallOption) (*ListTriggerResponse, error) {
	out := new(ListTriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.TriggerService/SearchTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerServiceServer is the server API for TriggerService service.
type TriggerServiceServer interface {
	// 根据关键字分页查找触发器
	ListTrigger(context.Context, *ByKeyAndPageRequest) (*ListTriggerResponse, error)
	// 更新触发器信息
	UpdateTrigger(context.Context, *TriggerRequest) (*TriggerResponse, error)
	// 批量更新触发器信息
	UpdateTriggers(context.Context, *ListTriggerRequest) (*EmptyResponse, error)
	// 创建触发器信息
	CreateTrigger(context.Context, *TriggerRequest) (*TriggerResponse, error)
	// 批量创建触发器信息
	CreateTriggers(context.Context, *ListTriggerRequest) (*EmptyResponse, error)
	// 删除触发器信息
	DeleteTrigger(context.Context, *ByTriggerIdRequest) (*DeleteTriggerResponse, error)
	// 批量删除触发器信息
	DeleteTriggers(context.Context, *ByTriggerIdsRequest) (*DeleteTriggerResponse, error)
	// 获取触发器信息
	GetTrigger(context.Context, *ByTriggerIdRequest) (*TriggerResponse, error)
	// 暂停触发器
	StopTrigger(context.Context, *ByTriggerIdRequest) (*TriggerResponse, error)
	// 恢复触发器
	ResumeTrigger(context.Context, *ByTriggerIdRequest) (*TriggerResponse, error)
	// 根据资源查找对应的触发器
	SearchTrigger(context.Context, *ByResourceRequest) (*ListTriggerResponse, error)
}

// UnimplementedTriggerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTriggerServiceServer struct {
}

func (*UnimplementedTriggerServiceServer) ListTrigger(ctx context.Context, req *ByKeyAndPageRequest) (*ListTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) UpdateTrigger(ctx context.Context, req *TriggerRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) UpdateTriggers(ctx context.Context, req *ListTriggerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTriggers not implemented")
}
func (*UnimplementedTriggerServiceServer) CreateTrigger(ctx context.Context, req *TriggerRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) CreateTriggers(ctx context.Context, req *ListTriggerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTriggers not implemented")
}
func (*UnimplementedTriggerServiceServer) DeleteTrigger(ctx context.Context, req *ByTriggerIdRequest) (*DeleteTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) DeleteTriggers(ctx context.Context, req *ByTriggerIdsRequest) (*DeleteTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTriggers not implemented")
}
func (*UnimplementedTriggerServiceServer) GetTrigger(ctx context.Context, req *ByTriggerIdRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) StopTrigger(ctx context.Context, req *ByTriggerIdRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) ResumeTrigger(ctx context.Context, req *ByTriggerIdRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeTrigger not implemented")
}
func (*UnimplementedTriggerServiceServer) SearchTrigger(ctx context.Context, req *ByResourceRequest) (*ListTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTrigger not implemented")
}

func RegisterTriggerServiceServer(s *grpc.Server, srv TriggerServiceServer) {
	s.RegisterService(&_TriggerService_serviceDesc, srv)
}

func _TriggerService_ListTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByKeyAndPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).ListTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/ListTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).ListTrigger(ctx, req.(*ByKeyAndPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_UpdateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).UpdateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/UpdateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).UpdateTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_UpdateTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).UpdateTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/UpdateTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).UpdateTriggers(ctx, req.(*ListTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_CreateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).CreateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/CreateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).CreateTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_CreateTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).CreateTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/CreateTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).CreateTriggers(ctx, req.(*ListTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_DeleteTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByTriggerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).DeleteTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/DeleteTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).DeleteTrigger(ctx, req.(*ByTriggerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_DeleteTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByTriggerIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).DeleteTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/DeleteTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).DeleteTriggers(ctx, req.(*ByTriggerIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_GetTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByTriggerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).GetTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/GetTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).GetTrigger(ctx, req.(*ByTriggerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_StopTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByTriggerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).StopTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/StopTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).StopTrigger(ctx, req.(*ByTriggerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_ResumeTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByTriggerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).ResumeTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/ResumeTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).ResumeTrigger(ctx, req.(*ByTriggerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_SearchTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).SearchTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TriggerService/SearchTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).SearchTrigger(ctx, req.(*ByResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TriggerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.TriggerService",
	HandlerType: (*TriggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTrigger",
			Handler:    _TriggerService_ListTrigger_Handler,
		},
		{
			MethodName: "UpdateTrigger",
			Handler:    _TriggerService_UpdateTrigger_Handler,
		},
		{
			MethodName: "UpdateTriggers",
			Handler:    _TriggerService_UpdateTriggers_Handler,
		},
		{
			MethodName: "CreateTrigger",
			Handler:    _TriggerService_CreateTrigger_Handler,
		},
		{
			MethodName: "CreateTriggers",
			Handler:    _TriggerService_CreateTriggers_Handler,
		},
		{
			MethodName: "DeleteTrigger",
			Handler:    _TriggerService_DeleteTrigger_Handler,
		},
		{
			MethodName: "DeleteTriggers",
			Handler:    _TriggerService_DeleteTriggers_Handler,
		},
		{
			MethodName: "GetTrigger",
			Handler:    _TriggerService_GetTrigger_Handler,
		},
		{
			MethodName: "StopTrigger",
			Handler:    _TriggerService_StopTrigger_Handler,
		},
		{
			MethodName: "ResumeTrigger",
			Handler:    _TriggerService_ResumeTrigger_Handler,
		},
		{
			MethodName: "SearchTrigger",
			Handler:    _TriggerService_SearchTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/trigger.proto",
}