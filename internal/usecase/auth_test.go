package usecase

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/domain/errors"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *mockUserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *mockUserRepository) UpdateCoins(ctx context.Context, id int, coins int) error {
	args := m.Called(ctx, id, coins)
	return args.Error(0)
}

func TestAuthUsecase_Register(t *testing.T) {
	repo := &mockUserRepository{}
	uc := NewAuthUsecase(repo, "secret")

	repo.On("Create", mock.Anything, mock.MatchedBy(func(u *entity.User) bool {
		return u.Username == "timur" && u.Coins == 1000
	})).Return(nil).Run(func(args mock.Arguments) {
		u := args.Get(1).(*entity.User)
		u.ID = 1
	})

	got, err := uc.Register(context.Background(), "timur", "pass")
	assert.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	repo.AssertExpectations(t)
}

func TestAuthUsecase_Login_Success(t *testing.T) {
	repo := &mockUserRepository{}
	uc := NewAuthUsecase(repo, "secret")

	hash, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := &entity.User{ID: 1, Username: "timur", PasswordHash: string(hash)}

	repo.On("GetByUsername", mock.Anything, "timur").Return(user, nil)

	token, err := uc.Login(context.Background(), "timur", "pass")
	assert.NoError(t, err)
	assert.Equal(t, "jwt-token", token)
}

func TestAuthUsecase_Login_InvalidPassword(t *testing.T) {
	repo := &mockUserRepository{}
	uc := NewAuthUsecase(repo, "secret")

	hash, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := &entity.User{PasswordHash: string(hash)}

	repo.On("GetByUsername", mock.Anything, "timur").Return(user, nil)

	_, err := uc.Login(context.Background(), "timur", "wrong")
	assert.ErrorIs(t, err, errors.ErrInvalidCredentials)
}
