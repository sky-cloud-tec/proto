// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: src/jrpc/irequest.proto

package jrpc

import (
	common "github.com/sky-cloud-tec/proto/v1/common"
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

// Standard response
type IResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     common.Retcode    `protobuf:"varint,10,opt,name=code,proto3,enum=common.Retcode" json:"code,omitempty"`
	Msg      string            `protobuf:"bytes,20,opt,name=msg,proto3" json:"msg,omitempty"`                                                                                                 // short msg like retcode but more readable for a glance
	Log      string            `protobuf:"bytes,23,opt,name=log,proto3" json:"log,omitempty"`                                                                                                 // error stack msg for debugging.
	Headers  map[string]string `protobuf:"bytes,25,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 复制/存储/转发请求时需要, 正常的http请求不需要
	PageInfo *common.PageInfo  `protobuf:"bytes,30,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	Body     *ResBody          `protobuf:"bytes,40,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *IResponse) Reset() {
	*x = IResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_irequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IResponse) ProtoMessage() {}

func (x *IResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_irequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IResponse.ProtoReflect.Descriptor instead.
func (*IResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_irequest_proto_rawDescGZIP(), []int{0}
}

func (x *IResponse) GetCode() common.Retcode {
	if x != nil {
		return x.Code
	}
	return common.Retcode_SUCCESS
}

func (x *IResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IResponse) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *IResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *IResponse) GetPageInfo() *common.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *IResponse) GetBody() *ResBody {
	if x != nil {
		return x.Body
	}
	return nil
}

// Standard request
type IRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string            `protobuf:"bytes,10,opt,name=url,proto3" json:"url,omitempty"`                                                                                                 // 复制/存储/转发请求时需要, 正常的http请求不需要
	Headers map[string]string `protobuf:"bytes,20,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 复制/存储/转发请求时需要, 正常的http请求不需要
	Body    *ReqBody          `protobuf:"bytes,30,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *IRequest) Reset() {
	*x = IRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_irequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRequest) ProtoMessage() {}

func (x *IRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_irequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRequest.ProtoReflect.Descriptor instead.
func (*IRequest) Descriptor() ([]byte, []int) {
	return file_src_jrpc_irequest_proto_rawDescGZIP(), []int{1}
}

func (x *IRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *IRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *IRequest) GetBody() *ReqBody {
	if x != nil {
		return x.Body
	}
	return nil
}

// 正常的 http 服务处理的是 ReqBody
type ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional request
	CreateNotificationChannelRequest *CreateNotificationChannelRequest `protobuf:"bytes,30,opt,name=create_notification_channel_request,json=createNotificationChannelRequest,proto3" json:"create_notification_channel_request,omitempty"`
	OperatorHotfixRequest            *OperatorHotfixRequest            `protobuf:"bytes,40,opt,name=operator_hotfix_request,json=operatorHotfixRequest,proto3" json:"operator_hotfix_request,omitempty"` // 如果没有找到对应的请求体，可能是 empty body 的情况
}

func (x *ReqBody) Reset() {
	*x = ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_irequest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqBody) ProtoMessage() {}

