// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: withdrawal.proto

package geiqin_srv_pay

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

// Client API for WithdrawalService service

type WithdrawalService interface {
	//创建提现
	Create(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error)
	//修改提现
	Update(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error)
	//删除提现
	Delete(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error)
	//获得提现信息
	Get(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error)
	//查询提现
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*WithdrawalResponse, error)
}

type withdrawalService struct {
	c    client.Client
	name string
}

func NewWithdrawalService(name string, c client.Client) WithdrawalService {
	return &withdrawalService{
		c:    c,
		name: name,
	}
}

func (c *withdrawalService) Create(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error) {
	req := c.c.NewRequest(c.name, "WithdrawalService.Create", in)
	out := new(WithdrawalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawalService) Update(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error) {
	req := c.c.NewRequest(c.name, "WithdrawalService.Update", in)
	out := new(WithdrawalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawalService) Delete(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error) {
	req := c.c.NewRequest(c.name, "WithdrawalService.Delete", in)
	out := new(WithdrawalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawalService) Get(ctx context.Context, in *Withdrawal, opts ...client.CallOption) (*WithdrawalResponse, error) {
	req := c.c.NewRequest(c.name, "WithdrawalService.Get", in)
	out := new(WithdrawalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawalService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*WithdrawalResponse, error) {
	req := c.c.NewRequest(c.name, "WithdrawalService.Search", in)
	out := new(WithdrawalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WithdrawalService service

type WithdrawalServiceHandler interface {
	//创建提现
	Create(context.Context, *Withdrawal, *WithdrawalResponse) error
	//修改提现
	Update(context.Context, *Withdrawal, *WithdrawalResponse) error
	//删除提现
	Delete(context.Context, *Withdrawal, *WithdrawalResponse) error
	//获得提现信息
	Get(context.Context, *Withdrawal, *WithdrawalResponse) error
	//查询提现
	Search(context.Context, *BaseWhere, *WithdrawalResponse) error
}

func RegisterWithdrawalServiceHandler(s server.Server, hdlr WithdrawalServiceHandler, opts ...server.HandlerOption) error {
	type withdrawalService interface {
		Create(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error
		Update(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error
		Delete(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error
		Get(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error
		Search(ctx context.Context, in *BaseWhere, out *WithdrawalResponse) error
	}
	type WithdrawalService struct {
		withdrawalService
	}
	h := &withdrawalServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WithdrawalService{h}, opts...))
}

type withdrawalServiceHandler struct {
	WithdrawalServiceHandler
}

func (h *withdrawalServiceHandler) Create(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error {
	return h.WithdrawalServiceHandler.Create(ctx, in, out)
}

func (h *withdrawalServiceHandler) Update(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error {
	return h.WithdrawalServiceHandler.Update(ctx, in, out)
}

func (h *withdrawalServiceHandler) Delete(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error {
	return h.WithdrawalServiceHandler.Delete(ctx, in, out)
}

func (h *withdrawalServiceHandler) Get(ctx context.Context, in *Withdrawal, out *WithdrawalResponse) error {
	return h.WithdrawalServiceHandler.Get(ctx, in, out)
}

func (h *withdrawalServiceHandler) Search(ctx context.Context, in *BaseWhere, out *WithdrawalResponse) error {
	return h.WithdrawalServiceHandler.Search(ctx, in, out)
}