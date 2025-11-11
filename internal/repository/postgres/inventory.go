package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/repository"
	"context"
)

type inventoryRepository struct {
	store *Store
}

func NewInventoryRepository(store *Store) repository.InventoryRepository {
	return &inventoryRepository{store: store}
}

func (r *inventoryRepository) AddItem(ctx context.Context, inv *entity.InventoryItem) error {
	query := r.store.Builder().
		Insert("user_inventory").
		Columns("user_id", "item_id", "quantity").
		Values(inv.UserID, inv.ItemID, inv.Quantity).
		Suffix("ON CONFLICT (user_id, item_id) DO UPDATE SET quantity = user_inventory.quantity + EXCLUDED.quantity")

	sql, args, _ := query.ToSql()
	_, err := r.store.Pool().Exec(ctx, sql, args...)
	return err
}

func (r *inventoryRepository) GetUserInventory(ctx context.Context, userID int) ([]entity.InventoryItem, error) {
	query := r.store.Builder().
		Select("ui.user_id, ui.item_id, it.name, ui.quantity").
		From("user_inventory ui").
		Join("items it ON ui.item_id = it.id").
		Where("ui.user_id = ?", userID)

	sql, args, _ := query.ToSql()

	rows, err := r.store.Pool().Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.InventoryItem
	for rows.Next() {
		var item entity.InventoryItem
		var itemName string
		if err := rows.Scan(&item.UserID, &item.ItemID, &itemName, &item.Quantity); err != nil {
			return nil, err
		}
		item.ItemName = itemName
		items = append(items, item)
	}
	return items, rows.Err()
}
