// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: remote-image.proto

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

type LocalImageEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SnapshotId *string `protobuf:"bytes,2,req,name=snapshot_id,json=snapshotId" json:"snapshot_id,omitempty"`
	OpenMode   *uint32 `protobuf:"varint,3,req,name=open_mode,json=openMode" json:"open_mode,omitempty"`
}

func (x *LocalImageEntry) Reset() {
	*x = LocalImageEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalImageEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalImageEntry) ProtoMessage() {}

func (x *LocalImageEntry) ProtoReflect() protoreflect.Message {
	mi := &file_remote_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalImageEntry.ProtoReflect.Descriptor instead.
func (*LocalImageEntry) Descriptor() ([]byte, []int) {
	return file_remote_image_proto_rawDescGZIP(), []int{0}
}

func (x *LocalImageEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LocalImageEntry) GetSnapshotId() string {
	if x != nil && x.SnapshotId != nil {
		return *x.SnapshotId
	}
	return ""
}

func (x *LocalImageEntry) GetOpenMode() uint32 {
	if x != nil && x.OpenMode != nil {
		return *x.OpenMode
	}
	return 0
}

type RemoteImageEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SnapshotId *string `protobuf:"bytes,2,req,name=snapshot_id,json=snapshotId" json:"snapshot_id,omitempty"`
	OpenMode   *uint32 `protobuf:"varint,3,req,name=open_mode,json=openMode" json:"open_mode,omitempty"`
	Size       *uint64 `protobuf:"varint,4,req,name=size" json:"size,omitempty"`
}

func (x *RemoteImageEntry) Reset() {
	*x = RemoteImageEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteImageEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteImageEntry) ProtoMessage() {}

func (x *RemoteImageEntry) ProtoReflect() protoreflect.Message {
	mi := &file_remote_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteImageEntry.ProtoReflect.Descriptor instead.
func (*RemoteImageEntry) Descriptor() ([]byte, []int) {
	return file_remote_image_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteImageEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RemoteImageEntry) GetSnapshotId() string {
	if x != nil && x.SnapshotId != nil {
		return *x.SnapshotId
	}
	return ""
}

func (x *RemoteImageEntry) GetOpenMode() uint32 {
	if x != nil && x.OpenMode != nil {
		return *x.OpenMode
	}
	return 0
}

func (x *RemoteImageEntry) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

type LocalImageReplyEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *uint32 `protobuf:"varint,1,req,name=error" json:"error,omitempty"`
}

func (x *LocalImageReplyEntry) Reset() {
	*x = LocalImageReplyEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalImageReplyEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalImageReplyEntry) ProtoMessage() {}

func (x *LocalImageReplyEntry) ProtoReflect() protoreflect.Message {
	mi := &file_remote_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalImageReplyEntry.ProtoReflect.Descriptor instead.
func (*LocalImageReplyEntry) Descriptor() ([]byte, []int) {
	return file_remote_image_proto_rawDescGZIP(), []int{2}
}

func (x *LocalImageReplyEntry) GetError() uint32 {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return 0
}

type SnapshotIdEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapshotId *string `protobuf:"bytes,1,req,name=snapshot_id,json=snapshotId" json:"snapshot_id,omitempty"`
}

func (x *SnapshotIdEntry) Reset() {
	*x = SnapshotIdEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotIdEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotIdEntry) ProtoMessage() {}

func (x *SnapshotIdEntry) ProtoReflect() protoreflect.Message {
	mi := &file_remote_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotIdEntry.ProtoReflect.Descriptor instead.
func (*SnapshotIdEntry) Descriptor() ([]byte, []int) {
	return file_remote_image_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotIdEntry) GetSnapshotId() string {
	if x != nil && x.SnapshotId != nil {
		return *x.SnapshotId
	}
	return ""
}

var File_remote_image_proto protoreflect.FileDescriptor

var file_remote_image_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x7a, 0x0a, 0x12, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2f, 0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x11, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
}

var (
	file_remote_image_proto_rawDescOnce sync.Once
	file_remote_image_proto_rawDescData = file_remote_image_proto_rawDesc
)

func file_remote_image_proto_rawDescGZIP() []byte {
	file_remote_image_proto_rawDescOnce.Do(func() {
		file_remote_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_image_proto_rawDescData)
	})
	return file_remote_image_proto_rawDescData
}

var file_remote_image_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_image_proto_goTypes = []interface{}{
	(*LocalImageEntry)(nil),      // 0: local_image_entry
	(*RemoteImageEntry)(nil),     // 1: remote_image_entry
	(*LocalImageReplyEntry)(nil), // 2: local_image_reply_entry
	(*SnapshotIdEntry)(nil),      // 3: snapshot_id_entry
}
var file_remote_image_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_remote_image_proto_init() }
func file_remote_image_proto_init() {
	if File_remote_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalImageEntry); i {
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
		file_remote_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteImageEntry); i {
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
		file_remote_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalImageReplyEntry); i {
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
		file_remote_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotIdEntry); i {
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
			RawDescriptor: file_remote_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_image_proto_goTypes,
		DependencyIndexes: file_remote_image_proto_depIdxs,
		MessageInfos:      file_remote_image_proto_msgTypes,
	}.Build()
	File_remote_image_proto = out.File
	file_remote_image_proto_rawDesc = nil
	file_remote_image_proto_goTypes = nil
	file_remote_image_proto_depIdxs = nil
}
