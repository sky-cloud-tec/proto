// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventti/eventti.proto

/*
Package eventti is a generated protocol buffer package.

It is generated from these files:
	eventti/eventti.proto

It has these top-level messages:
	From
	DeviceConfigUpdatedMsg
	Event
	EmailServerCfg
	SmsCfg
	Target
	CreateNotificationChannelReq
	NotificationChannel
	CreateNotificationChannelRes
	GetNotificationChannelReq
	GetNotificationChannelRes
*/
package eventti

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import common "github.com/sky-cloud-tec/proto/v1/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventType int32

const (
	EventType_UNKNOWN_EVENT_TYPE            EventType = 0
	EventType_DEVICE_CONFIG_UPDATED         EventType = 1
	EventType_DEVICE_RAW_CONFIG_SYNCING     EventType = 2
	EventType_DEVICE_RAW_CONFIG_SYNCED      EventType = 3
	EventType_DEVICE_OBJECT_CONFIG_ANALYSED EventType = 4
)

var EventType_name = map[int32]string{
	0: "UNKNOWN_EVENT_TYPE",
	1: "DEVICE_CONFIG_UPDATED",
	2: "DEVICE_RAW_CONFIG_SYNCING",
	3: "DEVICE_RAW_CONFIG_SYNCED",
	4: "DEVICE_OBJECT_CONFIG_ANALYSED",
}
var EventType_value = map[string]int32{
	"UNKNOWN_EVENT_TYPE":            0,
	"DEVICE_CONFIG_UPDATED":         1,
	"DEVICE_RAW_CONFIG_SYNCING":     2,
	"DEVICE_RAW_CONFIG_SYNCED":      3,
	"DEVICE_OBJECT_CONFIG_ANALYSED": 4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NotificationChannelType int32

const (
	NotificationChannelType_UNKNOWN_NOTIFICATION_CHANNEL_TYPE NotificationChannelType = 0
	NotificationChannelType_EMAIL                             NotificationChannelType = 10
	NotificationChannelType_PHONE                             NotificationChannelType = 20
	NotificationChannelType_WEBSOCKET                         NotificationChannelType = 30
)

var NotificationChannelType_name = map[int32]string{
	0:  "UNKNOWN_NOTIFICATION_CHANNEL_TYPE",
	10: "EMAIL",
	20: "PHONE",
	30: "WEBSOCKET",
}
var NotificationChannelType_value = map[string]int32{
	"UNKNOWN_NOTIFICATION_CHANNEL_TYPE": 0,
	"EMAIL":     10,
	"PHONE":     20,
	"WEBSOCKET": 30,
}

func (x NotificationChannelType) String() string {
	return proto.EnumName(NotificationChannelType_name, int32(x))
}
func (NotificationChannelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EmailProtocol int32

const (
	EmailProtocol_UNKNOWN_PROTOCOL EmailProtocol = 0
	EmailProtocol_SMTP             EmailProtocol = 10
	EmailProtocol_SECURE_SMTP      EmailProtocol = 20
)

var EmailProtocol_name = map[int32]string{
	0:  "UNKNOWN_PROTOCOL",
	10: "SMTP",
	20: "SECURE_SMTP",
}
var EmailProtocol_value = map[string]int32{
	"UNKNOWN_PROTOCOL": 0,
	"SMTP":             10,
	"SECURE_SMTP":      20,
}

func (x EmailProtocol) String() string {
	return proto.EnumName(EmailProtocol_name, int32(x))
}
func (EmailProtocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type From struct {
	Host   string `protobuf:"bytes,10,opt,name=host" json:"host,omitempty"`
	Module string `protobuf:"bytes,20,opt,name=module" json:"module,omitempty"`
}

func (m *From) Reset()                    { *m = From{} }
func (m *From) String() string            { return proto.CompactTextString(m) }
func (*From) ProtoMessage()               {}
func (*From) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *From) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *From) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

type DeviceConfigUpdatedMsg struct {
	DeviceId   string `protobuf:"bytes,10,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	DeviceName string `protobuf:"bytes,20,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	Message    string `protobuf:"bytes,30,opt,name=message" json:"message,omitempty"`
}

func (m *DeviceConfigUpdatedMsg) Reset()                    { *m = DeviceConfigUpdatedMsg{} }
func (m *DeviceConfigUpdatedMsg) String() string            { return proto.CompactTextString(m) }
func (*DeviceConfigUpdatedMsg) ProtoMessage()               {}
func (*DeviceConfigUpdatedMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceConfigUpdatedMsg) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceConfigUpdatedMsg) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *DeviceConfigUpdatedMsg) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Event struct {
	Id        string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type      EventType            `protobuf:"varint,20,opt,name=type,enum=eventti.EventType" json:"type,omitempty"`
	Timestamp int64                `protobuf:"varint,30,opt,name=timestamp" json:"timestamp,omitempty"`
	Metadata  *google_protobuf.Any `protobuf:"bytes,40,opt,name=metadata" json:"metadata,omitempty"`
	Target    *Target              `protobuf:"bytes,50,opt,name=target" json:"target,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetMetadata() *google_protobuf.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Event) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type EmailServerCfg struct {
	Account  string        `protobuf:"bytes,10,opt,name=account" json:"account,omitempty"`
	Password string        `protobuf:"bytes,20,opt,name=password" json:"password,omitempty"`
	Host     string        `protobuf:"bytes,30,opt,name=host" json:"host,omitempty"`
	Port     int32         `protobuf:"varint,40,opt,name=port" json:"port,omitempty"`
	Protocol EmailProtocol `protobuf:"varint,50,opt,name=protocol,enum=eventti.EmailProtocol" json:"protocol,omitempty"`
	Timeout  int32         `protobuf:"varint,60,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *EmailServerCfg) Reset()                    { *m = EmailServerCfg{} }
func (m *EmailServerCfg) String() string            { return proto.CompactTextString(m) }
func (*EmailServerCfg) ProtoMessage()               {}
func (*EmailServerCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EmailServerCfg) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EmailServerCfg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EmailServerCfg) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *EmailServerCfg) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EmailServerCfg) GetProtocol() EmailProtocol {
	if m != nil {
		return m.Protocol
	}
	return EmailProtocol_UNKNOWN_PROTOCOL
}

