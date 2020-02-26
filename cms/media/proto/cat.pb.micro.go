// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cat.proto

package geiqin_srv_cms_media

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

// Client API for CatService service

type CatService interface {
	Create(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error)
	Update(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error)
	Delete(ctx context.Context, in *IdInt, opts ...client.CallOption) (*CatResponse, error)
	Get(ctx context.Context, in *IdInt, opts ...client.CallOption) (*CatResponse, error)
	Search(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error)
	List(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error)
	Tree(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error)
}

type catService struct {
	c    client.Client
	name string
}

func NewCatService(name string, c client.Client) CatService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.cms.media"
	}
	return &catService{
		c:    c,
		name: name,
	}
}

func (c *catService) Create(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Create", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) Update(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Update", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) Delete(ctx context.Context, in *IdInt, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Delete", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) Get(ctx context.Context, in *IdInt, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Get", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) Search(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Search", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) List(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.List", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catService) Tree(ctx context.Context, in *Cat, opts ...client.CallOption) (*CatResponse, error) {
	req := c.c.NewRequest(c.name, "CatService.Tree", in)
	out := new(CatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CatService service

type CatServiceHandler interface {
	Create(context.Context, *Cat, *CatResponse) error
	Update(context.Context, *Cat, *CatResponse) error
	Delete(context.Context, *IdInt, *CatResponse) error
	Get(context.Context, *IdInt, *CatResponse) error
	Search(context.Context, *Cat, *CatResponse) error
	List(context.Context, *Cat, *CatResponse) error
	Tree(context.Context, *Cat, *CatResponse) error
}

func RegisterCatServiceHandler(s server.Server, hdlr CatServiceHandler, opts ...server.HandlerOption) error {
	type catService interface {
		Create(ctx context.Context, in *Cat, out *CatResponse) error
		Update(ctx context.Context, in *Cat, out *CatResponse) error
		Delete(ctx context.Context, in *IdInt, out *CatResponse) error
		Get(ctx context.Context, in *IdInt, out *CatResponse) error
		Search(ctx context.Context, in *Cat, out *CatResponse) error
		List(ctx context.Context, in *Cat, out *CatResponse) error
		Tree(ctx context.Context, in *Cat, out *CatResponse) error
	}
	type CatService struct {
		catService
	}
	h := &catServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CatService{h}, opts...))
}

type catServiceHandler struct {
	CatServiceHandler
}

func (h *catServiceHandler) Create(ctx context.Context, in *Cat, out *CatResponse) error {
	return h.CatServiceHandler.Create(ctx, in, out)
}

func (h *catServiceHandler) Update(ctx context.Context, in *Cat, out *CatResponse) error {
	return h.CatServiceHandler.Update(ctx, in, out)
}

func (h *catServiceHandler) Delete(ctx context.Context, in *IdInt, out *CatResponse) error {
	return h.CatServiceHandler.Delete(ctx, in, out)
}

func (h *catServiceHandler) Get(ctx context.Context, in *IdInt, out *CatResponse) error {
	return h.CatServiceHandler.Get(ctx, in, out)
}

func (h *catServiceHandler) Search(ctx context.Context, in *Cat, out *CatResponse) error {
	return h.CatServiceHandler.Search(ctx, in, out)
}

func (h *catServiceHandler) List(ctx context.Context, in *Cat, out *CatResponse) error {
	return h.CatServiceHandler.List(ctx, in, out)
}

func (h *catServiceHandler) Tree(ctx context.Context, in *Cat, out *CatResponse) error {
	return h.CatServiceHandler.Tree(ctx, in, out)
}
