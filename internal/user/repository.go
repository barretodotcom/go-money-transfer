package user

import (
	"context"

	"github.com/go-money-transfer/internal/database"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	FindUserByUsername(username string) (*User, error)
	Create(user User, tx pgx.Tx) error
}

type UserRepository struct {
	DB *pgxpool.Pool
}

var Repo *UserRepository

func InitRepository() {
	conn := database.GetConnection()

	Repo = &UserRepository{DB: conn}
}

func (r *UserRepository) FindUserByUsername(username string) (*User, error) {
	row := r.DB.QueryRow(context.TODO(), "SELECT id, password FROM users WHERE username = $1", username)

	var id, password string
	err := row.Scan(&id, &password)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	if id == "" {
		return nil, nil
	}

	return &User{ID: id, Username: username, Password: password}, nil
}

func (r *UserRepository) Create(user User, tx pgx.Tx) error {
	sql := "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)"
	_, err := tx.Exec(context.TODO(), sql, user.ID, user.Username, user.Password)

	return err
}
