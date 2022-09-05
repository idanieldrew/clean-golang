package mysql

import (
	"clean-golang/app/infrastructure/database/factories"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type (
	DbMysql struct {
		factories.Database
	}
)

func (m *DbMysql) Make() (interface{}, error) {
	connect, err := m.Connect()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return connect, nil
}

func NewMysql() factories.IDatabase {
	return &DbMysql{factories.Database{
		User: os.Getenv("DB_USERNAME"),
		Psd:  os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Db:   os.Getenv("DB_DATABASE"),
	}}
}

func (m *DbMysql) Connect() (*sql.DB, error) {
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