// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: fanWeixin.proto

package geiqin_srv_crm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for FanWeixinService service

type FanWeixinService interface {
	Create(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error)
	Update(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error)
	Delete(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error)
	Get(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*FanWeixinResponse, error)
}

type fanWeixinService struct {
	c    client.Client
	name string
}

func NewFanWeixinService(name string, c client.Client) FanWeixinService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.crm"
	}
	return &fanWeixinService{
		c:    c,
		name: name,
	}
}

func (c *fanWeixinService) Create(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error) {
	req := c.c.NewRequest(c.name, "FanWeixinService.Create", in)
	out := new(FanWeixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanWeixinService) Update(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error) {
	req := c.c.NewRequest(c.name, "FanWeixinService.Update", in)
	out := new(FanWeixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanWeixinService) Delete(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error) {
	req := c.c.NewRequest(c.name, "FanWeixinService.Delete", in)
	out := new(FanWeixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanWeixinService) Get(ctx context.Context, in *FanWeixin, opts ...client.CallOption) (*FanWeixinResponse, error) {
	req := c.c.NewRequest(c.name, "FanWeixinService.Get", in)
	out := new(FanWeixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanWeixinService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*FanWeixinResponse, error) {
	req := c.c.NewRequest(c.name, "FanWeixinService.Search", in)
	out := new(FanWeixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FanWeixinService service

type FanWeixinServiceHandler interface {
	Create(context.Context, *FanWeixin, *FanWeixinResponse) error
	Update(context.Context, *FanWeixin, *FanWeixinResponse) error
	Delete(context.Context, *FanWeixin, *FanWeixinResponse) error
	Get(context.Context, *FanWeixin, *FanWeixinResponse) error
	Search(context.Context, *BaseWhere, *FanWeixinResponse) error
}

func RegisterFanWeixinServiceHandler(s server.Server, hdlr FanWeixinServiceHandler, opts ...server.HandlerOption) error {
	type fanWeixinService interface {
		Create(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error
		Update(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error
		Delete(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error
		Get(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error
		Search(ctx context.Context, in *BaseWhere, out *FanWeixinResponse) error
	}
	type FanWeixinService struct {
		fanWeixinService
	}
	h := &fanWeixinServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FanWeixinService{h}, opts...))
}

type fanWeixinServiceHandler struct {
	FanWeixinServiceHandler
}

func (h *fanWeixinServiceHandler) Create(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error {
	return h.FanWeixinServiceHandler.Create(ctx, in, out)
}

func (h *fanWeixinServiceHandler) Update(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error {
	return h.FanWeixinServiceHandler.Update(ctx, in, out)
}

func (h *fanWeixinServiceHandler) Delete(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error {
	return h.FanWeixinServiceHandler.Delete(ctx, in, out)
}

func (h *fanWeixinServiceHandler) Get(ctx context.Context, in *FanWeixin, out *FanWeixinResponse) error {
	return h.FanWeixinServiceHandler.Get(ctx, in, out)
}

func (h *fanWeixinServiceHandler) Search(ctx context.Context, in *BaseWhere, out *FanWeixinResponse) error {
	return h.FanWeixinServiceHandler.Search(ctx, in, out)
}
