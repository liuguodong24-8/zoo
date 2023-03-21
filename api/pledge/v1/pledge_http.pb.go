// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.7
// source: api/pledge/v1/pledge.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPledgeCreatePledge = "/api.pledge.v1.Pledge/CreatePledge"
const OperationPledgeGetPledge = "/api.pledge.v1.Pledge/GetPledge"
const OperationPledgeUpdatePledgeStatus = "/api.pledge.v1.Pledge/UpdatePledgeStatus"

type PledgeHTTPServer interface {
	CreatePledge(context.Context, *CreatePledgeRequest) (*CreatePledgeReply, error)
	GetPledge(context.Context, *GetPledgeRequest) (*GetPledgeReply, error)
	UpdatePledgeStatus(context.Context, *UpdatePledgeStatusRequest) (*UpdatePledgeStatusReply, error)
}

func RegisterPledgeHTTPServer(s *http.Server, srv PledgeHTTPServer) {
	r := s.Route("/")
	r.POST("/pledge", _Pledge_CreatePledge0_HTTP_Handler(srv))
	r.POST("/pledge/status", _Pledge_UpdatePledgeStatus0_HTTP_Handler(srv))
	r.GET("/pledge/{address}", _Pledge_GetPledge0_HTTP_Handler(srv))
}

func _Pledge_CreatePledge0_HTTP_Handler(srv PledgeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePledgeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPledgeCreatePledge)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePledge(ctx, req.(*CreatePledgeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePledgeReply)
		return ctx.Result(200, reply)
	}
}

func _Pledge_UpdatePledgeStatus0_HTTP_Handler(srv PledgeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePledgeStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPledgeUpdatePledgeStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePledgeStatus(ctx, req.(*UpdatePledgeStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePledgeStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Pledge_GetPledge0_HTTP_Handler(srv PledgeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPledgeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPledgeGetPledge)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPledge(ctx, req.(*GetPledgeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPledgeReply)
		return ctx.Result(200, reply)
	}
}

type PledgeHTTPClient interface {
	CreatePledge(ctx context.Context, req *CreatePledgeRequest, opts ...http.CallOption) (rsp *CreatePledgeReply, err error)
	GetPledge(ctx context.Context, req *GetPledgeRequest, opts ...http.CallOption) (rsp *GetPledgeReply, err error)
	UpdatePledgeStatus(ctx context.Context, req *UpdatePledgeStatusRequest, opts ...http.CallOption) (rsp *UpdatePledgeStatusReply, err error)
}

type PledgeHTTPClientImpl struct {
	cc *http.Client
}

func NewPledgeHTTPClient(client *http.Client) PledgeHTTPClient {
	return &PledgeHTTPClientImpl{client}
}

func (c *PledgeHTTPClientImpl) CreatePledge(ctx context.Context, in *CreatePledgeRequest, opts ...http.CallOption) (*CreatePledgeReply, error) {
	var out CreatePledgeReply
	pattern := "/pledge"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPledgeCreatePledge))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PledgeHTTPClientImpl) GetPledge(ctx context.Context, in *GetPledgeRequest, opts ...http.CallOption) (*GetPledgeReply, error) {
	var out GetPledgeReply
	pattern := "/pledge/{address}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPledgeGetPledge))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PledgeHTTPClientImpl) UpdatePledgeStatus(ctx context.Context, in *UpdatePledgeStatusRequest, opts ...http.CallOption) (*UpdatePledgeStatusReply, error) {
	var out UpdatePledgeStatusReply
	pattern := "/pledge/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPledgeUpdatePledgeStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
