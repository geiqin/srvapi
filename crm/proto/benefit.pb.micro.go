// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: benefit.proto

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

// Client API for BenefitService service

type BenefitService interface {
	Create(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error)
	Update(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error)
	Delete(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error)
	Get(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*BenefitResponse, error)
}

type benefitService struct {
	c    client.Client
	name string
}

func NewBenefitService(name string, c client.Client) BenefitService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.crm"
	}
	return &benefitService{
		c:    c,
		name: name,
	}
}

func (c *benefitService) Create(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error) {
	req := c.c.NewRequest(c.name, "BenefitService.Create", in)
	out := new(BenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benefitService) Update(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error) {
	req := c.c.NewRequest(c.name, "BenefitService.Update", in)
	out := new(BenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benefitService) Delete(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error) {
	req := c.c.NewRequest(c.name, "BenefitService.Delete", in)
	out := new(BenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benefitService) Get(ctx context.Context, in *Benefit, opts ...client.CallOption) (*BenefitResponse, error) {
	req := c.c.NewRequest(c.name, "BenefitService.Get", in)
	out := new(BenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benefitService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*BenefitResponse, error) {
	req := c.c.NewRequest(c.name, "BenefitService.Search", in)
	out := new(BenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BenefitService service

type BenefitServiceHandler interface {
	Create(context.Context, *Benefit, *BenefitResponse) error
	Update(context.Context, *Benefit, *BenefitResponse) error
	Delete(context.Context, *Benefit, *BenefitResponse) error
	Get(context.Context, *Benefit, *BenefitResponse) error
	Search(context.Context, *BaseWhere, *BenefitResponse) error
}

func RegisterBenefitServiceHandler(s server.Server, hdlr BenefitServiceHandler, opts ...server.HandlerOption) error {
	type benefitService interface {
		Create(ctx context.Context, in *Benefit, out *BenefitResponse) error
		Update(ctx context.Context, in *Benefit, out *BenefitResponse) error
		Delete(ctx context.Context, in *Benefit, out *BenefitResponse) error
		Get(ctx context.Context, in *Benefit, out *BenefitResponse) error
		Search(ctx context.Context, in *BaseWhere, out *BenefitResponse) error
	}
	type BenefitService struct {
		benefitService
	}
	h := &benefitServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BenefitService{h}, opts...))
}

type benefitServiceHandler struct {
	BenefitServiceHandler
}

func (h *benefitServiceHandler) Create(ctx context.Context, in *Benefit, out *BenefitResponse) error {
	return h.BenefitServiceHandler.Create(ctx, in, out)
}

func (h *benefitServiceHandler) Update(ctx context.Context, in *Benefit, out *BenefitResponse) error {
	return h.BenefitServiceHandler.Update(ctx, in, out)
}

func (h *benefitServiceHandler) Delete(ctx context.Context, in *Benefit, out *BenefitResponse) error {
	return h.BenefitServiceHandler.Delete(ctx, in, out)
}

func (h *benefitServiceHandler) Get(ctx context.Context, in *Benefit, out *BenefitResponse) error {
	return h.BenefitServiceHandler.Get(ctx, in, out)
}

func (h *benefitServiceHandler) Search(ctx context.Context, in *BaseWhere, out *BenefitResponse) error {
	return h.BenefitServiceHandler.Search(ctx, in, out)
}
