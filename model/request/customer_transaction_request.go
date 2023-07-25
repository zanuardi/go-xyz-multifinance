package request

type CustomerTransactionRequest struct {
	CustomerId        string  `json:"customer_id"`
	ContractNumber    string  `json:"contract_number"`
	OTRPrice          float32 `json:"otr_price"`
	AdminFee          float32 `json:"admin_fee"`
	InstallmentAmount float32 `json:"installment_amount"`
	InterestAmount    int     `json:"interest_amount"`
	AssetName         string  `json:"asset_name"`
	Status            string  `json:"status"`
}
