package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupMemberService interface {
	GetGroupMembers(groupId int64) ([]*models.User, error)
}

type GroupMemberService struct {
	Logger                *log.Logger
	GroupMemberRepository models.IGroupUserRepository
}

func InitGroupMemberService(logger *log.Logger, groupMemberRepo *models.GroupUserRepository) *GroupMemberService {
	return &GroupMemberService{
		Logger:                logger,
		GroupMemberRepository: groupMemberRepo,
	}
}

func (s *GroupMemberService) GetGroupMembers(groupId int64) ([]*models.User, error) {

	members, err := s.GroupMemberRepository.GetGroupMembersByGroupId(groupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	return members, nil
}
