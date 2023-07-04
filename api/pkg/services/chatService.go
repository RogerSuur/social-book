package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"os"
	"sort"
	"time"
)

type IChatService interface {
	GetChatlist(userID int64) ([]ChatListUser, error)
	GetLastMessage(userId int64, otherId int64) (*MessageJSON, error)
	CreateMessage(message *models.Message) (int64, error)
	GetMessageHistory(userId int64, otherId int64) ([]*MessageJSON, error)
}

type ChatService struct {
	Logger   *log.Logger
	UserRepo models.IUserRepository
	ChatRepo models.IMessageRepository
}

func InitChatService(
	userRepo *models.UserRepository,
	chatRepo *models.MessageRepository,
) *ChatService {
	return &ChatService{
		Logger:   log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo: userRepo,
		ChatRepo: chatRepo,
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
	Content       string    `json:"content"`
	SentAt        time.Time `json:"sent_at"`
	ReadAt        time.Time `json:"read_at"`
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
		return chatlistData[i].Timestamp.Before(chatlistData[j].Timestamp)
	})

	return chatlistData, nil
}

func (s *ChatService) GetLastMessage(userId int64, otherId int64) (*MessageJSON, error) {

	//TODO

	return nil, nil
}

func (s *ChatService) CreateMessage(message *models.Message) (int64, error) {

	// check if users exist
	_, err := s.UserRepo.GetById(message.SenderId)
	if err != nil {
		s.Logger.Printf("User with id %d does not exist", message.SenderId)
		return -1, err
	}
	_, err = s.UserRepo.GetById(message.RecipientId)
	if err != nil {
		s.Logger.Printf("User with id %d does not exist", message.RecipientId)
		return -1, err
	}

	message.SentAt = time.Now()

	lastID, err := s.ChatRepo.Insert(message)
	if err != nil {
		return -1, err
	}

	s.Logger.Printf("Message created: %d", lastID)

	return lastID, nil
}

func (s *ChatService) GetMessageHistory(userId int64, otherId int64) ([]*MessageJSON, error) {

	// check if users exist
	userData, err := s.UserRepo.GetById(userId)
	if err != nil {
		s.Logger.Printf("User with id %d does not exist", userId)
		return nil, err
	}
	otherData, err := s.UserRepo.GetById(otherId)
	if err != nil {
		s.Logger.Printf("User with id %d does not exist", otherId)
		return nil, err
	}

	// get messages
	messages, err := s.ChatRepo.GetMessagesByUserIds(userId, otherId)
	if err != nil {
		return nil, err
	}

	messagesJSON := []*MessageJSON{}

	if len(messages) == 0 {
		return messagesJSON, nil
	}

	if userData.Nickname == "" {
		userData.Nickname = userData.FirstName + " " + userData.LastName
	}

	if otherData.Nickname == "" {
		otherData.Nickname = otherData.FirstName + " " + otherData.LastName
	}

	for _, message := range messages {
		messageJSON := &MessageJSON{
			Id:            message.Id,
			SenderId:      message.SenderId,
			SenderName:    userData.Nickname,
			RecipientId:   message.RecipientId,
			RecipientName: otherData.Nickname,
			GroupId:       0,
			GroupName:     "",
			Content:       message.Content,
			SentAt:        message.SentAt,
			ReadAt:        message.ReadAt,
		}
		messagesJSON = append(messagesJSON, messageJSON)
	}

	// mark messages as read
	err = s.ChatRepo.MarkMessagesAsRead(userId, otherId, false)
	if err != nil {
		return nil, err
	}

	return messagesJSON, nil
}
