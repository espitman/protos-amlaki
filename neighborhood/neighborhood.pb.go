// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: neighborhood/neighborhood.proto

package neighborhood

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

type Neighborhood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title",validate:"required"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title" validate:"required"`
	// @inject_tag: json:"name",validate:"alpha,required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"alpha,required"`
	// @inject_tag: json:"province",validate:"alphanum,required"
	Province string `protobuf:"bytes,3,opt,name=province,proto3" json:"province" validate:"alphanum,required"`
	// @inject_tag: json:"city",validate:"alphanum,required"
	City string `protobuf:"bytes,4,opt,name=city,proto3" json:"city" validate:"alphanum,required"`
}

func (x *Neighborhood) Reset() {
	*x = Neighborhood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Neighborhood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Neighborhood) ProtoMessage() {}

func (x *Neighborhood) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Neighborhood.ProtoReflect.Descriptor instead.
func (*Neighborhood) Descriptor() ([]byte, []int) {
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{0}
}

func (x *Neighborhood) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Neighborhood) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Neighborhood) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Neighborhood) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id",validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[1]
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
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{1}
}

func (x *RequestDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Neighborhood *Neighborhood `protobuf:"bytes,2,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	Creator      string        `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[2]
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
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseDetails) GetNeighborhood() *Neighborhood {
	if x != nil {
		return x.Neighborhood
	}
	return nil
}

func (x *ResponseDetails) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type RequestCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"neighborhood",validate:"required"
	Neighborhood *Neighborhood `protobuf:"bytes,1,opt,name=neighborhood,proto3" json:"neighborhood" validate:"required"`
	// @inject_tag: json:"-"
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"-"`
}

func (x *RequestCreate) Reset() {
	*x = RequestCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate) ProtoMessage() {}

func (x *RequestCreate) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[3]
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
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{3}
}

func (x *RequestCreate) GetNeighborhood() *Neighborhood {
	if x != nil {
		return x.Neighborhood
	}
	return nil
}

func (x *RequestCreate) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type RequestCheckExistByProvinceIdAndCityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Neighborhood *Neighborhood `protobuf:"bytes,1,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
}

func (x *RequestCheckExistByProvinceIdAndCityId) Reset() {
	*x = RequestCheckExistByProvinceIdAndCityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCheckExistByProvinceIdAndCityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCheckExistByProvinceIdAndCityId) ProtoMessage() {}

func (x *RequestCheckExistByProvinceIdAndCityId) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCheckExistByProvinceIdAndCityId.ProtoReflect.Descriptor instead.
func (*RequestCheckExistByProvinceIdAndCityId) Descriptor() ([]byte, []int) {
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{4}
}

func (x *RequestCheckExistByProvinceIdAndCityId) GetNeighborhood() *Neighborhood {
	if x != nil {
		return x.Neighborhood
	}
	return nil
}

type ResponseCheckExistByProvinceIdAndCityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool `protobuf:"varint,1,opt,name=isExist,proto3" json:"isExist,omitempty"`
}

func (x *ResponseCheckExistByProvinceIdAndCityId) Reset() {
	*x = ResponseCheckExistByProvinceIdAndCityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCheckExistByProvinceIdAndCityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCheckExistByProvinceIdAndCityId) ProtoMessage() {}

func (x *ResponseCheckExistByProvinceIdAndCityId) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCheckExistByProvinceIdAndCityId.ProtoReflect.Descriptor instead.
func (*ResponseCheckExistByProvinceIdAndCityId) Descriptor() ([]byte, []int) {
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseCheckExistByProvinceIdAndCityId) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

