package wishlist

import (
	"context"
	"log"
	"main/internal/services/wishlist"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) GetList(ctx context.Context, userId int64) ([]wishlist.GetList, error) {
	log.Println("user", userId)
	var wishlist []wishlist.GetList

	query := `SELECT
				wl.id,
				wl.product_id,
				wl.created_at,
				wl.user_id,
				p.name,
				p.price,
				p.images,
				p.views_count,
				p.discount_percent,
				p.category_id,
				p.rating,
				p.description
				FROM wishlists wl
				LEFT JOIN products p ON wl.product_id = p.id
				WHERE wl.user_id = ? AND wl.deleted_at IS NULL AND wl.status = true AND p.deleted_at IS NULL`

	rows, err := r.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = r.ScanRows(ctx, rows, &wishlist); err != nil {
		return nil, err
	}

	return wishlist, err
}

func (r Repository) Create(ctx context.Context, productId wishlist.Create, userId int64) (int64, error) {
	var id int64

	query := `INSERT INTO wishlists (product_id, user_id, created_at, created_by) VALUES (?, ?, NOW(), ?) RETURNING id`
	if err := r.QueryRowContext(ctx, query, productId.ProductId, userId, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) Delete(ctx context.Context, wishlistId int64, id int64) error {
	query := `UPDATE wishlists SET deleted_at = NOW(), deleted_by = ? WHERE id = ? AND deleted_at IS NULL`
	if _, err := r.ExecContext(ctx, query, id, wishlistId); err != nil {
		return err
	}

	return nil
}
