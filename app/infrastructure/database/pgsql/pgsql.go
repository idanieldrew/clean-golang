package pgsql

import (
	"clean-golang/app/infrastructure/database/contracts/database"
	"os"
)

type (
	Pgsql struct {
		database.Database
	}
)

func NewPgsql() database.Connection {
	return &Pgsql{database.Database{
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
