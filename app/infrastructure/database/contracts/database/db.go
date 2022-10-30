package database

type Connection interface {
	Connect() (interface{}, error)
	Make() (interface{}, error)
}

type (
	Database struct {
		User, Psd, Host, Port, Db string
	}
)
