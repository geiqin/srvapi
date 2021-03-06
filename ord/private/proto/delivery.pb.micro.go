// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: delivery.proto

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

// Client API for DeliveryService service

type DeliveryService interface {
	// 获取订单发货列表
	List(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error)
}

type deliveryService struct {
	c    client.Client
	name string
}

func NewDeliveryService(name string, c client.Client) DeliveryService {
	return &deliveryService{
		c:    c,
		name: name,
	}
}

func (c *deliveryService) List(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryService.List", in)
	out := new(DeliveryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeliveryService service

type DeliveryServiceHandler interface {
	// 获取订单发货列表
	List(context.Context, *Delivery, *DeliveryResponse) error
}

func RegisterDeliveryServiceHandler(s server.Server, hdlr DeliveryServiceHandler, opts ...server.HandlerOption) error {
	type deliveryService interface {
		List(ctx context.Context, in *Delivery, out *DeliveryResponse) error
	}
	type DeliveryService struct {
		deliveryService
	}
	h := &deliveryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeliveryService{h}, opts...))
}

type deliveryServiceHandler struct {
	DeliveryServiceHandler
}

func (h *deliveryServiceHandler) List(ctx context.Context, in *Delivery, out *DeliveryResponse) error {
	return h.DeliveryServiceHandler.List(ctx, in, out)
}
