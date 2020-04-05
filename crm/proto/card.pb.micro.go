// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: card.proto

package geiqin_srv_crm

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

// Client API for CardService service

type CardService interface {
	//创建卡
	Create(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//修改卡
	Update(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//删除卡
	Delete(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//禁用卡
	Disabled(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//启用卡
	Enabled(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//获得卡
	Get(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//查询卡列表(未删除的权益卡列表)
	List(ctx context.Context, in *CardWhere, opts ...client.CallOption) (*CardResponse, error)
	//有效卡列表
	ValidList(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error)
	//查询卡列表
	Search(ctx context.Context, in *CardWhere, opts ...client.CallOption) (*CardResponse, error)
}

type cardService struct {
	c    client.Client
	name string
}

func NewCardService(name string, c client.Client) CardService {
	return &cardService{
		c:    c,
		name: name,
	}
}

func (c *cardService) Create(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Create", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Update(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Update", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Delete(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Delete", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Disabled(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Disabled", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Enabled(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Enabled", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Get(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Get", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) List(ctx context.Context, in *CardWhere, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.List", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) ValidList(ctx context.Context, in *Card, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.ValidList", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardService) Search(ctx context.Context, in *CardWhere, opts ...client.CallOption) (*CardResponse, error) {
	req := c.c.NewRequest(c.name, "CardService.Search", in)
	out := new(CardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CardService service

type CardServiceHandler interface {
	//创建卡
	Create(context.Context, *Card, *CardResponse) error
	//修改卡
	Update(context.Context, *Card, *CardResponse) error
	//删除卡
	Delete(context.Context, *Card, *CardResponse) error
	//禁用卡
	Disabled(context.Context, *Card, *CardResponse) error
	//启用卡
	Enabled(context.Context, *Card, *CardResponse) error
	//获得卡
	Get(context.Context, *Card, *CardResponse) error
	//查询卡列表(未删除的权益卡列表)
	List(context.Context, *CardWhere, *CardResponse) error
	//有效卡列表
	ValidList(context.Context, *Card, *CardResponse) error
	//查询卡列表
	Search(context.Context, *CardWhere, *CardResponse) error
}

func RegisterCardServiceHandler(s server.Server, hdlr CardServiceHandler, opts ...server.HandlerOption) error {
	type cardService interface {
		Create(ctx context.Context, in *Card, out *CardResponse) error
		Update(ctx context.Context, in *Card, out *CardResponse) error
		Delete(ctx context.Context, in *Card, out *CardResponse) error
		Disabled(ctx context.Context, in *Card, out *CardResponse) error
		Enabled(ctx context.Context, in *Card, out *CardResponse) error
		Get(ctx context.Context, in *Card, out *CardResponse) error
		List(ctx context.Context, in *CardWhere, out *CardResponse) error
		ValidList(ctx context.Context, in *Card, out *CardResponse) error
		Search(ctx context.Context, in *CardWhere, out *CardResponse) error
	}
	type CardService struct {
		cardService
	}
	h := &cardServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CardService{h}, opts...))
}

type cardServiceHandler struct {
	CardServiceHandler
}

func (h *cardServiceHandler) Create(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Create(ctx, in, out)
}

func (h *cardServiceHandler) Update(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Update(ctx, in, out)
}

func (h *cardServiceHandler) Delete(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Delete(ctx, in, out)
}

func (h *cardServiceHandler) Disabled(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Disabled(ctx, in, out)
}

func (h *cardServiceHandler) Enabled(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Enabled(ctx, in, out)
}

func (h *cardServiceHandler) Get(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.Get(ctx, in, out)
}

func (h *cardServiceHandler) List(ctx context.Context, in *CardWhere, out *CardResponse) error {
	return h.CardServiceHandler.List(ctx, in, out)
}

func (h *cardServiceHandler) ValidList(ctx context.Context, in *Card, out *CardResponse) error {
	return h.CardServiceHandler.ValidList(ctx, in, out)
}

func (h *cardServiceHandler) Search(ctx context.Context, in *CardWhere, out *CardResponse) error {
	return h.CardServiceHandler.Search(ctx, in, out)
}