func (x *ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_irequest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqBody.ProtoReflect.Descriptor instead.
func (*ReqBody) Descriptor() ([]byte, []int) {
	return file_src_jrpc_irequest_proto_rawDescGZIP(), []int{2}
}

func (x *ReqBody) GetCreateNotificationChannelRequest() *CreateNotificationChannelRequest {
	if x != nil {
		return x.CreateNotificationChannelRequest
	}
	return nil
}

func (x *ReqBody) GetOperatorHotfixRequest() *OperatorHotfixRequest {
	if x != nil {
		return x.OperatorHotfixRequest
	}
	return nil
}

type ResBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional response
	CreateNotificationChannelResponse    *CreateNotificationChannelResponse    `protobuf:"bytes,10,opt,name=create_notification_channel_response,json=createNotificationChannelResponse,proto3" json:"create_notification_channel_response,omitempty"`
	StatBktRateAllTodayResponse          *StatBktRateAllTodayResponse          `protobuf:"bytes,20,opt,name=stat_bkt_rate_all_today_response,json=statBktRateAllTodayResponse,proto3" json:"stat_bkt_rate_all_today_response,omitempty"`                              // /api/stat/bkt/rate/all/today
	StatBktRateChannelTop10TodayResponse *StatBktRateChannelTop10TodayResponse `protobuf:"bytes,30,opt,name=stat_bkt_rate_channel_top10_today_response,json=statBktRateChannelTop10TodayResponse,proto3" json:"stat_bkt_rate_channel_top10_today_response,omitempty"` // /api/stat/bkt/rate/channel/top10/today
	StatBktVolAllTodayResponse           *StatBktVolAllTodayResponse           `protobuf:"bytes,40,opt,name=stat_bkt_vol_all_today_response,json=statBktVolAllTodayResponse,proto3" json:"stat_bkt_vol_all_today_response,omitempty"`                                 // /api/stat/bkt/vol/all/today
	StatBktVolChannelTop10TodayResponse  *StatBktVolChannelTop10TodayResponse  `protobuf:"bytes,50,opt,name=stat_bkt_vol_channel_top10_today_response,json=statBktVolChannelTop10TodayResponse,proto3" json:"stat_bkt_vol_channel_top10_today_response,omitempty"`    // /api/stat/bkt/vol/channel/top10/today
	StatBktLapseTodayResponse            *StatBktLapseTodayResponse            `protobuf:"bytes,60,opt,name=stat_bkt_lapse_today_response,json=statBktLapseTodayResponse,proto3" json:"stat_bkt_lapse_today_response,omitempty"`                                      // /api/stat/bkt/lapse/today
	StatBktRecordResponse                *StatBktRecordResponse                `protobuf:"bytes,70,opt,name=stat_bkt_record_response,json=statBktRecordResponse,proto3" json:"stat_bkt_record_response,omitempty"`                                                    // /api/stat/bkt/record?page=1&size=9
	StatBktNodeResponse                  *StatBktNodeResponse                  `protobuf:"bytes,80,opt,name=stat_bkt_node_response,json=statBktNodeResponse,proto3" json:"stat_bkt_node_response,omitempty"`                                                          // /api/stat/bkt/node?name=xxx
	GetNotificationChannelResponse       *GetNotificationChannelResponse       `protobuf:"bytes,90,opt,name=get_notification_channel_response,json=getNotificationChannelResponse,proto3" json:"get_notification_channel_response,omitempty"`
	StatBktActiveClientsResponse         *StatBktActiveClientsResponse         `protobuf:"bytes,100,opt,name=stat_bkt_active_clients_response,json=statBktActiveClientsResponse,proto3" json:"stat_bkt_active_clients_response,omitempty"` // /api/stat/bkt/active_clients
	OperatorHotfixResponse               *OperatorHotfixResponse               `protobuf:"bytes,110,opt,name=operator_hotfix_response,json=operatorHotfixResponse,proto3" json:"operator_hotfix_response,omitempty"`
}

func (x *ResBody) Reset() {
	*x = ResBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_irequest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResBody) ProtoMessage() {}

func (x *ResBody) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_irequest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResBody.ProtoReflect.Descriptor instead.
func (*ResBody) Descriptor() ([]byte, []int) {
	return file_src_jrpc_irequest_proto_rawDescGZIP(), []int{3}
}

func (x *ResBody) GetCreateNotificationChannelResponse() *CreateNotificationChannelResponse {
	if x != nil {
		return x.CreateNotificationChannelResponse
	}
	return nil
}

func (x *ResBody) GetStatBktRateAllTodayResponse() *StatBktRateAllTodayResponse {
	if x != nil {
		return x.StatBktRateAllTodayResponse
	}
	return nil
}

func (x *ResBody) GetStatBktRateChannelTop10TodayResponse() *StatBktRateChannelTop10TodayResponse {
	if x != nil {
		return x.StatBktRateChannelTop10TodayResponse
	}
	return nil
}

func (x *ResBody) GetStatBktVolAllTodayResponse() *StatBktVolAllTodayResponse {
	if x != nil {
		return x.StatBktVolAllTodayResponse
	}
	return nil
}

func (x *ResBody) GetStatBktVolChannelTop10TodayResponse() *StatBktVolChannelTop10TodayResponse {
	if x != nil {
		return x.StatBktVolChannelTop10TodayResponse
	}
	return nil
}

func (x *ResBody) GetStatBktLapseTodayResponse() *StatBktLapseTodayResponse {
	if x != nil {
		return x.StatBktLapseTodayResponse
	}
	return nil
}

func (x *ResBody) GetStatBktRecordResponse() *StatBktRecordResponse {
	if x != nil {
		return x.StatBktRecordResponse
	}
	return nil
}

func (x *ResBody) GetStatBktNodeResponse() *StatBktNodeResponse {
	if x != nil {
		return x.StatBktNodeResponse
	}
	return nil
}

