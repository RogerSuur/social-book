package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"sort"
	"time"
)

type IChatService interface {
	GetChatlist(userID int64) ([]ChatListUser, error)
	CreateMessage(message *models.Message) (int64, error)
	GetMessageHistory(userId int64, otherId int64, groupId int64, lastMessage int64) ([]*MessageJSON, error)
}

type ChatService struct {
	Logger    *log.Logger
	UserRepo  models.IUserRepository
	ChatRepo  models.IMessageRepository
	GroupRepo models.IGroupRepository
}

func InitChatService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	chatRepo *models.MessageRepository,
	groupRepo *models.GroupRepository,
) *ChatService {
	return &ChatService{
		Logger:    logger,
		UserRepo:  userRepo,
		ChatRepo:  chatRepo,
		GroupRepo: groupRepo,
	}
}

type ChatListUser struct {
	UserID      int       `json:"user_id"`
	GroupID     int       `json:"group_id"`
	Name        string    `json:"name"`
	Timestamp   time.Time `json:"timestamp"`
	AvatarImage string    `json:"avatar_image"`
	UnreadCount int       `json:"unread_count"`
}

type MessageJSON struct {
	Id            int64     `json:"id"`
	SenderId      int64     `json:"sender_id"`
	SenderName    string    `json:"sender_name"`
	RecipientId   int64     `json:"recipient_id"`
	RecipientName string    `json:"recipient_name"`
	GroupId       int64     `json:"group_id"`
	GroupName     string    `json:"group_name"`
	Content       string    `json:"body"`
	SentAt        time.Time `json:"timestamp"`
	//ReadAt        time.Time `json:"read_at"`
}

func (s *ChatService) GetChatlist(userID int64) ([]ChatListUser, error) {

	userList, err := s.ChatRepo.GetChatUsers(userID)
	if err != nil {
		return nil, err
	}

	chatlistData := []ChatListUser{}

	for _, user := range userList {
		if user.Nickname == "" {
			user.Nickname = user.FirstName + " " + user.LastName
		}

		lastMessage, err := s.ChatRepo.GetLastMessage(userID, int64(user.Id), false)
		if err != nil {
			return nil, err
		}
		if lastMessage.SentAt == (time.Time{}) {
			lastMessage.SentAt = user.CreatedAt
		}

		unreadCount, err := s.ChatRepo.GetUnreadCount(userID, int64(user.Id), false)
		if err != nil {
			return nil, err
		}

		chatData := ChatListUser{
			UserID:      int(user.Id),
			GroupID:     0,
			Name:        user.Nickname,
			Timestamp:   lastMessage.SentAt,
			AvatarImage: user.ImagePath,
			UnreadCount: int(unreadCount),
		}
		chatlistData = append(chatlistData, chatData)
	}

	groupList, err := s.ChatRepo.GetChatGroups(userID)
	if err != nil {
		return nil, err
	}

	for _, group := range groupList {
		lastMessage, err := s.ChatRepo.GetLastMessage(userID, int64(group.Id), true)
		if err != nil {
			return nil, err
		}
		if lastMessage.SentAt == (time.Time{}) {
			lastMessage.SentAt = group.CreatedAt
		}

		unreadCount, err := s.ChatRepo.GetUnreadCount(userID, int64(group.Id), true)
		if err != nil {
			return nil, err
		}

		chatData := ChatListUser{
			UserID:      0,
			GroupID:     int(group.Id),
			Name:        group.Title,
			Timestamp:   lastMessage.SentAt,
			AvatarImage: group.ImagePath,
			UnreadCount: int(unreadCount),
		}
		chatlistData = append(chatlistData, chatData)
	}

	// sort the chatlistData array by ChatListUser.Timestamp field in descending order

	sort.Slice(chatlistData, func(i, j int) bool {
		return chatlistData[i].Timestamp.After(chatlistData[j].Timestamp)
	})

	return chatlistData, nil
}

