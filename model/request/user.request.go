package request

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" validate:"required"`
}

type UserUpdatePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}
