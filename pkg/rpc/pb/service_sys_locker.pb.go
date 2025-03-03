// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_sys_locker.proto

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

// 获得锁
type SysLockerLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TimeoutSeconds int64  `protobuf:"varint,2,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds,omitempty"`
}

func (x *SysLockerLockRequest) Reset() {
	*x = SysLockerLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_sys_locker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLockerLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLockerLockRequest) ProtoMessage() {}

func (x *SysLockerLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_sys_locker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLockerLockRequest.ProtoReflect.Descriptor instead.
func (*SysLockerLockRequest) Descriptor() ([]byte, []int) {
	return file_service_sys_locker_proto_rawDescGZIP(), []int{0}
}

func (x *SysLockerLockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SysLockerLockRequest) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

type SysLockerLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SysLockerLockResponse) Reset() {
	*x = SysLockerLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_sys_locker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLockerLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLockerLockResponse) ProtoMessage() {}

func (x *SysLockerLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_sys_locker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLockerLockResponse.ProtoReflect.Descriptor instead.
func (*SysLockerLockResponse) Descriptor() ([]byte, []int) {
	return file_service_sys_locker_proto_rawDescGZIP(), []int{1}
}

func (x *SysLockerLockResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// 释放锁
type SysLockerUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SysLockerUnlockRequest) Reset() {
	*x = SysLockerUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_sys_locker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLockerUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLockerUnlockRequest) ProtoMessage() {}

func (x *SysLockerUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_sys_locker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLockerUnlockRequest.ProtoReflect.Descriptor instead.
func (*SysLockerUnlockRequest) Descriptor() ([]byte, []int) {
	return file_service_sys_locker_proto_rawDescGZIP(), []int{2}
}

func (x *SysLockerUnlockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_service_sys_locker_proto protoreflect.FileDescriptor

var file_service_sys_locker_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x14, 0x53, 0x79, 0x73,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x53,
	0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x2a, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x32, 0x97, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x73, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x53,
	0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_sys_locker_proto_rawDescOnce sync.Once
	file_service_sys_locker_proto_rawDescData = file_service_sys_locker_proto_rawDesc
)

func file_service_sys_locker_proto_rawDescGZIP() []byte {
	file_service_sys_locker_proto_rawDescOnce.Do(func() {
		file_service_sys_locker_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_sys_locker_proto_rawDescData)
	})
	return file_service_sys_locker_proto_rawDescData
}

var file_service_sys_locker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_sys_locker_proto_goTypes = []interface{}{
	(*SysLockerLockRequest)(nil),   // 0: pb.SysLockerLockRequest
	(*SysLockerLockResponse)(nil),  // 1: pb.SysLockerLockResponse
	(*SysLockerUnlockRequest)(nil), // 2: pb.SysLockerUnlockRequest
	(*RPCSuccess)(nil),             // 3: pb.RPCSuccess
}
var file_service_sys_locker_proto_depIdxs = []int32{
	0, // 0: pb.SysLockerService.SysLockerLock:input_type -> pb.SysLockerLockRequest
	2, // 1: pb.SysLockerService.SysLockerUnlock:input_type -> pb.SysLockerUnlockRequest
	1, // 2: pb.SysLockerService.SysLockerLock:output_type -> pb.SysLockerLockResponse
	3, // 3: pb.SysLockerService.SysLockerUnlock:output_type -> pb.RPCSuccess
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_sys_locker_proto_init() }
func file_service_sys_locker_proto_init() {
	if File_service_sys_locker_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_sys_locker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLockerLockRequest); i {
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
		file_service_sys_locker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLockerLockResponse); i {
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
		file_service_sys_locker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLockerUnlockRequest); i {
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
			RawDescriptor: file_service_sys_locker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_sys_locker_proto_goTypes,
		DependencyIndexes: file_service_sys_locker_proto_depIdxs,
		MessageInfos:      file_service_sys_locker_proto_msgTypes,
	}.Build()
	File_service_sys_locker_proto = out.File
	file_service_sys_locker_proto_rawDesc = nil
	file_service_sys_locker_proto_goTypes = nil
	file_service_sys_locker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SysLockerServiceClient is the client API for SysLockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SysLockerServiceClient interface {
	// 获得锁
	SysLockerLock(ctx context.Context, in *SysLockerLockRequest, opts ...grpc.CallOption) (*SysLockerLockResponse, error)
	// 释放锁
	SysLockerUnlock(ctx context.Context, in *SysLockerUnlockRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type sysLockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSysLockerServiceClient(cc grpc.ClientConnInterface) SysLockerServiceClient {
	return &sysLockerServiceClient{cc}
}

func (c *sysLockerServiceClient) SysLockerLock(ctx context.Context, in *SysLockerLockRequest, opts ...grpc.CallOption) (*SysLockerLockResponse, error) {
	out := new(SysLockerLockResponse)
	err := c.cc.Invoke(ctx, "/pb.SysLockerService/SysLockerLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysLockerServiceClient) SysLockerUnlock(ctx context.Context, in *SysLockerUnlockRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.SysLockerService/SysLockerUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysLockerServiceServer is the server API for SysLockerService service.
type SysLockerServiceServer interface {
	// 获得锁
	SysLockerLock(context.Context, *SysLockerLockRequest) (*SysLockerLockResponse, error)
	// 释放锁
	SysLockerUnlock(context.Context, *SysLockerUnlockRequest) (*RPCSuccess, error)
}

// UnimplementedSysLockerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSysLockerServiceServer struct {
}

func (*UnimplementedSysLockerServiceServer) SysLockerLock(context.Context, *SysLockerLockRequest) (*SysLockerLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLockerLock not implemented")
}
func (*UnimplementedSysLockerServiceServer) SysLockerUnlock(context.Context, *SysLockerUnlockRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLockerUnlock not implemented")
}

func RegisterSysLockerServiceServer(s *grpc.Server, srv SysLockerServiceServer) {
	s.RegisterService(&_SysLockerService_serviceDesc, srv)
}

func _SysLockerService_SysLockerLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLockerLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysLockerServiceServer).SysLockerLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SysLockerService/SysLockerLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysLockerServiceServer).SysLockerLock(ctx, req.(*SysLockerLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysLockerService_SysLockerUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLockerUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysLockerServiceServer).SysLockerUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SysLockerService/SysLockerUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysLockerServiceServer).SysLockerUnlock(ctx, req.(*SysLockerUnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SysLockerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SysLockerService",
	HandlerType: (*SysLockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysLockerLock",
			Handler:    _SysLockerService_SysLockerLock_Handler,
		},
		{
			MethodName: "SysLockerUnlock",
			Handler:    _SysLockerService_SysLockerUnlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_sys_locker.proto",
}
