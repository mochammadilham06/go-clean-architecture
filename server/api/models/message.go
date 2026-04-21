package models

//just example for implement message publisher, you can modify this struct as you need
type NotificationPayload struct {
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	EventType string `json:"event_type"`
	Message   string `json:"message"`
}
