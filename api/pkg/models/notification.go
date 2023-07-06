package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Notification struct {
	Id               int64
	ReceiverId       int64
	SenderId         int64
	NotificationType string
	EntityId         int64
	CreatedAt        time.Time
	SeenAt           time.Time
	Reaction         bool
}

type NotificationJSON struct {
	NotificationType string `json:"notification_type"`
	NotificationId   int64  `json:"id"`
	SenderId         int64  `json:"sender_id"`
	SenderName       string `json:"sender_name"`
	GroupId          int64  `json:"group_id"`
	GroupName        string `json:"group_name"`
	EventId          int64  `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDate        string `json:"event_datetime"`
}

type INotificationRepository interface {
	Insert(notification *Notification) (int64, error)
	Update(notification *Notification) error
	GetById(id int64) (*Notification, error)
	GetByReceiverId(receiverId int64) ([]*Notification, error)
	GetNotificationType(notificationType string) (int64, error)
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

	NotificationTypeID, err := repo.GetNotificationType(notification.NotificationType)

	if err != nil {
		repo.Logger.Printf("Error getting notification type: %s", err.Error())
		return -1, err
	}

	query := `INSERT INTO notification_details (sender_id, notification_type_id, entity_id, created_at)
	VALUES(?, ?, ?, ?)`

	args := []interface{}{
		notification.SenderId,
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

func (repo NotificationRepository) Update(notification *Notification) error {
	query := `UPDATE notifications SET seen_at = ?, reaction = ? WHERE id = ?`

	args := []interface{}{
		notification.SeenAt,
		notification.Reaction,
		notification.Id,
	}

	_, err := repo.DB.Exec(query, args...)

	if err != nil {
		repo.Logger.Printf("Error updating notification: %s", err.Error())
		return err
	}

	return nil
}

func (repo NotificationRepository) GetById(id int64) (*Notification, error) {
	query := `SELECT n.id, nd.sender_id, nd.created_at, n.seen_at, n.reaction FROM notifications n
	JOIN notification_details nd ON n.notification_details_id = nd.id
	WHERE n.id = ?`

	args := []interface{}{
		id,
	}

	var notification Notification

	err := repo.DB.QueryRow(query, args...).Scan(&notification.Id, &notification.SenderId, &notification.CreatedAt, &notification.SeenAt, &notification.Reaction)

	if err != nil {
		repo.Logger.Printf("Error getting notification: %s", err.Error())
		return nil, err
	}

	return &notification, nil
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

func (repo NotificationRepository) GetByReceiverId(userId int64) ([]*Notification, error) {
	/*
		query := `SELECT n.id, nt.name, nd.sender_id, nd.created_at, n.seen_at, n.reaction FROM notifications n
		JOIN notification_details nd ON n.notification_details_id = nd.id
		JOIN notification_types nt ON nd.notification_type_id = nt.id
		WHERE n.receiver_id = ?`
	*/

	query := `SELECT n.id, nt.name, nd.sender_id, n.reaction FROM notifications n
	JOIN notification_details nd ON n.notification_details_id = nd.id
	JOIN notification_types nt ON nd.notification_type_id = nt.id
	WHERE n.receiver_id = ?`

	args := []interface{}{
		userId,
	}

	rows, err := repo.DB.Query(query, args...)

	if err != nil {
		repo.Logger.Printf("Error getting notifications: %s", err.Error())
		return nil, err
	}

	defer rows.Close()

	notifications := []*Notification{}

	for rows.Next() {
		var notification Notification

		err := rows.Scan(&notification.Id, &notification.NotificationType, &notification.SenderId, &notification.Reaction)

		if err != nil {
			repo.Logger.Printf("Error scanning notification: %s", err.Error())
			return nil, err
		}

		notifications = append(notifications, &notification)
	}

	return notifications, nil
}
