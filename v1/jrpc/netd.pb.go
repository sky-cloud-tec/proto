// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/jrpc/netd.proto

package jrpc

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

type FixType int32

const (
	FixType_UNKNOWN_FIX_TYPE FixType = 0
	FixType_APPEND           FixType = 1
	FixType_REPLACE_ALL      FixType = 2
	FixType_DELETE_ONE       FixType = 3
	FixType_REPLACE_ONE      FixType = 4
)

var FixType_name = map[int32]string{
	0: "UNKNOWN_FIX_TYPE",
	1: "APPEND",
	2: "REPLACE_ALL",
	3: "DELETE_ONE",
	4: "REPLACE_ONE",
}

var FixType_value = map[string]int32{
	"UNKNOWN_FIX_TYPE": 0,
	"APPEND":           1,
	"REPLACE_ALL":      2,
	"DELETE_ONE":       3,
	"REPLACE_ONE":      4,
}

func (x FixType) String() string {
	return proto.EnumName(FixType_name, int32(x))
}

func (FixType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fe09a602917bb371, []int{0}
}

type CliRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CliRequest) Reset()         { *m = CliRequest{} }
func (m *CliRequest) String() string { return proto.CompactTextString(m) }
func (*CliRequest) ProtoMessage()    {}
func (*CliRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe09a602917bb371, []int{0}
}

