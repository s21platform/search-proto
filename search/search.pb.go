// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: search.proto

package search

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

// Данные запроса собществ
type GetSocietyIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// полное имя сообщества или часть имени, по которому будет осуществлен поиск
	PartName string `protobuf:"bytes,1,opt,name=part_name,json=partName,proto3" json:"part_name,omitempty"`
	Limit    int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetSocietyIn) Reset() {
	*x = GetSocietyIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSocietyIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocietyIn) ProtoMessage() {}

func (x *GetSocietyIn) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocietyIn.ProtoReflect.Descriptor instead.
func (*GetSocietyIn) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *GetSocietyIn) GetPartName() string {
	if x != nil {
		return x.PartName
	}
	return ""
}

func (x *GetSocietyIn) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetSocietyIn) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Объект описывающий сообщество
type Society struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Имя сообщества
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Аватарка сообщества
	AvatarLink string `protobuf:"bytes,2,opt,name=avatar_link,json=avatarLink,proto3" json:"avatar_link,omitempty"`
	// Признак публичности сообщества
	IsPrivate bool `protobuf:"varint,3,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	// Описание сообщества
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Society) Reset() {
	*x = Society{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Society) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Society) ProtoMessage() {}

func (x *Society) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Society.ProtoReflect.Descriptor instead.
func (*Society) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *Society) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Society) GetAvatarLink() string {
	if x != nil {
		return x.AvatarLink
	}
	return ""
}

func (x *Society) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Society) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Объект ответа на запрос поиска
type GetSocietyOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Массив объектов сообществ, удовлетворяющих поску
	Societies []*Society `protobuf:"bytes,1,rep,name=societies,proto3" json:"societies,omitempty"`
	// Общее количество сообществ
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetSocietyOut) Reset() {
	*x = GetSocietyOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSocietyOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocietyOut) ProtoMessage() {}

func (x *GetSocietyOut) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocietyOut.ProtoReflect.Descriptor instead.
func (*GetSocietyOut) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *GetSocietyOut) GetSocieties() []*Society {
	if x != nil {
		return x.Societies
	}
	return nil
}

func (x *GetSocietyOut) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// Объект запроса пользователей
type GetUserWithOffsetIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// количество запрашиваемых записей
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// offset записей
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// nickname пользователя которого ищем
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *GetUserWithOffsetIn) Reset() {
	*x = GetUserWithOffsetIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWithOffsetIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithOffsetIn) ProtoMessage() {}

func (x *GetUserWithOffsetIn) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithOffsetIn.ProtoReflect.Descriptor instead.
func (*GetUserWithOffsetIn) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserWithOffsetIn) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserWithOffsetIn) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserWithOffsetIn) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

// Объект ответа на запрос пользователей
type GetUserWithOffsetOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Срез пользователей
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// Количество возвращаемых пользователей
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetUserWithOffsetOut) Reset() {
	*x = GetUserWithOffsetOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWithOffsetOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithOffsetOut) ProtoMessage() {}

func (x *GetUserWithOffsetOut) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithOffsetOut.ProtoReflect.Descriptor instead.
func (*GetUserWithOffsetOut) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserWithOffsetOut) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetUserWithOffsetOut) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// Объект пользователь
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ник пользователя
	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// UUID пользователя
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// ссылка на последний аватар пользователя
	AvatarLink string `protobuf:"bytes,3,opt,name=avatar_link,json=avatarLink,proto3" json:"avatar_link,omitempty"`
	// имя пользователя
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// фамилия пользователя
	Surname string `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetAvatarLink() string {
	if x != nil {
		return x.AvatarLink
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x49, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7f, 0x0a, 0x07, 0x53, 0x6f, 0x63,
	0x69, 0x65, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x73,
	0x6f, 0x63, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x52, 0x09, 0x73, 0x6f, 0x63, 0x69, 0x65, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x49, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4f,
	0x75, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x82, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x12, 0x0d, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_search_proto_goTypes = []interface{}{
	(*GetSocietyIn)(nil),         // 0: GetSocietyIn
	(*Society)(nil),              // 1: Society
	(*GetSocietyOut)(nil),        // 2: GetSocietyOut
	(*GetUserWithOffsetIn)(nil),  // 3: GetUserWithOffsetIn
	(*GetUserWithOffsetOut)(nil), // 4: GetUserWithOffsetOut
	(*User)(nil),                 // 5: User
}
var file_search_proto_depIdxs = []int32{
	1, // 0: GetSocietyOut.societies:type_name -> Society
	5, // 1: GetUserWithOffsetOut.users:type_name -> User
	0, // 2: SearchService.GetSociety:input_type -> GetSocietyIn
	3, // 3: SearchService.GetUserWithOffset:input_type -> GetUserWithOffsetIn
	2, // 4: SearchService.GetSociety:output_type -> GetSocietyOut
	4, // 5: SearchService.GetUserWithOffset:output_type -> GetUserWithOffsetOut
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSocietyIn); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Society); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSocietyOut); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWithOffsetIn); i {
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
		file_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWithOffsetOut); i {
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
		file_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
