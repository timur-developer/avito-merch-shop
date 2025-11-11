package postgres

import (
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	store *Store
}

func NewUserRepository(store *Store) repository.UserRepository {
	return &userRepository{store: store}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := r.store.Builder().
		Insert("users").
		Columns("username", "password_hash", "coins").
		Values(user.Username, user.PasswordHash, user.Coins).
		Suffix("RETURNING id")

	sql, args, _ := query.ToSql()
	return r.store.Pool().QueryRow(ctx, sql, args...).Scan(&user.ID)
}
func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := r.store.Builder().
		Select("id, username, password_hash, coins, created_at").
		From("users").
		Where("username = ?", username)

	sql, args, _ := query.ToSql()
	var u entity.User
	err := r.store.Pool().QueryRow(ctx, sql, args...).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Coins, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	query := r.store.Builder().
		Select("id, username, password_hash, coins, created_at").
		From("users").
		Where("id = ?", id)

	sql, args, _ := query.ToSql()
	var u entity.User
	err := r.store.Pool().QueryRow(ctx, sql, args...).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Coins, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) UpdateCoins(ctx context.Context, id int, coins int) error {
	query := r.store.Builder().
		Update("users").
		Set("coins", coins).
		Where("id = ?", id)

	sql, args, _ := query.ToSql()
	_, err := r.store.Pool().Exec(ctx, sql, args...)
	return err
}
