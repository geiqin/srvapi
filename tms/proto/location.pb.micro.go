// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: location.proto

package geiqin_srv_tms

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

// Client API for LocationReturnService service

type LocationReturnService interface {
	Create(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error)
	Update(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error)
	Get(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	Delete(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	Search(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	List(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	GetDefaulted(ctx context.Context, in *Empty, opts ...client.CallOption) (*LocationResponse, error)
	Defaulted(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
}

type locationReturnService struct {
	c    client.Client
	name string
}

func NewLocationReturnService(name string, c client.Client) LocationReturnService {
	return &locationReturnService{
		c:    c,
		name: name,
	}
}

func (c *locationReturnService) Create(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Create", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) Update(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Update", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) Get(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Get", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) Delete(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Delete", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) Search(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Search", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) List(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.List", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) GetDefaulted(ctx context.Context, in *Empty, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.GetDefaulted", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReturnService) Defaulted(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationReturnService.Defaulted", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LocationReturnService service

type LocationReturnServiceHandler interface {
	Create(context.Context, *Location, *LocationResponse) error
	Update(context.Context, *Location, *LocationResponse) error
	Get(context.Context, *LocationWhere, *LocationResponse) error
	Delete(context.Context, *LocationWhere, *LocationResponse) error
	Search(context.Context, *LocationWhere, *LocationResponse) error
	List(context.Context, *LocationWhere, *LocationResponse) error
	GetDefaulted(context.Context, *Empty, *LocationResponse) error
	Defaulted(context.Context, *LocationWhere, *LocationResponse) error
}

func RegisterLocationReturnServiceHandler(s server.Server, hdlr LocationReturnServiceHandler, opts ...server.HandlerOption) error {
	type locationReturnService interface {
		Create(ctx context.Context, in *Location, out *LocationResponse) error
		Update(ctx context.Context, in *Location, out *LocationResponse) error
		Get(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		Delete(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		Search(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		List(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		GetDefaulted(ctx context.Context, in *Empty, out *LocationResponse) error
		Defaulted(ctx context.Context, in *LocationWhere, out *LocationResponse) error
	}
	type LocationReturnService struct {
		locationReturnService
	}
	h := &locationReturnServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LocationReturnService{h}, opts...))
}

type locationReturnServiceHandler struct {
	LocationReturnServiceHandler
}

func (h *locationReturnServiceHandler) Create(ctx context.Context, in *Location, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Create(ctx, in, out)
}

func (h *locationReturnServiceHandler) Update(ctx context.Context, in *Location, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Update(ctx, in, out)
}

func (h *locationReturnServiceHandler) Get(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Get(ctx, in, out)
}

func (h *locationReturnServiceHandler) Delete(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Delete(ctx, in, out)
}

func (h *locationReturnServiceHandler) Search(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Search(ctx, in, out)
}

func (h *locationReturnServiceHandler) List(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.List(ctx, in, out)
}

func (h *locationReturnServiceHandler) GetDefaulted(ctx context.Context, in *Empty, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.GetDefaulted(ctx, in, out)
}

func (h *locationReturnServiceHandler) Defaulted(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationReturnServiceHandler.Defaulted(ctx, in, out)
}

// Client API for LocationDeliverService service

type LocationDeliverService interface {
	Create(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error)
	Update(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error)
	Get(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	Delete(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	Search(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	List(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
	GetDefaulted(ctx context.Context, in *Empty, opts ...client.CallOption) (*LocationResponse, error)
	Defaulted(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error)
}

type locationDeliverService struct {
	c    client.Client
	name string
}

func NewLocationDeliverService(name string, c client.Client) LocationDeliverService {
	return &locationDeliverService{
		c:    c,
		name: name,
	}
}

func (c *locationDeliverService) Create(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Create", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) Update(ctx context.Context, in *Location, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Update", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) Get(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Get", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) Delete(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Delete", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) Search(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Search", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) List(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.List", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) GetDefaulted(ctx context.Context, in *Empty, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.GetDefaulted", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDeliverService) Defaulted(ctx context.Context, in *LocationWhere, opts ...client.CallOption) (*LocationResponse, error) {
	req := c.c.NewRequest(c.name, "LocationDeliverService.Defaulted", in)
	out := new(LocationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LocationDeliverService service

type LocationDeliverServiceHandler interface {
	Create(context.Context, *Location, *LocationResponse) error
	Update(context.Context, *Location, *LocationResponse) error
	Get(context.Context, *LocationWhere, *LocationResponse) error
	Delete(context.Context, *LocationWhere, *LocationResponse) error
	Search(context.Context, *LocationWhere, *LocationResponse) error
	List(context.Context, *LocationWhere, *LocationResponse) error
	GetDefaulted(context.Context, *Empty, *LocationResponse) error
	Defaulted(context.Context, *LocationWhere, *LocationResponse) error
}

func RegisterLocationDeliverServiceHandler(s server.Server, hdlr LocationDeliverServiceHandler, opts ...server.HandlerOption) error {
	type locationDeliverService interface {
		Create(ctx context.Context, in *Location, out *LocationResponse) error
		Update(ctx context.Context, in *Location, out *LocationResponse) error
		Get(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		Delete(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		Search(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		List(ctx context.Context, in *LocationWhere, out *LocationResponse) error
		GetDefaulted(ctx context.Context, in *Empty, out *LocationResponse) error
		Defaulted(ctx context.Context, in *LocationWhere, out *LocationResponse) error
	}
	type LocationDeliverService struct {
		locationDeliverService
	}
	h := &locationDeliverServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LocationDeliverService{h}, opts...))
}

type locationDeliverServiceHandler struct {
	LocationDeliverServiceHandler
}

func (h *locationDeliverServiceHandler) Create(ctx context.Context, in *Location, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Create(ctx, in, out)
}

func (h *locationDeliverServiceHandler) Update(ctx context.Context, in *Location, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Update(ctx, in, out)
}

func (h *locationDeliverServiceHandler) Get(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Get(ctx, in, out)
}

func (h *locationDeliverServiceHandler) Delete(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Delete(ctx, in, out)
}

func (h *locationDeliverServiceHandler) Search(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Search(ctx, in, out)
}

func (h *locationDeliverServiceHandler) List(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.List(ctx, in, out)
}

func (h *locationDeliverServiceHandler) GetDefaulted(ctx context.Context, in *Empty, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.GetDefaulted(ctx, in, out)
}

func (h *locationDeliverServiceHandler) Defaulted(ctx context.Context, in *LocationWhere, out *LocationResponse) error {
	return h.LocationDeliverServiceHandler.Defaulted(ctx, in, out)
}