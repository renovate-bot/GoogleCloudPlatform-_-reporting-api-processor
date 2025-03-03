// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: deprecation_report.proto

package securityreport

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

type DeprecationReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of API, e.g. websql
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// YYYY-MM-DD date format, e.g. "2020-01-01"
	AnticipatedRemoval string `protobuf:"bytes,2,opt,name=anticipated_removal,json=anticipatedRemoval,proto3" json:"anticipated_removal,omitempty"`
	// free form text, e.g. "WebSQL is deprecated and will be removed in Chrome 97 around January 2020"
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// where the call to the deprecated API happened, e.g. https://example.com/index.js
	SourceFile string `protobuf:"bytes,4,opt,name=source_file,json=sourceFile,proto3" json:"source_file,omitempty"`
	// 1-based
	LineNumber int32 `protobuf:"varint,5,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// 1-based
	ColumnNumber int32 `protobuf:"varint,6,opt,name=column_number,json=columnNumber,proto3" json:"column_number,omitempty"`
}

func (x *DeprecationReport) Reset() {
	*x = DeprecationReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deprecation_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeprecationReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecationReport) ProtoMessage() {}

func (x *DeprecationReport) ProtoReflect() protoreflect.Message {
	mi := &file_deprecation_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeprecationReport.ProtoReflect.Descriptor instead.
func (*DeprecationReport) Descriptor() ([]byte, []int) {
	return file_deprecation_report_proto_rawDescGZIP(), []int{0}
}

func (x *DeprecationReport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeprecationReport) GetAnticipatedRemoval() string {
	if x != nil {
		return x.AnticipatedRemoval
	}
	return ""
}

func (x *DeprecationReport) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeprecationReport) GetSourceFile() string {
	if x != nil {
		return x.SourceFile
	}
	return ""
}

func (x *DeprecationReport) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *DeprecationReport) GetColumnNumber() int32 {
	if x != nil {
		return x.ColumnNumber
	}
	return 0
}

var File_deprecation_report_proto protoreflect.FileDescriptor

var file_deprecation_report_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x11, 0x44,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x13, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deprecation_report_proto_rawDescOnce sync.Once
	file_deprecation_report_proto_rawDescData = file_deprecation_report_proto_rawDesc
)

func file_deprecation_report_proto_rawDescGZIP() []byte {
	file_deprecation_report_proto_rawDescOnce.Do(func() {
		file_deprecation_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_deprecation_report_proto_rawDescData)
	})
	return file_deprecation_report_proto_rawDescData
}

var file_deprecation_report_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_deprecation_report_proto_goTypes = []interface{}{
	(*DeprecationReport)(nil), // 0: securityreport.DeprecationReport
}
var file_deprecation_report_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deprecation_report_proto_init() }
func file_deprecation_report_proto_init() {
	if File_deprecation_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deprecation_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeprecationReport); i {
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
			RawDescriptor: file_deprecation_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deprecation_report_proto_goTypes,
		DependencyIndexes: file_deprecation_report_proto_depIdxs,
		MessageInfos:      file_deprecation_report_proto_msgTypes,
	}.Build()
	File_deprecation_report_proto = out.File
	file_deprecation_report_proto_rawDesc = nil
	file_deprecation_report_proto_goTypes = nil
	file_deprecation_report_proto_depIdxs = nil
}
