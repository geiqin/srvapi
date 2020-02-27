// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: address.proto

package geiqin_srv_crm

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

// Client API for AddressService service

type AddressService interface {
	Create(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Update(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Delete(ctx context.Context, in *Id, opts ...client.CallOption) (*AddressResponse, error)
	Get(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*AddressResponse, error)
}

type addressService struct {
	c    client.Client
	name string
}

func NewAddressService(name string, c client.Client) AddressService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.crm"
	}
	return &addressService{
		c:    c,
		name: name,
	}
}

func (c *addressService) Create(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Create", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressService) Update(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Update", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressService) Delete(ctx context.Context, in *Id, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Delete", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressService) Get(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Get", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Search", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddressService service

type AddressServiceHandler interface {
	Create(context.Context, *Address, *AddressResponse) error
	Update(context.Context, *Address, *AddressResponse) error
	Delete(context.Context, *Id, *AddressResponse) error
	Get(context.Context, *Address, *AddressResponse) error
	Search(context.Context, *BaseWhere, *AddressResponse) error
}

func RegisterAddressServiceHandler(s server.Server, hdlr AddressServiceHandler, opts ...server.HandlerOption) error {
	type addressService interface {
		Create(ctx context.Context, in *Address, out *AddressResponse) error
		Update(ctx context.Context, in *Address, out *AddressResponse) error
		Delete(ctx context.Context, in *Id, out *AddressResponse) error
		Get(ctx context.Context, in *Address, out *AddressResponse) error
		Search(ctx context.Context, in *BaseWhere, out *AddressResponse) error
	}
	type AddressService struct {
		addressService
	}
	h := &addressServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AddressService{h}, opts...))
}

type addressServiceHandler struct {
	AddressServiceHandler
}

func (h *addressServiceHandler) Create(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.AddressServiceHandler.Create(ctx, in, out)
}

func (h *addressServiceHandler) Update(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.AddressServiceHandler.Update(ctx, in, out)
}

func (h *addressServiceHandler) Delete(ctx context.Context, in *Id, out *AddressResponse) error {
	return h.AddressServiceHandler.Delete(ctx, in, out)
}

func (h *addressServiceHandler) Get(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.AddressServiceHandler.Get(ctx, in, out)
}

func (h *addressServiceHandler) Search(ctx context.Context, in *BaseWhere, out *AddressResponse) error {
	return h.AddressServiceHandler.Search(ctx, in, out)
}
