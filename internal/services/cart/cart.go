package cart

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

func (s Service) Create(ctx context.Context, cart Create, authHeader string) (int64, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, cart.ProductId, isValidToken.Id)
	if err != nil {
		return 0, err
	}

	cacheKey := "user_cart:" + strconv.FormatInt(isValidToken.Id, 10)
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

	return id, nil
}

func (s Service) UpdateCartItemTotal(ctx context.Context, cartItemId int64, authHeader string) error {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	if err := s.repo.Update(ctx, cartItemId, isValidToken.Id); err != nil {
		return err
	}

	cacheKey := "user_cart:" + strconv.FormatInt(isValidToken.Id, 10)
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

func (s Service) DeleteCartItem(ctx context.Context, cartItemId int64, authHeader string) error {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteCartItem(ctx, cartItemId, isValidToken.Id); err != nil {
		return err
	}

	cacheKey := "user_cart:" + strconv.FormatInt(isValidToken.Id, 10)
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

func (s Service) GetList(ctx context.Context, authHeader string) ([]Get, error) {
	isValidToken, err := s.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	var dest []Get
	cacheKey := "user_cart:" + strconv.FormatInt(isValidToken.Id, 10)

	if err := s.cache.Get(ctx, cacheKey, &dest); err == nil {
		return dest, nil
	}

	list, err := s.repo.GetList(ctx, isValidToken.Id)
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

	return list, nil
}
