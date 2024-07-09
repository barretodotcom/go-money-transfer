package config

import "errors"

var ErrDatabaseURLNotFound = errors.New("DATABASE_URL variable is not set.")
