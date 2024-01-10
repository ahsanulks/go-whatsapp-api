// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"app/configs"
	"app/internal/adapter/driven/user_repository"
	"app/internal/adapter/driven/whatsmeow"
	"app/internal/adapter/driver"
	"app/internal/infra"
	"app/internal/usecase/authentication"
	"app/internal/usecase/message"
	"app/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-playground/validator/v10"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(applicationConfig *configs.ApplicationConfig, dbConfig *configs.DBConfig, validate *validator.Validate, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(applicationConfig, logger)
	postgresDB, cleanup := infra.NewPostgresDB(dbConfig, logger)
	whatsmeowClient := whatsmeowclient.NewWhatsmeowClient(postgresDB)
	loginUsecase := authentication.NewLoginUsecase(validate, whatsmeowClient)
	loginHandler := driver.NewLoginHandler(loginUsecase)
	userRepository := userrepository.NewUserRepository(postgresDB)
	messageUsecase := message.NewMessageUsecase(validate, whatsmeowClient, userRepository)
	messageHandler := driver.NewMessageHandler(messageUsecase)
	httpServer := server.NewHTTPServer(applicationConfig, loginHandler, messageHandler, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
