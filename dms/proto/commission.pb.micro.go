// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: commission.proto

package geiqin_srv_dms

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

// Client API for CommissionService service

type CommissionService interface {
	//下单后创建佣金相关数据
	Create(ctx context.Context, in *Commission, opts ...client.CallOption) (*CommissionResponse, error)
	//查看结算详情
	Get(ctx context.Context, in *Commission, opts ...client.CallOption) (*CommissionResponse, error)
	//分页查询结算列表
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CommissionResponse, error)
	//分页查询业绩报表
	SearchPerformance(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CommissionResponse, error)
}

type commissionService struct {
	c    client.Client
	name string
}

func NewCommissionService(name string, c client.Client) CommissionService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geiqin.srv.dms"
	}
	return &commissionService{
		c:    c,
		name: name,
	}
}

func (c *commissionService) Create(ctx context.Context, in *Commission, opts ...client.CallOption) (*CommissionResponse, error) {
	req := c.c.NewRequest(c.name, "CommissionService.Create", in)
	out := new(CommissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commissionService) Get(ctx context.Context, in *Commission, opts ...client.CallOption) (*CommissionResponse, error) {
	req := c.c.NewRequest(c.name, "CommissionService.Get", in)
	out := new(CommissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commissionService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CommissionResponse, error) {
	req := c.c.NewRequest(c.name, "CommissionService.Search", in)
	out := new(CommissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commissionService) SearchPerformance(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*CommissionResponse, error) {
	req := c.c.NewRequest(c.name, "CommissionService.searchPerformance", in)
	out := new(CommissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommissionService service

type CommissionServiceHandler interface {
	//下单后创建佣金相关数据
	Create(context.Context, *Commission, *CommissionResponse) error
	//查看结算详情
	Get(context.Context, *Commission, *CommissionResponse) error
	//分页查询结算列表
	Search(context.Context, *BaseWhere, *CommissionResponse) error
	//分页查询业绩报表
	SearchPerformance(context.Context, *BaseWhere, *CommissionResponse) error
}

func RegisterCommissionServiceHandler(s server.Server, hdlr CommissionServiceHandler, opts ...server.HandlerOption) error {
	type commissionService interface {
		Create(ctx context.Context, in *Commission, out *CommissionResponse) error
		Get(ctx context.Context, in *Commission, out *CommissionResponse) error
		Search(ctx context.Context, in *BaseWhere, out *CommissionResponse) error
		SearchPerformance(ctx context.Context, in *BaseWhere, out *CommissionResponse) error
	}
	type CommissionService struct {
		commissionService
	}
	h := &commissionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CommissionService{h}, opts...))
}

type commissionServiceHandler struct {
	CommissionServiceHandler
}

func (h *commissionServiceHandler) Create(ctx context.Context, in *Commission, out *CommissionResponse) error {
	return h.CommissionServiceHandler.Create(ctx, in, out)
}

func (h *commissionServiceHandler) Get(ctx context.Context, in *Commission, out *CommissionResponse) error {
	return h.CommissionServiceHandler.Get(ctx, in, out)
}

func (h *commissionServiceHandler) Search(ctx context.Context, in *BaseWhere, out *CommissionResponse) error {
	return h.CommissionServiceHandler.Search(ctx, in, out)
}

func (h *commissionServiceHandler) SearchPerformance(ctx context.Context, in *BaseWhere, out *CommissionResponse) error {
	return h.CommissionServiceHandler.SearchPerformance(ctx, in, out)
}
