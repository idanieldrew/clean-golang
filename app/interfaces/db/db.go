package db

type Connection interface {
	Connect() (interface{}, error)
}
