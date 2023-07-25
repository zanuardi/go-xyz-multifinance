package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,max=200,min=2"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=200,min=2"`
}