func (s *ChatService) CreateMessage(message *models.Message) (int64, error) {

	// check if users exist
	_, err := s.UserRepo.GetById(message.SenderId)
	if err != nil {
		s.Logger.Printf("User with id %d does not exist", message.SenderId)
		return -1, err
	}

	if message.RecipientId != 0 {
		_, err = s.UserRepo.GetById(message.RecipientId)
		if err != nil {
			s.Logger.Printf("User with id %d does not exist", message.RecipientId)
			return -1, err
		}
	} else if message.GroupId != 0 {
		_, err = s.GroupRepo.GetById(message.GroupId)
		if err != nil {
			s.Logger.Printf("Group with id %d does not exist", message.GroupId)
			return -1, err
		}
	} else {
		s.Logger.Printf("Neither recipient nor group id is specified")
		return -1, errors.New("neither recipient nor group id is specified")
	}

	lastID, err := s.ChatRepo.Insert(message)
	if err != nil {
		return -1, err
	}

	s.Logger.Printf("Message created: %d", lastID)

	return lastID, nil
}

func (s *ChatService) GetMessageHistory(userId int64, otherId int64, groupId int64, lastMessage int64) ([]*MessageJSON, error) {

	// check if users exist

	var messages []*models.Message
	groupData := &models.Group{
		Title: "",
	}
	recipientData := &models.User{
		Nickname: "",
	}
	var err error

	if otherId != 0 {

		if lastMessage == 0 {
			lastFullMessage, err := s.ChatRepo.GetLastMessage(userId, otherId, false)
			if err != nil {
				return nil, err
			}
			lastMessage = lastFullMessage.Id + 1
		}

		messages, err = s.ChatRepo.GetMessagesByUserIds(userId, otherId, lastMessage)
		if err != nil {
			return nil, err
		}

	} else if groupId != 0 {

		if lastMessage == 0 {
			lastFullMessage, err := s.ChatRepo.GetLastMessage(userId, groupId, true)
			if err != nil {
				return nil, err
			}
			lastMessage = lastFullMessage.Id + 1
		}

		messages, err = s.ChatRepo.GetMessagesByGroupId(groupId, lastMessage)
		if err != nil {
			return nil, err
		}

	} else {
		s.Logger.Printf("Neither recipient nor group id is specified")
		return nil, errors.New("neither recipient nor group id is specified")
	}

	messagesJSON := []*MessageJSON{}

	if len(messages) == 0 {
		return messagesJSON, nil
	}

	for i := len(messages) - 1; i >= 0; i-- {
		message := messages[i]

		senderData, err := s.UserRepo.GetById(message.SenderId)
		if err != nil {
			s.Logger.Printf("User with id %d does not exist", userId)
			return nil, err
		}

		if senderData.Nickname == "" {
			senderData.Nickname = senderData.FirstName + " " + senderData.LastName
		}

		if otherId != 0 {
			recipientData, err := s.UserRepo.GetById(message.RecipientId)
			if err != nil {
				s.Logger.Printf("User with id %d does not exist", otherId)
				return nil, err
			}

			if recipientData.Nickname == "" {
				recipientData.Nickname = recipientData.FirstName + " " + recipientData.LastName
			}
		}

		if groupId != 0 {
			groupData, err = s.GroupRepo.GetById(groupId)
			if err != nil {
				s.Logger.Printf("Group with id %d does not exist", groupId)
				return nil, err
			}
		}

		messageJSON := &MessageJSON{
			Id:            message.Id,
			SenderId:      message.SenderId,
			SenderName:    senderData.Nickname,
			RecipientId:   message.RecipientId,
			RecipientName: recipientData.Nickname,
			GroupId:       groupId,
			GroupName:     groupData.Title,
			Content:       message.Content,
			SentAt:        message.SentAt,
			//ReadAt:        message.ReadAt,
		}

		/*
			if message.SenderId != userId && groupId == 0 {
				messageJSON.SenderName = otherData.Nickname
				messageJSON.RecipientName = userData.Nickname
			}
		*/

		messagesJSON = append(messagesJSON, messageJSON)
	}

	// mark messages as read
	if otherId != 0 {
		err = s.ChatRepo.MarkMessagesAsRead(userId, otherId, false)
		if err != nil {
			return nil, err
		}
	} else if groupId != 0 {
		err = s.ChatRepo.MarkMessagesAsRead(userId, groupId, true)
		if err != nil {
			return nil, err
		}
	}

	return messagesJSON, nil
}
