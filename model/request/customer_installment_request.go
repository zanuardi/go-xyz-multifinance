package request

type CustomerInstallmentRequest struct {
	CustomerTransactionId int     `json:"customer_transaction_id"`
	CustomerLimitId       int     `json:"customer_limit_id"`
	Tenor                 int     `json:"tenor"`
	TotalAmounts          float32 `json:"total_amounts"`
	RemainingAmounts      float32 `json:"remaining_amounts"`
	RemainingLimit        int     `json:"remaining_limit"`
}
