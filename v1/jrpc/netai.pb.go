// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: src/jrpc/netai.proto

package jrpc

import (
	netai "github.com/sky-cloud-tec/proto/v1/netai"
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

type StatBktRateAllTodayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int32   `protobuf:"varint,10,opt,name=success,proto3" json:"success,omitempty"` // 成功数
	Fail    int32   `protobuf:"varint,20,opt,name=fail,proto3" json:"fail,omitempty"`       // 失败数
	Rate    float32 `protobuf:"fixed32,30,opt,name=rate,proto3" json:"rate,omitempty"`      // 成功率
}

func (x *StatBktRateAllTodayResponse) Reset() {
	*x = StatBktRateAllTodayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktRateAllTodayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktRateAllTodayResponse) ProtoMessage() {}

func (x *StatBktRateAllTodayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktRateAllTodayResponse.ProtoReflect.Descriptor instead.
func (*StatBktRateAllTodayResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{0}
}

func (x *StatBktRateAllTodayResponse) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *StatBktRateAllTodayResponse) GetFail() int32 {
	if x != nil {
		return x.Fail
	}
	return 0
}

func (x *StatBktRateAllTodayResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type ChannelRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Rate float32 `protobuf:"fixed32,20,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *ChannelRate) Reset() {
	*x = ChannelRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelRate) ProtoMessage() {}

func (x *ChannelRate) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelRate.ProtoReflect.Descriptor instead.
func (*ChannelRate) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelRate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelRate) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type StatBktRateChannelTop10TodayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelRates []*ChannelRate `protobuf:"bytes,10,rep,name=channel_rates,json=channelRates,proto3" json:"channel_rates,omitempty"`
}

func (x *StatBktRateChannelTop10TodayResponse) Reset() {
	*x = StatBktRateChannelTop10TodayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktRateChannelTop10TodayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktRateChannelTop10TodayResponse) ProtoMessage() {}

func (x *StatBktRateChannelTop10TodayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktRateChannelTop10TodayResponse.ProtoReflect.Descriptor instead.
func (*StatBktRateChannelTop10TodayResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{2}
}

func (x *StatBktRateChannelTop10TodayResponse) GetChannelRates() []*ChannelRate {
	if x != nil {
		return x.ChannelRates
	}
	return nil
}

type StatBktVolAllTodayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cnt int32   `protobuf:"varint,10,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Vol float32 `protobuf:"fixed32,20,opt,name=vol,proto3" json:"vol,omitempty"`
}

func (x *StatBktVolAllTodayResponse) Reset() {
	*x = StatBktVolAllTodayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktVolAllTodayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktVolAllTodayResponse) ProtoMessage() {}

func (x *StatBktVolAllTodayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktVolAllTodayResponse.ProtoReflect.Descriptor instead.
func (*StatBktVolAllTodayResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{3}
}

func (x *StatBktVolAllTodayResponse) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *StatBktVolAllTodayResponse) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

type ChannelVol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Cnt  int32   `protobuf:"varint,20,opt,name=cnt,proto3" json:"cnt,omitempty"`  // 交易量
	Vol  float32 `protobuf:"fixed32,30,opt,name=vol,proto3" json:"vol,omitempty"` // 交易额
}

func (x *ChannelVol) Reset() {
	*x = ChannelVol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelVol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelVol) ProtoMessage() {}

func (x *ChannelVol) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelVol.ProtoReflect.Descriptor instead.
func (*ChannelVol) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelVol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelVol) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *ChannelVol) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

type StatBktVolChannelTop10TodayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByCnt []*ChannelVol `protobuf:"bytes,10,rep,name=by_cnt,json=byCnt,proto3" json:"by_cnt,omitempty"`
	ByVol []*ChannelVol `protobuf:"bytes,20,rep,name=by_vol,json=byVol,proto3" json:"by_vol,omitempty"`
}

func (x *StatBktVolChannelTop10TodayResponse) Reset() {
	*x = StatBktVolChannelTop10TodayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktVolChannelTop10TodayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktVolChannelTop10TodayResponse) ProtoMessage() {}

