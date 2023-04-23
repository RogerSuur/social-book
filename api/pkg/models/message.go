package models

import "database/sql"

type Message struct {
	// TODO
}

type IMessageRepository interface {
	Insert(event *Message) (int64, error)
}

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (m MessageRepository) Insert(event *Message) (int64, error) {

	//TODO
	//insert new message into database
	return 0, nil
}
