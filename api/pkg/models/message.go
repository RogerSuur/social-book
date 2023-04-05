package models

import "database/sql"

type Message struct {
	// TODO
}

type MessageModel struct {
	DB *sql.DB
}

func (g MessageModel) Insert(event *Message) (int64, error) {

	//TODO
	//insert new message into database
	return 0, nil
}
