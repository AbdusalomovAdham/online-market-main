package order

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, data Create) (Get, error) {
	return s.repo.Create(ctx, data)
}

func (s *Service) GetList(ctx context.Context, clientId int) ([]Get, error) {
	return s.repo.GetList(ctx, clientId)
}

func (s *Service) Update(ctx context.Context, data Update) error {
	return s.repo.Update(ctx, data)
}

func (s *Service) UpdateDriverLocation(ctx context.Context, data UpdateDriverLocation) error {
	return s.repo.UpdateDriverLocation(ctx, data)
}
