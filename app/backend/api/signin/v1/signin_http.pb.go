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

type SigninHTTPServer interface {
	ListSignin(context.Context, *ListSigninRequest) (*ListSigninReply, error)
	Signin(context.Context, *SigninRequest) (*SigninReply, error)
}

func RegisterSigninHTTPServer(s *http.Server, srv SigninHTTPServer) {
	r := s.Route("/")
	r.POST("/signin", _Signin_Signin0_HTTP_Handler(srv))
	r.GET("/signin/list", _Signin_ListSignin0_HTTP_Handler(srv))
}

func _Signin_Signin0_HTTP_Handler(srv SigninHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SigninRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.v1.Signin/Signin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Signin(ctx, req.(*SigninRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SigninReply)
		return ctx.Result(200, reply)
	}
}

func _Signin_ListSignin0_HTTP_Handler(srv SigninHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSigninRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.v1.Signin/ListSignin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSignin(ctx, req.(*ListSigninRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSigninReply)
		return ctx.Result(200, reply)
	}
}

type SigninHTTPClient interface {
	ListSignin(ctx context.Context, req *ListSigninRequest, opts ...http.CallOption) (rsp *ListSigninReply, err error)
	Signin(ctx context.Context, req *SigninRequest, opts ...http.CallOption) (rsp *SigninReply, err error)
}

type SigninHTTPClientImpl struct {
	cc *http.Client
}

func NewSigninHTTPClient(client *http.Client) SigninHTTPClient {
	return &SigninHTTPClientImpl{client}
}

func (c *SigninHTTPClientImpl) ListSignin(ctx context.Context, in *ListSigninRequest, opts ...http.CallOption) (*ListSigninReply, error) {
	var out ListSigninReply
	pattern := "/signin/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.signin.v1.Signin/ListSignin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SigninHTTPClientImpl) Signin(ctx context.Context, in *SigninRequest, opts ...http.CallOption) (*SigninReply, error) {
	var out SigninReply
	pattern := "/signin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.signin.v1.Signin/Signin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
