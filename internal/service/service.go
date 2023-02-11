package service

import (
	"github.com/google/wire"

	"github.com/thinkgos/etching/internal/service/helloworld"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	helloworld.NewGreeterService,
)
