package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
)

type itemRepository struct {
	store *Store
}

func NewItemRepository(store *Store) repository.ItemRepository {
	return &itemRepository{store: store}
}

func (r *itemRepository) GetByName(ctx context.Context, name string) (*entity.Item, error) {
	query := r.store.Builder().
		Select("id, name, price").
		From("items").
		Where("name = ?", name)

	sql, args, _ := query.ToSql()

	var item entity.Item
	err := r.store.pool.QueryRow(ctx, sql, args...).Scan(
		&item.ID, &item.Name, &item.Price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetAll(ctx context.Context) ([]entity.Item, error) {
	query := r.store.Builder().
		Select("id, name, price").
		From("items")

	sql, args, _ := query.ToSql()

	rows, err := r.store.Pool().Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.Item
	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
