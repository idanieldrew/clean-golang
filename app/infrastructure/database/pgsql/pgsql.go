package pgsql

import (
	"clean-golang/app/infrastructure/database/factories"
	"database/sql"
	"os"
)

type (
	Pgsql struct {
		factories.Database
	}
)

func NewPgsql() factories.IDatabase {
	return &Pgsql{factories.Database{
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

func (p *Pgsql) Connect() (*sql.DB, error) {
	return nil, nil
}
