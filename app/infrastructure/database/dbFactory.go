package database

import (
	"clean-golang/app/infrastructure/database/contracts/database"
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"fmt"
)

func NewDb() error {
	/* factory method pattern */

	// database.Connection[]{mysql.NewMysql(),mongo.NewMongo(),pgsql}
	err := database.Make(mysql.NewMysql())
	if err != nil {
		return err
	}

	return nil
}

func NewCache(s string) error {
	switch s {
	case "redis":
		r := redis.NewRedis()
		_, err := r.Make()
		if err != nil {
			return err
		}
	default:
		textErr := fmt.Sprintf("%s is not database cache", s)
		logger.Error(textErr)
	}
	return nil
}
