// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WxHTTPServer interface {
	Code2Session(context.Context, *Code2SessionRequest) (*Code2SessionReply, error)
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenReply, error)
	SubscribeSend(context.Context, *SubscribeSendRequest) (*SubscribeSendReply, error)
}

func RegisterWxHTTPServer(s *http.Server, srv WxHTTPServer) {
	r := s.Route("/")
	r.GET("/sns/jscode2session", _Wx_Code2Session0_HTTP_Handler(srv))
	r.POST("/cgi-bin/message/subscribe/send", _Wx_SubscribeSend0_HTTP_Handler(srv))
	r.GET("/cgi-bin/token", _Wx_GetAccessToken0_HTTP_Handler(srv))
}

func _Wx_Code2Session0_HTTP_Handler(srv WxHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Code2SessionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wx.v1.Wx/Code2Session")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Code2Session(ctx, req.(*Code2SessionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Code2SessionReply)
		return ctx.Result(200, reply)
	}
}

func _Wx_SubscribeSend0_HTTP_Handler(srv WxHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SubscribeSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wx.v1.Wx/SubscribeSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubscribeSend(ctx, req.(*SubscribeSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SubscribeSendReply)
		return ctx.Result(200, reply)
	}
}

func _Wx_GetAccessToken0_HTTP_Handler(srv WxHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAccessTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wx.v1.Wx/GetAccessToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAccessToken(ctx, req.(*GetAccessTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAccessTokenReply)
		return ctx.Result(200, reply)
	}
}

type WxHTTPClient interface {
	Code2Session(ctx context.Context, req *Code2SessionRequest, opts ...http.CallOption) (rsp *Code2SessionReply, err error)
	GetAccessToken(ctx context.Context, req *GetAccessTokenRequest, opts ...http.CallOption) (rsp *GetAccessTokenReply, err error)
	SubscribeSend(ctx context.Context, req *SubscribeSendRequest, opts ...http.CallOption) (rsp *SubscribeSendReply, err error)
}

type WxHTTPClientImpl struct {
	cc *http.Client
}

func NewWxHTTPClient(client *http.Client) WxHTTPClient {
	return &WxHTTPClientImpl{client}
}

func (c *WxHTTPClientImpl) Code2Session(ctx context.Context, in *Code2SessionRequest, opts ...http.CallOption) (*Code2SessionReply, error) {
	var out Code2SessionReply
	pattern := "/sns/jscode2session"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.wx.v1.Wx/Code2Session"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WxHTTPClientImpl) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...http.CallOption) (*GetAccessTokenReply, error) {
	var out GetAccessTokenReply
	pattern := "/cgi-bin/token"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.wx.v1.Wx/GetAccessToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WxHTTPClientImpl) SubscribeSend(ctx context.Context, in *SubscribeSendRequest, opts ...http.CallOption) (*SubscribeSendReply, error) {
	var out SubscribeSendReply
	pattern := "/cgi-bin/message/subscribe/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.wx.v1.Wx/SubscribeSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
