// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: idl/api/message/social.proto

package social

import (
	model "github.com/ozline/tiktok/api-gateway/biz/model/model"
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

// Social API
type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                                          // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty" form:"to_user_id" query:"to_user_id"`       // 对方用户ID
	ActionType int64  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty" form:"action_type" query:"action_type"` // 1-关注，2-取消关注
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type RelationFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"` // 用户ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                      // 用户鉴权token
}

func (x *RelationFollowListRequest) Reset() {
	*x = RelationFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListRequest) ProtoMessage() {}

func (x *RelationFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowListRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{2}
}

func (x *RelationFollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
	UserList   []*model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list" form:"user_list" query:"user_list"`                      // 用户信息列表
}

func (x *RelationFollowListResponse) Reset() {
	*x = RelationFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListResponse) ProtoMessage() {}

func (x *RelationFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowListResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowListResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowListResponse) GetUserList() []*model.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFollowerListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"` // 用户ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                      // 用户鉴权token
}

func (x *RelationFollowerListRequest) Reset() {
	*x = RelationFollowerListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListRequest) ProtoMessage() {}

func (x *RelationFollowerListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowerListRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{4}
}

func (x *RelationFollowerListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowerListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFollowerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
	UserList   []*model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list" form:"user_list" query:"user_list"`                      // 用户列表
}

func (x *RelationFollowerListResponse) Reset() {
	*x = RelationFollowerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListResponse) ProtoMessage() {}

func (x *RelationFollowerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowerListResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{5}
}

func (x *RelationFollowerListResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowerListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowerListResponse) GetUserList() []*model.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"` // 用户ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                      // 用户鉴权token
}

func (x *RelationFriendListRequest) Reset() {
	*x = RelationFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendListRequest) ProtoMessage() {}

