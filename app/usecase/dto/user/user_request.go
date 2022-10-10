package user

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,gte=8"`
	Password string `validate:"required"`
	RoleId   string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()
}
