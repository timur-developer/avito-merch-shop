package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	store := &Store{
		pool: mockPool,
		sb:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	repo := NewUserRepository(store)

	user := &entity.User{
		Username:     "timur",
		PasswordHash: "hash",
		Coins:        1000,
	}

	mockPool.ExpectQuery(`INSERT INTO users.*RETURNING id`).
		WithArgs("timur", "hash", 1000).
		WillReturnRows(mockPool.NewRows([]string{"id"}).AddRow(1))

	err = repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername_NotFound(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewUserRepository(store)

	mockPool.ExpectQuery(`SELECT.*FROM users WHERE username = \$1`).
		WithArgs("unknown").
		WillReturnError(pgx.ErrNoRows)

	got, err := repo.GetByUsername(context.Background(), "unknown")
	assert.NoError(t, err)
	assert.Nil(t, got)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}

func TestUserRepository_UpdateCoins(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewUserRepository(store)

	mockPool.ExpectExec(`UPDATE users SET coins = \$1 WHERE id = \$2`).
		WithArgs(500, 1).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := repo.UpdateCoins(context.Background(), 1, 500)
	assert.NoError(t, err)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}
