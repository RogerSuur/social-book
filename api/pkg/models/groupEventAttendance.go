package models

import (
	"database/sql"
	"log"
	"os"
)

type GroupEventAttendance struct {
	UserId int
	PostId int
}

type IGroupEventAttendanceRepository interface {
	Insert(attendance *GroupEventAttendance) (int64, error)
}

type GroupEventAttendanceRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewGroupEventAttendanceRepo(db *sql.DB) *GroupEventAttendanceRepository {
	return &GroupEventAttendanceRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo GroupEventAttendanceRepository) Insert(attendance *GroupEventAttendance) (int64, error) {
	query := `INSERT INTO group_event_attendance (user_id, event_id, is_attending)
	VALUES(?, ?)`

	args := []interface{}{
		attendance.PostId,
		attendance.UserId,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
