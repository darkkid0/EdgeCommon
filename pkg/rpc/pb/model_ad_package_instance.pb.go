// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_ad_package_instance.proto

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

// 高防产品实例
type ADPackageInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn           bool         `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	AdPackageId    int64        `protobuf:"varint,3,opt,name=adPackageId,proto3" json:"adPackageId,omitempty"`
	NodeClusterId  int64        `protobuf:"varint,4,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	NodeIds        []int64      `protobuf:"varint,5,rep,packed,name=nodeIds,proto3" json:"nodeIds,omitempty"`
	IpAddresses    []string     `protobuf:"bytes,6,rep,name=ipAddresses,proto3" json:"ipAddresses,omitempty"`
	UserId         int64        `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`                 // 租用用户ID
	UserDayTo      string       `protobuf:"bytes,8,opt,name=userDayTo,proto3" json:"userDayTo,omitempty"`            // 租用日期
	UserInstanceId int64        `protobuf:"varint,9,opt,name=userInstanceId,proto3" json:"userInstanceId,omitempty"` // 当前绑定的用户实例ID
	NodeCluster    *NodeCluster `protobuf:"bytes,30,opt,name=nodeCluster,proto3" json:"nodeCluster,omitempty"`
	AdPackage      *ADPackage   `protobuf:"bytes,31,opt,name=adPackage,proto3" json:"adPackage,omitempty"`
	User           *User        `protobuf:"bytes,32,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ADPackageInstance) Reset() {
	*x = ADPackageInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ad_package_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADPackageInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADPackageInstance) ProtoMessage() {}

func (x *ADPackageInstance) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ad_package_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADPackageInstance.ProtoReflect.Descriptor instead.
func (*ADPackageInstance) Descriptor() ([]byte, []int) {
	return file_models_model_ad_package_instance_proto_rawDescGZIP(), []int{0}
}

func (x *ADPackageInstance) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ADPackageInstance) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *ADPackageInstance) GetAdPackageId() int64 {
	if x != nil {
		return x.AdPackageId
	}
	return 0
}

func (x *ADPackageInstance) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *ADPackageInstance) GetNodeIds() []int64 {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

func (x *ADPackageInstance) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *ADPackageInstance) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ADPackageInstance) GetUserDayTo() string {
	if x != nil {
		return x.UserDayTo
	}
	return ""
}

func (x *ADPackageInstance) GetUserInstanceId() int64 {
	if x != nil {
		return x.UserInstanceId
	}
	return 0
}

func (x *ADPackageInstance) GetNodeCluster() *NodeCluster {
	if x != nil {
		return x.NodeCluster
	}
	return nil
}

func (x *ADPackageInstance) GetAdPackage() *ADPackage {
	if x != nil {
		return x.AdPackage
	}
	return nil
}

func (x *ADPackageInstance) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_models_model_ad_package_instance_proto protoreflect.FileDescriptor

var file_models_model_ad_package_instance_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x11, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x44, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x09, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ad_package_instance_proto_rawDescOnce sync.Once
	file_models_model_ad_package_instance_proto_rawDescData = file_models_model_ad_package_instance_proto_rawDesc
)

func file_models_model_ad_package_instance_proto_rawDescGZIP() []byte {
	file_models_model_ad_package_instance_proto_rawDescOnce.Do(func() {
		file_models_model_ad_package_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ad_package_instance_proto_rawDescData)
	})
	return file_models_model_ad_package_instance_proto_rawDescData
}

var file_models_model_ad_package_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ad_package_instance_proto_goTypes = []interface{}{
	(*ADPackageInstance)(nil), // 0: pb.ADPackageInstance
	(*NodeCluster)(nil),       // 1: pb.NodeCluster
	(*ADPackage)(nil),         // 2: pb.ADPackage
	(*User)(nil),              // 3: pb.User
}
var file_models_model_ad_package_instance_proto_depIdxs = []int32{
	1, // 0: pb.ADPackageInstance.nodeCluster:type_name -> pb.NodeCluster
	2, // 1: pb.ADPackageInstance.adPackage:type_name -> pb.ADPackage
	3, // 2: pb.ADPackageInstance.user:type_name -> pb.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_ad_package_instance_proto_init() }
func file_models_model_ad_package_instance_proto_init() {
	if File_models_model_ad_package_instance_proto != nil {
		return
	}
	file_models_model_node_cluster_proto_init()
	file_models_model_ad_package_proto_init()
	file_models_model_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ad_package_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADPackageInstance); i {
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
			RawDescriptor: file_models_model_ad_package_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ad_package_instance_proto_goTypes,
		DependencyIndexes: file_models_model_ad_package_instance_proto_depIdxs,
		MessageInfos:      file_models_model_ad_package_instance_proto_msgTypes,
	}.Build()
	File_models_model_ad_package_instance_proto = out.File
	file_models_model_ad_package_instance_proto_rawDesc = nil
	file_models_model_ad_package_instance_proto_goTypes = nil
	file_models_model_ad_package_instance_proto_depIdxs = nil
}
