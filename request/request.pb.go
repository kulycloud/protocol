// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: request.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	storage "github.com/kulycloud/protocol/storage"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpData    *HTTPRequest_HTTPData `protobuf:"bytes,1,opt,name=http_data,json=httpData,proto3" json:"http_data,omitempty"`
	KulyData    *HTTPRequest_KulyData `protobuf:"bytes,2,opt,name=kuly_data,json=kulyData,proto3" json:"kuly_data,omitempty"`
	ServiceData string                `protobuf:"bytes,3,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPRequest) GetHttpData() *HTTPRequest_HTTPData {
	if x != nil {
		return x.HttpData
	}
	return nil
}

func (x *HTTPRequest) GetKulyData() *HTTPRequest_KulyData {
	if x != nil {
		return x.KulyData
	}
	return nil
}

func (x *HTTPRequest) GetServiceData() string {
	if x != nil {
		return x.ServiceData
	}
	return ""
}

type HTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     int32             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Header     map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       string            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	RequestUid string            `protobuf:"bytes,4,opt,name=request_uid,json=requestUid,proto3" json:"request_uid,omitempty"`
}

func (x *HTTPResponse) Reset() {
	*x = HTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPResponse) ProtoMessage() {}

func (x *HTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPResponse.ProtoReflect.Descriptor instead.
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HTTPResponse) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HTTPResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *HTTPResponse) GetRequestUid() string {
	if x != nil {
		return x.RequestUid
	}
	return ""
}

type HTTPRequest_HTTPData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source  string            `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Body    string            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HTTPRequest_HTTPData) Reset() {
	*x = HTTPRequest_HTTPData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest_HTTPData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest_HTTPData) ProtoMessage() {}

func (x *HTTPRequest_HTTPData) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest_HTTPData.ProtoReflect.Descriptor instead.
func (*HTTPRequest_HTTPData) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HTTPRequest_HTTPData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HTTPRequest_HTTPData) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HTTPRequest_HTTPData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *HTTPRequest_HTTPData) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type HTTPRequest_KulyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteUid   string             `protobuf:"bytes,1,opt,name=route_uid,json=routeUid,proto3" json:"route_uid,omitempty"`
	Step       *storage.RouteStep `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	RequestUid string             `protobuf:"bytes,3,opt,name=request_uid,json=requestUid,proto3" json:"request_uid,omitempty"`
}

func (x *HTTPRequest_KulyData) Reset() {
	*x = HTTPRequest_KulyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest_KulyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest_KulyData) ProtoMessage() {}

func (x *HTTPRequest_KulyData) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest_KulyData.ProtoReflect.Descriptor instead.
func (*HTTPRequest_KulyData) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HTTPRequest_KulyData) GetRouteUid() string {
	if x != nil {
		return x.RouteUid
	}
	return ""
}

func (x *HTTPRequest_KulyData) GetStep() *storage.RouteStep {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *HTTPRequest_KulyData) GetRequestUid() string {
	if x != nil {
		return x.RequestUid
	}
	return ""
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x03, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x32, 0x0a, 0x09, 0x6b, 0x75, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4b, 0x75, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6b, 0x75,
	0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0xc4, 0x01, 0x0a, 0x08, 0x48, 0x54,
	0x54, 0x50, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x68, 0x0a, 0x08, 0x4b, 0x75, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x55, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x69, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0c, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x69, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_request_proto_goTypes = []interface{}{
	(*HTTPRequest)(nil),          // 0: HTTPRequest
	(*HTTPResponse)(nil),         // 1: HTTPResponse
	(*HTTPRequest_HTTPData)(nil), // 2: HTTPRequest.HTTPData
	(*HTTPRequest_KulyData)(nil), // 3: HTTPRequest.KulyData
	nil,                          // 4: HTTPRequest.HTTPData.HeadersEntry
	nil,                          // 5: HTTPResponse.HeaderEntry
	(*storage.RouteStep)(nil),    // 6: RouteStep
}
var file_request_proto_depIdxs = []int32{
	2, // 0: HTTPRequest.http_data:type_name -> HTTPRequest.HTTPData
	3, // 1: HTTPRequest.kuly_data:type_name -> HTTPRequest.KulyData
	5, // 2: HTTPResponse.header:type_name -> HTTPResponse.HeaderEntry
	4, // 3: HTTPRequest.HTTPData.headers:type_name -> HTTPRequest.HTTPData.HeadersEntry
	6, // 4: HTTPRequest.KulyData.step:type_name -> RouteStep
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPResponse); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest_HTTPData); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest_KulyData); i {
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
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
