package models

import (
	"database/sql"
	"log"
	"os"
)

type EventAttendance struct {
	UserId      int64
	EventId     int64
	IsAttending bool
}

type IEventAttendanceRepository interface {
	Insert(attendance *EventAttendance) (int64, error)
}

type EventAttendanceRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewEventAttendanceRepo(db *sql.DB) *EventAttendanceRepository {
	return &EventAttendanceRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo EventAttendanceRepository) Insert(attendance *EventAttendance) (int64, error) {
	query := `INSERT INTO group_event_attendance (user_id, event_id, is_attending)
	VALUES(?, ?, ?)`

	args := []interface{}{
		attendance.UserId,
		attendance.EventId,
		attendance.IsAttending,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return -1, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	repo.Logger.Printf("User %d added to attend event %d", attendance.UserId, attendance.EventId)

	return lastId, nil
}
