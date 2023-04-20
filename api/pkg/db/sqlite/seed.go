package database

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
)

// type Seed struct {
// 	db *sql.DB
// }

func Seed(db *sql.DB) {

	SeedUsers(db)
	SeedPosts(db)
	// SeedComments(db)
	// SeedGroups(db)
	// SeedFollowers(db)
	// SeedSessions(db)

	env := models.CreateEnv(db)

	//Single value test
	test, err := env.Sessions.GetByToken("aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL.aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL.aWiLyVUgPWdDqRVBVHqRKQghrxFigZFpNvWjpCWL")

	if err != nil {
		fmt.Printf("%+v\n", err)

	}

	//Array TEST
	tests, err := env.Groups.GetAllByCreatorId(1)

	if err != nil {
		fmt.Printf("%+v\n", err)

	} else {
		fmt.Printf("%+v\n", test)

		for _, v := range tests {
			fmt.Printf("%+v\n", v)

		}

	}

}

func SeedSessions(db *sql.DB) {
	env := models.CreateEnv(db)

	for i := 0; i < 10; i++ {
		session := &models.Session{
			UserId: i + 1,
			Token:  faker.Jwt(),
		}

		_, err := env.Sessions.Insert(session)

		fmt.Printf("%+v\n", session)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedFollowers(db *sql.DB) {
	env := models.CreateEnv(db)

	for i := 0; i < 10; i++ {
		tempFollower := &models.Follower{
			FollowingId: i + 1,
			FollowerId:  i + 11,
			Accepted:    true,
			Active:      true,
		}

		_, err := env.Followers.Insert(tempFollower)

		// fmt.Printf("%+v\n", tempFollower)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedPosts(db *sql.DB) {
	env := models.CreateEnv(db)

	for i := 0; i < 10; i++ {
		tempPost := &models.Post{
			Title:       faker.Word(),
			Content:     faker.Sentence(),
			UserId:      i + 1,
			PrivacyType: enums.SubPrivate,
		}

		_, err := env.Posts.Insert(tempPost)

		// fmt.Printf("%+v\n", tempPost)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedUsers(db *sql.DB) {
	env := models.CreateEnv(db)

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

		id, err := env.Users.Insert(tempUser)
		tempUser.Id = int(id)

		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("%+v\n", tempUser)
		// fmt.Println()

	}
}

func SeedComments(db *sql.DB) {
	env := models.CreateEnv(db)

	for i := 0; i < 10; i++ {
		tempComment := &models.Comment{
			Content: faker.Sentence(),
			UserId:  i + 1,
			PostId:  10 - i,
		}

		id, err := env.Comments.Insert(tempComment)

		tempComment.Id = int(id)

		// fmt.Printf("%+v\n", tempComment)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func SeedGroups(db *sql.DB) {
	env := models.CreateEnv(db)

	for i := 0; i < 10; i++ {
		tempGroup := &models.Group{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			CreatorId:   i + 1,
		}

		_, err := env.Groups.Insert(tempGroup)

		// fmt.Printf("%+v\n", tempGroup)

		if err != nil {
			fmt.Println(err)
		}

	}
}
