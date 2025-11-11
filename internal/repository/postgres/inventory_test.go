package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryRepository_AddItem(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewInventoryRepository(store)

	inv := &entity.InventoryItem{UserID: 1, ItemID: 3, Quantity: 2}

	mockPool.ExpectExec(`INSERT INTO user_inventory.*ON CONFLICT.*DO UPDATE`).
		WithArgs(1, 3, 2).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := repo.AddItem(context.Background(), inv)
	assert.NoError(t, err)
	assert.NoError(t, mockPool.ExpectationsWereMet())
}

func TestInventoryRepository_GetUserInventory(t *testing.T) {
	mockPool, _ := pgxmock.NewPool()
	defer mockPool.Close()

	store := &Store{pool: mockPool, sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
	repo := NewInventoryRepository(store)

	rows := mockPool.NewRows([]string{"user_id", "item_id", "name", "quantity"}).
		AddRow(1, 1, "t-shirt", 5)

	mockPool.ExpectQuery(`SELECT ui\.user_id, ui\.item_id, it\.name, ui\.quantity FROM user_inventory ui JOIN items it`).
		WithArgs(1).
		WillReturnRows(rows)

	items, err := repo.GetUserInventory(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "t-shirt", items[0].ItemName)
	assert.Equal(t, 5, items[0].Quantity)

	assert.NoError(t, mockPool.ExpectationsWereMet())
}
