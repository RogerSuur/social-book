package models

import (
	"database/sql"
)

type GroupEventAttendance struct {
	UserId int
	PostId int
}

type GroupEventAttendanceModel struct {
	DB *sql.DB
}

func (repo GroupEventAttendanceModel) Insert(attendance *GroupEventAttendance) (int64, error) {
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
