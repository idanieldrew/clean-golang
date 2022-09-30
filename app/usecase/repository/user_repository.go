package repository

import "clean-golang/app/entities"

type UserRepository interface {
	All() (entities.Users, error)
}
