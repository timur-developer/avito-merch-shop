package usecase

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/domain/errors"
	"avito-merch-shop/internal/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepository repository.UserRepository
	jwtKey         []byte
}

func NewAuthUsecase(userRepository repository.UserRepository, jwtKey string) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		jwtKey:         []byte(jwtKey),
	}
}

func (u *AuthUsecase) Register(ctx context.Context, username, password string) (*entity.User, error) {
	if username == "" || password == "" {
		return nil, errors.ErrInvalidCredentials
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServerError
	}

	user := &entity.User{
		Username:     username,
		PasswordHash: string(hash),
		Coins:        1000,
	}

	if err := u.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *AuthUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.ErrInvalidCredentials
	}

	// TODO Реализовать создание JWT токена
	return "jwt-token", nil
}
