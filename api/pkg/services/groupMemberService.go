package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupMemberService interface {
	GetGroupMembers(groupId int64, userId int64) ([]*models.GroupMember, error)
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

func (s *GroupMemberService) GetGroupMembers(groupId int64, userId int64) ([]*models.GroupMember, error) {

	isGroupMember, err := s.GroupMemberRepository.IsGroupMember(groupId, userId)

	if err != nil {
		s.Logger.Printf("Failed fetching group member status: %s", err)
		return nil, err
	}

	if !isGroupMember {
		return nil, nil
	}

	members, err := s.GroupMemberRepository.GetGroupMembersByGroupId(groupId)
	s.Logger.Println(members)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	return members, nil
}
