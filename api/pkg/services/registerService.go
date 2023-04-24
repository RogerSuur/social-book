package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"
	"regexp"

	uuid "github.com/satori/go.uuid"
)

func (s *Service) UserRegister(user *models.User) (string, error) {

	env := models.CreateEnv(s.DB)

	_, err := env.Users.GetByEmail(user.Email)
	if err == nil {
		log.Printf("User email already exists")
		return "", errors.New("email")
	}

	if len(user.Nickname) > 0 {
		err = env.Users.CheckIfNicknameExists(user.Nickname, 0)
		if err == nil {
			log.Printf("User nickname already exists")
			return "", errors.New("nickname")
		}
	}

	if !s.CheckPasswordStrength(user.Password) {
		log.Printf("Password is not strong enough")
		return "", errors.New("password")
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Cannot hash password: %s", err)
		return "", errors.New("server")
	}
	user.Password = hashedPassword

	lastID, err := env.Users.Insert(user)
	if err != nil {
		log.Printf("Cannot create user: %s", err)
		return "", errors.New("server")
	}
	fmt.Println("User succesfully created, ID", lastID)

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
		return "", errors.New("server")
	}
	fmt.Println("Last inserted ID:", lastID)
	return sessionToken, nil
}

func (s *Service) CheckPasswordStrength(password string) bool {
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	numberRegex := regexp.MustCompile(`[0-9]`)
	symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

	strong := lowercaseRegex.MatchString(password) && uppercaseRegex.MatchString(password) && numberRegex.MatchString(password) && symbolRegex.MatchString(password)

	return strong
}
