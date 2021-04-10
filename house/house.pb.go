// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: house/house.proto

package house

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backdrops []*Image `protobuf:"bytes,1,rep,name=backdrops,proto3" json:"backdrops,omitempty"`
	Posters   []*Image `protobuf:"bytes,2,rep,name=posters,proto3" json:"posters,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_house_house_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_house_house_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_house_house_proto_rawDescGZIP(), []int{0}
}

func (x *Images) GetBackdrops() []*Image {
	if x != nil {
		return x.Backdrops
	}
	return nil
}

func (x *Images) GetPosters() []*Image {
	if x != nil {
		return x.Posters
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_house_house_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_house_house_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_house_house_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id",validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
	// @inject_tag: json:"full"
	Full bool `protobuf:"varint,2,opt,name=full,proto3" json:"full"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_house_house_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_house_house_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDetails.ProtoReflect.Descriptor instead.
func (*RequestDetails) Descriptor() ([]byte, []int) {
	return file_house_house_proto_rawDescGZIP(), []int{2}
}

func (x *RequestDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RequestDetails) GetFull() bool {
	if x != nil {
		return x.Full
	}
	return false
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"_id"
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"_id"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	BackdropPath string `protobuf:"bytes,3,opt,name=backdrop_path,json=backdropPath,proto3" json:"backdrop_path,omitempty"`
	PosterPath   string `protobuf:"bytes,4,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_house_house_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_house_house_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDetails.ProtoReflect.Descriptor instead.
func (*ResponseDetails) Descriptor() ([]byte, []int) {
	return file_house_house_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseDetails) GetBackdropPath() string {
	if x != nil {
		return x.BackdropPath
	}
	return ""
}

func (x *ResponseDetails) GetPosterPath() string {
	if x != nil {
		return x.PosterPath
	}
	return ""
}

type RequestCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title",validate:"required"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title" validate:"required"`
}

func (x *RequestCreate) Reset() {
	*x = RequestCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_house_house_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate) ProtoMessage() {}

func (x *RequestCreate) ProtoReflect() protoreflect.Message {
	mi := &file_house_house_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate.ProtoReflect.Descriptor instead.
func (*RequestCreate) Descriptor() ([]byte, []int) {
	return file_house_house_proto_rawDescGZIP(), []int{4}
}

func (x *RequestCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_house_house_proto protoreflect.FileDescriptor

var file_house_house_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x06, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x73,
	0x12, 0x26, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x34,
	0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x66, 0x75, 0x6c, 0x6c, 0x22, 0x7d, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x48, 0x0a, 0x0c, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_house_house_proto_rawDescOnce sync.Once
	file_house_house_proto_rawDescData = file_house_house_proto_rawDesc
)

func file_house_house_proto_rawDescGZIP() []byte {
	file_house_house_proto_rawDescOnce.Do(func() {
		file_house_house_proto_rawDescData = protoimpl.X.CompressGZIP(file_house_house_proto_rawDescData)
	})
	return file_house_house_proto_rawDescData
}

var file_house_house_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_house_house_proto_goTypes = []interface{}{
	(*Images)(nil),          // 0: house.Images
	(*Image)(nil),           // 1: house.Image
	(*RequestDetails)(nil),  // 2: house.RequestDetails
	(*ResponseDetails)(nil), // 3: house.ResponseDetails
	(*RequestCreate)(nil),   // 4: house.RequestCreate
}
var file_house_house_proto_depIdxs = []int32{
	1, // 0: house.Images.backdrops:type_name -> house.Image
	1, // 1: house.Images.posters:type_name -> house.Image
	4, // 2: house.houseService.Create:input_type -> house.RequestCreate
	3, // 3: house.houseService.Create:output_type -> house.ResponseDetails
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_house_house_proto_init() }
func file_house_house_proto_init() {
	if File_house_house_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_house_house_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_house_house_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_house_house_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDetails); i {
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
		file_house_house_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDetails); i {
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
		file_house_house_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate); i {
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
			RawDescriptor: file_house_house_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_house_house_proto_goTypes,
		DependencyIndexes: file_house_house_proto_depIdxs,
		MessageInfos:      file_house_house_proto_msgTypes,
	}.Build()
	File_house_house_proto = out.File
	file_house_house_proto_rawDesc = nil
	file_house_house_proto_goTypes = nil
	file_house_house_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HouseServiceClient is the client API for HouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HouseServiceClient interface {
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error)
}

type houseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHouseServiceClient(cc grpc.ClientConnInterface) HouseServiceClient {
	return &houseServiceClient{cc}
}

func (c *houseServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/house.houseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HouseServiceServer is the server API for HouseService service.
type HouseServiceServer interface {
	Create(context.Context, *RequestCreate) (*ResponseDetails, error)
}

// UnimplementedHouseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHouseServiceServer struct {
}

func (*UnimplementedHouseServiceServer) Create(context.Context, *RequestCreate) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterHouseServiceServer(s *grpc.Server, srv HouseServiceServer) {
	s.RegisterService(&_HouseService_serviceDesc, srv)
}

func _HouseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/house.houseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

var _HouseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "house.houseService",
	HandlerType: (*HouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _HouseService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "house/house.proto",
}
