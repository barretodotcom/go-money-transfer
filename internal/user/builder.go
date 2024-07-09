package user

import (
	"github.com/go-money-transfer/internal/balance"
	"github.com/go-money-transfer/internal/database"
)

func BuildUserService() *UserService {
	txManager := database.GetTxManager()
	userRepository := UserRepository{DB: database.GetConnection()}
	balanceRepository := balance.BalanceRepository{DB: database.GetConnection()}
	service := &UserService{UserRepository: &userRepository, BalanceRepository: &balanceRepository, TxManager: txManager}

	return service
}
