// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: refunding.proto

package geiqin_srv_crm_integral

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

// Client API for RefundingService service

type RefundingService interface {
	// 计算退款时退回、扣除的积分及积分不足扣除时抵扣的金额
	Calculate(ctx context.Context, in *Refunding, opts ...client.CallOption) (*RefundingResponse, error)
}

type refundingService struct {
	c    client.Client
	name string
}

func NewRefundingService(name string, c client.Client) RefundingService {
	return &refundingService{
		c:    c,
		name: name,
	}
}

func (c *refundingService) Calculate(ctx context.Context, in *Refunding, opts ...client.CallOption) (*RefundingResponse, error) {
	req := c.c.NewRequest(c.name, "RefundingService.Calculate", in)
	out := new(RefundingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RefundingService service

type RefundingServiceHandler interface {
	// 计算退款时退回、扣除的积分及积分不足扣除时抵扣的金额
	Calculate(context.Context, *Refunding, *RefundingResponse) error
}

func RegisterRefundingServiceHandler(s server.Server, hdlr RefundingServiceHandler, opts ...server.HandlerOption) error {
	type refundingService interface {
		Calculate(ctx context.Context, in *Refunding, out *RefundingResponse) error
	}
	type RefundingService struct {
		refundingService
	}
	h := &refundingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RefundingService{h}, opts...))
}

type refundingServiceHandler struct {
	RefundingServiceHandler
}

func (h *refundingServiceHandler) Calculate(ctx context.Context, in *Refunding, out *RefundingResponse) error {
	return h.RefundingServiceHandler.Calculate(ctx, in, out)
}
