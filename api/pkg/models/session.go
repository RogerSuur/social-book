package models

import (
	"database/sql"
	"time"
)

type Session struct {
	Id        int
	UserId    int
	Token     string
	CreatedAt time.Time
}

type ISessionRepository interface {
	DeleteByToken(token string) error
	GetByToken(token string) (*Session, error)
	GetUserSessions(id int) ([]*Session, error)
	Insert(session *Session) (int64, error)
}

type SessionRepository struct {
	DB *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		DB: db,
	}
}

// Inserts a new user session to database
func (m SessionRepository) Insert(session *Session) (int64, error) {
	query := `INSERT INTO user_sessions (user_id, token, created_at)
	VALUES(?, ?, ?)`

	args := []interface{}{
		session.UserId,
		session.Token,
		time.Now(),
	}

	result, err := m.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

// Returns a session by token
func (m SessionRepository) GetByToken(token string) (*Session, error) {

	query := `SELECT id, user_id, token, created_at FROM user_sessions WHERE token = ?`
	row := m.DB.QueryRow(query, token)
	session := &Session{}

	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.CreatedAt)

	return session, err
}

// Returns all user sessions
func (m SessionRepository) GetUserSessions(id int) ([]*Session, error) {

	//TODO
	//Not sure if needed
	return nil, nil
}

// remove a session by token
func (m SessionRepository) DeleteByToken(token string) error {
	query := `DELETE FROM user_sessions WHERE token = ?`

	_, err := m.DB.Exec(query, token)

	return err
}
