// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: menu.proto

package geiqin_srv_uim

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

// Client API for MenuService service

type MenuService interface {
	Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*MenuResponse, error)
	List(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Tree(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
}

type menuService struct {
	c    client.Client
	name string
}

func NewMenuService(name string, c client.Client) MenuService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.uim"
	}
	return &menuService{
		c:    c,
		name: name,
	}
}

func (c *menuService) Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Get", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Search", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) List(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.List", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Tree(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Tree", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MenuService service

type MenuServiceHandler interface {
	Get(context.Context, *Menu, *MenuResponse) error
	Search(context.Context, *BaseWhere, *MenuResponse) error
	List(context.Context, *Menu, *MenuResponse) error
	Tree(context.Context, *Menu, *MenuResponse) error
}

func RegisterMenuServiceHandler(s server.Server, hdlr MenuServiceHandler, opts ...server.HandlerOption) error {
	type menuService interface {
		Get(ctx context.Context, in *Menu, out *MenuResponse) error
		Search(ctx context.Context, in *BaseWhere, out *MenuResponse) error
		List(ctx context.Context, in *Menu, out *MenuResponse) error
		Tree(ctx context.Context, in *Menu, out *MenuResponse) error
	}
	type MenuService struct {
		menuService
	}
	h := &menuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MenuService{h}, opts...))
}

type menuServiceHandler struct {
	MenuServiceHandler
}

func (h *menuServiceHandler) Get(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Get(ctx, in, out)
}

func (h *menuServiceHandler) Search(ctx context.Context, in *BaseWhere, out *MenuResponse) error {
	return h.MenuServiceHandler.Search(ctx, in, out)
}

func (h *menuServiceHandler) List(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.List(ctx, in, out)
}

func (h *menuServiceHandler) Tree(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Tree(ctx, in, out)
}
