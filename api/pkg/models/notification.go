package models

import (
	"database/sql"
	"log"
	"os"
)

type Notification struct {
	ReceiverId      int64
	FollowRequestId int64
	GroupInviteId   int64
	GroupRequestId  int64
	EventId         int64
	Reaction        bool
}

type INotificationRepository interface {
	Insert(notification *Notification) (int64, error)
}

type NotificationRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewNotificationRepo(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo NotificationRepository) Insert(notification *Notification) (int64, error) {
	query := `INSERT INTO notifications (receiver_id, follow_request_id, group_invite_id, group_request_id, event_id, reaction)
	VALUES(?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		notification.ReceiverId,
		notification.FollowRequestId,
		notification.GroupInviteId,
		notification.GroupRequestId,
		notification.EventId,
		notification.Reaction,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Inserted notification for user %d (last insert ID: %d)", notification.ReceiverId, lastId)

	return lastId, nil
}
