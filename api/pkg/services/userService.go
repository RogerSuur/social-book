package services

import (
	"SocialNetworkRestApi/api/pkg/models"
)

type UserService struct {
	user *models.UserModel
}

func (u *UserService) CreateUser(user *models.User) (int64, error) {
	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB
	return u.user.Insert(user)
}
