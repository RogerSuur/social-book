package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type GroupMember struct {
	UserId   int64
	GroupId  int64
	JoinedAt time.Time
}

type GroupMemberJSON struct {
	GroupId int   `json:"groupId"`
	UserIds []int `json:"userIds"`
}

type IGroupMemberRepository interface {
	Insert(groupMember *GroupMember) (int64, error)
	Update(groupMember *GroupMember) error
	Delete(groupMember *GroupMember) error
	GetGroupMembersByGroupId(groupId int64) ([]*User, error)
	IsGroupMember(group_id int64, userId int64) (bool, error)
	GetById(id int64) (*GroupMember, error)
}

type GroupMemberRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewGroupMemberRepo(db *sql.DB) *GroupMemberRepository {
	return &GroupMemberRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo GroupMemberRepository) Insert(groupMember *GroupMember) (int64, error) {
	query := `INSERT INTO user_groups (user_id, group_id, joined_at)
	VALUES(?, ?, ?)`

	args := []interface{}{
		groupMember.UserId,
		groupMember.GroupId,
		groupMember.JoinedAt,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Last inserted groupuser '%d' for user %d in group %d", lastId, groupMember.UserId, groupMember.GroupId)

	return lastId, nil
}

func (repo GroupMemberRepository) Update(groupMember *GroupMember) error {
	query := `UPDATE user_groups SET joined_at = ?
	WHERE user_id = ? AND group_id = ?`

	args := []interface{}{
		groupMember.JoinedAt,
		groupMember.UserId,
		groupMember.GroupId,
	}

	_, err := repo.DB.Exec(query, args...)

	if err != nil {
		repo.Logger.Printf("Error updating groupuser for user %d in group %d: %s", groupMember.UserId, groupMember.GroupId, err)
		return err
	}

	return nil
}

func (repo GroupMemberRepository) Delete(groupMember *GroupMember) error {
	query := `DELETE FROM user_groups WHERE user_id = ? AND group_id = ?`

	args := []interface{}{
		groupMember.UserId,
		groupMember.GroupId,
	}

	_, err := repo.DB.Exec(query, args...)

	if err != nil {
		return err
	}

	repo.Logger.Printf("Deleted groupuser for user %d in group %d", groupMember.UserId, groupMember.GroupId)

	return nil
}

func (repo GroupMemberRepository) GetGroupMembersByGroupId(groupId int64) ([]*User, error) {
	query := `SELECT ug.user_id, u.forname, u.surname, u.nickname, u.image_path FROM user_groups ug
	LEFT JOIN users u ON
	u.id = ug.user_id
	WHERE group_id = ?`

	// INNER JOIN user_groups ug ON
	// g.id = ug.group_id
	rows, err := repo.DB.Query(query, groupId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groupMembers := []*User{}

	for rows.Next() {
		groupMember := &User{}

		err := rows.Scan(&groupMember.Id, &groupMember.FirstName, &groupMember.LastName, &groupMember.Nickname, &groupMember.ImagePath)
		if err != nil {
			return nil, err
		}
		groupMembers = append(groupMembers, groupMember)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groupMembers, nil
}

func (repo GroupMemberRepository) IsGroupMember(groupId int64, userId int64) (bool, error) {
	query := `SELECT COUNT(id) FROM user_groups
	WHERE user_id = ? AND group_id = ? AND joined_at IS NOT NULL
	GROUP BY group_id`

	args := []interface{}{
		userId,
		groupId,
	}

	row := repo.DB.QueryRow(query, args...)

	var result int

	err := row.Scan(&result)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if result > 0 {
		return true, err
	}

	return false, nil
}

func (repo GroupMemberRepository) GetById(id int64) (*GroupMember, error) {
	query := `SELECT user_id, group_id, joined_at FROM user_groups WHERE id = ?`

	row := repo.DB.QueryRow(query, id)

	groupMember := &GroupMember{}

	err := row.Scan(&groupMember.UserId, &groupMember.GroupId, &groupMember.JoinedAt)

	if err != nil {
		return nil, err
	}

	return groupMember, nil
}
