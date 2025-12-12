package user

import (
	"context"
	"database/sql"
	"main/internal/entity"
	"main/internal/services/auth"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, user *auth.Create, phoneNumber string) error {
	var id int64
	query := `INSERT INTO users (first_name, last_name, created_at) VALUES (?, ?, ?, NOW()) WHERE phone_number = ? RETURNING id`

	err := r.QueryRowContext(ctx, query, user.FirstName, user.LastName, phoneNumber).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetOrCreateUserByPhone(ctx context.Context, phone string) (int64, error) {
	var id int64

	query := `SELECT id FROM users WHERE phone_number = ? LIMIT 1`
	err := r.QueryRowContext(ctx, query, phone).Scan(&id)
	if err == nil {
		return id, nil
	}

	if err != sql.ErrNoRows {
		return 0, err
	}

	insertQuery := `INSERT INTO users (phone_number) VALUES (?)`
	res, err := r.ExecContext(ctx, insertQuery, phone)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r Repository) GetById(ctx context.Context, id int) (entity.User, error) {
	var detail entity.User
	query := `SELECT id, first_name, last_name, phone_number FROM users WHERE id = ?`
	err := r.QueryRowContext(ctx, query, id).Scan(&detail.Id, &detail.FirstName, &detail.LastName, &detail.PhoneNumber)
	if err != nil {
		return entity.User{}, err
	}
	return detail, nil
}
