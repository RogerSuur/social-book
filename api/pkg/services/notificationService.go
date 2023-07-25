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
	CreateGroupRequest(senderId int64, groupId int64) (int64, error)
}

type NotificationService struct {
	Logger                 *log.Logger
	UserRepo               models.IUserRepository
	FollowerRepo           models.IFollowerRepository
	NotificationRepository models.INotificationRepository
	GroupRepo              models.IGroupRepository
	GroupMemberRepo        models.IGroupMemberRepository
	EventRepo              models.IEventRepository
}

func InitNotificationService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
	groupRepo *models.GroupRepository,
	groupMemberRepo *models.GroupMemberRepository,
	eventRepo *models.EventRepository,
) *NotificationService {
	return &NotificationService{
		Logger:                 logger,
		UserRepo:               userRepo,
		FollowerRepo:           followerRepo,
		NotificationRepository: notificationRepo,
		GroupRepo:              groupRepo,
		GroupMemberRepo:        groupMemberRepo,
		EventRepo:              eventRepo,
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
			ReceiverId:       userId,
			NotificationType: notification.NotificationType,
			NotificationId:   notification.Id,
			SenderId:         notification.SenderId,
		}
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

		if notification.NotificationType == "group_invite" || notification.NotificationType == "group_request" || notification.NotificationType == "event_invite" {
			group, err := s.GroupRepo.GetById(notification.EntityId)
			if err != nil {
				s.Logger.Printf("Cannot get group: %s", err)
				return nil, err
			}
			singleNotification.GroupId = group.Id
			singleNotification.GroupName = group.Title
		}

		if notification.NotificationType == "event_invite" {
			s.Logger.Printf("Getting event: %d", notification.EntityId)
			event, err := s.EventRepo.GetById(notification.EntityId)
			if err != nil {
				s.Logger.Printf("Cannot get event: %s", err)
				return nil, err
			}
			singleNotification.EventId = event.Id
			singleNotification.EventName = event.Title
			singleNotification.EventDate = event.EventTime
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

func (s *NotificationService) CreateGroupRequest(senderId int64, groupId int64) (int64, error) {

	// check if sender and group exist
	_, err := s.UserRepo.GetById(senderId)
	if err != nil {
		s.Logger.Printf("Sender not found: %s", err)
		return -1, err
	}
	_, err = s.GroupRepo.GetById(groupId)
	if err != nil {
		s.Logger.Printf("Group not found: %s", err)
		return -1, err
	}

	// check if user is already member of group
	isMember, err := s.GroupMemberRepo.IsGroupMember(senderId, groupId)
	if err == nil {
		return -1, errors.New("error in checking if user is already member of group")
	}
	if isMember {
		return -1, errors.New("user is already member of group")
	}

	// add member to group with joined at Zero
	groupMember := &models.GroupMember{
		UserId:   senderId,
		GroupId:  groupId,
		JoinedAt: time.Time{},
	}

	lastID, err := s.GroupMemberRepo.Insert(groupMember)
	if err != nil {
		s.Logger.Printf("Cannot insert group request: %s", err)
		return -1, err
	}

	s.Logger.Printf("Member added: %d", lastID)

	// create notification
	notification := models.Notification{
		ReceiverId:       groupId,
		NotificationType: "group_request",
		SenderId:         senderId,
		EntityId:         lastID,
		CreatedAt:        time.Now(),
		Reaction:         false,
	}

	notificationId, err := s.NotificationRepository.Insert(&notification)
	if err != nil {
		return -1, err
	}

	return notificationId, nil
}
