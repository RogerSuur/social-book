package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupEventService interface {
	GetGroupEvents(groupId int64) ([]*models.Event, error)
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
