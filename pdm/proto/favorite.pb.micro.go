// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: favorite.proto

package geiqin_srv_pdm

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

// Client API for FavoriteService service

type FavoriteService interface {
	Search(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error)
}

type favoriteService struct {
	c    client.Client
	name string
}

func NewFavoriteService(name string, c client.Client) FavoriteService {
	return &favoriteService{
		c:    c,
		name: name,
	}
}

func (c *favoriteService) Search(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.Search", in)
	out := new(FavoriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FavoriteService service

type FavoriteServiceHandler interface {
	Search(context.Context, *FavoriteWhere, *FavoriteResponse) error
}

func RegisterFavoriteServiceHandler(s server.Server, hdlr FavoriteServiceHandler, opts ...server.HandlerOption) error {
	type favoriteService interface {
		Search(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error
	}
	type FavoriteService struct {
		favoriteService
	}
	h := &favoriteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FavoriteService{h}, opts...))
}

type favoriteServiceHandler struct {
	FavoriteServiceHandler
}

func (h *favoriteServiceHandler) Search(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error {
	return h.FavoriteServiceHandler.Search(ctx, in, out)
}

// Client API for MyFavoriteService service

type MyFavoriteService interface {
	Create(ctx context.Context, in *Favorite, opts ...client.CallOption) (*FavoriteResponse, error)
	Delete(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error)
	Check(ctx context.Context, in *Favorite, opts ...client.CallOption) (*FavoriteResponse, error)
	Search(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error)
}

type myFavoriteService struct {
	c    client.Client
	name string
}

func NewMyFavoriteService(name string, c client.Client) MyFavoriteService {
	return &myFavoriteService{
		c:    c,
		name: name,
	}
}

func (c *myFavoriteService) Create(ctx context.Context, in *Favorite, opts ...client.CallOption) (*FavoriteResponse, error) {
	req := c.c.NewRequest(c.name, "MyFavoriteService.Create", in)
	out := new(FavoriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myFavoriteService) Delete(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error) {
	req := c.c.NewRequest(c.name, "MyFavoriteService.Delete", in)
	out := new(FavoriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myFavoriteService) Check(ctx context.Context, in *Favorite, opts ...client.CallOption) (*FavoriteResponse, error) {
	req := c.c.NewRequest(c.name, "MyFavoriteService.Check", in)
	out := new(FavoriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myFavoriteService) Search(ctx context.Context, in *FavoriteWhere, opts ...client.CallOption) (*FavoriteResponse, error) {
	req := c.c.NewRequest(c.name, "MyFavoriteService.Search", in)
	out := new(FavoriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyFavoriteService service

type MyFavoriteServiceHandler interface {
	Create(context.Context, *Favorite, *FavoriteResponse) error
	Delete(context.Context, *FavoriteWhere, *FavoriteResponse) error
	Check(context.Context, *Favorite, *FavoriteResponse) error
	Search(context.Context, *FavoriteWhere, *FavoriteResponse) error
}

func RegisterMyFavoriteServiceHandler(s server.Server, hdlr MyFavoriteServiceHandler, opts ...server.HandlerOption) error {
	type myFavoriteService interface {
		Create(ctx context.Context, in *Favorite, out *FavoriteResponse) error
		Delete(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error
		Check(ctx context.Context, in *Favorite, out *FavoriteResponse) error
		Search(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error
	}
	type MyFavoriteService struct {
		myFavoriteService
	}
	h := &myFavoriteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyFavoriteService{h}, opts...))
}

type myFavoriteServiceHandler struct {
	MyFavoriteServiceHandler
}

func (h *myFavoriteServiceHandler) Create(ctx context.Context, in *Favorite, out *FavoriteResponse) error {
	return h.MyFavoriteServiceHandler.Create(ctx, in, out)
}

func (h *myFavoriteServiceHandler) Delete(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error {
	return h.MyFavoriteServiceHandler.Delete(ctx, in, out)
}

func (h *myFavoriteServiceHandler) Check(ctx context.Context, in *Favorite, out *FavoriteResponse) error {
	return h.MyFavoriteServiceHandler.Check(ctx, in, out)
}

func (h *myFavoriteServiceHandler) Search(ctx context.Context, in *FavoriteWhere, out *FavoriteResponse) error {
	return h.MyFavoriteServiceHandler.Search(ctx, in, out)
}
