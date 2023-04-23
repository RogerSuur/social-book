package models

import "database/sql"

type Notification struct {
	// TODO
}

type INotificationRepository interface {
	Insert(notification *Notification) (int64, error)
}

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepo(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		DB: db,
	}
}

func (g NotificationRepository) Insert(notification *Notification) (int64, error) {

	//TODO
	//insert new notification into database
	return 0, nil
}
