package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Event struct {
	Id          int64
	GroupId     int64
	UserId      int64
	CreatedAt   time.Time
	EventTime   time.Time
	TimeSpan    time.Duration
	Title       string
	Description string
}

type IEventRepository interface {
	GetAllByGroupId(groupId int64) ([]*Event, error)
	GetAllByUserId(userId int64) ([]*Event, error)
	Insert(event *Event) (int64, error)
	InsertSeedEvent(event *Event) (int64, error)
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

func (repo EventRepository) InsertSeedEvent(event *Event) (int64, error) {
	query := `INSERT INTO group_events (group_id, user_id, created_at, event_time, timespan, title, description)
	VALUES(?, ?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		event.GroupId,
		event.UserId,
		event.CreatedAt,
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

// Get all group events
func (repo EventRepository) GetAllByGroupId(id int64) ([]*Event, error) {

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

// Get all events by user
func (repo EventRepository) GetAllByUserId(id int64) ([]*Event, error) {
	//TODO
	//Get all events by attending user
	query := `SELECT DISTINCT ge.id, group_id, ge.user_id, ge.created_at, ge.event_time, ge.timespan, ge.title, ge.description FROM group_events ge
	INNER JOIN group_event_atenndance gea
	ON gea.group_id = ge.group_id
	WHERE gea.user_id = ? AND WHERE (gea.attending = true OR gea.attending IS NULL)`

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
