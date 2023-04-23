package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

func (s *UserService) UserLogin(user *models.User) (string, error) {

	// env := models.CreateEnv(s.DB)

	// check if user exists
	dbUser, err := s.UserRepo.GetByEmail(user.Email)
	if err != nil {
		log.Printf("User email not found: %s", err)
		return "", errors.New("user email not found")
	}

	// check if password is correct
	if !CheckPasswordHash(user.Password, dbUser.Password) {
		log.Printf("Incorrect password")
		return "", errors.New("incorrect password")
	}

	// create session
	sessionToken := uuid.NewV4().String()
	session := models.Session{
		UserId: dbUser.Id,
		Token:  sessionToken,
	}

	// store session in DB
	lastID, err := s.SessionRepo.Insert(&session)
	if err != nil {
		log.Printf("Cannot create session: %s", err)
		return "", errors.New("cannot create session")
	}
	fmt.Println("Last inserted ID:", lastID)
	return sessionToken, nil
}
