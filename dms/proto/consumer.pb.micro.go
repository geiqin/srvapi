// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: consumer.proto

package geiqin_srv_dms

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

// Client API for MyConsumerService service

type MyConsumerService interface {
	//邀请新客户
	//rpc Invite (Consumer) returns (ConsumerResponse) {}
	//获取我的客户信息
	//rpc Get (Consumer) returns (ConsumerResponse) {}
	//查询我的客户
	Search(ctx context.Context, in *ConsumerWhere, opts ...client.CallOption) (*ConsumerResponse, error)
}

type myConsumerService struct {
	c    client.Client
	name string
}

func NewMyConsumerService(name string, c client.Client) MyConsumerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.dms"
	}
	return &myConsumerService{
		c:    c,
		name: name,
	}
}

func (c *myConsumerService) Search(ctx context.Context, in *ConsumerWhere, opts ...client.CallOption) (*ConsumerResponse, error) {
	req := c.c.NewRequest(c.name, "MyConsumerService.Search", in)
	out := new(ConsumerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyConsumerService service

type MyConsumerServiceHandler interface {
	//邀请新客户
	//rpc Invite (Consumer) returns (ConsumerResponse) {}
	//获取我的客户信息
	//rpc Get (Consumer) returns (ConsumerResponse) {}
	//查询我的客户
	Search(context.Context, *ConsumerWhere, *ConsumerResponse) error
}

func RegisterMyConsumerServiceHandler(s server.Server, hdlr MyConsumerServiceHandler, opts ...server.HandlerOption) error {
	type myConsumerService interface {
		Search(ctx context.Context, in *ConsumerWhere, out *ConsumerResponse) error
	}
	type MyConsumerService struct {
		myConsumerService
	}
	h := &myConsumerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyConsumerService{h}, opts...))
}

type myConsumerServiceHandler struct {
	MyConsumerServiceHandler
}

func (h *myConsumerServiceHandler) Search(ctx context.Context, in *ConsumerWhere, out *ConsumerResponse) error {
	return h.MyConsumerServiceHandler.Search(ctx, in, out)
}

// Client API for ConsumerService service

type ConsumerService interface {
	//分页查询客户
	Search(ctx context.Context, in *ConsumerWhere, opts ...client.CallOption) (*ConsumerResponse, error)
}

type consumerService struct {
	c    client.Client
	name string
}

func NewConsumerService(name string, c client.Client) ConsumerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.dms"
	}
	return &consumerService{
		c:    c,
		name: name,
	}
}

func (c *consumerService) Search(ctx context.Context, in *ConsumerWhere, opts ...client.CallOption) (*ConsumerResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumerService.Search", in)
	out := new(ConsumerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsumerService service

type ConsumerServiceHandler interface {
	//分页查询客户
	Search(context.Context, *ConsumerWhere, *ConsumerResponse) error
}

func RegisterConsumerServiceHandler(s server.Server, hdlr ConsumerServiceHandler, opts ...server.HandlerOption) error {
	type consumerService interface {
		Search(ctx context.Context, in *ConsumerWhere, out *ConsumerResponse) error
	}
	type ConsumerService struct {
		consumerService
	}
	h := &consumerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConsumerService{h}, opts...))
}

type consumerServiceHandler struct {
	ConsumerServiceHandler
}

func (h *consumerServiceHandler) Search(ctx context.Context, in *ConsumerWhere, out *ConsumerResponse) error {
	return h.ConsumerServiceHandler.Search(ctx, in, out)
}
