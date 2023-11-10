// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: tty28.proto

package pb

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

type GoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Authority       string `protobuf:"bytes,21,opt,name=authority,proto3" json:"authority,omitempty"`
	Origin          string `protobuf:"bytes,22,opt,name=origin,proto3" json:"origin,omitempty"`
	Referer         string `protobuf:"bytes,23,opt,name=referer,proto3" json:"referer,omitempty"`
	SecChUa         string `protobuf:"bytes,31,opt,name=secChUa,proto3" json:"secChUa,omitempty"`
	SecChUaPlatform string `protobuf:"bytes,32,opt,name=secChUaPlatform,proto3" json:"secChUaPlatform,omitempty"`
	UserAgent       string `protobuf:"bytes,33,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
}

func (x *GoldRequest) Reset() {
	*x = GoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty28_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoldRequest) ProtoMessage() {}

func (x *GoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tty28_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoldRequest.ProtoReflect.Descriptor instead.
func (*GoldRequest) Descriptor() ([]byte, []int) {
	return file_tty28_proto_rawDescGZIP(), []int{0}
}

func (x *GoldRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GoldRequest) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *GoldRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GoldRequest) GetReferer() string {
	if x != nil {
		return x.Referer
	}
	return ""
}

func (x *GoldRequest) GetSecChUa() string {
	if x != nil {
		return x.SecChUa
	}
	return ""
}

func (x *GoldRequest) GetSecChUaPlatform() string {
	if x != nil {
		return x.SecChUaPlatform
	}
	return ""
}

func (x *GoldRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type GoldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold int64 `protobuf:"varint,1,opt,name=gold,proto3" json:"gold,omitempty"`
}

func (x *GoldResponse) Reset() {
	*x = GoldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty28_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoldResponse) ProtoMessage() {}

func (x *GoldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tty28_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoldResponse.ProtoReflect.Descriptor instead.
func (*GoldResponse) Descriptor() ([]byte, []int) {
	return file_tty28_proto_rawDescGZIP(), []int{1}
}

func (x *GoldResponse) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

type BettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Authority       string `protobuf:"bytes,21,opt,name=authority,proto3" json:"authority,omitempty"`
	Origin          string `protobuf:"bytes,22,opt,name=origin,proto3" json:"origin,omitempty"`
	Referer         string `protobuf:"bytes,23,opt,name=referer,proto3" json:"referer,omitempty"`
	SecChUa         string `protobuf:"bytes,31,opt,name=secChUa,proto3" json:"secChUa,omitempty"`
	SecChUaPlatform string `protobuf:"bytes,32,opt,name=secChUaPlatform,proto3" json:"secChUaPlatform,omitempty"`
	UserAgent       string `protobuf:"bytes,33,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
}

func (x *BettingRequest) Reset() {
	*x = BettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty28_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BettingRequest) ProtoMessage() {}

func (x *BettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tty28_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BettingRequest.ProtoReflect.Descriptor instead.
func (*BettingRequest) Descriptor() ([]byte, []int) {
	return file_tty28_proto_rawDescGZIP(), []int{2}
}

func (x *BettingRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BettingRequest) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *BettingRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *BettingRequest) GetReferer() string {
	if x != nil {
		return x.Referer
	}
	return ""
}

func (x *BettingRequest) GetSecChUa() string {
	if x != nil {
		return x.SecChUa
	}
	return ""
}

func (x *BettingRequest) GetSecChUaPlatform() string {
	if x != nil {
		return x.SecChUaPlatform
	}
	return ""
}

func (x *BettingRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type BettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BettingResponse) Reset() {
	*x = BettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty28_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BettingResponse) ProtoMessage() {}

func (x *BettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tty28_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BettingResponse.ProtoReflect.Descriptor instead.
func (*BettingResponse) Descriptor() ([]byte, []int) {
	return file_tty28_proto_rawDescGZIP(), []int{3}
}

var File_tty28_proto protoreflect.FileDescriptor

var file_tty28_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x74, 0x79, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0b, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55,
	0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55, 0x61,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55, 0x61, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x43, 0x68,
	0x55, 0x61, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x6f, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x22, 0xd4, 0x01, 0x0a,
	0x0e, 0x42, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55, 0x61, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x73,
	0x65, 0x63, 0x43, 0x68, 0x55, 0x61, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x43, 0x68, 0x55, 0x61, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x40, 0x0a, 0x0b, 0x47, 0x6f, 0x6c, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4c, 0x0a, 0x0e, 0x42, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x42, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tty28_proto_rawDescOnce sync.Once
	file_tty28_proto_rawDescData = file_tty28_proto_rawDesc
)

func file_tty28_proto_rawDescGZIP() []byte {
	file_tty28_proto_rawDescOnce.Do(func() {
		file_tty28_proto_rawDescData = protoimpl.X.CompressGZIP(file_tty28_proto_rawDescData)
	})
	return file_tty28_proto_rawDescData
}

var file_tty28_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tty28_proto_goTypes = []interface{}{
	(*GoldRequest)(nil),     // 0: proto.GoldRequest
	(*GoldResponse)(nil),    // 1: proto.GoldResponse
	(*BettingRequest)(nil),  // 2: proto.BettingRequest
	(*BettingResponse)(nil), // 3: proto.BettingResponse
}
var file_tty28_proto_depIdxs = []int32{
	0, // 0: proto.GoldService.Gold:input_type -> proto.GoldRequest
	2, // 1: proto.BettingService.Betting:input_type -> proto.BettingRequest
	1, // 2: proto.GoldService.Gold:output_type -> proto.GoldResponse
	3, // 3: proto.BettingService.Betting:output_type -> proto.BettingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tty28_proto_init() }
func file_tty28_proto_init() {
	if File_tty28_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tty28_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoldRequest); i {
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
		file_tty28_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoldResponse); i {
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
		file_tty28_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BettingRequest); i {
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
		file_tty28_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BettingResponse); i {
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
			RawDescriptor: file_tty28_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_tty28_proto_goTypes,
		DependencyIndexes: file_tty28_proto_depIdxs,
		MessageInfos:      file_tty28_proto_msgTypes,
	}.Build()
	File_tty28_proto = out.File
	file_tty28_proto_rawDesc = nil
	file_tty28_proto_goTypes = nil
	file_tty28_proto_depIdxs = nil
}