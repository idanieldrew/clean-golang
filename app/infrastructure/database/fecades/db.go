package fecades

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/pgsql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"fmt"
	"os"
)

func NewDb(s string) error {
	switch s {
	case "mysql":
		m := mysql.NewMysql(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
		_, mErr := m.Make()
		if mErr != nil {
			return mErr
		}
	case "pgsql":
		p := pgsql.NewPgsql()
		_, pErr := p.Make()
		if pErr != nil {
			return pErr
		}
	case "mongo":
		m := mongo.NewMongo()
		_, mErr := m.Make()
		if mErr != nil {
			return mErr
		}
	default:
		textErr := fmt.Sprintf("%s is not database", s)
		logger.Error(textErr)
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
