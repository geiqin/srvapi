// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dict.proto

package geiqin_srv_bas

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

// Client API for DictService service

type DictService interface {
	Get(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*DictResponse, error)
	Tree(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error)
	List(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error)
}

type dictService struct {
	c    client.Client
	name string
}

func NewDictService(name string, c client.Client) DictService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.bas"
	}
	return &dictService{
		c:    c,
		name: name,
	}
}

func (c *dictService) Get(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error) {
	req := c.c.NewRequest(c.name, "DictService.Get", in)
	out := new(DictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*DictResponse, error) {
	req := c.c.NewRequest(c.name, "DictService.Search", in)
	out := new(DictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) Tree(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error) {
	req := c.c.NewRequest(c.name, "DictService.Tree", in)
	out := new(DictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) List(ctx context.Context, in *Dict, opts ...client.CallOption) (*DictResponse, error) {
	req := c.c.NewRequest(c.name, "DictService.List", in)
	out := new(DictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DictService service

type DictServiceHandler interface {
	Get(context.Context, *Dict, *DictResponse) error
	Search(context.Context, *BaseWhere, *DictResponse) error
	Tree(context.Context, *Dict, *DictResponse) error
	List(context.Context, *Dict, *DictResponse) error
}

func RegisterDictServiceHandler(s server.Server, hdlr DictServiceHandler, opts ...server.HandlerOption) error {
	type dictService interface {
		Get(ctx context.Context, in *Dict, out *DictResponse) error
		Search(ctx context.Context, in *BaseWhere, out *DictResponse) error
		Tree(ctx context.Context, in *Dict, out *DictResponse) error
		List(ctx context.Context, in *Dict, out *DictResponse) error
	}
	type DictService struct {
		dictService
	}
	h := &dictServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DictService{h}, opts...))
}

type dictServiceHandler struct {
	DictServiceHandler
}

func (h *dictServiceHandler) Get(ctx context.Context, in *Dict, out *DictResponse) error {
	return h.DictServiceHandler.Get(ctx, in, out)
}

func (h *dictServiceHandler) Search(ctx context.Context, in *BaseWhere, out *DictResponse) error {
	return h.DictServiceHandler.Search(ctx, in, out)
}

func (h *dictServiceHandler) Tree(ctx context.Context, in *Dict, out *DictResponse) error {
	return h.DictServiceHandler.Tree(ctx, in, out)
}

func (h *dictServiceHandler) List(ctx context.Context, in *Dict, out *DictResponse) error {
	return h.DictServiceHandler.List(ctx, in, out)
}
