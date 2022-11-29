package helper

import (
	"clean-golang/app/infrastructure/logger"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func MakeHash(t string) []byte {
	b, err := bcrypt.GenerateFromPassword([]byte(t), 8)
	if err != nil {
		s := fmt.Sprintf("problem in bcrypt %s", t)
		logger.Error(s)
		return nil
	}
	return b
}
