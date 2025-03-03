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
// source: csp_violation.proto

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

type CspReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URI of the document in which the violation occurred.
	DocumentUri string `protobuf:"bytes,4,opt,name=document_uri,json=documentUri,proto3" json:"document_uri,omitempty"`
	// The referrer of the document in which the violation occurred.
	Referrer string `protobuf:"bytes,5,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// The URI of the resource that was blocked from loading by the
	// Content Security Policy. If the blocked URI is from a different
	// origin than the document-uri, then the blocked URI is truncated
	// to contain just the scheme, host, and port.
	BlockedUri string `protobuf:"bytes,6,opt,name=blocked_uri,json=blockedUri,proto3" json:"blocked_uri,omitempty"`
	// The name of the policy section that was violated.
	ViolatedDirective string `protobuf:"bytes,7,opt,name=violated_directive,json=violatedDirective,proto3" json:"violated_directive,omitempty"`
	// The name of the policy directive that was violated.
	EffectiveDirective string `protobuf:"bytes,15,opt,name=effective_directive,json=effectiveDirective,proto3" json:"effective_directive,omitempty"`
	// The original policy as specified by the CSP HTTP header:
	// Content-Security-Policy, X-Content-Security-Policy (IE),
	// X-Webkit-CSP (old Safari, old Chrome).
	OriginalPolicy string `protobuf:"bytes,8,opt,name=original_policy,json=originalPolicy,proto3" json:"original_policy,omitempty"`
	// The URL of the resource where the violation occurred.
	SourceFile string `protobuf:"bytes,9,opt,name=source_file,json=sourceFile,proto3" json:"source_file,omitempty"`
	// The line number in source-file on which the violation occurred, 1-based.
	LineNumber int32 `protobuf:"varint,31,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// The column number in source-file on which the violation occurred, 1-based.
	ColumnNumber int32 `protobuf:"varint,32,opt,name=column_number,json=columnNumber,proto3" json:"column_number,omitempty"`
	// A snippet of the rejected script (first 40 bytes).
	ScriptSample string `protobuf:"bytes,21,opt,name=script_sample,json=scriptSample,proto3" json:"script_sample,omitempty"`
}

func (x *CspReport) Reset() {
	*x = CspReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csp_violation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CspReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CspReport) ProtoMessage() {}

func (x *CspReport) ProtoReflect() protoreflect.Message {
	mi := &file_csp_violation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CspReport.ProtoReflect.Descriptor instead.
func (*CspReport) Descriptor() ([]byte, []int) {
	return file_csp_violation_proto_rawDescGZIP(), []int{0}
}

func (x *CspReport) GetDocumentUri() string {
	if x != nil {
		return x.DocumentUri
	}
	return ""
}

func (x *CspReport) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *CspReport) GetBlockedUri() string {
	if x != nil {
		return x.BlockedUri
	}
	return ""
}

func (x *CspReport) GetViolatedDirective() string {
	if x != nil {
		return x.ViolatedDirective
	}
	return ""
}

func (x *CspReport) GetEffectiveDirective() string {
	if x != nil {
		return x.EffectiveDirective
	}
	return ""
}

func (x *CspReport) GetOriginalPolicy() string {
	if x != nil {
		return x.OriginalPolicy
	}
	return ""
}

func (x *CspReport) GetSourceFile() string {
	if x != nil {
		return x.SourceFile
	}
	return ""
}

func (x *CspReport) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *CspReport) GetColumnNumber() int32 {
	if x != nil {
		return x.ColumnNumber
	}
	return 0
}

func (x *CspReport) GetScriptSample() string {
	if x != nil {
		return x.ScriptSample
	}
	return ""
}

var File_csp_violation_proto protoreflect.FileDescriptor

var file_csp_violation_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x73, 0x70, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x80, 0x03, 0x0a, 0x09, 0x43, 0x73, 0x70, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x55, 0x72, 0x69, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_csp_violation_proto_rawDescOnce sync.Once
	file_csp_violation_proto_rawDescData = file_csp_violation_proto_rawDesc
)

func file_csp_violation_proto_rawDescGZIP() []byte {
	file_csp_violation_proto_rawDescOnce.Do(func() {
		file_csp_violation_proto_rawDescData = protoimpl.X.CompressGZIP(file_csp_violation_proto_rawDescData)
	})
	return file_csp_violation_proto_rawDescData
}

var file_csp_violation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_csp_violation_proto_goTypes = []interface{}{
	(*CspReport)(nil), // 0: securityreport.CspReport
}
var file_csp_violation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_csp_violation_proto_init() }
func file_csp_violation_proto_init() {
	if File_csp_violation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csp_violation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CspReport); i {
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
			RawDescriptor: file_csp_violation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_csp_violation_proto_goTypes,
		DependencyIndexes: file_csp_violation_proto_depIdxs,
		MessageInfos:      file_csp_violation_proto_msgTypes,
	}.Build()
	File_csp_violation_proto = out.File
	file_csp_violation_proto_rawDesc = nil
	file_csp_violation_proto_goTypes = nil
	file_csp_violation_proto_depIdxs = nil
}
