package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type txManager struct {
	tx  pgx.Tx
	ctx context.Context
	DB  *pgxpool.Pool
}

type TxManager interface {
	InitTx(ctx context.Context) (pgx.Tx, error)
	Rollback() error
	Commit() error
	GetTransaction() pgx.Tx
}

const transactionTimeout = 20 * time.Second

func GetTxManager() TxManager {
	return &txManager{DB: conn}
}

func (m *txManager) InitTx(ctx context.Context) (pgx.Tx, error) {
	var err error
	m.tx, err = m.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	m.ctx = ctx
	return m.tx, nil
}

func (m *txManager) Rollback() error {
	if m.tx == nil {
		log.Fatal("transaction is not initialized")
	}
	err := m.tx.Rollback(m.ctx)
	if err != nil {
		return err
	}
	m.tx = nil
	return nil
}

func (m *txManager) Commit() error {
	if m.tx == nil {
		log.Fatal("transaction is not initialized")
	}
	if m.tx.Conn().IsClosed() {
		log.Fatal("transaction is already closed")
	}
	return m.tx.Commit(m.ctx)
}

func (m *txManager) GetTransaction() pgx.Tx {
	return m.tx
}
