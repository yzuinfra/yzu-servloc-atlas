// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: protos/birpc.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ip            string                 `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Hostname      string                 `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Service       string                 `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Version       string                 `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_protos_birpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_birpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protos_birpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *RegisterRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RegisterRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Registered    bool                   `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_protos_birpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_birpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_protos_birpc_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

var File_protos_birpc_proto protoreflect.FileDescriptor

const file_protos_birpc_proto_rawDesc = "" +
	"\n" +
	"\x12protos/birpc.proto\"q\n" +
	"\x0fRegisterRequest\x12\x0e\n" +
	"\x02ip\x18\x01 \x01(\tR\x02ip\x12\x1a\n" +
	"\bhostname\x18\x02 \x01(\tR\bhostname\x12\x18\n" +
	"\aservice\x18\x03 \x01(\tR\aservice\x12\x18\n" +
	"\aversion\x18\x04 \x01(\tR\aversion\"2\n" +
	"\x10RegisterResponse\x12\x1e\n" +
	"\n" +
	"registered\x18\x01 \x01(\bR\n" +
	"registered2A\n" +
	"\x05BiRPC\x128\n" +
	"\rRegisterAgent\x12\x10.RegisterRequest\x1a\x11.RegisterResponse(\x010\x01B\x17Z\x15yzuinfra/atlas/protosb\x06proto3"

var (
	file_protos_birpc_proto_rawDescOnce sync.Once
	file_protos_birpc_proto_rawDescData []byte
)

func file_protos_birpc_proto_rawDescGZIP() []byte {
	file_protos_birpc_proto_rawDescOnce.Do(func() {
		file_protos_birpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_birpc_proto_rawDesc), len(file_protos_birpc_proto_rawDesc)))
	})
	return file_protos_birpc_proto_rawDescData
}

var file_protos_birpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_birpc_proto_goTypes = []any{
	(*RegisterRequest)(nil),  // 0: RegisterRequest
	(*RegisterResponse)(nil), // 1: RegisterResponse
}
var file_protos_birpc_proto_depIdxs = []int32{
	0, // 0: BiRPC.RegisterAgent:input_type -> RegisterRequest
	1, // 1: BiRPC.RegisterAgent:output_type -> RegisterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_birpc_proto_init() }
func file_protos_birpc_proto_init() {
	if File_protos_birpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_birpc_proto_rawDesc), len(file_protos_birpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_birpc_proto_goTypes,
		DependencyIndexes: file_protos_birpc_proto_depIdxs,
		MessageInfos:      file_protos_birpc_proto_msgTypes,
	}.Build()
	File_protos_birpc_proto = out.File
	file_protos_birpc_proto_goTypes = nil
	file_protos_birpc_proto_depIdxs = nil
}
