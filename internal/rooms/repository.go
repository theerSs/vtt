package rooms

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct{}

func newRepository(db *pgxpool.Pool) *repository {
	return &repository{}
}
