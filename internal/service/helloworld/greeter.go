package helloworld

import (
	"context"

	v1 "github.com/thinkgos/etching/genproto/helloworld/v1"
	bizHelloworld "github.com/thinkgos/etching/internal/biz/helloworld"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *bizHelloworld.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *bizHelloworld.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &bizHelloworld.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello, Number: 1111}, nil
}
