package seed

import (
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"log"
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
			ImagePath: faker.Word(),
			Birthday:  date,
		}

		id, err := repo.Insert(tempUser)
		seedUser.Id = int(id)

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
					UserId:      int64(seedUser.Id),
					PrivacyType: seedPost.PrivacyType,
					CreatedAt:   seedPost.CreatedAt,
				}

				postId, err := repos.PostRepo.InsertSeedPost(tempPost)
				seedPost.Id = int(postId)
				tempPost.Id = int(postId)

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
					FollowerId:  int64(seedUser.Id),
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

	}

}