func (m *EmailServerCfg) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type SmsCfg struct {
}

func (m *SmsCfg) Reset()                    { *m = SmsCfg{} }
func (m *SmsCfg) String() string            { return proto.CompactTextString(m) }
func (*SmsCfg) ProtoMessage()               {}
func (*SmsCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Target struct {
	Users []string `protobuf:"bytes,10,rep,name=users" json:"users,omitempty"`
	Roles []string `protobuf:"bytes,20,rep,name=roles" json:"roles,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Target) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Target) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type CreateNotificationChannelReq struct {
	Nc *NotificationChannel `protobuf:"bytes,10,opt,name=nc" json:"nc,omitempty"`
}

func (m *CreateNotificationChannelReq) Reset()                    { *m = CreateNotificationChannelReq{} }
func (m *CreateNotificationChannelReq) String() string            { return proto.CompactTextString(m) }
func (*CreateNotificationChannelReq) ProtoMessage()               {}
func (*CreateNotificationChannelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateNotificationChannelReq) GetNc() *NotificationChannel {
	if m != nil {
		return m.Nc
	}
	return nil
}

type NotificationChannel struct {
	Id     string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string                  `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
	Etypes []EventType             `protobuf:"varint,15,rep,packed,name=etypes,enum=eventti.EventType" json:"etypes,omitempty"`
	Type   NotificationChannelType `protobuf:"varint,20,opt,name=type,enum=eventti.NotificationChannelType" json:"type,omitempty"`
	Target *Target                 `protobuf:"bytes,30,opt,name=target" json:"target,omitempty"`
	// Types that are valid to be assigned to Cfg:
	//	*NotificationChannel_EmailServerCfg
	//	*NotificationChannel_SmsCfg
	Cfg isNotificationChannel_Cfg `protobuf_oneof:"cfg"`
}

func (m *NotificationChannel) Reset()                    { *m = NotificationChannel{} }
func (m *NotificationChannel) String() string            { return proto.CompactTextString(m) }
func (*NotificationChannel) ProtoMessage()               {}
func (*NotificationChannel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isNotificationChannel_Cfg interface {
	isNotificationChannel_Cfg()
}

type NotificationChannel_EmailServerCfg struct {
	EmailServerCfg *EmailServerCfg `protobuf:"bytes,40,opt,name=email_server_cfg,json=emailServerCfg,oneof"`
}
type NotificationChannel_SmsCfg struct {
	SmsCfg *SmsCfg `protobuf:"bytes,50,opt,name=sms_cfg,json=smsCfg,oneof"`
}

func (*NotificationChannel_EmailServerCfg) isNotificationChannel_Cfg() {}
func (*NotificationChannel_SmsCfg) isNotificationChannel_Cfg()         {}

func (m *NotificationChannel) GetCfg() isNotificationChannel_Cfg {
	if m != nil {
		return m.Cfg
	}
	return nil
}

func (m *NotificationChannel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NotificationChannel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotificationChannel) GetEtypes() []EventType {
	if m != nil {
		return m.Etypes
	}
	return nil
}

func (m *NotificationChannel) GetType() NotificationChannelType {
	if m != nil {
		return m.Type
	}
	return NotificationChannelType_UNKNOWN_NOTIFICATION_CHANNEL_TYPE
}

func (m *NotificationChannel) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *NotificationChannel) GetEmailServerCfg() *EmailServerCfg {
	if x, ok := m.GetCfg().(*NotificationChannel_EmailServerCfg); ok {
		return x.EmailServerCfg
	}
	return nil
}

func (m *NotificationChannel) GetSmsCfg() *SmsCfg {
	if x, ok := m.GetCfg().(*NotificationChannel_SmsCfg); ok {
		return x.SmsCfg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NotificationChannel) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NotificationChannel_OneofMarshaler, _NotificationChannel_OneofUnmarshaler, _NotificationChannel_OneofSizer, []interface{}{
		(*NotificationChannel_EmailServerCfg)(nil),
		(*NotificationChannel_SmsCfg)(nil),
	}
}

func _NotificationChannel_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NotificationChannel)
	// cfg
	switch x := m.Cfg.(type) {
	case *NotificationChannel_EmailServerCfg:
		b.EncodeVarint(40<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EmailServerCfg); err != nil {
			return err
		}
	case *NotificationChannel_SmsCfg:
		b.EncodeVarint(50<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SmsCfg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NotificationChannel.Cfg has unexpected type %T", x)
	}
	return nil
}

