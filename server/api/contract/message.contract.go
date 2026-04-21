package contract

import (
	"context"
	"go-clean-architecture/server/api/models"
)

// example contract for publisher, you can modify this contract as you need
type MessagePublisher interface {
	PublishNotification(ctx context.Context, payload models.NotificationPayload) error
}
