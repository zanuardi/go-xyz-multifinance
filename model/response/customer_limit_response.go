package response

import "time"

type CustomerLimitResponse struct {
	Id             int       `json:"id"`
	CustomerId     int       `json:"customer_id"`
	Limit1         float32   `json:"limit_1"`
	Limit2         float32   `json:"limit_2"`
	Limit3         float32   `json:"limit_3"`
	Limit4         float32   `json:"limit_4"`
	RemainingLimit float32   `json:"remaining_limit"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
