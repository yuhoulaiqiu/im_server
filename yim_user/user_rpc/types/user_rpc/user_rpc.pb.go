// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: user_rpc.proto

package user_rpc

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

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname       string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password       string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role           int32  `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
	Avatar         string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	OpenId         string `protobuf:"bytes,5,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	RegisterSource string `protobuf:"bytes,6,opt,name=register_source,json=registerSource,proto3" json:"register_source,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *UserCreateRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *UserCreateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserCreateRequest) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *UserCreateRequest) GetRegisterSource() string {
	if x != nil {
		return x.RegisterSource
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` //用户和用户配置的json
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfoResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_user_rpc_proto protoreflect.FileDescriptor

var file_user_rpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x26, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x93, 0x01, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_rpc_proto_rawDescOnce sync.Once
	file_user_rpc_proto_rawDescData = file_user_rpc_proto_rawDesc
)

func file_user_rpc_proto_rawDescGZIP() []byte {
	file_user_rpc_proto_rawDescOnce.Do(func() {
		file_user_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_rpc_proto_rawDescData)
	})
	return file_user_rpc_proto_rawDescData
}

var file_user_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_rpc_proto_goTypes = []interface{}{
	(*UserCreateRequest)(nil),  // 0: user_rpc.UserCreateRequest
	(*UserCreateResponse)(nil), // 1: user_rpc.UserCreateResponse
	(*UserInfoRequest)(nil),    // 2: user_rpc.UserInfoRequest
	(*UserInfoResponse)(nil),   // 3: user_rpc.UserInfoResponse
}
var file_user_rpc_proto_depIdxs = []int32{
	0, // 0: user_rpc.Users.UserCreate:input_type -> user_rpc.UserCreateRequest
	2, // 1: user_rpc.Users.UserInfo:input_type -> user_rpc.UserInfoRequest
	1, // 2: user_rpc.Users.UserCreate:output_type -> user_rpc.UserCreateResponse
	3, // 3: user_rpc.Users.UserInfo:output_type -> user_rpc.UserInfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_rpc_proto_init() }
func file_user_rpc_proto_init() {
	if File_user_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_user_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateResponse); i {
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
		file_user_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_user_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
			RawDescriptor: file_user_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_rpc_proto_goTypes,
		DependencyIndexes: file_user_rpc_proto_depIdxs,
		MessageInfos:      file_user_rpc_proto_msgTypes,
	}.Build()
	File_user_rpc_proto = out.File
	file_user_rpc_proto_rawDesc = nil
	file_user_rpc_proto_goTypes = nil
	file_user_rpc_proto_depIdxs = nil
}
