// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_user_verify_code.proto

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

// 发送重置密码验证码
type SendUserVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`     // 类型：重置密码（resetPassword）
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`   // 已验证邮箱地址
	Mobile string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"` // 已验证手机号
}

func (x *SendUserVerifyCodeRequest) Reset() {
	*x = SendUserVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_verify_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserVerifyCodeRequest) ProtoMessage() {}

func (x *SendUserVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_verify_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*SendUserVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_service_user_verify_code_proto_rawDescGZIP(), []int{0}
}

func (x *SendUserVerifyCodeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendUserVerifyCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendUserVerifyCodeRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SendUserVerifyCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeLength int32 `protobuf:"varint,1,opt,name=codeLength,proto3" json:"codeLength,omitempty"` // 验证码长度
}

func (x *SendUserVerifyCodeResponse) Reset() {
	*x = SendUserVerifyCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_verify_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserVerifyCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserVerifyCodeResponse) ProtoMessage() {}

func (x *SendUserVerifyCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_verify_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserVerifyCodeResponse.ProtoReflect.Descriptor instead.
func (*SendUserVerifyCodeResponse) Descriptor() ([]byte, []int) {
	return file_service_user_verify_code_proto_rawDescGZIP(), []int{1}
}

func (x *SendUserVerifyCodeResponse) GetCodeLength() int32 {
	if x != nil {
		return x.CodeLength
	}
	return 0
}

// 校验验证码
type ValidateUserVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`     // 类型：重置密码（resetPassword）
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`   // 已验证邮箱地址
	Mobile string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"` // 已验证手机号
	Code   string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`     // 验证码
	// 找回密码
	NewPassword string `protobuf:"bytes,10,opt,name=newPassword,proto3" json:"newPassword,omitempty"` // 新密码
}

func (x *ValidateUserVerifyCodeRequest) Reset() {
	*x = ValidateUserVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_verify_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserVerifyCodeRequest) ProtoMessage() {}

func (x *ValidateUserVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_verify_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*ValidateUserVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_service_user_verify_code_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateUserVerifyCodeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ValidateUserVerifyCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ValidateUserVerifyCodeRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ValidateUserVerifyCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ValidateUserVerifyCodeRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ValidateUserVerifyCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk         bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`                // 是否成功
	ErrorCode    string `protobuf:"bytes,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`       // 错误代号
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"` // 错误消息
}

func (x *ValidateUserVerifyCodeResponse) Reset() {
	*x = ValidateUserVerifyCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_verify_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserVerifyCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserVerifyCodeResponse) ProtoMessage() {}

func (x *ValidateUserVerifyCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_verify_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserVerifyCodeResponse.ProtoReflect.Descriptor instead.
func (*ValidateUserVerifyCodeResponse) Descriptor() ([]byte, []int) {
	return file_service_user_verify_code_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateUserVerifyCodeResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ValidateUserVerifyCodeResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *ValidateUserVerifyCodeResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_service_user_verify_code_proto protoreflect.FileDescriptor

var file_service_user_verify_code_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x5d, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x76, 0x0a, 0x1e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a,
	0x12, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_user_verify_code_proto_rawDescOnce sync.Once
	file_service_user_verify_code_proto_rawDescData = file_service_user_verify_code_proto_rawDesc
)

func file_service_user_verify_code_proto_rawDescGZIP() []byte {
	file_service_user_verify_code_proto_rawDescOnce.Do(func() {
		file_service_user_verify_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_verify_code_proto_rawDescData)
	})
	return file_service_user_verify_code_proto_rawDescData
}

