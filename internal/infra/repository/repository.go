package repository

import (
	"database/sql"

	"github.com/devfullcycle/imersao10-consolidacao/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
