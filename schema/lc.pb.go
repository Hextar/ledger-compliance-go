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

type TamperReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Key   []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Root  []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *TamperReport) Reset() {
	*x = TamperReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TamperReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TamperReport) ProtoMessage() {}

func (x *TamperReport) ProtoReflect() protoreflect.Message {
	mi := &file_lc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TamperReport.ProtoReflect.Descriptor instead.
func (*TamperReport) Descriptor() ([]byte, []int) {
	return file_lc_proto_rawDescGZIP(), []int{0}
}

func (x *TamperReport) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *TamperReport) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TamperReport) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

type ReportOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *TamperReport     `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature *schema.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ReportOptions) Reset() {
	*x = ReportOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportOptions) ProtoMessage() {}

func (x *ReportOptions) ProtoReflect() protoreflect.Message {
	mi := &file_lc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportOptions.ProtoReflect.Descriptor instead.
func (*ReportOptions) Descriptor() ([]byte, []int) {
	return file_lc_proto_rawDescGZIP(), []int{1}
}

func (x *ReportOptions) GetPayload() *TamperReport {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ReportOptions) GetSignature() *schema.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_lc_proto protoreflect.FileDescriptor

var file_lc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x63, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x0c, 0x54, 0x61, 0x6d, 0x70, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x22, 0x7a, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x61, 0x6d, 0x70, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x36, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xfb, 0x03, 0x0a, 0x09, 0x4c, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e,
	0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x14, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x6d, 0x75,
	0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x07, 0x53, 0x61, 0x66, 0x65, 0x53, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x69, 0x6d,
	0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61, 0x66, 0x65,
	0x53, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x14, 0x2e, 0x69, 0x6d, 0x6d,
	0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x53, 0x61, 0x66, 0x65, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e,
	0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61,
	0x66, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x69,
	0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x61, 0x66,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52,
	0x6f, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x64, 0x62,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x6d, 0x70,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x6c, 0x63, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x75, 0x73, 0x2f, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lc_proto_rawDescOnce sync.Once
	file_lc_proto_rawDescData = file_lc_proto_rawDesc
)

func file_lc_proto_rawDescGZIP() []byte {
	file_lc_proto_rawDescOnce.Do(func() {
		file_lc_proto_rawDescData = protoimpl.X.CompressGZIP(file_lc_proto_rawDescData)
	})
	return file_lc_proto_rawDescData
}

var file_lc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lc_proto_goTypes = []interface{}{
	(*TamperReport)(nil),          // 0: lc.schema.TamperReport
	(*ReportOptions)(nil),         // 1: lc.schema.ReportOptions
	(*schema.Signature)(nil),      // 2: immudb.schema.Signature
	(*schema.KeyValue)(nil),       // 3: immudb.schema.KeyValue
	(*schema.Key)(nil),            // 4: immudb.schema.Key
	(*schema.SafeSetOptions)(nil), // 5: immudb.schema.SafeSetOptions
	(*schema.SafeGetOptions)(nil), // 6: immudb.schema.SafeGetOptions
	(*empty.Empty)(nil),           // 7: google.protobuf.Empty
	(*schema.Index)(nil),          // 8: immudb.schema.Index
	(*schema.Item)(nil),           // 9: immudb.schema.Item
	(*schema.Proof)(nil),          // 10: immudb.schema.Proof
	(*schema.SafeItem)(nil),       // 11: immudb.schema.SafeItem
	(*schema.Root)(nil),           // 12: immudb.schema.Root
	(*schema.HealthResponse)(nil), // 13: immudb.schema.HealthResponse
	(*schema.ItemList)(nil),       // 14: immudb.schema.ItemList
}
var file_lc_proto_depIdxs = []int32{
	0,  // 0: lc.schema.ReportOptions.payload:type_name -> lc.schema.TamperReport
	2,  // 1: lc.schema.ReportOptions.signature:type_name -> immudb.schema.Signature
	3,  // 2: lc.schema.LcService.Set:input_type -> immudb.schema.KeyValue
	4,  // 3: lc.schema.LcService.Get:input_type -> immudb.schema.Key
	5,  // 4: lc.schema.LcService.SafeSet:input_type -> immudb.schema.SafeSetOptions
	6,  // 5: lc.schema.LcService.SafeGet:input_type -> immudb.schema.SafeGetOptions
	7,  // 6: lc.schema.LcService.CurrentRoot:input_type -> google.protobuf.Empty
	7,  // 7: lc.schema.LcService.Health:input_type -> google.protobuf.Empty
	4,  // 8: lc.schema.LcService.History:input_type -> immudb.schema.Key
	1,  // 9: lc.schema.LcService.ReportTamper:input_type -> lc.schema.ReportOptions
	8,  // 10: lc.schema.LcService.Set:output_type -> immudb.schema.Index
	9,  // 11: lc.schema.LcService.Get:output_type -> immudb.schema.Item
	10, // 12: lc.schema.LcService.SafeSet:output_type -> immudb.schema.Proof
	11, // 13: lc.schema.LcService.SafeGet:output_type -> immudb.schema.SafeItem
	12, // 14: lc.schema.LcService.CurrentRoot:output_type -> immudb.schema.Root
	13, // 15: lc.schema.LcService.Health:output_type -> immudb.schema.HealthResponse
	14, // 16: lc.schema.LcService.History:output_type -> immudb.schema.ItemList
	7,  // 17: lc.schema.LcService.ReportTamper:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_lc_proto_init() }
func file_lc_proto_init() {
	if File_lc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TamperReport); i {
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
		file_lc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportOptions); i {
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
			RawDescriptor: file_lc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lc_proto_goTypes,
		DependencyIndexes: file_lc_proto_depIdxs,
		MessageInfos:      file_lc_proto_msgTypes,
	}.Build()
	File_lc_proto = out.File
	file_lc_proto_rawDesc = nil
	file_lc_proto_goTypes = nil
	file_lc_proto_depIdxs = nil
}
