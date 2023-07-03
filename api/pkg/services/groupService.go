package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupService interface {
	GetUserGroups(userId int) ([]*models.UserGroup, error)
	GetUserCreatedGroups(userId int) ([]*models.UserGroup, error)
}

type GroupService struct {
	Logger          *log.Logger
	GroupRepository models.IGroupRepository
}

func InitGroupService(logger *log.Logger, groupRepo *models.GroupRepository) *GroupService {
	return &GroupService{
		Logger:          logger,
		GroupRepository: groupRepo,
	}
}

func (s *GroupService) GetUserGroups(userId int) ([]*models.UserGroup, error) {

	result, err := s.GroupRepository.GetAllByMemberId(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching groups: %s", err)
	}

	groups := []*models.UserGroup{}

	for _, p := range result {
		groups = append(groups, &models.UserGroup{
			Id:   p.Id,
			Name: p.Title,
		})
	}

	return groups, nil
}

func (s *GroupService) GetUserCreatedGroups(userId int) ([]*models.UserGroup, error) {

	result, err := s.GroupRepository.GetAllByCreatorId(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching groups: %s", err)
	}

	groups := []*models.UserGroup{}

	for _, p := range result {
		groups = append(groups, &models.UserGroup{
			Id:   p.Id,
			Name: p.Title,
		})
	}

	return groups, nil
}
