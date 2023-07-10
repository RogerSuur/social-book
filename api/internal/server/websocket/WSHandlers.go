package websocket

import (
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"errors"
	"time"
)

type PayloadHandler func(payload Payload, client *Client) error

var (
	ErrPayloadTypeNotSupported = errors.New("this payload type is not supported")
	ErrorInvalidPayload        = errors.New("invalid payload")
	ErrorInvalidNotification   = errors.New("invalid notification")
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

	notification, err := w.notificationService.GetById(int64(data.ID))
	if err != nil {
		return err
	}

	if notification == nil {
		w.Logger.Printf("Notification not found")
		return ErrorInvalidNotification
	}

	if notification.ReceiverId != int64(c.clientID) {
		w.Logger.Printf("Notification does not belong to user")
		return ErrorInvalidNotification
	}

	// perhaps case switch here?
	if notification.NotificationType == "follow_request" {
		w.Logger.Printf("User %v accepted follow request %v", c.clientID, data.ID)
		err = w.notificationService.HandleFollowRequest(int64(data.ID), data.Reaction)
		if err != nil {
			return err
		}
	}

	// TODO: handle other notification types

	return nil
}

func (w *WebsocketServer) FollowRequestHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to start following user %v", c.clientID, data.ID)

	followRequestId, sendNewChatlist, err := w.notificationService.CreateFollowRequest(int64(c.clientID), int64(data.ID))
	if err != nil {
		return err
	}

	w.Logger.Printf("Created follow request with id %v", followRequestId)

	// broadcast to recipient
	userData, err := w.userService.GetUserData(int64(c.clientID), int64(c.clientID))
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

	// send new chatlist to sender
	if sendNewChatlist {
		chatlist, err := w.chatService.GetChatlist(int64(c.clientID))
		if err != nil {
			return err
		}

		dataToSend, err := json.Marshal(
			&ChatListPayload{
				UserID:   int(c.clientID),
				Chatlist: chatlist,
			},
		)

		if err != nil {
			return err
		}

		c.gate <- Payload{
			Type: "chatlist",
			Data: dataToSend,
		}

		w.Logger.Printf("Sent new chatlist to sender %v", c.clientID)
	}

	return nil
}

func (w *WebsocketServer) UnfollowHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to unfollow user %v", c.clientID, data.ID)

	err = w.userService.Unfollow(int64(c.clientID), int64(data.ID))
	if err != nil {
		return err
	}
	w.Logger.Printf("User successfully %v unfollowed user %v", c.clientID, data.ID)

	return nil
}

func (w *WebsocketServer) RequestChatlistHandler(p Payload, c *Client) error {
	w.Logger.Printf("User %v has requested chatlist", c.clientID)

	chatlist, err := w.chatService.GetChatlist(int64(c.clientID))
	if err != nil {
		return err
	}
	w.Logger.Printf("Chatlist successfully retrieved (%v chats)", len(chatlist))

	dataToSend, err := json.Marshal(
		&ChatListPayload{
			UserID:   int(c.clientID),
			Chatlist: chatlist,
		},
	)

	if err != nil {
		return err
	}

	c.gate <- Payload{
		Type: "chatlist",
		Data: dataToSend,
	}

	w.Logger.Printf("Sent chatlist to user %v", c.clientID)

	return nil
}

func (w *WebsocketServer) MessageHistoryHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}

	if data.ID == 0 && data.GroupID > 0 {
		w.Logger.Printf("User %v requests message history with group %v starting from %v", c.clientID, data.GroupID, data.LastMessage)
	} else if data.GroupID == 0 && data.ID > 0 {
		w.Logger.Printf("User %v requests message history with user %v starting from %v", c.clientID, data.ID, data.LastMessage)
	} else {
		w.Logger.Printf("Invalid request payload")
		return ErrorInvalidPayload
	}

	messages, err := w.chatService.GetMessageHistory(int64(c.clientID), int64(data.ID), int64(data.GroupID), int64(data.LastMessage))
	if err != nil {
		return err
	}

	w.Logger.Printf("Message history successfully retrieved (%v messages)", len(messages))

	dataToSend, err := json.Marshal(messages)

	if err != nil {
		return err
	}

	c.gate <- Payload{
		Type: "message_history",
		Data: dataToSend,
	}

	return nil
}

func (w *WebsocketServer) NewMessageHandler(p Payload, c *Client) error {
	data := &MessagePayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}

	if data.GroupID == 0 && data.RecipientID > 0 {
		w.Logger.Printf("User %v sent message to user %v", c.clientID, data.RecipientID)
	} else if data.RecipientID == 0 && data.GroupID > 0 {
		w.Logger.Printf("User %v sent message to group %v", c.clientID, data.GroupID)
	} else {
		w.Logger.Printf("Invalid request payload")
		return ErrorInvalidPayload
	}

	messageData := &models.Message{
		SenderId:    c.clientID,
		RecipientId: int64(data.RecipientID),
		GroupId:     int64(data.GroupID),
		Content:     data.Content,
	}

	messageID, err := w.chatService.CreateMessage(messageData)
	if err != nil {
		return err
	}

	w.Logger.Printf("Message successfully created with id %v", messageID)

	// broadcast to recipient

	// todo: check if recipient is group

	recipientClient := w.getClientByUserID(int64(data.RecipientID))

	if recipientClient == nil {
		w.Logger.Printf("Recipient client not found (recipient offline)")
	} else {
		w.Logger.Printf("Recipient client found (recipient online)")

		userData, err := w.userService.GetUserData(c.clientID, c.clientID)
		if err != nil {
			return err
		}

		if userData.Nickname == "" {
			userData.Nickname = userData.FirstName + " " + userData.LastName
		}

		recipientData, err := w.userService.GetUserData(c.clientID, int64(data.RecipientID))
		if err != nil {
			return err
		}

		if recipientData.Nickname == "" {
			recipientData.Nickname = recipientData.FirstName + " " + recipientData.LastName
		}

		dataToSend, err := json.Marshal(
			&MessagePayload{
				MessageID:     int(messageID),
				SenderID:      int(c.clientID),
				SenderName:    userData.Nickname,
				SenderImage:   userData.AvatarImage,
				RecipientID:   recipientData.UserID,
				RecipientName: recipientData.Nickname,
				//GroupID:       data.GroupID,
				//GroupName:     data.GroupName,
				Content:   messageData.Content,
				Timestamp: time.Now(),
			},
		)

		if err != nil {
			return err
		}

		recipientClient.gate <- Payload{
			Type: "message",
			Data: dataToSend,
		}

		w.Logger.Printf("Sent message to recipient")

	}

	return nil
}

func (w *WebsocketServer) GroupRequestHandler(p Payload, c *Client) error {
	data := &RequestPayload{}
	err := json.Unmarshal(p.Data, &data)
	if err != nil {
		return err
	}
	w.Logger.Printf("User %v wants to join group %v", c.clientID, data.GroupID)
	return nil
}
