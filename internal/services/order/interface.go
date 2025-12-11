package order

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, data Create) (Get, error)
	GetList(ctx context.Context, clientId int) ([]Get, error)
	Update(ctx context.Context, data Update) error
	UpdateDriverLocation(ctx context.Context, data UpdateDriverLocation) error
}
