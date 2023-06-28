package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Message struct {
	Id            int64     `json:"id"`
	SenderId      int64     `json:"sender_id"`
	SenderName    string    `json:"sender_name"`
	RecipientId   int64     `json:"receiver_id"`
	RecipientName string    `json:"receiver_name"`
	GroupId       int64     `json:"group_id"`
	GroupName     string    `json:"group_name"`
	Content       string    `json:"body"`
	Timestamp     time.Time `json:"timestamp"`
}

type IMessageRepository interface {
	Insert(event *Message) (int64, error)
	GetChatUsers(id int64) ([]*User, error)
	GetLastMessage(userId int64, otherId int64) (*Message, error)
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

func (repo MessageRepository) GetChatUsers(id int64) ([]*User, error) {
	//TODO
	//get users from database
	return nil, nil
}

func (repo MessageRepository) GetLastMessage(userId int64, otherId int64) (*Message, error) {
	//TODO
	//get last message from database
	return nil, nil
}
