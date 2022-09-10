package repository

type UserRepository interface {
	All() (interface{}, error)
}
