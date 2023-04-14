package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		env := models.CreateEnv(s.DB)

		// check if cookie exists
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Printf("No cookie found: %s", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// check if session exists
		_, err = env.Sessions.GetByToken(cookie.Value)

		if err != nil {
			log.Printf("No session found: %s", err)
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		// finally, call the handler
		handler.ServeHTTP(w, r)
	}
}

func (s *Service) SetCookie(w http.ResponseWriter, sessionToken string) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  sessionToken,
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
