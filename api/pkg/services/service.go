package services

import (
	"database/sql"
)

type Service struct {
	DB *sql.DB
}
