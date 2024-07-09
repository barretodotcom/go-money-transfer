package errors

import "errors"

var ErrUserAlreadyExists = errors.New("User already exists.")
var ErrInvalidUsernameOrPassword = errors.New("Invalid username or password.")
var ErrTransactionAlreadyClosed = errors.New("database transaction is already closed.")
var ErrInsuficientBalance = errors.New("Insuficient balance.")
