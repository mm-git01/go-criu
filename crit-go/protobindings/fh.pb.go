// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: fh.proto

package protobindings

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

type FhEntrySizes int32

const (
	FhEntrySizes_min_entries FhEntrySizes = 16
)

// Enum value maps for FhEntrySizes.
var (
	FhEntrySizes_name = map[int32]string{
		16: "min_entries",
	}
	FhEntrySizes_value = map[string]int32{
		"min_entries": 16,
	}
)

func (x FhEntrySizes) Enum() *FhEntrySizes {
	p := new(FhEntrySizes)
	*p = x
	return p
}

func (x FhEntrySizes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FhEntrySizes) Descriptor() protoreflect.EnumDescriptor {
	return file_fh_proto_enumTypes[0].Descriptor()
}

func (FhEntrySizes) Type() protoreflect.EnumType {
	return &file_fh_proto_enumTypes[0]
}

func (x FhEntrySizes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FhEntrySizes) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FhEntrySizes(num)
	return nil
}

// Deprecated: Use FhEntrySizes.Descriptor instead.
func (FhEntrySizes) EnumDescriptor() ([]byte, []int) {
	return file_fh_proto_rawDescGZIP(), []int{0}
}

type FhEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes *uint32 `protobuf:"varint,1,req,name=bytes" json:"bytes,omitempty"`
	Type  *uint32 `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	// The minimum is fh_n_handle repetitions
	Handle []uint64 `protobuf:"varint,3,rep,name=handle" json:"handle,omitempty"`
	Path   *string  `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	MntId  *uint32  `protobuf:"varint,5,opt,name=mnt_id,json=mntId" json:"mnt_id,omitempty"`
}

func (x *FhEntry) Reset() {
	*x = FhEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FhEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FhEntry) ProtoMessage() {}

func (x *FhEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FhEntry.ProtoReflect.Descriptor instead.
func (*FhEntry) Descriptor() ([]byte, []int) {
	return file_fh_proto_rawDescGZIP(), []int{0}
}

func (x *FhEntry) GetBytes() uint32 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

func (x *FhEntry) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *FhEntry) GetHandle() []uint64 {
	if x != nil {
		return x.Handle
	}
	return nil
}

func (x *FhEntry) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *FhEntry) GetMntId() uint32 {
	if x != nil && x.MntId != nil {
		return *x.MntId
	}
	return 0
}

type IrmapCacheEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dev   *uint32 `protobuf:"varint,1,req,name=dev" json:"dev,omitempty"`
	Inode *uint64 `protobuf:"varint,2,req,name=inode" json:"inode,omitempty"`
	Path  *string `protobuf:"bytes,3,req,name=path" json:"path,omitempty"`
}

func (x *IrmapCacheEntry) Reset() {
	*x = IrmapCacheEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IrmapCacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IrmapCacheEntry) ProtoMessage() {}

func (x *IrmapCacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IrmapCacheEntry.ProtoReflect.Descriptor instead.
func (*IrmapCacheEntry) Descriptor() ([]byte, []int) {
	return file_fh_proto_rawDescGZIP(), []int{1}
}

func (x *IrmapCacheEntry) GetDev() uint32 {
	if x != nil && x.Dev != nil {
		return *x.Dev
	}
	return 0
}

func (x *IrmapCacheEntry) GetInode() uint64 {
	if x != nil && x.Inode != nil {
		return *x.Inode
	}
	return 0
}

func (x *IrmapCacheEntry) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

var File_fh_proto protoreflect.FileDescriptor

var file_fh_proto_rawDesc = []byte{
	0x0a, 0x08, 0x66, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x08, 0x66, 0x68, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x5b, 0x0a, 0x11, 0x69, 0x72, 0x6d, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x42, 0x0a, 0xd2, 0x3f, 0x02, 0x20, 0x01, 0xd2, 0x3f, 0x02, 0x28, 0x01, 0x52, 0x03, 0x64,
	0x65, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x04, 0x52, 0x05, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x21, 0x0a, 0x0e,
	0x66, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x10, 0x10, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
}

var (
	file_fh_proto_rawDescOnce sync.Once
	file_fh_proto_rawDescData = file_fh_proto_rawDesc
)

func file_fh_proto_rawDescGZIP() []byte {
	file_fh_proto_rawDescOnce.Do(func() {
		file_fh_proto_rawDescData = protoimpl.X.CompressGZIP(file_fh_proto_rawDescData)
	})
	return file_fh_proto_rawDescData
}

var file_fh_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fh_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fh_proto_goTypes = []interface{}{
	(FhEntrySizes)(0),       // 0: fh_entry_sizes
	(*FhEntry)(nil),         // 1: fh_entry
	(*IrmapCacheEntry)(nil), // 2: irmap_cache_entry
}
var file_fh_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fh_proto_init() }
func file_fh_proto_init() {
	if File_fh_proto != nil {
		return
	}
	file_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FhEntry); i {
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
		file_fh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IrmapCacheEntry); i {
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
			RawDescriptor: file_fh_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fh_proto_goTypes,
		DependencyIndexes: file_fh_proto_depIdxs,
		EnumInfos:         file_fh_proto_enumTypes,
		MessageInfos:      file_fh_proto_msgTypes,
	}.Build()
	File_fh_proto = out.File
	file_fh_proto_rawDesc = nil
	file_fh_proto_goTypes = nil
	file_fh_proto_depIdxs = nil
}
