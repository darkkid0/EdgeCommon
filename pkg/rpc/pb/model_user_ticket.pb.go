// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_user_ticket.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// 工单
type UserTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId          int64               `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	UserId              int64               `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Subject             string              `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                string              `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Status              string              `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt           int64               `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	LastLogAt           int64               `protobuf:"varint,8,opt,name=lastLogAt,proto3" json:"lastLogAt,omitempty"`
	UserTicketCategory  *UserTicketCategory `protobuf:"bytes,30,opt,name=userTicketCategory,proto3" json:"userTicketCategory,omitempty"`
	User                *User               `protobuf:"bytes,31,opt,name=user,proto3" json:"user,omitempty"`
	LatestUserTicketLog *UserTicketLog      `protobuf:"bytes,32,opt,name=latestUserTicketLog,proto3" json:"latestUserTicketLog,omitempty"`
}

func (x *UserTicket) Reset() {
	*x = UserTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTicket) ProtoMessage() {}

func (x *UserTicket) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTicket.ProtoReflect.Descriptor instead.
func (*UserTicket) Descriptor() ([]byte, []int) {
	return file_models_model_user_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *UserTicket) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTicket) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *UserTicket) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserTicket) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *UserTicket) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *UserTicket) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserTicket) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserTicket) GetLastLogAt() int64 {
	if x != nil {
		return x.LastLogAt
	}
	return 0
}

func (x *UserTicket) GetUserTicketCategory() *UserTicketCategory {
	if x != nil {
		return x.UserTicketCategory
	}
	return nil
}

func (x *UserTicket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserTicket) GetLatestUserTicketLog() *UserTicketLog {
	if x != nil {
		return x.LatestUserTicketLog
	}
	return nil
}

var File_models_model_user_ticket_proto protoreflect.FileDescriptor

var file_models_model_user_ticket_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x13, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_ticket_proto_rawDescOnce sync.Once
	file_models_model_user_ticket_proto_rawDescData = file_models_model_user_ticket_proto_rawDesc
)

func file_models_model_user_ticket_proto_rawDescGZIP() []byte {
	file_models_model_user_ticket_proto_rawDescOnce.Do(func() {
		file_models_model_user_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_ticket_proto_rawDescData)
	})
	return file_models_model_user_ticket_proto_rawDescData
}

var file_models_model_user_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_ticket_proto_goTypes = []interface{}{
	(*UserTicket)(nil),         // 0: pb.UserTicket
	(*UserTicketCategory)(nil), // 1: pb.UserTicketCategory
	(*User)(nil),               // 2: pb.User
	(*UserTicketLog)(nil),      // 3: pb.UserTicketLog
}
var file_models_model_user_ticket_proto_depIdxs = []int32{
	1, // 0: pb.UserTicket.userTicketCategory:type_name -> pb.UserTicketCategory
	2, // 1: pb.UserTicket.user:type_name -> pb.User
	3, // 2: pb.UserTicket.latestUserTicketLog:type_name -> pb.UserTicketLog
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_user_ticket_proto_init() }
func file_models_model_user_ticket_proto_init() {
	if File_models_model_user_ticket_proto != nil {
		return
	}
	file_models_model_user_ticket_category_proto_init()
	file_models_model_user_proto_init()
	file_models_model_user_ticket_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTicket); i {
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
			RawDescriptor: file_models_model_user_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_ticket_proto_goTypes,
		DependencyIndexes: file_models_model_user_ticket_proto_depIdxs,
		MessageInfos:      file_models_model_user_ticket_proto_msgTypes,
	}.Build()
	File_models_model_user_ticket_proto = out.File
	file_models_model_user_ticket_proto_rawDesc = nil
	file_models_model_user_ticket_proto_goTypes = nil
	file_models_model_user_ticket_proto_depIdxs = nil
}
