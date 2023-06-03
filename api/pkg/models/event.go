package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Event struct {
	Id          int
	GroupId     int
	UserId      int
	CreatedAt   time.Time
	EventTime   time.Time
	TimeSpan    time.Duration
	Title       string
	Description string
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
	query := `INSERT INTO group_events (group_id, user_id, created_at, event_time, timespan, title, description)
	VALUES(?, ?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		event.GroupId,
		event.UserId,
		time.Now(),
		event.EventTime,
		event.TimeSpan,
		event.Title,
		event.Description,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Last inserted event '%s' by user %d (last insert ID: %d)", event.Title, event.UserId, lastId)

	return lastId, nil
}

func (repo EventRepository) GetAllByGroupId(id int) ([]*Event, error) {

	query := `SELECT id, group_id, user_id, created_at, event_time, timespan, title, description FROM group_events WHERE group_id = ?`

	rows, err := repo.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	events := []*Event{}

	for rows.Next() {
		event := &Event{}

		err := rows.Scan(&event.Id, &event.GroupId, &event.UserId, &event.CreatedAt, &event.EventTime, &event.TimeSpan, &event.TimeSpan, &event.Title, &event.Description)
		if err != nil {
			return nil, err
		}
	}

	repo.Logger.Printf("Found %d events for group %d", len(events), id)

	return events, err
}

func (repo EventRepository) GetAllByUserId(userId int) ([]*Event, error) {
	//TODO
	//Get all events by attending user

	return nil, nil
}
