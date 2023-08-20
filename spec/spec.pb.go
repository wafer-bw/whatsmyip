// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: spec/spec.proto

package spec

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

type IPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IPRequest) Reset() {
	*x = IPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPRequest) ProtoMessage() {}

func (x *IPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPRequest.ProtoReflect.Descriptor instead.
func (*IPRequest) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{0}
}

type IPReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IPReply) Reset() {
	*x = IPReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPReply) ProtoMessage() {}

func (x *IPReply) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPReply.ProtoReflect.Descriptor instead.
func (*IPReply) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{1}
}

func (x *IPReply) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_spec_spec_proto protoreflect.FileDescriptor

var file_spec_spec_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x0b, 0x0a, 0x09, 0x49, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x07, 0x49, 0x50, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x42,
	0x10, 0x5a, 0x0e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x6d, 0x79, 0x69, 0x70, 0x2f, 0x73, 0x70, 0x65,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_spec_proto_rawDescOnce sync.Once
	file_spec_spec_proto_rawDescData = file_spec_spec_proto_rawDesc
)

func file_spec_spec_proto_rawDescGZIP() []byte {
	file_spec_spec_proto_rawDescOnce.Do(func() {
		file_spec_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_spec_proto_rawDescData)
	})
	return file_spec_spec_proto_rawDescData
}

var file_spec_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spec_spec_proto_goTypes = []interface{}{
	(*IPRequest)(nil), // 0: spec.IPRequest
	(*IPReply)(nil),   // 1: spec.IPReply
}
var file_spec_spec_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spec_spec_proto_init() }
func file_spec_spec_proto_init() {
	if File_spec_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPRequest); i {
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
		file_spec_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPReply); i {
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
			RawDescriptor: file_spec_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spec_spec_proto_goTypes,
		DependencyIndexes: file_spec_spec_proto_depIdxs,
		MessageInfos:      file_spec_spec_proto_msgTypes,
	}.Build()
	File_spec_spec_proto = out.File
	file_spec_spec_proto_rawDesc = nil
	file_spec_spec_proto_goTypes = nil
	file_spec_spec_proto_depIdxs = nil
}
