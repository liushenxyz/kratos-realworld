//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"realworld/internal/biz"
	"realworld/internal/conf"
	"realworld/internal/data"
	"realworld/internal/server"
	"realworld/internal/service"
)

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confAuth *conf.Auth, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
