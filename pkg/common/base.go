package common

import "github.com/jackc/pgx/v5/pgxpool"

// BaseCapsule database
type BaseCapsule struct {
	Database *pgxpool.Pool
}

func NewBaseRepository(dbp *pgxpool.Pool) *BaseCapsule {
	return &BaseCapsule{
		Database: dbp,
	}
}
