package mysql

import (
	db "clean-golang/app/infrastructure/database/contracts/database"
	factory "clean-golang/app/infrastructure/database/factories"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type (
	DbMysql struct {
		factory.Database
	}
	Mysql struct {
		Db *sql.DB
	}
)

var (
	Db *sql.DB
)

func NewMysql() db.Connection {
	return &DbMysql{factory.Database{
		User: os.Getenv("DB_USERNAME"),
		Psd:  os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Db:   os.Getenv("DB_DATABASE"),
	}}
}

func (m *DbMysql) Make() (interface{}, error) {
	connect, err := m.Connect()
	if err != nil {
		return nil, err
	}

	return connect, nil
}

func (m *DbMysql) Connect() (interface{}, error) {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.User, m.Psd, m.Host, m.Port, m.Db)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	return &Mysql{Db: Db}, nil
}
