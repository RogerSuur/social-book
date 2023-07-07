package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupEventService interface {
	GetGroupEvents(groupId int64) ([]*models.Event, error)
	CreateGroupEvent(formData *models.CreateGroupEventFormData, userId int64) (int64, error)
}

type GroupEventService struct {
	Logger                         *log.Logger
	GroupEventAttendanceRepository models.IGroupEventAttendanceRepository
	EventRepository                models.IEventRepository
}

func InitGroupEventService(logger *log.Logger, groupEventAttendanceRepo *models.GroupEventAttendanceRepository, groupEventRepo *models.EventRepository) *GroupEventService {
	return &GroupEventService{
		Logger:                         logger,
		GroupEventAttendanceRepository: groupEventAttendanceRepo,
		EventRepository:                groupEventRepo,
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

	return result, err
}
