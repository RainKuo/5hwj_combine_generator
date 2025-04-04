// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0
// source: proto/config_ClassicalCombines.proto

package classical_combine

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

type RivalType int32

const (
	RivalType_NORMAL_RIVAL_    RivalType = 0 // 常见有对抗
	RivalType_NORMAL_NO_RIVAL_ RivalType = 1 //常见无对抗
	RivalType_RARE_RIVAL_      RivalType = 2 //稀有有对抗
	RivalType_RARE_NO_RIVAL_   RivalType = 3 //稀有无对抗
)

// Enum value maps for RivalType.
var (
	RivalType_name = map[int32]string{
		0: "NORMAL_RIVAL_",
		1: "NORMAL_NO_RIVAL_",
		2: "RARE_RIVAL_",
		3: "RARE_NO_RIVAL_",
	}
	RivalType_value = map[string]int32{
		"NORMAL_RIVAL_":    0,
		"NORMAL_NO_RIVAL_": 1,
		"RARE_RIVAL_":      2,
		"RARE_NO_RIVAL_":   3,
	}
)

func (x RivalType) Enum() *RivalType {
	p := new(RivalType)
	*p = x
	return p
}

func (x RivalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RivalType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_ClassicalCombines_proto_enumTypes[0].Descriptor()
}

func (RivalType) Type() protoreflect.EnumType {
	return &file_proto_config_ClassicalCombines_proto_enumTypes[0]
}

func (x RivalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RivalType.Descriptor instead.
func (RivalType) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_ClassicalCombines_proto_rawDescGZIP(), []int{0}
}

type ClassicalCombineConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`                         //配置id
	ControlFlag    uint32 `protobuf:"varint,2,opt,name=ControlFlag,proto3" json:"ControlFlag,omitempty"`       //身份控制
	RivalType      uint32 `protobuf:"varint,3,opt,name=RivalType,proto3" json:"RivalType,omitempty"`           //炸弹类型
	Combine0       uint64 `protobuf:"varint,4,opt,name=Combine0,proto3" json:"Combine0,omitempty"`             //0号
	Combine1       uint64 `protobuf:"varint,5,opt,name=Combine1,proto3" json:"Combine1,omitempty"`             //1号
	Combine2       uint64 `protobuf:"varint,6,opt,name=Combine2,proto3" json:"Combine2,omitempty"`             //2号
	RemainBigCount uint32 `protobuf:"varint,7,opt,name=RemainBigCount,proto3" json:"RemainBigCount,omitempty"` //剩余大牌数量
}

func (x *ClassicalCombineConfig) Reset() {
	*x = ClassicalCombineConfig{}
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassicalCombineConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassicalCombineConfig) ProtoMessage() {}

func (x *ClassicalCombineConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassicalCombineConfig.ProtoReflect.Descriptor instead.
func (*ClassicalCombineConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_ClassicalCombines_proto_rawDescGZIP(), []int{0}
}

func (x *ClassicalCombineConfig) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ClassicalCombineConfig) GetControlFlag() uint32 {
	if x != nil {
		return x.ControlFlag
	}
	return 0
}

func (x *ClassicalCombineConfig) GetRivalType() uint32 {
	if x != nil {
		return x.RivalType
	}
	return 0
}

func (x *ClassicalCombineConfig) GetCombine0() uint64 {
	if x != nil {
		return x.Combine0
	}
	return 0
}

func (x *ClassicalCombineConfig) GetCombine1() uint64 {
	if x != nil {
		return x.Combine1
	}
	return 0
}

func (x *ClassicalCombineConfig) GetCombine2() uint64 {
	if x != nil {
		return x.Combine2
	}
	return 0
}

func (x *ClassicalCombineConfig) GetRemainBigCount() uint32 {
	if x != nil {
		return x.RemainBigCount
	}
	return 0
}

type ClassicalBaseCombine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BigCount    uint32 `protobuf:"varint,1,opt,name=BigCount,proto3" json:"BigCount,omitempty"`       //大牌数
	BombCount   uint32 `protobuf:"varint,2,opt,name=BombCount,proto3" json:"BombCount,omitempty"`     //炸弹数
	TripleCount uint32 `protobuf:"varint,3,opt,name=TripleCount,proto3" json:"TripleCount,omitempty"` //三张数
	PairCount   uint32 `protobuf:"varint,4,opt,name=PairCount,proto3" json:"PairCount,omitempty"`     //对子数
	SingleCount uint32 `protobuf:"varint,5,opt,name=SingleCount,proto3" json:"SingleCount,omitempty"` //单张数
}

func (x *ClassicalBaseCombine) Reset() {
	*x = ClassicalBaseCombine{}
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassicalBaseCombine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassicalBaseCombine) ProtoMessage() {}

func (x *ClassicalBaseCombine) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassicalBaseCombine.ProtoReflect.Descriptor instead.
func (*ClassicalBaseCombine) Descriptor() ([]byte, []int) {
	return file_proto_config_ClassicalCombines_proto_rawDescGZIP(), []int{1}
}

func (x *ClassicalBaseCombine) GetBigCount() uint32 {
	if x != nil {
		return x.BigCount
	}
	return 0
}

func (x *ClassicalBaseCombine) GetBombCount() uint32 {
	if x != nil {
		return x.BombCount
	}
	return 0
}

func (x *ClassicalBaseCombine) GetTripleCount() uint32 {
	if x != nil {
		return x.TripleCount
	}
	return 0
}

func (x *ClassicalBaseCombine) GetPairCount() uint32 {
	if x != nil {
		return x.PairCount
	}
	return 0
}

func (x *ClassicalBaseCombine) GetSingleCount() uint32 {
	if x != nil {
		return x.SingleCount
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs     []*ClassicalCombineConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`                            //牌库配置
	BaseConfigs []*ClassicalBaseCombine   `protobuf:"bytes,2,rep,name=base_configs,json=baseConfigs,proto3" json:"base_configs,omitempty"` //基础牌型
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_ClassicalCombines_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_config_ClassicalCombines_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetConfigs() []*ClassicalCombineConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *Config) GetBaseConfigs() []*ClassicalBaseCombine {
	if x != nil {
		return x.BaseConfigs
	}
	return nil
}

