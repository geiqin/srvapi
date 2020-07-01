// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: store.proto

package geiqin_srv_mts

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

// Client API for StoreService service

type StoreService interface {
	//创建店铺
	Create(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//修改店铺
	Update(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//修改店铺Logo
	UpdateLogo(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//修改店铺名称
	UpdateName(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//设置店铺地址
	SetAddress(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//删除店铺
	Delete(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//升级店铺数据库
	Upgrade(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//店铺授权
	Auth(ctx context.Context, in *StoreSecret, opts ...client.CallOption) (*TokenResponse, error)
	//切换店铺
	Switch(ctx context.Context, in *Store, opts ...client.CallOption) (*TokenResponse, error)
	//获取当前店铺信息
	Info(ctx context.Context, in *Empty, opts ...client.CallOption) (*StoreResponse, error)
	//获取店铺信息
	Get(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//锁定店铺
	Lock(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//解锁店铺
	Unlock(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//查询店铺
	Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*StoreResponse, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) Create(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Create", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Update(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Update", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) UpdateLogo(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.UpdateLogo", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) UpdateName(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.UpdateName", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) SetAddress(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.SetAddress", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Delete(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Delete", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Upgrade(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Upgrade", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Auth(ctx context.Context, in *StoreSecret, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Auth", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Switch(ctx context.Context, in *Store, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Switch", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Info(ctx context.Context, in *Empty, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Info", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Get(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Get", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Lock(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Lock", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Unlock(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Unlock", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Search(ctx context.Context, in *BaseWhere, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Search", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreService service

type StoreServiceHandler interface {
	//创建店铺
	Create(context.Context, *Store, *StoreResponse) error
	//修改店铺
	Update(context.Context, *Store, *StoreResponse) error
	//修改店铺Logo
	UpdateLogo(context.Context, *Store, *StoreResponse) error
	//修改店铺名称
	UpdateName(context.Context, *Store, *StoreResponse) error
	//设置店铺地址
	SetAddress(context.Context, *Store, *StoreResponse) error
	//删除店铺
	Delete(context.Context, *Store, *StoreResponse) error
	//升级店铺数据库
	Upgrade(context.Context, *Store, *StoreResponse) error
	//店铺授权
	Auth(context.Context, *StoreSecret, *TokenResponse) error
	//切换店铺
	Switch(context.Context, *Store, *TokenResponse) error
	//获取当前店铺信息
	Info(context.Context, *Empty, *StoreResponse) error
	//获取店铺信息
	Get(context.Context, *Store, *StoreResponse) error
	//锁定店铺
	Lock(context.Context, *Store, *StoreResponse) error
	//解锁店铺
	Unlock(context.Context, *Store, *StoreResponse) error
	//查询店铺
	Search(context.Context, *BaseWhere, *StoreResponse) error
}

func RegisterStoreServiceHandler(s server.Server, hdlr StoreServiceHandler, opts ...server.HandlerOption) error {
	type storeService interface {
		Create(ctx context.Context, in *Store, out *StoreResponse) error
		Update(ctx context.Context, in *Store, out *StoreResponse) error
		UpdateLogo(ctx context.Context, in *Store, out *StoreResponse) error
		UpdateName(ctx context.Context, in *Store, out *StoreResponse) error
		SetAddress(ctx context.Context, in *Store, out *StoreResponse) error
		Delete(ctx context.Context, in *Store, out *StoreResponse) error
		Upgrade(ctx context.Context, in *Store, out *StoreResponse) error
		Auth(ctx context.Context, in *StoreSecret, out *TokenResponse) error
		Switch(ctx context.Context, in *Store, out *TokenResponse) error
		Info(ctx context.Context, in *Empty, out *StoreResponse) error
		Get(ctx context.Context, in *Store, out *StoreResponse) error
		Lock(ctx context.Context, in *Store, out *StoreResponse) error
		Unlock(ctx context.Context, in *Store, out *StoreResponse) error
		Search(ctx context.Context, in *BaseWhere, out *StoreResponse) error
	}
	type StoreService struct {
		storeService
	}
	h := &storeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StoreService{h}, opts...))
}

type storeServiceHandler struct {
	StoreServiceHandler
}

func (h *storeServiceHandler) Create(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Create(ctx, in, out)
}

func (h *storeServiceHandler) Update(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Update(ctx, in, out)
}

func (h *storeServiceHandler) UpdateLogo(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.UpdateLogo(ctx, in, out)
}

func (h *storeServiceHandler) UpdateName(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.UpdateName(ctx, in, out)
}

func (h *storeServiceHandler) SetAddress(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.SetAddress(ctx, in, out)
}

func (h *storeServiceHandler) Delete(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Delete(ctx, in, out)
}

func (h *storeServiceHandler) Upgrade(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Upgrade(ctx, in, out)
}

func (h *storeServiceHandler) Auth(ctx context.Context, in *StoreSecret, out *TokenResponse) error {
	return h.StoreServiceHandler.Auth(ctx, in, out)
}

func (h *storeServiceHandler) Switch(ctx context.Context, in *Store, out *TokenResponse) error {
	return h.StoreServiceHandler.Switch(ctx, in, out)
}

func (h *storeServiceHandler) Info(ctx context.Context, in *Empty, out *StoreResponse) error {
	return h.StoreServiceHandler.Info(ctx, in, out)
}

func (h *storeServiceHandler) Get(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Get(ctx, in, out)
}

func (h *storeServiceHandler) Lock(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Lock(ctx, in, out)
}

func (h *storeServiceHandler) Unlock(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Unlock(ctx, in, out)
}

func (h *storeServiceHandler) Search(ctx context.Context, in *BaseWhere, out *StoreResponse) error {
	return h.StoreServiceHandler.Search(ctx, in, out)
}
