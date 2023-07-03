package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// Services contains all the controllers
type Services struct {
	UserService         IUserService
	NotificationService INotificationService
	PostService         IPostService
	CommentService      ICommentService
	ChatService         IChatService
}

// InitServices returns a new Controllers
func InitServices(repositories *models.Repositories, logger *log.Logger) *Services {
	return &Services{
		UserService: InitUserService(
			repositories.UserRepo,
			repositories.SessionRepo,
			repositories.FollowerRepo,
			repositories.NotificationRepo,
		),
		NotificationService: InitNotificationService(
			repositories.UserRepo,
			repositories.FollowerRepo,
			repositories.NotificationRepo,
		),
		PostService:    InitPostService(logger, repositories.PostRepo, repositories.AllowedPostRepo),
		CommentService: InitCommentService(repositories.CommentRepo),
		ChatService: InitChatService(
			repositories.UserRepo,
			repositories.MessageRepo,
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

func CheckPasswordStrength(password string) bool {
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	numberRegex := regexp.MustCompile(`[0-9]`)
	symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

	strong := lowercaseRegex.MatchString(password) && uppercaseRegex.MatchString(password) && numberRegex.MatchString(password) && symbolRegex.MatchString(password)

	return strong
}
