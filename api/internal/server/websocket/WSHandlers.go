package websocket

import (
	"encoding/json"
	"errors"
)

type PayloadHandler func(payload Payload, client *Client) error

var (
	ErrPayloadTypeNotSupported = errors.New("this payload type is not supported")
)

const (
	FollowRequest   = "follow_request"
	Unfollow        = "unfollow"
	RequestChatlist = "request_chatlist"
	MessageHistory  = "request_message_history"
	Message         = "message"
	GroupRequest    = "group_request"
	Response        = "response"
)

func (w *WebsocketServer) setupHandlers() {
	w.handlers[FollowRequest] = w.FollowRequestHandler
	w.handlers[Unfollow] = w.UnfollowHandler
	w.handlers[RequestChatlist] = w.RequestChatlistHandler
	w.handlers[MessageHistory] = w.MessageHistoryHandler
	w.handlers[Message] = w.NewMessageHandler
	w.handlers[GroupRequest] = w.GroupRequestHandler
	w.handlers[Response] = w.ResponseHandler
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

func (w *WebsocketServer) ResponseHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v responded to notification %v with %v", c.clientID, data.ID, data.Reaction)
	return nil
}

func (w *WebsocketServer) FollowRequestHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to start following user %v", c.clientID, data.ID)

	user, err := w.userService.GetUserData(int64(c.clientID))
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v found", user.Email)

	followRequestId, err := w.notificationService.CreateFollowRequest(int64(c.clientID), int64(data.ID))
	if err != nil {
		return err
	}

	w.Logger.Printf("Created follow request with id %v", followRequestId)

	// broadcast to recipient
	userData, err := w.userService.GetUserData(int64(c.clientID))
	if err != nil {
		return err
	}

	recipientClient := w.getClientByUserID(int64(data.ID))

	if recipientClient == nil {
		w.Logger.Printf("Recipient client not found (recipient offline)")
		return nil
	}

	w.Logger.Printf("Recipient client found (recipient online)")

	dataToSend, err := json.Marshal(
		&NotificationPayload{
			NotificationType: "follow_request",
			NotificationID:   int(followRequestId),
			SenderID:         int(c.clientID),
			SenderName:       userData.FirstName + " " + userData.LastName,
		},
	)

	if err != nil {
		return err
	}

	recipientClient.gate <- Payload{
		Type: "notification",
		Data: dataToSend,
	}

	w.Logger.Printf("Sent notification to recipient")

	return nil
}

func (w *WebsocketServer) UnfollowHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to unfollow user %v", c.clientID, data.ID)
	return nil
}

func (w *WebsocketServer) RequestChatlistHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}

func (w *WebsocketServer) MessageHistoryHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to open message history with user %v", c.clientID, data.ID)
	return nil
}

func (w *WebsocketServer) NewMessageHandler(p Payload, c *Client) error {
	w.Logger.Println(p)
	return nil
}

func (w *WebsocketServer) GroupRequestHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to join group %v", c.clientID, data.ID)
	return nil
}
