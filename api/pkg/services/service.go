package services

import (
	"SocialNetworkRestApi/api/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

// Services contains all the controllers
type Services struct {
	UserService IUserService
}

// InitServices returns a new Controllers
func InitServices(repositories *models.Repositories) *Services {
	return &Services{
		UserService: InitUserService(
			repositories.UserRepo,
			repositories.SessionRepo,
		),
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
