// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sf/solana/type/v1/type.proto

package typev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot       uint64                 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Hash       string                 `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentSlot uint64                 `protobuf:"varint,3,opt,name=parent_slot,json=parentSlot,proto3" json:"parent_slot,omitempty"`
	ParentHash string                 `protobuf:"bytes,4,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Accounts   []*Account             `protobuf:"bytes,7,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountBlock) Reset() {
	*x = AccountBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_solana_type_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBlock) ProtoMessage() {}

func (x *AccountBlock) ProtoReflect() protoreflect.Message {
	mi := &file_sf_solana_type_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBlock.ProtoReflect.Descriptor instead.
func (*AccountBlock) Descriptor() ([]byte, []int) {
	return file_sf_solana_type_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *AccountBlock) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *AccountBlock) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *AccountBlock) GetParentSlot() uint64 {
	if x != nil {
		return x.ParentSlot
	}
	return 0
}

func (x *AccountBlock) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *AccountBlock) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AccountBlock) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owner   []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Deleted bool   `protobuf:"varint,7,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_solana_type_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_sf_solana_type_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_sf_solana_type_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Account) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Account) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Account) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_sf_solana_type_v1_type_proto protoreflect.FileDescriptor

var file_sf_solana_type_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x66, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x73, 0x66, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x73,
	0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x67, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0xf6, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x66, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x53, 0x54, 0xaa, 0x02, 0x11, 0x53, 0x66, 0x2e, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x53, 0x66, 0x5c, 0x53, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x53, 0x66, 0x5c,
	0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x66, 0x3a,
	0x3a, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_solana_type_v1_type_proto_rawDescOnce sync.Once
	file_sf_solana_type_v1_type_proto_rawDescData = file_sf_solana_type_v1_type_proto_rawDesc
)

func file_sf_solana_type_v1_type_proto_rawDescGZIP() []byte {
	file_sf_solana_type_v1_type_proto_rawDescOnce.Do(func() {
		file_sf_solana_type_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_solana_type_v1_type_proto_rawDescData)
	})
	return file_sf_solana_type_v1_type_proto_rawDescData
}

var file_sf_solana_type_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_solana_type_v1_type_proto_goTypes = []interface{}{
	(*AccountBlock)(nil),          // 0: sf.solana.type.v1.AccountBlock
	(*Account)(nil),               // 1: sf.solana.type.v1.Account
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_sf_solana_type_v1_type_proto_depIdxs = []int32{
	2, // 0: sf.solana.type.v1.AccountBlock.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: sf.solana.type.v1.AccountBlock.accounts:type_name -> sf.solana.type.v1.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sf_solana_type_v1_type_proto_init() }
func file_sf_solana_type_v1_type_proto_init() {
	if File_sf_solana_type_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_solana_type_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBlock); i {
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
		file_sf_solana_type_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
			RawDescriptor: file_sf_solana_type_v1_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_solana_type_v1_type_proto_goTypes,
		DependencyIndexes: file_sf_solana_type_v1_type_proto_depIdxs,
		MessageInfos:      file_sf_solana_type_v1_type_proto_msgTypes,
	}.Build()
	File_sf_solana_type_v1_type_proto = out.File
	file_sf_solana_type_v1_type_proto_rawDesc = nil
	file_sf_solana_type_v1_type_proto_goTypes = nil
	file_sf_solana_type_v1_type_proto_depIdxs = nil
}