package pgsql

import (
	db "clean-golang/app/infrastructure/database/contracts/database"
	factory "clean-golang/app/infrastructure/database/factories"
	"os"
)

type (
	Pgsql struct {
		factory.Database
	}
)

func NewPgsql() db.Connection {
	return &Pgsql{factory.Database{
		User: os.Getenv("DB_USERNAME"),
		Psd:  os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Db:   os.Getenv("DB_DATABASE"),
	}}
}

func (p *Pgsql) Make() (interface{}, error) {
	return nil, nil
}

func (p *Pgsql) Connect() (interface{}, error) {
	return nil, nil
}
