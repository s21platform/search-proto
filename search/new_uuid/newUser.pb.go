// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: newUser.proto

package new_uuid

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

type UpdateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *UpdateUser) Reset() {
	*x = UpdateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUser) ProtoMessage() {}

func (x *UpdateUser) ProtoReflect() protoreflect.Message {
	mi := &file_newUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUser.ProtoReflect.Descriptor instead.
func (*UpdateUser) Descriptor() ([]byte, []int) {
	return file_newUser_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUser) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_newUser_proto protoreflect.FileDescriptor

var file_newUser_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x6e, 0x65,
	0x77, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newUser_proto_rawDescOnce sync.Once
	file_newUser_proto_rawDescData = file_newUser_proto_rawDesc
)

func file_newUser_proto_rawDescGZIP() []byte {
	file_newUser_proto_rawDescOnce.Do(func() {
		file_newUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_newUser_proto_rawDescData)
	})
	return file_newUser_proto_rawDescData
}

var file_newUser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_newUser_proto_goTypes = []interface{}{
	(*UpdateUser)(nil), // 0: UpdateUser
}
var file_newUser_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newUser_proto_init() }
func file_newUser_proto_init() {
	if File_newUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUser); i {
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
			RawDescriptor: file_newUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_newUser_proto_goTypes,
		DependencyIndexes: file_newUser_proto_depIdxs,
		MessageInfos:      file_newUser_proto_msgTypes,
	}.Build()
	File_newUser_proto = out.File
	file_newUser_proto_rawDesc = nil
	file_newUser_proto_goTypes = nil
	file_newUser_proto_depIdxs = nil
}
