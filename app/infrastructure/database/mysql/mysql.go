package mysql

import (
	"clean-golang/app/interfaces/db"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	user, psd, host, port, db string
}

func NewConnection(user, psd, host, port, db string) db.Connection {
	return &Connection{
		user: user,
		psd:  psd,
		host: host,
		port: port,
		db:   db,
	}
}

func (c *Connection) Connect() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.user, c.psd, c.host, c.port, c.db)
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
