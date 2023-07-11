package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type GroupMemberModel struct {
	UserId   int64
	GroupId  int64
	JoinedAt time.Time
}

type IGroupUserRepository interface {
	Insert(groupMember *GroupMemberModel) (int64, error)
	GetGroupMembersByGroupId(groupId int64) ([]*User, error)
	IsGroupMember(group_id int64, userId int64) (bool, error)
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

func (repo GroupUserRepository) Insert(groupMember *GroupMemberModel) (int64, error) {
	query := `INSERT INTO user_groups (user_id, group_id, joined_at)
	VALUES(?, ?, ?)`

	args := []interface{}{
		groupMember.UserId,
		groupMember.GroupId,
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

	repo.Logger.Printf("Last inserted groupuser '%d' for user %d in group %d", lastId, groupMember.UserId, groupMember.GroupId)

	return lastId, nil
}

func (repo GroupUserRepository) GetGroupMembersByGroupId(groupId int64) ([]*User, error) {
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

func (repo GroupUserRepository) IsGroupMember(groupId int64, userId int64) (bool, error) {
	query := `SELECT COUNT(id) FROM user_groups
	WHERE user_id = ? AND group_id = ?
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
