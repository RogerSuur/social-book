package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Message struct {
	Id          int64
	SenderId    int64
	RecipientId int64
	GroupId     int64
	Content     string
	//ImagePath   string
	SentAt time.Time
	ReadAt time.Time
}

type IMessageRepository interface {
	Insert(event *Message) (int64, error)
	GetMessagesByGroupId(groupId int64) ([]*Message, error)
	GetMessagesByUserIds(userId int64, secondUserId int64) ([]*Message, error)
	GetChatUsers(id int64) ([]*User, error)
	GetChatGroups(id int64) ([]*Group, error)
	GetLastMessage(userId int64, otherId int64, isGroup bool) (*Message, error)
	GetUnreadCount(userId int64, otherId int64, isGroup bool) (int64, error)
}

type MessageRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepository {
	return &MessageRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo MessageRepository) Insert(message *Message) (int64, error) {
	query := `INSERT INTO messages (sender_id, recipient_id, group_id, content, sent_at)
	VALUES(?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		message.SenderId,
		message.RecipientId,
		message.GroupId,
		message.Content,
		message.SentAt,
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Inserted message by user %d, to group/user %d/%d (last insert ID: %d)", message.SenderId, message.GroupId, message.RecipientId, lastId)

	return lastId, nil
}
func (repo MessageRepository) Update(message *Message) error {
	//TODO
	//Update method needed when readAt is being used
	return nil
}

func (repo MessageRepository) GetMessagesByGroupId(groupId int64) ([]*Message, error) {
	stmt := `SELECT id, sender_id, recipient_id, group_id, content, sent_at, read_at FROM messages m
	WHERE group_id = ?
    ORDER BY sent_at DESC`

	rows, err := repo.DB.Query(stmt, groupId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		message := &Message{}

		err := rows.Scan(&message.Id, &message.SenderId, &message.RecipientId, &message.GroupId, &message.Content, &message.SentAt, &message.ReadAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo MessageRepository) GetMessagesByUserIds(userId int64, secondUserId int64) ([]*Message, error) {
	stmt := `SELECT id, sender_id, recipient_id, group_id, content, sent_at, read_at FROM messages m
	WHERE (sender_id = ? AND recipient_id = ?) OR (sender_id = ? AND recipient_id = ?) 
    ORDER BY sent_at DESC`

	args := []interface{}{
		userId,
		secondUserId,
		secondUserId,
		userId,
	}

	rows, err := repo.DB.Query(stmt, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		message := &Message{}

		err := rows.Scan(&message.Id, &message.SenderId, &message.RecipientId, &message.GroupId, &message.Content, &message.SentAt, &message.ReadAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo MessageRepository) GetChatUsers(id int64) ([]*User, error) {
	//get all users from database except myself

	query := `SELECT id, forname, surname, nickname, image_path, created_at FROM users WHERE id != ?`

	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Nickname,
			&user.ImagePath,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo MessageRepository) GetChatGroups(id int64) ([]*Group, error) {
	//get all groups from user_groups where user is member and title and image_path from groups

	query := `SELECT g.id, g.title, g.created_at, g.image_path FROM groups g JOIN user_groups ug ON g.id = ug.group_id WHERE ug.user_id = ?`

	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := []*Group{}

	for rows.Next() {
		group := &Group{}
		err := rows.Scan(
			&group.Id,
			&group.Title,
			&group.CreatedAt,
			&group.ImagePath,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (repo MessageRepository) GetLastMessage(userId int64, otherId int64, isGroup bool) (*Message, error) {
	//get last message from database
	var query string
	var args []interface{}

	if isGroup {
		query = `SELECT id, sender_id, recipient_id, group_id, content, sent_at, read_at FROM messages WHERE group_id = ? ORDER BY sent_at DESC LIMIT 1`
		args = []interface{}{
			otherId,
		}
	} else {
		query = `SELECT id, sender_id, recipient_id, group_id, content, sent_at, read_at FROM messages WHERE (sender_id = ? AND recipient_id = ?) OR (sender_id = ? AND recipient_id = ?) ORDER BY sent_at DESC LIMIT 1`
		args = []interface{}{
			userId,
			otherId,
			otherId,
			userId,
		}
	}

	row := repo.DB.QueryRow(query, args...)

	message := &Message{}

	err := row.Scan(&message.Id, &message.SenderId, &message.RecipientId, &message.GroupId, &message.Content, &message.SentAt, &message.ReadAt)

	if err == sql.ErrNoRows {
		return &Message{}, nil
	}

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (repo MessageRepository) GetUnreadCount(userId int64, otherId int64, isGroup bool) (int64, error) {
	var query string
	var args []interface{}

	if isGroup {
		query = `SELECT COUNT(*) FROM messages WHERE group_id = ? AND read_at IS NULL`
		args = []interface{}{
			otherId,
		}
	} else {
		query = `SELECT COUNT(*) FROM messages WHERE (sender_id = ? AND recipient_id = ?) AND read_at IS NULL`
		args = []interface{}{
			otherId,
			userId,
		}
	}

	row := repo.DB.QueryRow(query, args...)

	var count int64

	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
