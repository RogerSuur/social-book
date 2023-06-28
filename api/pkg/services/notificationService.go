package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"os"
	"time"
)

type INotificationService interface {
	CreateFollowRequest(followerID int64, followingID int64) (int64, error)
}

type NotificationService struct {
	Logger                 *log.Logger
	UserRepo               models.IUserRepository
	FollowerRepo           models.IFollowerRepository
	NotificationRepository models.INotificationRepository
}

func InitNotificationService(
	userRepo *models.UserRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
) *NotificationService {
	return &NotificationService{
		Logger:                 log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo:               userRepo,
		FollowerRepo:           followerRepo,
		NotificationRepository: notificationRepo,
	}
}

func (s *NotificationService) CreateFollowRequest(followerId int64, followingId int64) (int64, error) {

	// check if follower and following exist
	_, err := s.UserRepo.GetById(followerId)
	if err != nil {
		s.Logger.Printf("Follower not found: %s", err)
		return -1, err
	}
	_, err = s.UserRepo.GetById(followingId)
	if err != nil {
		s.Logger.Printf("Following not found: %s", err)
		return -1, err
	}

	// check if follow request already exists
	_, err = s.FollowerRepo.GetByFollowerAndFollowing(followerId, followingId)
	if err == nil {
		return -1, errors.New("follow request already exists")
	}

	// create follow request
	follower := &models.Follower{
		FollowerId:  followerId,
		FollowingId: followingId,
		Accepted:    false,
	}

	lastID, err := s.FollowerRepo.Insert(follower)
	if err != nil {
		s.Logger.Printf("Cannot insert follow request: %s", err)
		return -1, err
	}

	s.Logger.Printf("Follow request created: %d", lastID)

	// create notification
	notification := models.Notification{
		ReceiverId:       followingId,
		NotificationType: "follow_request",
		SenderID:         followerId,
		EntityId:         lastID,
		CreatedAt:        time.Now(),
		Reaction:         false,
	}

	_, err = s.NotificationRepository.Insert(&notification)
	if err != nil {
		return -1, err
	}

	return lastID, nil
}
