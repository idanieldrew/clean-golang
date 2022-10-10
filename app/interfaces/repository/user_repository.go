package repository

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/logger"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	Connection *sql.DB
}

const (
	all = "SELECT id,name,email,phone,created_at,updated_at FROM users"
)

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
	user := entities.User{}
	users := entities.Users{}

	res, err := r.Connection.Query(all)
	if err != nil {
		textErr := fmt.Sprintf("problem in %s query", all)
		logger.Error(textErr)
		return nil, err
	}

	// close query
	defer c(res)

	for res.Next() {
		err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			logger.Error("problem in scan")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
