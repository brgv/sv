// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/v1/Person.proto

package rpc

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

// Запрос позиции
type PersonId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PersonId) Reset() {
	*x = PersonId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_Person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonId) ProtoMessage() {}

func (x *PersonId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_Person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonId.ProtoReflect.Descriptor instead.
func (*PersonId) Descriptor() ([]byte, []int) {
	return file_proto_v1_Person_proto_rawDescGZIP(), []int{0}
}

func (x *PersonId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Данные позиции
type PersonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Название записи
	LastName   string `protobuf:"bytes,10,opt,name=lastName,proto3" json:"lastName,omitempty"`
	FirstName  string `protobuf:"bytes,11,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName string `protobuf:"bytes,12,opt,name=middleName,proto3" json:"middleName,omitempty"`
}

func (x *PersonData) Reset() {
	*x = PersonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_Person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonData) ProtoMessage() {}

func (x *PersonData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_Person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonData.ProtoReflect.Descriptor instead.
func (*PersonData) Descriptor() ([]byte, []int) {
	return file_proto_v1_Person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PersonData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PersonData) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LastName   string `protobuf:"bytes,10,opt,name=lastName,proto3" json:"lastName,omitempty"`
	FirstName  string `protobuf:"bytes,11,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName string `protobuf:"bytes,12,opt,name=middleName,proto3" json:"middleName,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_Person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_Person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_v1_Person_proto_rawDescGZIP(), []int{2}
}

func (x *Person) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

// Запрос списка записей
type PersonListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Количество записей
	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Смещение записей
	Start uint64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *PersonListFilter) Reset() {
	*x = PersonListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_Person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonListFilter) ProtoMessage() {}

func (x *PersonListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_Person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonListFilter.ProtoReflect.Descriptor instead.
func (*PersonListFilter) Descriptor() ([]byte, []int) {
	return file_proto_v1_Person_proto_rawDescGZIP(), []int{3}
}

func (x *PersonListFilter) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PersonListFilter) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

