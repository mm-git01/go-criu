// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: core-arm.proto

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

type UserArmRegsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R0     *uint32 `protobuf:"varint,1,req,name=r0" json:"r0,omitempty"`
	R1     *uint32 `protobuf:"varint,2,req,name=r1" json:"r1,omitempty"`
	R2     *uint32 `protobuf:"varint,3,req,name=r2" json:"r2,omitempty"`
	R3     *uint32 `protobuf:"varint,4,req,name=r3" json:"r3,omitempty"`
	R4     *uint32 `protobuf:"varint,5,req,name=r4" json:"r4,omitempty"`
	R5     *uint32 `protobuf:"varint,6,req,name=r5" json:"r5,omitempty"`
	R6     *uint32 `protobuf:"varint,7,req,name=r6" json:"r6,omitempty"`
	R7     *uint32 `protobuf:"varint,8,req,name=r7" json:"r7,omitempty"`
	R8     *uint32 `protobuf:"varint,9,req,name=r8" json:"r8,omitempty"`
	R9     *uint32 `protobuf:"varint,10,req,name=r9" json:"r9,omitempty"`
	R10    *uint32 `protobuf:"varint,11,req,name=r10" json:"r10,omitempty"`
	Fp     *uint32 `protobuf:"varint,12,req,name=fp" json:"fp,omitempty"`
	Ip     *uint32 `protobuf:"varint,13,req,name=ip" json:"ip,omitempty"`
	Sp     *uint32 `protobuf:"varint,14,req,name=sp" json:"sp,omitempty"`
	Lr     *uint32 `protobuf:"varint,15,req,name=lr" json:"lr,omitempty"`
	Pc     *uint32 `protobuf:"varint,16,req,name=pc" json:"pc,omitempty"`
	Cpsr   *uint32 `protobuf:"varint,17,req,name=cpsr" json:"cpsr,omitempty"`
	OrigR0 *uint32 `protobuf:"varint,18,req,name=orig_r0,json=origR0" json:"orig_r0,omitempty"`
}

func (x *UserArmRegsEntry) Reset() {
	*x = UserArmRegsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_arm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserArmRegsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserArmRegsEntry) ProtoMessage() {}

func (x *UserArmRegsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_arm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserArmRegsEntry.ProtoReflect.Descriptor instead.
func (*UserArmRegsEntry) Descriptor() ([]byte, []int) {
	return file_core_arm_proto_rawDescGZIP(), []int{0}
}

func (x *UserArmRegsEntry) GetR0() uint32 {
	if x != nil && x.R0 != nil {
		return *x.R0
	}
	return 0
}

func (x *UserArmRegsEntry) GetR1() uint32 {
	if x != nil && x.R1 != nil {
		return *x.R1
	}
	return 0
}

func (x *UserArmRegsEntry) GetR2() uint32 {
	if x != nil && x.R2 != nil {
		return *x.R2
	}
	return 0
}

func (x *UserArmRegsEntry) GetR3() uint32 {
	if x != nil && x.R3 != nil {
		return *x.R3
	}
	return 0
}

func (x *UserArmRegsEntry) GetR4() uint32 {
	if x != nil && x.R4 != nil {
		return *x.R4
	}
	return 0
}

func (x *UserArmRegsEntry) GetR5() uint32 {
	if x != nil && x.R5 != nil {
		return *x.R5
	}
	return 0
}

func (x *UserArmRegsEntry) GetR6() uint32 {
	if x != nil && x.R6 != nil {
		return *x.R6
	}
	return 0
}

func (x *UserArmRegsEntry) GetR7() uint32 {
	if x != nil && x.R7 != nil {
		return *x.R7
	}
	return 0
}

func (x *UserArmRegsEntry) GetR8() uint32 {
	if x != nil && x.R8 != nil {
		return *x.R8
	}
	return 0
}

func (x *UserArmRegsEntry) GetR9() uint32 {
	if x != nil && x.R9 != nil {
		return *x.R9
	}
	return 0
}

func (x *UserArmRegsEntry) GetR10() uint32 {
	if x != nil && x.R10 != nil {
		return *x.R10
	}
	return 0
}

func (x *UserArmRegsEntry) GetFp() uint32 {
	if x != nil && x.Fp != nil {
		return *x.Fp
	}
	return 0
}

func (x *UserArmRegsEntry) GetIp() uint32 {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return 0
}

func (x *UserArmRegsEntry) GetSp() uint32 {
	if x != nil && x.Sp != nil {
		return *x.Sp
	}
	return 0
}

func (x *UserArmRegsEntry) GetLr() uint32 {
	if x != nil && x.Lr != nil {
		return *x.Lr
	}
	return 0
}

func (x *UserArmRegsEntry) GetPc() uint32 {
	if x != nil && x.Pc != nil {
		return *x.Pc
	}
	return 0
}

func (x *UserArmRegsEntry) GetCpsr() uint32 {
	if x != nil && x.Cpsr != nil {
		return *x.Cpsr
	}
	return 0
}

func (x *UserArmRegsEntry) GetOrigR0() uint32 {
	if x != nil && x.OrigR0 != nil {
		return *x.OrigR0
	}
	return 0
}

type UserArmVfpstateEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VfpRegs []uint64 `protobuf:"varint,1,rep,name=vfp_regs,json=vfpRegs" json:"vfp_regs,omitempty"`
	Fpscr   *uint32  `protobuf:"varint,2,req,name=fpscr" json:"fpscr,omitempty"`
	Fpexc   *uint32  `protobuf:"varint,3,req,name=fpexc" json:"fpexc,omitempty"`
	Fpinst  *uint32  `protobuf:"varint,4,req,name=fpinst" json:"fpinst,omitempty"`
	Fpinst2 *uint32  `protobuf:"varint,5,req,name=fpinst2" json:"fpinst2,omitempty"`
}

