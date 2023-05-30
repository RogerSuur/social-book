package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Notification struct {
	Id                    int
	ReceiverId            int
	NotificationDetailsId int
	SeenAt                time.Time
	Reaction              bool
}

type NotificationDetails struct {
	Id                    int
	SenderId              int
	NotificationDetailsId int
	EntityId              int
	CreatedAt             time.Time
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

	//TODO
	//insert new notification into database
	return 0, nil
}
