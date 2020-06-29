// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: packet-sock.proto

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

type PacketMclist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *uint32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Type  *uint32 `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	Addr  []byte  `protobuf:"bytes,3,req,name=addr" json:"addr,omitempty"`
}

func (x *PacketMclist) Reset() {
	*x = PacketMclist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_sock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketMclist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketMclist) ProtoMessage() {}

func (x *PacketMclist) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketMclist.ProtoReflect.Descriptor instead.
func (*PacketMclist) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{0}
}

func (x *PacketMclist) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *PacketMclist) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *PacketMclist) GetAddr() []byte {
	if x != nil {
		return x.Addr
	}
	return nil
}

type PacketRing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockSize  *uint32 `protobuf:"varint,1,req,name=block_size,json=blockSize" json:"block_size,omitempty"`
	BlockNr    *uint32 `protobuf:"varint,2,req,name=block_nr,json=blockNr" json:"block_nr,omitempty"`
	FrameSize  *uint32 `protobuf:"varint,3,req,name=frame_size,json=frameSize" json:"frame_size,omitempty"`
	FrameNr    *uint32 `protobuf:"varint,4,req,name=frame_nr,json=frameNr" json:"frame_nr,omitempty"`
	RetireTmo  *uint32 `protobuf:"varint,5,req,name=retire_tmo,json=retireTmo" json:"retire_tmo,omitempty"`
	SizeofPriv *uint32 `protobuf:"varint,6,req,name=sizeof_priv,json=sizeofPriv" json:"sizeof_priv,omitempty"`
	Features   *uint32 `protobuf:"varint,7,req,name=features" json:"features,omitempty"`
}

func (x *PacketRing) Reset() {
	*x = PacketRing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_sock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketRing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketRing) ProtoMessage() {}

func (x *PacketRing) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketRing.ProtoReflect.Descriptor instead.
func (*PacketRing) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{1}
}

func (x *PacketRing) GetBlockSize() uint32 {
	if x != nil && x.BlockSize != nil {
		return *x.BlockSize
	}
	return 0
}

func (x *PacketRing) GetBlockNr() uint32 {
	if x != nil && x.BlockNr != nil {
		return *x.BlockNr
	}
	return 0
}

func (x *PacketRing) GetFrameSize() uint32 {
	if x != nil && x.FrameSize != nil {
		return *x.FrameSize
	}
	return 0
}

func (x *PacketRing) GetFrameNr() uint32 {
	if x != nil && x.FrameNr != nil {
		return *x.FrameNr
	}
	return 0
}

func (x *PacketRing) GetRetireTmo() uint32 {
	if x != nil && x.RetireTmo != nil {
		return *x.RetireTmo
	}
	return 0
}

func (x *PacketRing) GetSizeofPriv() uint32 {
	if x != nil && x.SizeofPriv != nil {
		return *x.SizeofPriv
	}
	return 0
}

func (x *PacketRing) GetFeatures() uint32 {
	if x != nil && x.Features != nil {
		return *x.Features
	}
	return 0
}

type PacketSockEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint32         `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type       *uint32         `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	Protocol   *uint32         `protobuf:"varint,3,req,name=protocol" json:"protocol,omitempty"`
	Flags      *uint32         `protobuf:"varint,4,req,name=flags" json:"flags,omitempty"`
	Ifindex    *uint32         `protobuf:"varint,5,req,name=ifindex" json:"ifindex,omitempty"`
	Fown       *FownEntry      `protobuf:"bytes,6,req,name=fown" json:"fown,omitempty"`
	Opts       *SkOptsEntry    `protobuf:"bytes,7,req,name=opts" json:"opts,omitempty"`
	Version    *uint32         `protobuf:"varint,8,req,name=version" json:"version,omitempty"`
	Reserve    *uint32         `protobuf:"varint,9,req,name=reserve" json:"reserve,omitempty"`
	AuxData    *bool           `protobuf:"varint,10,req,name=aux_data,json=auxData" json:"aux_data,omitempty"`
	OrigDev    *bool           `protobuf:"varint,11,req,name=orig_dev,json=origDev" json:"orig_dev,omitempty"`
	VnetHdr    *bool           `protobuf:"varint,12,req,name=vnet_hdr,json=vnetHdr" json:"vnet_hdr,omitempty"`
	Loss       *bool           `protobuf:"varint,13,req,name=loss" json:"loss,omitempty"`
	Timestamp  *uint32         `protobuf:"varint,14,req,name=timestamp" json:"timestamp,omitempty"`
	CopyThresh *uint32         `protobuf:"varint,15,req,name=copy_thresh,json=copyThresh" json:"copy_thresh,omitempty"`
	Mclist     []*PacketMclist `protobuf:"bytes,16,rep,name=mclist" json:"mclist,omitempty"`
	Fanout     *uint32         `protobuf:"varint,17,opt,name=fanout,def=4294967295" json:"fanout,omitempty"`
	RxRing     *PacketRing     `protobuf:"bytes,18,opt,name=rx_ring,json=rxRing" json:"rx_ring,omitempty"`
	TxRing     *PacketRing     `protobuf:"bytes,19,opt,name=tx_ring,json=txRing" json:"tx_ring,omitempty"`
	NsId       *uint32         `protobuf:"varint,20,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
}

// Default values for PacketSockEntry fields.
const (
	Default_PacketSockEntry_Fanout = uint32(4294967295)
)

func (x *PacketSockEntry) Reset() {
	*x = PacketSockEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_sock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketSockEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketSockEntry) ProtoMessage() {}

