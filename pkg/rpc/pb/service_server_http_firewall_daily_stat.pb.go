// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_server_http_firewall_daily_stat.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 组合服务的Dashboard
type ComposeServerHTTPFirewallDashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day      string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ServerId int64  `protobuf:"varint,3,opt,name=serverId,proto3" json:"serverId,omitempty"`
}

func (x *ComposeServerHTTPFirewallDashboardRequest) Reset() {
	*x = ComposeServerHTTPFirewallDashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServerHTTPFirewallDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServerHTTPFirewallDashboardRequest) ProtoMessage() {}

func (x *ComposeServerHTTPFirewallDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServerHTTPFirewallDashboardRequest.ProtoReflect.Descriptor instead.
func (*ComposeServerHTTPFirewallDashboardRequest) Descriptor() ([]byte, []int) {
	return file_service_server_http_firewall_daily_stat_proto_rawDescGZIP(), []int{0}
}

func (x *ComposeServerHTTPFirewallDashboardRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ComposeServerHTTPFirewallDashboardRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type ComposeServerHTTPFirewallDashboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountDailyLog          int64                                                                   `protobuf:"varint,1,opt,name=countDailyLog,proto3" json:"countDailyLog,omitempty"`
	CountDailyBlock        int64                                                                   `protobuf:"varint,2,opt,name=countDailyBlock,proto3" json:"countDailyBlock,omitempty"`
	CountDailyCaptcha      int64                                                                   `protobuf:"varint,3,opt,name=countDailyCaptcha,proto3" json:"countDailyCaptcha,omitempty"`
	CountWeeklyBlock       int64                                                                   `protobuf:"varint,4,opt,name=countWeeklyBlock,proto3" json:"countWeeklyBlock,omitempty"`
	CountMonthlyBlock      int64                                                                   `protobuf:"varint,5,opt,name=countMonthlyBlock,proto3" json:"countMonthlyBlock,omitempty"`
	HttpFirewallRuleGroups []*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat `protobuf:"bytes,30,rep,name=httpFirewallRuleGroups,proto3" json:"httpFirewallRuleGroups,omitempty"`
	LogDailyStats          []*ComposeServerHTTPFirewallDashboardResponse_DailyStat                 `protobuf:"bytes,31,rep,name=logDailyStats,proto3" json:"logDailyStats,omitempty"`
	BlockDailyStats        []*ComposeServerHTTPFirewallDashboardResponse_DailyStat                 `protobuf:"bytes,32,rep,name=blockDailyStats,proto3" json:"blockDailyStats,omitempty"`
	CaptchaDailyStats      []*ComposeServerHTTPFirewallDashboardResponse_DailyStat                 `protobuf:"bytes,33,rep,name=captchaDailyStats,proto3" json:"captchaDailyStats,omitempty"`
}

func (x *ComposeServerHTTPFirewallDashboardResponse) Reset() {
	*x = ComposeServerHTTPFirewallDashboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServerHTTPFirewallDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServerHTTPFirewallDashboardResponse) ProtoMessage() {}

func (x *ComposeServerHTTPFirewallDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServerHTTPFirewallDashboardResponse.ProtoReflect.Descriptor instead.
func (*ComposeServerHTTPFirewallDashboardResponse) Descriptor() ([]byte, []int) {
	return file_service_server_http_firewall_daily_stat_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCountDailyLog() int64 {
	if x != nil {
		return x.CountDailyLog
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCountDailyBlock() int64 {
	if x != nil {
		return x.CountDailyBlock
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCountDailyCaptcha() int64 {
	if x != nil {
		return x.CountDailyCaptcha
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCountWeeklyBlock() int64 {
	if x != nil {
		return x.CountWeeklyBlock
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCountMonthlyBlock() int64 {
	if x != nil {
		return x.CountMonthlyBlock
	}
	return 0
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetHttpFirewallRuleGroups() []*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat {
	if x != nil {
		return x.HttpFirewallRuleGroups
	}
	return nil
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetLogDailyStats() []*ComposeServerHTTPFirewallDashboardResponse_DailyStat {
	if x != nil {
		return x.LogDailyStats
	}
	return nil
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetBlockDailyStats() []*ComposeServerHTTPFirewallDashboardResponse_DailyStat {
	if x != nil {
		return x.BlockDailyStats
	}
	return nil
}

func (x *ComposeServerHTTPFirewallDashboardResponse) GetCaptchaDailyStats() []*ComposeServerHTTPFirewallDashboardResponse_DailyStat {
	if x != nil {
		return x.CaptchaDailyStats
	}
	return nil
}

type ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFirewallRuleGroup *HTTPFirewallRuleGroup `protobuf:"bytes,1,opt,name=httpFirewallRuleGroup,proto3" json:"httpFirewallRuleGroup,omitempty"`
	Count                 int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) Reset() {
	*x = ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) ProtoMessage() {}

func (x *ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat.ProtoReflect.Descriptor instead.
func (*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) Descriptor() ([]byte, []int) {
	return file_service_server_http_firewall_daily_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) GetHttpFirewallRuleGroup() *HTTPFirewallRuleGroup {
	if x != nil {
		return x.HttpFirewallRuleGroup
	}
	return nil
}

func (x *ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ComposeServerHTTPFirewallDashboardResponse_DailyStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day   string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ComposeServerHTTPFirewallDashboardResponse_DailyStat) Reset() {
	*x = ComposeServerHTTPFirewallDashboardResponse_DailyStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServerHTTPFirewallDashboardResponse_DailyStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServerHTTPFirewallDashboardResponse_DailyStat) ProtoMessage() {}

func (x *ComposeServerHTTPFirewallDashboardResponse_DailyStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_http_firewall_daily_stat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServerHTTPFirewallDashboardResponse_DailyStat.ProtoReflect.Descriptor instead.
func (*ComposeServerHTTPFirewallDashboardResponse_DailyStat) Descriptor() ([]byte, []int) {
	return file_service_server_http_firewall_daily_stat_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ComposeServerHTTPFirewallDashboardResponse_DailyStat) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ComposeServerHTTPFirewallDashboardResponse_DailyStat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_server_http_firewall_daily_stat_proto protoreflect.FileDescriptor

var file_service_server_http_firewall_daily_stat_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x71, 0x0a, 0x29, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xed, 0x06, 0x0a, 0x2a, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x11,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x80, 0x01, 0x0a, 0x16, 0x68,
	0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x52, 0x16, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x5e, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x1f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0d,
	0x6c, 0x6f, 0x67, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x62, 0x0a,
	0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x66, 0x0a, 0x11, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x11, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x82, 0x01, 0x0a, 0x19, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x12, 0x4f, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x15, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x33,
	0x0a, 0x09, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xaa, 0x01, 0x0a, 0x22, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x2d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_http_firewall_daily_stat_proto_rawDescOnce sync.Once
	file_service_server_http_firewall_daily_stat_proto_rawDescData = file_service_server_http_firewall_daily_stat_proto_rawDesc
)

func file_service_server_http_firewall_daily_stat_proto_rawDescGZIP() []byte {
	file_service_server_http_firewall_daily_stat_proto_rawDescOnce.Do(func() {
		file_service_server_http_firewall_daily_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_http_firewall_daily_stat_proto_rawDescData)
	})
	return file_service_server_http_firewall_daily_stat_proto_rawDescData
}

var file_service_server_http_firewall_daily_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_server_http_firewall_daily_stat_proto_goTypes = []interface{}{
	(*ComposeServerHTTPFirewallDashboardRequest)(nil),                            // 0: pb.ComposeServerHTTPFirewallDashboardRequest
	(*ComposeServerHTTPFirewallDashboardResponse)(nil),                           // 1: pb.ComposeServerHTTPFirewallDashboardResponse
	(*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat)(nil), // 2: pb.ComposeServerHTTPFirewallDashboardResponse.HTTPFirewallRuleGroupStat
	(*ComposeServerHTTPFirewallDashboardResponse_DailyStat)(nil),                 // 3: pb.ComposeServerHTTPFirewallDashboardResponse.DailyStat
	(*HTTPFirewallRuleGroup)(nil),                                                // 4: pb.HTTPFirewallRuleGroup
}
var file_service_server_http_firewall_daily_stat_proto_depIdxs = []int32{
	2, // 0: pb.ComposeServerHTTPFirewallDashboardResponse.httpFirewallRuleGroups:type_name -> pb.ComposeServerHTTPFirewallDashboardResponse.HTTPFirewallRuleGroupStat
	3, // 1: pb.ComposeServerHTTPFirewallDashboardResponse.logDailyStats:type_name -> pb.ComposeServerHTTPFirewallDashboardResponse.DailyStat
	3, // 2: pb.ComposeServerHTTPFirewallDashboardResponse.blockDailyStats:type_name -> pb.ComposeServerHTTPFirewallDashboardResponse.DailyStat
	3, // 3: pb.ComposeServerHTTPFirewallDashboardResponse.captchaDailyStats:type_name -> pb.ComposeServerHTTPFirewallDashboardResponse.DailyStat
	4, // 4: pb.ComposeServerHTTPFirewallDashboardResponse.HTTPFirewallRuleGroupStat.httpFirewallRuleGroup:type_name -> pb.HTTPFirewallRuleGroup
	0, // 5: pb.ServerHTTPFirewallDailyStatService.composeServerHTTPFirewallDashboard:input_type -> pb.ComposeServerHTTPFirewallDashboardRequest
	1, // 6: pb.ServerHTTPFirewallDailyStatService.composeServerHTTPFirewallDashboard:output_type -> pb.ComposeServerHTTPFirewallDashboardResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_server_http_firewall_daily_stat_proto_init() }
func file_service_server_http_firewall_daily_stat_proto_init() {
	if File_service_server_http_firewall_daily_stat_proto != nil {
		return
	}
	file_models_model_http_firewall_rule_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_http_firewall_daily_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServerHTTPFirewallDashboardRequest); i {
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
		file_service_server_http_firewall_daily_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServerHTTPFirewallDashboardResponse); i {
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
		file_service_server_http_firewall_daily_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServerHTTPFirewallDashboardResponse_HTTPFirewallRuleGroupStat); i {
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
		file_service_server_http_firewall_daily_stat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServerHTTPFirewallDashboardResponse_DailyStat); i {
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
			RawDescriptor: file_service_server_http_firewall_daily_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_http_firewall_daily_stat_proto_goTypes,
		DependencyIndexes: file_service_server_http_firewall_daily_stat_proto_depIdxs,
		MessageInfos:      file_service_server_http_firewall_daily_stat_proto_msgTypes,
	}.Build()
	File_service_server_http_firewall_daily_stat_proto = out.File
	file_service_server_http_firewall_daily_stat_proto_rawDesc = nil
	file_service_server_http_firewall_daily_stat_proto_goTypes = nil
	file_service_server_http_firewall_daily_stat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerHTTPFirewallDailyStatServiceClient is the client API for ServerHTTPFirewallDailyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerHTTPFirewallDailyStatServiceClient interface {
	// 组合服务的Dashboard
	ComposeServerHTTPFirewallDashboard(ctx context.Context, in *ComposeServerHTTPFirewallDashboardRequest, opts ...grpc.CallOption) (*ComposeServerHTTPFirewallDashboardResponse, error)
}

type serverHTTPFirewallDailyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerHTTPFirewallDailyStatServiceClient(cc grpc.ClientConnInterface) ServerHTTPFirewallDailyStatServiceClient {
	return &serverHTTPFirewallDailyStatServiceClient{cc}
}

func (c *serverHTTPFirewallDailyStatServiceClient) ComposeServerHTTPFirewallDashboard(ctx context.Context, in *ComposeServerHTTPFirewallDashboardRequest, opts ...grpc.CallOption) (*ComposeServerHTTPFirewallDashboardResponse, error) {
	out := new(ComposeServerHTTPFirewallDashboardResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerHTTPFirewallDailyStatService/composeServerHTTPFirewallDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerHTTPFirewallDailyStatServiceServer is the server API for ServerHTTPFirewallDailyStatService service.
type ServerHTTPFirewallDailyStatServiceServer interface {
	// 组合服务的Dashboard
	ComposeServerHTTPFirewallDashboard(context.Context, *ComposeServerHTTPFirewallDashboardRequest) (*ComposeServerHTTPFirewallDashboardResponse, error)
}

// UnimplementedServerHTTPFirewallDailyStatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerHTTPFirewallDailyStatServiceServer struct {
}

func (*UnimplementedServerHTTPFirewallDailyStatServiceServer) ComposeServerHTTPFirewallDashboard(context.Context, *ComposeServerHTTPFirewallDashboardRequest) (*ComposeServerHTTPFirewallDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeServerHTTPFirewallDashboard not implemented")
}

func RegisterServerHTTPFirewallDailyStatServiceServer(s *grpc.Server, srv ServerHTTPFirewallDailyStatServiceServer) {
	s.RegisterService(&_ServerHTTPFirewallDailyStatService_serviceDesc, srv)
}

func _ServerHTTPFirewallDailyStatService_ComposeServerHTTPFirewallDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServerHTTPFirewallDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerHTTPFirewallDailyStatServiceServer).ComposeServerHTTPFirewallDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerHTTPFirewallDailyStatService/ComposeServerHTTPFirewallDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerHTTPFirewallDailyStatServiceServer).ComposeServerHTTPFirewallDashboard(ctx, req.(*ComposeServerHTTPFirewallDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerHTTPFirewallDailyStatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerHTTPFirewallDailyStatService",
	HandlerType: (*ServerHTTPFirewallDailyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "composeServerHTTPFirewallDashboard",
			Handler:    _ServerHTTPFirewallDailyStatService_ComposeServerHTTPFirewallDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_http_firewall_daily_stat.proto",
}
