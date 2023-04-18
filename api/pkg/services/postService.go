package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"errors"
	"log"
)

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
