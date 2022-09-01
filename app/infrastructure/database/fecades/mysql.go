package fecades

import (
	"clean-golang/app/infrastructure/database/mysql"
	"database/sql"
	"os"
)

func NewMysqlFacade() (*sql.DB, error) {
	conn := mysql.NewConnection(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))

	c, err := conn.Connect()
	if err != nil {
		return nil, err
	}

	return c, nil
}
