package wishlist

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

func NewService(repo Repository, auth Auth, cache Cache) *Service {
	return &Service{
		repo:  repo,
		auth:  auth,
		cache: cache,
	}
}

func (s *Service) GetList(ctx context.Context, authHeader string, lang string) ([]GetList, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	var dest []GetList
	cacheKey := "user_wishlist:" + strconv.FormatInt(isValidToken.Id, 10)
	if err := s.cache.Get(ctx, cacheKey, &dest); err == nil {
		return dest, nil
	}

	list, err := s.repo.GetList(ctx, isValidToken.Id, lang)
	if err != nil {
		return nil, err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		ctxBg := context.Background()
		if err := s.cache.Set(ctxBg, cacheKey, list); err != nil {
			fmt.Println("Redis set error:", err)
		}
	}()

	return list, nil
}

func (s *Service) Create(ctx context.Context, productId Create, authHeader string) (int64, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, productId, isValidToken.Id)
	if err != nil {
		return 0, err
	}

	cacheKey := "user_wishlist:" + strconv.FormatInt(isValidToken.Id, 10)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		ctxBg := context.Background()
		if err := s.cache.Delete(ctxBg, cacheKey); err != nil {
			fmt.Println("Redis delete error:", err)
		}
	}()

	return id, nil
}

func (s *Service) Delete(ctx context.Context, wishlistItemId int64, authHeader string) error {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	cacheKey := "user_wishlist:" + strconv.FormatInt(isValidToken.Id, 10)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Redis set panic:", r)
			}
		}()
		ctxBg := context.Background()
		if err := s.cache.Delete(ctxBg, cacheKey); err != nil {
			fmt.Println("Redis delete error:", err)
		}
	}()

	return s.repo.Delete(ctx, wishlistItemId, isValidToken.Id)
}
