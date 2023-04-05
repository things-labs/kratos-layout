package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	bizHelloworld "github.com/things-labs/kratos-layout/internal/biz/helloworld"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) bizHelloworld.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *bizHelloworld.Greeter) (*bizHelloworld.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *bizHelloworld.Greeter) (*bizHelloworld.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*bizHelloworld.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*bizHelloworld.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*bizHelloworld.Greeter, error) {
	return nil, nil
}
