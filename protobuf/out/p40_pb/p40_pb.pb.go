// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: p40_pb.proto

package p40_pb

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

type Cs_40001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System     *uint32  `protobuf:"varint,1,req,name=system" json:"system,omitempty"`
	ShipIdList []uint32 `protobuf:"varint,2,rep,name=ship_id_list,json=shipIdList" json:"ship_id_list,omitempty"`
	Data       *uint32  `protobuf:"varint,3,req,name=data" json:"data,omitempty"`
}

func (x *Cs_40001) Reset() {
	*x = Cs_40001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_40001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_40001) ProtoMessage() {}

func (x *Cs_40001) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_40001.ProtoReflect.Descriptor instead.
func (*Cs_40001) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Cs_40001) GetSystem() uint32 {
	if x != nil && x.System != nil {
		return *x.System
	}
	return 0
}

func (x *Cs_40001) GetShipIdList() []uint32 {
	if x != nil {
		return x.ShipIdList
	}
	return nil
}

func (x *Cs_40001) GetData() uint32 {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return 0
}

type Sc_40002 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result          *uint32            `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Key             *uint32            `protobuf:"varint,2,opt,name=key" json:"key,omitempty"`
	DropPerformance []*Dropperformance `protobuf:"bytes,3,rep,name=drop_performance,json=dropPerformance" json:"drop_performance,omitempty"`
}

func (x *Sc_40002) Reset() {
	*x = Sc_40002{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_40002) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_40002) ProtoMessage() {}

func (x *Sc_40002) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_40002.ProtoReflect.Descriptor instead.
func (*Sc_40002) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{1}
}

func (x *Sc_40002) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *Sc_40002) GetKey() uint32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *Sc_40002) GetDropPerformance() []*Dropperformance {
	if x != nil {
		return x.DropPerformance
	}
	return nil
}

type Cs_40003 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System     *uint32           `protobuf:"varint,1,req,name=system" json:"system,omitempty"`
	Data       *uint32           `protobuf:"varint,2,req,name=data" json:"data,omitempty"`
	Key        *uint32           `protobuf:"varint,3,req,name=key" json:"key,omitempty"`
	Score      *uint32           `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
	Statistics []*Statisticsinfo `protobuf:"bytes,5,rep,name=statistics" json:"statistics,omitempty"`
	KillIdList []uint32          `protobuf:"varint,6,rep,name=kill_id_list,json=killIdList" json:"kill_id_list,omitempty"`
	TotalTime  *uint32           `protobuf:"varint,7,req,name=total_time,json=totalTime" json:"total_time,omitempty"`
}

func (x *Cs_40003) Reset() {
	*x = Cs_40003{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_40003) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_40003) ProtoMessage() {}

func (x *Cs_40003) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_40003.ProtoReflect.Descriptor instead.
func (*Cs_40003) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Cs_40003) GetSystem() uint32 {
	if x != nil && x.System != nil {
		return *x.System
	}
	return 0
}

func (x *Cs_40003) GetData() uint32 {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return 0
}

func (x *Cs_40003) GetKey() uint32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *Cs_40003) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *Cs_40003) GetStatistics() []*Statisticsinfo {
	if x != nil {
		return x.Statistics
	}
	return nil
}

func (x *Cs_40003) GetKillIdList() []uint32 {
	if x != nil {
		return x.KillIdList
	}
	return nil
}

func (x *Cs_40003) GetTotalTime() uint32 {
	if x != nil && x.TotalTime != nil {
		return *x.TotalTime
	}
	return 0
}

