package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

type Service struct {
	repo    Repository
	auth    Auth
	cache   Cache
	sendSMS SendSMS
}

func NewService(repo Repository, cache Cache, sendSMS SendSMS, auth Auth) *Service {
	return &Service{
		repo:    repo,
		auth:    auth,
		cache:   cache,
		sendSMS: sendSMS,
	}
}

func GenerateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

func GenerateTokenUUID() string {
	return uuid.New().String()
}

func (s *Service) SendEmailCode(ctx context.Context, phone string) (int64, string, error) {
	code := GenerateOTP()
	token := GenerateTokenUUID()
	id, err := s.repo.GetOrCreateUserByPhone(ctx, phone)
	if err != nil {
		return 0, "", err
	}

	if err := s.sendSMS.SendSMS(phone, code, token); err != nil {
		return 0, "", err
	}

	if err := s.cache.Set(ctx, token, code); err != nil {
		return 0, "", err
	}
	return id, token, nil
}
