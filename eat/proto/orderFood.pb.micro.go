// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderFood.proto

package geiqin_srv_eat

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

// Client API for OrderFoodService service

type OrderFoodService interface {
	Create(ctx context.Context, in *OrderFood, opts ...client.CallOption) (*OrderFoodResponse, error)
	Get(ctx context.Context, in *OrderFoodWhere, opts ...client.CallOption) (*OrderFoodResponse, error)
}

type orderFoodService struct {
	c    client.Client
	name string
}

func NewOrderFoodService(name string, c client.Client) OrderFoodService {
	return &orderFoodService{
		c:    c,
		name: name,
	}
}

func (c *orderFoodService) Create(ctx context.Context, in *OrderFood, opts ...client.CallOption) (*OrderFoodResponse, error) {
	req := c.c.NewRequest(c.name, "OrderFoodService.Create", in)
	out := new(OrderFoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderFoodService) Get(ctx context.Context, in *OrderFoodWhere, opts ...client.CallOption) (*OrderFoodResponse, error) {
	req := c.c.NewRequest(c.name, "OrderFoodService.Get", in)
	out := new(OrderFoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderFoodService service

type OrderFoodServiceHandler interface {
	Create(context.Context, *OrderFood, *OrderFoodResponse) error
	Get(context.Context, *OrderFoodWhere, *OrderFoodResponse) error
}

func RegisterOrderFoodServiceHandler(s server.Server, hdlr OrderFoodServiceHandler, opts ...server.HandlerOption) error {
	type orderFoodService interface {
		Create(ctx context.Context, in *OrderFood, out *OrderFoodResponse) error
		Get(ctx context.Context, in *OrderFoodWhere, out *OrderFoodResponse) error
	}
	type OrderFoodService struct {
		orderFoodService
	}
	h := &orderFoodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderFoodService{h}, opts...))
}

type orderFoodServiceHandler struct {
	OrderFoodServiceHandler
}

func (h *orderFoodServiceHandler) Create(ctx context.Context, in *OrderFood, out *OrderFoodResponse) error {
	return h.OrderFoodServiceHandler.Create(ctx, in, out)
}

func (h *orderFoodServiceHandler) Get(ctx context.Context, in *OrderFoodWhere, out *OrderFoodResponse) error {
	return h.OrderFoodServiceHandler.Get(ctx, in, out)
}
