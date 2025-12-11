package user

import (
	"context"
	"errors"
	"main/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) IsValidToken(ctx context.Context, authHeader entity.User) (entity.User, error) {
	return s.IsValidToken(ctx, authHeader)
}

func (s Service) GetById(ctx context.Context, id int) (entity.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s Service) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s Service) CreateUser(ctx context.Context, data Create) (entity.User, error) {
	if data.Email == "" {
		return entity.User{}, errors.New("email is required")
	}
	return s.repo.Create(ctx, data)
}

func (s Service) Update(ctx context.Context, data Update) (entity.User, error) {
	return s.repo.Update(ctx, data)
}
