//go:build tools
// +build tools

package tools

import (
	_ "github.com/google/wire/cmd/wire"
	_ "golang.org/x/tools/cmd/stringer"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/go-kratos/kratos/cmd/kratos/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2"
)

// other need tools
// github.com/upx/upx
