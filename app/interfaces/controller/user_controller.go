package controller

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/logger"
	"database/sql"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := mysql.Db.Query("SELECT name FROM users")
	if err != nil {
		logger.Error("problem in prepare")
		return
	}

	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {
			logger.Error("close")
		}
	}(res)

	for res.Next() {
		var user entities.User
		err = res.Scan(&user.Name)
		if err != nil {
			logger.Error("problem in scan")

			return
		}
		fmt.Print(user)
	}
}
