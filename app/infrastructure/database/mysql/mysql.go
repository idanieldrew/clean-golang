package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
}

func NewConnection() *Connection {
	return new(Connection)
}

func (c *Connection) Connect(user, psd, host, port, db string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%i)/%s?parseTime=true", user, psd, host, port, db)
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
