package main

import (
	"avito-merch-shop/config"
	"avito-merch-shop/internal/domain/entity"
	"avito-merch-shop/internal/repository/postgres"
	"context"
	"fmt"
	"time"
)

func main() {
	// тест функци создании пользователя

	cfg := config.Load()

	store, err := postgres.NewStore(cfg)
	if err != nil {
		fmt.Println("fail to create store:", err)
	}
	ctx := context.Background()

	userRepo := postgres.NewUserRepository(store)

	userToCreate := &entity.User{
		ID:           2,
		Username:     "Timur1",
		PasswordHash: "sdfgdfg",
		Coins:        1000,
		CreatedAt:    time.Now(),
	}

	if err := userRepo.Create(ctx, userToCreate); err != nil {
		fmt.Println("fail to create user:", err)
	}

}
