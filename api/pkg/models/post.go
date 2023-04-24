package models

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"database/sql"
	"log"
	"os"
	"time"
)

type Post struct {
	Id          int
	UserId      int
	Title       string
	Content     string
	ImagePath   string
	CreatedAt   time.Time
	PrivacyType enums.PrivacyType
}

type IPostRepository interface {
	GetAllByUserId(id int64) ([]*Post, error)
	GetAllFeedPosts(id int64) ([]*Post, error)
	GetById(id int64) (*Post, error)
	Insert(post *Post) (int64, error)
}

type PostRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepository {
	return &PostRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo PostRepository) Insert(post *Post) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, created_at, image_path, privacy_type_id)
	VALUES(?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		post.UserId,
		post.Title,
		post.Content,
		time.Now(),
		post.ImagePath,
		post.PrivacyType,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Inserted post '%s' by user %d (last insert ID: %d)", post.Title, post.UserId, lastId)

	return lastId, nil
}

func (repo PostRepository) GetById(id int64) (*Post, error) {
	query := `SELECT id, user_id,  title, content, created_at, image_path, privacy_type_id FROM posts WHERE id = ?`
	row := repo.DB.QueryRow(query, id)
	post := &Post{}

	err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.ImagePath, &post.PrivacyType)

	return post, err
}

func (repo PostRepository) GetAllByUserId(id int64) ([]*Post, error) {

	stmt := `SELECT id, user_id,  title, content, created_at, image_path, privacy_type_id FROM posts p
	WHERE user_id = ?
    ORDER BY created_at DESC`

	rows, err := repo.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		post := &Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.ImagePath, &post.PrivacyType)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo PostRepository) GetAllFeedPosts(id int64) ([]*Post, error) {

	//TODO
	//return all posts to the current user
	return nil, nil
}
