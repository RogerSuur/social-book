package services

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProfileJSON struct {
	UserID      int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Birthday    string    `json:"birthday"`
	Nickname    string    `json:"nickname"`
	About       string    `json:"about"`
	AvatarImage string    `json:"avatarImage"`
	CreatedAt   time.Time `json:"createdAt"`
	IsPublic    bool      `json:"isPublic"`
	IsFollowed  bool      `json:"isFollowed"`
}

type ProfileUpdateJSON struct {
	UserID      int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Birthday    string    `json:"birthday"`
	Nickname    string    `json:"nickname"`
	About       string    `json:"about"`
	AvatarImage string    `json:"avatarImage"`
	CreatedAt   time.Time `json:"createdAt"`
	IsPublic    bool      `json:"isPublic"`
}

type FollowerData struct {
	UserID      int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Nickname    string `json:"nickname"`
	AvatarImage string `json:"avatarImage"`
	Accepted    bool   `json:"accepted"`
}

type IUserService interface {
	Authenticate(handler http.HandlerFunc) http.HandlerFunc
	CreateUser(user *models.User) (int64, error)
	UpdateUserData(userID int64, updateData ProfileUpdateJSON) error
	GetUserData(userID int64) (*ProfileJSON, error)
	GetUserID(r *http.Request) (int64, error)
	SetCookie(w http.ResponseWriter, sessionToken string)
	ClearCookie(w http.ResponseWriter)
	UserLogin(user *models.User) (string, error)
	UserLogout(r *http.Request) error
	UserRegister(user *models.User) (string, error)
	GetUserFollowers(userID int64) ([]FollowerData, error)
	GetUserFollowing(userID int64) ([]FollowerData, error)
	IsFollowed(followerID int64, followingID int64) bool
	Unfollow(followerID int64, followingID int64) error
	UpdateUserImage(userID int64, file multipart.File, fileHeader *multipart.FileHeader) error
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserService struct {
	Logger           *log.Logger
	UserRepo         models.IUserRepository
	SessionRepo      models.ISessionRepository
	FollowerRepo     models.IFollowerRepository
	NotificationRepo models.INotificationRepository
}

// InitUserService initializes the user controller.
func InitUserService(
	userRepo *models.UserRepository,
	sessionRepo *models.SessionRepository,
	followerRepo *models.FollowerRepository,
	notificationRepo *models.NotificationRepository,
) *UserService {
	return &UserService{
		Logger:           log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		UserRepo:         userRepo,
		SessionRepo:      sessionRepo,
		FollowerRepo:     followerRepo,
		NotificationRepo: notificationRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB

	return s.UserRepo.Insert(user)
}

func (s *UserService) UpdateUserData(userID int64, updateData ProfileUpdateJSON) error {

	user, err := s.UserRepo.GetById(userID)
	if err != nil {
		return err
	}

	// update the values of the user in case matching key is found in updateData

	switch {
	case updateData.Nickname != user.Nickname:
		if len(updateData.Nickname) < 3 {
			s.Logger.Printf("User nickname too short")
			return errors.New("nickname")
		}

		err = s.UserRepo.CheckIfNicknameExists(updateData.Nickname, userID)
		if err == nil {
			s.Logger.Printf("User nickname already exists")
			return errors.New("nickname")
		}
		user.Nickname = updateData.Nickname

	case updateData.About != user.About:
		user.About = updateData.About

	case updateData.IsPublic != user.IsPublic:
		user.IsPublic = updateData.IsPublic

	default:
		return errors.New("no data to update")

	}

	return s.UserRepo.Update(user)
}

func (s *UserService) GetUserData(userID int64) (*ProfileJSON, error) {
	user, err := s.UserRepo.GetById(userID)
	if err != nil {
		return nil, err
	}
	userJSON := &ProfileJSON{
		UserID:      int(userID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Birthday:    user.Birthday.Format("02/01/2006"),
		Nickname:    user.Nickname,
		About:       user.About,
		AvatarImage: user.ImagePath,
		CreatedAt:   user.CreatedAt,
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

	if len(user.Nickname) > 3 {
		err = s.UserRepo.CheckIfNicknameExists(user.Nickname, 0)
		if err == nil {
			log.Printf("User nickname already exists")
			return "", errors.New("nickname")
		}
	} else if len(user.Nickname) > 0 {
		log.Printf("Nickname is too short")
		return "", errors.New("nickname")
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

func (s *UserService) UserLogout(r *http.Request) error {
	cookie, err := r.Cookie("session")
	if err != nil {
		s.Logger.Printf("No cookie found: %s", err)
		return errors.New("no cookie found")
	}

	err = s.SessionRepo.DeleteByToken(cookie.Value)
	if err != nil {
		s.Logger.Printf("Cannot delete session: %s", err)
		return errors.New("cannot delete session")
	}
	return nil
}

func (s *UserService) Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session")
		if err != nil {
			s.Logger.Printf("No cookie found: %s", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = s.SessionRepo.GetByToken(cookie.Value)

		if err != nil {
			s.Logger.Printf("No session found: %s", err)
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		// required for auth handler
		if handler == nil {
			return
		}

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

func (s *UserService) ClearCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func (s *UserService) GetUserID(r *http.Request) (int64, error) {

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
		user, err := s.UserRepo.GetById(follower.FollowerId)
		if err != nil {
			return nil, err
		}
		follower := FollowerData{
			UserID:      int(user.Id),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Nickname:    user.Nickname,
			AvatarImage: user.ImagePath,
			Accepted:    follower.Accepted,
		}
		followersData = append(followersData, follower)
	}

	fmt.Println("followersData", followersData)

	return followersData, nil
}

func (s *UserService) GetUserFollowing(userID int64) ([]FollowerData, error) {

	following, err := s.FollowerRepo.GetFollowingById(userID)
	if err != nil {
		return nil, err
	}

	followingData := []FollowerData{}

	for _, follower := range following {
		user, err := s.UserRepo.GetById(follower.FollowingId)
		if err != nil {
			return nil, err
		}
		following := FollowerData{
			UserID:      int(user.Id),
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

func (s *UserService) IsFollowed(userID int64, followerID int64) bool {

	_, err := s.FollowerRepo.GetByFollowerAndFollowing(userID, followerID)
	return err == nil
}

func (s *UserService) Unfollow(userID int64, followerID int64) error {

	follower, err := s.FollowerRepo.GetByFollowerAndFollowing(userID, followerID)
	if err != nil {
		return err
	}

	err = s.FollowerRepo.Delete(follower)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserImage(userID int64, imageFile multipart.File, header *multipart.FileHeader) error {

	// check if user exists
	_, err := s.UserRepo.GetById(userID)
	if err != nil {
		s.Logger.Printf("User not found: %s", err)
		return err
	}

	// check if file is an image
	if !strings.HasPrefix(header.Header.Get("Content-Type"), "image") {
		s.Logger.Println("Not an image")
		return errors.New("not an image")
	}

	// save image
	imagePath, err := utils.SaveImage(userID, imageFile, header)
	if err != nil {
		s.Logger.Printf("Cannot save image: %s", err)
		return err
	}

	// update user image path
	err = s.UserRepo.UpdateImage(userID, imagePath)
	if err != nil {
		return err
	}

	return nil
}
