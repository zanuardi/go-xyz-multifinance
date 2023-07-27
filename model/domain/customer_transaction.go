package domain

import "time"

type CustomerTransaction struct {
	Id                int       `json:"id"`
	CustomerId        int       `json:"customer_id"`
	ContractNumber    string    `json:"contract_number"`
	OTRPrice          float32   `json:"otr_price"`
	AdminFee          float32   `json:"admin_fee"`
	InstallmentAmount float32   `json:"installment_amount"`
	InterestAmount    int       `json:"interest_amount"`
	AssetName         string    `json:"asset_name"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}
