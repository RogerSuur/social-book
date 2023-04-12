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