func (x *ResBody) GetGetNotificationChannelResponse() *GetNotificationChannelResponse {
	if x != nil {
		return x.GetNotificationChannelResponse
	}
	return nil
}

func (x *ResBody) GetStatBktActiveClientsResponse() *StatBktActiveClientsResponse {
	if x != nil {
		return x.StatBktActiveClientsResponse
	}
	return nil
}

func (x *ResBody) GetOperatorHotfixResponse() *OperatorHotfixResponse {
	if x != nil {
		return x.OperatorHotfixResponse
	}
	return nil
}

var File_src_jrpc_irequest_proto protoreflect.FileDescriptor

var file_src_jrpc_irequest_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6a, 0x72, 0x70, 0x63, 0x1a,
	0x1c, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73,
	0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x72, 0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x72,
	0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x09, 0x49, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x36,
	0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xb2, 0x01, 0x0a, 0x08, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd5, 0x01, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x75, 0x0a, 0x23, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x17,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x6f, 0x74, 0x66, 0x69, 0x78, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6a, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x6f, 0x74,
	0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x15, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x48, 0x6f, 0x74, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x9d, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x78, 0x0a,
	0x24, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6a, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x21, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x20, 0x73, 0x74, 0x61, 0x74, 0x5f,
	0x62, 0x6b, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x64,
	0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x1b, 0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x84, 0x01, 0x0a, 0x2a, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6b, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x6f, 0x70, 0x31,
	0x30, 0x5f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x6f, 0x70, 0x31, 0x30, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x24, 0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x6f, 0x70, 0x31, 0x30, 0x54, 0x6f, 0x64, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x1f, 0x73, 0x74, 0x61, 0x74,
	0x5f, 0x62, 0x6b, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x64,
	0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74,
	0x56, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x1a, 0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x56, 0x6f, 0x6c, 0x41,
	0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x81, 0x01, 0x0a, 0x29, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6b, 0x74, 0x5f, 0x76, 0x6f, 0x6c,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x6f, 0x70, 0x31, 0x30, 0x5f, 0x74,
	0x6f, 0x64, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42,
	0x6b, 0x74, 0x56, 0x6f, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x6f, 0x70, 0x31,
	0x30, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x23,
	0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x56, 0x6f, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x6f, 0x70, 0x31, 0x30, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x1d, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6b, 0x74, 0x5f,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x54, 0x6f,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x19, 0x73, 0x74, 0x61,
	0x74, 0x42, 0x6b, 0x74, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x18, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62,
	0x6b, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x15, 0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x16,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6b, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x73, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x21,
	0x67, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x1e, 0x67,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6b, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x1c, 0x73, 0x74, 0x61,
	0x74, 0x42, 0x6b, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x18, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x6f, 0x74, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x72,
	0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x6f, 0x74, 0x66, 0x69,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x48, 0x6f, 0x74, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x41, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_jrpc_irequest_proto_rawDescOnce sync.Once
	file_src_jrpc_irequest_proto_rawDescData = file_src_jrpc_irequest_proto_rawDesc
)

func file_src_jrpc_irequest_proto_rawDescGZIP() []byte {
	file_src_jrpc_irequest_proto_rawDescOnce.Do(func() {
		file_src_jrpc_irequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_jrpc_irequest_proto_rawDescData)
	})
	return file_src_jrpc_irequest_proto_rawDescData
}

