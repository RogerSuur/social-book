package models

import (
	"database/sql"
)

type Event struct {
	// TODO
}

type IEventRepository interface {
	GetAllByGroupId(groupId int) ([]*Event, error)
	GetAllByUserId(userId int) ([]*Event, error)
	Insert(event *Event) (int64, error)
}

type EventRepository struct {
	DB *sql.DB
}

func NewEventRepo(db *sql.DB) *EventRepository {
	return &EventRepository{
		DB: db,
	}
}

func (g EventRepository) Insert(event *Event) (int64, error) {

	//TODO
	//insert new event into database
	return 0, nil
}

func (g EventRepository) GetAllByGroupId(groupId int) ([]*Event, error) {

	//TODO
	//Get all events by group
	return nil, nil
}

func (g EventRepository) GetAllByUserId(userId int) ([]*Event, error) {

	//TODO
	//Get all events by attending user
	return nil, nil
}
