package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"time"
)

type IChatService interface {
	GetChatlist(userID int64) ([]ChatListUser, error)
	GetLastMessage(userId int64, otherId int64) (*Message, error)
}

type ChatService struct {
	Logger   *log.Logger
	UserRepo models.IUserRepository
	ChatRepo models.IMessageRepository
}

func InitChatService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	chatRepo *models.MessageRepository,
) *ChatService {
	return &ChatService{
		Logger:   logger,
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
}

func (s *ChatService) GetChatlist(userID int64) ([]ChatListUser, error) {

	chatlist, err := s.ChatRepo.GetChatUsers(userID)
	if err != nil {
		return nil, err
	}

	chatlistData := []ChatListUser{}

	for _, user := range chatlist {
		if user.Nickname == "" {
			user.Nickname = user.FirstName + " " + user.LastName
		}

		//lastMessage, err := s.ChatRepo.GetLastMessage(userID, int64(user.Id))
		if err != nil {
			return nil, err
		}

		chatData := ChatListUser{
			UserID:    int(user.Id),
			GroupID:   0,
			Name:      user.Nickname,
			Timestamp: time.Now(), // lastMessage.Timestamp
		}
		chatlistData = append(chatlistData, chatData)
	}

	return chatlistData, nil
}

func (s *ChatService) GetLastMessage(userId int64, otherId int64) (*Message, error) {

	//TODO

	return nil, nil
}
