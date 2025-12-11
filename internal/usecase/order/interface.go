package order

import (
	"context"
	"main/internal/entity"
	"main/internal/services/order"
)

type Order interface {
	Create(ctx context.Context, data order.Create) (order.Get, error)
	GetList(ctx context.Context, clientId int) ([]order.Get, error)
	Update(ctx context.Context, data order.Update) error
	UpdateDriverLocation(ctx context.Context, data order.UpdateDriverLocation) error
}

type Auth interface {
	HashPassword(password string) (string, error)
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
}

type Websocket interface {
	SendToDriver(driverID int, event string, data any)
	BroadcastExcept(exceptID int, event string, data any)
	Broadcast(event string, data any)
	SendToClient(clientID int, event string, data any)
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, dec interface{}) error
	Delete(ctx context.Context, key string) error
}
