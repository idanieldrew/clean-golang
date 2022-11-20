package user

import (
	"clean-golang/app/usecase/repository"
)

type Request struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *Request) Validation(r repository.UserRepository) bool {
	m := r.CountMail(u.Email)
	p := r.CountPhone(u.Phone)
	if (m == 0) && (p == 0) {
		return true
	}
	return false
}
