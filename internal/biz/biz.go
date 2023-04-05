package biz

import (
	"github.com/google/wire"
	"github.com/things-labs/kratos-layout/internal/biz/helloworld"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(helloworld.NewGreeterUsecase)
