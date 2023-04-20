package models

import "database/sql"

type Env struct {
	Users interface {
		Insert(*User) (int64, error)
		Update(*User) error
		GetById(id int) (*User, error)
		GetByEmail(email string) (*User, error)
		GetByUserName(userName string) (*User, error)
		GetAllUserFollowers(id int) ([]*User, error)
		GetAllFollowedBy(id int) ([]*User, error)
	}
	Followers     FollowerModel
	Posts         PostModel
	Comments      CommentModel
	Groups        GroupModel
	Events        EventModel
	Messages      MessageModel
	Notifications NotificationModel
	Sessions      SessionModel
}

func CreateEnv(db *sql.DB) Env {
	return Env{
		Users:         UserModel{DB: db},
		Posts:         PostModel{DB: db},
		Comments:      CommentModel{DB: db},
		Groups:        GroupModel{DB: db},
		Events:        EventModel{DB: db},
		Messages:      MessageModel{DB: db},
		Notifications: NotificationModel{DB: db},
		Followers:     FollowerModel{DB: db},
		Sessions:      SessionModel{DB: db},
	}
}