func (x *StatBktVolChannelTop10TodayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktVolChannelTop10TodayResponse.ProtoReflect.Descriptor instead.
func (*StatBktVolChannelTop10TodayResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{5}
}

func (x *StatBktVolChannelTop10TodayResponse) GetByCnt() []*ChannelVol {
	if x != nil {
		return x.ByCnt
	}
	return nil
}

func (x *StatBktVolChannelTop10TodayResponse) GetByVol() []*ChannelVol {
	if x != nil {
		return x.ByVol
	}
	return nil
}

type StatBktLapseTodayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avg float32 `protobuf:"fixed32,10,opt,name=avg,proto3" json:"avg,omitempty"` // 平均延时 nano secs 调用者自己处理精度
	Max int32   `protobuf:"varint,20,opt,name=max,proto3" json:"max,omitempty"`  // nano secs 调用者自己处理精度
	Min int32   `protobuf:"varint,30,opt,name=min,proto3" json:"min,omitempty"`  // nano secs
}

func (x *StatBktLapseTodayResponse) Reset() {
	*x = StatBktLapseTodayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktLapseTodayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktLapseTodayResponse) ProtoMessage() {}

func (x *StatBktLapseTodayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktLapseTodayResponse.ProtoReflect.Descriptor instead.
func (*StatBktLapseTodayResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{6}
}

func (x *StatBktLapseTodayResponse) GetAvg() float32 {
	if x != nil {
		return x.Avg
	}
	return 0
}

func (x *StatBktLapseTodayResponse) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *StatBktLapseTodayResponse) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

type BktRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       string  `protobuf:"bytes,10,opt,name=ts,proto3" json:"ts,omitempty"`
	FlowNo   string  `protobuf:"bytes,20,opt,name=flow_no,json=flowNo,proto3" json:"flow_no,omitempty"`
	Channel  string  `protobuf:"bytes,30,opt,name=channel,proto3" json:"channel,omitempty"`
	Trade    string  `protobuf:"bytes,40,opt,name=trade,proto3" json:"trade,omitempty"`
	Cnt      float32 `protobuf:"fixed32,50,opt,name=cnt,proto3" json:"cnt,omitempty"`
	ClientId int32   `protobuf:"varint,60,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Branch   string  `protobuf:"bytes,70,opt,name=branch,proto3" json:"branch,omitempty"`
	CreateAt int64   `protobuf:"varint,80,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	Src      string  `protobuf:"bytes,81,opt,name=src,proto3" json:"src,omitempty"`
	Dst      string  `protobuf:"bytes,82,opt,name=dst,proto3" json:"dst,omitempty"`
	Code     int32   `protobuf:"varint,83,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string  `protobuf:"bytes,84,opt,name=msg,proto3" json:"msg,omitempty"`
	In       int64   `protobuf:"varint,90,opt,name=in,proto3" json:"in,omitempty"`        // unix nano
	Out      int64   `protobuf:"varint,100,opt,name=out,proto3" json:"out,omitempty"`     // unix nano
	Lapse    int64   `protobuf:"varint,110,opt,name=lapse,proto3" json:"lapse,omitempty"` // nano secs
}

func (x *BktRecord) Reset() {
	*x = BktRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BktRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BktRecord) ProtoMessage() {}

func (x *BktRecord) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BktRecord.ProtoReflect.Descriptor instead.
func (*BktRecord) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{7}
}

func (x *BktRecord) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *BktRecord) GetFlowNo() string {
	if x != nil {
		return x.FlowNo
	}
	return ""
}

func (x *BktRecord) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *BktRecord) GetTrade() string {
	if x != nil {
		return x.Trade
	}
	return ""
}

func (x *BktRecord) GetCnt() float32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *BktRecord) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *BktRecord) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BktRecord) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *BktRecord) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *BktRecord) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *BktRecord) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BktRecord) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BktRecord) GetIn() int64 {
	if x != nil {
		return x.In
	}
	return 0
}

func (x *BktRecord) GetOut() int64 {
	if x != nil {
		return x.Out
	}
	return 0
}

func (x *BktRecord) GetLapse() int64 {
	if x != nil {
		return x.Lapse
	}
	return 0
}

type StatBktRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*BktRecord `protobuf:"bytes,10,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *StatBktRecordResponse) Reset() {
	*x = StatBktRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktRecordResponse) ProtoMessage() {}

