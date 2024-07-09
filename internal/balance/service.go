package balance

import (
	"context"
	"fmt"
	"time"

	"github.com/go-money-transfer/internal/database"
	"github.com/go-money-transfer/internal/errors"
)

type BalanceService struct {
	txManager         database.TxManager
	balanceRepository Repository
}

func (s *BalanceService) Transfer(transfer TransferRequest) error {
	beneficiaryBalance, err := s.balanceRepository.FindById(transfer.BeneficiaryID)
	if err != nil {
		return err
	}

	debtorBalance, err := s.balanceRepository.FindById(transfer.DebtorID)
	fmt.Println(transfer.DebtorID)
	if err != nil {
		return err
	}
	if debtorBalance.Amount-transfer.Amount < 0 {
		return errors.ErrInsuficientBalance
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	tx, err := s.txManager.InitTx(ctx)

	beneficiaryBalance.Amount += transfer.Amount
	err = s.balanceRepository.Update(beneficiaryBalance, tx)
	if err != nil {
		s.txManager.Rollback()
		return err
	}
	debtorBalance.Amount -= transfer.Amount
	err = s.balanceRepository.Update(debtorBalance, tx)
	if err != nil {
		s.txManager.Rollback()
		return err
	}

	return s.txManager.Commit()

}
