package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

type IPostService interface {
	CreatePost(post *models.Post) error
	GetFeedPosts(userId int, offset int) ([]*feedPostJSON, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type PostService struct {
	Logger         *log.Logger
	PostRepository models.IPostRepository
}

func InitPostService(postRepo *models.PostRepository) *PostService {
	return &PostService{
		PostRepository: postRepo,
	}
}

type feedPostJSON struct {
	Id           int       `json:"id"`
	UserId       int       `json:"userId"`
	Content      string    `json:"content"`
	ImagePath    string    `json:"imagePath"`
	CommentCount int       `json:"commentCount"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (s *PostService) CreatePost(post *models.Post) error {

	if len(post.Content) == 0 {
		err := errors.New("content too short")
		log.Printf("CreatePost error: %s", err)
		return err
	}

	_, err := s.PostRepository.Insert(post)

	if err != nil {
		log.Printf("CreatePost error: %s", err)
	}

	return err
}

func (s *PostService) GetFeedPosts(userId int, offset int) ([]*feedPostJSON, error) {

	posts, err := s.PostRepository.GetAllFeedPosts(userId, offset)

	if err != nil {
		fmt.Println(err)
	}

	feedPosts := []*feedPostJSON{}

	for _, p := range posts {
		commentCount, err := s.PostRepository.GetCommentCount(p.Id)

		if err != nil {
			s.Logger.Printf("GetFeedPosts error: %s", err)
		}

		feedPosts = append(feedPosts, &feedPostJSON{
			p.Id,
			p.UserId,
			p.Content,
			p.ImagePath,
			commentCount,
			p.CreatedAt,
		})
	}

	return feedPosts, nil
}
