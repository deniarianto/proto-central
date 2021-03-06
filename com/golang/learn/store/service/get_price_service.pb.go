// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: get_price_service.proto

package service

import (
	context "context"
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

type CalculationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *CalculationRequest) Reset() {
	*x = CalculationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_price_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationRequest) ProtoMessage() {}

func (x *CalculationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_price_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationRequest.ProtoReflect.Descriptor instead.
func (*CalculationRequest) Descriptor() ([]byte, []int) {
	return file_get_price_service_proto_rawDescGZIP(), []int{0}
}

func (x *CalculationRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CalculationRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type CalculationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculationResponse) Reset() {
	*x = CalculationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_price_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationResponse) ProtoMessage() {}

func (x *CalculationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_price_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationResponse.ProtoReflect.Descriptor instead.
func (*CalculationResponse) Descriptor() ([]byte, []int) {
	return file_get_price_service_proto_rawDescGZIP(), []int{1}
}

func (x *CalculationResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_get_price_service_proto protoreflect.FileDescriptor

var file_get_price_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x22, 0x30, 0x0a, 0x12, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x62, 0x22, 0x2d, 0x0a, 0x13, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x99, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12,
	0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_price_service_proto_rawDescOnce sync.Once
	file_get_price_service_proto_rawDescData = file_get_price_service_proto_rawDesc
)

func file_get_price_service_proto_rawDescGZIP() []byte {
	file_get_price_service_proto_rawDescOnce.Do(func() {
		file_get_price_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_price_service_proto_rawDescData)
	})
	return file_get_price_service_proto_rawDescData
}

var file_get_price_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_get_price_service_proto_goTypes = []interface{}{
	(*CalculationRequest)(nil),  // 0: store.calculationRequest
	(*CalculationResponse)(nil), // 1: store.calculationResponse
}
var file_get_price_service_proto_depIdxs = []int32{
	0, // 0: store.CalculationService.add:input_type -> store.calculationRequest
	0, // 1: store.CalculationService.multiply:input_type -> store.calculationRequest
	1, // 2: store.CalculationService.add:output_type -> store.calculationResponse
	1, // 3: store.CalculationService.multiply:output_type -> store.calculationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_get_price_service_proto_init() }
func file_get_price_service_proto_init() {
	if File_get_price_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_get_price_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationRequest); i {
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
		file_get_price_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationResponse); i {
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
			RawDescriptor: file_get_price_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_get_price_service_proto_goTypes,
		DependencyIndexes: file_get_price_service_proto_depIdxs,
		MessageInfos:      file_get_price_service_proto_msgTypes,
	}.Build()
	File_get_price_service_proto = out.File
	file_get_price_service_proto_rawDesc = nil
	file_get_price_service_proto_goTypes = nil
	file_get_price_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculationServiceClient is the client API for CalculationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculationServiceClient interface {
	Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Multiply(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
}

type calculationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculationServiceClient(cc grpc.ClientConnInterface) CalculationServiceClient {
	return &calculationServiceClient{cc}
}

func (c *calculationServiceClient) Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/store.CalculationService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculationServiceClient) Multiply(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/store.CalculationService/multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculationServiceServer is the server API for CalculationService service.
type CalculationServiceServer interface {
	Add(context.Context, *CalculationRequest) (*CalculationResponse, error)
	Multiply(context.Context, *CalculationRequest) (*CalculationResponse, error)
}

// UnimplementedCalculationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculationServiceServer struct {
}

func (*UnimplementedCalculationServiceServer) Add(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculationServiceServer) Multiply(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterCalculationServiceServer(s *grpc.Server, srv CalculationServiceServer) {
	s.RegisterService(&_CalculationService_serviceDesc, srv)
}

func _CalculationService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculationServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.CalculationService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculationServiceServer).Add(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculationService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculationServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.CalculationService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculationServiceServer).Multiply(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.CalculationService",
	HandlerType: (*CalculationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _CalculationService_Add_Handler,
		},
		{
			MethodName: "multiply",
			Handler:    _CalculationService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_price_service.proto",
}
