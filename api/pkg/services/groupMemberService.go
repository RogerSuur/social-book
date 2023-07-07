package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
)

type IGroupMemberService interface {
	GetGroupMembers(groupId int64) ([]*models.GroupMember, error)
	IsGroupMember(groupId int64, userId int64) (bool, error)
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

func (s *GroupMemberService) GetGroupMembers(groupId int64) ([]*models.GroupMember, error) {

	members, err := s.GroupMemberRepository.GetGroupMembersByGroupId(groupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	return members, nil
}

func (s *GroupMemberService) IsGroupMember(groupId int64, userId int64) (bool, error) {
	isgroupMember, err := s.GroupMemberRepository.IsGroupMember(groupId, userId)

	if err != nil {
		s.Logger.Printf("Cannot validate user: %s", err)
	}

	return isgroupMember, err
}
