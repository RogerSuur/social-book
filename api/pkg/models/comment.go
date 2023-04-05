package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id        int
	PostId    int
	UserId    int
	Content   string
	ImagePath string
	CreatedAt time.Time
}

type CommentModel struct {
	DB *sql.DB
}

func (m CommentModel) Insert(comment *Comment) (int64, error) {
	query := `INSERT INTO comments (post_id, user_id, content, image_path, created_at)
	VALUES(?, ?, ?, ?, ?)`

	args := []interface{}{
		comment.PostId,
		comment.UserId,
		comment.Content,
		comment.ImagePath,
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

func (m CommentModel) Update(comment *Comment) error {

	//TODO
	//update comment in database
	return nil
}

func (m CommentModel) GetById(id int64) (*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE id = ?`
	row := m.DB.QueryRow(query, id)
	comment := &Comment{}

	err := row.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.ImagePath, &comment.CreatedAt)

	return comment, err
}

func (m CommentModel) GetAllByPostId(postId int) ([]*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC`
	rows, err := m.DB.Query(query, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	comments := []*Comment{}

	for rows.Next() {
		comment := &Comment{}

		err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.ImagePath, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (m CommentModel) GetAllByUserId(userId int) ([]*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE user_id = ? ORDER BY created_at DESC`
	rows, err := m.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	comments := []*Comment{}

	for rows.Next() {
		comment := &Comment{}

		err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.ImagePath, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
