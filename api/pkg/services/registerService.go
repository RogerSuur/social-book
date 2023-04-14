package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

func (s *Service) UserRegister(user *models.User) (string, error) {

	env := models.CreateEnv(s.DB)

	// check if user exists
	_, err := env.Users.GetByEmail(user.Email)
	if err == nil {
		log.Printf("User email already exists")
		return "", errors.New("user email already exists")
	}

	// hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Cannot hash password: %s", err)
		return "", errors.New("cannot hash password")
	}
	user.Password = hashedPassword

	// create user
	lastID, err := env.Users.Insert(user)
	if err != nil {
		log.Printf("Cannot create user: %s", err)
		return "", errors.New("cannot create user")
	}
	fmt.Println("Last inserted ID:", lastID)

	// create session
	sessionToken := uuid.NewV4().String()
	session := models.Session{
		UserId: user.Id,
		Token:  sessionToken,
	}

	// store session in DB
	lastID, err = env.Sessions.Insert(&session)
	if err != nil {
		log.Printf("Cannot create session: %s", err)
		return "", errors.New("cannot create session")
	}
	fmt.Println("Last inserted ID:", lastID)
	return sessionToken, nil
}
