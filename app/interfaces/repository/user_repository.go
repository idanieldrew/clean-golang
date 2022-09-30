package repository

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/logger"
	"database/sql"
)

type UserRepository struct {
	Connection *sql.DB
}

// close query
func c(res *sql.Rows) {
	func(res *sql.Rows) {
		err := res.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(res)
}

func (r *UserRepository) All() (entities.Users, error) {
	var user entities.User
	var users entities.Users

	res, err := r.Connection.Query("SELECT name FROM users")
	if err != nil {
		logger.Error("problem in prepare")
		return nil, err
	}

	// close query
	defer c(res)

	for res.Next() {
		err = res.Scan(&user.Name)
		if err != nil {
			logger.Error("problem in scan")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
