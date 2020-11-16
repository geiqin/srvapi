// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: fypCall.proto

package geiqin_srv_tms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FypCallService service

type FypCallService interface {
	Auth(ctx context.Context, in *Empty, opts ...client.CallOption) (*FypCallResponse, error)
	Submit(ctx context.Context, in *FypCall, opts ...client.CallOption) (*FypCallResponse, error)
}

type fypCallService struct {
	c    client.Client
	name string
}

func NewFypCallService(name string, c client.Client) FypCallService {
	return &fypCallService{
		c:    c,
		name: name,
	}
}

func (c *fypCallService) Auth(ctx context.Context, in *Empty, opts ...client.CallOption) (*FypCallResponse, error) {
	req := c.c.NewRequest(c.name, "FypCallService.Auth", in)
	out := new(FypCallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fypCallService) Submit(ctx context.Context, in *FypCall, opts ...client.CallOption) (*FypCallResponse, error) {
	req := c.c.NewRequest(c.name, "FypCallService.Submit", in)
	out := new(FypCallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FypCallService service

type FypCallServiceHandler interface {
	Auth(context.Context, *Empty, *FypCallResponse) error
	Submit(context.Context, *FypCall, *FypCallResponse) error
}

func RegisterFypCallServiceHandler(s server.Server, hdlr FypCallServiceHandler, opts ...server.HandlerOption) error {
	type fypCallService interface {
		Auth(ctx context.Context, in *Empty, out *FypCallResponse) error
		Submit(ctx context.Context, in *FypCall, out *FypCallResponse) error
	}
	type FypCallService struct {
		fypCallService
	}
	h := &fypCallServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FypCallService{h}, opts...))
}

type fypCallServiceHandler struct {
	FypCallServiceHandler
}

func (h *fypCallServiceHandler) Auth(ctx context.Context, in *Empty, out *FypCallResponse) error {
	return h.FypCallServiceHandler.Auth(ctx, in, out)
}

func (h *fypCallServiceHandler) Submit(ctx context.Context, in *FypCall, out *FypCallResponse) error {
	return h.FypCallServiceHandler.Submit(ctx, in, out)
}
