// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.0--rc1
// source: discovery.proto

package api

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

type DiscoverPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscoverPeersRequest) Reset() {
	*x = DiscoverPeersRequest{}
	mi := &file_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverPeersRequest) ProtoMessage() {}

func (x *DiscoverPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverPeersRequest.ProtoReflect.Descriptor instead.
func (*DiscoverPeersRequest) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{0}
}

type DiscoverPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *DiscoverPeersResponse) Reset() {
	*x = DiscoverPeersResponse{}
	mi := &file_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverPeersResponse) ProtoMessage() {}

func (x *DiscoverPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverPeersResponse.ProtoReflect.Descriptor instead.
func (*DiscoverPeersResponse) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoverPeersResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_discovery_proto protoreflect.FileDescriptor

var file_discovery_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x15,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x32, 0x53, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discovery_proto_rawDescOnce sync.Once
	file_discovery_proto_rawDescData = file_discovery_proto_rawDesc
)

func file_discovery_proto_rawDescGZIP() []byte {
	file_discovery_proto_rawDescOnce.Do(func() {
		file_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_proto_rawDescData)
	})
	return file_discovery_proto_rawDescData
}

var file_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discovery_proto_goTypes = []any{
	(*DiscoverPeersRequest)(nil),  // 0: api.DiscoverPeersRequest
	(*DiscoverPeersResponse)(nil), // 1: api.DiscoverPeersResponse
	(*Peer)(nil),                  // 2: api.Peer
}
var file_discovery_proto_depIdxs = []int32{
	2, // 0: api.DiscoverPeersResponse.peers:type_name -> api.Peer
	0, // 1: api.Discovery.DiscoverPeers:input_type -> api.DiscoverPeersRequest
	1, // 2: api.Discovery.DiscoverPeers:output_type -> api.DiscoverPeersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_discovery_proto_init() }
func file_discovery_proto_init() {
	if File_discovery_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discovery_proto_goTypes,
		DependencyIndexes: file_discovery_proto_depIdxs,
		MessageInfos:      file_discovery_proto_msgTypes,
	}.Build()
	File_discovery_proto = out.File
	file_discovery_proto_rawDesc = nil
	file_discovery_proto_goTypes = nil
	file_discovery_proto_depIdxs = nil
}
