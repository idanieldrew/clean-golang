package database

type Connection interface {
	Connect() (interface{}, error)
	Make() (interface{}, error)
}
