// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.7
// source: api/helloworld1/v1/student.proto

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

const OperationStudentGetStudent = "/api.helloworld1.v1.Student/GetStudent"
const OperationStudentHello = "/api.helloworld1.v1.Student/Hello"
const OperationStudentListStudent = "/api.helloworld1.v1.Student/ListStudent"

type StudentHTTPServer interface {
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentReply, error)
	Hello(context.Context, *HelloReq) (*HelloResp, error)
	ListStudent(context.Context, *ListStudentRequest) (*ListStudentReply, error)
}

func RegisterStudentHTTPServer(s *http.Server, srv StudentHTTPServer) {
	r := s.Route("/")
	r.GET("/student/{id}", _Student_GetStudent0_HTTP_Handler(srv))
	r.GET("/students", _Student_ListStudent0_HTTP_Handler(srv))
	r.POST("/hello/{id}/sayhello/{sayname}", _Student_Hello0_HTTP_Handler(srv))
	r.GET("/hello/{name}", _Student_Hello1_HTTP_Handler(srv))
}

func _Student_GetStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentGetStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStudent(ctx, req.(*GetStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_ListStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentListStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListStudent(ctx, req.(*ListStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_Hello0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Hello(ctx, req.(*HelloReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloResp)
		return ctx.Result(200, reply)
	}
}

func _Student_Hello1_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Hello(ctx, req.(*HelloReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloResp)
		return ctx.Result(200, reply)
	}
}

type StudentHTTPClient interface {
	GetStudent(ctx context.Context, req *GetStudentRequest, opts ...http.CallOption) (rsp *GetStudentReply, err error)
	Hello(ctx context.Context, req *HelloReq, opts ...http.CallOption) (rsp *HelloResp, err error)
	ListStudent(ctx context.Context, req *ListStudentRequest, opts ...http.CallOption) (rsp *ListStudentReply, err error)
}

type StudentHTTPClientImpl struct {
	cc *http.Client
}

func NewStudentHTTPClient(client *http.Client) StudentHTTPClient {
	return &StudentHTTPClientImpl{client}
}

func (c *StudentHTTPClientImpl) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...http.CallOption) (*GetStudentReply, error) {
	var out GetStudentReply
	pattern := "/student/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentGetStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) Hello(ctx context.Context, in *HelloReq, opts ...http.CallOption) (*HelloResp, error) {
	var out HelloResp
	pattern := "/hello/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) ListStudent(ctx context.Context, in *ListStudentRequest, opts ...http.CallOption) (*ListStudentReply, error) {
	var out ListStudentReply
	pattern := "/students"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentListStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
