package models

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"database/sql"
	"fmt"
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
	GetAllFeedPosts(offset int) ([]*Post, error)
	GetById(id int64) (*Post, error)
	Insert(post *Post) (int64, error)
	GetCommentCount(postId int) (int, error)
}

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

const FeedLimit = 100

func (m PostRepository) Insert(post *Post) (int64, error) {
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

func (m PostRepository) GetById(id int64) (*Post, error) {
	query := `SELECT id, user_id,  title, content, created_at, image_path, privacy_type_id FROM posts WHERE id = ?`
	row := m.DB.QueryRow(query, id)
	post := &Post{}

	err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.ImagePath, &post.PrivacyType)

	return post, err
}

func (m PostRepository) GetAllByUserId(id int64) ([]*Post, error) {

	stmt := `SELECT id, user_id,  title, content, created_at, image_path, privacy_type_id FROM posts p
	WHERE user_id = ?
    ORDER BY created_at DESC`

	rows, err := m.DB.Query(stmt, id)
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

// Return all posts to the current user by offset
func (m PostRepository) GetAllFeedPosts(offset int) ([]*Post, error) {

	currentUserId := 12

	stmt := `SELECT p.id, p.user_id, p.content, p.created_at, p.image_path, privacy_type_id FROM posts p 
	LEFT JOIN  followers f ON  
	p.user_id = f.following_id
	LEFT JOIN allowed_private_posts app ON
	p.id = app.post_id
	WHERE privacy_type_id = 1 
	OR privacy_type_id = 2 AND f.id IS NOT NULL AND f.follower_id = ? AND f.accepted = 1 AND f.active = 1
	OR privacy_type_id = 3 AND f.id IS NOT NULL AND f.follower_id = ? AND f.accepted = 1 AND f.active = 1 AND app.id IS NOT NULL AND app.user_id = ?
	OR p.user_id = ?
	GROUP BY p.id
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?`

	args := []interface{}{
		currentUserId,
		currentUserId,
		currentUserId,
		currentUserId,
		FeedLimit,
		offset,
	}

	rows, err := m.DB.Query(stmt, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		post := &Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.ImagePath, &post.PrivacyType)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	fmt.Println(posts)

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m PostRepository) GetCommentCount(postId int) (int, error) {
	query := `SELECT COUNT(*) FROM comments WHERE post_id = ?`
	row := m.DB.QueryRow(query, postId)
	var commentCount int

	err := row.Scan(&commentCount)

	if err != nil {
		return -1, err
	}

	return commentCount, nil
}
