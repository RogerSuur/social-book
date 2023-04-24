package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Follower struct {
	Id          int
	FollowingId int
	FollowerId  int
	Accepted    bool
	Active      bool
}

type IFollowerRepository interface {
	GetById(id int64) (*Follower, error)
	Insert(follower *Follower) (int64, error)
	Update(follower *Follower) error
}

type FollowerRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewFollowerRepo(db *sql.DB) *FollowerRepository {
	return &FollowerRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo FollowerRepository) Insert(follower *Follower) (int64, error) {
	query := `INSERT INTO followers (following_id, follower_id, accepted, active)
	VALUES(?, ?, ?, ?)`

	args := []interface{}{
		follower.FollowingId,
		follower.FollowerId,
		follower.Accepted,
		follower.Active,
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

	repo.Logger.Printf("Inserted follower %d to start following %d (Last insert ID: %d)", follower.FollowerId, follower.FollowingId, lastId)

	return lastId, nil
}

func (repo FollowerRepository) Update(follower *Follower) error {
	query := `UPDATE followers SET accepted = ?, active = ? WHERE id = ?`

	args := []interface{}{
		follower.Accepted,
		follower.Active,
		follower.Id,
	}

	_, err := repo.DB.Exec(query, args...)

	return err
}

func (repo FollowerRepository) GetById(id int64) (*Follower, error) {
	query := `SELECT id, following_id,  follower_id, accepted, active FROM followers WHERE id = ?`
	row := repo.DB.QueryRow(query, id)
	follower := &Follower{}

	err := row.Scan(&follower.Id, &follower.FollowingId, &follower.FollowerId, &follower.Accepted, &follower.Active)

	return follower, err
}
