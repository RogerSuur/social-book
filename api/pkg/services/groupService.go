package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupService interface {
	GetUserGroups(userId int64) ([]*models.UserGroup, error)
	GetUserCreatedGroups(userId int64) ([]*models.UserGroup, error)
	GetGroupById(groupId int64) (models.GroupJSON, error)
	SearchGroupsAndUsers(searchString string) ([]*models.SearchResult, error)
	CreateGroup(groupFormData *models.GroupJSON, userId int64) (int64, error)
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

func (s *GroupService) GetUserGroups(userId int64) ([]*models.UserGroup, error) {

	result, err := s.GroupRepository.GetAllByMemberId(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching groups: %s", err)
	}

	groups := []*models.UserGroup{}

	for _, p := range result {
		groups = append(groups, &models.UserGroup{
			Id:    p.Id,
			Title: p.Title,
		})
	}

	return groups, nil
}

func (s *GroupService) GetUserCreatedGroups(userId int64) ([]*models.UserGroup, error) {

	result, err := s.GroupRepository.GetAllByCreatorId(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching groups: %s", err)
	}

	groups := []*models.UserGroup{}

	for _, p := range result {
		groups = append(groups, &models.UserGroup{
			Id:    p.Id,
			Title: p.Title,
		})
	}

	return groups, nil
}

func (s *GroupService) GetGroupById(groupId int64) (models.GroupJSON, error) {
	result, err := s.GroupRepository.GetById(groupId)

	group := models.GroupJSON{
		Title:       result.Title,
		Description: result.Description,
		ImagePath:   result.ImagePath,
	}

	if err != nil {
		s.Logger.Printf("Failed fetching groups: %s", err)
	}

	return group, err
}

func (s *GroupService) SearchGroupsAndUsers(searchString string) ([]*models.SearchResult, error) {

	result, err := s.GroupRepository.SearchGroupsAndUsersByString(searchString)

	if err != nil {
		s.Logger.Printf("Failed searching groups: %s", err)
	}
	//TODO
	return result, err
}

func (s *GroupService) CreateGroup(groupFormData *models.GroupJSON, userId int64) (int64, error) {
	group := &models.Group{
		CreatorId:   userId,
		ImagePath:   groupFormData.ImagePath,
		Title:       groupFormData.Title,
		Description: groupFormData.Description,
	}

	result, err := s.GroupRepository.Insert(group)

	if err != nil {
		s.Logger.Printf("Failed inserting group: %s", err)
	}

	return result, err
}
