package websocket

import "encoding/json"

type Payload struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type RequestPayload struct {
	ID       int  `json:"id"`
	Reaction bool `json:"reaction"`
	GroupID  int  `json:"group_id"`
}

type MessagePayload struct {
	MessageID     int    `json:"message_id"`
	SenderID      int    `json:"sender_id"`
	RecipientID   int    `json:"recipient_id"`
	RecipientName string `json:"recipient_name"`
	GroupID       int    `json:"group_id"`
	GroupName     string `json:"group_name"`
	MessageBody   string `json:"body"`
	Timestamp     string `json:"timestamp"`
}
