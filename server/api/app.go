package api

import (
	pkg "go-clean-architecture/pkg/rabbitmq"
	"go-clean-architecture/server/api/contract"
	"go-clean-architecture/server/api/handler"
)

type App struct {
	Handler   *handler.Handler
	RabbitMQ  *pkg.RabbitMQClient
	Publisher contract.MessagePublisher
}

func NewApp(h *handler.Handler, rmq *pkg.RabbitMQClient, pub contract.MessagePublisher) *App {
	return &App{
		Handler:   h,
		RabbitMQ:  rmq,
		Publisher: pub,
	}
}
