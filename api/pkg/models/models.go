package models

import "database/sql"

type Env struct {
	Users         UserModel
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
