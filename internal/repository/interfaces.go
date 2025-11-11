package repository

import (
	"avito-merch-shop/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
	UpdateCoins(ctx context.Context, id int, coins int) error
}

type ItemRepository interface {
	GetByName(ctx context.Context, name string) (*entity.Item, error)
	GetAll(ctx context.Context) ([]entity.Item, error)
}

type InventoryRepository interface {
	AddItem(ctx context.Context, inv *entity.InventoryItem) error
	GetUserInventory(ctx context.Context, userID int) ([]entity.InventoryItem, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, tx *entity.Transaction) error
	GetUserHistory(ctx context.Context, userID int) ([]entity.Transaction, error)
}
