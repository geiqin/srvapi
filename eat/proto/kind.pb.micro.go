// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: kind.proto

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

// Client API for InKindService service

type InKindService interface {
	Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Create(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Update(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Delete(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	// 菜单拖动排序
	Resort(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
}

type inKindService struct {
	c    client.Client
	name string
}

func NewInKindService(name string, c client.Client) InKindService {
	return &inKindService{
		c:    c,
		name: name,
	}
}

func (c *inKindService) Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Search", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.List", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Get", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) Create(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Create", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) Update(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Update", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) Delete(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Delete", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inKindService) Resort(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "InKindService.Resort", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InKindService service

type InKindServiceHandler interface {
	Search(context.Context, *KindWhere, *KindResponse) error
	List(context.Context, *KindWhere, *KindResponse) error
	Get(context.Context, *Kind, *KindResponse) error
	Create(context.Context, *Kind, *KindResponse) error
	Update(context.Context, *Kind, *KindResponse) error
	Delete(context.Context, *KindWhere, *KindResponse) error
	// 菜单拖动排序
	Resort(context.Context, *KindWhere, *KindResponse) error
}

func RegisterInKindServiceHandler(s server.Server, hdlr InKindServiceHandler, opts ...server.HandlerOption) error {
	type inKindService interface {
		Search(ctx context.Context, in *KindWhere, out *KindResponse) error
		List(ctx context.Context, in *KindWhere, out *KindResponse) error
		Get(ctx context.Context, in *Kind, out *KindResponse) error
		Create(ctx context.Context, in *Kind, out *KindResponse) error
		Update(ctx context.Context, in *Kind, out *KindResponse) error
		Delete(ctx context.Context, in *KindWhere, out *KindResponse) error
		Resort(ctx context.Context, in *KindWhere, out *KindResponse) error
	}
	type InKindService struct {
		inKindService
	}
	h := &inKindServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&InKindService{h}, opts...))
}

type inKindServiceHandler struct {
	InKindServiceHandler
}

func (h *inKindServiceHandler) Search(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.InKindServiceHandler.Search(ctx, in, out)
}

func (h *inKindServiceHandler) List(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.InKindServiceHandler.List(ctx, in, out)
}

func (h *inKindServiceHandler) Get(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.InKindServiceHandler.Get(ctx, in, out)
}

func (h *inKindServiceHandler) Create(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.InKindServiceHandler.Create(ctx, in, out)
}

func (h *inKindServiceHandler) Update(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.InKindServiceHandler.Update(ctx, in, out)
}

func (h *inKindServiceHandler) Delete(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.InKindServiceHandler.Delete(ctx, in, out)
}

func (h *inKindServiceHandler) Resort(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.InKindServiceHandler.Resort(ctx, in, out)
}

// Client API for MyInKindService service

type MyInKindService interface {
	Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
}

type myInKindService struct {
	c    client.Client
	name string
}

func NewMyInKindService(name string, c client.Client) MyInKindService {
	return &myInKindService{
		c:    c,
		name: name,
	}
}

func (c *myInKindService) Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyInKindService.Search", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myInKindService) List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyInKindService.List", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myInKindService) Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyInKindService.Get", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyInKindService service

type MyInKindServiceHandler interface {
	Search(context.Context, *KindWhere, *KindResponse) error
	List(context.Context, *KindWhere, *KindResponse) error
	Get(context.Context, *Kind, *KindResponse) error
}

func RegisterMyInKindServiceHandler(s server.Server, hdlr MyInKindServiceHandler, opts ...server.HandlerOption) error {
	type myInKindService interface {
		Search(ctx context.Context, in *KindWhere, out *KindResponse) error
		List(ctx context.Context, in *KindWhere, out *KindResponse) error
		Get(ctx context.Context, in *Kind, out *KindResponse) error
	}
	type MyInKindService struct {
		myInKindService
	}
	h := &myInKindServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyInKindService{h}, opts...))
}

type myInKindServiceHandler struct {
	MyInKindServiceHandler
}

func (h *myInKindServiceHandler) Search(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.MyInKindServiceHandler.Search(ctx, in, out)
}

func (h *myInKindServiceHandler) List(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.MyInKindServiceHandler.List(ctx, in, out)
}

func (h *myInKindServiceHandler) Get(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.MyInKindServiceHandler.Get(ctx, in, out)
}

// Client API for OutKindService service

