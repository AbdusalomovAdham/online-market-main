package user

import (
	"context"
	"main/internal/entity"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

// CREATE
func (r Repository) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (login, password, age)
	          VALUES ($1, $2, $3) RETURNING id`

	return r.DB.QueryRowContext(ctx, query,
		u.Login, u.Password, u.Age,
	).Scan(&u.ID)
}

// READ: by ID
func (r Repository) GetByID(ctx context.Context, id int64) (User, error) {
	var usr User
	query := `SELECT id, login, password, age FROM users WHERE id = $1`

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&usr.ID, &usr.Login, &usr.Password, &usr.Age,
	)
	return usr, err
}

func (r Repository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("email = ?", email).Scan(ctx)

	return detail, err
}

// UPDATE
func (r Repository) Update(ctx context.Context, u User) error {
	query := `UPDATE users SET login=$1, password=$2, age=$3 WHERE id=$4`

	_, err := r.DB.ExecContext(ctx, query,
		u.Login, u.Password, u.Age, u.ID,
	)
	return err
}

// DELETE
func (r Repository) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx,
		"DELETE FROM users WHERE id = $1", id)
	return err
}

// SORT + FILTER (age)
func (r Repository) ListUsers(ctx context.Context, minAge int, orderBy string) ([]User, error) {
	var list []User

	query := `SELECT id, login, password, age
	          FROM users WHERE age >= $1 ORDER BY ` + orderBy

	rows, err := r.DB.QueryContext(ctx, query, minAge)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Login, &u.Password, &u.Age); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, rows.Err()
}

// PAGINATION
func (r Repository) GetUsersPage(ctx context.Context, limit, offset int) ([]User, error) {
	var list []User

	query := `
		SELECT id, login, password, age
		FROM users
		ORDER BY id
		LIMIT $1 OFFSET $2
	`
	rows, err := r.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Login, &u.Password, &u.Age); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, rows.Err()
}

// JOIN (User → Orders)
func (r Repository) GetOrdersWithUser(ctx context.Context, userID int64) ([]Order, error) {
	var list []Order

	query := `
		SELECT o.id, o.user_id, o.amount
		FROM orders o
		JOIN users u ON u.id = o.user_id
		WHERE u.id = $1
	`
	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Amount); err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}
