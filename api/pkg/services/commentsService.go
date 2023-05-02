package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"log"
	"os"
	"time"
)

type ICommentService interface {
	GetPostComments(userId int, offset int) ([]*commentJSON, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type CommentService struct {
	Logger            *log.Logger
	CommentRepository models.ICommentRepository
}

func InitCommentService(commentRepo *models.CommentRepository) *CommentService {
	return &CommentService{
		Logger:            log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		CommentRepository: commentRepo,
	}
}

type commentJSON struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	Content   string    `json:"content"`
	ImagePath string    `json:"imagePath"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *CommentService) GetPostComments(postId int, offset int) ([]*commentJSON, error) {

	result, err := s.CommentRepository.GetAllByPostId(postId, offset)

	if err != nil {
		s.Logger.Printf("Failed fetching comments: %s", err)
	}

	comments := []*commentJSON{}

	for _, p := range result {
		comments = append(comments, &commentJSON{
			p.Id,
			p.UserId,
			p.Content,
			p.ImagePath,
			p.CreatedAt,
		})
	}

	return comments, nil
}
