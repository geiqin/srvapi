// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: specValue.proto

package geiqin_srv_pdm

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

// Client API for SpecValueService service

type SpecValueService interface {
	Create(ctx context.Context, in *SpecValue, opts ...client.CallOption) (*SpecValueResponse, error)
	Delete(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error)
	Get(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error)
	List(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error)
}

type specValueService struct {
	c    client.Client
	name string
}

func NewSpecValueService(name string, c client.Client) SpecValueService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.pdm"
	}
	return &specValueService{
		c:    c,
		name: name,
	}
}

func (c *specValueService) Create(ctx context.Context, in *SpecValue, opts ...client.CallOption) (*SpecValueResponse, error) {
	req := c.c.NewRequest(c.name, "SpecValueService.Create", in)
	out := new(SpecValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specValueService) Delete(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error) {
	req := c.c.NewRequest(c.name, "SpecValueService.Delete", in)
	out := new(SpecValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specValueService) Get(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error) {
	req := c.c.NewRequest(c.name, "SpecValueService.Get", in)
	out := new(SpecValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specValueService) List(ctx context.Context, in *IdInt, opts ...client.CallOption) (*SpecValueResponse, error) {
	req := c.c.NewRequest(c.name, "SpecValueService.List", in)
	out := new(SpecValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpecValueService service

type SpecValueServiceHandler interface {
	Create(context.Context, *SpecValue, *SpecValueResponse) error
	Delete(context.Context, *IdInt, *SpecValueResponse) error
	Get(context.Context, *IdInt, *SpecValueResponse) error
	List(context.Context, *IdInt, *SpecValueResponse) error
}

func RegisterSpecValueServiceHandler(s server.Server, hdlr SpecValueServiceHandler, opts ...server.HandlerOption) error {
	type specValueService interface {
		Create(ctx context.Context, in *SpecValue, out *SpecValueResponse) error
		Delete(ctx context.Context, in *IdInt, out *SpecValueResponse) error
		Get(ctx context.Context, in *IdInt, out *SpecValueResponse) error
		List(ctx context.Context, in *IdInt, out *SpecValueResponse) error
	}
	type SpecValueService struct {
		specValueService
	}
	h := &specValueServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SpecValueService{h}, opts...))
}

type specValueServiceHandler struct {
	SpecValueServiceHandler
}

func (h *specValueServiceHandler) Create(ctx context.Context, in *SpecValue, out *SpecValueResponse) error {
	return h.SpecValueServiceHandler.Create(ctx, in, out)
}

func (h *specValueServiceHandler) Delete(ctx context.Context, in *IdInt, out *SpecValueResponse) error {
	return h.SpecValueServiceHandler.Delete(ctx, in, out)
}

func (h *specValueServiceHandler) Get(ctx context.Context, in *IdInt, out *SpecValueResponse) error {
	return h.SpecValueServiceHandler.Get(ctx, in, out)
}

func (h *specValueServiceHandler) List(ctx context.Context, in *IdInt, out *SpecValueResponse) error {
	return h.SpecValueServiceHandler.List(ctx, in, out)
}