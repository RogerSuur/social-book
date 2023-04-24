package models

import (
	"database/sql"
	"log"
	"os"
)

type Message struct {
	// TODO
}

type IMessageRepository interface {
	Insert(event *Message) (int64, error)
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

func (repo MessageRepository) Insert(event *Message) (int64, error) {

	//TODO
	//insert new message into database
	return 0, nil
}
