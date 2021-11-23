// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/wx/v1/wx.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//    属性	类型	默认值	必填	说明
//    appid	string		是	小程序 appId
//    secret	string		是	小程序 appSecret
//    js_code	string		是	登录时获取的 code
//    grant_type	string		是	授权类型，此处只需填写 authorization_code
type Code2SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Secret    string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	JsCode    string `protobuf:"bytes,3,opt,name=js_code,json=jsCode,proto3" json:"js_code,omitempty"`
	GrantType string `protobuf:"bytes,4,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
}

func (x *Code2SessionRequest) Reset() {
	*x = Code2SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionRequest) ProtoMessage() {}

func (x *Code2SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionRequest.ProtoReflect.Descriptor instead.
func (*Code2SessionRequest) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{0}
}

func (x *Code2SessionRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Code2SessionRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Code2SessionRequest) GetJsCode() string {
	if x != nil {
		return x.JsCode
	}
	return ""
}

func (x *Code2SessionRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

//    属性	类型	说明
//    openid	string	用户唯一标识
//    session_key	string	会话密钥
//    unionid	string	用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
//    errcode	number	错误码
//    errmsg	string	错误信息
type Code2SessionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid     string `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	SessionKey string `protobuf:"bytes,2,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
	Unionid    string `protobuf:"bytes,3,opt,name=unionid,proto3" json:"unionid,omitempty"`
	Errcode    int64  `protobuf:"varint,4,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Errmsg     string `protobuf:"bytes,5,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *Code2SessionReply) Reset() {
	*x = Code2SessionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionReply) ProtoMessage() {}

func (x *Code2SessionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionReply.ProtoReflect.Descriptor instead.
func (*Code2SessionReply) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{1}
}

func (x *Code2SessionReply) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *Code2SessionReply) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *Code2SessionReply) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *Code2SessionReply) GetErrcode() int64 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *Code2SessionReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

type SubscribeSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string               `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Touser      string               `protobuf:"bytes,2,opt,name=touser,proto3" json:"touser,omitempty"`
	TemplateId  string               `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Data        map[string]*TmplData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubscribeSendRequest) Reset() {
	*x = SubscribeSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeSendRequest) ProtoMessage() {}

func (x *SubscribeSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeSendRequest.ProtoReflect.Descriptor instead.
func (*SubscribeSendRequest) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeSendRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SubscribeSendRequest) GetTouser() string {
	if x != nil {
		return x.Touser
	}
	return ""
}

func (x *SubscribeSendRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SubscribeSendRequest) GetData() map[string]*TmplData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TmplData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TmplData) Reset() {
	*x = TmplData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TmplData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TmplData) ProtoMessage() {}

func (x *TmplData) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TmplData.ProtoReflect.Descriptor instead.
func (*TmplData) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{3}
}

func (x *TmplData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SubscribeSendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errcode int64  `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *SubscribeSendReply) Reset() {
	*x = SubscribeSendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeSendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeSendReply) ProtoMessage() {}

func (x *SubscribeSendReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeSendReply.ProtoReflect.Descriptor instead.
func (*SubscribeSendReply) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeSendReply) GetErrcode() int64 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *SubscribeSendReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

type GetAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantType string `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	Appid     string `protobuf:"bytes,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *GetAccessTokenRequest) Reset() {
	*x = GetAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenRequest) ProtoMessage() {}

func (x *GetAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccessTokenRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *GetAccessTokenRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetAccessTokenRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type GetAccessTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn   int64  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	Errcode     int64  `protobuf:"varint,3,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Errmsg      string `protobuf:"bytes,4,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *GetAccessTokenReply) Reset() {
	*x = GetAccessTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wx_v1_wx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenReply) ProtoMessage() {}

func (x *GetAccessTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_wx_v1_wx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenReply.ProtoReflect.Descriptor instead.
func (*GetAccessTokenReply) Descriptor() ([]byte, []int) {
	return file_api_wx_v1_wx_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccessTokenReply) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetAccessTokenReply) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *GetAccessTokenReply) GetErrcode() int64 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *GetAccessTokenReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

var File_api_wx_v1_wx_proto protoreflect.FileDescriptor

