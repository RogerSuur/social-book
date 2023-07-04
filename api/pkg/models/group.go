package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Group struct {
	Id          int
	CreatorId   int
	Title       string
	Description string
	ImagePath   string
	CreatedAt   time.Time
}

type UserGroup struct {
	Id    int    `json:"groupId"`
	Title string `json:"groupName"`
}

type GroupJSON struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"imagePath"`
}

type IGroupRepository interface {
	GetAllByCreatorId(userId int) ([]*Group, error)
	GetAllByMemberId(userId int) ([]*Group, error)
	GetById(id int64) (*Group, error)
	Insert(group *Group) (int64, error)
}

type GroupRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewGroupRepo(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo GroupRepository) Insert(group *Group) (int64, error) {

	query := `INSERT INTO groups (creator_id, title, description, created_at, image_path)
	VALUES(?, ?, ?, ?, ?)`

	args := []interface{}{
		group.CreatorId,
		group.Title,
		group.Description,
		time.Now(),
		group.ImagePath,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Last inserted group '%s' by user %d (last insert ID: %d)", group.Title, group.CreatorId, lastId)

	return lastId, nil
}

func (p GroupRepository) GetById(id int64) (*Group, error) {
	query := `SELECT id, creator_id,  title, description, created_at, image_path FROM groups WHERE id = ?`
	row := p.DB.QueryRow(query, id)
	group := &Group{}

	err := row.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt, &group.ImagePath)

	return group, err
}

func (repo GroupRepository) GetAllByCreatorId(userId int) ([]*Group, error) {

	stmt := `SELECT id, creator_id,  title, description, created_at, image_path FROM groups
	WHERE creator_id = ?
    ORDER BY title ASC`

	rows, err := repo.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := []*Group{}

	for rows.Next() {
		group := &Group{}

		err := rows.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt, &group.ImagePath)
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

func (repo GroupRepository) GetAllByMemberId(userId int) ([]*Group, error) {

	stmt := `SELECT DISTINCT g.id, g.creator_id,  g.title, g.description, g.created_at, g.image_path FROM groups g
	INNER JOIN user_groups ug ON
	g.id = ug.group_id
	WHERE ug.user_id = ?
    ORDER BY title ASC`

	rows, err := repo.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := []*Group{}

	for rows.Next() {
		group := &Group{}

		err := rows.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt, &group.ImagePath)
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
