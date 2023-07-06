package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"time"
)

type INotificationService interface {
	GetById(notificationId int64) (*models.Notification, error)
	GetUserNotifications(userId int64) ([]*models.NotificationJSON, error)
	CreateFollowRequest(followerId int64, followingId int64) (int64, bool, error)
	HandleFollowRequest(notificationId int64, accepted bool) error
}

type NotificationService struct {
	Logger                 *log.Logger
	UserRepo               models.IUserRepository
	FollowerRepo           models.IFollowerRepository
	NotificationRepository models.INotificationRepository
}

func InitNotificationService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
) *NotificationService {
	return &NotificationService{
		Logger:                 logger,
		UserRepo:               userRepo,
		FollowerRepo:           followerRepo,
		NotificationRepository: notificationRepo,
	}
}

func (s *NotificationService) GetById(notificationId int64) (*models.Notification, error) {

	notification, err := s.NotificationRepository.GetById(notificationId)
	if err != nil {
		s.Logger.Printf("Cannot get notification: %s", err)
		return nil, err
	}

	s.Logger.Printf("Notification returned: %d", notification.Id)

	return notification, nil
}

func (s *NotificationService) GetUserNotifications(userId int64) ([]*models.NotificationJSON, error) {

	notifications, err := s.NotificationRepository.GetByReceiverId(userId)
	if err != nil {
		s.Logger.Printf("Cannot get user notifications: %s", err)
		return nil, err
	}

	NotificationJSON := []*models.NotificationJSON{}

	for _, notification := range notifications {
		singleNotification := &models.NotificationJSON{
			NotificationType: notification.NotificationType,
			NotificationId:   notification.Id,
			SenderId:         notification.SenderId,
		}
		switch notification.NotificationType {
		case "follow_request":
			sender, err := s.UserRepo.GetById(notification.SenderId)
			if err != nil {
				s.Logger.Printf("Cannot get sender: %s", err)
				return nil, err
			}
			if sender.Nickname == "" {
				singleNotification.SenderName = sender.FirstName + " " + sender.LastName
			} else {
				singleNotification.SenderName = sender.Nickname
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
		NotificationJSON = append(NotificationJSON, singleNotification)
	}

	s.Logger.Printf("User notifications returned: %d", len(NotificationJSON))

	return NotificationJSON, nil
}

func (s *NotificationService) CreateFollowRequest(followerId int64, followingId int64) (int64, bool, error) {

	// check if follower and following exist
	_, err := s.UserRepo.GetById(followerId)
	if err != nil {
		s.Logger.Printf("Follower not found: %s", err)
		return -1, false, err
	}
	following, err := s.UserRepo.GetById(followingId)
	if err != nil {
		s.Logger.Printf("Following not found: %s", err)
		return -1, false, err
	}

	// check if follow request already exists
	_, err = s.FollowerRepo.GetByFollowerAndFollowing(followerId, followingId)
	if err == nil {
		return -1, false, errors.New("follow request already exists")
	}

	// check if following is private
	follower := &models.Follower{
		FollowerId:  followerId,
		FollowingId: followingId,
		Accepted:    following.IsPublic,
	}

	// create follow reque

	lastID, err := s.FollowerRepo.Insert(follower)
	if err != nil {
		s.Logger.Printf("Cannot insert follow request: %s", err)
		return -1, false, err
	}

	s.Logger.Printf("Follow request created: %d", lastID)

	// create notification
	notification := models.Notification{
		ReceiverId:       followingId,
		NotificationType: "follow_request",
		SenderId:         followerId,
		EntityId:         lastID,
		CreatedAt:        time.Now(),
		Reaction:         false,
	}

	notificationId, err := s.NotificationRepository.Insert(&notification)
	if err != nil {
		return -1, false, err
	}

	return notificationId, following.IsPublic, nil
}

func (s *NotificationService) HandleFollowRequest(notificationId int64, accepted bool) error {

	notification, err := s.NotificationRepository.GetById(notificationId)
	if err != nil {
		s.Logger.Printf("Cannot get notification: %s", err)
		return err
	}

	// check if follow request already handled
	if notification.Reaction {
		return errors.New("follow request already handled")
	}

	// check if follow request exists
	follower, err := s.FollowerRepo.GetById(notification.EntityId)
	if err != nil {
		s.Logger.Printf("Cannot get follow request: %s", err)
		return err
	}

	// check if follow request is accepted
	if follower.Accepted {
		return errors.New("follow request already accepted")
	}

	// update follow request
	follower.Accepted = accepted
	err = s.FollowerRepo.Update(follower)
	if err != nil {
		s.Logger.Printf("Cannot update follow request: %s", err)
		return err
	}

	s.Logger.Printf("Follow request updated: %d", follower.Id)

	// update notification
	notification.Reaction = true
	err = s.NotificationRepository.Update(notification)
	if err != nil {
		s.Logger.Printf("Cannot update notification: %s", err)
		return err
	}

	s.Logger.Printf("Notification updated: %d", notification.Id)

	return nil
}
