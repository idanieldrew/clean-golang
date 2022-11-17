package database

import (
	"clean-golang/app/infrastructure/logger"
)

type Connection interface {
	Connect() (interface{}, error)
	Make() (interface{}, error)
}

type (
	Database struct {
		User, Psd, Host, Port, Db string
	}
)

func Make(c Connection) error {
	_, err := c.Make()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
