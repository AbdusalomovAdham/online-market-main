package order

import (
	"context"
	"fmt"
	"strconv"
)

type Service struct {
	repo  Repository
	auth  Auth
	cache Cache
}

func NewService(repo Repository, auth Auth, cache Cache) Service {
	return Service{
		repo:  repo,
		auth:  auth,
		cache: cache,
	}
}

func (s *Service) Create(ctx context.Context, order Create, authHeader string) error {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	if err := s.repo.Create(ctx, order, isValidToken.Id); err != nil {
		return err
	}

	cacheKey := "user_order:" + strconv.FormatInt(isValidToken.Id, 10)
	go func() {
		ctxBg := context.Background()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		if err := s.cache.Delete(ctxBg, cacheKey); err != nil {
			fmt.Println("Redis delete error:", err)
		}
	}()

	return nil
}

func (s *Service) GetList(ctx context.Context, authHeader string, lang string) ([]Get, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	var dest []Get
	cacheKey := "user_order:" + strconv.FormatInt(isValidToken.Id, 10)

	if err := s.cache.Get(ctx, cacheKey, &dest); err == nil {
		return dest, nil
	}

	list, err := s.repo.GetList(ctx, isValidToken.Id, lang)
	if err != nil {
		return []Get{}, err
	}

	go func() {
		ctxBg := context.Background()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		if err := s.cache.Set(ctxBg, cacheKey, list); err != nil {
			fmt.Println("Redis set error:", err)
		}
	}()

	return list, err
}

func (s *Service) GetById(ctx context.Context, orderId int64, authHeader string) (Get, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return Get{}, err
	}

	cacheKey := "user_order:" + strconv.FormatInt(isValidToken.Id, 10)
	go func() {
		ctxBg := context.Background()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		if err := s.cache.Delete(ctxBg, cacheKey); err != nil {
			fmt.Println("Redis delete error:", err)
		}
	}()

	return s.repo.GetById(ctx, orderId, isValidToken.Id)
}

func (s *Service) Delete(ctx context.Context, orderId int64, authHeader string) error {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, orderId, isValidToken.Id); err != nil {
		return err
	}

	cacheKey := "user_order:" + strconv.FormatInt(isValidToken.Id, 10)
	go func() {
		ctxBg := context.Background()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		if err := s.cache.Delete(ctxBg, cacheKey); err != nil {
			fmt.Println("Redis delete error:", err)
		}
	}()

	return nil
}
