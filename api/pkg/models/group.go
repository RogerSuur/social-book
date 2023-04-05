package models

import (
	"database/sql"
	"time"
)

type Group struct {
	Id          int
	CreatorId   int
	Title       string
	Description string
	CreatedAt   time.Time
}

type GroupModel struct {
	DB *sql.DB
}

func (m GroupModel) Insert(group *Group) (int64, error) {

	query := `INSERT INTO groups (creator_id, title, description, created_at)
	VALUES(?, ?, ?, ?)`

	args := []interface{}{
		group.CreatorId,
		group.Title,
		group.Description,
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

func (p GroupModel) GetById(id int64) (*Group, error) {
	query := `SELECT id, creator_id,  title, description, created_at FROM groups WHERE id = ?`
	row := p.DB.QueryRow(query, id)
	group := &Group{}

	err := row.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt)

	return group, err
}

func (m GroupModel) GetAllByCreatorId(userId int) ([]*Group, error) {

	stmt := `SELECT id, creator_id,  title, description, created_at FROM groups
	WHERE creator_id = ?
    ORDER BY title ASC`

	rows, err := m.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := []*Group{}

	for rows.Next() {
		group := &Group{}

		err := rows.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (g GroupModel) GetAllByMemberId(userId int) ([]*Group, error) {

	//TODO
	//Get all groups the user is member of, by id
	return nil, nil
}
