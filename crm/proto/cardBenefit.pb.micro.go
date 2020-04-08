// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cardBenefit.proto

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

// Client API for CardBenefitService service

type CardBenefitService interface {
	Create(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error)
	Update(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error)
	Delete(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error)
	Get(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CardBenefitResponse, error)
}

type cardBenefitService struct {
	c    client.Client
	name string
}

func NewCardBenefitService(name string, c client.Client) CardBenefitService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.crm"
	}
	return &cardBenefitService{
		c:    c,
		name: name,
	}
}

func (c *cardBenefitService) Create(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error) {
	req := c.c.NewRequest(c.name, "CardBenefitService.Create", in)
	out := new(CardBenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardBenefitService) Update(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error) {
	req := c.c.NewRequest(c.name, "CardBenefitService.Update", in)
	out := new(CardBenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardBenefitService) Delete(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error) {
	req := c.c.NewRequest(c.name, "CardBenefitService.Delete", in)
	out := new(CardBenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardBenefitService) Get(ctx context.Context, in *CardBenefit, opts ...client.CallOption) (*CardBenefitResponse, error) {
	req := c.c.NewRequest(c.name, "CardBenefitService.Get", in)
	out := new(CardBenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardBenefitService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CardBenefitResponse, error) {
	req := c.c.NewRequest(c.name, "CardBenefitService.Search", in)
	out := new(CardBenefitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CardBenefitService service

type CardBenefitServiceHandler interface {
	Create(context.Context, *CardBenefit, *CardBenefitResponse) error
	Update(context.Context, *CardBenefit, *CardBenefitResponse) error
	Delete(context.Context, *CardBenefit, *CardBenefitResponse) error
	Get(context.Context, *CardBenefit, *CardBenefitResponse) error
	Search(context.Context, *BaseWhere, *CardBenefitResponse) error
}

func RegisterCardBenefitServiceHandler(s server.Server, hdlr CardBenefitServiceHandler, opts ...server.HandlerOption) error {
	type cardBenefitService interface {
		Create(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error
		Update(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error
		Delete(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error
		Get(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error
		Search(ctx context.Context, in *BaseWhere, out *CardBenefitResponse) error
	}
	type CardBenefitService struct {
		cardBenefitService
	}
	h := &cardBenefitServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CardBenefitService{h}, opts...))
}

type cardBenefitServiceHandler struct {
	CardBenefitServiceHandler
}

func (h *cardBenefitServiceHandler) Create(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error {
	return h.CardBenefitServiceHandler.Create(ctx, in, out)
}

func (h *cardBenefitServiceHandler) Update(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error {
	return h.CardBenefitServiceHandler.Update(ctx, in, out)
}

func (h *cardBenefitServiceHandler) Delete(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error {
	return h.CardBenefitServiceHandler.Delete(ctx, in, out)
}

func (h *cardBenefitServiceHandler) Get(ctx context.Context, in *CardBenefit, out *CardBenefitResponse) error {
	return h.CardBenefitServiceHandler.Get(ctx, in, out)
}

func (h *cardBenefitServiceHandler) Search(ctx context.Context, in *BaseWhere, out *CardBenefitResponse) error {
	return h.CardBenefitServiceHandler.Search(ctx, in, out)
}
