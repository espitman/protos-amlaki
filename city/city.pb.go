// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: city/city.proto

package city

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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title",validate:"alphanumunicode,required"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title" validate:"alphanumunicode,required"`
	// @inject_tag: json:"name",validate:"alpha,required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"alpha,required"`
	// @inject_tag: json:"province",validate:"alphanum,required"
	Province string `protobuf:"bytes,3,opt,name=province,proto3" json:"province" validate:"alphanum,required"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *City) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

type CityFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Name     string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Province *Province `protobuf:"bytes,4,opt,name=province,proto3" json:"province,omitempty"`
}

func (x *CityFull) Reset() {
	*x = CityFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityFull) ProtoMessage() {}

func (x *CityFull) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityFull.ProtoReflect.Descriptor instead.
func (*CityFull) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{1}
}

func (x *CityFull) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CityFull) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CityFull) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CityFull) GetProvince() *Province {
	if x != nil {
		return x.Province
	}
	return nil
}

type Province struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Province) Reset() {
	*x = Province{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Province) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Province) ProtoMessage() {}

func (x *Province) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Province.ProtoReflect.Descriptor instead.
func (*Province) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{2}
}

func (x *Province) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Province) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Province) GetName() string {
	if x != nil {
		return x.Name
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
		mi := &file_city_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[3]
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
	return file_city_city_proto_rawDescGZIP(), []int{3}
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

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	City    *City  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[4]
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
	return file_city_city_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseDetails) GetCity() *City {
	if x != nil {
		return x.City
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

	// @inject_tag: json:"city",validate:"required"
	City *City `protobuf:"bytes,1,opt,name=city,proto3" json:"city" validate:"required"`
	// @inject_tag: json:"-"
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"-"`
}

func (x *RequestCreate) Reset() {
	*x = RequestCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate) ProtoMessage() {}

func (x *RequestCreate) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[5]
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
	return file_city_city_proto_rawDescGZIP(), []int{5}
}

func (x *RequestCreate) GetCity() *City {
	if x != nil {
		return x.City
	}
	return nil
}

func (x *RequestCreate) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type ResponseIsExist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *ResponseIsExist) Reset() {
	*x = ResponseIsExist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseIsExist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseIsExist) ProtoMessage() {}

func (x *ResponseIsExist) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseIsExist.ProtoReflect.Descriptor instead.
func (*ResponseIsExist) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseIsExist) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

type RequestCheckProvinceHasCityByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId string `protobuf:"bytes,1,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
	CityName   string `protobuf:"bytes,2,opt,name=cityName,proto3" json:"cityName,omitempty"`
}

func (x *RequestCheckProvinceHasCityByName) Reset() {
	*x = RequestCheckProvinceHasCityByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCheckProvinceHasCityByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCheckProvinceHasCityByName) ProtoMessage() {}

func (x *RequestCheckProvinceHasCityByName) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCheckProvinceHasCityByName.ProtoReflect.Descriptor instead.
func (*RequestCheckProvinceHasCityByName) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{7}
}

func (x *RequestCheckProvinceHasCityByName) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *RequestCheckProvinceHasCityByName) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

type ResponseCheckProvinceHasCityByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
}

func (x *ResponseCheckProvinceHasCityByName) Reset() {
	*x = ResponseCheckProvinceHasCityByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCheckProvinceHasCityByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCheckProvinceHasCityByName) ProtoMessage() {}

func (x *ResponseCheckProvinceHasCityByName) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCheckProvinceHasCityByName.ProtoReflect.Descriptor instead.
func (*ResponseCheckProvinceHasCityByName) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{8}
}

func (x *ResponseCheckProvinceHasCityByName) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

type RequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"count",validate:"number,min=10,required"
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count" validate:"number,min=10,required"`
	// @inject_tag: json:"page",validate:"number,min=1,required"
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" validate:"number,min=1,required"`
}

func (x *RequestList) Reset() {
	*x = RequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestList) ProtoMessage() {}

func (x *RequestList) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestList.ProtoReflect.Descriptor instead.
func (*RequestList) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{9}
}

func (x *RequestList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RequestList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cities []*CityFull `protobuf:"bytes,1,rep,name=cities,proto3" json:"cities,omitempty"`
}

func (x *ResponseList) Reset() {
	*x = ResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseList) ProtoMessage() {}

func (x *ResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseList.ProtoReflect.Descriptor instead.
func (*ResponseList) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{10}
}

func (x *ResponseList) GetCities() []*CityFull {
	if x != nil {
		return x.Cities
	}
	return nil
}

var File_city_city_proto protoreflect.FileDescriptor

var file_city_city_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x4c, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x46, 0x75, 0x6c,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x22, 0x5f, 0x0a, 0x21, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x73, 0x43, 0x69, 0x74, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x22, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x73, 0x43, 0x69, 0x74,
	0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x61, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x46, 0x75,
	0x6c, 0x6c, 0x52, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xa3, 0x02, 0x0a, 0x0b, 0x63,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x1a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x73,
	0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x73, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x48, 0x61, 0x73, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_city_proto_rawDescOnce sync.Once
	file_city_city_proto_rawDescData = file_city_city_proto_rawDesc
)

func file_city_city_proto_rawDescGZIP() []byte {
	file_city_city_proto_rawDescOnce.Do(func() {
		file_city_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_city_proto_rawDescData)
	})
	return file_city_city_proto_rawDescData
}

