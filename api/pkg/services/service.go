package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"database/sql"
)

type Service struct {
	DB *sql.DB
}

// Services contains all the controllers
type Services struct {
	UserService IUserService
}

// InitServices returns a new Controllers
func InitServices(repositories *models.Repositories) *Services {
	return &Services{
		UserService: InitUserService(repositories.UserRepo, repositories.SessionRepo),
	}
}