type OutKindService interface {
	Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Create(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Update(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
	Delete(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	// 拖动菜单排序
	Resort(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
}

type outKindService struct {
	c    client.Client
	name string
}

func NewOutKindService(name string, c client.Client) OutKindService {
	return &outKindService{
		c:    c,
		name: name,
	}
}

func (c *outKindService) Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Search", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.List", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Get", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) Create(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Create", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) Update(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Update", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) Delete(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Delete", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outKindService) Resort(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "OutKindService.Resort", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OutKindService service

type OutKindServiceHandler interface {
	Search(context.Context, *KindWhere, *KindResponse) error
	List(context.Context, *KindWhere, *KindResponse) error
	Get(context.Context, *Kind, *KindResponse) error
	Create(context.Context, *Kind, *KindResponse) error
	Update(context.Context, *Kind, *KindResponse) error
	Delete(context.Context, *KindWhere, *KindResponse) error
	// 拖动菜单排序
	Resort(context.Context, *KindWhere, *KindResponse) error
}

func RegisterOutKindServiceHandler(s server.Server, hdlr OutKindServiceHandler, opts ...server.HandlerOption) error {
	type outKindService interface {
		Search(ctx context.Context, in *KindWhere, out *KindResponse) error
		List(ctx context.Context, in *KindWhere, out *KindResponse) error
		Get(ctx context.Context, in *Kind, out *KindResponse) error
		Create(ctx context.Context, in *Kind, out *KindResponse) error
		Update(ctx context.Context, in *Kind, out *KindResponse) error
		Delete(ctx context.Context, in *KindWhere, out *KindResponse) error
		Resort(ctx context.Context, in *KindWhere, out *KindResponse) error
	}
	type OutKindService struct {
		outKindService
	}
	h := &outKindServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OutKindService{h}, opts...))
}

type outKindServiceHandler struct {
	OutKindServiceHandler
}

func (h *outKindServiceHandler) Search(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.OutKindServiceHandler.Search(ctx, in, out)
}

func (h *outKindServiceHandler) List(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.OutKindServiceHandler.List(ctx, in, out)
}

func (h *outKindServiceHandler) Get(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.OutKindServiceHandler.Get(ctx, in, out)
}

func (h *outKindServiceHandler) Create(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.OutKindServiceHandler.Create(ctx, in, out)
}

func (h *outKindServiceHandler) Update(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.OutKindServiceHandler.Update(ctx, in, out)
}

func (h *outKindServiceHandler) Delete(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.OutKindServiceHandler.Delete(ctx, in, out)
}

func (h *outKindServiceHandler) Resort(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.OutKindServiceHandler.Resort(ctx, in, out)
}

// Client API for MyOutKindService service

type MyOutKindService interface {
	Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error)
	Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error)
}

type myOutKindService struct {
	c    client.Client
	name string
}

func NewMyOutKindService(name string, c client.Client) MyOutKindService {
	return &myOutKindService{
		c:    c,
		name: name,
	}
}

func (c *myOutKindService) Search(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyOutKindService.Search", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myOutKindService) List(ctx context.Context, in *KindWhere, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyOutKindService.List", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myOutKindService) Get(ctx context.Context, in *Kind, opts ...client.CallOption) (*KindResponse, error) {
	req := c.c.NewRequest(c.name, "MyOutKindService.Get", in)
	out := new(KindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyOutKindService service

type MyOutKindServiceHandler interface {
	Search(context.Context, *KindWhere, *KindResponse) error
	List(context.Context, *KindWhere, *KindResponse) error
	Get(context.Context, *Kind, *KindResponse) error
}

func RegisterMyOutKindServiceHandler(s server.Server, hdlr MyOutKindServiceHandler, opts ...server.HandlerOption) error {
	type myOutKindService interface {
		Search(ctx context.Context, in *KindWhere, out *KindResponse) error
		List(ctx context.Context, in *KindWhere, out *KindResponse) error
		Get(ctx context.Context, in *Kind, out *KindResponse) error
	}
	type MyOutKindService struct {
		myOutKindService
	}
	h := &myOutKindServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyOutKindService{h}, opts...))
}

type myOutKindServiceHandler struct {
	MyOutKindServiceHandler
}

func (h *myOutKindServiceHandler) Search(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.MyOutKindServiceHandler.Search(ctx, in, out)
}

func (h *myOutKindServiceHandler) List(ctx context.Context, in *KindWhere, out *KindResponse) error {
	return h.MyOutKindServiceHandler.List(ctx, in, out)
}

func (h *myOutKindServiceHandler) Get(ctx context.Context, in *Kind, out *KindResponse) error {
	return h.MyOutKindServiceHandler.Get(ctx, in, out)
}