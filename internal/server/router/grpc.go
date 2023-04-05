package router

import (
	v1 "github.com/things-labs/kratos-layout/genproto/helloworld/v1"
	"google.golang.org/grpc"
)

func RegisterGrpc(srv grpc.ServiceRegistrar, d *DependencyService) {
	v1.RegisterGreeterServer(srv, d.GreeterService)
}