func (x *StatBktRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktRecordResponse.ProtoReflect.Descriptor instead.
func (*StatBktRecordResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{8}
}

func (x *StatBktRecordResponse) GetRecords() []*BktRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type StatBktNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traffic netai.TrafficLoad `protobuf:"varint,5,opt,name=traffic,proto3,enum=netai.TrafficLoad" json:"traffic,omitempty"`
	Avg     float32           `protobuf:"fixed32,10,opt,name=avg,proto3" json:"avg,omitempty"` // 平均延时 nano secs 调用者自己处理精度
	Max     int32             `protobuf:"varint,20,opt,name=max,proto3" json:"max,omitempty"`  // nano secs 调用者自己处理精度
	Min     int32             `protobuf:"varint,30,opt,name=min,proto3" json:"min,omitempty"`  // nano secs
	Cnt     int32             `protobuf:"varint,40,opt,name=cnt,proto3" json:"cnt,omitempty"`  // 交易量
	Vol     float32           `protobuf:"fixed32,50,opt,name=vol,proto3" json:"vol,omitempty"` // 交易额
}

func (x *StatBktNodeResponse) Reset() {
	*x = StatBktNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktNodeResponse) ProtoMessage() {}

func (x *StatBktNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktNodeResponse.ProtoReflect.Descriptor instead.
func (*StatBktNodeResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{9}
}

func (x *StatBktNodeResponse) GetTraffic() netai.TrafficLoad {
	if x != nil {
		return x.Traffic
	}
	return netai.TrafficLoad_HIGH
}

func (x *StatBktNodeResponse) GetAvg() float32 {
	if x != nil {
		return x.Avg
	}
	return 0
}

func (x *StatBktNodeResponse) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *StatBktNodeResponse) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *StatBktNodeResponse) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *StatBktNodeResponse) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

type StatBktActiveClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cnt int32 `protobuf:"varint,10,opt,name=cnt,proto3" json:"cnt,omitempty"`
}

func (x *StatBktActiveClientsResponse) Reset() {
	*x = StatBktActiveClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_jrpc_netai_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBktActiveClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBktActiveClientsResponse) ProtoMessage() {}

func (x *StatBktActiveClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_jrpc_netai_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBktActiveClientsResponse.ProtoReflect.Descriptor instead.
func (*StatBktActiveClientsResponse) Descriptor() ([]byte, []int) {
	return file_src_jrpc_netai_proto_rawDescGZIP(), []int{10}
}

func (x *StatBktActiveClientsResponse) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

var File_src_jrpc_netai_proto protoreflect.FileDescriptor

var file_src_jrpc_netai_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x72, 0x63, 0x2f, 0x6a, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6a, 0x72, 0x70, 0x63, 0x1a, 0x15, 0x73, 0x72,
	0x63, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x1b, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x61, 0x69, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x5e, 0x0a, 0x24, 0x53,
	0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x6f, 0x70, 0x31, 0x30, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x1a, 0x53,
	0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x56, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x6f, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x6f, 0x6c, 0x22, 0x44, 0x0a,
	0x0a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x6f, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x76, 0x6f, 0x6c, 0x22, 0x77, 0x0a, 0x23, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x56, 0x6f,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x6f, 0x70, 0x31, 0x30, 0x54, 0x6f, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x62, 0x79,
	0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x6f, 0x6c, 0x52, 0x05, 0x62, 0x79,
	0x43, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x62, 0x79, 0x5f, 0x76, 0x6f, 0x6c, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x56, 0x6f, 0x6c, 0x52, 0x05, 0x62, 0x79, 0x56, 0x6f, 0x6c, 0x22, 0x51, 0x0a, 0x19,
	0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x61, 0x76, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x22,
	0xca, 0x02, 0x0a, 0x09, 0x42, 0x6b, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72,
	0x63, 0x18, 0x51, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x73, 0x74, 0x18, 0x52, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x53, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x54, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x15,
	0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6b,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x9d, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x61,
	0x69, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x61, 0x76, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69,
	0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x76, 0x6f, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x6f, 0x6c,
	0x22, 0x30, 0x0a, 0x1c, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6b, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63,
	0x6e, 0x74, 0x42, 0x41, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x72, 0x70, 0x63, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x74, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x6a, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_jrpc_netai_proto_rawDescOnce sync.Once
	file_src_jrpc_netai_proto_rawDescData = file_src_jrpc_netai_proto_rawDesc
)

