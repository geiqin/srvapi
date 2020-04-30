// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: team.proto

package geiqin_srv_dms

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

// Client API for MyTeamService service

type MyTeamService interface {
	//查询我的团队成员
	Search(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*MyTeamResponse, error)
}

type myTeamService struct {
	c    client.Client
	name string
}

func NewMyTeamService(name string, c client.Client) MyTeamService {
	return &myTeamService{
		c:    c,
		name: name,
	}
}

func (c *myTeamService) Search(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*MyTeamResponse, error) {
	req := c.c.NewRequest(c.name, "MyTeamService.Search", in)
	out := new(MyTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyTeamService service

type MyTeamServiceHandler interface {
	//查询我的团队成员
	Search(context.Context, *TeamWhere, *MyTeamResponse) error
}

func RegisterMyTeamServiceHandler(s server.Server, hdlr MyTeamServiceHandler, opts ...server.HandlerOption) error {
	type myTeamService interface {
		Search(ctx context.Context, in *TeamWhere, out *MyTeamResponse) error
	}
	type MyTeamService struct {
		myTeamService
	}
	h := &myTeamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyTeamService{h}, opts...))
}

type myTeamServiceHandler struct {
	MyTeamServiceHandler
}

func (h *myTeamServiceHandler) Search(ctx context.Context, in *TeamWhere, out *MyTeamResponse) error {
	return h.MyTeamServiceHandler.Search(ctx, in, out)
}

// Client API for TeamService service

type TeamService interface {
	//查询团队列表
	Search(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*TeamResponse, error)
	//查询团队成员
	MembersSearch(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*MyTeamResponse, error)
}

type teamService struct {
	c    client.Client
	name string
}

func NewTeamService(name string, c client.Client) TeamService {
	return &teamService{
		c:    c,
		name: name,
	}
}

func (c *teamService) Search(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*TeamResponse, error) {
	req := c.c.NewRequest(c.name, "TeamService.Search", in)
	out := new(TeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) MembersSearch(ctx context.Context, in *TeamWhere, opts ...client.CallOption) (*MyTeamResponse, error) {
	req := c.c.NewRequest(c.name, "TeamService.MembersSearch", in)
	out := new(MyTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TeamService service

type TeamServiceHandler interface {
	//查询团队列表
	Search(context.Context, *TeamWhere, *TeamResponse) error
	//查询团队成员
	MembersSearch(context.Context, *TeamWhere, *MyTeamResponse) error
}

func RegisterTeamServiceHandler(s server.Server, hdlr TeamServiceHandler, opts ...server.HandlerOption) error {
	type teamService interface {
		Search(ctx context.Context, in *TeamWhere, out *TeamResponse) error
		MembersSearch(ctx context.Context, in *TeamWhere, out *MyTeamResponse) error
	}
	type TeamService struct {
		teamService
	}
	h := &teamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TeamService{h}, opts...))
}

type teamServiceHandler struct {
	TeamServiceHandler
}

func (h *teamServiceHandler) Search(ctx context.Context, in *TeamWhere, out *TeamResponse) error {
	return h.TeamServiceHandler.Search(ctx, in, out)
}

func (h *teamServiceHandler) MembersSearch(ctx context.Context, in *TeamWhere, out *MyTeamResponse) error {
	return h.TeamServiceHandler.MembersSearch(ctx, in, out)
}