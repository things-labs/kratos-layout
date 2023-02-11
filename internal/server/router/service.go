package router

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/thinkgos/kratos-layout/internal/service/helloworld"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(wire.Struct(new(DependencyService), "*"))

type DependencyService struct {
	log.Logger
	*helloworld.GreeterService
}
