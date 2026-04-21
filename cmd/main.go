package main

import (
	"fmt"

	"go-clean-architecture/docs"
	"go-clean-architecture/server/api"
	"go-clean-architecture/server/lib/environment"
	"go-clean-architecture/server/lib/logger"
	"go-clean-architecture/server/lib/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"     // Dependency Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // Dependency Swagger

	_ "go-clean-architecture/docs" // Swagger Docs
)

func main() {
	cfg, err := environment.ProvideConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to Load Configuration: %v", err))
	}

	appLogger := logger.ProvideLogger(cfg.APP_ENV, cfg.APP_NAME)
	defer func() {
		_ = appLogger.Sync()
	}()

	app, err := api.InitializeAPI(cfg, appLogger)
	if err != nil {
		appLogger.Fatal("Failed to initialize application (Wire)", zap.Error(err))
	}

	// 2. Setup Graceful Shutdown untuk RabbitMQ
	defer func() {
		appLogger.Info("Closing RabbitMQ connection...")
		if app.RabbitMQ != nil {
			app.RabbitMQ.Close()
		}
	}()

	if cfg.APP_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//initial gin
	r := gin.New()

	// use gin logger and recovery
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//use custom middleware cors and rate limiter
	r.Use(middleware.CorsMiddleware(cfg))
	r.Use(middleware.RateLimitMiddleware(cfg))

	//register gin and swagger
	app.Handler.Register(r)
	registerSwagger(r, cfg, appLogger)

	//run app
	appPort := fmt.Sprintf(":%d", cfg.APP_HTTP_PORT)
	appLogger.Info("Running Server on...", zap.String("port", appPort))

	if err := r.Run(appPort); err != nil {
		appLogger.Fatal("Server crash", zap.Error(err))
	}

}

func registerSwagger(r *gin.Engine, cfg *environment.Config, logger *logger.Logger) {

	if cfg.APP_ENV == "production" {
		logger.Info("Disable swagger for production")
		return
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.APP_HTTP_PORT)
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
