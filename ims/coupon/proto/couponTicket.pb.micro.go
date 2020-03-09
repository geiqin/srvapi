// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: couponTicket.proto

package geiqin_srv_ims_coupon

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

// Client API for CouponTicketService service

type CouponTicketService interface {
	//发放优惠劵凭证
	Give(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error)
	//核销优惠劵凭证
	Use(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error)
	//获取优惠劵凭证信息
	Get(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error)
	//查询优惠劵凭证
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
}

type couponTicketService struct {
	c    client.Client
	name string
}

func NewCouponTicketService(name string, c client.Client) CouponTicketService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.ims.coupon"
	}
	return &couponTicketService{
		c:    c,
		name: name,
	}
}

func (c *couponTicketService) Give(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Give", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) Use(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Use", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) Get(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Get", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Search", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CouponTicketService service

type CouponTicketServiceHandler interface {
	//发放优惠劵凭证
	Give(context.Context, *CouponTicket, *CouponTicketResponse) error
	//核销优惠劵凭证
	Use(context.Context, *CouponTicket, *CouponTicketResponse) error
	//获取优惠劵凭证信息
	Get(context.Context, *CouponTicket, *CouponTicketResponse) error
	//查询优惠劵凭证
	Search(context.Context, *BaseWhere, *CouponTicketResponse) error
}

func RegisterCouponTicketServiceHandler(s server.Server, hdlr CouponTicketServiceHandler, opts ...server.HandlerOption) error {
	type couponTicketService interface {
		Give(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error
		Use(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error
		Get(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error
		Search(ctx context.Context, in *BaseWhere, out *CouponTicketResponse) error
	}
	type CouponTicketService struct {
		couponTicketService
	}
	h := &couponTicketServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CouponTicketService{h}, opts...))
}

type couponTicketServiceHandler struct {
	CouponTicketServiceHandler
}

func (h *couponTicketServiceHandler) Give(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Give(ctx, in, out)
}

func (h *couponTicketServiceHandler) Use(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Use(ctx, in, out)
}

func (h *couponTicketServiceHandler) Get(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Get(ctx, in, out)
}

func (h *couponTicketServiceHandler) Search(ctx context.Context, in *BaseWhere, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Search(ctx, in, out)
}