//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/thinkgos/etching/internal/biz"
	"github.com/thinkgos/etching/internal/config"
	"github.com/thinkgos/etching/internal/data"
	"github.com/thinkgos/etching/internal/server"
	"github.com/thinkgos/etching/internal/server/router"
	"github.com/thinkgos/etching/internal/service"

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
