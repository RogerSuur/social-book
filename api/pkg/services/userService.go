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

type FollowerData struct {
	UserID      int
	FirstName   string
	LastName    string
	Nickname    string
	AvatarImage string
	Accepted    bool
}

type IUserService interface {
	Authenticate(handler http.HandlerFunc) http.HandlerFunc
	CreateUser(user *models.User) (int64, error)
	GetUserData(userID int64) (*ProfileJSON, error)
	GetUserID(r *http.Request) (int, error)
	SetCookie(w http.ResponseWriter, sessionToken string)
	UserLogin(user *models.User) (string, error)
	UserRegister(user *models.User) (string, error)
	GetUserFollowers(userID int64) ([]FollowerData, error)
	GetUserFollowing(userID int64) ([]FollowerData, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserService struct {
	Logger       *log.Logger
	UserRepo     models.IUserRepository
	SessionRepo  models.ISessionRepository
	FollowerRepo models.IFollowerRepository
}

// InitUserService initializes the user controller.
func InitUserService(
	userRepo *models.UserRepository,
	sessionRepo *models.SessionRepository,
	followerRepo *models.FollowerRepository,
) *UserService {
	return &UserService{
		Logger:       log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo:     userRepo,
		SessionRepo:  sessionRepo,
		FollowerRepo: followerRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB

	return s.UserRepo.Insert(user)
}

func (s *UserService) GetUserData(userID int64) (*ProfileJSON, error) {
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
		return "", errors.New("email")
	}

	if len(user.Nickname) > 0 {
		err = s.UserRepo.CheckIfNicknameExists(user.Nickname, 0)
		if err == nil {
			log.Printf("User nickname already exists")
			return "", errors.New("nickname")
		}
	}

	if !CheckPasswordStrength(user.Password) {
		log.Printf("Password is not strong enough")
		return "", errors.New("password")
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		s.Logger.Printf("Cannot hash password: %s", err)
		return "", errors.New("server")
	}
	user.Password = hashedPassword

	// create user
	lastID, err := s.UserRepo.Insert(user)
	if err != nil {
		s.Logger.Printf("Cannot create user: %s", err)
		return "", errors.New("server")
	}

	s.Logger.Printf("User successfully registered (Last inserted ID: %v)", lastID)

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
	s.Logger.Printf("Session initiated, (last inserted ID %v:)", lastID)

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

func (s *UserService) GetUserFollowers(userID int64) ([]FollowerData, error) {

	followers, err := s.FollowerRepo.GetFollowersById(userID)
	if err != nil {
		return nil, err
	}

	followersData := []FollowerData{}

	for _, follower := range followers {
		user, err := s.UserRepo.GetById(int64(follower.FollowerId))
		if err != nil {
			return nil, err
		}
		follower := FollowerData{
			UserID:      user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Nickname:    user.Nickname,
			AvatarImage: user.ImagePath,
			Accepted:    follower.Accepted,
		}
		followersData = append(followersData, follower)
	}

	return followersData, nil
}

func (s *UserService) GetUserFollowing(userID int64) ([]FollowerData, error) {

	following, err := s.FollowerRepo.GetFollowingById(userID)
	if err != nil {
		return nil, err
	}

	followingData := []FollowerData{}

	for _, follower := range following {
		user, err := s.UserRepo.GetById(int64(follower.FollowingId))
		if err != nil {
			return nil, err
		}
		following := FollowerData{
			UserID:      user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Nickname:    user.Nickname,
			AvatarImage: user.ImagePath,
			Accepted:    follower.Accepted,
		}
		followingData = append(followingData, following)
	}

	return followingData, nil
}
