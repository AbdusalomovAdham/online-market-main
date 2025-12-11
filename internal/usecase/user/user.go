package user

import (
	"context"
	"main/internal/entity"
	"main/internal/services/user"
	"mime/multipart"
)

type UseCase struct {
	user User
	auth Auth
	file File
}

func NewUseCase(user User, auth Auth, file File) *UseCase {
	return &UseCase{user: user, auth: auth, file: file}
}

func (uc UseCase) Update(ctx context.Context, data user.Update, authHeader string) (entity.User, error) {
	tokenClaims, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.User{}, err
	}

	detail, err := uc.user.GetById(ctx, tokenClaims.Id)
	if err != nil {
		return entity.User{}, err
	}

	if detail.Avatar.Path != "" {
		if err := uc.file.Delete(ctx, detail.Avatar.Path); err != nil {
			return entity.User{}, err
		}
	}

	return uc.user.Update(ctx, data)
}

func (uc UseCase) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (entity.File, error) {
	return uc.file.Upload(ctx, file, folder)
}

func (uc UseCase) Delete(ctx context.Context, authHeader string) error {
	tokenClaims, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}
	return uc.user.Delete(ctx, tokenClaims.Id)
}