func (x *PacketSockEntry) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketSockEntry.ProtoReflect.Descriptor instead.
func (*PacketSockEntry) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{2}
}

func (x *PacketSockEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PacketSockEntry) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *PacketSockEntry) GetProtocol() uint32 {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return 0
}

func (x *PacketSockEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *PacketSockEntry) GetIfindex() uint32 {
	if x != nil && x.Ifindex != nil {
		return *x.Ifindex
	}
	return 0
}

func (x *PacketSockEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *PacketSockEntry) GetOpts() *SkOptsEntry {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *PacketSockEntry) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *PacketSockEntry) GetReserve() uint32 {
	if x != nil && x.Reserve != nil {
		return *x.Reserve
	}
	return 0
}

func (x *PacketSockEntry) GetAuxData() bool {
	if x != nil && x.AuxData != nil {
		return *x.AuxData
	}
	return false
}

func (x *PacketSockEntry) GetOrigDev() bool {
	if x != nil && x.OrigDev != nil {
		return *x.OrigDev
	}
	return false
}

func (x *PacketSockEntry) GetVnetHdr() bool {
	if x != nil && x.VnetHdr != nil {
		return *x.VnetHdr
	}
	return false
}

func (x *PacketSockEntry) GetLoss() bool {
	if x != nil && x.Loss != nil {
		return *x.Loss
	}
	return false
}

func (x *PacketSockEntry) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *PacketSockEntry) GetCopyThresh() uint32 {
	if x != nil && x.CopyThresh != nil {
		return *x.CopyThresh
	}
	return 0
}

func (x *PacketSockEntry) GetMclist() []*PacketMclist {
	if x != nil {
		return x.Mclist
	}
	return nil
}

func (x *PacketSockEntry) GetFanout() uint32 {
	if x != nil && x.Fanout != nil {
		return *x.Fanout
	}
	return Default_PacketSockEntry_Fanout
}

func (x *PacketSockEntry) GetRxRing() *PacketRing {
	if x != nil {
		return x.RxRing
	}
	return nil
}

func (x *PacketSockEntry) GetTxRing() *PacketRing {
	if x != nil {
		return x.TxRing
	}
	return nil
}

func (x *PacketSockEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

var File_packet_sock_proto protoreflect.FileDescriptor

var file_packet_sock_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2d, 0x73, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x6b, 0x2d,
	0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0d, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x63, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6e, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x72, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x09, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x54, 0x6d, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x69, 0x7a, 0x65, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x6f, 0x66, 0x50, 0x72, 0x69, 0x76, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xd6, 0x04, 0x0a, 0x11, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x1b, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x05,
	0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x69,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x18, 0x06,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x18, 0x09, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x75, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x02, 0x28,
	0x08, 0x52, 0x07, 0x61, 0x75, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x5f, 0x64, 0x65, 0x76, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x72,
	0x69, 0x67, 0x44, 0x65, 0x76, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6e, 0x65, 0x74, 0x5f, 0x68, 0x64,
	0x72, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x76, 0x6e, 0x65, 0x74, 0x48, 0x64, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x08, 0x52, 0x04,
	0x6c, 0x6f, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x0e, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x70, 0x79, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x63, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x63, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x06, 0x6d, 0x63, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x66,
	0x61, 0x6e, 0x6f, 0x75, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x0a, 0x34, 0x32, 0x39,
	0x34, 0x39, 0x36, 0x37, 0x32, 0x39, 0x35, 0x52, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x12,
	0x25, 0x0a, 0x07, 0x72, 0x78, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x72, 0x78, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x78, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x0a,
	0x05, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x73,
	0x49, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
}

var (
	file_packet_sock_proto_rawDescOnce sync.Once
	file_packet_sock_proto_rawDescData = file_packet_sock_proto_rawDesc
)

func file_packet_sock_proto_rawDescGZIP() []byte {
	file_packet_sock_proto_rawDescOnce.Do(func() {
		file_packet_sock_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_sock_proto_rawDescData)
	})
	return file_packet_sock_proto_rawDescData
}

var file_packet_sock_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packet_sock_proto_goTypes = []interface{}{
	(*PacketMclist)(nil),    // 0: packet_mclist
	(*PacketRing)(nil),      // 1: packet_ring
	(*PacketSockEntry)(nil), // 2: packet_sock_entry
	(*FownEntry)(nil),       // 3: fown_entry
	(*SkOptsEntry)(nil),     // 4: sk_opts_entry
}
var file_packet_sock_proto_depIdxs = []int32{
	3, // 0: packet_sock_entry.fown:type_name -> fown_entry
	4, // 1: packet_sock_entry.opts:type_name -> sk_opts_entry
	0, // 2: packet_sock_entry.mclist:type_name -> packet_mclist
	1, // 3: packet_sock_entry.rx_ring:type_name -> packet_ring
	1, // 4: packet_sock_entry.tx_ring:type_name -> packet_ring
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_packet_sock_proto_init() }
func file_packet_sock_proto_init() {
	if File_packet_sock_proto != nil {
		return
	}
	file_opts_proto_init()
	file_fown_proto_init()
	file_sk_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_packet_sock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketMclist); i {
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
		file_packet_sock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketRing); i {
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
		file_packet_sock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketSockEntry); i {
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
			RawDescriptor: file_packet_sock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_sock_proto_goTypes,
		DependencyIndexes: file_packet_sock_proto_depIdxs,
		MessageInfos:      file_packet_sock_proto_msgTypes,
	}.Build()
	File_packet_sock_proto = out.File
	file_packet_sock_proto_rawDesc = nil
	file_packet_sock_proto_goTypes = nil
	file_packet_sock_proto_depIdxs = nil
}
