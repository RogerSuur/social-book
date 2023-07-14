package models

import (
	"database/sql"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo             *UserRepository
	SessionRepo          *SessionRepository
	FollowerRepo         *FollowerRepository
	PostRepo             *PostRepository
	CommentRepo          *CommentRepository
	GroupRepo            *GroupRepository
	EventRepo            *EventRepository
	MessageRepo          *MessageRepository
	NotificationRepo     *NotificationRepository
	GroupMemberRepo      *GroupMemberRepository
	AllowedPostRepo      *AllowedPostRepository
	GroupEventAttendance *GroupEventAttendanceRepository
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	userRepo := NewUserRepo(db)
	sessionRepo := NewSessionRepo(db)
	followerRepo := NewFollowerRepo(db)
	postRepo := NewPostRepo(db)
	commentRepo := NewCommentRepo(db)
	groupRepo := NewGroupRepo(db)
	eventRepo := NewEventRepo(db)
	messageRepo := NewMessageRepo(db)
	notificationRepo := NewNotificationRepo(db)
	groupMemberRepo := NewGroupMemberRepo(db)
	allowedPostRepo := NewAllowedPostRepo(db)
	groupEventAttendance := NewGroupEventAttendanceRepo(db)

	return &Repositories{
		UserRepo:             userRepo,
		SessionRepo:          sessionRepo,
		FollowerRepo:         followerRepo,
		PostRepo:             postRepo,
		CommentRepo:          commentRepo,
		GroupRepo:            groupRepo,
		EventRepo:            eventRepo,
		MessageRepo:          messageRepo,
		NotificationRepo:     notificationRepo,
		GroupMemberRepo:      groupMemberRepo,
		AllowedPostRepo:      allowedPostRepo,
		GroupEventAttendance: groupEventAttendance,
	}
}
