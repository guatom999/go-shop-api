package model

type (
	ItemCreatingReq struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=64"`
		Description string `json:"description" validate:"required,max=128"`
		Picture     string `json:"picture" validate:"required"`
		Price       uint64 `json:"price" validate:"required,max=64"`
	}

	ItemEditingReq struct {
		AdminID     string
		Name        string `json:"name" validate:"omitempty,max=64"`
		Description string `json:"description" validate:"omitempty,max=128"`
		Picture     string `json:"picture" validate:"omitempty"`
		Price       uint64 `json:"price" validate:"omitempty,max=64"`
	}
)
