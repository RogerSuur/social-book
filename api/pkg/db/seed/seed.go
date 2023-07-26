package seed

import (
	"SocialNetworkRestApi/api/pkg/enums"
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
			Birthday:  date,
			ImagePath: seedUser.ImagePath,
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
						Content:   comments.Content,
						UserId:    commentUser.Id,
						PostId:    postId,
						CreatedAt: seedPost.CreatedAt.Add(comments.PostOffSet),
					}

					id, err := repos.CommentRepo.InsertSeedComment(tempComment)

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
			ImagePath:   group.ImagePath,
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
			tempGroupUser := &models.GroupMember{
				UserId:  groupUser.Id,
				GroupId: id,
			}

			_, err = repos.GroupMemberRepo.Insert(tempGroupUser)
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
				GroupId:      id,
				UserId:       eventCreator.Id,
				CreatedAt:    event.CreatedAt,
				EventTime:    event.EventTime,
				EventEndTime: event.EventTime.Add(event.TimeSpan),
				Title:        event.Title,
				Description:  event.Description,
			}
			_, err = repos.EventRepo.InsertSeedEvent(tempEvent)
			if err != nil {
				logger.Printf("%+v\n", err)
			}

		}

		//Add group posts
		for _, post := range group.SeedGroupPostsData {
			eventCreator, err := repos.UserRepo.GetByEmail(post.CreatorEmail)

			if err != nil {
				logger.Printf("%+v\n", err)
			}

			tempPost := &models.Post{
				GroupId:     id,
				UserId:      eventCreator.Id,
				Content:     post.Content,
				CreatedAt:   post.CreatedAt,
				PrivacyType: enums.PrivacyType(enums.None),
			}

			_, err = repos.PostRepo.InsertSeedPost(tempPost)
			if err != nil {
				logger.Printf("%+v\n", err)
			}
		}

	}

	// Add group event attendees
	for _, attendance := range SeedEventAttendanceDataAccepted {
		_, err := repos.EventAttendanceRepo.Insert(&models.EventAttendance{
			EventId:     attendance.EventId,
			UserId:      attendance.UserId,
			IsAttending: attendance.IsAttending,
		})

		if err != nil {
			logger.Printf("%+v\n", err)
		}
	}

	for _, invitation := range SeedEventAttendanceDataPending {
		event, err := repos.EventRepo.GetById(invitation.EventId)
		if err != nil {
			logger.Printf("%+v\n", err)
		}

		_, err = repos.NotificationRepo.Insert(&models.Notification{
			ReceiverId:       invitation.UserId,
			SenderId:         event.UserId,
			EntityId:         invitation.EventId,
			NotificationType: "event_invite",
			CreatedAt:        time.Now(),
		})
		if err != nil {
			logger.Printf("%+v\n", err)
		}

	}

}
