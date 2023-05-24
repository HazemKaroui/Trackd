// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: tracked.proto

package tracker

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

type CreateTrackedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Cover       string `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Genre       string `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Rating      string `protobuf:"bytes,7,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *CreateTrackedRequest) Reset() {
	*x = CreateTrackedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrackedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrackedRequest) ProtoMessage() {}

func (x *CreateTrackedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrackedRequest.ProtoReflect.Descriptor instead.
func (*CreateTrackedRequest) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTrackedRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateTrackedRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTrackedRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTrackedRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreateTrackedRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateTrackedRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *CreateTrackedRequest) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

type CreateTrackedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateTrackedResponse) Reset() {
	*x = CreateTrackedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrackedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrackedResponse) ProtoMessage() {}

func (x *CreateTrackedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrackedResponse.ProtoReflect.Descriptor instead.
func (*CreateTrackedResponse) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTrackedResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTrackedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTrackedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetTrackedsRequest) Reset() {
	*x = GetTrackedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackedsRequest) ProtoMessage() {}

func (x *GetTrackedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackedsRequest.ProtoReflect.Descriptor instead.
func (*GetTrackedsRequest) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrackedsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetTrackedsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTrackedsResponse) Reset() {
	*x = GetTrackedsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackedsResponse) ProtoMessage() {}

func (x *GetTrackedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackedsResponse.ProtoReflect.Descriptor instead.
func (*GetTrackedsResponse) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrackedsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetTrackedsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTrackedsResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type UpdateTrackedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Cover       string `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Genre       string `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Rating      string `protobuf:"bytes,7,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *UpdateTrackedRequest) Reset() {
	*x = UpdateTrackedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrackedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrackedRequest) ProtoMessage() {}

func (x *UpdateTrackedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrackedRequest.ProtoReflect.Descriptor instead.
func (*UpdateTrackedRequest) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTrackedRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTrackedRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTrackedRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTrackedRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *UpdateTrackedRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateTrackedRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *UpdateTrackedRequest) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

type UpdateTrackedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateTrackedResponse) Reset() {
	*x = UpdateTrackedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrackedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrackedResponse) ProtoMessage() {}

func (x *UpdateTrackedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrackedResponse.ProtoReflect.Descriptor instead.
func (*UpdateTrackedResponse) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTrackedResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateTrackedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteTrackedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTrackedRequest) Reset() {
	*x = DeleteTrackedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackedRequest) ProtoMessage() {}

func (x *DeleteTrackedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrackedRequest.ProtoReflect.Descriptor instead.
func (*DeleteTrackedRequest) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTrackedRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTrackedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTrackedResponse) Reset() {
	*x = DeleteTrackedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracked_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackedResponse) ProtoMessage() {}

func (x *DeleteTrackedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracked_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrackedResponse.ProtoReflect.Descriptor instead.
func (*DeleteTrackedResponse) Descriptor() ([]byte, []int) {
	return file_tracked_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTrackedResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteTrackedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_tracked_proto protoreflect.FileDescriptor

var file_tracked_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x22, 0xc6, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x4b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5d, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd2,
	0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x64, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracked_proto_rawDescOnce sync.Once
	file_tracked_proto_rawDescData = file_tracked_proto_rawDesc
)

func file_tracked_proto_rawDescGZIP() []byte {
	file_tracked_proto_rawDescOnce.Do(func() {
		file_tracked_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracked_proto_rawDescData)
	})
	return file_tracked_proto_rawDescData
}

var file_tracked_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tracked_proto_goTypes = []interface{}{
	(*CreateTrackedRequest)(nil),  // 0: tracker.CreateTrackedRequest
	(*CreateTrackedResponse)(nil), // 1: tracker.CreateTrackedResponse
	(*GetTrackedsRequest)(nil),    // 2: tracker.GetTrackedsRequest
	(*GetTrackedsResponse)(nil),   // 3: tracker.GetTrackedsResponse
	(*UpdateTrackedRequest)(nil),  // 4: tracker.UpdateTrackedRequest
	(*UpdateTrackedResponse)(nil), // 5: tracker.UpdateTrackedResponse
	(*DeleteTrackedRequest)(nil),  // 6: tracker.DeleteTrackedRequest
	(*DeleteTrackedResponse)(nil), // 7: tracker.DeleteTrackedResponse
}
var file_tracked_proto_depIdxs = []int32{
	0, // 0: tracker.TrackedService.CreateTracked:input_type -> tracker.CreateTrackedRequest
	2, // 1: tracker.TrackedService.GetTrackeds:input_type -> tracker.GetTrackedsRequest
	4, // 2: tracker.TrackedService.UpdateTracked:input_type -> tracker.UpdateTrackedRequest
	6, // 3: tracker.TrackedService.DeleteTracked:input_type -> tracker.DeleteTrackedRequest
	1, // 4: tracker.TrackedService.CreateTracked:output_type -> tracker.CreateTrackedResponse
	3, // 5: tracker.TrackedService.GetTrackeds:output_type -> tracker.GetTrackedsResponse
	5, // 6: tracker.TrackedService.UpdateTracked:output_type -> tracker.UpdateTrackedResponse
	7, // 7: tracker.TrackedService.DeleteTracked:output_type -> tracker.DeleteTrackedResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tracked_proto_init() }
func file_tracked_proto_init() {
	if File_tracked_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracked_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTrackedRequest); i {
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
		file_tracked_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTrackedResponse); i {
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
		file_tracked_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackedsRequest); i {
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
		file_tracked_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackedsResponse); i {
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
		file_tracked_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrackedRequest); i {
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
		file_tracked_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrackedResponse); i {
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
		file_tracked_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTrackedRequest); i {
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
		file_tracked_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTrackedResponse); i {
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
			RawDescriptor: file_tracked_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracked_proto_goTypes,
		DependencyIndexes: file_tracked_proto_depIdxs,
		MessageInfos:      file_tracked_proto_msgTypes,
	}.Build()
	File_tracked_proto = out.File
	file_tracked_proto_rawDesc = nil
	file_tracked_proto_goTypes = nil
	file_tracked_proto_depIdxs = nil
}