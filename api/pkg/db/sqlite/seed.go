package database

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"log"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

func Seed(repos models.Repositories) {

	//SeedUsers(repos.UserRepo)
	//SeedSessions(repos.SessionRepo)
	//SeedPosts(repos.PostRepo)

	//SeedComments(repos.CommentRepo)
	// SeedGroups(db)
	// SeedFollowers(db)

	// env := models.CreateEnv(db)

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

func SeedSessions(repo *models.SessionRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		session := &models.Session{
			UserId: i + 1,
			Token:  faker.UUIDHyphenated(),
		}

		_, err := repo.Insert(session)

		//logger.Printf("%+v\n", session)

		if err != nil {
			logger.Println(err)
		}

	}
}

func SeedFollowers(repo *models.FollowerRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		tempFollower := &models.Follower{
			FollowingId: i + 1,
			FollowerId:  i + 11,
			Accepted:    true,
			Active:      true,
		}

		_, err := repo.Insert(tempFollower)

		// logger.Printf("%+v\n", tempFollower)

		if err != nil {
			logger.Println(err)
		}

	}
}

func SeedPosts(repo *models.PostRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		tempPost := &models.Post{
			Content:     faker.Sentence(),
			UserId:      i + 1,
			PrivacyType: enums.SubPrivate,
		}

		_, err := repo.Insert(tempPost)

		// logger.Printf("%+v\n", tempPost)

		if err != nil {
			logger.Println(err)
		}

	}
}

func SeedUsers(repo *models.UserRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {

		date, _ := time.Parse("2006-01-02", faker.Date())
		pwd, _ := services.HashPassword("something")
		tempUser := &models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Password:  pwd,
			Nickname:  faker.Username(),
			About:     faker.Sentence(),
			ImagePath: faker.Word(),
			Birthday:  date,
		}

		id, err := repo.Insert(tempUser)
		tempUser.Id = int(id)

		if err != nil {
			logger.Println(err)
		}

		// logger.Printf("%+v\n", tempUser)
		// logger.Println()

	}
}

func SeedComments(repo *models.CommentRepository) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	for i := 0; i < 10; i++ {
		tempComment := &models.Comment{
			Content: faker.Sentence(),
			UserId:  i + 1,
			PostId:  10 - i,
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
