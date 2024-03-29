// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: rpc_secretKey.proto

package proto_files

import (
	_ "github.com/mwitkow/go-proto-validators"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 回传的提示码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 回传使用的提示信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 用户id
	Uid int32 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// 过期时间
	ExpTime int64 `protobuf:"varint,4,opt,name=exp_time,json=expTime,proto3" json:"exp_time,omitempty"`
	// 返回绑定的token
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_secretKey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_secretKey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rpc_secretKey_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *User) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *User) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *User) GetExpTime() int64 {
	if x != nil {
		return x.ExpTime
	}
	return 0
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RsaKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 公钥
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// 私钥
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// 客户端id,与私钥绑定，获取私钥使用
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *RsaKey) Reset() {
	*x = RsaKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_secretKey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RsaKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RsaKey) ProtoMessage() {}

func (x *RsaKey) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_secretKey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RsaKey.ProtoReflect.Descriptor instead.
func (*RsaKey) Descriptor() ([]byte, []int) {
	return file_rpc_secretKey_proto_rawDescGZIP(), []int{1}
}

func (x *RsaKey) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *RsaKey) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *RsaKey) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type WechatTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid     string `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	SessionKey string `protobuf:"bytes,2,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
	Unionid    string `protobuf:"bytes,3,opt,name=unionid,proto3" json:"unionid,omitempty"`
	Errcode    int32  `protobuf:"varint,4,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Errmsg     string `protobuf:"bytes,5,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *WechatTokenInfo) Reset() {
	*x = WechatTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_secretKey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatTokenInfo) ProtoMessage() {}

func (x *WechatTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_secretKey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatTokenInfo.ProtoReflect.Descriptor instead.
func (*WechatTokenInfo) Descriptor() ([]byte, []int) {
	return file_rpc_secretKey_proto_rawDescGZIP(), []int{2}
}

func (x *WechatTokenInfo) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *WechatTokenInfo) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *WechatTokenInfo) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *WechatTokenInfo) GetErrcode() int32 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *WechatTokenInfo) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

type UserVcCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BindInfo string `protobuf:"bytes,3,opt,name=bind_info,json=bindInfo,proto3" json:"bind_info,omitempty"`
	VcCode   string `protobuf:"bytes,4,opt,name=vc_code,json=vcCode,proto3" json:"vc_code,omitempty"`
}

func (x *UserVcCode) Reset() {
	*x = UserVcCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_secretKey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVcCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVcCode) ProtoMessage() {}

func (x *UserVcCode) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_secretKey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVcCode.ProtoReflect.Descriptor instead.
func (*UserVcCode) Descriptor() ([]byte, []int) {
	return file_rpc_secretKey_proto_rawDescGZIP(), []int{3}
}

func (x *UserVcCode) GetBindInfo() string {
	if x != nil {
		return x.BindInfo
	}
	return ""
}

func (x *UserVcCode) GetVcCode() string {
	if x != nil {
		return x.VcCode
	}
	return ""
}

var File_rpc_secretKey_proto protoreflect.FileDescriptor

var file_rpc_secretKey_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x1a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x40, 0x76, 0x30, 0x2e, 0x33, 0x2e, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x07, 0x72, 0x73, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x11,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x55, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76,
	0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x07, 0x76,
	0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf,
	0x1f, 0x03, 0x80, 0x01, 0x06, 0x52, 0x06, 0x76, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x23, 0x5a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x6f, 0x75, 0x79, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_secretKey_proto_rawDescOnce sync.Once
	file_rpc_secretKey_proto_rawDescData = file_rpc_secretKey_proto_rawDesc
)

func file_rpc_secretKey_proto_rawDescGZIP() []byte {
	file_rpc_secretKey_proto_rawDescOnce.Do(func() {
		file_rpc_secretKey_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_secretKey_proto_rawDescData)
	})
	return file_rpc_secretKey_proto_rawDescData
}

var file_rpc_secretKey_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_secretKey_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: base.common.rpc_cent.user
	(*RsaKey)(nil),          // 1: base.common.rpc_cent.rsa_key
	(*WechatTokenInfo)(nil), // 2: base.common.rpc_cent.wechat_token_info
	(*UserVcCode)(nil),      // 3: base.common.rpc_cent.user_vc_code
}
var file_rpc_secretKey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_secretKey_proto_init() }
func file_rpc_secretKey_proto_init() {
	if File_rpc_secretKey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_secretKey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_rpc_secretKey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RsaKey); i {
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
		file_rpc_secretKey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatTokenInfo); i {
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
		file_rpc_secretKey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVcCode); i {
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
			RawDescriptor: file_rpc_secretKey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_secretKey_proto_goTypes,
		DependencyIndexes: file_rpc_secretKey_proto_depIdxs,
		MessageInfos:      file_rpc_secretKey_proto_msgTypes,
	}.Build()
	File_rpc_secretKey_proto = out.File
	file_rpc_secretKey_proto_rawDesc = nil
	file_rpc_secretKey_proto_goTypes = nil
	file_rpc_secretKey_proto_depIdxs = nil
}
