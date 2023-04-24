package models

import (
	"database/sql"
	"log"
	"os"
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

type ICommentRepository interface {
	GetAllByPostId(postId int) ([]*Comment, error)
	GetAllByUserId(userId int) ([]*Comment, error)
	GetById(id int64) (*Comment, error)
	Insert(comment *Comment) (int64, error)
	Update(comment *Comment) error
}

type CommentRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo CommentRepository) Insert(comment *Comment) (int64, error) {
	query := `INSERT INTO comments (post_id, user_id, content, image_path, created_at)
	VALUES(?, ?, ?, ?, ?)`

	args := []interface{}{
		comment.PostId,
		comment.UserId,
		comment.Content,
		comment.ImagePath,
		time.Now(),
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Inserted comment by user %d for post %d (last insert ID: %d)", comment.UserId, comment.PostId, lastId)

	return lastId, nil
}

func (repo CommentRepository) Update(comment *Comment) error {

	//TODO
	//update comment in database
	return nil
}

func (repo CommentRepository) GetById(id int64) (*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE id = ?`
	row := repo.DB.QueryRow(query, id)
	comment := &Comment{}

	err := row.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.ImagePath, &comment.CreatedAt)

	return comment, err
}

func (repo CommentRepository) GetAllByPostId(postId int) ([]*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC`
	rows, err := repo.DB.Query(query, postId)
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

func (repo CommentRepository) GetAllByUserId(userId int) ([]*Comment, error) {
	query := `SELECT id, post_id, user_id, content,  image_path, created_at FROM comments WHERE user_id = ? ORDER BY created_at DESC`
	rows, err := repo.DB.Query(query, userId)
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
