// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: itemPrice.proto

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

// Client API for ItemPriceService service

type ItemPriceService interface {
	Set(ctx context.Context, in *ItemPrice, opts ...client.CallOption) (*ItemPriceResponse, error)
	Get(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemPriceResponse, error)
}

type itemPriceService struct {
	c    client.Client
	name string
}

func NewItemPriceService(name string, c client.Client) ItemPriceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.pdm"
	}
	return &itemPriceService{
		c:    c,
		name: name,
	}
}

func (c *itemPriceService) Set(ctx context.Context, in *ItemPrice, opts ...client.CallOption) (*ItemPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ItemPriceService.Set", in)
	out := new(ItemPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemPriceService) Get(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ItemPriceService.Get", in)
	out := new(ItemPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemPriceService service

type ItemPriceServiceHandler interface {
	Set(context.Context, *ItemPrice, *ItemPriceResponse) error
	Get(context.Context, *Id, *ItemPriceResponse) error
}

func RegisterItemPriceServiceHandler(s server.Server, hdlr ItemPriceServiceHandler, opts ...server.HandlerOption) error {
	type itemPriceService interface {
		Set(ctx context.Context, in *ItemPrice, out *ItemPriceResponse) error
		Get(ctx context.Context, in *Id, out *ItemPriceResponse) error
	}
	type ItemPriceService struct {
		itemPriceService
	}
	h := &itemPriceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ItemPriceService{h}, opts...))
}

type itemPriceServiceHandler struct {
	ItemPriceServiceHandler
}

func (h *itemPriceServiceHandler) Set(ctx context.Context, in *ItemPrice, out *ItemPriceResponse) error {
	return h.ItemPriceServiceHandler.Set(ctx, in, out)
}

func (h *itemPriceServiceHandler) Get(ctx context.Context, in *Id, out *ItemPriceResponse) error {
	return h.ItemPriceServiceHandler.Get(ctx, in, out)
}