// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: rpc_cent.proto

package proto_files

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_rpc_cent_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Config) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Value    []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_rpc_cent_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Log) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_rpc_cent_proto protoreflect.FileDescriptor

var file_rpc_cent_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x11, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x37, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xf8, 0x02, 0x0a, 0x07,
	0x43, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x12, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x11,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x75, 0x6c,
	0x6c, 0x1a, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x75, 0x6c, 0x6c,
	0x1a, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x52, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x35, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x75, 0x6c, 0x6c, 0x1a, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x0d,
	0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6c, 0x6f, 0x67, 0x1a,
	0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x75,
	0x6c, 0x6c, 0x22, 0x00, 0x28, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x6f,
	0x75, 0x79, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_cent_proto_rawDescOnce sync.Once
	file_rpc_cent_proto_rawDescData = file_rpc_cent_proto_rawDesc
)

func file_rpc_cent_proto_rawDescGZIP() []byte {
	file_rpc_cent_proto_rawDescOnce.Do(func() {
		file_rpc_cent_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_cent_proto_rawDescData)
	})
	return file_rpc_cent_proto_rawDescData
}

var file_rpc_cent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_cent_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: base.common.config
	(*Log)(nil),    // 1: base.common.log
	(*Null)(nil),   // 2: base.common.null
}
var file_rpc_cent_proto_depIdxs = []int32{
	2, // 0: base.common.CentApi.GetBusinessConfFile:input_type -> base.common.null
	0, // 1: base.common.CentApi.SetBusinessConfFile:input_type -> base.common.config
	2, // 2: base.common.CentApi.GetRpcServerConfFile:input_type -> base.common.null
	0, // 3: base.common.CentApi.SetRpcServerConfFile:input_type -> base.common.config
	2, // 4: base.common.CentApi.FlushConfFile:input_type -> base.common.null
	1, // 5: base.common.CentApi.PushLogStream:input_type -> base.common.log
	0, // 6: base.common.CentApi.GetBusinessConfFile:output_type -> base.common.config
	2, // 7: base.common.CentApi.SetBusinessConfFile:output_type -> base.common.null
	0, // 8: base.common.CentApi.GetRpcServerConfFile:output_type -> base.common.config
	2, // 9: base.common.CentApi.SetRpcServerConfFile:output_type -> base.common.null
	2, // 10: base.common.CentApi.FlushConfFile:output_type -> base.common.null
	2, // 11: base.common.CentApi.PushLogStream:output_type -> base.common.null
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_cent_proto_init() }
func file_rpc_cent_proto_init() {
	if File_rpc_cent_proto != nil {
		return
	}
	file_rpc_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_cent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_rpc_cent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_rpc_cent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_cent_proto_goTypes,
		DependencyIndexes: file_rpc_cent_proto_depIdxs,
		MessageInfos:      file_rpc_cent_proto_msgTypes,
	}.Build()
	File_rpc_cent_proto = out.File
	file_rpc_cent_proto_rawDesc = nil
	file_rpc_cent_proto_goTypes = nil
	file_rpc_cent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CentApiClient is the client API for CentApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CentApiClient interface {
	// 配置中心接口
	GetBusinessConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Config, error)
	SetBusinessConfFile(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Null, error)
	GetRpcServerConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Config, error)
	SetRpcServerConfFile(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Null, error)
	// 刷新配置文件
	FlushConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Null, error)
	// 日志接口
	PushLogStream(ctx context.Context, opts ...grpc.CallOption) (CentApi_PushLogStreamClient, error)
}

type centApiClient struct {
	cc grpc.ClientConnInterface
}

func NewCentApiClient(cc grpc.ClientConnInterface) CentApiClient {
	return &centApiClient{cc}
}

func (c *centApiClient) GetBusinessConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/base.common.CentApi/GetBusinessConfFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centApiClient) SetBusinessConfFile(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/base.common.CentApi/SetBusinessConfFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centApiClient) GetRpcServerConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/base.common.CentApi/GetRpcServerConfFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centApiClient) SetRpcServerConfFile(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/base.common.CentApi/SetRpcServerConfFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centApiClient) FlushConfFile(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/base.common.CentApi/FlushConfFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centApiClient) PushLogStream(ctx context.Context, opts ...grpc.CallOption) (CentApi_PushLogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CentApi_serviceDesc.Streams[0], "/base.common.CentApi/PushLogStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &centApiPushLogStreamClient{stream}
	return x, nil
}