var file_api_wx_v1_wx_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x07,
	0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x11, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0xa3, 0x02, 0x0a, 0x14, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a,
	0x06, 0x74, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x4c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6d, 0x70, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29,
	0x0a, 0x08, 0x54, 0x6d, 0x70, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x22, 0x7f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x32, 0xd5, 0x02, 0x0a, 0x02, 0x57, 0x78, 0x12, 0x69,
	0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x73, 0x6e, 0x73, 0x2f, 0x6a, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x78, 0x0a, 0x0d, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x12, 0x1f, 0x2f, 0x63, 0x67, 0x69, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x12, 0x6a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x63, 0x67, 0x69, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x47, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2f, 0x77, 0x78, 0x61, 0x70, 0x70, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x77, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_wx_v1_wx_proto_rawDescOnce sync.Once
	file_api_wx_v1_wx_proto_rawDescData = file_api_wx_v1_wx_proto_rawDesc
)

func file_api_wx_v1_wx_proto_rawDescGZIP() []byte {
	file_api_wx_v1_wx_proto_rawDescOnce.Do(func() {
		file_api_wx_v1_wx_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_wx_v1_wx_proto_rawDescData)
	})
	return file_api_wx_v1_wx_proto_rawDescData
}

var file_api_wx_v1_wx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_wx_v1_wx_proto_goTypes = []interface{}{
	(*Code2SessionRequest)(nil),   // 0: api.wx.v1.Code2SessionRequest
	(*Code2SessionReply)(nil),     // 1: api.wx.v1.Code2SessionReply
	(*SubscribeSendRequest)(nil),  // 2: api.wx.v1.SubscribeSendRequest
	(*TmplData)(nil),              // 3: api.wx.v1.TmplData
	(*SubscribeSendReply)(nil),    // 4: api.wx.v1.SubscribeSendReply
	(*GetAccessTokenRequest)(nil), // 5: api.wx.v1.GetAccessTokenRequest
	(*GetAccessTokenReply)(nil),   // 6: api.wx.v1.GetAccessTokenReply
	nil,                           // 7: api.wx.v1.SubscribeSendRequest.DataEntry
}
var file_api_wx_v1_wx_proto_depIdxs = []int32{
	7, // 0: api.wx.v1.SubscribeSendRequest.data:type_name -> api.wx.v1.SubscribeSendRequest.DataEntry
	3, // 1: api.wx.v1.SubscribeSendRequest.DataEntry.value:type_name -> api.wx.v1.TmplData
	0, // 2: api.wx.v1.Wx.Code2Session:input_type -> api.wx.v1.Code2SessionRequest
	2, // 3: api.wx.v1.Wx.SubscribeSend:input_type -> api.wx.v1.SubscribeSendRequest
	5, // 4: api.wx.v1.Wx.GetAccessToken:input_type -> api.wx.v1.GetAccessTokenRequest
	1, // 5: api.wx.v1.Wx.Code2Session:output_type -> api.wx.v1.Code2SessionReply
	4, // 6: api.wx.v1.Wx.SubscribeSend:output_type -> api.wx.v1.SubscribeSendReply
	6, // 7: api.wx.v1.Wx.GetAccessToken:output_type -> api.wx.v1.GetAccessTokenReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_wx_v1_wx_proto_init() }
func file_api_wx_v1_wx_proto_init() {
	if File_api_wx_v1_wx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_wx_v1_wx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionRequest); i {
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
		file_api_wx_v1_wx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionReply); i {
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
		file_api_wx_v1_wx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeSendRequest); i {
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
		file_api_wx_v1_wx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TmplData); i {
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
		file_api_wx_v1_wx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeSendReply); i {
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
		file_api_wx_v1_wx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenRequest); i {
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
		file_api_wx_v1_wx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenReply); i {
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
			RawDescriptor: file_api_wx_v1_wx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_wx_v1_wx_proto_goTypes,
		DependencyIndexes: file_api_wx_v1_wx_proto_depIdxs,
		MessageInfos:      file_api_wx_v1_wx_proto_msgTypes,
	}.Build()
	File_api_wx_v1_wx_proto = out.File
	file_api_wx_v1_wx_proto_rawDesc = nil
	file_api_wx_v1_wx_proto_goTypes = nil
	file_api_wx_v1_wx_proto_depIdxs = nil
}
