package models

type Mail struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required,phone"`
	Message string `json:"message" validate:"required"`
}
