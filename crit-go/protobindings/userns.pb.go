// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: userns.proto

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

type UidGidExtent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First      *uint32 `protobuf:"varint,1,req,name=first" json:"first,omitempty"`
	LowerFirst *uint32 `protobuf:"varint,2,req,name=lower_first,json=lowerFirst" json:"lower_first,omitempty"`
	Count      *uint32 `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
}

func (x *UidGidExtent) Reset() {
	*x = UidGidExtent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidGidExtent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidGidExtent) ProtoMessage() {}

func (x *UidGidExtent) ProtoReflect() protoreflect.Message {
	mi := &file_userns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidGidExtent.ProtoReflect.Descriptor instead.
func (*UidGidExtent) Descriptor() ([]byte, []int) {
	return file_userns_proto_rawDescGZIP(), []int{0}
}

func (x *UidGidExtent) GetFirst() uint32 {
	if x != nil && x.First != nil {
		return *x.First
	}
	return 0
}

func (x *UidGidExtent) GetLowerFirst() uint32 {
	if x != nil && x.LowerFirst != nil {
		return *x.LowerFirst
	}
	return 0
}

func (x *UidGidExtent) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type UsernsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UidMap []*UidGidExtent `protobuf:"bytes,1,rep,name=uid_map,json=uidMap" json:"uid_map,omitempty"`
	GidMap []*UidGidExtent `protobuf:"bytes,2,rep,name=gid_map,json=gidMap" json:"gid_map,omitempty"`
}

func (x *UsernsEntry) Reset() {
	*x = UsernsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernsEntry) ProtoMessage() {}

func (x *UsernsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_userns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernsEntry.ProtoReflect.Descriptor instead.
func (*UsernsEntry) Descriptor() ([]byte, []int) {
	return file_userns_proto_rawDescGZIP(), []int{1}
}

func (x *UsernsEntry) GetUidMap() []*UidGidExtent {
	if x != nil {
		return x.UidMap
	}
	return nil
}

func (x *UsernsEntry) GetGidMap() []*UidGidExtent {
	if x != nil {
		return x.GidMap
	}
	return nil
}

var File_userns_proto protoreflect.FileDescriptor

var file_userns_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d,
	0x0a, 0x0e, 0x75, 0x69, 0x64, 0x5f, 0x67, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a,
	0x07, 0x75, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x75, 0x69, 0x64, 0x5f, 0x67, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x28, 0x0a, 0x07, 0x67, 0x69, 0x64, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x69, 0x64, 0x5f, 0x67,
	0x69, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x67, 0x69, 0x64, 0x4d, 0x61,
	0x70, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
}

var (
	file_userns_proto_rawDescOnce sync.Once
	file_userns_proto_rawDescData = file_userns_proto_rawDesc
)

func file_userns_proto_rawDescGZIP() []byte {
	file_userns_proto_rawDescOnce.Do(func() {
		file_userns_proto_rawDescData = protoimpl.X.CompressGZIP(file_userns_proto_rawDescData)
	})
	return file_userns_proto_rawDescData
}

var file_userns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userns_proto_goTypes = []interface{}{
	(*UidGidExtent)(nil), // 0: uid_gid_extent
	(*UsernsEntry)(nil),  // 1: userns_entry
}
var file_userns_proto_depIdxs = []int32{
	0, // 0: userns_entry.uid_map:type_name -> uid_gid_extent
	0, // 1: userns_entry.gid_map:type_name -> uid_gid_extent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userns_proto_init() }
func file_userns_proto_init() {
	if File_userns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidGidExtent); i {
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
		file_userns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernsEntry); i {
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
			RawDescriptor: file_userns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userns_proto_goTypes,
		DependencyIndexes: file_userns_proto_depIdxs,
		MessageInfos:      file_userns_proto_msgTypes,
	}.Build()
	File_userns_proto = out.File
	file_userns_proto_rawDesc = nil
	file_userns_proto_goTypes = nil
	file_userns_proto_depIdxs = nil
}
