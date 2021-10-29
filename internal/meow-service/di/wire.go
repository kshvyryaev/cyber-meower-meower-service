//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/command"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	controller "github.com/kshvyryaev/cyber-meower/internal/meow-service/controller/http"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/repository"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/service"
)

func InitializeHttpServer(config *config.Config) (*controller.HttpServer, func(), error) {
	panic(wire.Build(
		repository.PostgresConnectionSet,
		repository.PostgresMeowRepositorySet,
		service.MeowTranslatorServiceSet,
		command.СreateMeowCommandHandlerSet,
		controller.MeowControllerSet,
		controller.HttpServerSet,
	))
}
