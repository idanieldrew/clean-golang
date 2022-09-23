package interactor

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/logger"
	"database/sql"
	_ "database/sql"
	"log"
)

func c(res *sql.Rows) {
	func(res *sql.Rows) {
		err := res.Close()
		if err != nil {
			logger.Error("c")
		}
	}(res)
}

func Index() entities.Users {
	var user entities.User
	var users entities.Users

	res, err := mysql.Db.Query("SELECT name FROM users")
	if err != nil {
		logger.Error("problem in prepare")
		return nil
	}

	// close query
	defer c(res)

	for res.Next() {
		err = res.Scan(&user.Name)
		if err != nil {
			log.Fatalln(22)
			logger.Error("problem in scan")
			return nil
		}
		users = append(users, user)
	}
	return users
}
