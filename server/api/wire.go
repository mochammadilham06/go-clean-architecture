//go:build wireinject
// +build wireinject

package api

import (
	pkg "go-clean-architecture/pkg/rabbitmq"
	"go-clean-architecture/server/api/contract"
	"go-clean-architecture/server/api/handler"
	"go-clean-architecture/server/api/repository"
	"go-clean-architecture/server/api/service"
	"go-clean-architecture/server/lib/database"
	"go-clean-architecture/server/lib/environment"
	"go-clean-architecture/server/lib/logger"

	"github.com/google/wire"
)

var ProjectSet = wire.NewSet(
	repository.NewProjectRepository,
	wire.Bind(new(contract.IProjectRepository), new(*repository.ProjectRepository)),
	service.NewProjectService,
	wire.Bind(new(contract.IProjectService), new(*service.ProjectService)),
)

var ExperienceSet = wire.NewSet(
	repository.NewExperienceRepository,
	wire.Bind(new(contract.IExperienceRepository), new(*repository.ExperienceRepository)),
	service.NewExperienceService,
	wire.Bind(new(contract.IExperienceService), new(*service.ExperienceService)),
)

var RabbitMQSet = wire.NewSet(
	pkg.NewRabbitMQConnectionFromConfig,
	repository.NewRabbitMQPublisherFromConfig,
)

func InitializeAPI(cfg *environment.Config, log *logger.Logger) (*App, error) {
	wire.Build(
		database.ProvideSQLDatabase,
		ProjectSet,
		ExperienceSet,
		handler.NewHandler,
		RabbitMQSet,

		NewApp,
	)
	return &App{}, nil
}
