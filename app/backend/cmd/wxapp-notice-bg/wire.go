// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/conf"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/data"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/server"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.WxAppConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
