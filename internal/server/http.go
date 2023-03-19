package server

import (
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/swagger-api/openapiv2"
	v1 "helloworld/api/helloworld/v1"
	v2 "helloworld/api/helloworld1/v1"
	p "helloworld/api/pledge/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, student *service.StudentService, pledge *service.PledgeService, animal *service.AnimalService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(), // 链路追踪
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	v2.RegisterStudentHTTPServer(srv, student)
	p.RegisterPledgeHTTPServer(srv, pledge)
	p.RegisterAnimalHTTPServer(srv, animal)
	return srv
}
