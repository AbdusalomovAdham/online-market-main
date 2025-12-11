package user

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.User, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	Delete(ctx context.Context, id int) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Update(ctx context.Context, data Update) (entity.User, error)
}
