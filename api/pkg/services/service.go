package services

import (
	"SocialNetworkRestApi/api/pkg/models"
	"database/sql"
)

type Service struct {
	DB  *sql.DB
	Env models.Env
}