func file_src_jrpc_netai_proto_rawDescGZIP() []byte {
	file_src_jrpc_netai_proto_rawDescOnce.Do(func() {
		file_src_jrpc_netai_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_jrpc_netai_proto_rawDescData)
	})
	return file_src_jrpc_netai_proto_rawDescData
}

var file_src_jrpc_netai_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_src_jrpc_netai_proto_goTypes = []interface{}{
	(*StatBktRateAllTodayResponse)(nil),          // 0: jrpc.StatBktRateAllTodayResponse
	(*ChannelRate)(nil),                          // 1: jrpc.ChannelRate
	(*StatBktRateChannelTop10TodayResponse)(nil), // 2: jrpc.StatBktRateChannelTop10TodayResponse
	(*StatBktVolAllTodayResponse)(nil),           // 3: jrpc.StatBktVolAllTodayResponse
	(*ChannelVol)(nil),                           // 4: jrpc.ChannelVol
	(*StatBktVolChannelTop10TodayResponse)(nil),  // 5: jrpc.StatBktVolChannelTop10TodayResponse
	(*StatBktLapseTodayResponse)(nil),            // 6: jrpc.StatBktLapseTodayResponse
	(*BktRecord)(nil),                            // 7: jrpc.BktRecord
	(*StatBktRecordResponse)(nil),                // 8: jrpc.StatBktRecordResponse
	(*StatBktNodeResponse)(nil),                  // 9: jrpc.StatBktNodeResponse
	(*StatBktActiveClientsResponse)(nil),         // 10: jrpc.StatBktActiveClientsResponse
	(netai.TrafficLoad)(0),                       // 11: netai.TrafficLoad
}
var file_src_jrpc_netai_proto_depIdxs = []int32{
	1,  // 0: jrpc.StatBktRateChannelTop10TodayResponse.channel_rates:type_name -> jrpc.ChannelRate
	4,  // 1: jrpc.StatBktVolChannelTop10TodayResponse.by_cnt:type_name -> jrpc.ChannelVol
	4,  // 2: jrpc.StatBktVolChannelTop10TodayResponse.by_vol:type_name -> jrpc.ChannelVol
	7,  // 3: jrpc.StatBktRecordResponse.records:type_name -> jrpc.BktRecord
	11, // 4: jrpc.StatBktNodeResponse.traffic:type_name -> netai.TrafficLoad
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_src_jrpc_netai_proto_init() }
func file_src_jrpc_netai_proto_init() {
	if File_src_jrpc_netai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_jrpc_netai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktRateAllTodayResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelRate); i {
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
		file_src_jrpc_netai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktRateChannelTop10TodayResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktVolAllTodayResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelVol); i {
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
		file_src_jrpc_netai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktVolChannelTop10TodayResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktLapseTodayResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BktRecord); i {
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
		file_src_jrpc_netai_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktRecordResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktNodeResponse); i {
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
		file_src_jrpc_netai_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBktActiveClientsResponse); i {
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
			RawDescriptor: file_src_jrpc_netai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_jrpc_netai_proto_goTypes,
		DependencyIndexes: file_src_jrpc_netai_proto_depIdxs,
		MessageInfos:      file_src_jrpc_netai_proto_msgTypes,
	}.Build()
	File_src_jrpc_netai_proto = out.File
	file_src_jrpc_netai_proto_rawDesc = nil
	file_src_jrpc_netai_proto_goTypes = nil
	file_src_jrpc_netai_proto_depIdxs = nil
}
