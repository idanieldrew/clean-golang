package user

import (
	"clean-golang/app/infrastructure/logger"
	user_request "clean-golang/app/interfaces/request/user"
	"clean-golang/app/usecase/dto/user"
	user2 "clean-golang/app/usecase/repository/user"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

type UserInteract struct {
	UserRepository user2.UserRepository
}

func (u *UserInteract) Index() ([]user.PublicResponse, int) {
	// Repository
	users, err := u.UserRepository.All()
	if err != nil {
		logger.Error("problem")
		return nil, http.StatusInternalServerError
	}

	r := user.UserResponse{}
	res := r.Public(users)
	return res, http.StatusOK
}

func (u *UserInteract) Register(req *user_request.Request) (int, string) {
	// check validation usecase
	request := &user.Request{
		Email: req.Email,
		Phone: req.Phone,
	}
	if req := request.Validation(u.UserRepository); !req {
		return http.StatusUnprocessableEntity, "incorrect mail or phone"
	}

	// register user
	err := u.UserRepository.Register(req)
	if err != nil {
		return http.StatusInternalServerError, "server problem"
	}
	go sendMail([]string{req.Email})

	return http.StatusCreated, "success"
}

func sendMail(to []string) {
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")
	from := "cleangolang@gmail.com"
	message := []byte(fmt.Sprintf(`To: %s 

From: %s 

Subject: %s`,
		to[0], "cleangolang.gmail.com", "Why aren’t you using Mailtrap yet?\n\n\t\tHere’s the space for your great sales pitch"))
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(host+":25", auth, from, to, message)
	if err != nil {
		logger.Error(err.Error())
	}
}
