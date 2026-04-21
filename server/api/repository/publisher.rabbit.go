package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pkg "go-clean-architecture/pkg/rabbitmq"
	"go-clean-architecture/server/api/contract"
	"go-clean-architecture/server/api/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQPublisher struct {
	client   *pkg.RabbitMQClient
	exchange string
}

// constructor
func NewRabbitMQPublisher(client *pkg.RabbitMQClient, exchangeName string) contract.MessagePublisher {
	err := client.Channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable (bertahan meskipun rabbitmq restart)
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange %s: %v", exchangeName, err)
	}

	return &rabbitMQPublisher{
		client:   client,
		exchange: exchangeName,
	}
}

// wrapper method for publish message to rabbitmq, you can modify this method as you need
func NewRabbitMQPublisherFromConfig(client *pkg.RabbitMQClient) contract.MessagePublisher {
	return NewRabbitMQPublisher(client, "notification_exchange")
}

// PublishNotification is a method to publish notification messages to RabbitMQ, you can modify this method as you need
func (r *rabbitMQPublisher) PublishNotification(ctx context.Context, payload models.NotificationPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal notification payload: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	routingKey := "notification.email"
	err = r.client.Channel.PublishWithContext(
		ctx,
		r.exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		log.Printf("Failed to publish message to RabbitMQ: %v", err)
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Successfully published event %s for user %d to RabbitMQ", payload.EventType, payload.UserID)
	return nil
}