var File_proto_config_ClassicalCombines_proto protoreflect.FileDescriptor

var file_proto_config_ClassicalCombines_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x22, 0xe6, 0x01, 0x0a,
	0x18, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x52,
	0x69, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x52, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x65, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x65, 0x30, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x31, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x32, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x32, 0x12, 0x26, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x69, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x16, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x42, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x42, 0x6f, 0x6d, 0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x42, 0x6f, 0x6d, 0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72,
	0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x50, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa3, 0x01, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2a, 0x59, 0x0a, 0x09, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x5f,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x4e, 0x4f, 0x5f,
	0x52, 0x49, 0x56, 0x41, 0x4c, 0x5f, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x41, 0x52, 0x45,
	0x5f, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x5f, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x41, 0x52,
	0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x5f, 0x10, 0x03, 0x42, 0x1d, 0x5a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x75, 0x74, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_ClassicalCombines_proto_rawDescOnce sync.Once
	file_proto_config_ClassicalCombines_proto_rawDescData = file_proto_config_ClassicalCombines_proto_rawDesc
)

func file_proto_config_ClassicalCombines_proto_rawDescGZIP() []byte {
	file_proto_config_ClassicalCombines_proto_rawDescOnce.Do(func() {
		file_proto_config_ClassicalCombines_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_ClassicalCombines_proto_rawDescData)
	})
	return file_proto_config_ClassicalCombines_proto_rawDescData
}

var file_proto_config_ClassicalCombines_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_config_ClassicalCombines_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_config_ClassicalCombines_proto_goTypes = []any{
	(RivalType)(0),                 // 0: ClassicalCombines.pb.RivalType
	(*ClassicalCombineConfig)(nil), // 1: ClassicalCombines.pb.classical_combine_config
	(*ClassicalBaseCombine)(nil),   // 2: ClassicalCombines.pb.classical_base_combine
	(*Config)(nil),                 // 3: ClassicalCombines.pb.Config
}
var file_proto_config_ClassicalCombines_proto_depIdxs = []int32{
	1, // 0: ClassicalCombines.pb.Config.configs:type_name -> ClassicalCombines.pb.classical_combine_config
	2, // 1: ClassicalCombines.pb.Config.base_configs:type_name -> ClassicalCombines.pb.classical_base_combine
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_config_ClassicalCombines_proto_init() }
func file_proto_config_ClassicalCombines_proto_init() {
	if File_proto_config_ClassicalCombines_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_config_ClassicalCombines_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_ClassicalCombines_proto_goTypes,
		DependencyIndexes: file_proto_config_ClassicalCombines_proto_depIdxs,
		EnumInfos:         file_proto_config_ClassicalCombines_proto_enumTypes,
		MessageInfos:      file_proto_config_ClassicalCombines_proto_msgTypes,
	}.Build()
	File_proto_config_ClassicalCombines_proto = out.File
	file_proto_config_ClassicalCombines_proto_rawDesc = nil
	file_proto_config_ClassicalCombines_proto_goTypes = nil
	file_proto_config_ClassicalCombines_proto_depIdxs = nil
}