func (x *RelationFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendListRequest.ProtoReflect.Descriptor instead.
func (*RelationFriendListRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{6}
}

func (x *RelationFriendListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFriendListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
	UserList   []*model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list" form:"user_list" query:"user_list"`                      // 用户列表
}

func (x *RelationFriendListResponse) Reset() {
	*x = RelationFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendListResponse) ProtoMessage() {}

func (x *RelationFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendListResponse.ProtoReflect.Descriptor instead.
func (*RelationFriendListResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{7}
}

func (x *RelationFriendListResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFriendListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFriendListResponse) GetUserList() []*model.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type MessageSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                                          // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty" form:"to_user_id" query:"to_user_id"`       // 对方用户ID
	ActionType int64  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty" form:"action_type" query:"action_type"` // 1-发送消息，2-删除消息
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty" form:"content" query:"content"`                                  // 消息内容
}

func (x *MessageSendRequest) Reset() {
	*x = MessageSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendRequest) ProtoMessage() {}

func (x *MessageSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendRequest.ProtoReflect.Descriptor instead.
func (*MessageSendRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{8}
}

func (x *MessageSendRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MessageSendRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *MessageSendRequest) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *MessageSendRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessageSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
}

func (x *MessageSendResponse) Reset() {
	*x = MessageSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendResponse) ProtoMessage() {}

func (x *MessageSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendResponse.ProtoReflect.Descriptor instead.
func (*MessageSendResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{9}
}

func (x *MessageSendResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *MessageSendResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type MessageChatMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`                                    // 用户鉴权token
	ToUserId int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty" form:"to_user_id" query:"to_user_id"` // 对方用户ID
}

func (x *MessageChatMsgRequest) Reset() {
	*x = MessageChatMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageChatMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatMsgRequest) ProtoMessage() {}

func (x *MessageChatMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatMsgRequest.ProtoReflect.Descriptor instead.
func (*MessageChatMsgRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{10}
}

func (x *MessageChatMsgRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MessageChatMsgRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

type MessageChatMsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int64            `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg   string           `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
	MessageList []*model.Message `protobuf:"bytes,3,rep,name=message_list,json=messageList,proto3" json:"message_list" form:"message_list" query:"message_list"`       // 消息列表
}

func (x *MessageChatMsgResponse) Reset() {
	*x = MessageChatMsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_message_social_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageChatMsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatMsgResponse) ProtoMessage() {}

func (x *MessageChatMsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_message_social_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatMsgResponse.ProtoReflect.Descriptor instead.
func (*MessageChatMsgResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_message_social_proto_rawDescGZIP(), []int{11}
}

func (x *MessageChatMsgResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *MessageChatMsgResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *MessageChatMsgResponse) GetMessageList() []*model.Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

var File_idl_api_message_social_proto protoreflect.FileDescriptor

var file_idl_api_message_social_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x22, 0x4a, 0x0a, 0x19, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a,
	0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x1b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x86, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x19, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x83, 0x01,
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x4b, 0x0a, 0x15, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x2f, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x7a, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_api_message_social_proto_rawDescOnce sync.Once
	file_idl_api_message_social_proto_rawDescData = file_idl_api_message_social_proto_rawDesc
)

func file_idl_api_message_social_proto_rawDescGZIP() []byte {
	file_idl_api_message_social_proto_rawDescOnce.Do(func() {
		file_idl_api_message_social_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_api_message_social_proto_rawDescData)
	})
	return file_idl_api_message_social_proto_rawDescData
}

var file_idl_api_message_social_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_idl_api_message_social_proto_goTypes = []interface{}{
	(*RelationActionRequest)(nil),        // 0: api.RelationActionRequest
	(*RelationActionResponse)(nil),       // 1: api.RelationActionResponse
	(*RelationFollowListRequest)(nil),    // 2: api.RelationFollowListRequest
	(*RelationFollowListResponse)(nil),   // 3: api.RelationFollowListResponse
	(*RelationFollowerListRequest)(nil),  // 4: api.RelationFollowerListRequest
	(*RelationFollowerListResponse)(nil), // 5: api.RelationFollowerListResponse
	(*RelationFriendListRequest)(nil),    // 6: api.RelationFriendListRequest
	(*RelationFriendListResponse)(nil),   // 7: api.RelationFriendListResponse
	(*MessageSendRequest)(nil),           // 8: api.MessageSendRequest
	(*MessageSendResponse)(nil),          // 9: api.MessageSendResponse
	(*MessageChatMsgRequest)(nil),        // 10: api.MessageChatMsgRequest
	(*MessageChatMsgResponse)(nil),       // 11: api.MessageChatMsgResponse
	(*model.User)(nil),                   // 12: api.User
	(*model.Message)(nil),                // 13: api.Message
}
var file_idl_api_message_social_proto_depIdxs = []int32{
	12, // 0: api.RelationFollowListResponse.user_list:type_name -> api.User
	12, // 1: api.RelationFollowerListResponse.user_list:type_name -> api.User
	12, // 2: api.RelationFriendListResponse.user_list:type_name -> api.User
	13, // 3: api.MessageChatMsgResponse.message_list:type_name -> api.Message
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_idl_api_message_social_proto_init() }
func file_idl_api_message_social_proto_init() {
	if File_idl_api_message_social_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_api_message_social_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
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
		file_idl_api_message_social_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListResponse); i {
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
		file_idl_api_message_social_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListResponse); i {
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
		file_idl_api_message_social_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendListRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendListResponse); i {
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
		file_idl_api_message_social_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendResponse); i {
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
		file_idl_api_message_social_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageChatMsgRequest); i {
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
		file_idl_api_message_social_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageChatMsgResponse); i {
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
			RawDescriptor: file_idl_api_message_social_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_idl_api_message_social_proto_goTypes,
		DependencyIndexes: file_idl_api_message_social_proto_depIdxs,
		MessageInfos:      file_idl_api_message_social_proto_msgTypes,
	}.Build()
	File_idl_api_message_social_proto = out.File
	file_idl_api_message_social_proto_rawDesc = nil
	file_idl_api_message_social_proto_goTypes = nil
	file_idl_api_message_social_proto_depIdxs = nil
}
