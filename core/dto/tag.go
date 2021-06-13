package dto

type UserDto struct {
	NickName string `json:"nick_name" validate:"required,min=13,max=32"`
	IsActive bool   `json:"is_active" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
}