type Sc_40004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result      *uint32     `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	DropInfo    []*Dropinfo `protobuf:"bytes,2,rep,name=drop_info,json=dropInfo" json:"drop_info,omitempty"`
	PlayerExp   *uint32     `protobuf:"varint,3,req,name=player_exp,json=playerExp" json:"player_exp,omitempty"`
	ShipExpList []*ShipExp  `protobuf:"bytes,4,rep,name=ship_exp_list,json=shipExpList" json:"ship_exp_list,omitempty"`
}

func (x *Sc_40004) Reset() {
	*x = Sc_40004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_40004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_40004) ProtoMessage() {}

func (x *Sc_40004) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_40004.ProtoReflect.Descriptor instead.
func (*Sc_40004) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Sc_40004) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *Sc_40004) GetDropInfo() []*Dropinfo {
	if x != nil {
		return x.DropInfo
	}
	return nil
}

func (x *Sc_40004) GetPlayerExp() uint32 {
	if x != nil && x.PlayerExp != nil {
		return *x.PlayerExp
	}
	return 0
}

func (x *Sc_40004) GetShipExpList() []*ShipExp {
	if x != nil {
		return x.ShipExpList
	}
	return nil
}

type Cs_40005 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System *uint32 `protobuf:"varint,1,req,name=system" json:"system,omitempty"`
}

func (x *Cs_40005) Reset() {
	*x = Cs_40005{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_40005) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_40005) ProtoMessage() {}

func (x *Cs_40005) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_40005.ProtoReflect.Descriptor instead.
func (*Cs_40005) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Cs_40005) GetSystem() uint32 {
	if x != nil && x.System != nil {
		return *x.System
	}
	return 0
}

type Sc_40006 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
}

func (x *Sc_40006) Reset() {
	*x = Sc_40006{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_40006) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_40006) ProtoMessage() {}

func (x *Sc_40006) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_40006.ProtoReflect.Descriptor instead.
func (*Sc_40006) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{5}
}

func (x *Sc_40006) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

type Dropperformance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnemyId     *uint32 `protobuf:"varint,1,req,name=enemy_id,json=enemyId" json:"enemy_id,omitempty"`
	ResourceNum *uint32 `protobuf:"varint,2,req,name=resource_num,json=resourceNum" json:"resource_num,omitempty"`
	OtherNum    *uint32 `protobuf:"varint,3,req,name=other_num,json=otherNum" json:"other_num,omitempty"`
}

func (x *Dropperformance) Reset() {
	*x = Dropperformance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dropperformance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dropperformance) ProtoMessage() {}

func (x *Dropperformance) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dropperformance.ProtoReflect.Descriptor instead.
func (*Dropperformance) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{6}
}

func (x *Dropperformance) GetEnemyId() uint32 {
	if x != nil && x.EnemyId != nil {
		return *x.EnemyId
	}
	return 0
}

func (x *Dropperformance) GetResourceNum() uint32 {
	if x != nil && x.ResourceNum != nil {
		return *x.ResourceNum
	}
	return 0
}

func (x *Dropperformance) GetOtherNum() uint32 {
	if x != nil && x.OtherNum != nil {
		return *x.OtherNum
	}
	return 0
}

type Dropinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   *uint32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Id     *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Number *uint32 `protobuf:"varint,3,req,name=number" json:"number,omitempty"`
}

func (x *Dropinfo) Reset() {
	*x = Dropinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dropinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dropinfo) ProtoMessage() {}

func (x *Dropinfo) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dropinfo.ProtoReflect.Descriptor instead.
func (*Dropinfo) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{7}
}

func (x *Dropinfo) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *Dropinfo) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Dropinfo) GetNumber() uint32 {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return 0
}

type Statisticsinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipId        *uint32 `protobuf:"varint,1,req,name=ship_id,json=shipId" json:"ship_id,omitempty"`
	DamageCause   *uint32 `protobuf:"varint,2,req,name=damage_cause,json=damageCause" json:"damage_cause,omitempty"`
	DamageCaused  *uint32 `protobuf:"varint,3,req,name=damage_caused,json=damageCaused" json:"damage_caused,omitempty"`
	HpRest        *uint32 `protobuf:"varint,4,req,name=hp_rest,json=hpRest" json:"hp_rest,omitempty"`
	MaxDamageOnce *uint32 `protobuf:"varint,5,req,name=max_damage_once,json=maxDamageOnce" json:"max_damage_once,omitempty"`
}

func (x *Statisticsinfo) Reset() {
	*x = Statisticsinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statisticsinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statisticsinfo) ProtoMessage() {}

func (x *Statisticsinfo) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statisticsinfo.ProtoReflect.Descriptor instead.
func (*Statisticsinfo) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{8}
}

func (x *Statisticsinfo) GetShipId() uint32 {
	if x != nil && x.ShipId != nil {
		return *x.ShipId
	}
	return 0
}

func (x *Statisticsinfo) GetDamageCause() uint32 {
	if x != nil && x.DamageCause != nil {
		return *x.DamageCause
	}
	return 0
}

func (x *Statisticsinfo) GetDamageCaused() uint32 {
	if x != nil && x.DamageCaused != nil {
		return *x.DamageCaused
	}
	return 0
}

func (x *Statisticsinfo) GetHpRest() uint32 {
	if x != nil && x.HpRest != nil {
		return *x.HpRest
	}
	return 0
}

func (x *Statisticsinfo) GetMaxDamageOnce() uint32 {
	if x != nil && x.MaxDamageOnce != nil {
		return *x.MaxDamageOnce
	}
	return 0
}

type ShipExp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipId   *uint32 `protobuf:"varint,1,req,name=ship_id,json=shipId" json:"ship_id,omitempty"`
	Exp      *uint32 `protobuf:"varint,2,req,name=exp" json:"exp,omitempty"`
	Intimacy *uint32 `protobuf:"varint,3,req,name=intimacy" json:"intimacy,omitempty"`
}

func (x *ShipExp) Reset() {
	*x = ShipExp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p40_pb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipExp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipExp) ProtoMessage() {}

func (x *ShipExp) ProtoReflect() protoreflect.Message {
	mi := &file_p40_pb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipExp.ProtoReflect.Descriptor instead.
func (*ShipExp) Descriptor() ([]byte, []int) {
	return file_p40_pb_proto_rawDescGZIP(), []int{9}
}

func (x *ShipExp) GetShipId() uint32 {
	if x != nil && x.ShipId != nil {
		return *x.ShipId
	}
	return 0
}

func (x *ShipExp) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *ShipExp) GetIntimacy() uint32 {
	if x != nil && x.Intimacy != nil {
		return *x.Intimacy
	}
	return 0
}

var File_p40_pb_proto protoreflect.FileDescriptor

var file_p40_pb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x34, 0x30, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x34, 0x30, 0x5f, 0x70, 0x62, 0x22,
	0x58, 0x0a, 0x08, 0x63, 0x73, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x73, 0x63,
	0x5f, 0x34, 0x30, 0x30, 0x30, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x4b, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x7a, 0x75,
	0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x34, 0x30, 0x5f, 0x70, 0x62, 0x2e, 0x64, 0x72, 0x6f,
	0x70, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x64, 0x72,
	0x6f, 0x70, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xe0, 0x01,
	0x0a, 0x08, 0x63, 0x73, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x34,
	0x30, 0x5f, 0x70, 0x62, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x69,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x20, 0x0a, 0x0c, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xb8, 0x01, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x34, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x70, 0x34, 0x30, 0x5f, 0x70, 0x62, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x69,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x12, 0x3d, 0x0a, 0x0d,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70,
	0x34, 0x30, 0x5f, 0x70, 0x62, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x52, 0x0b,
	0x73, 0x68, 0x69, 0x70, 0x45, 0x78, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x63,
	0x73, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22,
	0x22, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x36, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x6c, 0x0a, 0x0f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x22, 0x46, 0x0a, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb2, 0x01, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x70, 0x52, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0d, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x63, 0x65, 0x22, 0x51,
	0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x69,
	0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x63,
	0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x66, 0x6c, 0x79, 0x71,
	0x69, 0x65, 0x2f, 0x62, 0x6c, 0x68, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x68, 0x78, 0x5f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x34, 0x30, 0x5f, 0x70, 0x62,
}

var (
	file_p40_pb_proto_rawDescOnce sync.Once
	file_p40_pb_proto_rawDescData = file_p40_pb_proto_rawDesc
)

func file_p40_pb_proto_rawDescGZIP() []byte {
	file_p40_pb_proto_rawDescOnce.Do(func() {
		file_p40_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_p40_pb_proto_rawDescData)
	})
	return file_p40_pb_proto_rawDescData
}

var file_p40_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_p40_pb_proto_goTypes = []interface{}{
	(*Cs_40001)(nil),        // 0: azurlane.p40_pb.cs_40001
	(*Sc_40002)(nil),        // 1: azurlane.p40_pb.sc_40002
	(*Cs_40003)(nil),        // 2: azurlane.p40_pb.cs_40003
	(*Sc_40004)(nil),        // 3: azurlane.p40_pb.sc_40004
	(*Cs_40005)(nil),        // 4: azurlane.p40_pb.cs_40005
	(*Sc_40006)(nil),        // 5: azurlane.p40_pb.sc_40006
	(*Dropperformance)(nil), // 6: azurlane.p40_pb.dropperformance
	(*Dropinfo)(nil),        // 7: azurlane.p40_pb.dropinfo
	(*Statisticsinfo)(nil),  // 8: azurlane.p40_pb.statisticsinfo
	(*ShipExp)(nil),         // 9: azurlane.p40_pb.ship_exp
}
var file_p40_pb_proto_depIdxs = []int32{
	6, // 0: azurlane.p40_pb.sc_40002.drop_performance:type_name -> azurlane.p40_pb.dropperformance
	8, // 1: azurlane.p40_pb.cs_40003.statistics:type_name -> azurlane.p40_pb.statisticsinfo
	7, // 2: azurlane.p40_pb.sc_40004.drop_info:type_name -> azurlane.p40_pb.dropinfo
	9, // 3: azurlane.p40_pb.sc_40004.ship_exp_list:type_name -> azurlane.p40_pb.ship_exp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_p40_pb_proto_init() }
func file_p40_pb_proto_init() {
	if File_p40_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p40_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_40001); i {
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
		file_p40_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_40002); i {
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
		file_p40_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_40003); i {
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
		file_p40_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_40004); i {
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
		file_p40_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_40005); i {
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
		file_p40_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_40006); i {
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
		file_p40_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dropperformance); i {
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
		file_p40_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dropinfo); i {
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
		file_p40_pb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statisticsinfo); i {
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
		file_p40_pb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipExp); i {
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
			RawDescriptor: file_p40_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p40_pb_proto_goTypes,
		DependencyIndexes: file_p40_pb_proto_depIdxs,
		MessageInfos:      file_p40_pb_proto_msgTypes,
	}.Build()
	File_p40_pb_proto = out.File
	file_p40_pb_proto_rawDesc = nil
	file_p40_pb_proto_goTypes = nil
	file_p40_pb_proto_depIdxs = nil
}
