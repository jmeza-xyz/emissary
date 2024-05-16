// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/early_data/v3/default_early_data_policy.proto

package early_datav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// [#extension: envoy.route.early_data_policy.default]
// The default rule to allow/disallow a request to be sent as early data. It's an empty config now. Configuring it will disallow any request to be sent over early data.
type DefaultEarlyDataPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DefaultEarlyDataPolicy) Reset() {
	*x = DefaultEarlyDataPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_early_data_v3_default_early_data_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultEarlyDataPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultEarlyDataPolicy) ProtoMessage() {}

func (x *DefaultEarlyDataPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_early_data_v3_default_early_data_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultEarlyDataPolicy.ProtoReflect.Descriptor instead.
func (*DefaultEarlyDataPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescGZIP(), []int{0}
}

var File_envoy_extensions_early_data_v3_default_early_data_policy_proto protoreflect.FileDescriptor

var file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33,
	0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x18, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0xa9, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x0a, 0x2c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x33, 0x42, 0x1b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x61, 0x72, 0x6c, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x61, 0x72, 0x6c, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x3b, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescOnce sync.Once
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescData = file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDesc
)

func file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescData)
	})
	return file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDescData
}

var file_envoy_extensions_early_data_v3_default_early_data_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_early_data_v3_default_early_data_policy_proto_goTypes = []interface{}{
	(*DefaultEarlyDataPolicy)(nil), // 0: envoy.extensions.early_data.v3.DefaultEarlyDataPolicy
}
var file_envoy_extensions_early_data_v3_default_early_data_policy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_early_data_v3_default_early_data_policy_proto_init() }
func file_envoy_extensions_early_data_v3_default_early_data_policy_proto_init() {
	if File_envoy_extensions_early_data_v3_default_early_data_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_early_data_v3_default_early_data_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultEarlyDataPolicy); i {
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
			RawDescriptor: file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_early_data_v3_default_early_data_policy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_early_data_v3_default_early_data_policy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_early_data_v3_default_early_data_policy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_early_data_v3_default_early_data_policy_proto = out.File
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_rawDesc = nil
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_goTypes = nil
	file_envoy_extensions_early_data_v3_default_early_data_policy_proto_depIdxs = nil
}
