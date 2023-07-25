package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"time"
)

type EventJSON struct {
	Id           int64     `json:"id"`
	GroupId      int64     `json:"group_id"`
	GroupName    string    `json:"group_name"`
	UserId       int64     `json:"creator_id"`
	NickName     string    `json:"creator_name"`
	CreatedAt    time.Time `json:"created_at"`
	EventTime    time.Time `json:"event_time"`
	EventEndTime time.Time `json:"event_end_time"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
}

type IGroupEventService interface {
	GetGroupEvents(groupId int64) ([]*models.Event, error)
	CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) ([]*models.NotificationJSON, error)
	GetUserEvents(userId int64) ([]*EventJSON, error)
}

type GroupEventService struct {
	Logger                         *log.Logger
	GroupEventAttendanceRepository models.IEventAttendanceRepository
	EventRepository                models.IEventRepository
	GroupRepository                models.IGroupRepository
	GroupMemberRepository          models.IGroupMemberRepository
	UserRepository                 models.IUserRepository
	NotificationRepository         models.INotificationRepository
}

func InitGroupEventService(
	logger *log.Logger,
	groupEventAttendanceRepo *models.EventAttendanceRepository,
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

	s.Logger.Printf("Fetched %d events", len(events))

	return events, nil
}

func (s *GroupEventService) CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) ([]*models.NotificationJSON, error) {

	s.Logger.Printf("Event timestring: %s", formData.EventTime)
	sTime, err := time.Parse("2006-01-02T15:04", formData.EventTime)

	if err != nil {
		s.Logger.Printf("Failed parsing event start time: %s", err)
	}

	eTime, err := time.Parse("2006-01-02T15:04", formData.EventEndTime)
	if err != nil {
		s.Logger.Printf("Failed parsing event start time: %s", err)
	}

	event := &models.Event{
		GroupId:      int64(formData.GroupId),
		UserId:       userId,
		EventTime:    sTime,
		EventEndTime: eTime,
		Title:        formData.Title,
		Description:  formData.Description,
	}

	result, err := s.EventRepository.Insert(event)

	if err != nil {
		s.Logger.Printf("Failed inserting event: %s", err)
	}

	// Send notification to all group members

	groupMembers, err := s.GroupMemberRepository.GetGroupMembersByGroupId(int64(formData.GroupId))

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	userData, err := s.UserRepository.GetById(userId)
	if err != nil {
		s.Logger.Printf("Failed fetching user data: %s", err)
		return nil, err
	}

	if userData.Nickname == "" {
		userData.Nickname = userData.FirstName + " " + userData.LastName
	}

	groupData, err := s.GroupRepository.GetById(int64(formData.GroupId))
	if err != nil {
		s.Logger.Printf("Failed fetching group data: %s", err)
		return nil, err
	}

	notificationsToBroadcast := []*models.NotificationJSON{}

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
			ReceiverId:       member.Id,
			NotificationType: "event_invite",
			NotificationId:   notificationId,
			SenderId:         userId,
			SenderName:       userData.Nickname,
			GroupId:          int64(formData.GroupId),
			GroupName:        groupData.Title,
			EventId:          result,
			EventName:        formData.Title,
			// EventDate:        formData.EventTime,
		}

		s.Logger.Printf("Broadcasting notification: %v", notificationJSON)

		notificationsToBroadcast = append(notificationsToBroadcast, notificationJSON)

	}

	return notificationsToBroadcast, err
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
			Id:           event.Id,
			GroupId:      event.GroupId,
			GroupName:    groupName.Title,
			UserId:       event.UserId,
			NickName:     userData.Nickname,
			CreatedAt:    event.CreatedAt,
			EventTime:    event.EventTime,
			EventEndTime: event.EventEndTime,
			Title:        event.Title,
			Description:  event.Description,
		})
	}

	return eventJSON, nil
}

func (s *GroupEventService) InviteUser(userId int64, eventId int64) error {

	event, err := s.EventRepository.GetById(eventId)

	if err != nil {
		s.Logger.Printf("Failed fetching event: %s", err)
		return err
	}

	notification := &models.Notification{
		ReceiverId:       userId,
		SenderId:         event.UserId,
		EntityId:         eventId,
		NotificationType: "event_invite",
		CreatedAt:        time.Now(),
	}

	_, err = s.NotificationRepository.Insert(notification)

	if err != nil {
		s.Logger.Printf("Failed inserting notification: %s", err)
		return err
	}

	return nil
}