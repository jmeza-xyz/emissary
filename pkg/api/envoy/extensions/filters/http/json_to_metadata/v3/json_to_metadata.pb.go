// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/filters/http/json_to_metadata/v3/json_to_metadata.proto

package json_to_metadatav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JsonToMetadata_ValueType int32

const (
	// The value is a serialized `protobuf.Value
	// <https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/struct.proto#L62>`_.
	JsonToMetadata_PROTOBUF_VALUE JsonToMetadata_ValueType = 0
	JsonToMetadata_STRING         JsonToMetadata_ValueType = 1
	JsonToMetadata_NUMBER         JsonToMetadata_ValueType = 2
)

// Enum value maps for JsonToMetadata_ValueType.
var (
	JsonToMetadata_ValueType_name = map[int32]string{
		0: "PROTOBUF_VALUE",
		1: "STRING",
		2: "NUMBER",
	}
	JsonToMetadata_ValueType_value = map[string]int32{
		"PROTOBUF_VALUE": 0,
		"STRING":         1,
		"NUMBER":         2,
	}
)

func (x JsonToMetadata_ValueType) Enum() *JsonToMetadata_ValueType {
	p := new(JsonToMetadata_ValueType)
	*p = x
	return p
}

func (x JsonToMetadata_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JsonToMetadata_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_enumTypes[0].Descriptor()
}

func (JsonToMetadata_ValueType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_enumTypes[0]
}

func (x JsonToMetadata_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JsonToMetadata_ValueType.Descriptor instead.
func (JsonToMetadata_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type JsonToMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// At least one of request_rules and response_rules must be provided.
	// Rules to match json body of requests.
	RequestRules *JsonToMetadata_MatchRules `protobuf:"bytes,1,opt,name=request_rules,json=requestRules,proto3" json:"request_rules,omitempty"`
	// Rules to match json body of responses.
	ResponseRules *JsonToMetadata_MatchRules `protobuf:"bytes,2,opt,name=response_rules,json=responseRules,proto3" json:"response_rules,omitempty"`
}

func (x *JsonToMetadata) Reset() {
	*x = JsonToMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonToMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonToMetadata) ProtoMessage() {}

func (x *JsonToMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonToMetadata.ProtoReflect.Descriptor instead.
func (*JsonToMetadata) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *JsonToMetadata) GetRequestRules() *JsonToMetadata_MatchRules {
	if x != nil {
		return x.RequestRules
	}
	return nil
}

func (x *JsonToMetadata) GetResponseRules() *JsonToMetadata_MatchRules {
	if x != nil {
		return x.ResponseRules
	}
	return nil
}

// [#next-free-field: 6]
type JsonToMetadata_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to ValueType:
	//
	//	*JsonToMetadata_KeyValuePair_Value
	ValueType isJsonToMetadata_KeyValuePair_ValueType `protobuf_oneof:"value_type"`
	// The value's type — defaults to protobuf.Value.
	Type JsonToMetadata_ValueType `protobuf:"varint,4,opt,name=type,proto3,enum=envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata_ValueType" json:"type,omitempty"`
	// False if we want to overwrite the existing metadata value. Default to false.
	PreserveExistingMetadataValue bool `protobuf:"varint,5,opt,name=preserve_existing_metadata_value,json=preserveExistingMetadataValue,proto3" json:"preserve_existing_metadata_value,omitempty"`
}

func (x *JsonToMetadata_KeyValuePair) Reset() {
	*x = JsonToMetadata_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonToMetadata_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonToMetadata_KeyValuePair) ProtoMessage() {}

