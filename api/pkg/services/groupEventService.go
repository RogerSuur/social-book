package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"time"
)

type EventJSON struct {
	Id          int           `json:"id"`
	GroupId     int           `json:"group_id"`
	GroupName   string        `json:"group_name"`
	UserId      int           `json:"creator_id"`
	NickName    string        `json:"creator_name"`
	CreatedAt   time.Time     `json:"created_at"`
	EventTime   time.Time     `json:"event_time"`
	TimeSpan    time.Duration `json:"timespan"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

type IGroupEventService interface {
	GetGroupEvents(groupId int64) ([]*models.Event, error)
	CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) (int64, error)
	GetUserEvents(userId int64) ([]*EventJSON, error)
}

type GroupEventService struct {
	Logger                         *log.Logger
	GroupEventAttendanceRepository models.IGroupEventAttendanceRepository
	EventRepository                models.IEventRepository
	GroupRepository                models.IGroupRepository
	GroupMemberRepository          models.IGroupMemberRepository
	UserRepository                 models.IUserRepository
	NotificationRepository         models.INotificationRepository
}

func InitGroupEventService(
	logger *log.Logger,
	groupEventAttendanceRepo *models.GroupEventAttendanceRepository,
	groupEventRepo *models.EventRepository,
	groupRepo *models.GroupRepository,
	GroupMemberRepository *models.GroupMemberRepository,
	userRepo *models.UserRepository,
	notificationRepo *models.NotificationRepository,
) *GroupEventService {
	return &GroupEventService{
		Logger:                         logger,
		GroupEventAttendanceRepository: groupEventAttendanceRepo,
		EventRepository:                groupEventRepo,
		GroupRepository:                groupRepo,
		GroupMemberRepository:          GroupMemberRepository,
		UserRepository:                 userRepo,
		NotificationRepository:         notificationRepo,
	}
}

func (s *GroupEventService) GetGroupEvents(groupId int64) ([]*models.Event, error) {

	events, err := s.EventRepository.GetAllByGroupId(groupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	return events, nil
}

func (s *GroupEventService) CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) (int64, error) {

	event := &models.Event{
		GroupId:     formData.GroupId,
		UserId:      userId,
		EventTime:   formData.EventTime,
		TimeSpan:    formData.TimeSpan,
		Title:       formData.Title,
		Description: formData.Description,
	}

	result, err := s.EventRepository.Insert(event)

	if err != nil {
		s.Logger.Printf("Failed inserting event: %s", err)
	}

	// Send notification to all group members

	groupMembers, err := s.GroupMemberRepository.GetGroupMembersByGroupId(formData.GroupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return -1, err
	}

	userData, err := s.UserRepository.GetById(userId)
	if err != nil {
		s.Logger.Printf("Failed fetching user data: %s", err)
		return -1, err
	}

	if userData.Nickname == "" {
		userData.Nickname = userData.FirstName + " " + userData.LastName
	}

	groupData, err := s.GroupRepository.GetById(formData.GroupId)
	if err != nil {
		s.Logger.Printf("Failed fetching group data: %s", err)
		return -1, err
	}

	for _, member := range groupMembers {
		notification := &models.Notification{
			ReceiverId:       member.Id,
			SenderId:         userId,
			EntityId:         result,
			NotificationType: "group_events",
			CreatedAt:        time.Now(),
		}

		// should be added to group event attendance as false?

		notificationId, err := s.NotificationRepository.Insert(notification)

		if err != nil {
			s.Logger.Printf("Failed inserting notification: %s", err)
		}

		// broadcast notification to all users

		notificationJSON := &models.NotificationJSON{
			NotificationType: "event_invite",
			NotificationId:   notificationId,
			SenderId:         userId,
			SenderName:       userData.Nickname,
			GroupId:          formData.GroupId,
			GroupName:        groupData.Title,
			EventId:          result,
			EventName:        formData.Title,
			EventDate:        formData.EventTime,
		}

		s.Logger.Printf("Broadcasting notification: %v", notificationJSON)

		// TODO: broadcast by WS to member

	}

	return result, err
}

func (s *GroupEventService) GetUserEvents(userId int64) ([]*EventJSON, error) {

	events, err := s.EventRepository.GetAllByUserId(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching user events: %s", err)
		return nil, err
	}

	var eventJSON []*EventJSON

	for _, event := range events {

		groupName, err := s.GroupRepository.GetById(event.GroupId)
		if err != nil {
			s.Logger.Printf("Failed fetching group name: %s", err)
			return nil, err
		}

		userData, err := s.UserRepository.GetById(event.UserId)
		if err != nil {
			s.Logger.Printf("Failed fetching user data: %s", err)
			return nil, err
		}

		if userData.Nickname == "" {
			userData.Nickname = userData.FirstName + " " + userData.LastName
		}

		eventJSON = append(eventJSON, &EventJSON{
			Id:          int(event.Id),
			GroupId:     int(event.GroupId),
			GroupName:   groupName.Title,
			UserId:      int(event.UserId),
			NickName:    userData.Nickname,
			CreatedAt:   event.CreatedAt,
			EventTime:   event.EventTime,
			TimeSpan:    event.TimeSpan,
			Title:       event.Title,
			Description: event.Description,
		})
	}

	return eventJSON, nil
}
