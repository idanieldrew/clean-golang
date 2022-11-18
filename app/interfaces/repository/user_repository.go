package repository

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/logger"
	user_request "clean-golang/app/interfaces/request/user"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type UserRepository struct {
	Connection *sql.DB
	Cache      *redis.Client
}

const (
	all         = "SELECT id,name,email,phone,created_at,updated_at FROM users"
	register    = "INSERT INTO users (name,phone,email,password) VALUES (?,?,?,?)"
	email_count = "SELECT COUNT(email) AS EmailCount FROM users WHERE email =?"
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
	ctx := context.Background()

	res, existErr := r.Cache.Get(ctx, "users").Result()

	if existErr == redis.Nil {
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		//fmt.Println("from db")
		res, err := r.Connection.QueryContext(ctx, all)
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
		b, _ := json.Marshal(users)

		existErr = r.Cache.SetNX(ctx, "users", b, 15*time.Minute).Err()
		return users, nil
	} else {
		//fmt.Println("from cache")
		uErr := json.Unmarshal([]byte(res), &users)
		if uErr != nil {
			return nil, uErr
		}
		return users, nil
	}
}

func (r *UserRepository) Register(req *user_request.Request) error {
	stmt, pErr := r.Connection.Prepare(register)
	if pErr != nil {
		logger.Error(pErr.Error())
		return pErr
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(req.Name, req.Phone, req.Email, req.Password)
	if execErr != nil {
		logger.Error(execErr.Error())

		return execErr
	}

	return nil
}

func (r *UserRepository) CountMail(email string) int {
	var count int
	err := r.Connection.QueryRow(email_count, email).Scan(&count)
	switch {
	case err != nil:
		logger.Error(err.Error())
		return count
	default:
		return count
	}
}
