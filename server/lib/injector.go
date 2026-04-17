//go:build wireinject
// +build wireinject

package lib

import (
	"go-portfolio/server/lib/database"
	"go-portfolio/server/lib/environment"
	"go-portfolio/server/lib/logger"

	"github.com/google/wire"
)

var AppModule = wire.NewSet(
	environment.ProvideConfig,
	database.ProvideSQLDatabase,
	logger.ProvideLogger,
)
