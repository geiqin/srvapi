// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: applier.proto

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

// Client API for MyApplierService service

type MyApplierService interface {
	// 检查用户是否存在可申请的分销等级
	Check(ctx context.Context, in *Empty, opts ...client.CallOption) (*ApplierResponse, error)
	// 发起申请销售员
	Apply(ctx context.Context, in *Empty, opts ...client.CallOption) (*ApplyInfoResponse, error)
	// 提交销售员申请
	Submit(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error)
	// 获取申请者信息
	Get(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error)
	// 检查用户是否已经提交申请信息
	Exists(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error)
}

type myApplierService struct {
	c    client.Client
	name string
}

func NewMyApplierService(name string, c client.Client) MyApplierService {
	return &myApplierService{
		c:    c,
		name: name,
	}
}

func (c *myApplierService) Check(ctx context.Context, in *Empty, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "MyApplierService.Check", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myApplierService) Apply(ctx context.Context, in *Empty, opts ...client.CallOption) (*ApplyInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyApplierService.Apply", in)
	out := new(ApplyInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myApplierService) Submit(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "MyApplierService.Submit", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myApplierService) Get(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "MyApplierService.Get", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myApplierService) Exists(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "MyApplierService.Exists", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyApplierService service

type MyApplierServiceHandler interface {
	// 检查用户是否存在可申请的分销等级
	Check(context.Context, *Empty, *ApplierResponse) error
	// 发起申请销售员
	Apply(context.Context, *Empty, *ApplyInfoResponse) error
	// 提交销售员申请
	Submit(context.Context, *Applier, *ApplierResponse) error
	// 获取申请者信息
	Get(context.Context, *Applier, *ApplierResponse) error
	// 检查用户是否已经提交申请信息
	Exists(context.Context, *Applier, *ApplierResponse) error
}

func RegisterMyApplierServiceHandler(s server.Server, hdlr MyApplierServiceHandler, opts ...server.HandlerOption) error {
	type myApplierService interface {
		Check(ctx context.Context, in *Empty, out *ApplierResponse) error
		Apply(ctx context.Context, in *Empty, out *ApplyInfoResponse) error
		Submit(ctx context.Context, in *Applier, out *ApplierResponse) error
		Get(ctx context.Context, in *Applier, out *ApplierResponse) error
		Exists(ctx context.Context, in *Applier, out *ApplierResponse) error
	}
	type MyApplierService struct {
		myApplierService
	}
	h := &myApplierServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyApplierService{h}, opts...))
}

type myApplierServiceHandler struct {
	MyApplierServiceHandler
}

func (h *myApplierServiceHandler) Check(ctx context.Context, in *Empty, out *ApplierResponse) error {
	return h.MyApplierServiceHandler.Check(ctx, in, out)
}

func (h *myApplierServiceHandler) Apply(ctx context.Context, in *Empty, out *ApplyInfoResponse) error {
	return h.MyApplierServiceHandler.Apply(ctx, in, out)
}

func (h *myApplierServiceHandler) Submit(ctx context.Context, in *Applier, out *ApplierResponse) error {
	return h.MyApplierServiceHandler.Submit(ctx, in, out)
}

func (h *myApplierServiceHandler) Get(ctx context.Context, in *Applier, out *ApplierResponse) error {
	return h.MyApplierServiceHandler.Get(ctx, in, out)
}

func (h *myApplierServiceHandler) Exists(ctx context.Context, in *Applier, out *ApplierResponse) error {
	return h.MyApplierServiceHandler.Exists(ctx, in, out)
}

// Client API for ApplierService service

type ApplierService interface {
	//审核申请者
	Check(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error)
	//获取申请者信息
	Get(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error)
	//分页查询申请者
	Search(ctx context.Context, in *ApplierWhere, opts ...client.CallOption) (*ApplierResponse, error)
}

type applierService struct {
	c    client.Client
	name string
}

func NewApplierService(name string, c client.Client) ApplierService {
	return &applierService{
		c:    c,
		name: name,
	}
}

func (c *applierService) Check(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "ApplierService.Check", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applierService) Get(ctx context.Context, in *Applier, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "ApplierService.Get", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applierService) Search(ctx context.Context, in *ApplierWhere, opts ...client.CallOption) (*ApplierResponse, error) {
	req := c.c.NewRequest(c.name, "ApplierService.Search", in)
	out := new(ApplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplierService service

type ApplierServiceHandler interface {
	//审核申请者
	Check(context.Context, *Applier, *ApplierResponse) error
	//获取申请者信息
	Get(context.Context, *Applier, *ApplierResponse) error
	//分页查询申请者
	Search(context.Context, *ApplierWhere, *ApplierResponse) error
}

func RegisterApplierServiceHandler(s server.Server, hdlr ApplierServiceHandler, opts ...server.HandlerOption) error {
	type applierService interface {
		Check(ctx context.Context, in *Applier, out *ApplierResponse) error
		Get(ctx context.Context, in *Applier, out *ApplierResponse) error
		Search(ctx context.Context, in *ApplierWhere, out *ApplierResponse) error
	}
	type ApplierService struct {
		applierService
	}
	h := &applierServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ApplierService{h}, opts...))
}

type applierServiceHandler struct {
	ApplierServiceHandler
}

func (h *applierServiceHandler) Check(ctx context.Context, in *Applier, out *ApplierResponse) error {
	return h.ApplierServiceHandler.Check(ctx, in, out)
}

func (h *applierServiceHandler) Get(ctx context.Context, in *Applier, out *ApplierResponse) error {
	return h.ApplierServiceHandler.Get(ctx, in, out)
}

func (h *applierServiceHandler) Search(ctx context.Context, in *ApplierWhere, out *ApplierResponse) error {
	return h.ApplierServiceHandler.Search(ctx, in, out)
}
