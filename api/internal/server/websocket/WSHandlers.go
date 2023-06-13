package websocket

import (
	"encoding/json"
	"errors"
)

type Payload struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type PayloadHandler func(payload Payload, client *Client) error

var (
	ErrPayloadTypeNotSupported = errors.New("this payload type is not supported")
)

const (
	FollowRequest = "follow_request"
	FollowAccept  = "follow_accept"
	NewMessage    = "new_message"
	GroupInvite   = "group_invite"
	GroupAccept   = "group_accept"
)

type SendMessagePayload struct {
	RecipientID int    `json:"recipient_id"`
	Message     string `json:"message"`
}

func (w *WebsocketServer) setupHandlers() {
	w.handlers[FollowRequest] = w.FollowRequestHandler
	w.handlers[FollowAccept] = w.FollowAcceptHandler
	w.handlers[NewMessage] = w.NewMessageHandler
	w.handlers[GroupInvite] = w.GroupInviteHandler
	w.handlers[GroupAccept] = w.GroupAcceptHandler
}

func (w *WebsocketServer) routePayloads(payload Payload, client *Client) error {
	handler, ok := w.handlers[payload.Type]
	if !ok {
		w.Logger.Printf("No handler for event %s", payload.Type)
		return ErrPayloadTypeNotSupported
	}
	if err := handler(payload, client); err != nil {
		return err
	}
	return nil
}

func (w *WebsocketServer) FollowRequestHandler(p Payload, c *Client) error {
	// c.clientID
	var followerID int
	err := json.Unmarshal(p.Data, &followerID)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to start following user %v", c.clientID, followerID)
	return nil
}

func (w *WebsocketServer) FollowAcceptHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}

func (w *WebsocketServer) NewMessageHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}

func (w *WebsocketServer) GroupInviteHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}

func (w *WebsocketServer) GroupAcceptHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}