func (m *CliRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliRequest.Unmarshal(m, b)
}
func (m *CliRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliRequest.Marshal(b, m, deterministic)
}
func (m *CliRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliRequest.Merge(m, src)
}
func (m *CliRequest) XXX_Size() int {
	return xxx_messageInfo_CliRequest.Size(m)
}
func (m *CliRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CliRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CliRequest proto.InternalMessageInfo

type OperatorHotfixRequest struct {
	FixType              FixType  `protobuf:"varint,5,opt,name=fix_type,json=fixType,proto3,enum=jrpc.FixType" json:"fix_type,omitempty"`
	Vendor               string   `protobuf:"bytes,10,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Type                 string   `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
	Version              string   `protobuf:"bytes,30,opt,name=version,proto3" json:"version,omitempty"`
	Mode                 string   `protobuf:"bytes,40,opt,name=mode,proto3" json:"mode,omitempty"`
	Prompts              []string `protobuf:"bytes,50,rep,name=prompts,proto3" json:"prompts,omitempty"`
	Errs                 []string `protobuf:"bytes,60,rep,name=errs,proto3" json:"errs,omitempty"`
	Index                int32    `protobuf:"varint,70,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorHotfixRequest) Reset()         { *m = OperatorHotfixRequest{} }
func (m *OperatorHotfixRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorHotfixRequest) ProtoMessage()    {}
func (*OperatorHotfixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe09a602917bb371, []int{1}
}

func (m *OperatorHotfixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorHotfixRequest.Unmarshal(m, b)
}
func (m *OperatorHotfixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorHotfixRequest.Marshal(b, m, deterministic)
}
func (m *OperatorHotfixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorHotfixRequest.Merge(m, src)
}
func (m *OperatorHotfixRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorHotfixRequest.Size(m)
}
func (m *OperatorHotfixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorHotfixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorHotfixRequest proto.InternalMessageInfo

func (m *OperatorHotfixRequest) GetFixType() FixType {
	if m != nil {
		return m.FixType
	}
	return FixType_UNKNOWN_FIX_TYPE
}

func (m *OperatorHotfixRequest) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *OperatorHotfixRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OperatorHotfixRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OperatorHotfixRequest) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *OperatorHotfixRequest) GetPrompts() []string {
	if m != nil {
		return m.Prompts
	}
	return nil
}

func (m *OperatorHotfixRequest) GetErrs() []string {
	if m != nil {
		return m.Errs
	}
	return nil
}

func (m *OperatorHotfixRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type OperatorHotfixResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorHotfixResponse) Reset()         { *m = OperatorHotfixResponse{} }
func (m *OperatorHotfixResponse) String() string { return proto.CompactTextString(m) }
func (*OperatorHotfixResponse) ProtoMessage()    {}
func (*OperatorHotfixResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe09a602917bb371, []int{2}
}

func (m *OperatorHotfixResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorHotfixResponse.Unmarshal(m, b)
}
func (m *OperatorHotfixResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorHotfixResponse.Marshal(b, m, deterministic)
}
func (m *OperatorHotfixResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorHotfixResponse.Merge(m, src)
}
func (m *OperatorHotfixResponse) XXX_Size() int {
	return xxx_messageInfo_OperatorHotfixResponse.Size(m)
}
func (m *OperatorHotfixResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorHotfixResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorHotfixResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("jrpc.FixType", FixType_name, FixType_value)
	proto.RegisterType((*CliRequest)(nil), "jrpc.CliRequest")
	proto.RegisterType((*OperatorHotfixRequest)(nil), "jrpc.OperatorHotfixRequest")
	proto.RegisterType((*OperatorHotfixResponse)(nil), "jrpc.OperatorHotfixResponse")
}

func init() { proto.RegisterFile("src/jrpc/netd.proto", fileDescriptor_fe09a602917bb371) }

var fileDescriptor_fe09a602917bb371 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x86, 0xe7, 0xe5, 0x6b, 0x3b, 0xdb, 0x32, 0xa3, 0x65, 0x99, 0xae, 0x86, 0xc9, 0xc5, 0x30,
	0x83, 0xd8, 0x2c, 0xbb, 0xdd, 0x8d, 0x97, 0x28, 0xb4, 0xd4, 0x38, 0xc6, 0xa4, 0xf4, 0x03, 0x8a,
	0x69, 0x6c, 0xa5, 0x75, 0x93, 0x58, 0xaa, 0xa4, 0x04, 0xfb, 0xd7, 0xf6, 0xaf, 0x14, 0x29, 0x2e,
	0x94, 0xde, 0xbd, 0xef, 0xa3, 0x47, 0x82, 0xa3, 0x03, 0xdf, 0xa4, 0xc8, 0xfc, 0x07, 0xc1, 0x33,
	0xbf, 0xa4, 0x2a, 0xf7, 0xb8, 0x60, 0x8a, 0xa1, 0xb6, 0x06, 0xa3, 0xcf, 0x00, 0xd3, 0x6d, 0x91,
	0xd0, 0xc7, 0x3d, 0x95, 0x6a, 0xf4, 0x64, 0xc1, 0xf7, 0x05, 0xa7, 0xe2, 0x56, 0x31, 0x71, 0xc2,
	0xd4, 0xba, 0xa8, 0x9a, 0x13, 0xe4, 0xc2, 0x87, 0x75, 0x51, 0xa5, 0xaa, 0xe6, 0x14, 0x77, 0x1c,
	0xcb, 0xed, 0x4f, 0xbe, 0x78, 0xfa, 0x01, 0x6f, 0x5e, 0x54, 0xcb, 0x9a, 0xd3, 0xa4, 0xb7, 0x3e,
	0x06, 0x34, 0x84, 0xee, 0x81, 0x96, 0x39, 0x13, 0x18, 0x1c, 0xcb, 0xfd, 0x98, 0x34, 0x0d, 0x21,
	0x68, 0x9b, 0xdb, 0x03, 0x43, 0x4d, 0x46, 0x18, 0x7a, 0x07, 0x2a, 0x64, 0xc1, 0x4a, 0xfc, 0xd3,
	0xe0, 0x97, 0xaa, 0xed, 0x1d, 0xcb, 0x29, 0x76, 0x8f, 0xb6, 0xce, 0xda, 0xe6, 0x82, 0xed, 0xb8,
	0x92, 0x78, 0xe2, 0xb4, 0xb4, 0xdd, 0x54, 0x6d, 0x53, 0x21, 0x24, 0xfe, 0x67, 0xb0, 0xc9, 0x68,
	0x00, 0x9d, 0xa2, 0xcc, 0x69, 0x85, 0xe7, 0x8e, 0xe5, 0x76, 0x92, 0x63, 0x19, 0x61, 0x18, 0xbe,
	0x1d, 0x50, 0x72, 0x56, 0x4a, 0xfa, 0xfb, 0x06, 0x7a, 0xcd, 0x2c, 0x68, 0x00, 0xf6, 0x79, 0x74,
	0x16, 0x2d, 0x2e, 0xa2, 0x74, 0x7e, 0x7a, 0x99, 0x2e, 0xaf, 0x62, 0x62, 0xbf, 0x43, 0x00, 0xdd,
	0x20, 0x8e, 0x49, 0x34, 0xb3, 0x2d, 0xf4, 0x15, 0x3e, 0x25, 0x24, 0x0e, 0x83, 0x29, 0x49, 0x83,
	0x30, 0xb4, 0xdf, 0xa3, 0x3e, 0xc0, 0x8c, 0x84, 0x64, 0x49, 0xd2, 0x45, 0x44, 0xec, 0xd6, 0x6b,
	0x41, 0x83, 0xf6, 0xff, 0x00, 0x7e, 0x94, 0x54, 0x79, 0x72, 0x53, 0x67, 0x5b, 0xb6, 0x6f, 0x96,
	0x60, 0xbe, 0xf0, 0xfa, 0xd7, 0x5d, 0xa1, 0xee, 0xf7, 0x2b, 0x2f, 0x63, 0x3b, 0x5f, 0x6e, 0xea,
	0xb1, 0x11, 0xc6, 0x8a, 0x66, 0xbe, 0x91, 0xfc, 0xc3, 0x1f, 0xb3, 0xbc, 0x55, 0xd7, 0xd4, 0xbf,
	0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x42, 0x4a, 0x98, 0xcf, 0x01, 0x00, 0x00,
}
