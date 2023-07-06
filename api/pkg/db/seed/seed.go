package seed

import (
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

func Seed(repos *models.Repositories) {

	SeedUsers(repos.UserRepo)
	SeedFollowers(repos)

	SeedPosts(repos)
	SeedGroups(repos)
}

// Seed database users from predefined dataset
func SeedUsers(repo *models.UserRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for _, seedUser := range SeedUserData {

		date, _ := time.Parse("2006-01-02", faker.Date())
		pwd, _ := services.HashPassword("123")

		tempUser := &models.User{
			FirstName: seedUser.FirstName,
			LastName:  seedUser.LastName,
			Email:     seedUser.Email,
			Password:  pwd,
			Nickname:  seedUser.Nickname,
			About:     seedUser.About,
			Birthday:  date,
		}

		id, err := repo.Insert(tempUser)
		seedUser.Id = id

		if err != nil {
			logger.Println(err)
		}

	}

}

func SeedPosts(repos *models.Repositories) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for _, seedUser := range SeedUserData {
		if seedUser.Id > 0 {
			for _, seedPost := range seedUser.PostSet {
				tempPost := &models.Post{
					Content:     seedPost.Content,
					UserId:      seedUser.Id,
					PrivacyType: seedPost.PrivacyType,
					CreatedAt:   seedPost.CreatedAt,
				}

				postId, err := repos.PostRepo.InsertSeedPost(tempPost)
				seedPost.Id = int(postId)
				tempPost.Id = postId

				//Insert post comments
				for _, comments := range seedPost.CommentSet {
					commentUser, err := repos.UserRepo.GetByEmail(comments.UserEmail)

					if err != nil {
						logger.Printf("%+v\n", err)
					}

					tempComment := &models.Comment{
						Content: comments.Content,
						UserId:  commentUser.Id,
						PostId:  postId,
					}

					id, err := repos.CommentRepo.Insert(tempComment)

					tempComment.Id = id

					// logger.Printf("%+v\n", tempComment)

					if err != nil {
						logger.Printf("%+v\n", err)
					}
				}

				for i := 0; i < seedPost.LoremComments; i++ {

					loremUser, err := repos.UserRepo.GetByEmail("l@l.com")

					if err != nil {
						logger.Printf("%+v\n", err)
					}
					tempComment := &models.Comment{
						Content: faker.Sentence(),
						UserId:  loremUser.Id,
						PostId:  postId,
					}

					id, err := repos.CommentRepo.Insert(tempComment)

					tempComment.Id = id

					// logger.Printf("%+v\n", tempComment)

					if err != nil {
						logger.Printf("%+v\n", err)
					}
				}

				// logger.Printf("%+v\n", tempPost)

				if err != nil {
					logger.Printf("%+v\n", err)
				}
			}
		}
	}

}

func SeedFollowers(repos *models.Repositories) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for _, seedUser := range SeedUserData {
		if seedUser.Id > 0 {
			for _, seedFollowing := range seedUser.FollowingEmails {

				followedUser, err := repos.UserRepo.GetByEmail(seedFollowing)
				if err != nil {
					logger.Printf("%+v\n", err)
				}

				tempFollowing := &models.Follower{
					FollowingId: followedUser.Id,
					FollowerId:  seedUser.Id,
					Accepted:    true,
				}

				_, err = repos.FollowerRepo.Insert(tempFollowing)

				logger.Printf("%+v\n", tempFollowing)

				if err != nil {
					logger.Printf("%+v\n", err)
				}
			}
		}
	}

}

func SeedGroups(repos *models.Repositories) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for _, group := range SeedGroupsData {

		user, err := repos.UserRepo.GetByEmail(group.CreatorEmail)

		if err != nil {
			logger.Printf("%+v\n", err)
		}

		tempGroup := &models.Group{
			Title:       group.Title,
			Description: group.Description,
			CreatorId:   user.Id,
			ImageBase64: toBase64Format(group.ImagePath),
		}

		id, err := repos.GroupRepo.Insert(tempGroup)

		if err != nil {
			logger.Printf("%+v\n", err)
		}

		//Add group users
		for _, groupUserEmail := range group.Users {
			groupUser, err := repos.UserRepo.GetByEmail(groupUserEmail)
			if err != nil {
				logger.Printf("%+v\n", err)
			}
			tempGroupUser := &models.GroupMemberModel{
				UserId:  groupUser.Id,
				GroupId: id,
			}

			_, err = repos.GroupUserRepo.Insert(tempGroupUser)
			if err != nil {
				logger.Printf("%+v\n", err)
			}
		}

		//Add group events
		for _, event := range group.SeedEventsData {
			eventCreator, err := repos.UserRepo.GetByEmail(event.CreatorEmail)
			if err != nil {
				logger.Printf("%+v\n", err)
			}

			tempEvent := &models.Event{
				GroupId:     id,
				UserId:      eventCreator.Id,
				CreatedAt:   event.CreatedAt,
				EventTime:   event.EventTime,
				TimeSpan:    event.TimeSpan,
				Title:       event.Title,
				Description: event.Description,
			}
			_, err = repos.EventRepo.InsertSeedEvent(tempEvent)
			if err != nil {
				logger.Printf("%+v\n", err)
			}

		}

	}

}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func toBase64Format(imageName string) string {

	if imageName != "" && imageName != "null" {

		filepath := fmt.Sprintf("./api/pkg/db/seed/images/%s", imageName)

		bytes, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}

		var base64Encoding string

		// Determine the content type of the image file
		mimeType := http.DetectContentType(bytes)

		// Prepend the appropriate URI scheme header depending
		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
		}

		// Append the base64 encoded output
		base64Encoding += toBase64(bytes)

		// Print the full base64 representation of the image
		// fmt.Println(base64Encoding)
		return base64Encoding
	}

	return ""

}
