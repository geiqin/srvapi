// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cashAccount.proto

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

// Client API for CashAccountService service

type CashAccountService interface {
	//创建提现账户
	Create(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error)
	//删除提现账户
	Delete(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error)
	//获取提现账户
	Get(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error)
	//查询提现账户
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CashAccountResponse, error)
}

type cashAccountService struct {
	c    client.Client
	name string
}

func NewCashAccountService(name string, c client.Client) CashAccountService {
	return &cashAccountService{
		c:    c,
		name: name,
	}
}

func (c *cashAccountService) Create(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error) {
	req := c.c.NewRequest(c.name, "CashAccountService.Create", in)
	out := new(CashAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashAccountService) Delete(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error) {
	req := c.c.NewRequest(c.name, "CashAccountService.Delete", in)
	out := new(CashAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashAccountService) Get(ctx context.Context, in *CashAccount, opts ...client.CallOption) (*CashAccountResponse, error) {
	req := c.c.NewRequest(c.name, "CashAccountService.Get", in)
	out := new(CashAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashAccountService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CashAccountResponse, error) {
	req := c.c.NewRequest(c.name, "CashAccountService.Search", in)
	out := new(CashAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CashAccountService service

type CashAccountServiceHandler interface {
	//创建提现账户
	Create(context.Context, *CashAccount, *CashAccountResponse) error
	//删除提现账户
	Delete(context.Context, *CashAccount, *CashAccountResponse) error
	//获取提现账户
	Get(context.Context, *CashAccount, *CashAccountResponse) error
	//查询提现账户
	Search(context.Context, *BaseWhere, *CashAccountResponse) error
}

func RegisterCashAccountServiceHandler(s server.Server, hdlr CashAccountServiceHandler, opts ...server.HandlerOption) error {
	type cashAccountService interface {
		Create(ctx context.Context, in *CashAccount, out *CashAccountResponse) error
		Delete(ctx context.Context, in *CashAccount, out *CashAccountResponse) error
		Get(ctx context.Context, in *CashAccount, out *CashAccountResponse) error
		Search(ctx context.Context, in *BaseWhere, out *CashAccountResponse) error
	}
	type CashAccountService struct {
		cashAccountService
	}
	h := &cashAccountServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CashAccountService{h}, opts...))
}

type cashAccountServiceHandler struct {
	CashAccountServiceHandler
}

func (h *cashAccountServiceHandler) Create(ctx context.Context, in *CashAccount, out *CashAccountResponse) error {
	return h.CashAccountServiceHandler.Create(ctx, in, out)
}

func (h *cashAccountServiceHandler) Delete(ctx context.Context, in *CashAccount, out *CashAccountResponse) error {
	return h.CashAccountServiceHandler.Delete(ctx, in, out)
}

func (h *cashAccountServiceHandler) Get(ctx context.Context, in *CashAccount, out *CashAccountResponse) error {
	return h.CashAccountServiceHandler.Get(ctx, in, out)
}

func (h *cashAccountServiceHandler) Search(ctx context.Context, in *BaseWhere, out *CashAccountResponse) error {
	return h.CashAccountServiceHandler.Search(ctx, in, out)
}
