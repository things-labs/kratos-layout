// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/things-labs/kratos-layout/internal/biz/helloworld"
	"github.com/things-labs/kratos-layout/internal/config"
	"github.com/things-labs/kratos-layout/internal/data"
	"github.com/things-labs/kratos-layout/internal/server"
	"github.com/things-labs/kratos-layout/internal/server/router"
	helloworld2 "github.com/things-labs/kratos-layout/internal/service/helloworld"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(configServer *config.Server, configData *config.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(configData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := helloworld.NewGreeterUsecase(greeterRepo, logger)
	greeterService := helloworld2.NewGreeterService(greeterUsecase)
	dependencyService := &router.DependencyService{
		Logger:         logger,
		GreeterService: greeterService,
	}
	grpcServer := server.NewGRPCServer(configServer, dependencyService)
	httpServer := server.NewHTTPServer(configServer, dependencyService)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
