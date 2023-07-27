package domain

import "time"

type CustomerInstallment struct {
	Id                    int       `json:"id"`
	CustomerTransactionId int       `json:"customer_transaction_id"`
	CustomerLimitId       int       `json:"customer_limit_id"`
	Tenor                 int       `json:"tenor"`
	TotalAmounts          float32   `json:"total_amounts"`
	RemainingAmounts      float32   `json:"remaining_amounts"`
	RemainingLimit        int       `json:"remaining_limit"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
}
