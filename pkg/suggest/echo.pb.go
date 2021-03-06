// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: echo.proto

package suggest

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SuggestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *SuggestRequest) Reset() {
	*x = SuggestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestRequest) ProtoMessage() {}

func (x *SuggestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestRequest.ProtoReflect.Descriptor instead.
func (*SuggestRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Position uint32 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Suggestion) Reset() {
	*x = Suggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestion) ProtoMessage() {}

func (x *Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestion.ProtoReflect.Descriptor instead.
func (*Suggestion) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Suggestion) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Suggestion) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type SuggestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggestions []*Suggestion `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
}

func (x *SuggestResponse) Reset() {
	*x = SuggestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestResponse) ProtoMessage() {}

func (x *SuggestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestResponse.ProtoReflect.Descriptor instead.
func (*SuggestResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{2}
}

func (x *SuggestResponse) GetSuggestions() []*Suggestion {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

var File_echo_proto protoreflect.FileDescriptor

var file_echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x3c, 0x0a, 0x0a, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0f, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x61, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x56,
	0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_proto_rawDescOnce sync.Once
	file_echo_proto_rawDescData = file_echo_proto_rawDesc
)

func file_echo_proto_rawDescGZIP() []byte {
	file_echo_proto_rawDescOnce.Do(func() {
		file_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_proto_rawDescData)
	})
	return file_echo_proto_rawDescData
}

var file_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_echo_proto_goTypes = []interface{}{
	(*SuggestRequest)(nil),  // 0: suggest.SuggestRequest
	(*Suggestion)(nil),      // 1: suggest.Suggestion
	(*SuggestResponse)(nil), // 2: suggest.SuggestResponse
}
var file_echo_proto_depIdxs = []int32{
	1, // 0: suggest.SuggestResponse.suggestions:type_name -> suggest.Suggestion
	0, // 1: suggest.Suggest.Input:input_type -> suggest.SuggestRequest
	2, // 2: suggest.Suggest.Input:output_type -> suggest.SuggestResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_echo_proto_init() }
func file_echo_proto_init() {
	if File_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestRequest); i {
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
		file_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggestion); i {
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
		file_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestResponse); i {
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
			RawDescriptor: file_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_proto_goTypes,
		DependencyIndexes: file_echo_proto_depIdxs,
		MessageInfos:      file_echo_proto_msgTypes,
	}.Build()
	File_echo_proto = out.File
	file_echo_proto_rawDesc = nil
	file_echo_proto_goTypes = nil
	file_echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SuggestClient is the client API for Suggest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuggestClient interface {
	Input(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error)
}

type suggestClient struct {
	cc grpc.ClientConnInterface
}

func NewSuggestClient(cc grpc.ClientConnInterface) SuggestClient {
	return &suggestClient{cc}
}

func (c *suggestClient) Input(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error) {
	out := new(SuggestResponse)
	err := c.cc.Invoke(ctx, "/suggest.Suggest/Input", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuggestServer is the server API for Suggest service.
type SuggestServer interface {
	Input(context.Context, *SuggestRequest) (*SuggestResponse, error)
}

// UnimplementedSuggestServer can be embedded to have forward compatible implementations.
type UnimplementedSuggestServer struct {
}

func (*UnimplementedSuggestServer) Input(context.Context, *SuggestRequest) (*SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Input not implemented")
}

func RegisterSuggestServer(s *grpc.Server, srv SuggestServer) {
	s.RegisterService(&_Suggest_serviceDesc, srv)
}

func _Suggest_Input_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestServer).Input(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suggest.Suggest/Input",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestServer).Input(ctx, req.(*SuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Suggest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "suggest.Suggest",
	HandlerType: (*SuggestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Input",
			Handler:    _Suggest_Input_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
