// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/baas/baas.proto

package baas

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

type UnitState int32

const (
	UnitState_UNKNOWN_WORKFLOW_STATE UnitState = 0
	UnitState_CREATED                UnitState = 1
	UnitState_RUNNING                UnitState = 2
	UnitState_SUCCEEDED              UnitState = 3
	UnitState_FAILED                 UnitState = 4
	UnitState_SUSPEND                UnitState = 5
	UnitState_PENDING                UnitState = 6
)

var UnitState_name = map[int32]string{
	0: "UNKNOWN_WORKFLOW_STATE",
	1: "CREATED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "SUSPEND",
	6: "PENDING",
}

var UnitState_value = map[string]int32{
	"UNKNOWN_WORKFLOW_STATE": 0,
	"CREATED":                1,
	"RUNNING":                2,
	"SUCCEEDED":              3,
	"FAILED":                 4,
	"SUSPEND":                5,
	"PENDING":                6,
}

func (x UnitState) String() string {
	return proto.EnumName(UnitState_name, int32(x))
}

func (UnitState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_080da5ced2d0d84a, []int{0}
}

func init() {
	proto.RegisterEnum("baas.UnitState", UnitState_name, UnitState_value)
}

func init() { proto.RegisterFile("src/baas/baas.proto", fileDescriptor_080da5ced2d0d84a) }

var fileDescriptor_080da5ced2d0d84a = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x4a, 0xd6,
	0x4f, 0x4a, 0x4c, 0x2c, 0x06, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6,
	0x56, 0x19, 0x17, 0x67, 0x68, 0x5e, 0x66, 0x49, 0x70, 0x49, 0x62, 0x49, 0xaa, 0x90, 0x14, 0x97,
	0x58, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x5f, 0x7c, 0xb8, 0x7f, 0x90, 0xb7, 0x9b, 0x8f, 0x7f,
	0x78, 0x7c, 0x70, 0x88, 0x63, 0x88, 0xab, 0x00, 0x83, 0x10, 0x37, 0x17, 0xbb, 0x73, 0x90, 0xab,
	0x63, 0x88, 0xab, 0x8b, 0x00, 0x23, 0x88, 0x13, 0x14, 0xea, 0xe7, 0xe7, 0xe9, 0xe7, 0x2e, 0xc0,
	0x24, 0xc4, 0xcb, 0xc5, 0x19, 0x1c, 0xea, 0xec, 0xec, 0xea, 0xea, 0xe2, 0xea, 0x22, 0xc0, 0x2c,
	0xc4, 0xc5, 0xc5, 0xe6, 0xe6, 0xe8, 0xe9, 0xe3, 0xea, 0x22, 0xc0, 0x02, 0x52, 0x17, 0x1c, 0x1a,
	0x1c, 0xe0, 0xea, 0xe7, 0x22, 0xc0, 0x0a, 0xe2, 0x80, 0x58, 0x20, 0x4d, 0x6c, 0x4e, 0x8e, 0x5c,
	0xe2, 0x79, 0xa9, 0x25, 0x7a, 0xc5, 0xd9, 0x95, 0xc9, 0x39, 0xf9, 0xa5, 0x29, 0x10, 0x37, 0xe9,
	0x81, 0x9c, 0x14, 0xa5, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f,
	0x9c, 0x5d, 0xa9, 0x0b, 0x56, 0xa0, 0x5b, 0x92, 0x9a, 0xac, 0x0f, 0x56, 0xa4, 0x5f, 0x66, 0x08,
	0xf6, 0x46, 0x12, 0x1b, 0x98, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x60, 0xc6, 0x51,
	0xde, 0x00, 0x00, 0x00,
}
