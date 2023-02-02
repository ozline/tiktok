// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: gateway.proto

package user

import (
	_ "github.com/ozline/tiktok/api-gateway/biz/model/api"
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

type Code int32

const (
	Code_Success      Code = 0
	Code_ParamInvalid Code = 1
	Code_DBErr        Code = 2
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "Success",
		1: "ParamInvalid",
		2: "DBErr",
	}
	Code_value = map[string]int32{
		"Success":      0,
		"ParamInvalid": 1,
		"DBErr":        2,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_Unknown Gender = 0
	Gender_Male    Gender = 1
	Gender_Female  Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "Unknown",
		1: "Male",
		2: "Female",
	}
	Gender_value = map[string]int32{
		"Unknown": 0,
		"Male":    1,
		"Female":  2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty" form:"UserID" query:"UserID"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" form:"Name" query:"Name"`
	Gender    Gender `protobuf:"varint,3,opt,name=Gender,proto3,enum=user.Gender" json:"Gender,omitempty" form:"Gender" query:"Gender"`
	Age       int64  `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty" form:"Age" query:"Age"`
	Introduce string `protobuf:"bytes,5,opt,name=Introduce,proto3" json:"Introduce,omitempty" form:"Introduce" query:"Introduce"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
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
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_Unknown
}

func (x *User) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name,omitempty" form:"name" vd:"(len($) > 0 && len($) < 100)"`
	Gender    Gender `protobuf:"varint,2,opt,name=Gender,proto3,enum=user.Gender" json:"gender,omitempty" form:"gender" vd:"($ == 1||$ == 2)"`
	Age       int64  `protobuf:"varint,3,opt,name=Age,proto3" json:"age,omitempty" form:"age" vd:"$>0"`
	Introduce string `protobuf:"bytes,4,opt,name=Introduce,proto3" json:"introduce,omitempty" form:"introduce" vd:"(len($) > 0 && len($) < 1000)"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserReq) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_Unknown
}

func (x *CreateUserReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateUserReq) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

type CreateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
}

func (x *CreateUserResp) Reset() {
	*x = CreateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResp) ProtoMessage() {}

func (x *CreateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResp.ProtoReflect.Descriptor instead.
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *CreateUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type QueryUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword  string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"keyword,omitempty" form:"keyword"`
	Page     int64  `protobuf:"varint,2,opt,name=Page,proto3" json:"page,omitempty" form:"page" vd:"$>0"`
	PageSize int64  `protobuf:"varint,3,opt,name=PageSize,proto3" json:"page_size,omitempty" form:"page_size" vd:"($ > 0 || $ <= 100)"`
}

func (x *QueryUserReq) Reset() {
	*x = QueryUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserReq) ProtoMessage() {}

func (x *QueryUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserReq.ProtoReflect.Descriptor instead.
func (*QueryUserReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *QueryUserReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *QueryUserReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryUserReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  Code    `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg   string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
	User  []*User `protobuf:"bytes,3,rep,name=User,proto3" json:"User" form:"User" query:"User"`
	Total int64   `protobuf:"varint,4,opt,name=Total,proto3" json:"Total,omitempty" form:"Total" query:"Total"`
}

func (x *QueryUserResp) Reset() {
	*x = QueryUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserResp) ProtoMessage() {}

func (x *QueryUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserResp.ProtoReflect.Descriptor instead.
func (*QueryUserResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *QueryUserResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *QueryUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *QueryUserResp) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *QueryUserResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty" path:"user_id" vd:"$>0"`
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type DeleteUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
}

func (x *DeleteUserResp) Reset() {
	*x = DeleteUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResp) ProtoMessage() {}

func (x *DeleteUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResp.ProtoReflect.Descriptor instead.
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUserResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *DeleteUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty" path:"user_id" vd:"$>0"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name,omitempty" form:"name" vd:"(len($) > 0 && len($) < 100)"`
	Gender    Gender `protobuf:"varint,3,opt,name=Gender,proto3,enum=user.Gender" json:"gender,omitempty" form:"gender" vd:"($ == 1||$ == 2)"`
	Age       int64  `protobuf:"varint,4,opt,name=Age,proto3" json:"age,omitempty" form:"age" vd:"$>0"`
	Introduce string `protobuf:"bytes,5,opt,name=Introduce,proto3" form:"introduce" form:"introduce" json:"introduce,omitempty" vd:"(len($) > 0 && len($) < 1000)"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserReq) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_Unknown
}

func (x *UpdateUserReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserReq) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

type UpdateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
}

func (x *UpdateUserResp) Reset() {
	*x = UpdateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResp) ProtoMessage() {}

func (x *UpdateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResp.ProtoReflect.Descriptor instead.
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *UpdateUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type OtherResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
}

func (x *OtherResp) Reset() {
	*x = OtherResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherResp) ProtoMessage() {}

