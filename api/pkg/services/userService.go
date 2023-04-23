package services

import (
	"SocialNetworkRestApi/api/pkg/models"
)

func (s *Service) CreateUser(user *models.User) (int64, error) {
	env := models.CreateEnv(s.DB)

	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB
	return env.Users.Insert(user)
}

func (s *Service) GetUserData(userID int64) (*models.User, error) {
	env := models.CreateEnv(s.DB)
	return env.Users.GetById(userID)
}

type FollowerData struct {
	UserID      int
	FirstName   string
	LastName    string
	Nickname    string
	AvatarImage string
	Accepted    bool
}

func (s *Service) GetUserFollowers(userID int64) ([]FollowerData, error) {
	env := models.CreateEnv(s.DB)
	followers, err := env.Followers.GetFollowersById(userID)
	if err != nil {
		return nil, err
	}

	followersData := []FollowerData{}

	for _, follower := range followers {
		user, err := env.Users.GetById(int64(follower.FollowerId))
		if err != nil {
			return nil, err
		}
		followersData = append(followersData, FollowerData{
			UserID:      user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Nickname:    user.Nickname,
			AvatarImage: user.ImagePath,
			Accepted:    follower.Accepted,
		})
	}

	return followersData, nil
}

func (s *Service) GetUserFollowing(userID int64) ([]FollowerData, error) {
	env := models.CreateEnv(s.DB)
	following, err := env.Followers.GetFollowingById(userID)
	if err != nil {
		return nil, err
	}

	followingData := []FollowerData{}

	for _, follower := range following {
		user, err := env.Users.GetById(int64(follower.FollowingId))
		if err != nil {
			return nil, err
		}
		followingData = append(followingData, FollowerData{
			UserID:      user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Nickname:    user.Nickname,
			AvatarImage: user.ImagePath,
			Accepted:    follower.Accepted,
		})
	}

	return followingData, nil
}
