package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

type feedPostJSON struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	Content   string    `json:"content"`
	ImagePath string    `json:"imagePath"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Service) CreatePost(post *models.Post) error {

	env := models.CreateEnv(s.DB)

	if len(post.Title) == 0 {
		err := errors.New("title too short")
		log.Printf("CreatePost error: %s", err)
		return err
	}

	if len(post.Content) == 0 {
		err := errors.New("content too short")
		log.Printf("CreatePost error: %s", err)
		return err
	}

	_, err := env.Posts.Insert(post)

	if err != nil {
		log.Printf("CreatePost error: %s", err)
	}

	return err
}

func (s *Service) GetFeedPosts(offset string) ([]*feedPostJSON, error) {

	env := models.CreateEnv(s.DB)

	posts, _ := env.Posts.GetAllFeedPosts(1)

	fmt.Println(posts)

	feedPosts := []*feedPostJSON{}

	for _, p := range posts {
		feedPosts = append(feedPosts, &feedPostJSON{
			p.Id,
			p.UserId,
			p.Content,
			p.ImagePath,
			p.CreatedAt,
		})
	}

	return feedPosts, nil
}
