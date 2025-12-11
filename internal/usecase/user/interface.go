package user

import (
	"context"
	"main/internal/entity"
	"main/internal/services/user"
	"mime/multipart"
)

type User interface {
	CreateUser(ctx context.Context, data user.Create) (entity.User, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data user.Update) (entity.User, error)
}

type Auth interface {
	HashPassword(password string) (string, error)
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (entity.File, error)
	Delete(ctx context.Context, url string) error
}
