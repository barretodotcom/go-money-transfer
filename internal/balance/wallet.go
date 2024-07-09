package balance

import (
	"math/big"
	"time"
)

type Wallet struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`
	Amount    big.Int   `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}
