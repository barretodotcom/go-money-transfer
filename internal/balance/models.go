package balance

import "time"

type Balance struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`
	Amount    int       `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TransferRequest struct {
	DebtorID      string `json:"debtorId"`
	BeneficiaryID string `json:"beneficiaryId"`
	Amount        int    `json:"amount"`
}
