// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/grpc/eventti.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/sky-cloud-tec/proto/v1/common"
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

type EventItem struct {
	User                 string   `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,20,opt,name=message,proto3" json:"message,omitempty"`
	Time                 string   `protobuf:"bytes,30,opt,name=time,proto3" json:"time,omitempty"`
	Type                 string   `protobuf:"bytes,40,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventItem) Reset()         { *m = EventItem{} }
func (m *EventItem) String() string { return proto.CompactTextString(m) }
func (*EventItem) ProtoMessage()    {}
func (*EventItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20a27e87a146f3e, []int{0}
}

func (m *EventItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventItem.Unmarshal(m, b)
}
func (m *EventItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventItem.Marshal(b, m, deterministic)
}
func (m *EventItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventItem.Merge(m, src)
}
func (m *EventItem) XXX_Size() int {
	return xxx_messageInfo_EventItem.Size(m)
}
func (m *EventItem) XXX_DiscardUnknown() {
	xxx_messageInfo_EventItem.DiscardUnknown(m)
}

var xxx_messageInfo_EventItem proto.InternalMessageInfo

func (m *EventItem) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *EventItem) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EventItem) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *EventItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ByEventFieldsAndPageRequest struct {
	Page                 int32    `protobuf:"varint,10,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,20,opt,name=size,proto3" json:"size,omitempty"`
	User                 string   `protobuf:"bytes,30,opt,name=user,proto3" json:"user,omitempty"`
	Type                 string   `protobuf:"bytes,40,opt,name=type,proto3" json:"type,omitempty"`
	Message              string   `protobuf:"bytes,50,opt,name=message,proto3" json:"message,omitempty"`
	Start                int64    `protobuf:"varint,60,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,70,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByEventFieldsAndPageRequest) Reset()         { *m = ByEventFieldsAndPageRequest{} }
func (m *ByEventFieldsAndPageRequest) String() string { return proto.CompactTextString(m) }
func (*ByEventFieldsAndPageRequest) ProtoMessage()    {}
func (*ByEventFieldsAndPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20a27e87a146f3e, []int{1}
}

func (m *ByEventFieldsAndPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByEventFieldsAndPageRequest.Unmarshal(m, b)
}
func (m *ByEventFieldsAndPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByEventFieldsAndPageRequest.Marshal(b, m, deterministic)
}
func (m *ByEventFieldsAndPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByEventFieldsAndPageRequest.Merge(m, src)
}
func (m *ByEventFieldsAndPageRequest) XXX_Size() int {
	return xxx_messageInfo_ByEventFieldsAndPageRequest.Size(m)
}
func (m *ByEventFieldsAndPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByEventFieldsAndPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByEventFieldsAndPageRequest proto.InternalMessageInfo

func (m *ByEventFieldsAndPageRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ByEventFieldsAndPageRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ByEventFieldsAndPageRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ByEventFieldsAndPageRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ByEventFieldsAndPageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ByEventFieldsAndPageRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ByEventFieldsAndPageRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type SearchEventResponse struct {
	PageInfo             *common.PageInfo `protobuf:"bytes,10,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	Items                []*EventItem     `protobuf:"bytes,20,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchEventResponse) Reset()         { *m = SearchEventResponse{} }
func (m *SearchEventResponse) String() string { return proto.CompactTextString(m) }
func (*SearchEventResponse) ProtoMessage()    {}
func (*SearchEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20a27e87a146f3e, []int{2}
}

func (m *SearchEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchEventResponse.Unmarshal(m, b)
}
func (m *SearchEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchEventResponse.Marshal(b, m, deterministic)
}
func (m *SearchEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchEventResponse.Merge(m, src)
}
func (m *SearchEventResponse) XXX_Size() int {
	return xxx_messageInfo_SearchEventResponse.Size(m)
}
func (m *SearchEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchEventResponse proto.InternalMessageInfo

func (m *SearchEventResponse) GetPageInfo() *common.PageInfo {
	if m != nil {
		return m.PageInfo
	}
	return nil
}

func (m *SearchEventResponse) GetItems() []*EventItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EventItem)(nil), "rpc.EventItem")
	proto.RegisterType((*ByEventFieldsAndPageRequest)(nil), "rpc.ByEventFieldsAndPageRequest")
	proto.RegisterType((*SearchEventResponse)(nil), "rpc.SearchEventResponse")
}

func init() { proto.RegisterFile("src/grpc/eventti.proto", fileDescriptor_e20a27e87a146f3e) }

var fileDescriptor_e20a27e87a146f3e = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0xcb, 0xd3, 0x30,
	0x14, 0xc5, 0x1d, 0xb5, 0xea, 0x32, 0x91, 0x8f, 0x38, 0x3f, 0xc2, 0x84, 0x51, 0x8a, 0x42, 0x5f,
	0xd6, 0x62, 0x7d, 0x15, 0xc4, 0x81, 0x83, 0xe1, 0x8b, 0x64, 0x6f, 0xfa, 0x20, 0x5d, 0x7a, 0xd7,
	0xc5, 0xad, 0x49, 0x4c, 0xd2, 0x41, 0xfd, 0xa7, 0xfc, 0x17, 0x25, 0xb7, 0x73, 0x4c, 0x18, 0xdf,
	0xdb, 0xb9, 0xe7, 0x9e, 0x70, 0x7e, 0xbd, 0x25, 0xf7, 0xce, 0x8a, 0xa2, 0xb1, 0x46, 0x14, 0x70,
	0x02, 0xe5, 0xbd, 0xcc, 0x8d, 0xd5, 0x5e, 0xd3, 0xc8, 0x1a, 0x31, 0x7b, 0x15, 0x96, 0x42, 0xb7,
	0xad, 0x56, 0x85, 0xa9, 0x1a, 0x18, 0x76, 0x69, 0x45, 0xc6, 0x9f, 0x43, 0x78, 0xed, 0xa1, 0xa5,
	0x94, 0x3c, 0xee, 0x1c, 0x58, 0x46, 0x92, 0x51, 0x36, 0xe6, 0xa8, 0x29, 0x23, 0x4f, 0x5b, 0x70,
	0xae, 0x6a, 0x80, 0x4d, 0xd1, 0xfe, 0x37, 0x86, 0xb4, 0x97, 0x2d, 0xb0, 0xf9, 0x90, 0x0e, 0x1a,
	0xbd, 0xde, 0x00, 0xcb, 0xce, 0x5e, 0x6f, 0x20, 0xfd, 0x33, 0x22, 0xaf, 0x97, 0x3d, 0xb6, 0xac,
	0x24, 0x1c, 0x6b, 0xf7, 0x49, 0xd5, 0x5f, 0xab, 0x06, 0x38, 0xfc, 0xea, 0xc0, 0xf9, 0xf0, 0x26,
	0x00, 0x61, 0x6b, 0xcc, 0x51, 0x07, 0xcf, 0xc9, 0xdf, 0x43, 0x65, 0xcc, 0x51, 0x5f, 0xe8, 0xe6,
	0x57, 0x74, 0x37, 0xfa, 0xae, 0x89, 0xcb, 0xff, 0x89, 0xa7, 0x24, 0x76, 0xbe, 0xb2, 0x9e, 0x7d,
	0x48, 0x46, 0x59, 0xc4, 0x87, 0x81, 0xde, 0x91, 0x08, 0x54, 0xcd, 0x56, 0xe8, 0x05, 0x99, 0xfe,
	0x24, 0x2f, 0x37, 0x50, 0x59, 0xb1, 0x47, 0x68, 0x0e, 0xce, 0x68, 0xe5, 0x80, 0x2e, 0xc8, 0x38,
	0xc0, 0xfd, 0x90, 0x6a, 0xa7, 0x91, 0x76, 0x52, 0xde, 0xe5, 0xc3, 0x49, 0xf3, 0xf0, 0x41, 0x6b,
	0xb5, 0xd3, 0xfc, 0x99, 0x39, 0x2b, 0xfa, 0x86, 0xc4, 0xd2, 0x43, 0xeb, 0xd8, 0x34, 0x89, 0xb2,
	0x49, 0xf9, 0x22, 0xb7, 0x46, 0xe4, 0x97, 0x63, 0xf3, 0x61, 0x59, 0x7e, 0x27, 0xcf, 0xd1, 0xdb,
	0x80, 0x3d, 0x49, 0x01, 0xf4, 0x0b, 0x99, 0x5c, 0x75, 0xd3, 0x04, 0x5f, 0x3d, 0x70, 0xbe, 0x19,
	0xc3, 0xc4, 0x0d, 0xde, 0xf4, 0xd1, 0xf2, 0x23, 0xb9, 0x57, 0xe0, 0x73, 0x77, 0xe8, 0xc5, 0x51,
	0x77, 0xf5, 0xf0, 0xcf, 0x43, 0xfe, 0xdb, 0xdb, 0x46, 0xfa, 0x7d, 0xb7, 0x0d, 0xf8, 0x85, 0x3b,
	0xf4, 0x0b, 0xdc, 0x2f, 0x3c, 0x88, 0x02, 0x33, 0xc5, 0xe9, 0x5d, 0x61, 0x8d, 0xd8, 0x3e, 0xc1,
	0xe9, 0xfd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x72, 0x90, 0x32, 0x5b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	SearchEvent(ctx context.Context, in *ByEventFieldsAndPageRequest, opts ...grpc.CallOption) (*SearchEventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) SearchEvent(ctx context.Context, in *ByEventFieldsAndPageRequest, opts ...grpc.CallOption) (*SearchEventResponse, error) {
	out := new(SearchEventResponse)
	err := c.cc.Invoke(ctx, "/rpc.EventService/SearchEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	SearchEvent(context.Context, *ByEventFieldsAndPageRequest) (*SearchEventResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) SearchEvent(ctx context.Context, req *ByEventFieldsAndPageRequest) (*SearchEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEvent not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_SearchEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEventFieldsAndPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SearchEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.EventService/SearchEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SearchEvent(ctx, req.(*ByEventFieldsAndPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchEvent",
			Handler:    _EventService_SearchEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/eventti.proto",
}