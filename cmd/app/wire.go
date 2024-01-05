//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"app/configs"
	"app/internal/adapter/driven"
	"app/internal/adapter/driver"
	"app/internal/infra"
	portdriven "app/internal/port/driven"
	portdriver "app/internal/port/driver"
	"app/internal/service"
	"app/internal/usecase"
	"app/internal/usecase/authentication"
	"app/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*configs.ApplicationConfig, *configs.DBConfig, *validator.Validate, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			infra.ProviderSet,
			usecase.ProviderSet,
			service.ProviderSet,
			driven.ProviderSet,
			driver.ProviderSet,
			newApp,
			wire.Bind(new(usecase.GreeterRepo), new(*infra.GreeterRepo)),
			wire.Bind(new(portdriven.DeviceProvider), new(*driven.WhatsmewClient)),
			wire.Bind(new(portdriver.AuthenticationUsecase), new(*authentication.LoginUsecase)),
		),
	)
}
