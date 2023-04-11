package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	user *models.UserModel
}

func (u *UserService) CreateUser(user *models.User) (int64, error) {
	// do validation/business rule validation here
	// .. more user stuff
	// finally, insert into the DB
	return u.user.Insert(user)
}

func Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check if cookie exists
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Printf("No cookie found: %s", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// check if session exists
		sessionModel := models.SessionModel{}
		_, err = sessionModel.GetByToken(cookie.Value)

		if err != nil {
			log.Printf("No session found: %s", err)
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		// finally, call the handler
		handler.ServeHTTP(w, r)
	}
}

func UserLogin(user *models.User) (string, error) {

	// check if user exists
	userModel := models.UserModel{}
	dbUser, err := userModel.GetByEmail(user.Email)
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
	sessionModel := models.SessionModel{}
	_, err = sessionModel.Insert(&session)
	if err != nil {
		log.Printf("Cannot create session: %s", err)
		return "", errors.New("cannot create session")
	}
	return sessionToken, nil
}

func SetCookie(w http.ResponseWriter, sessionToken string) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  sessionToken,
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
