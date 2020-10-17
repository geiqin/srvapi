// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: table.proto

package geiqin_srv_eat

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

// Client API for TableService service

type TableService interface {
	Search(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
	List(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
	Get(ctx context.Context, in *Table, opts ...client.CallOption) (*TableResponse, error)
	Create(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
	Update(ctx context.Context, in *Table, opts ...client.CallOption) (*TableResponse, error)
	Delete(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
	OpenOrClose(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
	// 导出桌台码
	Export(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error)
}

type tableService struct {
	c    client.Client
	name string
}

func NewTableService(name string, c client.Client) TableService {
	return &tableService{
		c:    c,
		name: name,
	}
}

func (c *tableService) Search(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Search", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) List(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.List", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Get(ctx context.Context, in *Table, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Get", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Create(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Create", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Update(ctx context.Context, in *Table, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Update", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Delete(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Delete", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) OpenOrClose(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.OpenOrClose", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Export(ctx context.Context, in *TableWhere, opts ...client.CallOption) (*TableResponse, error) {
	req := c.c.NewRequest(c.name, "TableService.Export", in)
	out := new(TableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TableService service

type TableServiceHandler interface {
	Search(context.Context, *TableWhere, *TableResponse) error
	List(context.Context, *TableWhere, *TableResponse) error
	Get(context.Context, *Table, *TableResponse) error
	Create(context.Context, *TableWhere, *TableResponse) error
	Update(context.Context, *Table, *TableResponse) error
	Delete(context.Context, *TableWhere, *TableResponse) error
	OpenOrClose(context.Context, *TableWhere, *TableResponse) error
	// 导出桌台码
	Export(context.Context, *TableWhere, *TableResponse) error
}

func RegisterTableServiceHandler(s server.Server, hdlr TableServiceHandler, opts ...server.HandlerOption) error {
	type tableService interface {
		Search(ctx context.Context, in *TableWhere, out *TableResponse) error
		List(ctx context.Context, in *TableWhere, out *TableResponse) error
		Get(ctx context.Context, in *Table, out *TableResponse) error
		Create(ctx context.Context, in *TableWhere, out *TableResponse) error
		Update(ctx context.Context, in *Table, out *TableResponse) error
		Delete(ctx context.Context, in *TableWhere, out *TableResponse) error
		OpenOrClose(ctx context.Context, in *TableWhere, out *TableResponse) error
		Export(ctx context.Context, in *TableWhere, out *TableResponse) error
	}
	type TableService struct {
		tableService
	}
	h := &tableServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TableService{h}, opts...))
}

type tableServiceHandler struct {
	TableServiceHandler
}

func (h *tableServiceHandler) Search(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.Search(ctx, in, out)
}

func (h *tableServiceHandler) List(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.List(ctx, in, out)
}

func (h *tableServiceHandler) Get(ctx context.Context, in *Table, out *TableResponse) error {
	return h.TableServiceHandler.Get(ctx, in, out)
}

func (h *tableServiceHandler) Create(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.Create(ctx, in, out)
}

func (h *tableServiceHandler) Update(ctx context.Context, in *Table, out *TableResponse) error {
	return h.TableServiceHandler.Update(ctx, in, out)
}

func (h *tableServiceHandler) Delete(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.Delete(ctx, in, out)
}

func (h *tableServiceHandler) OpenOrClose(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.OpenOrClose(ctx, in, out)
}

func (h *tableServiceHandler) Export(ctx context.Context, in *TableWhere, out *TableResponse) error {
	return h.TableServiceHandler.Export(ctx, in, out)
}
