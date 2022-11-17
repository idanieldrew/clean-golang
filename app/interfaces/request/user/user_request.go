package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Name     string `json:"name" validate:"gte=3,lte=12"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,gte=8"`
	Password string `json:"password" validate:"required"`
}

var v = validator.New()

func (u *Request) Validation() error {
	err := v.Struct(u)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Println(v.Namespace())
			fmt.Println(v.Tag())
			fmt.Println()
		}
		return err
	}
	return nil
}
