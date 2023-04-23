package services

import (
	"SocialNetworkRestApi/api/pkg/models"
)

// type Service struct {
// 	DB  *sql.DB
// 	Env models.Env
// }

// Services contains all the controllers
type Services struct {
	UserService IUserService
	PostService IPostService
}

// InitServices returns a new Controllers
func InitServices(repositories *models.Repositories) *Services {
	return &Services{
		UserService: InitUserService(repositories.UserRepo, repositories.SessionRepo),
	}
}
