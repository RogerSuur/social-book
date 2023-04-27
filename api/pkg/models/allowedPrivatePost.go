package models

import (
	"database/sql"
)

type AllowedPost struct {
	UserId int
	PostId int
}

type AllowedPostModel struct {
	DB *sql.DB
}

func (m AllowedPostModel) Insert(allowedPost *AllowedPost) (int64, error) {
	query := `INSERT INTO allowed_private_posts (post_id, user_id)
	VALUES(?, ?)`

	args := []interface{}{
		allowedPost.PostId,
		allowedPost.UserId,
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
