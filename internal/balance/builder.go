package balance

import "github.com/go-money-transfer/internal/database"

func BuildBalanceService() *BalanceService {
	txManager := database.GetTxManager()
	balanceRepository := &BalanceRepository{DB: database.GetConnection()}

	return &BalanceService{txManager: txManager, balanceRepository: balanceRepository}
}
