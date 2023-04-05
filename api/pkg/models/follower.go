package models

import (
	"database/sql"
	"time"
)

type Follower struct {
	Id          int
	FollowingId int
	FollowerId  int
	Accepted    bool
	Active      bool
}

type FollowerModel struct {
	DB *sql.DB
}

func (m FollowerModel) Insert(follower *Follower) (int64, error) {
	query := `INSERT INTO followers (following_id, follower_id, accepted, active)
	VALUES(?, ?, ?, ?)`

	args := []interface{}{
		follower.FollowingId,
		follower.FollowerId,
		follower.Accepted,
		follower.Active,
		time.Now(),
	}

	result, err := m.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (m FollowerModel) Update(follower *Follower) error {
	query := `UPDATE followers SET accepted = ?, active = ? WHERE id = ?`

	args := []interface{}{
		follower.Accepted,
		follower.Active,
		follower.Id,
	}

	_, err := m.DB.Exec(query, args...)

	return err
}

func (p FollowerModel) GetById(id int64) (*Follower, error) {
	query := `SELECT id, following_id,  follower_id, accepted, active FROM followers WHERE id = ?`
	row := p.DB.QueryRow(query, id)
	follower := &Follower{}

	err := row.Scan(&follower.Id, &follower.FollowingId, &follower.FollowerId, &follower.Accepted, &follower.Active)

	return follower, err
}
