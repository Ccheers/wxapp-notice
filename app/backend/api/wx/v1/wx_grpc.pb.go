// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WxClient is the client API for Wx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WxClient interface {
	Code2Session(ctx context.Context, in *Code2SessionRequest, opts ...grpc.CallOption) (*Code2SessionReply, error)
	SubscribeSend(ctx context.Context, in *SubscribeSendRequest, opts ...grpc.CallOption) (*SubscribeSendReply, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenReply, error)
}

type wxClient struct {
	cc grpc.ClientConnInterface
}

func NewWxClient(cc grpc.ClientConnInterface) WxClient {
	return &wxClient{cc}
}

func (c *wxClient) Code2Session(ctx context.Context, in *Code2SessionRequest, opts ...grpc.CallOption) (*Code2SessionReply, error) {
	out := new(Code2SessionReply)
	err := c.cc.Invoke(ctx, "/api.wx.v1.Wx/Code2Session", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxClient) SubscribeSend(ctx context.Context, in *SubscribeSendRequest, opts ...grpc.CallOption) (*SubscribeSendReply, error) {
	out := new(SubscribeSendReply)
	err := c.cc.Invoke(ctx, "/api.wx.v1.Wx/SubscribeSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenReply, error) {
	out := new(GetAccessTokenReply)
	err := c.cc.Invoke(ctx, "/api.wx.v1.Wx/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WxServer is the server API for Wx service.
// All implementations must embed UnimplementedWxServer
// for forward compatibility
type WxServer interface {
	Code2Session(context.Context, *Code2SessionRequest) (*Code2SessionReply, error)
	SubscribeSend(context.Context, *SubscribeSendRequest) (*SubscribeSendReply, error)
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenReply, error)
	mustEmbedUnimplementedWxServer()
}

// UnimplementedWxServer must be embedded to have forward compatible implementations.
type UnimplementedWxServer struct {
}

func (UnimplementedWxServer) Code2Session(context.Context, *Code2SessionRequest) (*Code2SessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code2Session not implemented")
}
func (UnimplementedWxServer) SubscribeSend(context.Context, *SubscribeSendRequest) (*SubscribeSendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeSend not implemented")
}
func (UnimplementedWxServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedWxServer) mustEmbedUnimplementedWxServer() {}

// UnsafeWxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WxServer will
// result in compilation errors.
type UnsafeWxServer interface {
	mustEmbedUnimplementedWxServer()
}

func RegisterWxServer(s grpc.ServiceRegistrar, srv WxServer) {
	s.RegisterService(&Wx_ServiceDesc, srv)
}

func _Wx_Code2Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code2SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WxServer).Code2Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wx.v1.Wx/Code2Session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WxServer).Code2Session(ctx, req.(*Code2SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wx_SubscribeSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WxServer).SubscribeSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wx.v1.Wx/SubscribeSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WxServer).SubscribeSend(ctx, req.(*SubscribeSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wx_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WxServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wx.v1.Wx/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WxServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wx_ServiceDesc is the grpc.ServiceDesc for Wx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wx.v1.Wx",
	HandlerType: (*WxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Code2Session",
			Handler:    _Wx_Code2Session_Handler,
		},
		{
			MethodName: "SubscribeSend",
			Handler:    _Wx_SubscribeSend_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _Wx_GetAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/wx/v1/wx.proto",
}
