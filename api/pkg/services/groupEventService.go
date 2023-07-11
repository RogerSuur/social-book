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
	GetUserEvents(userId int64) ([]*EventJSON, error)
}

type GroupEventService struct {
	Logger                         *log.Logger
	GroupEventAttendanceRepository models.IGroupEventAttendanceRepository
	EventRepository                models.IEventRepository
	GroupRepository                models.IGroupRepository
	UserRepository                 models.IUserRepository
}

func InitGroupEventService(
	logger *log.Logger,
	groupEventAttendanceRepo *models.GroupEventAttendanceRepository,
	groupEventRepo *models.EventRepository,
	groupRepo *models.GroupRepository,
	userRepo *models.UserRepository,
) *GroupEventService {
	return &GroupEventService{
		Logger:                         logger,
		GroupEventAttendanceRepository: groupEventAttendanceRepo,
		EventRepository:                groupEventRepo,
		GroupRepository:                groupRepo,
		UserRepository:                 userRepo,
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
