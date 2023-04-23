// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/v1/cache.proto

package v1

import (
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

type KVRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVRecord) Reset() {
	*x = KVRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVRecord) ProtoMessage() {}

func (x *KVRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVRecord.ProtoReflect.Descriptor instead.
func (*KVRecord) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{0}
}

func (x *KVRecord) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVRecord) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *KVRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Ttl    int32     `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{1}
}

func (x *StoreRequest) GetRecord() *KVRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *StoreRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{2}
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{3}
}

func (x *FetchRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{4}
}

func (x *FetchResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AllKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern string `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *AllKeysRequest) Reset() {
	*x = AllKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllKeysRequest) ProtoMessage() {}

func (x *AllKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllKeysRequest.ProtoReflect.Descriptor instead.
func (*AllKeysRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{5}
}

func (x *AllKeysRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

type AllKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *AllKeysResponse) Reset() {
	*x = AllKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllKeysResponse) ProtoMessage() {}

func (x *AllKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllKeysResponse.ProtoReflect.Descriptor instead.
func (*AllKeysResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{6}
}

func (x *AllKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type AllRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *AllRecordsRequest) Reset() {
	*x = AllRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRecordsRequest) ProtoMessage() {}

func (x *AllRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRecordsRequest.ProtoReflect.Descriptor instead.
func (*AllRecordsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{7}
}

func (x *AllRecordsRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type AllRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *KVRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Ttl    int32     `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *AllRecordsResponse) Reset() {
	*x = AllRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cache_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRecordsResponse) ProtoMessage() {}

func (x *AllRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cache_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRecordsResponse.ProtoReflect.Descriptor instead.
func (*AllRecordsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cache_proto_rawDescGZIP(), []int{8}
}

func (x *AllRecordsResponse) GetRecord() *KVRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *AllRecordsResponse) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

var File_api_v1_cache_proto protoreflect.FileDescriptor

var file_api_v1_cache_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x22, 0x32, 0x0a, 0x08,
	0x4b, 0x56, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x4a, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x0f, 0x0a, 0x0d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a,
	0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x25, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x22, 0x50, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x31, 0x2e, 0x4b, 0x56, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x74, 0x74, 0x6c, 0x32, 0xf6, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x34,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x41, 0x6c,
	0x6c, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x7a, 0x61, 0x61,
	0x6b, 0x64, 0x61, 0x6c, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_cache_proto_rawDescOnce sync.Once
	file_api_v1_cache_proto_rawDescData = file_api_v1_cache_proto_rawDesc
)

func file_api_v1_cache_proto_rawDescGZIP() []byte {
	file_api_v1_cache_proto_rawDescOnce.Do(func() {
		file_api_v1_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_cache_proto_rawDescData)
	})
	return file_api_v1_cache_proto_rawDescData
}

var file_api_v1_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_cache_proto_goTypes = []interface{}{
	(*KVRecord)(nil),           // 0: api_v1.KVRecord
	(*StoreRequest)(nil),       // 1: api_v1.StoreRequest
	(*StoreResponse)(nil),      // 2: api_v1.StoreResponse
	(*FetchRequest)(nil),       // 3: api_v1.FetchRequest
	(*FetchResponse)(nil),      // 4: api_v1.FetchResponse
	(*AllKeysRequest)(nil),     // 5: api_v1.AllKeysRequest
	(*AllKeysResponse)(nil),    // 6: api_v1.AllKeysResponse
	(*AllRecordsRequest)(nil),  // 7: api_v1.AllRecordsRequest
	(*AllRecordsResponse)(nil), // 8: api_v1.AllRecordsResponse
}
var file_api_v1_cache_proto_depIdxs = []int32{
	0, // 0: api_v1.StoreRequest.record:type_name -> api_v1.KVRecord
	0, // 1: api_v1.AllRecordsResponse.record:type_name -> api_v1.KVRecord
	1, // 2: api_v1.Cache.Store:input_type -> api_v1.StoreRequest
	3, // 3: api_v1.Cache.Fetch:input_type -> api_v1.FetchRequest
	5, // 4: api_v1.Cache.AllKeys:input_type -> api_v1.AllKeysRequest
	7, // 5: api_v1.Cache.AllRecords:input_type -> api_v1.AllRecordsRequest
	2, // 6: api_v1.Cache.Store:output_type -> api_v1.StoreResponse
	4, // 7: api_v1.Cache.Fetch:output_type -> api_v1.FetchResponse
	6, // 8: api_v1.Cache.AllKeys:output_type -> api_v1.AllKeysResponse
	8, // 9: api_v1.Cache.AllRecords:output_type -> api_v1.AllRecordsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_cache_proto_init() }
func file_api_v1_cache_proto_init() {
	if File_api_v1_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVRecord); i {
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
		file_api_v1_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_api_v1_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
		file_api_v1_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_api_v1_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_api_v1_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllKeysRequest); i {
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
		file_api_v1_cache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllKeysResponse); i {
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
		file_api_v1_cache_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRecordsRequest); i {
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
		file_api_v1_cache_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRecordsResponse); i {
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
			RawDescriptor: file_api_v1_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_cache_proto_goTypes,
		DependencyIndexes: file_api_v1_cache_proto_depIdxs,
		MessageInfos:      file_api_v1_cache_proto_msgTypes,
	}.Build()
	File_api_v1_cache_proto = out.File
	file_api_v1_cache_proto_rawDesc = nil
	file_api_v1_cache_proto_goTypes = nil
	file_api_v1_cache_proto_depIdxs = nil
}
