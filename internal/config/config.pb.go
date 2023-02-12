// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/config/config.proto

package config

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

type RepositoryType int32

const (
	RepositoryType_UNSPECIFIED RepositoryType = 0
	RepositoryType_IN_MEMORY   RepositoryType = 1
	RepositoryType_SQL         RepositoryType = 2
)

// Enum value maps for RepositoryType.
var (
	RepositoryType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "IN_MEMORY",
		2: "SQL",
	}
	RepositoryType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"IN_MEMORY":   1,
		"SQL":         2,
	}
)

func (x RepositoryType) Enum() *RepositoryType {
	p := new(RepositoryType)
	*p = x
	return p
}

func (x RepositoryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RepositoryType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_config_proto_enumTypes[0].Descriptor()
}

func (RepositoryType) Type() protoreflect.EnumType {
	return &file_proto_config_config_proto_enumTypes[0]
}

func (x RepositoryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RepositoryType.Descriptor instead.
func (RepositoryType) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_config_proto_rawDescGZIP(), []int{0}
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                  string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	RepositoryType           RepositoryType `protobuf:"varint,11,opt,name=repository_type,json=repositoryType,proto3,enum=config.RepositoryType" json:"repository_type,omitempty"`
	DatabaseDsn              string         `protobuf:"bytes,12,opt,name=database_dsn,json=databaseDsn,proto3" json:"database_dsn,omitempty"`
	FileStorage              string         `protobuf:"bytes,13,opt,name=file_storage,json=fileStorage,proto3" json:"file_storage,omitempty"`
	TokenRenewalTimeMinutes  int64          `protobuf:"varint,21,opt,name=token_renewal_time_minutes,json=tokenRenewalTimeMinutes,proto3" json:"token_renewal_time_minutes,omitempty"`
	TokenValidityTimeMinutes int64          `protobuf:"varint,22,opt,name=token_validity_time_minutes,json=tokenValidityTimeMinutes,proto3" json:"token_validity_time_minutes,omitempty"`
	FileStaleTimeMinutes     int64          `protobuf:"varint,23,opt,name=file_stale_time_minutes,json=fileStaleTimeMinutes,proto3" json:"file_stale_time_minutes,omitempty"`
	SecretEncryptionEnabled  bool           `protobuf:"varint,31,opt,name=secret_encryption_enabled,json=secretEncryptionEnabled,proto3" json:"secret_encryption_enabled,omitempty"`
	ConfigFile               string         `protobuf:"bytes,99,opt,name=config_file,json=configFile,proto3" json:"config_file,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ServerConfig) GetRepositoryType() RepositoryType {
	if x != nil {
		return x.RepositoryType
	}
	return RepositoryType_UNSPECIFIED
}

func (x *ServerConfig) GetDatabaseDsn() string {
	if x != nil {
		return x.DatabaseDsn
	}
	return ""
}

func (x *ServerConfig) GetFileStorage() string {
	if x != nil {
		return x.FileStorage
	}
	return ""
}

func (x *ServerConfig) GetTokenRenewalTimeMinutes() int64 {
	if x != nil {
		return x.TokenRenewalTimeMinutes
	}
	return 0
}

func (x *ServerConfig) GetTokenValidityTimeMinutes() int64 {
	if x != nil {
		return x.TokenValidityTimeMinutes
	}
	return 0
}

func (x *ServerConfig) GetFileStaleTimeMinutes() int64 {
	if x != nil {
		return x.FileStaleTimeMinutes
	}
	return 0
}

func (x *ServerConfig) GetSecretEncryptionEnabled() bool {
	if x != nil {
		return x.SecretEncryptionEnabled
	}
	return false
}

func (x *ServerConfig) GetConfigFile() string {
	if x != nil {
		return x.ConfigFile
	}
	return ""
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	Login         string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	StaticToken   string `protobuf:"bytes,4,opt,name=static_token,json=staticToken,proto3" json:"static_token,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *ClientConfig) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ClientConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ClientConfig) GetStaticToken() string {
	if x != nil {
		return x.StaticToken
	}
	return ""
}

var File_proto_config_config_proto protoreflect.FileDescriptor

var file_proto_config_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0xbf, 0x03, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3f,
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x73, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x44,
	0x73, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x14, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x2a, 0x39, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x4f,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69, 0x67,
	0x61, 0x6e, 0x67, 0x2f, 0x47, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_config_proto_rawDescOnce sync.Once
	file_proto_config_config_proto_rawDescData = file_proto_config_config_proto_rawDesc
)

func file_proto_config_config_proto_rawDescGZIP() []byte {
	file_proto_config_config_proto_rawDescOnce.Do(func() {
		file_proto_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_config_proto_rawDescData)
	})
	return file_proto_config_config_proto_rawDescData
}

var file_proto_config_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_config_config_proto_goTypes = []interface{}{
	(RepositoryType)(0),  // 0: config.RepositoryType
	(*ServerConfig)(nil), // 1: config.ServerConfig
	(*ClientConfig)(nil), // 2: config.ClientConfig
}
var file_proto_config_config_proto_depIdxs = []int32{
	0, // 0: config.ServerConfig.repository_type:type_name -> config.RepositoryType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_config_config_proto_init() }
func file_proto_config_config_proto_init() {
	if File_proto_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
		file_proto_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
			RawDescriptor: file_proto_config_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_config_proto_goTypes,
		DependencyIndexes: file_proto_config_config_proto_depIdxs,
		EnumInfos:         file_proto_config_config_proto_enumTypes,
		MessageInfos:      file_proto_config_config_proto_msgTypes,
	}.Build()
	File_proto_config_config_proto = out.File
	file_proto_config_config_proto_rawDesc = nil
	file_proto_config_config_proto_goTypes = nil
	file_proto_config_config_proto_depIdxs = nil
}
