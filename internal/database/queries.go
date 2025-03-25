package database

import (
	"context"

	"github.com/google/uuid"
)

func (q *Queries) GetUserNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, `
        SELECT name 
        FROM users 
        WHERE id = $1
    `, id)

	var name string
	err := row.Scan(&name)
	return name, err
}
