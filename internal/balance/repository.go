package balance

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Update(balance Balance, tx pgx.Tx) error
	Create(balance Balance, tx pgx.Tx) error
	FindById(id string) (Balance, error)
	FindByUserId(userId string) (Balance, error)
}

type BalanceRepository struct {
	DB *pgxpool.Pool
}

func (r *BalanceRepository) Create(balance Balance, tx pgx.Tx) error {
	sql := "INSERT INTO balances (id, user_id, amount, updated_at) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(context.TODO(), sql, balance.ID, balance.UserId, balance.Amount, balance.UpdatedAt)
	return err
}

func (r *BalanceRepository) Update(balance Balance, tx pgx.Tx) error {
	sql := "UPDATE balances SET amount = $1, updated_at = $2 WHERE id = $3"
	_, err := tx.Exec(context.TODO(), sql, balance.Amount, time.Now(), balance.ID)
	return err
}

func (r *BalanceRepository) FindById(id string) (Balance, error) {
	sql := "SELECT user_id, amount, updated_at FROM balances WHERE id = $1"

	row := r.DB.QueryRow(context.TODO(), sql, id)
	var userId string
	var amount int
	var updatedAt time.Time
	err := row.Scan(&userId, &amount, &updatedAt)
	if err != nil {
		return Balance{}, err
	}

	return Balance{ID: id, UserId: userId, Amount: amount, UpdatedAt: updatedAt}, err
}

func (r *BalanceRepository) FindByUserId(userId string) (Balance, error) {
	sql := "SELECT id, amount, updated_at FROM balances WHERE user_id = $1"

	row := r.DB.QueryRow(context.TODO(), sql, userId)
	var id string
	var amount int
	var updatedAt time.Time
	err := row.Scan(&id, &amount, &updatedAt)
	if err != nil {
		return Balance{}, err
	}

	return Balance{ID: id, UserId: userId, Amount: amount, UpdatedAt: updatedAt}, err
}
