package auth

import (
	"context"
	"main/internal/entity"
	"main/internal/services/auth"
	"main/internal/services/user"
	//"main/internal/services/email"
)

type Auth interface {
	GenerateToken(ctx context.Context, data auth.GenerateToken) (string, error)
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	GenerateResetToken(n int) (string, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, dec interface{}) error
	Delete(ctx context.Context, key string) error
}

type Email interface {
	SendMailSimple(subject, body string, to []string) error
	GenerateCode(n int) string
}

type User interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user user.Create) (entity.User, error)
}
