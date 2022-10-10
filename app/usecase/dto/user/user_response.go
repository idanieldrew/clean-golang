package user

import (
	"clean-golang/app/entities"
	"time"
)

type (
	UserResponse struct {
	}

	PublicResponse struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Password  string    `json:"password"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	PrivateResponse struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)

var public []PublicResponse
var private []PrivateResponse

func (u *UserResponse) Public(users entities.Users) []PublicResponse {
	for _, user := range users {
		res := PublicResponse{
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Password:  user.Password,
			UpdatedAt: user.UpdatedAt,
		}
		public = append(public, res)
	}

	return public
}

func (u *UserResponse) Private(users entities.Users) []PrivateResponse {
	for _, u := range users {
		res := PrivateResponse{
			Id:        u.Id,
			Name:      u.Name,
			Email:     u.Email,
			Phone:     u.Phone,
			Password:  u.Password,
			UpdatedAt: u.UpdatedAt,
		}
		private = append(private, res)
	}
	return private
}