var file_src_jrpc_irequest_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_jrpc_irequest_proto_goTypes = []interface{}{
	(*IResponse)(nil),                        // 0: jrpc.IResponse
	(*IRequest)(nil),                         // 1: jrpc.IRequest
	(*ReqBody)(nil),                          // 2: jrpc.ReqBody
	(*ResBody)(nil),                          // 3: jrpc.ResBody
	nil,                                      // 4: jrpc.IResponse.HeadersEntry
	nil,                                      // 5: jrpc.IRequest.HeadersEntry
	(common.Retcode)(0),                      // 6: common.Retcode
	(*common.PageInfo)(nil),                  // 7: common.PageInfo
	(*CreateNotificationChannelRequest)(nil), // 8: jrpc.CreateNotificationChannelRequest
	(*OperatorHotfixRequest)(nil),            // 9: jrpc.OperatorHotfixRequest
	(*CreateNotificationChannelResponse)(nil),    // 10: jrpc.CreateNotificationChannelResponse
	(*StatBktRateAllTodayResponse)(nil),          // 11: jrpc.StatBktRateAllTodayResponse
	(*StatBktRateChannelTop10TodayResponse)(nil), // 12: jrpc.StatBktRateChannelTop10TodayResponse
	(*StatBktVolAllTodayResponse)(nil),           // 13: jrpc.StatBktVolAllTodayResponse
	(*StatBktVolChannelTop10TodayResponse)(nil),  // 14: jrpc.StatBktVolChannelTop10TodayResponse
	(*StatBktLapseTodayResponse)(nil),            // 15: jrpc.StatBktLapseTodayResponse
	(*StatBktRecordResponse)(nil),                // 16: jrpc.StatBktRecordResponse
	(*StatBktNodeResponse)(nil),                  // 17: jrpc.StatBktNodeResponse
	(*GetNotificationChannelResponse)(nil),       // 18: jrpc.GetNotificationChannelResponse
	(*StatBktActiveClientsResponse)(nil),         // 19: jrpc.StatBktActiveClientsResponse
	(*OperatorHotfixResponse)(nil),               // 20: jrpc.OperatorHotfixResponse
}
var file_src_jrpc_irequest_proto_depIdxs = []int32{
	6,  // 0: jrpc.IResponse.code:type_name -> common.Retcode
	4,  // 1: jrpc.IResponse.headers:type_name -> jrpc.IResponse.HeadersEntry
	7,  // 2: jrpc.IResponse.page_info:type_name -> common.PageInfo
	3,  // 3: jrpc.IResponse.body:type_name -> jrpc.ResBody
	5,  // 4: jrpc.IRequest.headers:type_name -> jrpc.IRequest.HeadersEntry
	2,  // 5: jrpc.IRequest.body:type_name -> jrpc.ReqBody
	8,  // 6: jrpc.ReqBody.create_notification_channel_request:type_name -> jrpc.CreateNotificationChannelRequest
	9,  // 7: jrpc.ReqBody.operator_hotfix_request:type_name -> jrpc.OperatorHotfixRequest
	10, // 8: jrpc.ResBody.create_notification_channel_response:type_name -> jrpc.CreateNotificationChannelResponse
	11, // 9: jrpc.ResBody.stat_bkt_rate_all_today_response:type_name -> jrpc.StatBktRateAllTodayResponse
	12, // 10: jrpc.ResBody.stat_bkt_rate_channel_top10_today_response:type_name -> jrpc.StatBktRateChannelTop10TodayResponse
	13, // 11: jrpc.ResBody.stat_bkt_vol_all_today_response:type_name -> jrpc.StatBktVolAllTodayResponse
	14, // 12: jrpc.ResBody.stat_bkt_vol_channel_top10_today_response:type_name -> jrpc.StatBktVolChannelTop10TodayResponse
	15, // 13: jrpc.ResBody.stat_bkt_lapse_today_response:type_name -> jrpc.StatBktLapseTodayResponse
	16, // 14: jrpc.ResBody.stat_bkt_record_response:type_name -> jrpc.StatBktRecordResponse
	17, // 15: jrpc.ResBody.stat_bkt_node_response:type_name -> jrpc.StatBktNodeResponse
	18, // 16: jrpc.ResBody.get_notification_channel_response:type_name -> jrpc.GetNotificationChannelResponse
	19, // 17: jrpc.ResBody.stat_bkt_active_clients_response:type_name -> jrpc.StatBktActiveClientsResponse
	20, // 18: jrpc.ResBody.operator_hotfix_response:type_name -> jrpc.OperatorHotfixResponse
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_src_jrpc_irequest_proto_init() }
func file_src_jrpc_irequest_proto_init() {
	if File_src_jrpc_irequest_proto != nil {
		return
	}
	file_src_jrpc_eventti_proto_init()
	file_src_jrpc_netai_proto_init()
	file_src_jrpc_netd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_jrpc_irequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IResponse); i {
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
		file_src_jrpc_irequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRequest); i {
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
		file_src_jrpc_irequest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqBody); i {
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
		file_src_jrpc_irequest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResBody); i {
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
			RawDescriptor: file_src_jrpc_irequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_jrpc_irequest_proto_goTypes,
		DependencyIndexes: file_src_jrpc_irequest_proto_depIdxs,
		MessageInfos:      file_src_jrpc_irequest_proto_msgTypes,
	}.Build()
	File_src_jrpc_irequest_proto = out.File
	file_src_jrpc_irequest_proto_rawDesc = nil
	file_src_jrpc_irequest_proto_goTypes = nil
	file_src_jrpc_irequest_proto_depIdxs = nil
}
