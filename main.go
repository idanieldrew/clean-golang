package main

import (
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/logger"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if initErr := recover(); initErr != nil {
			logger.Error(string(debug.Stack()))
		}
		os.Exit(1)
	}()

	conn := mysql.NewConnection()
	db, err := conn.Connect(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		logger.Error(err.Error())
	}
	_ = db
}
