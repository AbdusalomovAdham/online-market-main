package auth

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/services/auth"
	"main/internal/services/user"
)

type UseCase struct {
	auth  Auth
	cache Cache
	email Email
	user  User
}

func NewUseCase(auth Auth, cache Cache, email Email, user User) *UseCase {
	return &UseCase{auth: auth, cache: cache, email: email, user: user}
}

func (au UseCase) SendEmailCode(ctx context.Context, email string) (string, string, error) {
	emailCode := au.email.GenerateCode(6)

	if err := au.email.SendMailSimple("Email Verification Code", "Your verification code is: "+emailCode, []string{email}); err != nil {
		return "", "", err
	}

	token, err := au.auth.GenerateResetToken(16)
	if err != nil {
		return "", "", err
	}

	data := auth.ResendData{Email: email, Code: emailCode}
	if err := au.cache.Set(ctx, token, data); err != nil {
		return "", "", err
	}

	return token, email, nil
}

func (au UseCase) CheckCode(ctx context.Context, code, token string) (entity.User, string, error) {
	var data auth.ResendData
	if err := au.cache.Get(ctx, token, &data); err != nil {
		return entity.User{}, "", err
	}

	if code != data.Code {
		return entity.User{}, "", errors.New("code error")
	}

	detail, err := au.user.GetByEmail(ctx, data.Email)

	if err != nil {
		newUser := user.Create{
			Email: data.Email,
		}

		detail, err := au.user.CreateUser(ctx, newUser)
		if err != nil {
			return entity.User{}, token, err
		}

		tokenGenerator := auth.GenerateToken{
			Role: detail.Role,
			Id:   detail.Id,
		}

		if err := au.cache.Delete(ctx, token); err != nil {
			return entity.User{}, "", err
		}

		token, err = au.auth.GenerateToken(ctx, tokenGenerator)
		if err != nil {
			return entity.User{}, "", err
		}

		return detail, token, nil
	}

	if err := au.cache.Delete(ctx, token); err != nil {
		return entity.User{}, "", err
	}

	tokenGenerator := auth.GenerateToken{
		Role: detail.Role,
		Id:   detail.Id,
	}

	token, err = au.auth.GenerateToken(ctx, tokenGenerator)
	if err != nil {
		return entity.User{}, "", err
	}
	return detail, token, nil
}

func (au UseCase) ResendCode(ctx context.Context, token string) error {
	var resetData auth.ResendData

	if err := au.cache.Get(ctx, token, &resetData); err != nil {
		return err
	}
	code := au.email.GenerateCode(6)

	err := au.email.SendMailSimple("Password Reset Code", "Your reset code is: "+code, []string{resetData.Email})
	if err != nil {
		return err
	}

	resetData.Code = code
	if err := au.cache.Set(ctx, token, resetData); err != nil {
		return err
	}

	return nil
}