func (x *OtherResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherResp.ProtoReflect.Descriptor instead.
func (*OtherResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *OtherResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x44, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x30, 0xca, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xda, 0xbb, 0x18,
	0x1c, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20,
	0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x29, 0xe2, 0xbb, 0x18,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x28, 0xca, 0xbb, 0x18, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0xda, 0xbb, 0x18, 0x10, 0x28, 0x24, 0x20, 0x3d, 0x3d, 0x20, 0x31,
	0x7c, 0x7c, 0x24, 0x20, 0x3d, 0x3d, 0x20, 0x32, 0x29, 0xe2, 0xbb, 0x18, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x03, 0x41,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0xca, 0xbb, 0x18, 0x03, 0x61, 0x67,
	0x65, 0xda, 0xbb, 0x18, 0x03, 0x24, 0x3e, 0x30, 0xe2, 0xbb, 0x18, 0x03, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x41, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0xca, 0xbb, 0x18, 0x09, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0xda, 0xbb, 0x18, 0x1d, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24,
	0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20,
	0x3c, 0x20, 0x31, 0x30, 0x30, 0x30, 0x29, 0xe2, 0xbb, 0x18, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x22,
	0x42, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xca, 0xbb, 0x18, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0xe2, 0xbb, 0x18, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x17, 0xca, 0xbb, 0x18, 0x04, 0x70, 0x61, 0x67, 0x65, 0xda, 0xbb,
	0x18, 0x03, 0x24, 0x3e, 0x30, 0xe2, 0xbb, 0x18, 0x04, 0x70, 0x61, 0x67, 0x65, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x31, 0xca, 0xbb, 0x18, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0xda, 0xbb, 0x18, 0x13, 0x28, 0x24, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x7c,
	0x7c, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x31, 0x30, 0x30, 0x29, 0xe2, 0xbb, 0x18, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x77, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x3b, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x12, 0xd2, 0xbb,
	0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xda, 0xbb, 0x18, 0x03, 0x24, 0x3e, 0x30,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x42, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0xd5, 0x02, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2a,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x12,
	0xd2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xda, 0xbb, 0x18, 0x03, 0x24,
	0x3e, 0x30, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xca, 0xbb, 0x18, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0xda, 0xbb, 0x18, 0x1c, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20,
	0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30,
	0x30, 0x29, 0xe2, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x28,
	0xca, 0xbb, 0x18, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0xda, 0xbb, 0x18, 0x10, 0x28, 0x24,
	0x20, 0x3d, 0x3d, 0x20, 0x31, 0x7c, 0x7c, 0x24, 0x20, 0x3d, 0x3d, 0x20, 0x32, 0x29, 0xe2, 0xbb,
	0x18, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0xca,
	0xbb, 0x18, 0x03, 0x61, 0x67, 0x65, 0xda, 0xbb, 0x18, 0x03, 0x24, 0x3e, 0x30, 0xe2, 0xbb, 0x18,
	0x03, 0x61, 0x67, 0x65, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0xca, 0xbb,
	0x18, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0xda, 0xbb, 0x18, 0x1d, 0x28,
	0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65,
	0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x30, 0x29, 0xe2, 0xbb, 0x18, 0x09,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x1d, 0x0a, 0x09, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x2a, 0x30, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x42, 0x45, 0x72, 0x72, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x65,
	0x6d, 0x61, 0x6c, 0x65, 0x10, 0x02, 0x32, 0xf3, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0xd2, 0xc1, 0x18, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x11,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xd2, 0xc1, 0x18, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x5d,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x5d, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c,
	0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x2f, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6c, 0x69, 0x6e,
	0x65, 0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gateway_proto_goTypes = []interface{}{
	(Code)(0),              // 0: user.Code
	(Gender)(0),            // 1: user.Gender
	(*User)(nil),           // 2: user.User
	(*CreateUserReq)(nil),  // 3: user.CreateUserReq
	(*CreateUserResp)(nil), // 4: user.CreateUserResp
	(*QueryUserReq)(nil),   // 5: user.QueryUserReq
	(*QueryUserResp)(nil),  // 6: user.QueryUserResp
	(*DeleteUserReq)(nil),  // 7: user.DeleteUserReq
	(*DeleteUserResp)(nil), // 8: user.DeleteUserResp
	(*UpdateUserReq)(nil),  // 9: user.UpdateUserReq
	(*UpdateUserResp)(nil), // 10: user.UpdateUserResp
	(*OtherResp)(nil),      // 11: user.OtherResp
}
var file_gateway_proto_depIdxs = []int32{
	1,  // 0: user.User.Gender:type_name -> user.Gender
	1,  // 1: user.CreateUserReq.Gender:type_name -> user.Gender
	0,  // 2: user.CreateUserResp.Code:type_name -> user.Code
	0,  // 3: user.QueryUserResp.Code:type_name -> user.Code
	2,  // 4: user.QueryUserResp.User:type_name -> user.User
	0,  // 5: user.DeleteUserResp.Code:type_name -> user.Code
	1,  // 6: user.UpdateUserReq.Gender:type_name -> user.Gender
	0,  // 7: user.UpdateUserResp.Code:type_name -> user.Code
	3,  // 8: user.UserService.CreateUserResponse:input_type -> user.CreateUserReq
	5,  // 9: user.UserService.QueryUserResponse:input_type -> user.QueryUserReq
	9,  // 10: user.UserService.UpdateUserResponse:input_type -> user.UpdateUserReq
	7,  // 11: user.UserService.DeleteUserResponse:input_type -> user.DeleteUserReq
	4,  // 12: user.UserService.CreateUserResponse:output_type -> user.CreateUserResp
	6,  // 13: user.UserService.QueryUserResponse:output_type -> user.QueryUserResp
	10, // 14: user.UserService.UpdateUserResponse:output_type -> user.UpdateUserResp
	8,  // 15: user.UserService.DeleteUserResponse:output_type -> user.DeleteUserResp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResp); i {
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
		file_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserReq); i {
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
		file_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserResp); i {
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
		file_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReq); i {
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
		file_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResp); i {
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
		file_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
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
		file_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResp); i {
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
		file_gateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherResp); i {
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
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		EnumInfos:         file_gateway_proto_enumTypes,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}