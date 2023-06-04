package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Message struct {
	Id          int
	SenderId    int
	RecipientId int
	GroupId     int
	Content     string
	ImagePath   string
	SentAt      time.Time
	ReadAt      time.Time
}

type IMessageRepository interface {
	Insert(event *Message) (int64, error)
	GetMessagesByGroupId(groupId int) ([]*Message, error)
	GetMessagesByUserIds(userId int, secondUserId int) ([]*Message, error)
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
	query := `INSERT INTO messages (sender_id, recipient_id, group_id, content, image_path, sent_at)
	VALUES(?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		message.SenderId,
		message.RecipientId,
		message.GroupId,
		message.Content,
		message.ImagePath,
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
	//Update methid needed when readAt is being used
	return nil
}

func (repo MessageRepository) GetMessagesByGroupId(groupId int) ([]*Message, error) {
	stmt := `SELECT id, sender_id, recipient_id, group_id, content, image_path, sent_at, read_at FROM messages m
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

		err := rows.Scan(&message.Id, &message.SenderId, &message.RecipientId, &message.GroupId, &message.Content, &message.ImagePath, &message.SentAt, &message.ReadAt)
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

func (repo MessageRepository) GetMessagesByUserIds(userId int, secondUserId int) ([]*Message, error) {
	stmt := `SELECT id, sender_id, recipient_id, group_id, content, image_path, sent_at, read_at FROM messages m
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

		err := rows.Scan(&message.Id, &message.SenderId, &message.RecipientId, &message.GroupId, &message.Content, &message.ImagePath, &message.SentAt, &message.ReadAt)
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
