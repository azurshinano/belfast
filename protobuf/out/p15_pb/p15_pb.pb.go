// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: p15_pb.proto

package p15_pb

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

type Sc_15001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*Iteminfo `protobuf:"bytes,1,rep,name=item_list,json=itemList" json:"item_list,omitempty"`
}

func (x *Sc_15001) Reset() {
	*x = Sc_15001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_15001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_15001) ProtoMessage() {}

func (x *Sc_15001) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_15001.ProtoReflect.Descriptor instead.
func (*Sc_15001) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Sc_15001) GetItemList() []*Iteminfo {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type Cs_15002 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Count *uint32 `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
}

func (x *Cs_15002) Reset() {
	*x = Cs_15002{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_15002) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_15002) ProtoMessage() {}

func (x *Cs_15002) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_15002.ProtoReflect.Descriptor instead.
func (*Cs_15002) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{1}
}

func (x *Cs_15002) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Cs_15002) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type Sc_15003 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *uint32     `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	DropList []*Dropinfo `protobuf:"bytes,2,rep,name=drop_list,json=dropList" json:"drop_list,omitempty"`
}

func (x *Sc_15003) Reset() {
	*x = Sc_15003{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_15003) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_15003) ProtoMessage() {}

func (x *Sc_15003) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_15003.ProtoReflect.Descriptor instead.
func (*Sc_15003) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Sc_15003) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *Sc_15003) GetDropList() []*Dropinfo {
	if x != nil {
		return x.DropList
	}
	return nil
}

type Cs_15004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Count *uint32 `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
}

func (x *Cs_15004) Reset() {
	*x = Cs_15004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_15004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_15004) ProtoMessage() {}

func (x *Cs_15004) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_15004.ProtoReflect.Descriptor instead.
func (*Cs_15004) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Cs_15004) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Cs_15004) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type Sc_15005 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
}

func (x *Sc_15005) Reset() {
	*x = Sc_15005{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_15005) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_15005) ProtoMessage() {}

func (x *Sc_15005) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_15005.ProtoReflect.Descriptor instead.
func (*Sc_15005) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Sc_15005) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

type Cs_15006 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Num *uint32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
}

func (x *Cs_15006) Reset() {
	*x = Cs_15006{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_15006) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_15006) ProtoMessage() {}

func (x *Cs_15006) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_15006.ProtoReflect.Descriptor instead.
func (*Cs_15006) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{5}
}

func (x *Cs_15006) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Cs_15006) GetNum() uint32 {
	if x != nil && x.Num != nil {
		return *x.Num
	}
	return 0
}

type Sc_15007 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
}

func (x *Sc_15007) Reset() {
	*x = Sc_15007{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_15007) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_15007) ProtoMessage() {}

func (x *Sc_15007) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_15007.ProtoReflect.Descriptor instead.
func (*Sc_15007) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{6}
}

func (x *Sc_15007) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

type Iteminfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Count *uint32 `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
}

func (x *Iteminfo) Reset() {
	*x = Iteminfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p15_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iteminfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iteminfo) ProtoMessage() {}

func (x *Iteminfo) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iteminfo.ProtoReflect.Descriptor instead.
func (*Iteminfo) Descriptor() ([]byte, []int) {
	return file_p15_pb_proto_rawDescGZIP(), []int{7}
}

func (x *Iteminfo) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Iteminfo) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
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
		mi := &file_p15_pb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dropinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dropinfo) ProtoMessage() {}

func (x *Dropinfo) ProtoReflect() protoreflect.Message {
	mi := &file_p15_pb_proto_msgTypes[8]
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
	return file_p15_pb_proto_rawDescGZIP(), []int{8}
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

var File_p15_pb_proto protoreflect.FileDescriptor

var file_p15_pb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x31, 0x35, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x31, 0x35, 0x5f, 0x70, 0x62, 0x22,
	0x42, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x31, 0x35, 0x30, 0x30, 0x31, 0x12, 0x36, 0x0a, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x31, 0x35, 0x5f, 0x70, 0x62,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x63, 0x73, 0x5f, 0x31, 0x35, 0x30, 0x30, 0x32, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x31, 0x35, 0x30, 0x30,
	0x33, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x72, 0x6f,
	0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61,
	0x7a, 0x75, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x31, 0x35, 0x5f, 0x70, 0x62, 0x2e, 0x64,
	0x72, 0x6f, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x30, 0x0a, 0x08, 0x63, 0x73, 0x5f, 0x31, 0x35, 0x30, 0x30, 0x34, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x31, 0x35, 0x30, 0x30, 0x35, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x08, 0x63, 0x73, 0x5f, 0x31, 0x35,
	0x30, 0x30, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x22, 0x0a, 0x08, 0x73, 0x63, 0x5f, 0x31, 0x35, 0x30, 0x30,
	0x37, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x08, 0x64,
	0x72, 0x6f, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x66, 0x6c,
	0x79, 0x71, 0x69, 0x65, 0x2f, 0x62, 0x6c, 0x68, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x68, 0x78, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x31, 0x35, 0x5f, 0x70,
	0x62,
}

var (
	file_p15_pb_proto_rawDescOnce sync.Once
	file_p15_pb_proto_rawDescData = file_p15_pb_proto_rawDesc
)

func file_p15_pb_proto_rawDescGZIP() []byte {
	file_p15_pb_proto_rawDescOnce.Do(func() {
		file_p15_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_p15_pb_proto_rawDescData)
	})
	return file_p15_pb_proto_rawDescData
}

var file_p15_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_p15_pb_proto_goTypes = []interface{}{
	(*Sc_15001)(nil), // 0: azurlane.p15_pb.sc_15001
	(*Cs_15002)(nil), // 1: azurlane.p15_pb.cs_15002
	(*Sc_15003)(nil), // 2: azurlane.p15_pb.sc_15003
	(*Cs_15004)(nil), // 3: azurlane.p15_pb.cs_15004
	(*Sc_15005)(nil), // 4: azurlane.p15_pb.sc_15005
	(*Cs_15006)(nil), // 5: azurlane.p15_pb.cs_15006
	(*Sc_15007)(nil), // 6: azurlane.p15_pb.sc_15007
	(*Iteminfo)(nil), // 7: azurlane.p15_pb.iteminfo
	(*Dropinfo)(nil), // 8: azurlane.p15_pb.dropinfo
}
var file_p15_pb_proto_depIdxs = []int32{
	7, // 0: azurlane.p15_pb.sc_15001.item_list:type_name -> azurlane.p15_pb.iteminfo
	8, // 1: azurlane.p15_pb.sc_15003.drop_list:type_name -> azurlane.p15_pb.dropinfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_p15_pb_proto_init() }
func file_p15_pb_proto_init() {
	if File_p15_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p15_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_15001); i {
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
		file_p15_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_15002); i {
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
		file_p15_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_15003); i {
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
		file_p15_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_15004); i {
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
		file_p15_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_15005); i {
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
		file_p15_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_15006); i {
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
		file_p15_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_15007); i {
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
		file_p15_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iteminfo); i {
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
		file_p15_pb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p15_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p15_pb_proto_goTypes,
		DependencyIndexes: file_p15_pb_proto_depIdxs,
		MessageInfos:      file_p15_pb_proto_msgTypes,
	}.Build()
	File_p15_pb_proto = out.File
	file_p15_pb_proto_rawDesc = nil
	file_p15_pb_proto_goTypes = nil
	file_p15_pb_proto_depIdxs = nil
}