func (x *JsonToMetadata_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonToMetadata_KeyValuePair.ProtoReflect.Descriptor instead.
func (*JsonToMetadata_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JsonToMetadata_KeyValuePair) GetMetadataNamespace() string {
	if x != nil {
		return x.MetadataNamespace
	}
	return ""
}

func (x *JsonToMetadata_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *JsonToMetadata_KeyValuePair) GetValueType() isJsonToMetadata_KeyValuePair_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (x *JsonToMetadata_KeyValuePair) GetValue() *structpb.Value {
	if x, ok := x.GetValueType().(*JsonToMetadata_KeyValuePair_Value); ok {
		return x.Value
	}
	return nil
}

func (x *JsonToMetadata_KeyValuePair) GetType() JsonToMetadata_ValueType {
	if x != nil {
		return x.Type
	}
	return JsonToMetadata_PROTOBUF_VALUE
}

func (x *JsonToMetadata_KeyValuePair) GetPreserveExistingMetadataValue() bool {
	if x != nil {
		return x.PreserveExistingMetadataValue
	}
	return false
}

type isJsonToMetadata_KeyValuePair_ValueType interface {
	isJsonToMetadata_KeyValuePair_ValueType()
}

type JsonToMetadata_KeyValuePair_Value struct {
	// The value to pair with the given key.
	//
	// When used for on_present case, if value is non-empty it'll be used instead
	// of the the value of the JSON key. If both are empty, the the value of the
	// JSON key is used as-is.
	//
	// When used for on_missing/on_error case, a non-empty value
	// must be provided.
	//
	// It ignores ValueType, i.e., not type casting.
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

func (*JsonToMetadata_KeyValuePair_Value) isJsonToMetadata_KeyValuePair_ValueType() {}

type JsonToMetadata_Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Selector:
	//
	//	*JsonToMetadata_Selector_Key
	Selector isJsonToMetadata_Selector_Selector `protobuf_oneof:"selector"`
}

func (x *JsonToMetadata_Selector) Reset() {
	*x = JsonToMetadata_Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonToMetadata_Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonToMetadata_Selector) ProtoMessage() {}

func (x *JsonToMetadata_Selector) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonToMetadata_Selector.ProtoReflect.Descriptor instead.
func (*JsonToMetadata_Selector) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0, 1}
}

func (m *JsonToMetadata_Selector) GetSelector() isJsonToMetadata_Selector_Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (x *JsonToMetadata_Selector) GetKey() string {
	if x, ok := x.GetSelector().(*JsonToMetadata_Selector_Key); ok {
		return x.Key
	}
	return ""
}

type isJsonToMetadata_Selector_Selector interface {
	isJsonToMetadata_Selector_Selector()
}

type JsonToMetadata_Selector_Key struct {
	// key to match
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*JsonToMetadata_Selector_Key) isJsonToMetadata_Selector_Selector() {}

// A Rule defines what metadata to apply when a key-value is present, missing in the json
// or fail to parse the payload.
type JsonToMetadata_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies that a match will be performed on the value of a property.
	// Here's an example to match on 1 in {"foo": {"bar": 1}, "bar": 2}
	//
	// selectors:
	// - key: foo
	// - key: bar
	Selectors []*JsonToMetadata_Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// If the attribute is present, apply this metadata KeyValuePair.
	OnPresent *JsonToMetadata_KeyValuePair `protobuf:"bytes,2,opt,name=on_present,json=onPresent,proto3" json:"on_present,omitempty"`
	// If the attribute is missing, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set.
	OnMissing *JsonToMetadata_KeyValuePair `protobuf:"bytes,3,opt,name=on_missing,json=onMissing,proto3" json:"on_missing,omitempty"`
	// If the body is too large or fail to parse or content-type is mismatched, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set.
	OnError *JsonToMetadata_KeyValuePair `protobuf:"bytes,4,opt,name=on_error,json=onError,proto3" json:"on_error,omitempty"`
}

func (x *JsonToMetadata_Rule) Reset() {
	*x = JsonToMetadata_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonToMetadata_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonToMetadata_Rule) ProtoMessage() {}

func (x *JsonToMetadata_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonToMetadata_Rule.ProtoReflect.Descriptor instead.
func (*JsonToMetadata_Rule) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0, 2}
}

func (x *JsonToMetadata_Rule) GetSelectors() []*JsonToMetadata_Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

func (x *JsonToMetadata_Rule) GetOnPresent() *JsonToMetadata_KeyValuePair {
	if x != nil {
		return x.OnPresent
	}
	return nil
}

func (x *JsonToMetadata_Rule) GetOnMissing() *JsonToMetadata_KeyValuePair {
	if x != nil {
		return x.OnMissing
	}
	return nil
}

func (x *JsonToMetadata_Rule) GetOnError() *JsonToMetadata_KeyValuePair {
	if x != nil {
		return x.OnError
	}
	return nil
}

type JsonToMetadata_MatchRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules to apply.
	Rules []*JsonToMetadata_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// Allowed content-type for json to metadata transformation.
	// Default to “{"application/json"}“.
	//
	// Set “allow_empty_content_type“ if empty/missing content-type header
	// is allowed.
	AllowContentTypes []string `protobuf:"bytes,2,rep,name=allow_content_types,json=allowContentTypes,proto3" json:"allow_content_types,omitempty"`
	// Allowed empty content-type for json to metadata transformation.
	// Default to false.
	AllowEmptyContentType bool `protobuf:"varint,3,opt,name=allow_empty_content_type,json=allowEmptyContentType,proto3" json:"allow_empty_content_type,omitempty"`
	// Allowed content-type by regex match for json to metadata transformation.
	// This can be used in parallel with “allow_content_types“.
	AllowContentTypesRegex *v3.RegexMatcher `protobuf:"bytes,4,opt,name=allow_content_types_regex,json=allowContentTypesRegex,proto3" json:"allow_content_types_regex,omitempty"`
}

