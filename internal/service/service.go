package service

import (
	"github.com/google/wire"

	"github.com/things-labs/kratos-layout/internal/service/helloworld"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	helloworld.NewGreeterService,
)
