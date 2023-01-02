package v1

type CreateUserRequest struct {
	Name       string `json:"name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Repassword string `json:"repassword" validate:"required"`
}
