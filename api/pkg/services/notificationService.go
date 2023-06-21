package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"os"
)

type INotificationService interface {
	CreateFollowRequestNotification(followerId int64, followingId int64) error
}

type NotificationService struct {
	Logger                 *log.Logger
	UserRepo               models.IUserRepository
	SessionRepo            models.ISessionRepository
	FollowerRepo           models.IFollowerRepository
	NotificationRepository models.INotificationRepository
}

func InitNotificationService(
	Logger *log.Logger,
	userRepo *models.UserRepository,
	sessionRepo *models.SessionRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
) *NotificationService {
	return &NotificationService{
		Logger:                 log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo:               userRepo,
		SessionRepo:            sessionRepo,
		FollowerRepo:           followerRepo,
		NotificationRepository: notificationRepo,
	}
}

func (service *NotificationService) CreateFollowRequestNotification(followingId int64, followReqId int64) error {

	notification := models.Notification{
		ReceiverId:      followingId,
		FollowRequestId: followReqId,
		Reaction:        false,
	}

	_, err := service.NotificationRepository.Insert(&notification)
	if err != nil {
		return err
	}

	return nil
}
