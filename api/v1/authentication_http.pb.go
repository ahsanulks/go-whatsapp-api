// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.12.4
// source: v1/authentication.proto

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

const OperationAuthenticationLogin = "/api.v1.Authentication/Login"

type AuthenticationHTTPServer interface {
	Login(context.Context, *LoginAuthenticationRequest) (*LoginAuthenticationResponse, error)
}

func RegisterAuthenticationHTTPServer(s *http.Server, srv AuthenticationHTTPServer) {
	r := s.Route("/")
	r.POST("/login", _Authentication_Login0_HTTP_Handler(srv))
}

func _Authentication_Login0_HTTP_Handler(srv AuthenticationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginAuthenticationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginAuthenticationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginAuthenticationResponse)
		return ctx.Result(200, reply)
	}
}

type AuthenticationHTTPClient interface {
	Login(ctx context.Context, req *LoginAuthenticationRequest, opts ...http.CallOption) (rsp *LoginAuthenticationResponse, err error)
}

type AuthenticationHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthenticationHTTPClient(client *http.Client) AuthenticationHTTPClient {
	return &AuthenticationHTTPClientImpl{client}
}

func (c *AuthenticationHTTPClientImpl) Login(ctx context.Context, in *LoginAuthenticationRequest, opts ...http.CallOption) (*LoginAuthenticationResponse, error) {
	var out LoginAuthenticationResponse
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthenticationLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
