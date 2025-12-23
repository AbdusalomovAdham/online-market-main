package wishlist

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	GetList(ctx context.Context, userId int64, lang string) ([]GetList, error)
	Create(ctx context.Context, productId Create, userId int64) (int64, error)
	Delete(ctx context.Context, productId, userId int64) error
}

type Auth interface {
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
}
