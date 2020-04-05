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

// Client API for AddressService service

type AddressService interface {
	Create(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Update(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Delete(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Get(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Search(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error)
	List(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error)
}

type addressService struct {
	c    client.Client
	name string
}

func NewAddressService(name string, c client.Client) AddressService {
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

func (c *addressService) Delete(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
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

func (c *addressService) Search(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.Search", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressService) List(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "AddressService.List", in)
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
	Delete(context.Context, *Address, *AddressResponse) error
	Get(context.Context, *Address, *AddressResponse) error
	Search(context.Context, *AddressWhere, *AddressResponse) error
	List(context.Context, *AddressWhere, *AddressResponse) error
}

func RegisterAddressServiceHandler(s server.Server, hdlr AddressServiceHandler, opts ...server.HandlerOption) error {
	type addressService interface {
		Create(ctx context.Context, in *Address, out *AddressResponse) error
		Update(ctx context.Context, in *Address, out *AddressResponse) error
		Delete(ctx context.Context, in *Address, out *AddressResponse) error
		Get(ctx context.Context, in *Address, out *AddressResponse) error
		Search(ctx context.Context, in *AddressWhere, out *AddressResponse) error
		List(ctx context.Context, in *AddressWhere, out *AddressResponse) error
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

func (h *addressServiceHandler) Delete(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.AddressServiceHandler.Delete(ctx, in, out)
}

func (h *addressServiceHandler) Get(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.AddressServiceHandler.Get(ctx, in, out)
}

func (h *addressServiceHandler) Search(ctx context.Context, in *AddressWhere, out *AddressResponse) error {
	return h.AddressServiceHandler.Search(ctx, in, out)
}

func (h *addressServiceHandler) List(ctx context.Context, in *AddressWhere, out *AddressResponse) error {
	return h.AddressServiceHandler.List(ctx, in, out)
}

// Client API for MyAddressService service

type MyAddressService interface {
	Create(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Update(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Delete(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	Get(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error)
	GetDefault(ctx context.Context, in *Empty, opts ...client.CallOption) (*AddressResponse, error)
	List(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error)
}

type myAddressService struct {
	c    client.Client
	name string
}

func NewMyAddressService(name string, c client.Client) MyAddressService {
	return &myAddressService{
		c:    c,
		name: name,
	}
}

func (c *myAddressService) Create(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.Create", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAddressService) Update(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.Update", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAddressService) Delete(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.Delete", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAddressService) Get(ctx context.Context, in *Address, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.Get", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAddressService) GetDefault(ctx context.Context, in *Empty, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.GetDefault", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAddressService) List(ctx context.Context, in *AddressWhere, opts ...client.CallOption) (*AddressResponse, error) {
	req := c.c.NewRequest(c.name, "MyAddressService.List", in)
	out := new(AddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyAddressService service

type MyAddressServiceHandler interface {
	Create(context.Context, *Address, *AddressResponse) error
	Update(context.Context, *Address, *AddressResponse) error
	Delete(context.Context, *Address, *AddressResponse) error
	Get(context.Context, *Address, *AddressResponse) error
	GetDefault(context.Context, *Empty, *AddressResponse) error
	List(context.Context, *AddressWhere, *AddressResponse) error
}

func RegisterMyAddressServiceHandler(s server.Server, hdlr MyAddressServiceHandler, opts ...server.HandlerOption) error {
	type myAddressService interface {
		Create(ctx context.Context, in *Address, out *AddressResponse) error
		Update(ctx context.Context, in *Address, out *AddressResponse) error
		Delete(ctx context.Context, in *Address, out *AddressResponse) error
		Get(ctx context.Context, in *Address, out *AddressResponse) error
		GetDefault(ctx context.Context, in *Empty, out *AddressResponse) error
		List(ctx context.Context, in *AddressWhere, out *AddressResponse) error
	}
	type MyAddressService struct {
		myAddressService
	}
	h := &myAddressServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyAddressService{h}, opts...))
}

type myAddressServiceHandler struct {
	MyAddressServiceHandler
}

func (h *myAddressServiceHandler) Create(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.MyAddressServiceHandler.Create(ctx, in, out)
}

func (h *myAddressServiceHandler) Update(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.MyAddressServiceHandler.Update(ctx, in, out)
}

func (h *myAddressServiceHandler) Delete(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.MyAddressServiceHandler.Delete(ctx, in, out)
}

func (h *myAddressServiceHandler) Get(ctx context.Context, in *Address, out *AddressResponse) error {
	return h.MyAddressServiceHandler.Get(ctx, in, out)
}

func (h *myAddressServiceHandler) GetDefault(ctx context.Context, in *Empty, out *AddressResponse) error {
	return h.MyAddressServiceHandler.GetDefault(ctx, in, out)
}

func (h *myAddressServiceHandler) List(ctx context.Context, in *AddressWhere, out *AddressResponse) error {
	return h.MyAddressServiceHandler.List(ctx, in, out)
}
