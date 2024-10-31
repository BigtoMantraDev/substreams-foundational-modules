// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/sf/substreams/solana/type/v1/type.proto

package typev1

import (
	v1 "github.com/streamingfast/substreams-foundational-modules/solana-accounts-common/pb/sf/solana/type/v1"
	_ "github.com/streamingfast/substreams-foundational-modules/solana-accounts-common/pb/sf/substreams/index/v1"
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

type FilteredAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*v1.Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *FilteredAccounts) Reset() {
	*x = FilteredAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sf_substreams_solana_type_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilteredAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteredAccounts) ProtoMessage() {}

func (x *FilteredAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sf_substreams_solana_type_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteredAccounts.ProtoReflect.Descriptor instead.
func (*FilteredAccounts) Descriptor() ([]byte, []int) {
	return file_proto_sf_substreams_solana_type_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *FilteredAccounts) GetAccounts() []*v1.Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_proto_sf_substreams_solana_type_v1_type_proto protoreflect.FileDescriptor

var file_proto_sf_substreams_solana_type_v1_type_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x73,
	0x66, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73, 0x66, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x10, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0xc0, 0x02, 0x0a, 0x20, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x7c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x53, 0x53, 0x53,
	0x54, 0xaa, 0x02, 0x1c, 0x53, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1c, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x5c, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x28, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x5c,
	0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x53, 0x66, 0x3a,
	0x3a, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x3a, 0x53, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sf_substreams_solana_type_v1_type_proto_rawDescOnce sync.Once
	file_proto_sf_substreams_solana_type_v1_type_proto_rawDescData = file_proto_sf_substreams_solana_type_v1_type_proto_rawDesc
)

func file_proto_sf_substreams_solana_type_v1_type_proto_rawDescGZIP() []byte {
	file_proto_sf_substreams_solana_type_v1_type_proto_rawDescOnce.Do(func() {
		file_proto_sf_substreams_solana_type_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sf_substreams_solana_type_v1_type_proto_rawDescData)
	})
	return file_proto_sf_substreams_solana_type_v1_type_proto_rawDescData
}

var file_proto_sf_substreams_solana_type_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_sf_substreams_solana_type_v1_type_proto_goTypes = []interface{}{
	(*FilteredAccounts)(nil), // 0: sf.substreams.solana.type.v1.FilteredAccounts
	(*v1.Account)(nil),       // 1: sf.solana.type.v1.Account
}
var file_proto_sf_substreams_solana_type_v1_type_proto_depIdxs = []int32{
	1, // 0: sf.substreams.solana.type.v1.FilteredAccounts.accounts:type_name -> sf.solana.type.v1.Account
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_sf_substreams_solana_type_v1_type_proto_init() }
func file_proto_sf_substreams_solana_type_v1_type_proto_init() {
	if File_proto_sf_substreams_solana_type_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sf_substreams_solana_type_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilteredAccounts); i {
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
			RawDescriptor: file_proto_sf_substreams_solana_type_v1_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_sf_substreams_solana_type_v1_type_proto_goTypes,
		DependencyIndexes: file_proto_sf_substreams_solana_type_v1_type_proto_depIdxs,
		MessageInfos:      file_proto_sf_substreams_solana_type_v1_type_proto_msgTypes,
	}.Build()
	File_proto_sf_substreams_solana_type_v1_type_proto = out.File
	file_proto_sf_substreams_solana_type_v1_type_proto_rawDesc = nil
	file_proto_sf_substreams_solana_type_v1_type_proto_goTypes = nil
	file_proto_sf_substreams_solana_type_v1_type_proto_depIdxs = nil
}