// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: creds.proto

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

type CredsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid           *uint32  `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Gid           *uint32  `protobuf:"varint,2,req,name=gid" json:"gid,omitempty"`
	Euid          *uint32  `protobuf:"varint,3,req,name=euid" json:"euid,omitempty"`
	Egid          *uint32  `protobuf:"varint,4,req,name=egid" json:"egid,omitempty"`
	Suid          *uint32  `protobuf:"varint,5,req,name=suid" json:"suid,omitempty"`
	Sgid          *uint32  `protobuf:"varint,6,req,name=sgid" json:"sgid,omitempty"`
	Fsuid         *uint32  `protobuf:"varint,7,req,name=fsuid" json:"fsuid,omitempty"`
	Fsgid         *uint32  `protobuf:"varint,8,req,name=fsgid" json:"fsgid,omitempty"`
	CapInh        []uint32 `protobuf:"varint,9,rep,name=cap_inh,json=capInh" json:"cap_inh,omitempty"`
	CapPrm        []uint32 `protobuf:"varint,10,rep,name=cap_prm,json=capPrm" json:"cap_prm,omitempty"`
	CapEff        []uint32 `protobuf:"varint,11,rep,name=cap_eff,json=capEff" json:"cap_eff,omitempty"`
	CapBnd        []uint32 `protobuf:"varint,12,rep,name=cap_bnd,json=capBnd" json:"cap_bnd,omitempty"`
	Secbits       *uint32  `protobuf:"varint,13,req,name=secbits" json:"secbits,omitempty"`
	Groups        []uint32 `protobuf:"varint,14,rep,name=groups" json:"groups,omitempty"`
	LsmProfile    *string  `protobuf:"bytes,15,opt,name=lsm_profile,json=lsmProfile" json:"lsm_profile,omitempty"`
	LsmSockcreate *string  `protobuf:"bytes,16,opt,name=lsm_sockcreate,json=lsmSockcreate" json:"lsm_sockcreate,omitempty"`
}

func (x *CredsEntry) Reset() {
	*x = CredsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredsEntry) ProtoMessage() {}

func (x *CredsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_creds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredsEntry.ProtoReflect.Descriptor instead.
func (*CredsEntry) Descriptor() ([]byte, []int) {
	return file_creds_proto_rawDescGZIP(), []int{0}
}

func (x *CredsEntry) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *CredsEntry) GetGid() uint32 {
	if x != nil && x.Gid != nil {
		return *x.Gid
	}
	return 0
}

func (x *CredsEntry) GetEuid() uint32 {
	if x != nil && x.Euid != nil {
		return *x.Euid
	}
	return 0
}

func (x *CredsEntry) GetEgid() uint32 {
	if x != nil && x.Egid != nil {
		return *x.Egid
	}
	return 0
}

func (x *CredsEntry) GetSuid() uint32 {
	if x != nil && x.Suid != nil {
		return *x.Suid
	}
	return 0
}

func (x *CredsEntry) GetSgid() uint32 {
	if x != nil && x.Sgid != nil {
		return *x.Sgid
	}
	return 0
}

func (x *CredsEntry) GetFsuid() uint32 {
	if x != nil && x.Fsuid != nil {
		return *x.Fsuid
	}
	return 0
}

func (x *CredsEntry) GetFsgid() uint32 {
	if x != nil && x.Fsgid != nil {
		return *x.Fsgid
	}
	return 0
}

func (x *CredsEntry) GetCapInh() []uint32 {
	if x != nil {
		return x.CapInh
	}
	return nil
}

func (x *CredsEntry) GetCapPrm() []uint32 {
	if x != nil {
		return x.CapPrm
	}
	return nil
}

func (x *CredsEntry) GetCapEff() []uint32 {
	if x != nil {
		return x.CapEff
	}
	return nil
}

func (x *CredsEntry) GetCapBnd() []uint32 {
	if x != nil {
		return x.CapBnd
	}
	return nil
}

func (x *CredsEntry) GetSecbits() uint32 {
	if x != nil && x.Secbits != nil {
		return *x.Secbits
	}
	return 0
}

func (x *CredsEntry) GetGroups() []uint32 {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *CredsEntry) GetLsmProfile() string {
	if x != nil && x.LsmProfile != nil {
		return *x.LsmProfile
	}
	return ""
}

func (x *CredsEntry) GetLsmSockcreate() string {
	if x != nil && x.LsmSockcreate != nil {
		return *x.LsmSockcreate
	}
	return ""
}

var File_creds_proto protoreflect.FileDescriptor

var file_creds_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x67, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x04, 0x65, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x67, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x04, 0x65, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x67, 0x69, 0x64, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x67, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x05, 0x66, 0x73, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x67, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x73, 0x67, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x68, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x63, 0x61, 0x70, 0x49, 0x6e, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x5f, 0x70, 0x72,
	0x6d, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x70, 0x50, 0x72, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x5f, 0x65, 0x66, 0x66, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x06, 0x63, 0x61, 0x70, 0x45, 0x66, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x5f,
	0x62, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x70, 0x42, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x62, 0x69, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x63, 0x62, 0x69, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x73, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x73, 0x6d, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x73, 0x6d, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x73,
	0x6d, 0x53, 0x6f, 0x63, 0x6b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
}

var (
	file_creds_proto_rawDescOnce sync.Once
	file_creds_proto_rawDescData = file_creds_proto_rawDesc
)

func file_creds_proto_rawDescGZIP() []byte {
	file_creds_proto_rawDescOnce.Do(func() {
		file_creds_proto_rawDescData = protoimpl.X.CompressGZIP(file_creds_proto_rawDescData)
	})
	return file_creds_proto_rawDescData
}

var file_creds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_creds_proto_goTypes = []interface{}{
	(*CredsEntry)(nil), // 0: creds_entry
}
var file_creds_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_creds_proto_init() }
func file_creds_proto_init() {
	if File_creds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredsEntry); i {
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
			RawDescriptor: file_creds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creds_proto_goTypes,
		DependencyIndexes: file_creds_proto_depIdxs,
		MessageInfos:      file_creds_proto_msgTypes,
	}.Build()
	File_creds_proto = out.File
	file_creds_proto_rawDesc = nil
	file_creds_proto_goTypes = nil
	file_creds_proto_depIdxs = nil
}
