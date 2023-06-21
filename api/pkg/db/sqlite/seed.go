package database

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

	SeedComments(repos.CommentRepo)

	//Single value test
	/* test, err := repos.SessionRepo.GetByToken("b48976a7-64e0-4ff5-a816-6362dbcb1aa0")

	if err != nil {
		repos.UserRepo.Logger.Printf("%+v\n", err)
	} */

	// //Array TEST
	// tests, err := env.Groups.GetAllByCreatorId(1)

	// if err != nil {
	// 	repos.UserRepo.Logger.Printf("%+v\n", err)

	// } else {
	// repos.UserRepo.Logger.Printf("%+v\n", test)

	// 	for _, v := range tests {
	// 		repos.UserRepo.Logger.Printf("%+v\n", v)

	// 	}

	// }

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
				}

				postId, err := repos.PostRepo.Insert(tempPost)
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
						PostId:  int(postId),
					}

					id, err := repos.CommentRepo.Insert(tempComment)

					tempComment.Id = int(id)

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
						PostId:  int(postId),
					}

					id, err := repos.CommentRepo.Insert(tempComment)

					tempComment.Id = int(id)

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
					FollowingId: int64(followedUser.Id),
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

	// for i := 0; i < 10; i++ {
	// 	tempFollower := &models.Follower{
	// 		FollowingId: i + 1,
	// 		FollowerId:  i + 11,
	// 		Accepted:    true,
	// 	}

	// 	_, err := repo.Insert(tempFollower)

	// 	// logger.Printf("%+v\n", tempFollower)

	// 	if err != nil {
	// 		logger.Println(err)
	// 	}

	// }
}

func SeedComments(repo *models.CommentRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		tempComment := &models.Comment{
			Content: faker.Sentence(),
			UserId:  int64(i + 1),
			PostId:  i + 1,
		}

		id, err := repo.Insert(tempComment)

		tempComment.Id = int(id)

		// logger.Printf("%+v\n", tempComment)

		if err != nil {
			logger.Println(err)
		}

	}
}

func SeedGroups(repo *models.GroupRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		tempGroup := &models.Group{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			CreatorId:   i + 1,
		}

		_, err := repo.Insert(tempGroup)

		// logger.Printf("%+v\n", tempGroup)

		if err != nil {
			logger.Println(err)
		}

	}
}

// func SeedSessions(repo *models.SessionRepository) {
// 	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

// 	for i := 0; i < 10; i++ {
// 		session := &models.Session{
// 			UserId: i + 1,
// 			Token:  faker.UUIDHyphenated(),
// 		}

// 		_, err := repo.Insert(session)

// 		//logger.Printf("%+v\n", session)

// 		if err != nil {
// 			logger.Println(err)
// 		}

// 	}
// }
