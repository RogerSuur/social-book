package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"time"
)

type EventJSON struct {
	Id           int64                  `json:"id"`
	GroupId      int64                  `json:"groupId"`
	GroupName    string                 `json:"groupName"`
	UserId       int64                  `json:"creatorId"`
	NickName     string                 `json:"creatorName"`
	CreatedAt    time.Time              `json:"createdAt"`
	EventTime    time.Time              `json:"eventTime"`
	EventEndTime time.Time              `json:"eventEndTime"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Members      []*models.AttendeeJSON `json:"members"`
}

type IGroupEventService interface {
	GetGroupEvents(groupId int64) ([]*EventJSON, error)
	CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) ([]*models.NotificationJSON, error)
	GetUserEvents(userId int64) ([]*EventJSON, error)
	GetEventById(eventId int64) (*EventJSON, error)
	ParseEventJSON(events []*models.Event) ([]*EventJSON, error)
}

type GroupEventService struct {
	Logger                    *log.Logger
	EventAttendanceRepository models.IEventAttendanceRepository
	EventRepository           models.IEventRepository
	GroupRepository           models.IGroupRepository
	GroupMemberRepository     models.IGroupMemberRepository
	UserRepository            models.IUserRepository
	NotificationRepository    models.INotificationRepository
}

func InitGroupEventService(
	logger *log.Logger,
	eventAttendanceRepo *models.EventAttendanceRepository,
	groupEventRepo *models.EventRepository,
	groupRepo *models.GroupRepository,
	GroupMemberRepository *models.GroupMemberRepository,
	userRepo *models.UserRepository,
	notificationRepo *models.NotificationRepository,
) *GroupEventService {
	return &GroupEventService{
		Logger:                    logger,
		EventAttendanceRepository: eventAttendanceRepo,
		EventRepository:           groupEventRepo,
		GroupRepository:           groupRepo,
		GroupMemberRepository:     GroupMemberRepository,
		UserRepository:            userRepo,
		NotificationRepository:    notificationRepo,
	}
}

func (s *GroupEventService) GetGroupEvents(groupId int64) ([]*EventJSON, error) {

	events, err := s.EventRepository.GetAllByGroupId(groupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	s.Logger.Printf("Fetched %d events", len(events))

	eventJSON, err := s.ParseEventJSON(events)
	if err != nil {
		s.Logger.Printf("Failed parsing event json: %s", err)
		return nil, err
	}

	return eventJSON, nil
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

	s.Logger.Printf("Fetched %d events", len(events))

	eventJSON, err := s.ParseEventJSON(events)
	if err != nil {
		s.Logger.Printf("Failed parsing event json: %s", err)
		return nil, err
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

func (s *GroupEventService) GetEventById(eventId int64) (*EventJSON, error) {

	event, err := s.EventRepository.GetById(eventId)

	if err != nil {
		s.Logger.Printf("Failed fetching event: %s", err)
		return nil, err
	}

	attendees, err := s.EventAttendanceRepository.GetAttendeesByEventId(eventId)

	if err != nil {
		s.Logger.Printf("Failed fetching event attendance: %s", err)
		return nil, err
	}

	attendeesJSON := []*models.AttendeeJSON{}

	for _, attendee := range attendees {

		singleJSON := &models.AttendeeJSON{}

		user, err := s.UserRepository.GetById(attendee.UserId)
		if err != nil {
			s.Logger.Printf("Failed fetching user: %s", err)
			return nil, err
		}

		if user.Nickname == "" {
			user.Nickname = user.FirstName + " " + user.LastName
		}

		singleJSON.Id = int(attendee.UserId)
		singleJSON.Nickname = user.Nickname
		singleJSON.ImagePath = user.ImagePath
		singleJSON.IsAttending = attendee.IsAttending

		attendeesJSON = append(attendeesJSON, singleJSON)
	}

	group, err := s.GroupRepository.GetById(event.GroupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group: %s", err)
		return nil, err
	}

	eventJSON := &EventJSON{
		Id:           event.Id,
		GroupId:      event.GroupId,
		GroupName:    group.Title,
		CreatedAt:    event.CreatedAt,
		EventTime:    event.EventTime,
		EventEndTime: event.EventEndTime,
		Title:        event.Title,
		Description:  event.Description,
		Members:      attendeesJSON,
	}

	//s.Logger.Printf("Fetched event: %v", eventJSON)

	return eventJSON, nil
}

func (s *GroupEventService) ParseEventJSON(events []*models.Event) ([]*EventJSON, error) {

	var eventJSON []*EventJSON

	for _, event := range events {

		groupData, err := s.GroupRepository.GetById(event.GroupId)
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
			GroupName:    groupData.Title,
			UserId:       event.UserId,
			NickName:     userData.Nickname,
			CreatedAt:    event.CreatedAt,
			EventTime:    event.EventTime,
			EventEndTime: event.EventEndTime,
			Title:        event.Title,
			Description:  event.Description,
		})
	}

	//s.Logger.Println("Parsed events", eventJSON)

	return eventJSON, nil
}
