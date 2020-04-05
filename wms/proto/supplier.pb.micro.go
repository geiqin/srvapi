// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: supplier.proto

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

// Client API for SupplierService service

type SupplierService interface {
	Create(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error)
	Update(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error)
	Delete(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error)
	Get(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error)
	Search(ctx context.Context, in *SupplierWhere, opts ...client.CallOption) (*SupplierResponse, error)
}

type supplierService struct {
	c    client.Client
	name string
}

func NewSupplierService(name string, c client.Client) SupplierService {
	return &supplierService{
		c:    c,
		name: name,
	}
}

func (c *supplierService) Create(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierService.Create", in)
	out := new(SupplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierService) Update(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierService.Update", in)
	out := new(SupplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierService) Delete(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierService.Delete", in)
	out := new(SupplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierService) Get(ctx context.Context, in *Supplier, opts ...client.CallOption) (*SupplierResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierService.Get", in)
	out := new(SupplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierService) Search(ctx context.Context, in *SupplierWhere, opts ...client.CallOption) (*SupplierResponse, error) {
	req := c.c.NewRequest(c.name, "SupplierService.Search", in)
	out := new(SupplierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SupplierService service

type SupplierServiceHandler interface {
	Create(context.Context, *Supplier, *SupplierResponse) error
	Update(context.Context, *Supplier, *SupplierResponse) error
	Delete(context.Context, *Supplier, *SupplierResponse) error
	Get(context.Context, *Supplier, *SupplierResponse) error
	Search(context.Context, *SupplierWhere, *SupplierResponse) error
}

func RegisterSupplierServiceHandler(s server.Server, hdlr SupplierServiceHandler, opts ...server.HandlerOption) error {
	type supplierService interface {
		Create(ctx context.Context, in *Supplier, out *SupplierResponse) error
		Update(ctx context.Context, in *Supplier, out *SupplierResponse) error
		Delete(ctx context.Context, in *Supplier, out *SupplierResponse) error
		Get(ctx context.Context, in *Supplier, out *SupplierResponse) error
		Search(ctx context.Context, in *SupplierWhere, out *SupplierResponse) error
	}
	type SupplierService struct {
		supplierService
	}
	h := &supplierServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SupplierService{h}, opts...))
}

type supplierServiceHandler struct {
	SupplierServiceHandler
}

func (h *supplierServiceHandler) Create(ctx context.Context, in *Supplier, out *SupplierResponse) error {
	return h.SupplierServiceHandler.Create(ctx, in, out)
}

func (h *supplierServiceHandler) Update(ctx context.Context, in *Supplier, out *SupplierResponse) error {
	return h.SupplierServiceHandler.Update(ctx, in, out)
}

func (h *supplierServiceHandler) Delete(ctx context.Context, in *Supplier, out *SupplierResponse) error {
	return h.SupplierServiceHandler.Delete(ctx, in, out)
}

func (h *supplierServiceHandler) Get(ctx context.Context, in *Supplier, out *SupplierResponse) error {
	return h.SupplierServiceHandler.Get(ctx, in, out)
}

func (h *supplierServiceHandler) Search(ctx context.Context, in *SupplierWhere, out *SupplierResponse) error {
	return h.SupplierServiceHandler.Search(ctx, in, out)
}
