package factories

type (
	IDatabase interface {
		Make() (interface{}, error)
	}

	Database struct {
		User, Psd, Host, Port, Db string
	}
)
