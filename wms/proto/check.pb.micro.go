// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: check.proto

package geiqin_srv_wms

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

// Client API for CheckService service

type CheckService interface {
	Create(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Update(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Delete(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Confirm(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Approve(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Get(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CheckResponse, error)
	Details(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CheckDetailResponse, error)
}

type checkService struct {
	c    client.Client
	name string
}

func NewCheckService(name string, c client.Client) CheckService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.wms"
	}
	return &checkService{
		c:    c,
		name: name,
	}
}

func (c *checkService) Create(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Create", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Update(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Update", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Delete(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Delete", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Confirm(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Confirm", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Approve(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Approve", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Get(ctx context.Context, in *Check, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Get", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Search", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkService) Details(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CheckDetailResponse, error) {
	req := c.c.NewRequest(c.name, "CheckService.Details", in)
	out := new(CheckDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckService service

type CheckServiceHandler interface {
	Create(context.Context, *Check, *CheckResponse) error
	Update(context.Context, *Check, *CheckResponse) error
	Delete(context.Context, *Check, *CheckResponse) error
	Confirm(context.Context, *Check, *CheckResponse) error
	Approve(context.Context, *Check, *CheckResponse) error
	Get(context.Context, *Check, *CheckResponse) error
	Search(context.Context, *BaseWhere, *CheckResponse) error
	Details(context.Context, *BaseWhere, *CheckDetailResponse) error
}

func RegisterCheckServiceHandler(s server.Server, hdlr CheckServiceHandler, opts ...server.HandlerOption) error {
	type checkService interface {
		Create(ctx context.Context, in *Check, out *CheckResponse) error
		Update(ctx context.Context, in *Check, out *CheckResponse) error
		Delete(ctx context.Context, in *Check, out *CheckResponse) error
		Confirm(ctx context.Context, in *Check, out *CheckResponse) error
		Approve(ctx context.Context, in *Check, out *CheckResponse) error
		Get(ctx context.Context, in *Check, out *CheckResponse) error
		Search(ctx context.Context, in *BaseWhere, out *CheckResponse) error
		Details(ctx context.Context, in *BaseWhere, out *CheckDetailResponse) error
	}
	type CheckService struct {
		checkService
	}
	h := &checkServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckService{h}, opts...))
}

type checkServiceHandler struct {
	CheckServiceHandler
}

func (h *checkServiceHandler) Create(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Create(ctx, in, out)
}

func (h *checkServiceHandler) Update(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Update(ctx, in, out)
}

func (h *checkServiceHandler) Delete(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Delete(ctx, in, out)
}

func (h *checkServiceHandler) Confirm(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Confirm(ctx, in, out)
}

func (h *checkServiceHandler) Approve(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Approve(ctx, in, out)
}

func (h *checkServiceHandler) Get(ctx context.Context, in *Check, out *CheckResponse) error {
	return h.CheckServiceHandler.Get(ctx, in, out)
}

func (h *checkServiceHandler) Search(ctx context.Context, in *BaseWhere, out *CheckResponse) error {
	return h.CheckServiceHandler.Search(ctx, in, out)
}

func (h *checkServiceHandler) Details(ctx context.Context, in *BaseWhere, out *CheckDetailResponse) error {
	return h.CheckServiceHandler.Details(ctx, in, out)
}