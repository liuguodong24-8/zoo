package server

import (
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	v1 "helloworld/api/helloworld/v1"
	v2 "helloworld/api/helloworld1/v1"
	p "helloworld/api/pledge/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, student *service.StudentService, pledge *service.PledgeService, animal *service.AnimalService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(), // 链路追踪
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	v2.RegisterStudentServer(srv, student)
	p.RegisterPledgeServer(srv, pledge)
	p.RegisterAnimalServer(srv, animal)
	return srv
}
