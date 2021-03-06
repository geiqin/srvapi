// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: present.proto

package geiqin_srv_ims_present

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

// Client API for PresentService service

type PresentService interface {
	Create(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error)
	Update(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error)
	Delete(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error)
	Get(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error)
	List(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*PresentResponse, error)
}

type presentService struct {
	c    client.Client
	name string
}

func NewPresentService(name string, c client.Client) PresentService {
	return &presentService{
		c:    c,
		name: name,
	}
}

func (c *presentService) Create(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.Create", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentService) Update(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.Update", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentService) Delete(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.Delete", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentService) Get(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.Get", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentService) List(ctx context.Context, in *Present, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.List", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*PresentResponse, error) {
	req := c.c.NewRequest(c.name, "PresentService.Search", in)
	out := new(PresentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PresentService service

type PresentServiceHandler interface {
	Create(context.Context, *Present, *PresentResponse) error
	Update(context.Context, *Present, *PresentResponse) error
	Delete(context.Context, *Present, *PresentResponse) error
	Get(context.Context, *Present, *PresentResponse) error
	List(context.Context, *Present, *PresentResponse) error
	Search(context.Context, *BaseWhere, *PresentResponse) error
}

func RegisterPresentServiceHandler(s server.Server, hdlr PresentServiceHandler, opts ...server.HandlerOption) error {
	type presentService interface {
		Create(ctx context.Context, in *Present, out *PresentResponse) error
		Update(ctx context.Context, in *Present, out *PresentResponse) error
		Delete(ctx context.Context, in *Present, out *PresentResponse) error
		Get(ctx context.Context, in *Present, out *PresentResponse) error
		List(ctx context.Context, in *Present, out *PresentResponse) error
		Search(ctx context.Context, in *BaseWhere, out *PresentResponse) error
	}
	type PresentService struct {
		presentService
	}
	h := &presentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PresentService{h}, opts...))
}

type presentServiceHandler struct {
	PresentServiceHandler
}

func (h *presentServiceHandler) Create(ctx context.Context, in *Present, out *PresentResponse) error {
	return h.PresentServiceHandler.Create(ctx, in, out)
}

func (h *presentServiceHandler) Update(ctx context.Context, in *Present, out *PresentResponse) error {
	return h.PresentServiceHandler.Update(ctx, in, out)
}

func (h *presentServiceHandler) Delete(ctx context.Context, in *Present, out *PresentResponse) error {
	return h.PresentServiceHandler.Delete(ctx, in, out)
}

func (h *presentServiceHandler) Get(ctx context.Context, in *Present, out *PresentResponse) error {
	return h.PresentServiceHandler.Get(ctx, in, out)
}

func (h *presentServiceHandler) List(ctx context.Context, in *Present, out *PresentResponse) error {
	return h.PresentServiceHandler.List(ctx, in, out)
}

func (h *presentServiceHandler) Search(ctx context.Context, in *BaseWhere, out *PresentResponse) error {
	return h.PresentServiceHandler.Search(ctx, in, out)
}
