// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: supplierContact.proto

package geiqin_srv_wms

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

// Client API for SupplierContactService service

type SupplierContactService interface {
	Create(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error)
	Update(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error)
	Delete(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error)
	Get(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*SupplierContactResponse, error)
}

type supplierContactService struct {
	c    client.Client
	name string
}

func NewSupplierContactService(name string, c client.Client) SupplierContactService {
	return &supplierContactService{
		c:    c,
		name: name,
	}
}

func (c *supplierContactService) Create(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierContactService.Create", in)
	out := new(SupplierContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierContactService) Update(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierContactService.Update", in)
	out := new(SupplierContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierContactService) Delete(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierContactService.Delete", in)
	out := new(SupplierContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierContactService) Get(ctx context.Context, in *SupplierContact, opts ...client.CallOption) (*SupplierContactResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierContactService.Get", in)
	out := new(SupplierContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierContactService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*SupplierContactResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierContactService.Search", in)
	out := new(SupplierContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SupplierContactService service

type SupplierContactServiceHandler interface {
	Create(context.Context, *SupplierContact, *SupplierContactResponse) error
	Update(context.Context, *SupplierContact, *SupplierContactResponse) error
	Delete(context.Context, *SupplierContact, *SupplierContactResponse) error
	Get(context.Context, *SupplierContact, *SupplierContactResponse) error
	Search(context.Context, *BaseWhere, *SupplierContactResponse) error
}

func RegisterSupplierContactServiceHandler(s server.Server, hdlr SupplierContactServiceHandler, opts ...server.HandlerOption) error {
	type supplierContactService interface {
		Create(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error
		Update(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error
		Delete(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error
		Get(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error
		Search(ctx context.Context, in *BaseWhere, out *SupplierContactResponse) error
	}
	type SupplierContactService struct {
		supplierContactService
	}
	h := &supplierContactServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SupplierContactService{h}, opts...))
}

type supplierContactServiceHandler struct {
	SupplierContactServiceHandler
}

func (h *supplierContactServiceHandler) Create(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error {
	return h.SupplierContactServiceHandler.Create(ctx, in, out)
}

func (h *supplierContactServiceHandler) Update(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error {
	return h.SupplierContactServiceHandler.Update(ctx, in, out)
}

func (h *supplierContactServiceHandler) Delete(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error {
	return h.SupplierContactServiceHandler.Delete(ctx, in, out)
}

func (h *supplierContactServiceHandler) Get(ctx context.Context, in *SupplierContact, out *SupplierContactResponse) error {
	return h.SupplierContactServiceHandler.Get(ctx, in, out)
}

func (h *supplierContactServiceHandler) Search(ctx context.Context, in *BaseWhere, out *SupplierContactResponse) error {
	return h.SupplierContactServiceHandler.Search(ctx, in, out)
}
