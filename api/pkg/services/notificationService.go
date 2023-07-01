package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"os"
	"time"
)

type INotificationService interface {
	GetUserNotifications(userId int64) ([]*models.NotificationJSON, error)
	CreateFollowRequest(followerId int64, followingId int64) (int64, error)
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

func (s *NotificationService) GetUserNotifications(userId int64) ([]*models.NotificationJSON, error) {

	notifications, err := s.NotificationRepository.GetByReceiverId(userId)
	if err != nil {
		s.Logger.Printf("Cannot get user notifications: %s", err)
		return nil, err
	}

	for _, notification := range notifications {

		switch notification.NotificationType {
		case "follow_request":
			sender, err := s.UserRepo.GetById(notification.SenderId)
			if err != nil {
				s.Logger.Printf("Cannot get sender: %s", err)
				return nil, err
			}
			if sender.Nickname == "" {
				notification.SenderName = sender.FirstName + " " + sender.LastName
			} else {
				notification.SenderName = sender.Nickname
			}
			// COMMENTED OUT UNTIL GROUPS AND EVENTS ARE IMPLEMENTED
			/*
				case "group_invite":
					group, err := s.GroupRepo.GetById(notification.EntityId)
					if err != nil {
						s.Logger.Printf("Cannot get group: %s", err)
						return nil, err
					}
					notification.GroupName = group.Name
				case "group_request":
					group, err := s.GroupRepo.GetById(notification.EntityId)
					if err != nil {
						s.Logger.Printf("Cannot get group: %s", err)
						return nil, err
					}
					notification.GroupName = group.Name
				case "event_invite":
					event, err := s.EventRepo.GetById(notification.EntityId)
					if err != nil {
						s.Logger.Printf("Cannot get event: %s", err)
						return nil, err
					}
					notification.EventName = event.Name
					notification.EventDate = event.Date.Format(time.RFC3339)
			*/
		}
	}

	s.Logger.Printf("User notifications returned: %d", len(notifications))

	return notifications, nil
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
