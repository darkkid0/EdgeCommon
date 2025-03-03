// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_reverse_proxy.proto

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

type ReverseProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SchedulingJSON     []byte `protobuf:"bytes,2,opt,name=schedulingJSON,proto3" json:"schedulingJSON,omitempty"`
	PrimaryOriginsJSON []byte `protobuf:"bytes,3,opt,name=primaryOriginsJSON,proto3" json:"primaryOriginsJSON,omitempty"`
	BackupOriginsJSON  []byte `protobuf:"bytes,4,opt,name=backupOriginsJSON,proto3" json:"backupOriginsJSON,omitempty"`
}

func (x *ReverseProxy) Reset() {
	*x = ReverseProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_reverse_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseProxy) ProtoMessage() {}

func (x *ReverseProxy) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_reverse_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseProxy.ProtoReflect.Descriptor instead.
func (*ReverseProxy) Descriptor() ([]byte, []int) {
	return file_models_model_reverse_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *ReverseProxy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReverseProxy) GetSchedulingJSON() []byte {
	if x != nil {
		return x.SchedulingJSON
	}
	return nil
}

func (x *ReverseProxy) GetPrimaryOriginsJSON() []byte {
	if x != nil {
		return x.PrimaryOriginsJSON
	}
	return nil
}

func (x *ReverseProxy) GetBackupOriginsJSON() []byte {
	if x != nil {
		return x.BackupOriginsJSON
	}
	return nil
}

var File_models_model_reverse_proxy_proto protoreflect.FileDescriptor

var file_models_model_reverse_proxy_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x2e, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x2c, 0x0a, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_reverse_proxy_proto_rawDescOnce sync.Once
	file_models_model_reverse_proxy_proto_rawDescData = file_models_model_reverse_proxy_proto_rawDesc
)

func file_models_model_reverse_proxy_proto_rawDescGZIP() []byte {
	file_models_model_reverse_proxy_proto_rawDescOnce.Do(func() {
		file_models_model_reverse_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_reverse_proxy_proto_rawDescData)
	})
	return file_models_model_reverse_proxy_proto_rawDescData
}

var file_models_model_reverse_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_reverse_proxy_proto_goTypes = []interface{}{
	(*ReverseProxy)(nil), // 0: pb.ReverseProxy
}
var file_models_model_reverse_proxy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_reverse_proxy_proto_init() }
func file_models_model_reverse_proxy_proto_init() {
	if File_models_model_reverse_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_reverse_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseProxy); i {
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
			RawDescriptor: file_models_model_reverse_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_reverse_proxy_proto_goTypes,
		DependencyIndexes: file_models_model_reverse_proxy_proto_depIdxs,
		MessageInfos:      file_models_model_reverse_proxy_proto_msgTypes,
	}.Build()
	File_models_model_reverse_proxy_proto = out.File
	file_models_model_reverse_proxy_proto_rawDesc = nil
	file_models_model_reverse_proxy_proto_goTypes = nil
	file_models_model_reverse_proxy_proto_depIdxs = nil
}
