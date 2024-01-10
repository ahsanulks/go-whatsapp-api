// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.12.4
// source: v1/message.proto

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

const OperationMessageSendMessage = "/api.v1.Message/SendMessage"

type MessageHTTPServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}

func RegisterMessageHTTPServer(s *http.Server, srv MessageHTTPServer) {
	r := s.Route("/")
	r.POST("/message", _Message_SendMessage0_HTTP_Handler(srv))
}

func _Message_SendMessage0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageSendMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMessage(ctx, req.(*SendMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageResponse)
		return ctx.Result(200, reply)
	}
}

type MessageHTTPClient interface {
	SendMessage(ctx context.Context, req *SendMessageRequest, opts ...http.CallOption) (rsp *SendMessageResponse, err error)
}

type MessageHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageHTTPClient(client *http.Client) MessageHTTPClient {
	return &MessageHTTPClientImpl{client}
}

func (c *MessageHTTPClientImpl) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...http.CallOption) (*SendMessageResponse, error) {
	var out SendMessageResponse
	pattern := "/message"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageSendMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}