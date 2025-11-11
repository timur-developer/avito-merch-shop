package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemRepository_GetByName(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{
		pool: mockPool,
		sb:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	repo := NewItemRepository(store)

	rows := mockPool.NewRows([]string{"id", "name", "price"}).AddRow(1, "t-shirt", 80)
	mockPool.ExpectQuery(`SELECT id, name, price FROM items WHERE name = \$1`).
		WithArgs("t-shirt").
		WillReturnRows(rows)

	item, err := repo.GetByName(context.Background(), "t-shirt")
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "t-shirt", item.Name)
	assert.Equal(t, 80, item.Price)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}

func TestItemRepository_GetAll(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewItemRepository(store)

	rows := mockPool.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "t-shirt", 80).
		AddRow(2, "cup", 20)
	mockPool.ExpectQuery(`SELECT id, name, price FROM items`).
		WillReturnRows(rows)

	items, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, "cup", items[1].Name)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}
