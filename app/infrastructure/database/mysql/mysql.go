package mysql

import (
	"clean-golang/app/infrastructure/database/factories"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type (
	DbMysql struct {
		factories.Database
	}
)

func NewMysql() factories.IDatabase {
	return &DbMysql{factories.Database{
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
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.User, m.Psd, m.Host, m.Port, m.Db)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
