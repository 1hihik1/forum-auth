package usecase

import (
	"fmt"
	"github.com/DrusGalkin/auth-grpc/internal/entity"
	"github.com/DrusGalkin/auth-grpc/pkg/auth"
)

func (uc *UserUseCase) Authenticate(email, password string) (string, string, int64, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		if err == entity.ErrorUserNotFound {
			return "", "", 0, entity.ErrorWrongPassword
		}
		return "", "", 0, err
	}

	if !user.Active {
		return "", "", 0, fmt.Errorf("пользователь не активен")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", "", 0, entity.ErrorWrongPassword
	}

	accessToken, expiresIn, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", 0, fmt.Errorf("ошибка генерации access токена: %w", err)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", 0, fmt.Errorf("ошибка генерации refresh токена: %w", err)
	}

	return accessToken, refreshToken, expiresIn, nil
}
