package database

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
)

// type Seed struct {
// 	db *sql.DB
// }

func Seed(repos models.Repositories) {

	SeedUsers(repos.UserRepo)
	SeedSessions(repos.SessionRepo)
	SeedPosts(repos.PostRepo)

	// SeedComments(db)
	// SeedGroups(db)
	// SeedFollowers(db)

	// env := models.CreateEnv(db)

	//Single value test
	test, err := repos.SessionRepo.GetByToken("aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL.aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL.aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL")

	// if err != nil {
	// 	 fmt.Printf("%+v\n", err)

	// }

	// //Array TEST
	// tests, err := env.Groups.GetAllByCreatorId(1)

	// if err != nil {
	// 	fmt.Printf("%+v\n", err)

	// } else {
	fmt.Printf("%+v\n", test)

	// 	for _, v := range tests {
	// 		fmt.Printf("%+v\n", v)

	// 	}

	// }

	// }
}

func SeedSessions(repo *models.SessionRepository) {

	for i := 0; i < 10; i++ {
		session := &models.Session{
			UserId: i + 1,
			Token:  faker.Jwt(),
		}

		_, err := repo.Insert(session)

		fmt.Printf("%+v\n", session)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedFollowers(repo *models.FollowerRepository) {

	// for i := 0; i < 10; i++ {
	tempFollower := &models.Follower{
		FollowingId: 11,
		FollowerId:  12,
		Accepted:    true,
		Active:      true,
	}

		_, err := repo.Insert(tempFollower)

	// fmt.Printf("%+v\n", tempFollower)

	if err != nil {
		fmt.Println(err)
	}

	// }
}

func SeedPosts(repo *models.PostRepository) {

	for i := 0; i < 10; i++ {

		res := (i % 3) + 1

		tempPost := &models.Post{
			Title:       faker.Word(),
			Content:     faker.Sentence(),
			UserId:      12,
			PrivacyType: enums.PrivacyType(res),
		}

		_, err := repo.Insert(tempPost)

		// fmt.Printf("%+v\n", tempPost)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedUsers(repo *models.UserRepository) {

	for i := 0; i < 10; i++ {

		date, _ := time.Parse("2006-01-02", faker.Date())
		tempUser := &models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Password:  "something",
			Nickname:  faker.Username(),
			About:     faker.Sentence(),
			ImagePath: faker.Word(),
			Birthday:  date,
		}

		id, err := repo.Insert(tempUser)
		tempUser.Id = int(id)

		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("%+v\n", tempUser)
		// fmt.Println()

	}
}

func SeedComments(repo *models.CommentRepository) {

	for i := 0; i < 10; i++ {
		tempComment := &models.Comment{
			Content: faker.Sentence(),
			UserId:  i + 1,
			PostId:  10 - i,
		}

		id, err := repo.Insert(tempComment)

		tempComment.Id = int(id)

		// fmt.Printf("%+v\n", tempComment)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedGroups(repo *models.GroupRepository) {

	for i := 0; i < 10; i++ {
		tempGroup := &models.Group{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			CreatorId:   i + 1,
		}

		_, err := repo.Insert(tempGroup)

		// fmt.Printf("%+v\n", tempGroup)

		if err != nil {
			fmt.Println(err)
		}

	}
}
