package repository

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/logger"
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
	All = "SELECT id,name,email,phone,created_at,updated_at FROM users"
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
		//fmt.Println("from db")
		res, err := r.Connection.Query(All)
		if err != nil {
			textErr := fmt.Sprintf("problem in %s query", All)
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