func (x *UserArmVfpstateEntry) Reset() {
	*x = UserArmVfpstateEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_arm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserArmVfpstateEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserArmVfpstateEntry) ProtoMessage() {}

func (x *UserArmVfpstateEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_arm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserArmVfpstateEntry.ProtoReflect.Descriptor instead.
func (*UserArmVfpstateEntry) Descriptor() ([]byte, []int) {
	return file_core_arm_proto_rawDescGZIP(), []int{1}
}

func (x *UserArmVfpstateEntry) GetVfpRegs() []uint64 {
	if x != nil {
		return x.VfpRegs
	}
	return nil
}

func (x *UserArmVfpstateEntry) GetFpscr() uint32 {
	if x != nil && x.Fpscr != nil {
		return *x.Fpscr
	}
	return 0
}

func (x *UserArmVfpstateEntry) GetFpexc() uint32 {
	if x != nil && x.Fpexc != nil {
		return *x.Fpexc
	}
	return 0
}

func (x *UserArmVfpstateEntry) GetFpinst() uint32 {
	if x != nil && x.Fpinst != nil {
		return *x.Fpinst
	}
	return 0
}

func (x *UserArmVfpstateEntry) GetFpinst2() uint32 {
	if x != nil && x.Fpinst2 != nil {
		return *x.Fpinst2
	}
	return 0
}

type ThreadInfoArm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearTidAddr *uint64               `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Tls          *uint32               `protobuf:"varint,2,req,name=tls" json:"tls,omitempty"`
	Gpregs       *UserArmRegsEntry     `protobuf:"bytes,3,req,name=gpregs" json:"gpregs,omitempty"`
	Fpstate      *UserArmVfpstateEntry `protobuf:"bytes,4,req,name=fpstate" json:"fpstate,omitempty"`
}

func (x *ThreadInfoArm) Reset() {
	*x = ThreadInfoArm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_arm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadInfoArm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfoArm) ProtoMessage() {}

func (x *ThreadInfoArm) ProtoReflect() protoreflect.Message {
	mi := &file_core_arm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfoArm.ProtoReflect.Descriptor instead.
func (*ThreadInfoArm) Descriptor() ([]byte, []int) {
	return file_core_arm_proto_rawDescGZIP(), []int{2}
}

func (x *ThreadInfoArm) GetClearTidAddr() uint64 {
	if x != nil && x.ClearTidAddr != nil {
		return *x.ClearTidAddr
	}
	return 0
}

func (x *ThreadInfoArm) GetTls() uint32 {
	if x != nil && x.Tls != nil {
		return *x.Tls
	}
	return 0
}

func (x *ThreadInfoArm) GetGpregs() *UserArmRegsEntry {
	if x != nil {
		return x.Gpregs
	}
	return nil
}

func (x *ThreadInfoArm) GetFpstate() *UserArmVfpstateEntry {
	if x != nil {
		return x.Fpstate
	}
	return nil
}

var File_core_arm_proto protoreflect.FileDescriptor

var file_core_arm_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a,
	0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x30, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x31, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x32, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x33, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x34, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x35, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x36, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x37, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x37, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x38, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x39, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x72, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x30, 0x18, 0x0b, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x03, 0x72, 0x31, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x70, 0x18, 0x0c, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x02, 0x66, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0d, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x70, 0x18, 0x0e, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x02, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x72, 0x18, 0x0f, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x02, 0x6c, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x63, 0x18, 0x10, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x02, 0x70, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70, 0x73, 0x72, 0x18, 0x11,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x70, 0x73, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x72,
	0x69, 0x67, 0x5f, 0x72, 0x30, 0x18, 0x12, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x52, 0x30, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6d,
	0x5f, 0x76, 0x66, 0x70, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x66, 0x70, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x07, 0x76, 0x66, 0x70, 0x52, 0x65, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x70,
	0x73, 0x63, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x70, 0x73, 0x63, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x70, 0x65, 0x78, 0x63, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x05, 0x66, 0x70, 0x65, 0x78, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x70, 0x69, 0x6e, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x70, 0x69, 0x6e, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x70, 0x69, 0x6e, 0x73, 0x74, 0x32, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x07, 0x66, 0x70, 0x69, 0x6e, 0x73, 0x74, 0x32, 0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x72, 0x6d, 0x12, 0x2b, 0x0a, 0x0e,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x69, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x54, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x67,
	0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73,
	0x12, 0x32, 0x0a, 0x07, 0x66, 0x70, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6d, 0x5f, 0x76, 0x66, 0x70,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x70, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c,
}

var (
	file_core_arm_proto_rawDescOnce sync.Once
	file_core_arm_proto_rawDescData = file_core_arm_proto_rawDesc
)

func file_core_arm_proto_rawDescGZIP() []byte {
	file_core_arm_proto_rawDescOnce.Do(func() {
		file_core_arm_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_arm_proto_rawDescData)
	})
	return file_core_arm_proto_rawDescData
}

var file_core_arm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_arm_proto_goTypes = []interface{}{
	(*UserArmRegsEntry)(nil),     // 0: user_arm_regs_entry
	(*UserArmVfpstateEntry)(nil), // 1: user_arm_vfpstate_entry
	(*ThreadInfoArm)(nil),        // 2: thread_info_arm
}
var file_core_arm_proto_depIdxs = []int32{
	0, // 0: thread_info_arm.gpregs:type_name -> user_arm_regs_entry
	1, // 1: thread_info_arm.fpstate:type_name -> user_arm_vfpstate_entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_arm_proto_init() }
func file_core_arm_proto_init() {
	if File_core_arm_proto != nil {
		return
	}
	file_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_arm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserArmRegsEntry); i {
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
		file_core_arm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserArmVfpstateEntry); i {
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
		file_core_arm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadInfoArm); i {
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
			RawDescriptor: file_core_arm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_arm_proto_goTypes,
		DependencyIndexes: file_core_arm_proto_depIdxs,
		MessageInfos:      file_core_arm_proto_msgTypes,
	}.Build()
	File_core_arm_proto = out.File
	file_core_arm_proto_rawDesc = nil
	file_core_arm_proto_goTypes = nil
	file_core_arm_proto_depIdxs = nil
}
