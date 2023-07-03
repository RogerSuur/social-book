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
	GetLastMessage(userId int64, otherId int64) (*Message, error)
	CreateMessage(message *models.Message) (int64, error)
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

type Message struct {
}

type ChatListUser struct {
	UserID      int       `json:"user_id"`
	GroupID     int       `json:"group_id"`
	Name        string    `json:"name"`
	Timestamp   time.Time `json:"timestamp"`
	AvatarImage string    `json:"avatar_image"`
	UnreadCount int       `json:"unread_count"`
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

func (s *ChatService) GetLastMessage(userId int64, otherId int64) (*Message, error) {

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

	lastID, err := s.ChatRepo.Insert(message)
	if err != nil {
		return -1, err
	}

	s.Logger.Printf("Message created: %d", lastID)

	return lastID, nil
}