// Ответ с списком записей
type PersonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Список записей
	Data []*Person `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PersonList) Reset() {
	*x = PersonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_Person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonList) ProtoMessage() {}

func (x *PersonList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_Person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonList.ProtoReflect.Descriptor instead.
func (*PersonList) Descriptor() ([]byte, []int) {
	return file_proto_v1_Person_proto_rawDescGZIP(), []int{4}
}

func (x *PersonList) GetData() []*Person {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_v1_Person_proto protoreflect.FileDescriptor

var file_proto_v1_Person_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a,
	0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x0a, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x72, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa6, 0x03, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x2a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x76, 0x31,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a,
	0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x72, 0x67, 0x76, 0x2f, 0x73, 0x76, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_Person_proto_rawDescOnce sync.Once
	file_proto_v1_Person_proto_rawDescData = file_proto_v1_Person_proto_rawDesc
)

func file_proto_v1_Person_proto_rawDescGZIP() []byte {
	file_proto_v1_Person_proto_rawDescOnce.Do(func() {
		file_proto_v1_Person_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_Person_proto_rawDescData)
	})
	return file_proto_v1_Person_proto_rawDescData
}

var file_proto_v1_Person_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_v1_Person_proto_goTypes = []interface{}{
	(*PersonId)(nil),         // 0: v1.rpc.PersonId
	(*PersonData)(nil),       // 1: v1.rpc.PersonData
	(*Person)(nil),           // 2: v1.rpc.Person
	(*PersonListFilter)(nil), // 3: v1.rpc.PersonListFilter
	(*PersonList)(nil),       // 4: v1.rpc.PersonList
}
var file_proto_v1_Person_proto_depIdxs = []int32{
	2, // 0: v1.rpc.PersonList.data:type_name -> v1.rpc.Person
	0, // 1: v1.rpc.PersonService.GetPerson:input_type -> v1.rpc.PersonId
	0, // 2: v1.rpc.PersonService.DestroyPerson:input_type -> v1.rpc.PersonId
	1, // 3: v1.rpc.PersonService.CreatePerson:input_type -> v1.rpc.PersonData
	2, // 4: v1.rpc.PersonService.UpdatePerson:input_type -> v1.rpc.Person
	3, // 5: v1.rpc.PersonService.GetPersonList:input_type -> v1.rpc.PersonListFilter
	2, // 6: v1.rpc.PersonService.GetPerson:output_type -> v1.rpc.Person
	2, // 7: v1.rpc.PersonService.DestroyPerson:output_type -> v1.rpc.Person
	2, // 8: v1.rpc.PersonService.CreatePerson:output_type -> v1.rpc.Person
	2, // 9: v1.rpc.PersonService.UpdatePerson:output_type -> v1.rpc.Person
	4, // 10: v1.rpc.PersonService.GetPersonList:output_type -> v1.rpc.PersonList
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1_Person_proto_init() }
func file_proto_v1_Person_proto_init() {
	if File_proto_v1_Person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_Person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonId); i {
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
		file_proto_v1_Person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonData); i {
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
		file_proto_v1_Person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_proto_v1_Person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonListFilter); i {
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
		file_proto_v1_Person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonList); i {
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
			RawDescriptor: file_proto_v1_Person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_Person_proto_goTypes,
		DependencyIndexes: file_proto_v1_Person_proto_depIdxs,
		MessageInfos:      file_proto_v1_Person_proto_msgTypes,
	}.Build()
	File_proto_v1_Person_proto = out.File
	file_proto_v1_Person_proto_rawDesc = nil
	file_proto_v1_Person_proto_goTypes = nil
	file_proto_v1_Person_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonServiceClient interface {
	// Получить запись по уникальному идентификатору
	GetPerson(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Person, error)
	// Удалить запись по уникальному идентификатору
	DestroyPerson(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Person, error)
	// Создать запись
	CreatePerson(ctx context.Context, in *PersonData, opts ...grpc.CallOption) (*Person, error)
	// Обновить запись по уникальному идентификатору
	UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	// Получить список записей
	GetPersonList(ctx context.Context, in *PersonListFilter, opts ...grpc.CallOption) (*PersonList, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/v1.rpc.PersonService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) DestroyPerson(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/v1.rpc.PersonService/DestroyPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *PersonData, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/v1.rpc.PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/v1.rpc.PersonService/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPersonList(ctx context.Context, in *PersonListFilter, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, "/v1.rpc.PersonService/GetPersonList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
type PersonServiceServer interface {
	// Получить запись по уникальному идентификатору
	GetPerson(context.Context, *PersonId) (*Person, error)
	// Удалить запись по уникальному идентификатору
	DestroyPerson(context.Context, *PersonId) (*Person, error)
	// Создать запись
	CreatePerson(context.Context, *PersonData) (*Person, error)
	// Обновить запись по уникальному идентификатору
	UpdatePerson(context.Context, *Person) (*Person, error)
	// Получить список записей
	GetPersonList(context.Context, *PersonListFilter) (*PersonList, error)
}

// UnimplementedPersonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (*UnimplementedPersonServiceServer) GetPerson(context.Context, *PersonId) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedPersonServiceServer) DestroyPerson(context.Context, *PersonId) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyPerson not implemented")
}
func (*UnimplementedPersonServiceServer) CreatePerson(context.Context, *PersonData) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (*UnimplementedPersonServiceServer) UpdatePerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (*UnimplementedPersonServiceServer) GetPersonList(context.Context, *PersonListFilter) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonList not implemented")
}

func RegisterPersonServiceServer(s *grpc.Server, srv PersonServiceServer) {
	s.RegisterService(&_PersonService_serviceDesc, srv)
}

func _PersonService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rpc.PersonService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPerson(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_DestroyPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).DestroyPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rpc.PersonService/DestroyPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).DestroyPerson(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rpc.PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*PersonData))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rpc.PersonService/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).UpdatePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPersonList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPersonList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rpc.PersonService/GetPersonList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPersonList(ctx, req.(*PersonListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.rpc.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _PersonService_GetPerson_Handler,
		},
		{
			MethodName: "DestroyPerson",
			Handler:    _PersonService_DestroyPerson_Handler,
		},
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _PersonService_UpdatePerson_Handler,
		},
		{
			MethodName: "GetPersonList",
			Handler:    _PersonService_GetPersonList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/Person.proto",
}
