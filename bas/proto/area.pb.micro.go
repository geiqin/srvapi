// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: area.proto

package geiqin_srv_bas

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

// Client API for AreaService service

type AreaService interface {
	Get(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*AreaResponse, error)
	Tree(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error)
	List(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error)
}

type areaService struct {
	c    client.Client
	name string
}

func NewAreaService(name string, c client.Client) AreaService {
	return &areaService{
		c:    c,
		name: name,
	}
}

func (c *areaService) Get(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error) {
	req := c.c.NewRequest(c.name, "AreaService.Get", in)
	out := new(AreaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*AreaResponse, error) {
	req := c.c.NewRequest(c.name, "AreaService.Search", in)
	out := new(AreaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Tree(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error) {
	req := c.c.NewRequest(c.name, "AreaService.Tree", in)
	out := new(AreaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) List(ctx context.Context, in *Area, opts ...client.CallOption) (*AreaResponse, error) {
	req := c.c.NewRequest(c.name, "AreaService.List", in)
	out := new(AreaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AreaService service

type AreaServiceHandler interface {
	Get(context.Context, *Area, *AreaResponse) error
	Search(context.Context, *BaseWhere, *AreaResponse) error
	Tree(context.Context, *Area, *AreaResponse) error
	List(context.Context, *Area, *AreaResponse) error
}

func RegisterAreaServiceHandler(s server.Server, hdlr AreaServiceHandler, opts ...server.HandlerOption) error {
	type areaService interface {
		Get(ctx context.Context, in *Area, out *AreaResponse) error
		Search(ctx context.Context, in *BaseWhere, out *AreaResponse) error
		Tree(ctx context.Context, in *Area, out *AreaResponse) error
		List(ctx context.Context, in *Area, out *AreaResponse) error
	}
	type AreaService struct {
		areaService
	}
	h := &areaServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AreaService{h}, opts...))
}

type areaServiceHandler struct {
	AreaServiceHandler
}

func (h *areaServiceHandler) Get(ctx context.Context, in *Area, out *AreaResponse) error {
	return h.AreaServiceHandler.Get(ctx, in, out)
}

func (h *areaServiceHandler) Search(ctx context.Context, in *BaseWhere, out *AreaResponse) error {
	return h.AreaServiceHandler.Search(ctx, in, out)
}

func (h *areaServiceHandler) Tree(ctx context.Context, in *Area, out *AreaResponse) error {
	return h.AreaServiceHandler.Tree(ctx, in, out)
}

func (h *areaServiceHandler) List(ctx context.Context, in *Area, out *AreaResponse) error {
	return h.AreaServiceHandler.List(ctx, in, out)
}
