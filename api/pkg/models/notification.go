package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Notification struct {
	ReceiverId       int64
	NotificationType string
	SenderID         int64
	EntityId         int64
	CreatedAt        time.Time `json:"created_at"`
	SeenAt           time.Time `json:"seen_at"`
	Reaction         bool
}

type INotificationRepository interface {
	GetNotificationType(notificationType string) (int64, error)
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

func (repo NotificationRepository) GetNotificationType(notificationType string) (int64, error) {
	query := `SELECT id FROM notification_types WHERE name = ?`

	args := []interface{}{
		notificationType,
	}

	var id int64

	err := repo.DB.QueryRow(query, args...).Scan(&id)

	if err != nil {
		repo.Logger.Printf("Error getting notification type: %s", err.Error())
		return -1, err
	}

	return id, nil
}

func (repo NotificationRepository) Insert(notification *Notification) (int64, error) {

	NotificationTypeID, err := repo.GetNotificationType(notification.NotificationType)

	if err != nil {
		repo.Logger.Printf("Error getting notification type: %s", err.Error())
		return -1, err
	}

	query := `INSERT INTO notification_details (sender_id, notification_type_id, entity_id, created_at)
	VALUES(?, ?, ?, ?)`

	args := []interface{}{
		notification.SenderID,
		NotificationTypeID,
		notification.EntityId,
		notification.CreatedAt,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		repo.Logger.Printf("Error inserting notification details: %s", err.Error())
		return -1, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		repo.Logger.Printf("Error getting last insert ID: %s", err.Error())
		return -1, err
	}

	query = `INSERT INTO notifications (receiver_id, notification_details_id, seen_at, reaction)
	VALUES(?, ?, ?, ?)`

	args = []interface{}{
		notification.ReceiverId,
		lastId,
		notification.SeenAt,
		notification.Reaction,
	}

	result, err = repo.DB.Exec(query, args...)

	if err != nil {
		repo.Logger.Printf("Error inserting notification: %s", err.Error())
		return -1, err
	}

	lastId, err = result.LastInsertId()

	if err != nil {
		return -1, err
	}

	repo.Logger.Printf("Inserted notification for user %d (last insert ID: %d)", notification.ReceiverId, lastId)

	return lastId, nil
}
