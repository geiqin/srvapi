// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: item.proto

package geiqin_srv_pdm

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

// Client API for ItemService service

type ItemService interface {
	Create(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error)
	Update(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error)
	Delete(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemResponse, error)
	Lock(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Unlock(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Listing(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Delisting(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Recovery(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Destroy(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error)
	Get(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemResponse, error)
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*ItemResponse, error)
	SearchDeleted(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*ItemResponse, error)
}

type itemService struct {
	c    client.Client
	name string
}

func NewItemService(name string, c client.Client) ItemService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.pdm"
	}
	return &itemService{
		c:    c,
		name: name,
	}
}

func (c *itemService) Create(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Create", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Update(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Update", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Delete(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Delete", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Lock(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Lock", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Unlock(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Unlock", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Listing(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Listing", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Delisting(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Delisting", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Recovery(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Recovery", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Destroy(ctx context.Context, in *Ids, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Destroy", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Get(ctx context.Context, in *Id, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Get", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Search", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) SearchDeleted(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.SearchDeleted", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemService service

type ItemServiceHandler interface {
	Create(context.Context, *Item, *ItemResponse) error
	Update(context.Context, *Item, *ItemResponse) error
	Delete(context.Context, *Id, *ItemResponse) error
	Lock(context.Context, *Ids, *ItemResponse) error
	Unlock(context.Context, *Ids, *ItemResponse) error
	Listing(context.Context, *Ids, *ItemResponse) error
	Delisting(context.Context, *Ids, *ItemResponse) error
	Recovery(context.Context, *Ids, *ItemResponse) error
	Destroy(context.Context, *Ids, *ItemResponse) error
	Get(context.Context, *Id, *ItemResponse) error
	Search(context.Context, *BaseWhere, *ItemResponse) error
	SearchDeleted(context.Context, *BaseWhere, *ItemResponse) error
}

func RegisterItemServiceHandler(s server.Server, hdlr ItemServiceHandler, opts ...server.HandlerOption) error {
	type itemService interface {
		Create(ctx context.Context, in *Item, out *ItemResponse) error
		Update(ctx context.Context, in *Item, out *ItemResponse) error
		Delete(ctx context.Context, in *Id, out *ItemResponse) error
		Lock(ctx context.Context, in *Ids, out *ItemResponse) error
		Unlock(ctx context.Context, in *Ids, out *ItemResponse) error
		Listing(ctx context.Context, in *Ids, out *ItemResponse) error
		Delisting(ctx context.Context, in *Ids, out *ItemResponse) error
		Recovery(ctx context.Context, in *Ids, out *ItemResponse) error
		Destroy(ctx context.Context, in *Ids, out *ItemResponse) error
		Get(ctx context.Context, in *Id, out *ItemResponse) error
		Search(ctx context.Context, in *BaseWhere, out *ItemResponse) error
		SearchDeleted(ctx context.Context, in *BaseWhere, out *ItemResponse) error
	}
	type ItemService struct {
		itemService
	}
	h := &itemServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ItemService{h}, opts...))
}

type itemServiceHandler struct {
	ItemServiceHandler
}

func (h *itemServiceHandler) Create(ctx context.Context, in *Item, out *ItemResponse) error {
	return h.ItemServiceHandler.Create(ctx, in, out)
}

func (h *itemServiceHandler) Update(ctx context.Context, in *Item, out *ItemResponse) error {
	return h.ItemServiceHandler.Update(ctx, in, out)
}

func (h *itemServiceHandler) Delete(ctx context.Context, in *Id, out *ItemResponse) error {
	return h.ItemServiceHandler.Delete(ctx, in, out)
}

func (h *itemServiceHandler) Lock(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Lock(ctx, in, out)
}

func (h *itemServiceHandler) Unlock(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Unlock(ctx, in, out)
}

func (h *itemServiceHandler) Listing(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Listing(ctx, in, out)
}

func (h *itemServiceHandler) Delisting(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Delisting(ctx, in, out)
}

func (h *itemServiceHandler) Recovery(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Recovery(ctx, in, out)
}

func (h *itemServiceHandler) Destroy(ctx context.Context, in *Ids, out *ItemResponse) error {
	return h.ItemServiceHandler.Destroy(ctx, in, out)
}

func (h *itemServiceHandler) Get(ctx context.Context, in *Id, out *ItemResponse) error {
	return h.ItemServiceHandler.Get(ctx, in, out)
}

func (h *itemServiceHandler) Search(ctx context.Context, in *BaseWhere, out *ItemResponse) error {
	return h.ItemServiceHandler.Search(ctx, in, out)
}

func (h *itemServiceHandler) SearchDeleted(ctx context.Context, in *BaseWhere, out *ItemResponse) error {
	return h.ItemServiceHandler.SearchDeleted(ctx, in, out)
}
