package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type GroupUser struct {
	UserId   int64
	GroupId  int64
	JoinedAt time.Time
}

type IGroupUserRepository interface {
	Insert(groupUser *GroupUser) (int64, error)
}

type GroupUserRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewGroupUserRepo(db *sql.DB) *GroupUserRepository {
	return &GroupUserRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo GroupUserRepository) Insert(groupUser *GroupUser) (int64, error) {
	query := `INSERT INTO user_groups (user_id, group_id, joined_at)
	VALUES(?, ?, ?)`

	args := []interface{}{
		groupUser.UserId,
		groupUser.GroupId,
		time.Now(),
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Last inserted groupuser '%d' for user %d in group %d", lastId, groupUser.UserId, groupUser.GroupId)

	return lastId, nil
}
