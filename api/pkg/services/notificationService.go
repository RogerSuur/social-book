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
	HandleGroupInvite(notificationID int64, accepted bool) error
	HandleGroupRequest(creatorID int64, notificationID int64, accepted bool) error
	HandleEventInvite(notificationID int64, accepted bool) error
}

type NotificationService struct {
	Logger                 *log.Logger
	UserRepo               models.IUserRepository
	FollowerRepo           models.IFollowerRepository
	NotificationRepository models.INotificationRepository
	GroupRepo              models.IGroupRepository
	GroupMemberRepo        models.IGroupMemberRepository
	EventRepo              models.IEventRepository
	EventAttendanceRepo    models.IEventAttendanceRepository
}

func InitNotificationService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
	groupRepo *models.GroupRepository,
	groupMemberRepo *models.GroupMemberRepository,
	eventRepo *models.EventRepository,
	eventAttendanceRepo *models.EventAttendanceRepository,
) *NotificationService {
	return &NotificationService{
		Logger:                 logger,
		UserRepo:               userRepo,
		FollowerRepo:           followerRepo,
		NotificationRepository: notificationRepo,
		GroupRepo:              groupRepo,
		GroupMemberRepo:        groupMemberRepo,
		EventRepo:              eventRepo,
		EventAttendanceRepo:    eventAttendanceRepo,
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

		if notification.NotificationType == "group_invite" || notification.NotificationType == "event_invite" {
			group, err := s.GroupRepo.GetById(notification.EntityId)
			if err != nil {
				s.Logger.Printf("Cannot get group: %s", err)
				return nil, err
			}
			singleNotification.GroupId = group.Id
			singleNotification.GroupName = group.Title
		}

		if notification.NotificationType == "group_request" {
			member, err := s.GroupMemberRepo.GetById(notification.EntityId)

			if err != nil {
				s.Logger.Printf("Cannot get group member: %s", err)
				return nil, err
			}

			group, err := s.GroupRepo.GetById(member.GroupId)
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
	isMember, err := s.GroupMemberRepo.IsGroupMember(groupId, senderId)
	if err != nil {
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

func (s *NotificationService) HandleGroupInvite(notificationID int64, accepted bool) error {
	notification, err := s.NotificationRepository.GetById(notificationID)
	if err != nil {
		s.Logger.Printf("Cannot get notification: %s", err)
		return err
	}

	// check if group invite already handled
	if notification.Reaction {
		return errors.New("group invite already handled")
	}

	// check if group invite exists
	groupMember, err := s.GroupMemberRepo.GetById(notification.EntityId)
	if err != nil {
		s.Logger.Printf("Cannot get group invite: %s", err)
		return err
	}

	// check if group invite is accepted
	if groupMember.JoinedAt != (time.Time{}) {
		return errors.New("group invite already accepted")
	}

	// update group invite
	if accepted {
		groupMember.JoinedAt = time.Now()
		err = s.GroupMemberRepo.Update(groupMember)
		if err != nil {
			s.Logger.Printf("Cannot update group invite: %s", err)
			return err
		}
	} else {
		err = s.GroupMemberRepo.Delete(groupMember)
		if err != nil {
			s.Logger.Printf("Cannot delete group invite: %s", err)
			return err
		}
	}

	s.Logger.Printf("Group invite updated: %d", notification.EntityId)

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

func (s *NotificationService) HandleGroupRequest(creatorID int64, notificationID int64, accepted bool) error {

	notification, err := s.NotificationRepository.GetById(notificationID)
	if err != nil {
		s.Logger.Printf("Cannot get notification: %s", err)
		return err
	}

	// check if group request already handled
	if notification.Reaction {
		return errors.New("group request already handled")
	}

	// check if group request exists
	groupMember, err := s.GroupMemberRepo.GetById(notification.EntityId)
	if err != nil {
		s.Logger.Printf("Cannot get group request: %s", err)
		return err
	}

	// check if group request is accepted
	if groupMember.JoinedAt != (time.Time{}) {
		return errors.New("group request already accepted")
	}

	// check if user is creator of group
	group, err := s.GroupRepo.GetById(groupMember.GroupId)
	if err != nil {
		s.Logger.Printf("Cannot get group: %s", err)
		return err
	}

	if group.CreatorId != creatorID {
		return errors.New("user is not creator of group")
	}

	// update group request
	if accepted {
		groupMember.JoinedAt = time.Now()
		err = s.GroupMemberRepo.Update(groupMember)
		if err != nil {
			s.Logger.Printf("Cannot update group request: %s", err)
			return err
		}
	} else {
		err = s.GroupMemberRepo.Delete(groupMember)
		if err != nil {
			s.Logger.Printf("Cannot delete group request: %s", err)
			return err
		}
	}

	s.Logger.Printf("Group request updated: %d", notification.EntityId)

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

func (s *NotificationService) HandleEventInvite(notificationID int64, accepted bool) error {

	notification, err := s.NotificationRepository.GetById(notificationID)
	if err != nil {
		s.Logger.Printf("Cannot get notification: %s", err)
		return err
	}

	// check if event invite already handled
	if notification.Reaction {
		return errors.New("event invite already handled")
	}

	// check if event exists
	event, err := s.EventRepo.GetById(notification.EntityId)
	if err != nil {
		s.Logger.Printf("Cannot get event invite: %s", err)
		return err
	}

	// check if event invite has been processed
	attendees, err := s.EventAttendanceRepo.GetAttendeesByEventId(event.Id)
	if err != nil {
		s.Logger.Printf("Cannot get event attendees: %s", err)
		return err
	}

	for _, attendee := range attendees {
		if attendee.UserId == notification.ReceiverId {
			return errors.New("event invite already processed")
		}
	}

	// update event invite
	eventAttendance := &models.EventAttendance{
		UserId:      notification.ReceiverId,
		EventId:     event.Id,
		IsAttending: accepted,
	}
	_, err = s.EventAttendanceRepo.Insert(eventAttendance)
	if err != nil {
		s.Logger.Printf("Cannot insert event attendance: %s", err)
		return err
	}

	// update notification
	notification.Reaction = true
	err = s.NotificationRepository.Update(notification)
	if err != nil {
		s.Logger.Printf("Cannot update notification: %s", err)
		return err
	}

	s.Logger.Printf("Event attendance and notification updated: %d", notification.EntityId)

	return nil

}
