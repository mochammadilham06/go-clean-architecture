//go:build wireinject
// +build wireinject

package api

import (
	"go-portfolio/server/api/contract"
	"go-portfolio/server/api/handler"
	"go-portfolio/server/api/repository"
	"go-portfolio/server/api/service"
	"go-portfolio/server/lib/database"
	"go-portfolio/server/lib/environment"
	"go-portfolio/server/lib/logger"

	"github.com/google/wire"
)

// Grouping services
var ProjectSet = wire.NewSet(

	//1.project services
	repository.NewProjectRepository,
	wire.Bind(new(contract.IProjectRepository), new(*repository.ProjectRepository)),

	service.NewProjectService,
	wire.Bind(new(contract.IProjectService), new(*service.ProjectService)),
)

// 2.experience services
var ExperienceSet = wire.NewSet(
	repository.NewExperienceRepository,
	wire.Bind(new(contract.IExperienceRepository), new(*repository.ExperienceRepository)),

	service.NewExperienceService,
	wire.Bind(new(contract.IExperienceService), new(*service.ExperienceService)),
)

// 3. Injector
func InitializeAPI(cfg *environment.Config, log *logger.Logger) (*handler.Handler, error) {
	wire.Build(
		database.ProvideSQLDatabase,

		// Domain sets
		ProjectSet,
		ExperienceSet,

		//handler
		handler.NewHandler,
	)
	return &handler.Handler{}, nil
}
