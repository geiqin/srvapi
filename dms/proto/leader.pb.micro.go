// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: leader.proto

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

// Client API for MyLeaderService service

type MyLeaderService interface {
	//获取团长信息
	Get(ctx context.Context, in *Leader, opts ...client.CallOption) (*LeaderResponse, error)
	//检查用户是否是团长
	Exists(ctx context.Context, in *LeaderWhere, opts ...client.CallOption) (*LeaderResponse, error)
}

type myLeaderService struct {
	c    client.Client
	name string
}

func NewMyLeaderService(name string, c client.Client) MyLeaderService {
	return &myLeaderService{
		c:    c,
		name: name,
	}
}

func (c *myLeaderService) Get(ctx context.Context, in *Leader, opts ...client.CallOption) (*LeaderResponse, error) {
	req := c.c.NewRequest(c.name, "MyLeaderService.Get", in)
	out := new(LeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myLeaderService) Exists(ctx context.Context, in *LeaderWhere, opts ...client.CallOption) (*LeaderResponse, error) {
	req := c.c.NewRequest(c.name, "MyLeaderService.Exists", in)
	out := new(LeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyLeaderService service

type MyLeaderServiceHandler interface {
	//获取团长信息
	Get(context.Context, *Leader, *LeaderResponse) error
	//检查用户是否是团长
	Exists(context.Context, *LeaderWhere, *LeaderResponse) error
}

func RegisterMyLeaderServiceHandler(s server.Server, hdlr MyLeaderServiceHandler, opts ...server.HandlerOption) error {
	type myLeaderService interface {
		Get(ctx context.Context, in *Leader, out *LeaderResponse) error
		Exists(ctx context.Context, in *LeaderWhere, out *LeaderResponse) error
	}
	type MyLeaderService struct {
		myLeaderService
	}
	h := &myLeaderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyLeaderService{h}, opts...))
}

type myLeaderServiceHandler struct {
	MyLeaderServiceHandler
}

func (h *myLeaderServiceHandler) Get(ctx context.Context, in *Leader, out *LeaderResponse) error {
	return h.MyLeaderServiceHandler.Get(ctx, in, out)
}

func (h *myLeaderServiceHandler) Exists(ctx context.Context, in *LeaderWhere, out *LeaderResponse) error {
	return h.MyLeaderServiceHandler.Exists(ctx, in, out)
}

// Client API for LeaderService service

type LeaderService interface {
	//获取团长信息
	Get(ctx context.Context, in *Leader, opts ...client.CallOption) (*LeaderResponse, error)
	//查询团长列表
	Search(ctx context.Context, in *LeaderWhere, opts ...client.CallOption) (*LeaderResponse, error)
}

type leaderService struct {
	c    client.Client
	name string
}

func NewLeaderService(name string, c client.Client) LeaderService {
	return &leaderService{
		c:    c,
		name: name,
	}
}

func (c *leaderService) Get(ctx context.Context, in *Leader, opts ...client.CallOption) (*LeaderResponse, error) {
	req := c.c.NewRequest(c.name, "LeaderService.Get", in)
	out := new(LeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderService) Search(ctx context.Context, in *LeaderWhere, opts ...client.CallOption) (*LeaderResponse, error) {
	req := c.c.NewRequest(c.name, "LeaderService.Search", in)
	out := new(LeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LeaderService service

type LeaderServiceHandler interface {
	//获取团长信息
	Get(context.Context, *Leader, *LeaderResponse) error
	//查询团长列表
	Search(context.Context, *LeaderWhere, *LeaderResponse) error
}

func RegisterLeaderServiceHandler(s server.Server, hdlr LeaderServiceHandler, opts ...server.HandlerOption) error {
	type leaderService interface {
		Get(ctx context.Context, in *Leader, out *LeaderResponse) error
		Search(ctx context.Context, in *LeaderWhere, out *LeaderResponse) error
	}
	type LeaderService struct {
		leaderService
	}
	h := &leaderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LeaderService{h}, opts...))
}

type leaderServiceHandler struct {
	LeaderServiceHandler
}

func (h *leaderServiceHandler) Get(ctx context.Context, in *Leader, out *LeaderResponse) error {
	return h.LeaderServiceHandler.Get(ctx, in, out)
}

func (h *leaderServiceHandler) Search(ctx context.Context, in *LeaderWhere, out *LeaderResponse) error {
	return h.LeaderServiceHandler.Search(ctx, in, out)
}