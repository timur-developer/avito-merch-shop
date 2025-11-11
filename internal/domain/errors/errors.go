package errors

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInsufficientFunds   = errors.New("insufficient funds")
	ErrSelfTransfer        = errors.New("cant transfer yourself")
	ErrItemNotFound        = errors.New("item not found")
	ErrNegativeAmount      = errors.New("amount must > 0")
	ErrInvalidToken        = errors.New("invalid token")
	ErrInternalServerError = errors.New("internal server err")
)
