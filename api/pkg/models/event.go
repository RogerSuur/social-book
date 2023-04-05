package models

import (
	"database/sql"
)

type Event struct {
	// TODO
}

type EventModel struct {
	DB *sql.DB
}

func (g EventModel) Insert(event *Event) (int64, error) {

	//TODO
	//insert new event into database
	return 0, nil
}

func (g EventModel) GetAllByGroupId(groupId int) ([]*Event, error) {

	//TODO
	//Get all events by group
	return nil, nil
}

func (g EventModel) GetAllByUserId(userId int) ([]*Event, error) {

	//TODO
	//Get all events by attending user
	return nil, nil
}
