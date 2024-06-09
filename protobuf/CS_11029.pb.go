// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: CS_11029.proto

package protobuf

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

type CS_11029 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackTyp *uint32 `protobuf:"varint,1,req,name=track_typ,json=trackTyp" json:"track_typ,omitempty"`
	IntArg1  *uint32 `protobuf:"varint,2,req,name=int_arg1,json=intArg1" json:"int_arg1,omitempty"`
	IntArg2  *uint32 `protobuf:"varint,3,req,name=int_arg2,json=intArg2" json:"int_arg2,omitempty"`
	IntArg3  *uint32 `protobuf:"varint,4,req,name=int_arg3,json=intArg3" json:"int_arg3,omitempty"`
	StrArg1  *string `protobuf:"bytes,5,req,name=str_arg1,json=strArg1" json:"str_arg1,omitempty"`
}

func (x *CS_11029) Reset() {
	*x = CS_11029{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_11029_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_11029) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_11029) ProtoMessage() {}

func (x *CS_11029) ProtoReflect() protoreflect.Message {
	mi := &file_CS_11029_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_11029.ProtoReflect.Descriptor instead.
func (*CS_11029) Descriptor() ([]byte, []int) {
	return file_CS_11029_proto_rawDescGZIP(), []int{0}
}

func (x *CS_11029) GetTrackTyp() uint32 {
	if x != nil && x.TrackTyp != nil {
		return *x.TrackTyp
	}
	return 0
}

func (x *CS_11029) GetIntArg1() uint32 {
	if x != nil && x.IntArg1 != nil {
		return *x.IntArg1
	}
	return 0
}

func (x *CS_11029) GetIntArg2() uint32 {
	if x != nil && x.IntArg2 != nil {
		return *x.IntArg2
	}
	return 0
}

func (x *CS_11029) GetIntArg3() uint32 {
	if x != nil && x.IntArg3 != nil {
		return *x.IntArg3
	}
	return 0
}

func (x *CS_11029) GetStrArg1() string {
	if x != nil && x.StrArg1 != nil {
		return *x.StrArg1
	}
	return ""
}

var File_CS_11029_proto protoreflect.FileDescriptor

var file_CS_11029_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x31, 0x31, 0x30, 0x32, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x43,
	0x53, 0x5f, 0x31, 0x31, 0x30, 0x32, 0x39, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x31,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x31, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x32, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x5f, 0x61, 0x72, 0x67, 0x33, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e,
	0x74, 0x41, 0x72, 0x67, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f, 0x61, 0x72, 0x67,
	0x31, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x41, 0x72, 0x67, 0x31,
	0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
}

var (
	file_CS_11029_proto_rawDescOnce sync.Once
	file_CS_11029_proto_rawDescData = file_CS_11029_proto_rawDesc
)

func file_CS_11029_proto_rawDescGZIP() []byte {
	file_CS_11029_proto_rawDescOnce.Do(func() {
		file_CS_11029_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_11029_proto_rawDescData)
	})
	return file_CS_11029_proto_rawDescData
}

var file_CS_11029_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_11029_proto_goTypes = []interface{}{
	(*CS_11029)(nil), // 0: protobuf.CS_11029
}
var file_CS_11029_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CS_11029_proto_init() }
func file_CS_11029_proto_init() {
	if File_CS_11029_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CS_11029_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_11029); i {
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
			RawDescriptor: file_CS_11029_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_11029_proto_goTypes,
		DependencyIndexes: file_CS_11029_proto_depIdxs,
		MessageInfos:      file_CS_11029_proto_msgTypes,
	}.Build()
	File_CS_11029_proto = out.File
	file_CS_11029_proto_rawDesc = nil
	file_CS_11029_proto_goTypes = nil
	file_CS_11029_proto_depIdxs = nil
}