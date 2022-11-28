package user

import (
	"clean-golang/app/infrastructure/logger"
	"errors"
	"regexp"
)

type Request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (u *Request) Validation() error {
	mail := checkMail(u.Email)
	if !mail {
		logger.Error("email is incorrect")
		return errors.New("email is incorrect")
	}
	name := checkSize(u.Name, 3, 12)
	if !name {
		logger.Error("name is incorrect")
		return errors.New("name is incorrect")
	}
	return nil
}

func checkMail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func checkSize(f string, l, h int) bool {
	i := len(f)
	if i < l || i > h {
		return false
	}
	return true
}
