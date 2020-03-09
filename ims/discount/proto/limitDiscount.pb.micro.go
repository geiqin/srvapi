// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: limitDiscount.proto

package geiqin_srv_ims_discount

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

// Client API for LimitDiscountService service

type LimitDiscountService interface {
	//创建限时折扣活动
	Create(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error)
	//编辑限时折扣活动
	Update(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error)
	//删除活动
	Delete(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error)
	//设置限时折扣活动失效
	Cancel(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error)
	//获取活动详情
	Get(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error)
	//查询活动列表
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*LimitDiscountResponse, error)
}

type limitDiscountService struct {
	c    client.Client
	name string
}

func NewLimitDiscountService(name string, c client.Client) LimitDiscountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.ims.discount"
	}
	return &limitDiscountService{
		c:    c,
		name: name,
	}
}

func (c *limitDiscountService) Create(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Create", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountService) Update(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Update", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountService) Delete(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Delete", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountService) Cancel(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Cancel", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountService) Get(ctx context.Context, in *LimitDiscount, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Get", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*LimitDiscountResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountService.Search", in)
	out := new(LimitDiscountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LimitDiscountService service

type LimitDiscountServiceHandler interface {
	//创建限时折扣活动
	Create(context.Context, *LimitDiscount, *LimitDiscountResponse) error
	//编辑限时折扣活动
	Update(context.Context, *LimitDiscount, *LimitDiscountResponse) error
	//删除活动
	Delete(context.Context, *LimitDiscount, *LimitDiscountResponse) error
	//设置限时折扣活动失效
	Cancel(context.Context, *LimitDiscount, *LimitDiscountResponse) error
	//获取活动详情
	Get(context.Context, *LimitDiscount, *LimitDiscountResponse) error
	//查询活动列表
	Search(context.Context, *BaseWhere, *LimitDiscountResponse) error
}

func RegisterLimitDiscountServiceHandler(s server.Server, hdlr LimitDiscountServiceHandler, opts ...server.HandlerOption) error {
	type limitDiscountService interface {
		Create(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error
		Update(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error
		Delete(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error
		Cancel(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error
		Get(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error
		Search(ctx context.Context, in *BaseWhere, out *LimitDiscountResponse) error
	}
	type LimitDiscountService struct {
		limitDiscountService
	}
	h := &limitDiscountServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LimitDiscountService{h}, opts...))
}

type limitDiscountServiceHandler struct {
	LimitDiscountServiceHandler
}

func (h *limitDiscountServiceHandler) Create(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Create(ctx, in, out)
}

func (h *limitDiscountServiceHandler) Update(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Update(ctx, in, out)
}

func (h *limitDiscountServiceHandler) Delete(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Delete(ctx, in, out)
}

func (h *limitDiscountServiceHandler) Cancel(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Cancel(ctx, in, out)
}

func (h *limitDiscountServiceHandler) Get(ctx context.Context, in *LimitDiscount, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Get(ctx, in, out)
}

func (h *limitDiscountServiceHandler) Search(ctx context.Context, in *BaseWhere, out *LimitDiscountResponse) error {
	return h.LimitDiscountServiceHandler.Search(ctx, in, out)
}
