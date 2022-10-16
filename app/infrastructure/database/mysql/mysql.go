package mysql

import (
	"clean-golang/app/infrastructure/database/factories"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type (
	DbMysql struct {
		factories.Database
	}
	Mysql struct {
		Db *sql.DB
	}
)

var (
	Db *sql.DB
)

func NewMysql(User, Psd, Host, Port, Db string) factories.IDatabase {
	return &DbMysql{factories.Database{
		User: User,
		Psd:  Psd,
		Host: Host,
		Port: Port,
		Db:   Db,
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
