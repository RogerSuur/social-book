package services

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
	"strconv"
	"time"
)

type IPostService interface {
	CreatePost(post *models.Post) error
	GetFeedPosts(userId int64, offset int64) ([]*feedPostJSON, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type PostService struct {
	Logger                *log.Logger
	PostRepository        models.IPostRepository
	AllowedPostRepository models.IAllowedPostRepository
}

func InitPostService(logger *log.Logger, postRepo *models.PostRepository, allowedPostRepo *models.AllowedPostRepository) *PostService {
	return &PostService{
		Logger:                logger,
		PostRepository:        postRepo,
		AllowedPostRepository: allowedPostRepo,
	}
}

type feedPostJSON struct {
	Id           int       `json:"id"`
	UserId       int       `json:"userId"`
	UserName     string    `json:"userName"`
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

	postId, err := s.PostRepository.Insert(post)

	if post.PrivacyType == enums.SubPrivate {
		for _, receiver := range post.Receivers {

			receiverId, err := strconv.Atoi(receiver)

			if err != nil {
				s.Logger.Printf("CreatePost atoi parse error: %s", err)
			}

			allowedPost := models.AllowedPost{
				UserId: receiverId,
				PostId: int(postId),
			}

			s.AllowedPostRepository.Insert(&allowedPost)
		}
	}

	if err != nil {
		log.Printf("CreatePost error: %s", err)
	}

	return err
}

func (s *PostService) GetFeedPosts(userId int64, offset int64) ([]*feedPostJSON, error) {
	// fmt.Println("userId", userId)

	if offset == 0 {
		lastPostId, err := s.PostRepository.GetLastPostId()
		if err != nil {
			s.Logger.Printf("GetFeedPosts error: %s", err)
			return nil, err
		}
		offset = lastPostId + 1
	}

	posts, err := s.PostRepository.GetAllFeedPosts(userId, offset)

	if err != nil {
		s.Logger.Printf("GetFeedPosts error: %s", err)
	}

	feedPosts := []*feedPostJSON{}

	for _, p := range posts {
		// commentCount, err := s.PostRepository.GetCommentCount(p.Id)

		feedPosts = append(feedPosts, &feedPostJSON{
			p.Id,
			p.UserId,
			p.UserName,
			p.Content,
			p.ImagePath,
			p.CommentCount,
			p.CreatedAt,
		})
	}
	// fmt.Println("feedPosts:", feedPosts)

	s.Logger.Printf("Retrived feed posts: %d", len(feedPosts))

	return feedPosts, nil
}
