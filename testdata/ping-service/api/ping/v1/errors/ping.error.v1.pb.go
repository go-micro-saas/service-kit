// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: testdata/ping-service/api/ping/v1/errors/ping.error.v1.proto

package errorv1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

// ERROR .
type ERROR int32

const (
	ERROR_UNKNOWN         ERROR = 0
	ERROR_CONTENT_MISSING ERROR = 1
	ERROR_CONTENT_ERROR   ERROR = 2
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0: "UNKNOWN",
		1: "CONTENT_MISSING",
		2: "CONTENT_ERROR",
	}
	ERROR_value = map[string]int32{
		"UNKNOWN":         0,
		"CONTENT_MISSING": 1,
		"CONTENT_ERROR":   2,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto protoreflect.FileDescriptor

var file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x54, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x12, 0x11, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x17, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03,
	0x42, 0x84, 0x01, 0x0a, 0x15, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x42, 0x12, 0x53, 0x61, 0x61, 0x73,
	0x41, 0x70, 0x69, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x31, 0x50, 0x01,
	0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescOnce sync.Once
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescData = file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDesc
)

func file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescGZIP() []byte {
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescOnce.Do(func() {
		file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescData)
	})
	return file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDescData
}

var file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_goTypes = []any{
	(ERROR)(0), // 0: saas.api.ping.errorv1.ERROR
}
var file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_init() }
func file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_init() {
	if File_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_goTypes,
		DependencyIndexes: file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_depIdxs,
		EnumInfos:         file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_enumTypes,
	}.Build()
	File_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto = out.File
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_rawDesc = nil
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_goTypes = nil
	file_testdata_ping_service_api_ping_v1_errors_ping_error_v1_proto_depIdxs = nil
}