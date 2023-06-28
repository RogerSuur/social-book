package websocket

import (
	"SocialNetworkRestApi/api/pkg/services"
	"encoding/json"
)

type Payload struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type RequestPayload struct {
	ID       int  `json:"id"`
	Reaction bool `json:"reaction"`
	GroupID  int  `json:"group_id"`
}

type NotificationPayload struct {
	NotificationType string `json:"notification_type"`
	NotificationID   int    `json:"notification_id"`
	SenderID         int    `json:"sender_id"`
	SenderName       string `json:"sender_name"`
	GroupID          int    `json:"group_id"`
	GroupName        string `json:"group_name"`
	EventID          int    `json:"event_id"`
	EventName        string `json:"event_name"`
}

type MessagePayload struct {
	MessageID     int    `json:"message_id"`
	SenderID      int    `json:"sender_id"`
	SenderName    string `json:"sender_name"`
	RecipientID   int    `json:"recipient_id"`
	RecipientName string `json:"recipient_name"`
	GroupID       int    `json:"group_id"`
	GroupName     string `json:"group_name"`
	MessageBody   string `json:"body"`
	Timestamp     string `json:"timestamp"`
}

type ChatListPayload struct {
	UserID   int                     `json:"user_id"`
	Chatlist []services.ChatListUser `json:"chatlist"`
}
