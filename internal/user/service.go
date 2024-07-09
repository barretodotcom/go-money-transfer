package user

import (
	"context"
	"time"

	"github.com/go-money-transfer/internal/balance"
	"github.com/go-money-transfer/internal/database"
	"github.com/go-money-transfer/internal/errors"
	"github.com/go-money-transfer/pkg/hash"
	"github.com/google/uuid"
)

type UserService struct {
	TxManager         database.TxManager
	UserRepository    Repository
	BalanceRepository balance.Repository
}

func (s *UserService) CreateUser(createUser CreateUser) error {
	userExists, err := s.UserRepository.FindUserByUsername(createUser.Username)
	if err != nil {
		return err
	}

	if userExists != nil {
		return errors.ErrUserAlreadyExists
	}

	hashedPassword, err := hash.HashPassword(createUser.Password)
	if err != nil {
		return err
	}

	user := User{ID: uuid.New().String(), Username: createUser.Username, Password: hashedPassword}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	tx, err := s.TxManager.InitTx(ctx)
	if err != nil {
		return err
	}
	err = s.UserRepository.Create(user, tx)
	if err != nil {
		s.TxManager.Rollback()
		return err
	}

	balance := balance.Balance{
		ID:        uuid.NewString(),
		UserId:    user.ID,
		Amount:    1000,
		UpdatedAt: time.Now(),
	}
	err = s.BalanceRepository.Create(balance, tx)
	if err != nil {
		s.TxManager.Rollback()
		return err
	}

	return s.TxManager.Commit()
}
