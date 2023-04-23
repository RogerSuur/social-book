package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"net/http"
)

type IUserService interface {
	Authenticate(handler http.HandlerFunc) http.HandlerFunc
	CreateUser(user *models.User) (int64, error)
	GetUserData(userID int64) (*models.User, error)
	GetUserID(r *http.Request) (int, error)
	SetCookie(w http.ResponseWriter, sessionToken string)
	UserLogin(user *models.User) (string, error)
	UserRegister(user *models.User) (string, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserService struct {
	UserRepo    models.IUserRepository
	SessionRepo models.ISessionRepository
}

// InitUserService initializes the user controller.
func InitUserService(userRepo *models.UserRepository, sessionRepo *models.SessionRepository) *UserService {
	return &UserService{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	// env := models.CreateEnv(s.DB)

	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB

	return s.UserRepo.Insert(user)
	// return env.Users.Insert(user)
}

func (s *UserService) GetUserData(userID int64) (*models.User, error) {
	// env := models.CreateEnv(s.DB)
	// return env.Users.GetById(userID)

	return s.UserRepo.GetById(userID)
}
