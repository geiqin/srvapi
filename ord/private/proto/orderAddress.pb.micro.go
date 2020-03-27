// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderAddress.proto

package geiqin_srv_ord_private

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

// Client API for OrderAddressService service

type OrderAddressService interface {
	Create(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error)
	Update(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error)
	Delete(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error)
	Get(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*OrderAddressResponse, error)
}

type orderAddressService struct {
	c    client.Client
	name string
}

func NewOrderAddressService(name string, c client.Client) OrderAddressService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.ord.private"
	}
	return &orderAddressService{
		c:    c,
		name: name,
	}
}

func (c *orderAddressService) Create(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAddressService.Create", in)
	out := new(OrderAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAddressService) Update(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAddressService.Update", in)
	out := new(OrderAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAddressService) Delete(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAddressService.Delete", in)
	out := new(OrderAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAddressService) Get(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderAddressResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAddressService.Get", in)
	out := new(OrderAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAddressService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*OrderAddressResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAddressService.Search", in)
	out := new(OrderAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderAddressService service

type OrderAddressServiceHandler interface {
	Create(context.Context, *OrderAddress, *OrderAddressResponse) error
	Update(context.Context, *OrderAddress, *OrderAddressResponse) error
	Delete(context.Context, *OrderAddress, *OrderAddressResponse) error
	Get(context.Context, *OrderAddress, *OrderAddressResponse) error
	Search(context.Context, *BaseWhere, *OrderAddressResponse) error
}

func RegisterOrderAddressServiceHandler(s server.Server, hdlr OrderAddressServiceHandler, opts ...server.HandlerOption) error {
	type orderAddressService interface {
		Create(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error
		Update(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error
		Delete(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error
		Get(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error
		Search(ctx context.Context, in *BaseWhere, out *OrderAddressResponse) error
	}
	type OrderAddressService struct {
		orderAddressService
	}
	h := &orderAddressServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderAddressService{h}, opts...))
}

type orderAddressServiceHandler struct {
	OrderAddressServiceHandler
}

func (h *orderAddressServiceHandler) Create(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error {
	return h.OrderAddressServiceHandler.Create(ctx, in, out)
}

func (h *orderAddressServiceHandler) Update(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error {
	return h.OrderAddressServiceHandler.Update(ctx, in, out)
}

func (h *orderAddressServiceHandler) Delete(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error {
	return h.OrderAddressServiceHandler.Delete(ctx, in, out)
}

func (h *orderAddressServiceHandler) Get(ctx context.Context, in *OrderAddress, out *OrderAddressResponse) error {
	return h.OrderAddressServiceHandler.Get(ctx, in, out)
}

func (h *orderAddressServiceHandler) Search(ctx context.Context, in *BaseWhere, out *OrderAddressResponse) error {
	return h.OrderAddressServiceHandler.Search(ctx, in, out)
}
