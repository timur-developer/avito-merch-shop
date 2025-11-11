package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionRepository_Create(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewTransactionRepository(store)

	tx := &entity.Transaction{
		FromUserID:      1,
		ToUserID:        2,
		Amount:          150,
		TransactionType: "purchase",
	}

	mockPool.ExpectExec(`INSERT INTO transactions.*`).
		WithArgs(1, 2, 150, "purchase").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := repo.Create(context.Background(), tx)
	assert.NoError(t, err)
	assert.NoError(t, mockPool.ExpectationsWereMet())
}

func TestTransactionRepository_GetUserHistory(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewTransactionRepository(store)

	now := time.Now().Truncate(time.Second)
	rows := mockPool.NewRows([]string{"id", "from_user_id", "to_user_id", "amount", "transaction_type", "created_at"}).
		AddRow(1, 1, 2, 100, "transfer", now)

	mockPool.ExpectQuery(`SELECT.*FROM transactions WHERE.*ORDER BY created_at DESC`).
		WithArgs(1, 1).
		WillReturnRows(rows)

	history, err := repo.GetUserHistory(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, history, 1)
	assert.Equal(t, "transfer", history[0].TransactionType)
	assert.Equal(t, now, history[0].CreatedAt.Truncate(time.Second))

	assert.NoError(t, mockPool.ExpectationsWereMet())
}
