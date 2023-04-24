package services

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

type ProfileJSON struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
	Nickname    string `json:"nickname"`
	About       string `json:"about"`
	AvatarImage string `json:"avatarImage"`
	CreatedAt   string `json:"createdAt"`
	IsPublic    bool   `json:"isPublic"`
}

type IUserService interface {
	Authenticate(handler http.HandlerFunc) http.HandlerFunc
	CreateUser(user *models.User) (int64, error)
	GetUserData(userID int64) (*ProfileJSON, error)
	GetUserID(r *http.Request) (int, error)
	SetCookie(w http.ResponseWriter, sessionToken string)
	UserLogin(user *models.User) (string, error)
	UserRegister(user *models.User) (string, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserService struct {
	Logger      *log.Logger
	UserRepo    models.IUserRepository
	SessionRepo models.ISessionRepository
}

// InitUserService initializes the user controller.
func InitUserService(userRepo *models.UserRepository, sessionRepo *models.SessionRepository) *UserService {
	return &UserService{
		Logger:      log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB

	return s.UserRepo.Insert(user)
}

func (s *UserService) GetUserData(userID int64) (*ProfileJSON, error) {
	user := &models.User{}
	user, err := s.UserRepo.GetById(userID)
	if err != nil {
		return nil, err
	}
	userJSON := &ProfileJSON{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Birthday:    user.Birthday.Format("02/01/2006"),
		Nickname:    user.Nickname,
		About:       user.About,
		AvatarImage: user.ImagePath,
		CreatedAt:   user.CreatedAt.Format("02/01/2006 15:04:05"),
		IsPublic:    user.IsPublic,
	}

	return userJSON, nil
}

func (s *UserService) UserRegister(user *models.User) (string, error) {

	// check if user exists
	_, err := s.UserRepo.GetByEmail(user.Email)
	if err == nil {
		s.Logger.Printf("User email already exists")
		return "", errors.New("user email already exists")
	}

	// hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		s.Logger.Printf("Cannot hash password: %s", err)
		return "", errors.New("cannot hash password")
	}
	user.Password = hashedPassword

	// create user
	lastID, err := s.UserRepo.Insert(user)
	if err != nil {
		s.Logger.Printf("Cannot create user: %s", err)
		return "", errors.New("cannot create user")
	}
	s.Logger.Println("Last inserted ID:", lastID)

	// create session
	sessionToken := uuid.NewV4().String()
	session := models.Session{
		UserId: user.Id,
		Token:  sessionToken,
	}

	// store session in DB
	lastID, err = s.SessionRepo.Insert(&session)
	if err != nil {
		s.Logger.Printf("Cannot create session: %s", err)
		return "", errors.New("cannot create session")
	}
	s.Logger.Println("Last inserted ID:", lastID)
	return sessionToken, nil
}

func (s *UserService) UserLogin(user *models.User) (string, error) {

	// check if user exists
	dbUser, err := s.UserRepo.GetByEmail(user.Email)
	if err != nil {
		s.Logger.Printf("User email not found: %s", err)
		return "", errors.New("user email not found")
	}

	// check if password is correct
	if !CheckPasswordHash(user.Password, dbUser.Password) {
		s.Logger.Printf("Incorrect password")
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
		s.Logger.Printf("Cannot create session: %s", err)
		return "", errors.New("cannot create session")
	}
	s.Logger.Println("Last inserted ID:", lastID)
	return sessionToken, nil
}

func (s *UserService) Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetCors(&w, r)

		// check if cookie exists
		cookie, err := r.Cookie("session")
		if err != nil {
			s.Logger.Printf("No cookie found: %s", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// check if session exists
		_, err = s.SessionRepo.GetByToken(cookie.Value)

		if err != nil {
			s.Logger.Printf("No session found: %s", err)
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		// finally, call the handler
		handler.ServeHTTP(w, r)
	}
}

func (s *UserService) SetCookie(w http.ResponseWriter, sessionToken string) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  sessionToken,
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
}

func (s *UserService) GetUserID(r *http.Request) (int, error) {

	cookie, err := r.Cookie("session")
	if err != nil {
		return 0, err
	}

	session, err := s.SessionRepo.GetByToken(cookie.Value)
	if err != nil {
		return 0, err
	}

	return session.UserId, nil
}