var file_city_city_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_city_city_proto_goTypes = []interface{}{
	(*City)(nil),                               // 0: city.City
	(*CityFull)(nil),                           // 1: city.CityFull
	(*Province)(nil),                           // 2: city.Province
	(*RequestDetails)(nil),                     // 3: city.RequestDetails
	(*ResponseDetails)(nil),                    // 4: city.ResponseDetails
	(*RequestCreate)(nil),                      // 5: city.RequestCreate
	(*ResponseIsExist)(nil),                    // 6: city.ResponseIsExist
	(*RequestCheckProvinceHasCityByName)(nil),  // 7: city.RequestCheckProvinceHasCityByName
	(*ResponseCheckProvinceHasCityByName)(nil), // 8: city.ResponseCheckProvinceHasCityByName
	(*RequestList)(nil),                        // 9: city.RequestList
	(*ResponseList)(nil),                       // 10: city.ResponseList
}
var file_city_city_proto_depIdxs = []int32{
	2,  // 0: city.CityFull.province:type_name -> city.Province
	0,  // 1: city.ResponseDetails.city:type_name -> city.City
	0,  // 2: city.RequestCreate.city:type_name -> city.City
	1,  // 3: city.ResponseList.cities:type_name -> city.CityFull
	5,  // 4: city.cityService.Create:input_type -> city.RequestCreate
	3,  // 5: city.cityService.isExist:input_type -> city.RequestDetails
	7,  // 6: city.cityService.checkProvinceHasCityByName:input_type -> city.RequestCheckProvinceHasCityByName
	9,  // 7: city.cityService.list:input_type -> city.RequestList
	4,  // 8: city.cityService.Create:output_type -> city.ResponseDetails
	6,  // 9: city.cityService.isExist:output_type -> city.ResponseIsExist
	8,  // 10: city.cityService.checkProvinceHasCityByName:output_type -> city.ResponseCheckProvinceHasCityByName
	10, // 11: city.cityService.list:output_type -> city.ResponseList
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_city_city_proto_init() }
func file_city_city_proto_init() {
	if File_city_city_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_city_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*City); i {
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
		file_city_city_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityFull); i {
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
		file_city_city_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Province); i {
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
		file_city_city_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_city_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_city_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_city_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseIsExist); i {
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
		file_city_city_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCheckProvinceHasCityByName); i {
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
		file_city_city_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCheckProvinceHasCityByName); i {
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
		file_city_city_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestList); i {
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
		file_city_city_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseList); i {
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
			RawDescriptor: file_city_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_city_proto_goTypes,
		DependencyIndexes: file_city_city_proto_depIdxs,
		MessageInfos:      file_city_city_proto_msgTypes,
	}.Build()
	File_city_city_proto = out.File
	file_city_city_proto_rawDesc = nil
	file_city_city_proto_goTypes = nil
	file_city_city_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CityServiceClient interface {
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error)
	IsExist(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseIsExist, error)
	CheckProvinceHasCityByName(ctx context.Context, in *RequestCheckProvinceHasCityByName, opts ...grpc.CallOption) (*ResponseCheckProvinceHasCityByName, error)
	List(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*ResponseList, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/city.cityService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) IsExist(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseIsExist, error) {
	out := new(ResponseIsExist)
	err := c.cc.Invoke(ctx, "/city.cityService/isExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) CheckProvinceHasCityByName(ctx context.Context, in *RequestCheckProvinceHasCityByName, opts ...grpc.CallOption) (*ResponseCheckProvinceHasCityByName, error) {
	out := new(ResponseCheckProvinceHasCityByName)
	err := c.cc.Invoke(ctx, "/city.cityService/checkProvinceHasCityByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) List(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*ResponseList, error) {
	out := new(ResponseList)
	err := c.cc.Invoke(ctx, "/city.cityService/list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
type CityServiceServer interface {
	Create(context.Context, *RequestCreate) (*ResponseDetails, error)
	IsExist(context.Context, *RequestDetails) (*ResponseIsExist, error)
	CheckProvinceHasCityByName(context.Context, *RequestCheckProvinceHasCityByName) (*ResponseCheckProvinceHasCityByName, error)
	List(context.Context, *RequestList) (*ResponseList, error)
}

// UnimplementedCityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (*UnimplementedCityServiceServer) Create(context.Context, *RequestCreate) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCityServiceServer) IsExist(context.Context, *RequestDetails) (*ResponseIsExist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExist not implemented")
}
func (*UnimplementedCityServiceServer) CheckProvinceHasCityByName(context.Context, *RequestCheckProvinceHasCityByName) (*ResponseCheckProvinceHasCityByName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProvinceHasCityByName not implemented")
}
func (*UnimplementedCityServiceServer) List(context.Context, *RequestList) (*ResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterCityServiceServer(s *grpc.Server, srv CityServiceServer) {
	s.RegisterService(&_CityService_serviceDesc, srv)
}

func _CityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.cityService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_IsExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).IsExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.cityService/IsExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).IsExist(ctx, req.(*RequestDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_CheckProvinceHasCityByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckProvinceHasCityByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).CheckProvinceHasCityByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.cityService/CheckProvinceHasCityByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).CheckProvinceHasCityByName(ctx, req.(*RequestCheckProvinceHasCityByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.cityService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).List(ctx, req.(*RequestList))
	}
	return interceptor(ctx, in, info, handler)
}

var _CityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "city.cityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CityService_Create_Handler,
		},
		{
			MethodName: "isExist",
			Handler:    _CityService_IsExist_Handler,
		},
		{
			MethodName: "checkProvinceHasCityByName",
			Handler:    _CityService_CheckProvinceHasCityByName_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CityService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city/city.proto",
}
