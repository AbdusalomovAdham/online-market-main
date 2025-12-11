package auth

import (
	"context"
	"main/internal/entity"
)

type User interface {
	GetById(ctx context.Context, id int) (entity.User, error)
}
