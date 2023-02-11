package router

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/thinkgos/kratos-layout/genproto/helloworld/v1"
)

func RegisterHttpApiV1(srv *http.Server, d *DependencyService) {
	Swagger(srv)

	v1.RegisterGreeterHTTPServer(srv, d.GreeterService)
}
