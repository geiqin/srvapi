// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: eatOutOrder.proto

package geiqin_srv_ord_private

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

// Client API for MyEatOutOrderService service

type MyEatOutOrderService interface {
	// 下单确认
	Confirm(ctx context.Context, in *BuyingRequest, opts ...client.CallOption) (*BuyingResponse, error)
	// 提交订单
	Submit(ctx context.Context, in *BuyingRequest, opts ...client.CallOption) (*OrderResponse, error)
}

type myEatOutOrderService struct {
	c    client.Client
	name string
}

func NewMyEatOutOrderService(name string, c client.Client) MyEatOutOrderService {
	return &myEatOutOrderService{
		c:    c,
		name: name,
	}
}

func (c *myEatOutOrderService) Confirm(ctx context.Context, in *BuyingRequest, opts ...client.CallOption) (*BuyingResponse, error) {
	req := c.c.NewRequest(c.name, "MyEatOutOrderService.Confirm", in)
	out := new(BuyingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myEatOutOrderService) Submit(ctx context.Context, in *BuyingRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "MyEatOutOrderService.Submit", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyEatOutOrderService service

type MyEatOutOrderServiceHandler interface {
	// 下单确认
	Confirm(context.Context, *BuyingRequest, *BuyingResponse) error
	// 提交订单
	Submit(context.Context, *BuyingRequest, *OrderResponse) error
}

func RegisterMyEatOutOrderServiceHandler(s server.Server, hdlr MyEatOutOrderServiceHandler, opts ...server.HandlerOption) error {
	type myEatOutOrderService interface {
		Confirm(ctx context.Context, in *BuyingRequest, out *BuyingResponse) error
		Submit(ctx context.Context, in *BuyingRequest, out *OrderResponse) error
	}
	type MyEatOutOrderService struct {
		myEatOutOrderService
	}
	h := &myEatOutOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyEatOutOrderService{h}, opts...))
}

type myEatOutOrderServiceHandler struct {
	MyEatOutOrderServiceHandler
}

func (h *myEatOutOrderServiceHandler) Confirm(ctx context.Context, in *BuyingRequest, out *BuyingResponse) error {
	return h.MyEatOutOrderServiceHandler.Confirm(ctx, in, out)
}

func (h *myEatOutOrderServiceHandler) Submit(ctx context.Context, in *BuyingRequest, out *OrderResponse) error {
	return h.MyEatOutOrderServiceHandler.Submit(ctx, in, out)
}
