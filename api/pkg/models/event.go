package models

import (
	"database/sql"
	"log"
	"os"
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
	Logger *log.Logger
	DB     *sql.DB
}

func NewEventRepo(db *sql.DB) *EventRepository {
	return &EventRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo EventRepository) Insert(event *Event) (int64, error) {

	//TODO
	//insert new event into database
	return 0, nil
}

func (repo EventRepository) GetAllByGroupId(groupId int) ([]*Event, error) {

	//TODO
	//Get all events by group
	return nil, nil
}

func (repo EventRepository) GetAllByUserId(userId int) ([]*Event, error) {

	//TODO
	//Get all events by attending user
	return nil, nil
}
