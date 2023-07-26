package domain

import "time"

type Customer struct {
	Id          int       `json:"id"`
	NIK         string    `json:"nik" validate:"required"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      int       `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