func _NotificationChannel_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NotificationChannel)
	switch tag {
	case 40: // cfg.email_server_cfg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EmailServerCfg)
		err := b.DecodeMessage(msg)
		m.Cfg = &NotificationChannel_EmailServerCfg{msg}
		return true, err
	case 50: // cfg.sms_cfg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SmsCfg)
		err := b.DecodeMessage(msg)
		m.Cfg = &NotificationChannel_SmsCfg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NotificationChannel_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NotificationChannel)
	// cfg
	switch x := m.Cfg.(type) {
	case *NotificationChannel_EmailServerCfg:
		s := proto.Size(x.EmailServerCfg)
		n += proto.SizeVarint(40<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NotificationChannel_SmsCfg:
		s := proto.Size(x.SmsCfg)
		n += proto.SizeVarint(50<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CreateNotificationChannelRes struct {
	Ret *common.BaseRet `protobuf:"bytes,10,opt,name=ret" json:"ret,omitempty"`
	Id  int32           `protobuf:"varint,20,opt,name=id" json:"id,omitempty"`
}

func (m *CreateNotificationChannelRes) Reset()                    { *m = CreateNotificationChannelRes{} }
func (m *CreateNotificationChannelRes) String() string            { return proto.CompactTextString(m) }
func (*CreateNotificationChannelRes) ProtoMessage()               {}
func (*CreateNotificationChannelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateNotificationChannelRes) GetRet() *common.BaseRet {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *CreateNotificationChannelRes) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNotificationChannelReq struct {
	Page *common.PageInfo `protobuf:"bytes,10,opt,name=page" json:"page,omitempty"`
}

func (m *GetNotificationChannelReq) Reset()                    { *m = GetNotificationChannelReq{} }
func (m *GetNotificationChannelReq) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationChannelReq) ProtoMessage()               {}
func (*GetNotificationChannelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetNotificationChannelReq) GetPage() *common.PageInfo {
	if m != nil {
		return m.Page
	}
	return nil
}

type GetNotificationChannelRes struct {
	Ret  *common.BaseRet        `protobuf:"bytes,10,opt,name=ret" json:"ret,omitempty"`
	Page *common.PageInfo       `protobuf:"bytes,20,opt,name=page" json:"page,omitempty"`
	Ncs  []*NotificationChannel `protobuf:"bytes,30,rep,name=ncs" json:"ncs,omitempty"`
}

func (m *GetNotificationChannelRes) Reset()                    { *m = GetNotificationChannelRes{} }
func (m *GetNotificationChannelRes) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationChannelRes) ProtoMessage()               {}
func (*GetNotificationChannelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetNotificationChannelRes) GetRet() *common.BaseRet {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *GetNotificationChannelRes) GetPage() *common.PageInfo {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *GetNotificationChannelRes) GetNcs() []*NotificationChannel {
	if m != nil {
		return m.Ncs
	}
	return nil
}

func init() {
	proto.RegisterType((*From)(nil), "eventti.From")
	proto.RegisterType((*DeviceConfigUpdatedMsg)(nil), "eventti.DeviceConfigUpdatedMsg")
	proto.RegisterType((*Event)(nil), "eventti.Event")
	proto.RegisterType((*EmailServerCfg)(nil), "eventti.EmailServerCfg")
	proto.RegisterType((*SmsCfg)(nil), "eventti.SmsCfg")
	proto.RegisterType((*Target)(nil), "eventti.Target")
	proto.RegisterType((*CreateNotificationChannelReq)(nil), "eventti.CreateNotificationChannelReq")
	proto.RegisterType((*NotificationChannel)(nil), "eventti.NotificationChannel")
	proto.RegisterType((*CreateNotificationChannelRes)(nil), "eventti.CreateNotificationChannelRes")
	proto.RegisterType((*GetNotificationChannelReq)(nil), "eventti.GetNotificationChannelReq")
	proto.RegisterType((*GetNotificationChannelRes)(nil), "eventti.GetNotificationChannelRes")
	proto.RegisterEnum("eventti.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("eventti.NotificationChannelType", NotificationChannelType_name, NotificationChannelType_value)
	proto.RegisterEnum("eventti.EmailProtocol", EmailProtocol_name, EmailProtocol_value)
}

func init() { proto.RegisterFile("eventti/eventti.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x53, 0xe3, 0x46,
	0x10, 0xc5, 0x9f, 0xd8, 0x4d, 0xad, 0x51, 0xcd, 0x7a, 0x59, 0x41, 0x58, 0xc7, 0xa8, 0xf2, 0x41,
	0x5c, 0x59, 0x3b, 0x71, 0xf6, 0x98, 0x43, 0x6c, 0x79, 0x00, 0x67, 0x8d, 0xe4, 0xc8, 0x62, 0x29,
	0x72, 0x71, 0x0d, 0xd2, 0x58, 0xa8, 0xd6, 0xd2, 0x28, 0x9a, 0x31, 0x29, 0xff, 0x8f, 0x9c, 0xf3,
	0x3b, 0x72, 0xd8, 0x1f, 0x97, 0xd2, 0x48, 0x32, 0x90, 0x00, 0xc9, 0x49, 0xdd, 0xef, 0xb5, 0x5b,
	0x3d, 0xfd, 0x9e, 0x47, 0xf0, 0x8a, 0xde, 0xd2, 0x50, 0x08, 0xbf, 0x97, 0x3d, 0xbb, 0x51, 0xcc,
	0x04, 0x43, 0xdb, 0x59, 0x7a, 0xb0, 0xef, 0x31, 0xe6, 0x2d, 0x69, 0x4f, 0xc2, 0xd7, 0xab, 0x45,
	0x8f, 0x84, 0xeb, 0xb4, 0xe6, 0xe0, 0xa5, 0xc3, 0x82, 0x80, 0x85, 0xbd, 0xf4, 0x91, 0x82, 0x5a,
	0x1f, 0xca, 0x27, 0x31, 0x0b, 0x10, 0x82, 0xf2, 0x0d, 0xe3, 0x42, 0x85, 0x76, 0xe1, 0xb8, 0x6e,
	0xc9, 0x18, 0xed, 0x41, 0x35, 0x60, 0xee, 0x6a, 0x49, 0xd5, 0xa6, 0x44, 0xb3, 0x4c, 0x8b, 0x60,
	0x6f, 0x44, 0x6f, 0x7d, 0x87, 0xea, 0x2c, 0x5c, 0xf8, 0xde, 0x45, 0xe4, 0x12, 0x41, 0xdd, 0x73,
	0xee, 0xa1, 0xcf, 0xa0, 0xee, 0x4a, 0x66, 0xee, 0xbb, 0x59, 0xab, 0x5a, 0x0a, 0x8c, 0x5d, 0xf4,
	0x39, 0xec, 0x64, 0x64, 0x48, 0x82, 0xbc, 0x27, 0xa4, 0x90, 0x41, 0x02, 0x8a, 0x54, 0xd8, 0x0e,
	0x28, 0xe7, 0xc4, 0xa3, 0x6a, 0x4b, 0x92, 0x79, 0xaa, 0xfd, 0x55, 0x80, 0x0a, 0x4e, 0x4e, 0x88,
	0x1a, 0x50, 0xf4, 0x5d, 0xb5, 0x20, 0xe9, 0xa2, 0xef, 0xa2, 0xaf, 0xa0, 0x2c, 0xd6, 0x51, 0xda,
	0xad, 0xd1, 0x47, 0xdd, 0x7c, 0x2d, 0xb2, 0xda, 0x5e, 0x47, 0xd4, 0x92, 0x3c, 0x3a, 0x84, 0xba,
	0xf0, 0x03, 0xca, 0x05, 0x09, 0x22, 0xd9, 0xbd, 0x64, 0xdd, 0x01, 0xe8, 0x3b, 0xa8, 0x05, 0x54,
	0x10, 0x97, 0x08, 0xa2, 0x1e, 0xb7, 0x0b, 0xc7, 0x3b, 0xfd, 0x66, 0x37, 0x5d, 0x64, 0x37, 0x5f,
	0x64, 0x77, 0x10, 0xae, 0xad, 0x4d, 0x15, 0xfa, 0x1a, 0xaa, 0x82, 0xc4, 0x1e, 0x15, 0x6a, 0x5f,
	0xd6, 0xef, 0x6e, 0xde, 0x6c, 0x4b, 0xd8, 0xca, 0x68, 0xed, 0x53, 0x01, 0x1a, 0x38, 0x20, 0xfe,
	0x72, 0x46, 0xe3, 0x5b, 0x1a, 0xeb, 0x0b, 0x2f, 0x39, 0x27, 0x71, 0x1c, 0xb6, 0x0a, 0xf3, 0x75,
	0xe7, 0x29, 0x3a, 0x80, 0x5a, 0x44, 0x38, 0xff, 0x9d, 0xc5, 0x6e, 0xb6, 0x9f, 0x4d, 0xbe, 0x51,
	0xa8, 0x75, 0x4f, 0x21, 0x04, 0xe5, 0x88, 0xc5, 0x42, 0xce, 0x5c, 0xb1, 0x64, 0x8c, 0xfa, 0x50,
	0x93, 0x33, 0x3b, 0x6c, 0x29, 0x67, 0x6b, 0xf4, 0xf7, 0xee, 0xb6, 0x92, 0x0c, 0x32, 0xcd, 0x58,
	0x6b, 0x53, 0x97, 0x4c, 0x94, 0x2c, 0x83, 0xad, 0x84, 0xfa, 0xa3, 0x6c, 0x95, 0xa7, 0x5a, 0x0d,
	0xaa, 0xb3, 0x80, 0xeb, 0x0b, 0x4f, 0x7b, 0x07, 0xd5, 0xf4, 0x68, 0xa8, 0x09, 0x95, 0x15, 0xa7,
	0x31, 0x57, 0xa1, 0x5d, 0x3a, 0xae, 0x5b, 0x69, 0x92, 0xa0, 0x31, 0x5b, 0x52, 0xae, 0x36, 0x53,
	0x54, 0x26, 0xda, 0x04, 0x0e, 0xf5, 0x98, 0x12, 0x41, 0x0d, 0x26, 0xfc, 0x85, 0xef, 0x10, 0xe1,
	0xb3, 0x50, 0xbf, 0x21, 0x61, 0x48, 0x97, 0x16, 0xfd, 0x0d, 0x7d, 0x0b, 0xc5, 0xd0, 0x91, 0x6b,
	0xd8, 0xe9, 0x1f, 0x6e, 0xe6, 0x7c, 0xac, 0xb8, 0x18, 0x3a, 0xda, 0xa7, 0x22, 0xbc, 0x7c, 0x84,
	0xfb, 0x97, 0x2b, 0x10, 0x94, 0xa5, 0xc7, 0x32, 0x37, 0x27, 0x31, 0xea, 0x40, 0x95, 0x26, 0x56,
	0xe0, 0xea, 0x6e, 0xbb, 0xf4, 0x84, 0x57, 0xb2, 0x0a, 0xf4, 0xee, 0x81, 0xab, 0xda, 0xcf, 0xcd,
	0x75, 0xcf, 0x63, 0x77, 0x9e, 0x68, 0x3d, 0xeb, 0x09, 0xa4, 0x83, 0x42, 0x13, 0x25, 0xe6, 0x5c,
	0x7a, 0x62, 0xee, 0x2c, 0xbc, 0xcc, 0x76, 0xaf, 0x1f, 0x4a, 0xb5, 0xf1, 0xcc, 0xd9, 0x96, 0xd5,
	0xa0, 0x0f, 0x5d, 0xd4, 0x81, 0x6d, 0x1e, 0x70, 0xf9, 0xdb, 0x7f, 0x5a, 0x30, 0x55, 0xec, 0x6c,
	0xcb, 0xaa, 0x72, 0x19, 0x0d, 0x2b, 0x50, 0x72, 0x16, 0x9e, 0xf6, 0xcb, 0xb3, 0x62, 0x70, 0x74,
	0x04, 0xa5, 0x98, 0x8a, 0x4c, 0x8d, 0xdd, 0x6e, 0x76, 0x51, 0x0c, 0x09, 0xa7, 0x16, 0x15, 0x56,
	0xc2, 0x65, 0x9b, 0x6e, 0x4a, 0x93, 0x14, 0x7d, 0x57, 0x1b, 0xc0, 0xfe, 0x29, 0x15, 0x4f, 0x88,
	0xfb, 0x05, 0x94, 0xa3, 0xe4, 0xdf, 0x9c, 0x36, 0x54, 0xf2, 0x86, 0x53, 0xe2, 0xd1, 0x71, 0xb8,
	0x60, 0x96, 0x64, 0xb5, 0x3f, 0x0a, 0x4f, 0xf7, 0xf8, 0x5f, 0x33, 0xe5, 0xaf, 0x69, 0x3e, 0xf7,
	0x1a, 0xd4, 0x85, 0x52, 0xe8, 0x70, 0xb5, 0xd5, 0x2e, 0xfd, 0xa7, 0xd5, 0x92, 0xc2, 0xce, 0x9f,
	0x05, 0xa8, 0x6f, 0x9c, 0x81, 0xf6, 0x00, 0x5d, 0x18, 0xef, 0x0d, 0xf3, 0xd2, 0x98, 0xe3, 0x0f,
	0xd8, 0xb0, 0xe7, 0xf6, 0xd5, 0x14, 0x2b, 0x5b, 0x68, 0x1f, 0x5e, 0x8d, 0xf0, 0x87, 0xb1, 0x8e,
	0xe7, 0xba, 0x69, 0x9c, 0x8c, 0x4f, 0xe7, 0x17, 0xd3, 0xd1, 0xc0, 0xc6, 0x23, 0xa5, 0x80, 0xde,
	0xc0, 0x7e, 0x46, 0x59, 0x83, 0xcb, 0x9c, 0x9e, 0x5d, 0x19, 0xfa, 0xd8, 0x38, 0x55, 0x8a, 0xe8,
	0x10, 0xd4, 0xc7, 0x69, 0x3c, 0x52, 0x4a, 0xe8, 0x08, 0xde, 0x64, 0xac, 0x39, 0xfc, 0x19, 0xeb,
	0x76, 0x5e, 0x30, 0x30, 0x06, 0x93, 0xab, 0x19, 0x1e, 0x29, 0xe5, 0x0e, 0x85, 0xd7, 0x4f, 0xf8,
	0x11, 0x7d, 0x09, 0x47, 0xf9, 0xb4, 0x86, 0x69, 0x8f, 0x4f, 0xc6, 0xfa, 0xc0, 0x1e, 0x9b, 0xc6,
	0x5c, 0x3f, 0x1b, 0x18, 0x06, 0x9e, 0xe4, 0xc3, 0xd7, 0xa1, 0x82, 0xcf, 0x07, 0xe3, 0x89, 0x02,
	0x49, 0x38, 0x3d, 0x33, 0x0d, 0xac, 0x34, 0xd1, 0x0b, 0xa8, 0x5f, 0xe2, 0xe1, 0xcc, 0xd4, 0xdf,
	0x63, 0x5b, 0x69, 0x75, 0x7e, 0x82, 0x17, 0x0f, 0xae, 0x0d, 0xd4, 0x04, 0x25, 0x6f, 0x3e, 0xb5,
	0x4c, 0xdb, 0xd4, 0xcd, 0x89, 0xb2, 0x85, 0x6a, 0x50, 0x9e, 0x9d, 0xdb, 0x53, 0x05, 0xd0, 0x2e,
	0xec, 0xcc, 0xb0, 0x7e, 0x61, 0xe1, 0xb9, 0x04, 0x9a, 0xc3, 0x53, 0x38, 0x08, 0xa9, 0xe8, 0xf2,
	0x8f, 0x6b, 0x67, 0xc9, 0x56, 0x6e, 0x7a, 0xa5, 0xe6, 0x02, 0xfc, 0xfa, 0x8d, 0xe7, 0x8b, 0x9b,
	0xd5, 0x75, 0xa2, 0x5a, 0x8f, 0x7f, 0x5c, 0xbf, 0x95, 0x35, 0x6f, 0x05, 0x75, 0xd2, 0x6f, 0x58,
	0xef, 0xf6, 0xfb, 0xfc, 0x5b, 0x77, 0x5d, 0x95, 0xc8, 0x0f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x8a, 0x64, 0xc7, 0x05, 0x07, 0x00, 0x00,
}
