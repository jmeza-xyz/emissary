// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/transport_socket/tap/v2alpha/tap.proto

package envoy_config_transport_socket_tap_v2alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/v2/pkg/api/envoy/api/v2/core"
	v2alpha "github.com/datawire/ambassador/v2/pkg/api/envoy/config/common/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Tap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonConfig    *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	TransportSocket *core.TransportSocket          `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
}

func (x *Tap) Reset() {
	*x = Tap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_transport_socket_tap_v2alpha_tap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tap) ProtoMessage() {}

func (x *Tap) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_transport_socket_tap_v2alpha_tap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tap.ProtoReflect.Descriptor instead.
func (*Tap) Descriptor() ([]byte, []int) {
	return file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescGZIP(), []int{0}
}

func (x *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

func (x *Tap) GetTransportSocket() *core.TransportSocket {
	if x != nil {
		return x.TransportSocket
	}
	return nil
}

var File_envoy_config_transport_socket_tap_v2alpha_tap_proto protoreflect.FileDescriptor

var file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x74, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x70, 0x12, 0x65, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x86, 0x01, 0x0a,
	0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x08, 0x54, 0x61, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2b, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescOnce sync.Once
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescData = file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDesc
)

func file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescGZIP() []byte {
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescOnce.Do(func() {
		file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescData)
	})
	return file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDescData
}

var file_envoy_config_transport_socket_tap_v2alpha_tap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_transport_socket_tap_v2alpha_tap_proto_goTypes = []interface{}{
	(*Tap)(nil),                           // 0: envoy.config.transport_socket.tap.v2alpha.Tap
	(*v2alpha.CommonExtensionConfig)(nil), // 1: envoy.config.common.tap.v2alpha.CommonExtensionConfig
	(*core.TransportSocket)(nil),          // 2: envoy.api.v2.core.TransportSocket
}
var file_envoy_config_transport_socket_tap_v2alpha_tap_proto_depIdxs = []int32{
	1, // 0: envoy.config.transport_socket.tap.v2alpha.Tap.common_config:type_name -> envoy.config.common.tap.v2alpha.CommonExtensionConfig
	2, // 1: envoy.config.transport_socket.tap.v2alpha.Tap.transport_socket:type_name -> envoy.api.v2.core.TransportSocket
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_transport_socket_tap_v2alpha_tap_proto_init() }
func file_envoy_config_transport_socket_tap_v2alpha_tap_proto_init() {
	if File_envoy_config_transport_socket_tap_v2alpha_tap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_transport_socket_tap_v2alpha_tap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tap); i {
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
			RawDescriptor: file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_transport_socket_tap_v2alpha_tap_proto_goTypes,
		DependencyIndexes: file_envoy_config_transport_socket_tap_v2alpha_tap_proto_depIdxs,
		MessageInfos:      file_envoy_config_transport_socket_tap_v2alpha_tap_proto_msgTypes,
	}.Build()
	File_envoy_config_transport_socket_tap_v2alpha_tap_proto = out.File
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_rawDesc = nil
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_goTypes = nil
	file_envoy_config_transport_socket_tap_v2alpha_tap_proto_depIdxs = nil
}
