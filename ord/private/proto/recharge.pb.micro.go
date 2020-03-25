// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: recharge.proto

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

// Client API for RechargeService service

type RechargeService interface {
	//创建充值单（管理员后台下单）
	Create(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error)
	//删除充值单
	Delete(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error)
	//获取充值单
	Get(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error)
	//查询充值单
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*RechargeResponse, error)
}

type rechargeService struct {
	c    client.Client
	name string
}

func NewRechargeService(name string, c client.Client) RechargeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.ord.private"
	}
	return &rechargeService{
		c:    c,
		name: name,
	}
}

func (c *rechargeService) Create(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeService.Create", in)
	out := new(RechargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeService) Delete(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeService.Delete", in)
	out := new(RechargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeService) Get(ctx context.Context, in *Recharge, opts ...client.CallOption) (*RechargeResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeService.Get", in)
	out := new(RechargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*RechargeResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeService.Search", in)
	out := new(RechargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RechargeService service

type RechargeServiceHandler interface {
	//创建充值单（管理员后台下单）
	Create(context.Context, *Recharge, *RechargeResponse) error
	//删除充值单
	Delete(context.Context, *Recharge, *RechargeResponse) error
	//获取充值单
	Get(context.Context, *Recharge, *RechargeResponse) error
	//查询充值单
	Search(context.Context, *BaseWhere, *RechargeResponse) error
}

func RegisterRechargeServiceHandler(s server.Server, hdlr RechargeServiceHandler, opts ...server.HandlerOption) error {
	type rechargeService interface {
		Create(ctx context.Context, in *Recharge, out *RechargeResponse) error
		Delete(ctx context.Context, in *Recharge, out *RechargeResponse) error
		Get(ctx context.Context, in *Recharge, out *RechargeResponse) error
		Search(ctx context.Context, in *BaseWhere, out *RechargeResponse) error
	}
	type RechargeService struct {
		rechargeService
	}
	h := &rechargeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RechargeService{h}, opts...))
}

type rechargeServiceHandler struct {
	RechargeServiceHandler
}

func (h *rechargeServiceHandler) Create(ctx context.Context, in *Recharge, out *RechargeResponse) error {
	return h.RechargeServiceHandler.Create(ctx, in, out)
}

func (h *rechargeServiceHandler) Delete(ctx context.Context, in *Recharge, out *RechargeResponse) error {
	return h.RechargeServiceHandler.Delete(ctx, in, out)
}

func (h *rechargeServiceHandler) Get(ctx context.Context, in *Recharge, out *RechargeResponse) error {
	return h.RechargeServiceHandler.Get(ctx, in, out)
}

func (h *rechargeServiceHandler) Search(ctx context.Context, in *BaseWhere, out *RechargeResponse) error {
	return h.RechargeServiceHandler.Search(ctx, in, out)
}