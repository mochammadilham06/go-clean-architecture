package pkg

import (
	"fmt"
	"go-clean-architecture/server/lib/environment"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConnection(amqpURL string) (*RabbitMQClient, error) {
	var conn *amqp.Connection
	var err error

	maxRetries := 5
	for i := 1; i <= maxRetries; i++ {
		log.Printf("Attempting to connect to RabbitMQ (Attempt %d/%d)...", i, maxRetries)
		conn, err = amqp.Dial(amqpURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ: %v. Retrying in 2 seconds...", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ after %d attempts: %w", maxRetries, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	log.Println("Successfully connected to RabbitMQ")

	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func NewRabbitMQConnectionFromConfig(cfg *environment.Config) (*RabbitMQClient, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.QUEUE_USER, cfg.QUEUE_PASSWORD, cfg.QUEUE_HOST, cfg.QUEUE_PORT)
	return NewRabbitMQConnection(url)
}

func (r *RabbitMQClient) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
	log.Println("RabbitMQ connection closed gracefully")
}
