package dal

import (
	"context"
	"database/sql"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
)

type DAL struct {
	Db *sql.DB
	ctx context.Context
	tx  *sql.Tx
}


type AccessInterface interface {
	SelectUserByEmail(email string) (user *models.UserProfile, errVal error)
}

