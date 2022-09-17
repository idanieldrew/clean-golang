package entities

import "time"

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	EmailVerifiedAt string `json:"emailVerifiedAt"`
	Password        string `json:"password"`
	//RoleId          string    `json:"role_Id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type Users []User
