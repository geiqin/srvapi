// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payment.proto

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

// Client API for Payment service

type PaymentService interface {
	//选择支付方式
	Choose(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//余额支付
	Balance(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//信用支付
	Credit(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//微信APP支付
	WxApp(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//微信小程序支付
	WxMini(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//支付宝手机网页支付
	AliWap(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
	//支付宝APP支付
	AliApp(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Choose(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.Choose", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Balance(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.Balance", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Credit(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.Credit", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) WxApp(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.WxApp", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) WxMini(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.WxMini", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) AliWap(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.AliWap", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) AliApp(ctx context.Context, in *PayRequest, opts ...client.CallOption) (*PayResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.AliApp", in)
	out := new(PayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payment service

type PaymentHandler interface {
	//选择支付方式
	Choose(context.Context, *PayRequest, *PayResponse) error
	//余额支付
	Balance(context.Context, *PayRequest, *PayResponse) error
	//信用支付
	Credit(context.Context, *PayRequest, *PayResponse) error
	//微信APP支付
	WxApp(context.Context, *PayRequest, *PayResponse) error
	//微信小程序支付
	WxMini(context.Context, *PayRequest, *PayResponse) error
	//支付宝手机网页支付
	AliWap(context.Context, *PayRequest, *PayResponse) error
	//支付宝APP支付
	AliApp(context.Context, *PayRequest, *PayResponse) error
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
		Choose(ctx context.Context, in *PayRequest, out *PayResponse) error
		Balance(ctx context.Context, in *PayRequest, out *PayResponse) error
		Credit(ctx context.Context, in *PayRequest, out *PayResponse) error
		WxApp(ctx context.Context, in *PayRequest, out *PayResponse) error
		WxMini(ctx context.Context, in *PayRequest, out *PayResponse) error
		AliWap(ctx context.Context, in *PayRequest, out *PayResponse) error
		AliApp(ctx context.Context, in *PayRequest, out *PayResponse) error
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

func (h *paymentHandler) Choose(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.Choose(ctx, in, out)
}

func (h *paymentHandler) Balance(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.Balance(ctx, in, out)
}

func (h *paymentHandler) Credit(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.Credit(ctx, in, out)
}

func (h *paymentHandler) WxApp(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.WxApp(ctx, in, out)
}

func (h *paymentHandler) WxMini(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.WxMini(ctx, in, out)
}

func (h *paymentHandler) AliWap(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.AliWap(ctx, in, out)
}

func (h *paymentHandler) AliApp(ctx context.Context, in *PayRequest, out *PayResponse) error {
	return h.PaymentHandler.AliApp(ctx, in, out)
}
