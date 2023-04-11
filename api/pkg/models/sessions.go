package models

import "database/sql"

type Session struct {
	Id     int
	UserId int
	Token  string
}

type SessionModel struct {
	DB *sql.DB
}

func (m SessionModel) Insert(session *Session) (int64, error) {
	query := `INSERT INTO sessions (user_id, session) VALUES(?, ?)`

	args := []interface{}{
		session.UserId,
		session.Token,
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

func (m SessionModel) GetByToken(token string) (*Session, error) {
	query := `SELECT id, user_id, session FROM sessions WHERE session = ?`
	row := m.DB.QueryRow(query, token)
	session := &Session{}

	err := row.Scan(&session.Id, &session.UserId, &session.Token)

	return session, err
}

func (m SessionModel) DeleteByToken(token string) error {
	query := `DELETE FROM sessions WHERE session = ?`

	_, err := m.DB.Exec(query, token)

	return err
}
