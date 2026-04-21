//go:build wireinject
// +build wireinject

package lib

import (
	"go-clean-architecture/server/lib/database"
	"go-clean-architecture/server/lib/environment"
	"go-clean-architecture/server/lib/logger"

	"github.com/google/wire"
)

var AppModule = wire.NewSet(
	environment.ProvideConfig,
	database.ProvideSQLDatabase,
	logger.ProvideLogger,
)
