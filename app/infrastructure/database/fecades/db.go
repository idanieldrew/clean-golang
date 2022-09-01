package fecades

import (
	"clean-golang/app/infrastructure/database/mysql"
	"database/sql"
	"os"
)

func DbFacade(connection string) (*sql.DB, error) {
	if connection == "mysql" {
		conn := mysql.NewConnection(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
		c, err := conn.Connect()
		if err != nil {
			return nil, err
		}
		return c, nil
	} else if connection == "pgsql" {
		return nil, nil
	} else {
		return nil, nil
	}
}
