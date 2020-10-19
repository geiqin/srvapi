// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: food.proto

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

// Client API for FoodService service

type FoodService interface {
	Create(ctx context.Context, in *Food, opts ...client.CallOption) (*FoodResponse, error)
	Update(ctx context.Context, in *Food, opts ...client.CallOption) (*FoodResponse, error)
	Get(ctx context.Context, in *FoodWhere, opts ...client.CallOption) (*FoodResponse, error)
	Get1(ctx context.Context, in *FoodWhere, opts ...client.CallOption) (*FoodResponse, error)
}

type foodService struct {
	c    client.Client
	name string
}

func NewFoodService(name string, c client.Client) FoodService {
	return &foodService{
		c:    c,
		name: name,
	}
}

func (c *foodService) Create(ctx context.Context, in *Food, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Create", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Update(ctx context.Context, in *Food, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Update", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Get(ctx context.Context, in *FoodWhere, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Get", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Get1(ctx context.Context, in *FoodWhere, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Get1", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FoodService service

type FoodServiceHandler interface {
	Create(context.Context, *Food, *FoodResponse) error
	Update(context.Context, *Food, *FoodResponse) error
	Get(context.Context, *FoodWhere, *FoodResponse) error
	Get1(context.Context, *FoodWhere, *FoodResponse) error
}

func RegisterFoodServiceHandler(s server.Server, hdlr FoodServiceHandler, opts ...server.HandlerOption) error {
	type foodService interface {
		Create(ctx context.Context, in *Food, out *FoodResponse) error
		Update(ctx context.Context, in *Food, out *FoodResponse) error
		Get(ctx context.Context, in *FoodWhere, out *FoodResponse) error
		Get1(ctx context.Context, in *FoodWhere, out *FoodResponse) error
	}
	type FoodService struct {
		foodService
	}
	h := &foodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FoodService{h}, opts...))
}

type foodServiceHandler struct {
	FoodServiceHandler
}

func (h *foodServiceHandler) Create(ctx context.Context, in *Food, out *FoodResponse) error {
	return h.FoodServiceHandler.Create(ctx, in, out)
}

func (h *foodServiceHandler) Update(ctx context.Context, in *Food, out *FoodResponse) error {
	return h.FoodServiceHandler.Update(ctx, in, out)
}

func (h *foodServiceHandler) Get(ctx context.Context, in *FoodWhere, out *FoodResponse) error {
	return h.FoodServiceHandler.Get(ctx, in, out)
}

func (h *foodServiceHandler) Get1(ctx context.Context, in *FoodWhere, out *FoodResponse) error {
	return h.FoodServiceHandler.Get1(ctx, in, out)
}