package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/repository"
	"context"
)

type transactionRepository struct {
	store *Store
}

func NewTransactionRepository(store *Store) repository.TransactionRepository {
	return &transactionRepository{store: store}
}

func (r *transactionRepository) Create(ctx context.Context, tx *entity.Transaction) error {
	query := r.store.Builder().
		Insert("transactions").
		Columns("from_user_id", "to_user_id", "amount", "transaction_type").
		Values(tx.FromUserID, tx.ToUserID, tx.Amount, tx.TransactionType)

	sql, args, _ := query.ToSql()
	_, err := r.store.Pool().Exec(ctx, sql, args...)
	return err
}

func (r *transactionRepository) GetUserHistory(ctx context.Context, userID int) ([]entity.Transaction, error) {
	query := r.store.Builder().
		Select("id, from_user_id, to_user_id, amount, transaction_type, created_at").
		From("transactions").
		Where("from_user_id = ? OR to_user_id = ?", userID, userID).
		OrderBy("created_at DESC")

	sql, args, _ := query.ToSql()

	rows, err := r.store.Pool().Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []entity.Transaction
	for rows.Next() {
		var t entity.Transaction
		if err := rows.Scan(&t.ID, &t.FromUserID, &t.ToUserID, &t.Amount, &t.TransactionType, &t.CreatedAt); err != nil {
			return nil, err
		}
		history = append(history, t)
	}
	return history, rows.Err()
}
