// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderInfo.proto

package geiqin_srv_dms

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

// Client API for MyOrderService service

type MyOrderService interface {
	// 获取推广订单详情
	Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
	// 查询推广订单列表
	Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
}

type myOrderService struct {
	c    client.Client
	name string
}

func NewMyOrderService(name string, c client.Client) MyOrderService {
	return &myOrderService{
		c:    c,
		name: name,
	}
}

func (c *myOrderService) Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyOrderService.Get", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myOrderService) Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyOrderService.Search", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyOrderService service

type MyOrderServiceHandler interface {
	// 获取推广订单详情
	Get(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
	// 查询推广订单列表
	Search(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
}

func RegisterMyOrderServiceHandler(s server.Server, hdlr MyOrderServiceHandler, opts ...server.HandlerOption) error {
	type myOrderService interface {
		Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
		Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
	}
	type MyOrderService struct {
		myOrderService
	}
	h := &myOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyOrderService{h}, opts...))
}

type myOrderServiceHandler struct {
	MyOrderServiceHandler
}

func (h *myOrderServiceHandler) Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.MyOrderServiceHandler.Get(ctx, in, out)
}

func (h *myOrderServiceHandler) Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.MyOrderServiceHandler.Search(ctx, in, out)
}

// Client API for OrderService service

type OrderService interface {
	// 获取推广订单详情
	Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
	// 查询推广订单列表
	Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Get", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Search", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	// 获取推广订单详情
	Get(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
	// 查询推广订单列表
	Search(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
		Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.OrderServiceHandler.Get(ctx, in, out)
}

func (h *orderServiceHandler) Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.OrderServiceHandler.Search(ctx, in, out)
}

// Client API for BonusOrderService service

type BonusOrderService interface {
	// 获取分红订单详情
	Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
	// 查询分红订单列表
	Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
}

type bonusOrderService struct {
	c    client.Client
	name string
}

func NewBonusOrderService(name string, c client.Client) BonusOrderService {
	return &bonusOrderService{
		c:    c,
		name: name,
	}
}

func (c *bonusOrderService) Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "BonusOrderService.Get", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bonusOrderService) Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "BonusOrderService.Search", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BonusOrderService service

type BonusOrderServiceHandler interface {
	// 获取分红订单详情
	Get(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
	// 查询分红订单列表
	Search(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
}

func RegisterBonusOrderServiceHandler(s server.Server, hdlr BonusOrderServiceHandler, opts ...server.HandlerOption) error {
	type bonusOrderService interface {
		Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
		Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
	}
	type BonusOrderService struct {
		bonusOrderService
	}
	h := &bonusOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BonusOrderService{h}, opts...))
}

type bonusOrderServiceHandler struct {
	BonusOrderServiceHandler
}

func (h *bonusOrderServiceHandler) Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.BonusOrderServiceHandler.Get(ctx, in, out)
}

func (h *bonusOrderServiceHandler) Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.BonusOrderServiceHandler.Search(ctx, in, out)
}

// Client API for MyBonusOrderService service

type MyBonusOrderService interface {
	// 获取分红订单详情
	Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
	// 查询分红订单列表
	Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error)
}

type myBonusOrderService struct {
	c    client.Client
	name string
}

func NewMyBonusOrderService(name string, c client.Client) MyBonusOrderService {
	return &myBonusOrderService{
		c:    c,
		name: name,
	}
}

func (c *myBonusOrderService) Get(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyBonusOrderService.Get", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myBonusOrderService) Search(ctx context.Context, in *OrderInfoWhere, opts ...client.CallOption) (*OrderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyBonusOrderService.Search", in)
	out := new(OrderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyBonusOrderService service

type MyBonusOrderServiceHandler interface {
	// 获取分红订单详情
	Get(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
	// 查询分红订单列表
	Search(context.Context, *OrderInfoWhere, *OrderInfoResponse) error
}

func RegisterMyBonusOrderServiceHandler(s server.Server, hdlr MyBonusOrderServiceHandler, opts ...server.HandlerOption) error {
	type myBonusOrderService interface {
		Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
		Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error
	}
	type MyBonusOrderService struct {
		myBonusOrderService
	}
	h := &myBonusOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyBonusOrderService{h}, opts...))
}

type myBonusOrderServiceHandler struct {
	MyBonusOrderServiceHandler
}

func (h *myBonusOrderServiceHandler) Get(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.MyBonusOrderServiceHandler.Get(ctx, in, out)
}

func (h *myBonusOrderServiceHandler) Search(ctx context.Context, in *OrderInfoWhere, out *OrderInfoResponse) error {
	return h.MyBonusOrderServiceHandler.Search(ctx, in, out)
}
