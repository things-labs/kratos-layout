////go:build swag
//// +build swag

package router

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"

	"github.com/thinkgos/etching/docs"
)

func Swagger(srv *http.Server) {
	swag.Register(swag.Name, new(docs.Docs))
	srv.HandlePrefix("/swagger", httpSwagger.Handler())
}