type RequestCheckValidNeighborhood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Province string `protobuf:"bytes,2,opt,name=province,proto3" json:"province,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *RequestCheckValidNeighborhood) Reset() {
	*x = RequestCheckValidNeighborhood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCheckValidNeighborhood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCheckValidNeighborhood) ProtoMessage() {}

func (x *RequestCheckValidNeighborhood) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCheckValidNeighborhood.ProtoReflect.Descriptor instead.
func (*RequestCheckValidNeighborhood) Descriptor() ([]byte, []int) {
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{6}
}

func (x *RequestCheckValidNeighborhood) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestCheckValidNeighborhood) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *RequestCheckValidNeighborhood) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type ResponseCheckValidNeighborhood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *ResponseCheckValidNeighborhood) Reset() {
	*x = ResponseCheckValidNeighborhood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neighborhood_neighborhood_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCheckValidNeighborhood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCheckValidNeighborhood) ProtoMessage() {}

func (x *ResponseCheckValidNeighborhood) ProtoReflect() protoreflect.Message {
	mi := &file_neighborhood_neighborhood_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCheckValidNeighborhood.ProtoReflect.Descriptor instead.
func (*ResponseCheckValidNeighborhood) Descriptor() ([]byte, []int) {
	return file_neighborhood_neighborhood_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseCheckValidNeighborhood) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_neighborhood_neighborhood_proto protoreflect.FileDescriptor

var file_neighborhood_neighborhood_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2f, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x22,
	0x68, 0x0a, 0x0c, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e,
	0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68,
	0x6f, 0x6f, 0x64, 0x2e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64,
	0x52, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x69, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x4e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x52, 0x0c, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x22, 0x68, 0x0a, 0x26, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3e, 0x0a,
	0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f,
	0x6f, 0x64, 0x2e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x52,
	0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x22, 0x43, 0x0a,
	0x27, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x41,
	0x6e, 0x64, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x22, 0x63, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68,
	0x6f, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x32, 0xe7, 0x02, 0x0a, 0x13, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x68, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f,
	0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x1f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x41, 0x6e,
	0x64, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x34, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62,
	0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a, 0x35, 0x2e,
	0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64,
	0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x1a, 0x2c, 0x2e,
	0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neighborhood_neighborhood_proto_rawDescOnce sync.Once
	file_neighborhood_neighborhood_proto_rawDescData = file_neighborhood_neighborhood_proto_rawDesc
)

func file_neighborhood_neighborhood_proto_rawDescGZIP() []byte {
	file_neighborhood_neighborhood_proto_rawDescOnce.Do(func() {
		file_neighborhood_neighborhood_proto_rawDescData = protoimpl.X.CompressGZIP(file_neighborhood_neighborhood_proto_rawDescData)
	})
	return file_neighborhood_neighborhood_proto_rawDescData
}

var file_neighborhood_neighborhood_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_neighborhood_neighborhood_proto_goTypes = []interface{}{
	(*Neighborhood)(nil),                            // 0: neighborhood.Neighborhood
	(*RequestDetails)(nil),                          // 1: neighborhood.RequestDetails
	(*ResponseDetails)(nil),                         // 2: neighborhood.ResponseDetails
	(*RequestCreate)(nil),                           // 3: neighborhood.RequestCreate
	(*RequestCheckExistByProvinceIdAndCityId)(nil),  // 4: neighborhood.RequestCheckExistByProvinceIdAndCityId
	(*ResponseCheckExistByProvinceIdAndCityId)(nil), // 5: neighborhood.ResponseCheckExistByProvinceIdAndCityId
	(*RequestCheckValidNeighborhood)(nil),           // 6: neighborhood.RequestCheckValidNeighborhood
	(*ResponseCheckValidNeighborhood)(nil),          // 7: neighborhood.ResponseCheckValidNeighborhood
}
var file_neighborhood_neighborhood_proto_depIdxs = []int32{
	0, // 0: neighborhood.ResponseDetails.neighborhood:type_name -> neighborhood.Neighborhood
	0, // 1: neighborhood.RequestCreate.neighborhood:type_name -> neighborhood.Neighborhood
	0, // 2: neighborhood.RequestCheckExistByProvinceIdAndCityId.neighborhood:type_name -> neighborhood.Neighborhood
	3, // 3: neighborhood.neighborhoodService.Create:input_type -> neighborhood.RequestCreate
	4, // 4: neighborhood.neighborhoodService.CheckExistByProvinceIdAndCityId:input_type -> neighborhood.RequestCheckExistByProvinceIdAndCityId
	6, // 5: neighborhood.neighborhoodService.CheckValidNeighborhood:input_type -> neighborhood.RequestCheckValidNeighborhood
	2, // 6: neighborhood.neighborhoodService.Create:output_type -> neighborhood.ResponseDetails
	5, // 7: neighborhood.neighborhoodService.CheckExistByProvinceIdAndCityId:output_type -> neighborhood.ResponseCheckExistByProvinceIdAndCityId
	7, // 8: neighborhood.neighborhoodService.CheckValidNeighborhood:output_type -> neighborhood.ResponseCheckValidNeighborhood
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_neighborhood_neighborhood_proto_init() }
func file_neighborhood_neighborhood_proto_init() {
	if File_neighborhood_neighborhood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neighborhood_neighborhood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Neighborhood); i {
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
		file_neighborhood_neighborhood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_neighborhood_neighborhood_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_neighborhood_neighborhood_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_neighborhood_neighborhood_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCheckExistByProvinceIdAndCityId); i {
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
		file_neighborhood_neighborhood_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCheckExistByProvinceIdAndCityId); i {
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
		file_neighborhood_neighborhood_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCheckValidNeighborhood); i {
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
		file_neighborhood_neighborhood_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCheckValidNeighborhood); i {
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
			RawDescriptor: file_neighborhood_neighborhood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neighborhood_neighborhood_proto_goTypes,
		DependencyIndexes: file_neighborhood_neighborhood_proto_depIdxs,
		MessageInfos:      file_neighborhood_neighborhood_proto_msgTypes,
	}.Build()
	File_neighborhood_neighborhood_proto = out.File
	file_neighborhood_neighborhood_proto_rawDesc = nil
	file_neighborhood_neighborhood_proto_goTypes = nil
	file_neighborhood_neighborhood_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NeighborhoodServiceClient is the client API for NeighborhoodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NeighborhoodServiceClient interface {
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error)
	CheckExistByProvinceIdAndCityId(ctx context.Context, in *RequestCheckExistByProvinceIdAndCityId, opts ...grpc.CallOption) (*ResponseCheckExistByProvinceIdAndCityId, error)
	CheckValidNeighborhood(ctx context.Context, in *RequestCheckValidNeighborhood, opts ...grpc.CallOption) (*ResponseCheckValidNeighborhood, error)
}

type neighborhoodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNeighborhoodServiceClient(cc grpc.ClientConnInterface) NeighborhoodServiceClient {
	return &neighborhoodServiceClient{cc}
}

func (c *neighborhoodServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/neighborhood.neighborhoodService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neighborhoodServiceClient) CheckExistByProvinceIdAndCityId(ctx context.Context, in *RequestCheckExistByProvinceIdAndCityId, opts ...grpc.CallOption) (*ResponseCheckExistByProvinceIdAndCityId, error) {
	out := new(ResponseCheckExistByProvinceIdAndCityId)
	err := c.cc.Invoke(ctx, "/neighborhood.neighborhoodService/CheckExistByProvinceIdAndCityId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neighborhoodServiceClient) CheckValidNeighborhood(ctx context.Context, in *RequestCheckValidNeighborhood, opts ...grpc.CallOption) (*ResponseCheckValidNeighborhood, error) {
	out := new(ResponseCheckValidNeighborhood)
	err := c.cc.Invoke(ctx, "/neighborhood.neighborhoodService/CheckValidNeighborhood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NeighborhoodServiceServer is the server API for NeighborhoodService service.
type NeighborhoodServiceServer interface {
	Create(context.Context, *RequestCreate) (*ResponseDetails, error)
	CheckExistByProvinceIdAndCityId(context.Context, *RequestCheckExistByProvinceIdAndCityId) (*ResponseCheckExistByProvinceIdAndCityId, error)
	CheckValidNeighborhood(context.Context, *RequestCheckValidNeighborhood) (*ResponseCheckValidNeighborhood, error)
}

// UnimplementedNeighborhoodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNeighborhoodServiceServer struct {
}

func (*UnimplementedNeighborhoodServiceServer) Create(context.Context, *RequestCreate) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNeighborhoodServiceServer) CheckExistByProvinceIdAndCityId(context.Context, *RequestCheckExistByProvinceIdAndCityId) (*ResponseCheckExistByProvinceIdAndCityId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistByProvinceIdAndCityId not implemented")
}
func (*UnimplementedNeighborhoodServiceServer) CheckValidNeighborhood(context.Context, *RequestCheckValidNeighborhood) (*ResponseCheckValidNeighborhood, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckValidNeighborhood not implemented")
}

func RegisterNeighborhoodServiceServer(s *grpc.Server, srv NeighborhoodServiceServer) {
	s.RegisterService(&_NeighborhoodService_serviceDesc, srv)
}

func _NeighborhoodService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeighborhoodServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neighborhood.neighborhoodService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeighborhoodServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeighborhoodService_CheckExistByProvinceIdAndCityId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckExistByProvinceIdAndCityId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeighborhoodServiceServer).CheckExistByProvinceIdAndCityId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neighborhood.neighborhoodService/CheckExistByProvinceIdAndCityId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeighborhoodServiceServer).CheckExistByProvinceIdAndCityId(ctx, req.(*RequestCheckExistByProvinceIdAndCityId))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeighborhoodService_CheckValidNeighborhood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckValidNeighborhood)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeighborhoodServiceServer).CheckValidNeighborhood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neighborhood.neighborhoodService/CheckValidNeighborhood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeighborhoodServiceServer).CheckValidNeighborhood(ctx, req.(*RequestCheckValidNeighborhood))
	}
	return interceptor(ctx, in, info, handler)
}

var _NeighborhoodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neighborhood.neighborhoodService",
	HandlerType: (*NeighborhoodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NeighborhoodService_Create_Handler,
		},
		{
			MethodName: "CheckExistByProvinceIdAndCityId",
			Handler:    _NeighborhoodService_CheckExistByProvinceIdAndCityId_Handler,
		},
		{
			MethodName: "CheckValidNeighborhood",
			Handler:    _NeighborhoodService_CheckValidNeighborhood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neighborhood/neighborhood.proto",
}