func (x *JsonToMetadata_MatchRules) Reset() {
	*x = JsonToMetadata_MatchRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonToMetadata_MatchRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonToMetadata_MatchRules) ProtoMessage() {}

func (x *JsonToMetadata_MatchRules) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonToMetadata_MatchRules.ProtoReflect.Descriptor instead.
func (*JsonToMetadata_MatchRules) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP(), []int{0, 3}
}

func (x *JsonToMetadata_MatchRules) GetRules() []*JsonToMetadata_Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *JsonToMetadata_MatchRules) GetAllowContentTypes() []string {
	if x != nil {
		return x.AllowContentTypes
	}
	return nil
}

func (x *JsonToMetadata_MatchRules) GetAllowEmptyContentType() bool {
	if x != nil {
		return x.AllowEmptyContentType
	}
	return false
}

func (x *JsonToMetadata_MatchRules) GetAllowContentTypesRegex() *v3.RegexMatcher {
	if x != nil {
		return x.AllowContentTypesRegex
	}
	return nil
}

var File_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x76, 0x33, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74,
	0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x21, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x0b, 0x0a, 0x0e, 0x4a, 0x73, 0x6f, 0x6e, 0x54,
	0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x71, 0x0a, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x1a, 0xca, 0x02, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x69, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a,
	0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x33,
	0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0xc3, 0x03, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x72, 0x0a, 0x09,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x6d, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12,
	0x6d, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x09, 0x6f, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x69,
	0x0a, 0x08, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x07, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xcb, 0x02, 0x0a, 0x0a, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x66, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x73, 0x6f, 0x6e,
	0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x3c, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0xfa,
	0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x37,
	0x0a, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x67, 0x65, 0x78, 0x22, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x42, 0x55, 0x46,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02,
	0x42, 0xcd, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3f, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x6f,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x4a, 0x73,
	0x6f, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74,
	0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x3b, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescData = file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDesc
)

func file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDescData
}

var file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_goTypes = []interface{}{
	(JsonToMetadata_ValueType)(0),       // 0: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.ValueType
	(*JsonToMetadata)(nil),              // 1: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata
	(*JsonToMetadata_KeyValuePair)(nil), // 2: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair
	(*JsonToMetadata_Selector)(nil),     // 3: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Selector
	(*JsonToMetadata_Rule)(nil),         // 4: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule
	(*JsonToMetadata_MatchRules)(nil),   // 5: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.MatchRules
	(*structpb.Value)(nil),              // 6: google.protobuf.Value
	(*v3.RegexMatcher)(nil),             // 7: envoy.type.matcher.v3.RegexMatcher
}
var file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_depIdxs = []int32{
	5,  // 0: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.request_rules:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.MatchRules
	5,  // 1: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.response_rules:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.MatchRules
	6,  // 2: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair.value:type_name -> google.protobuf.Value
	0,  // 3: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair.type:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.ValueType
	3,  // 4: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule.selectors:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Selector
	2,  // 5: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule.on_present:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair
	2,  // 6: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule.on_missing:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair
	2,  // 7: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule.on_error:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.KeyValuePair
	4,  // 8: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.MatchRules.rules:type_name -> envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.Rule
	7,  // 9: envoy.extensions.filters.http.json_to_metadata.v3.JsonToMetadata.MatchRules.allow_content_types_regex:type_name -> envoy.type.matcher.v3.RegexMatcher
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_init() }
func file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_init() {
	if File_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonToMetadata); i {
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
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonToMetadata_KeyValuePair); i {
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
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonToMetadata_Selector); i {
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
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonToMetadata_Rule); i {
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
		file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonToMetadata_MatchRules); i {
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
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JsonToMetadata_KeyValuePair_Value)(nil),
	}
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*JsonToMetadata_Selector_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto = out.File
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_rawDesc = nil
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_goTypes = nil
	file_envoy_extensions_filters_http_json_to_metadata_v3_json_to_metadata_proto_depIdxs = nil
}
