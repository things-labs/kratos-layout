//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/things-labs/kratos-layout/internal/biz"
	"github.com/things-labs/kratos-layout/internal/config"
	"github.com/things-labs/kratos-layout/internal/data"
	"github.com/things-labs/kratos-layout/internal/server"
	"github.com/things-labs/kratos-layout/internal/server/router"
	"github.com/things-labs/kratos-layout/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*config.Server, *config.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		newApp,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		router.ProviderSet,
	))
}
