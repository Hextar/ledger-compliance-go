// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: lc.proto

package schema

import (
	schema "github.com/codenotary/immudb/pkg/api/schema"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

var File_lc_proto protoreflect.FileDescriptor

var file_lc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x63, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xfd, 0x02, 0x0a, 0x09, 0x4c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x14,
	0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e,
	0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4b, 0x65,
	0x79, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x53, 0x61, 0x66, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x14, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x53, 0x61,
	0x66, 0x65, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1d, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x75, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2d,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_lc_proto_goTypes = []interface{}{
	(*schema.KeyValue)(nil),       // 0: immudb.schema.KeyValue
	(*schema.Key)(nil),            // 1: immudb.schema.Key
	(*schema.SafeSetOptions)(nil), // 2: immudb.schema.SafeSetOptions
	(*schema.SafeGetOptions)(nil), // 3: immudb.schema.SafeGetOptions
	(*empty.Empty)(nil),           // 4: google.protobuf.Empty
	(*schema.Index)(nil),          // 5: immudb.schema.Index
	(*schema.Item)(nil),           // 6: immudb.schema.Item
	(*schema.Proof)(nil),          // 7: immudb.schema.Proof
	(*schema.SafeItem)(nil),       // 8: immudb.schema.SafeItem
	(*schema.Root)(nil),           // 9: immudb.schema.Root
	(*schema.HealthResponse)(nil), // 10: immudb.schema.HealthResponse
}
var file_lc_proto_depIdxs = []int32{
	0,  // 0: lc.schema.LcService.Set:input_type -> immudb.schema.KeyValue
	1,  // 1: lc.schema.LcService.Get:input_type -> immudb.schema.Key
	2,  // 2: lc.schema.LcService.SafeSet:input_type -> immudb.schema.SafeSetOptions
	3,  // 3: lc.schema.LcService.SafeGet:input_type -> immudb.schema.SafeGetOptions
	4,  // 4: lc.schema.LcService.CurrentRoot:input_type -> google.protobuf.Empty
	4,  // 5: lc.schema.LcService.Health:input_type -> google.protobuf.Empty
	5,  // 6: lc.schema.LcService.Set:output_type -> immudb.schema.Index
	6,  // 7: lc.schema.LcService.Get:output_type -> immudb.schema.Item
	7,  // 8: lc.schema.LcService.SafeSet:output_type -> immudb.schema.Proof
	8,  // 9: lc.schema.LcService.SafeGet:output_type -> immudb.schema.SafeItem
	9,  // 10: lc.schema.LcService.CurrentRoot:output_type -> immudb.schema.Root
	10, // 11: lc.schema.LcService.Health:output_type -> immudb.schema.HealthResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_lc_proto_init() }
func file_lc_proto_init() {
	if File_lc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lc_proto_goTypes,
		DependencyIndexes: file_lc_proto_depIdxs,
	}.Build()
	File_lc_proto = out.File
	file_lc_proto_rawDesc = nil
	file_lc_proto_goTypes = nil
	file_lc_proto_depIdxs = nil
}