var file_service_user_verify_code_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_user_verify_code_proto_goTypes = []interface{}{
	(*SendUserVerifyCodeRequest)(nil),      // 0: pb.SendUserVerifyCodeRequest
	(*SendUserVerifyCodeResponse)(nil),     // 1: pb.SendUserVerifyCodeResponse
	(*ValidateUserVerifyCodeRequest)(nil),  // 2: pb.ValidateUserVerifyCodeRequest
	(*ValidateUserVerifyCodeResponse)(nil), // 3: pb.ValidateUserVerifyCodeResponse
}
var file_service_user_verify_code_proto_depIdxs = []int32{
	0, // 0: pb.UserVerifyCodeService.sendUserVerifyCode:input_type -> pb.SendUserVerifyCodeRequest
	2, // 1: pb.UserVerifyCodeService.validateUserVerifyCode:input_type -> pb.ValidateUserVerifyCodeRequest
	1, // 2: pb.UserVerifyCodeService.sendUserVerifyCode:output_type -> pb.SendUserVerifyCodeResponse
	3, // 3: pb.UserVerifyCodeService.validateUserVerifyCode:output_type -> pb.ValidateUserVerifyCodeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_user_verify_code_proto_init() }
func file_service_user_verify_code_proto_init() {
	if File_service_user_verify_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_user_verify_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserVerifyCodeRequest); i {
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
		file_service_user_verify_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserVerifyCodeResponse); i {
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
		file_service_user_verify_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserVerifyCodeRequest); i {
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
		file_service_user_verify_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserVerifyCodeResponse); i {
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
			RawDescriptor: file_service_user_verify_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_verify_code_proto_goTypes,
		DependencyIndexes: file_service_user_verify_code_proto_depIdxs,
		MessageInfos:      file_service_user_verify_code_proto_msgTypes,
	}.Build()
	File_service_user_verify_code_proto = out.File
	file_service_user_verify_code_proto_rawDesc = nil
	file_service_user_verify_code_proto_goTypes = nil
	file_service_user_verify_code_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserVerifyCodeServiceClient is the client API for UserVerifyCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserVerifyCodeServiceClient interface {
	// 发送重置密码验证码
	SendUserVerifyCode(ctx context.Context, in *SendUserVerifyCodeRequest, opts ...grpc.CallOption) (*SendUserVerifyCodeResponse, error)
	// 校验验证码
	ValidateUserVerifyCode(ctx context.Context, in *ValidateUserVerifyCodeRequest, opts ...grpc.CallOption) (*ValidateUserVerifyCodeResponse, error)
}

type userVerifyCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserVerifyCodeServiceClient(cc grpc.ClientConnInterface) UserVerifyCodeServiceClient {
	return &userVerifyCodeServiceClient{cc}
}

func (c *userVerifyCodeServiceClient) SendUserVerifyCode(ctx context.Context, in *SendUserVerifyCodeRequest, opts ...grpc.CallOption) (*SendUserVerifyCodeResponse, error) {
	out := new(SendUserVerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserVerifyCodeService/sendUserVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userVerifyCodeServiceClient) ValidateUserVerifyCode(ctx context.Context, in *ValidateUserVerifyCodeRequest, opts ...grpc.CallOption) (*ValidateUserVerifyCodeResponse, error) {
	out := new(ValidateUserVerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserVerifyCodeService/validateUserVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserVerifyCodeServiceServer is the server API for UserVerifyCodeService service.
type UserVerifyCodeServiceServer interface {
	// 发送重置密码验证码
	SendUserVerifyCode(context.Context, *SendUserVerifyCodeRequest) (*SendUserVerifyCodeResponse, error)
	// 校验验证码
	ValidateUserVerifyCode(context.Context, *ValidateUserVerifyCodeRequest) (*ValidateUserVerifyCodeResponse, error)
}

// UnimplementedUserVerifyCodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserVerifyCodeServiceServer struct {
}

func (*UnimplementedUserVerifyCodeServiceServer) SendUserVerifyCode(context.Context, *SendUserVerifyCodeRequest) (*SendUserVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserVerifyCode not implemented")
}
func (*UnimplementedUserVerifyCodeServiceServer) ValidateUserVerifyCode(context.Context, *ValidateUserVerifyCodeRequest) (*ValidateUserVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserVerifyCode not implemented")
}

func RegisterUserVerifyCodeServiceServer(s *grpc.Server, srv UserVerifyCodeServiceServer) {
	s.RegisterService(&_UserVerifyCodeService_serviceDesc, srv)
}

func _UserVerifyCodeService_SendUserVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserVerifyCodeServiceServer).SendUserVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserVerifyCodeService/SendUserVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserVerifyCodeServiceServer).SendUserVerifyCode(ctx, req.(*SendUserVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserVerifyCodeService_ValidateUserVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserVerifyCodeServiceServer).ValidateUserVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserVerifyCodeService/ValidateUserVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserVerifyCodeServiceServer).ValidateUserVerifyCode(ctx, req.(*ValidateUserVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserVerifyCodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserVerifyCodeService",
	HandlerType: (*UserVerifyCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendUserVerifyCode",
			Handler:    _UserVerifyCodeService_SendUserVerifyCode_Handler,
		},
		{
			MethodName: "validateUserVerifyCode",
			Handler:    _UserVerifyCodeService_ValidateUserVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_verify_code.proto",
}