type CentApi_PushLogStreamClient interface {
	Send(*Log) error
	CloseAndRecv() (*Null, error)
	grpc.ClientStream
}

type centApiPushLogStreamClient struct {
	grpc.ClientStream
}

func (x *centApiPushLogStreamClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *centApiPushLogStreamClient) CloseAndRecv() (*Null, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Null)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CentApiServer is the server API for CentApi service.
type CentApiServer interface {
	// 配置中心接口
	GetBusinessConfFile(context.Context, *Null) (*Config, error)
	SetBusinessConfFile(context.Context, *Config) (*Null, error)
	GetRpcServerConfFile(context.Context, *Null) (*Config, error)
	SetRpcServerConfFile(context.Context, *Config) (*Null, error)
	// 刷新配置文件
	FlushConfFile(context.Context, *Null) (*Null, error)
	// 日志接口
	PushLogStream(CentApi_PushLogStreamServer) error
}

// UnimplementedCentApiServer can be embedded to have forward compatible implementations.
type UnimplementedCentApiServer struct {
}

func (*UnimplementedCentApiServer) GetBusinessConfFile(context.Context, *Null) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessConfFile not implemented")
}
func (*UnimplementedCentApiServer) SetBusinessConfFile(context.Context, *Config) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBusinessConfFile not implemented")
}
func (*UnimplementedCentApiServer) GetRpcServerConfFile(context.Context, *Null) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRpcServerConfFile not implemented")
}
func (*UnimplementedCentApiServer) SetRpcServerConfFile(context.Context, *Config) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRpcServerConfFile not implemented")
}
func (*UnimplementedCentApiServer) FlushConfFile(context.Context, *Null) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushConfFile not implemented")
}
func (*UnimplementedCentApiServer) PushLogStream(CentApi_PushLogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PushLogStream not implemented")
}

func RegisterCentApiServer(s *grpc.Server, srv CentApiServer) {
	s.RegisterService(&_CentApi_serviceDesc, srv)
}

func _CentApi_GetBusinessConfFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentApiServer).GetBusinessConfFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.common.CentApi/GetBusinessConfFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentApiServer).GetBusinessConfFile(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentApi_SetBusinessConfFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentApiServer).SetBusinessConfFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.common.CentApi/SetBusinessConfFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentApiServer).SetBusinessConfFile(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentApi_GetRpcServerConfFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentApiServer).GetRpcServerConfFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.common.CentApi/GetRpcServerConfFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentApiServer).GetRpcServerConfFile(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentApi_SetRpcServerConfFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentApiServer).SetRpcServerConfFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.common.CentApi/SetRpcServerConfFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentApiServer).SetRpcServerConfFile(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentApi_FlushConfFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentApiServer).FlushConfFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.common.CentApi/FlushConfFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentApiServer).FlushConfFile(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentApi_PushLogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CentApiServer).PushLogStream(&centApiPushLogStreamServer{stream})
}

type CentApi_PushLogStreamServer interface {
	SendAndClose(*Null) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type centApiPushLogStreamServer struct {
	grpc.ServerStream
}

func (x *centApiPushLogStreamServer) SendAndClose(m *Null) error {
	return x.ServerStream.SendMsg(m)
}

func (x *centApiPushLogStreamServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CentApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.common.CentApi",
	HandlerType: (*CentApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusinessConfFile",
			Handler:    _CentApi_GetBusinessConfFile_Handler,
		},
		{
			MethodName: "SetBusinessConfFile",
			Handler:    _CentApi_SetBusinessConfFile_Handler,
		},
		{
			MethodName: "GetRpcServerConfFile",
			Handler:    _CentApi_GetRpcServerConfFile_Handler,
		},
		{
			MethodName: "SetRpcServerConfFile",
			Handler:    _CentApi_SetRpcServerConfFile_Handler,
		},
		{
			MethodName: "FlushConfFile",
			Handler:    _CentApi_FlushConfFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushLogStream",
			Handler:       _CentApi_PushLogStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "rpc_cent.proto",
}