// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/auth/auth.proto

package auth

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

// app用户信息
type UserRoles struct {
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Roles                []string `protobuf:"bytes,20,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRoles) Reset()         { *m = UserRoles{} }
func (m *UserRoles) String() string { return proto.CompactTextString(m) }
func (*UserRoles) ProtoMessage()    {}
func (*UserRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_048b2b9424705cfd, []int{0}
}

func (m *UserRoles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRoles.Unmarshal(m, b)
}
func (m *UserRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRoles.Marshal(b, m, deterministic)
}
func (m *UserRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoles.Merge(m, src)
}
func (m *UserRoles) XXX_Size() int {
	return xxx_messageInfo_UserRoles.Size(m)
}
func (m *UserRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoles.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoles proto.InternalMessageInfo

func (m *UserRoles) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRoles) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

// app用户组信息
type UserGroupRoles struct {
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Roles                []string `protobuf:"bytes,20,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGroupRoles) Reset()         { *m = UserGroupRoles{} }
func (m *UserGroupRoles) String() string { return proto.CompactTextString(m) }
func (*UserGroupRoles) ProtoMessage()    {}
func (*UserGroupRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_048b2b9424705cfd, []int{1}
}

func (m *UserGroupRoles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupRoles.Unmarshal(m, b)
}
func (m *UserGroupRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupRoles.Marshal(b, m, deterministic)
}
func (m *UserGroupRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupRoles.Merge(m, src)
}
func (m *UserGroupRoles) XXX_Size() int {
	return xxx_messageInfo_UserGroupRoles.Size(m)
}
func (m *UserGroupRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupRoles.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupRoles proto.InternalMessageInfo

func (m *UserGroupRoles) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserGroupRoles) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*UserRoles)(nil), "auth.UserRoles")
	proto.RegisterType((*UserGroupRoles)(nil), "auth.UserGroupRoles")
}

func init() { proto.RegisterFile("src/auth/auth.proto", fileDescriptor_048b2b9424705cfd) }

var fileDescriptor_048b2b9424705cfd = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x4a, 0xd6,
	0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6,
	0x92, 0x29, 0x17, 0x67, 0x68, 0x71, 0x6a, 0x51, 0x50, 0x7e, 0x4e, 0x6a, 0xb1, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x97, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc2,
	0xc5, 0x5a, 0x04, 0x92, 0x94, 0x10, 0x51, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xac, 0xb8,
	0xf8, 0x40, 0xda, 0xdc, 0x8b, 0xf2, 0x4b, 0x0b, 0x48, 0xd4, 0xeb, 0xe4, 0xc8, 0x25, 0x9e, 0x97,
	0x5a, 0xa2, 0x57, 0x9c, 0x5d, 0x99, 0x9c, 0x93, 0x5f, 0x9a, 0x02, 0x71, 0x8e, 0x1e, 0xc8, 0x35,
	0x51, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5, 0xd9, 0x95,
	0xba, 0x60, 0x05, 0xba, 0x25, 0xa9, 0xc9, 0xfa, 0x60, 0x45, 0xfa, 0x65, 0x86, 0x60, 0x1f, 0x24,
	0xb1, 0x81, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x4e, 0x6f, 0x0d, 0xd9, 0x00,
	0x00, 0x00,
}
