package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"time"
)

type IGroupMemberService interface {
	GetGroupMembers(groupId int64) ([]*models.User, error)
	IsGroupMember(groupId int64, userId int64) (bool, error)
	AddMembers(userId int64, members models.GroupMemberJSON) (*models.GroupMemberJSON, error)
}

type GroupMemberService struct {
	Logger                *log.Logger
	GroupMemberRepository models.IGroupMemberRepository
}

func InitGroupMemberService(logger *log.Logger, groupMemberRepo *models.GroupMemberRepository) *GroupMemberService {
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

func (s *GroupMemberService) IsGroupMember(groupId int64, userId int64) (bool, error) {
	isgroupMember, err := s.GroupMemberRepository.IsGroupMember(groupId, userId)

	if err != nil {
		s.Logger.Printf("Cannot validate user: %s", err)
	}

	return isgroupMember, err
}

func (s *GroupMemberService) AddMembers(userId int64, members models.GroupMemberJSON) (*models.GroupMemberJSON, error) {

	isGroupMember, err := s.IsGroupMember(int64(members.GroupId), userId)

	if err != nil {
		s.Logger.Printf("Cannot validate user: %s", err)
		return nil, err
	}

	if !isGroupMember {
		s.Logger.Printf("User %d is not a member of this group", userId)
		return nil, errors.New("not a member of this group")
	}

	addedMembers := make([]int, 0)

	for _, userIdToAdd := range members.UserIds {

		isGroupMember, err := s.IsGroupMember(int64(members.GroupId), int64(userIdToAdd))

		if err != nil {
			s.Logger.Printf("Cannot validate user: %s", err)
			return nil, err
		}

		if isGroupMember {
			s.Logger.Printf("User %d is already a member of this group", userIdToAdd)
			continue
		}

		groupMember := &models.GroupMember{
			UserId:   int64(userIdToAdd),
			GroupId:  int64(members.GroupId),
			JoinedAt: time.Now(),
		}

		_, err = s.GroupMemberRepository.Insert(groupMember)

		if err != nil {
			s.Logger.Printf("Cannot add user %d to group %d: %s", userIdToAdd, members.GroupId, err)
			return nil, err
		}

		addedMembers = append(addedMembers, userIdToAdd)

		s.Logger.Printf("User %d added to group %d", userIdToAdd, members.GroupId)
	}

	result := &models.GroupMemberJSON{
		GroupId: members.GroupId,
		UserIds: addedMembers,
	}

	return result, nil
}
