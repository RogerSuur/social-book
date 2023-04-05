package models

import "database/sql"

type Notification struct {
	// TODO
}

type NotificationModel struct {
	DB *sql.DB
}

func (g NotificationModel) Insert(event *Notification) (int64, error) {

	//TODO
	//insert new notification into database
	return 0, nil
}
