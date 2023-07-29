package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"time"
)

type IGroupMemberService interface {
	GetGroupMembers(groupId int64) ([]*models.SimpleUserJSON, error)
	IsGroupMember(groupId int64, userId int64) (bool, error)
	AddMembers(userId int64, members models.GroupMemberJSON) ([]*models.NotificationJSON, error)
	GetMembersToAdd(groupId int64, userId int64) ([]*models.SimpleUserJSON, error)
}

type GroupMemberService struct {
	Logger                 *log.Logger
	UserRepository         models.IUserRepository
	NotificationRepository models.INotificationRepository
	GroupMemberRepository  models.IGroupMemberRepository
}

func InitGroupMemberService(
	logger *log.Logger,
	userRepo *models.UserRepository,
	notificationsRepo *models.NotificationRepository,
	groupMemberRepo *models.GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{
		Logger:                 logger,
		UserRepository:         userRepo,
		NotificationRepository: notificationsRepo,
		GroupMemberRepository:  groupMemberRepo,
	}
}

func (s *GroupMemberService) GetGroupMembers(groupId int64) ([]*models.SimpleUserJSON, error) {

	members, err := s.GroupMemberRepository.GetGroupMembersByGroupId(groupId)

	if err != nil {
		s.Logger.Printf("Failed fetching group members: %s", err)
		return nil, err
	}

	simpleMembers := []*models.SimpleUserJSON{}

	for _, member := range members {

		if !member.Accepted {
			continue
		}

		userData, err := s.UserRepository.GetById(member.UserId)

		if err != nil {
			s.Logger.Printf("Failed fetching user data: %s", err)
			return nil, err
		}

		if userData.Nickname == "" {
			userData.Nickname = userData.FirstName + " " + userData.LastName
		}

		simpleMember := &models.SimpleUserJSON{
			Id:        int(member.UserId),
			Nickname:  userData.Nickname,
			ImagePath: userData.ImagePath,
		}

		simpleMembers = append(simpleMembers, simpleMember)
	}

	return simpleMembers, nil
}

func (s *GroupMemberService) IsGroupMember(groupId int64, userId int64) (bool, error) {
	isgroupMember, err := s.GroupMemberRepository.IsGroupMember(groupId, userId)

	if err != nil {
		s.Logger.Printf("Cannot validate user: %s", err)
	}

	s.Logger.Printf("User %d is group member: %t", userId, isgroupMember)

	return isgroupMember, err
}

func (s *GroupMemberService) AddMembers(userId int64, members models.GroupMemberJSON) ([]*models.NotificationJSON, error) {

	isGroupMember, err := s.IsGroupMember(int64(members.GroupId), userId)

	if err != nil {
		s.Logger.Printf("Cannot validate user: %s", err)
		return nil, err
	}

	if !isGroupMember {
		s.Logger.Printf("User %d is not a member of this group", userId)
		return nil, errors.New("not a member of this group")
	}

	notificationDetails := &models.NotificationDetails{
		SenderId:         userId,
		NotificationType: "group_invite",
		EntityId:         int64(members.GroupId),
		CreatedAt:        time.Now(),
	}

	detailsId, err := s.NotificationRepository.InsertDetails(notificationDetails)

	if err != nil {
		s.Logger.Printf("Cannot insert notification details: %s", err)
		return nil, err
	}

	notificationsToBroadcast := []*models.NotificationJSON{}

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
			Accepted: false,
		}

		_, err = s.GroupMemberRepository.Insert(groupMember)

		if err != nil {
			s.Logger.Printf("Cannot add user %d to group %d: %s", userIdToAdd, members.GroupId, err)
			return nil, err
		}

		s.Logger.Printf("User %d added to group %d", userIdToAdd, members.GroupId)

		// send notification to user

		notification := &models.Notification{
			ReceiverId:            int64(userIdToAdd),
			NotificationDetailsId: detailsId,
		}

		notificationId, err := s.NotificationRepository.InsertNotification(notification)

		if err != nil {
			s.Logger.Printf("Cannot insert notification: %s", err)
			return nil, err
		}

		// broadcast notification to users

		notificationJSON := &models.NotificationJSON{
			ReceiverId:       int64(userIdToAdd),
			NotificationType: notificationDetails.NotificationType,
			NotificationId:   notificationId,
			SenderId:         userId,
			SenderName:       "",
			GroupId:          int64(members.GroupId),
			GroupName:        "",
		}

		notificationsToBroadcast = append(notificationsToBroadcast, notificationJSON)

	}

	return notificationsToBroadcast, nil
}

func (s *GroupMemberService) GetMembersToAdd(groupId int64, userId int64) ([]*models.SimpleUserJSON, error) {

	followers, err := s.UserRepository.GetAllUserFollowers(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching user followers: %s", err)
		return nil, err
	}

	simpleMembers := map[int64]*models.SimpleUserJSON{}

	for _, follower := range followers {

		isGroupMember, err := s.IsGroupMember(groupId, follower.Id)

		if err != nil {
			s.Logger.Printf("Cannot validate user: %s", err)
			return nil, err
		}

		if isGroupMember {
			continue
		}

		simpleMember := &models.SimpleUserJSON{
			Id:        int(follower.Id),
			Nickname:  follower.Nickname,
			ImagePath: follower.ImagePath,
		}

		simpleMembers[follower.Id] = simpleMember
	}

	following, err := s.UserRepository.GetAllFollowedBy(userId)

	if err != nil {
		s.Logger.Printf("Failed fetching users following: %s", err)
		return nil, err
	}

	for _, followed := range following {

		if simpleMembers[followed.Id] != nil {
			continue
		}

		isGroupMember, err := s.IsGroupMember(groupId, followed.Id)

		if err != nil {
			s.Logger.Printf("Cannot validate user: %s", err)
			return nil, err
		}

		if isGroupMember {
			continue
		}

		simpleMember := &models.SimpleUserJSON{
			Id:        int(followed.Id),
			Nickname:  followed.Nickname,
			ImagePath: followed.ImagePath,
		}

		simpleMembers[followed.Id] = simpleMember
	}

	simpleMembersArray := make([]*models.SimpleUserJSON, 0, len(simpleMembers))

	for _, simpleMember := range simpleMembers {
		simpleMembersArray = append(simpleMembersArray, simpleMember)
	}

	return simpleMembersArray, nil
}
