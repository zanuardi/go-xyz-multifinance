package request

import "time"

type CustomerRequest struct {
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      int       `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
}
