package biz

import (
	"github.com/google/wire"
	"github.com/thinkgos/etching/internal/biz/helloworld"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(helloworld.NewGreeterUsecase)
