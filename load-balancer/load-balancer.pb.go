// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: load-balancer.proto

package load_balancer

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/kulycloud/protocol/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_load_balancer_proto protoreflect.FileDescriptor

var file_load_balancer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x65, 0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x27, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x0d, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6c, 0x79, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x61, 0x64,
	0x2d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x3b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_load_balancer_proto_goTypes = []interface{}{
	(*common.EndpointList)(nil), // 0: EndpointList
	(*common.Empty)(nil),        // 1: Empty
}
var file_load_balancer_proto_depIdxs = []int32{
	0, // 0: LoadBalancer.SetStorageEndpoints:input_type -> EndpointList
	0, // 1: LoadBalancer.SetEndpoints:input_type -> EndpointList
	1, // 2: LoadBalancer.SetStorageEndpoints:output_type -> Empty
	1, // 3: LoadBalancer.SetEndpoints:output_type -> Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_load_balancer_proto_init() }
func file_load_balancer_proto_init() {
	if File_load_balancer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_load_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_load_balancer_proto_goTypes,
		DependencyIndexes: file_load_balancer_proto_depIdxs,
	}.Build()
	File_load_balancer_proto = out.File
	file_load_balancer_proto_rawDesc = nil
	file_load_balancer_proto_goTypes = nil
	file_load_balancer_proto_depIdxs = nil
}
