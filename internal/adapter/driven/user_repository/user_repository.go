package userrepository

import (
	"app/internal/infra"
	"context"
)

type UserRepository struct {
	db *infra.PostgresDB
}

func NewUserRepository(db *infra.PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) IsPhoneValid(ctx context.Context, id, phone string) bool {
	var count int64
	err := ur.db.Conn().QueryRow(ctx, `
		SELECT
			COUNT(*)
		FROM
			user_whatsmeow_map
		WHERE
			user_id = $1
		AND
			phone = $2
		LIMIT 1
	`, id, phone).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
