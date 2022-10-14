package redis

import (
	"clean-golang/app/infrastructure/database/factories"
	"github.com/go-redis/redis/v8"
	"os"
)

type (
	DbRedis struct {
		Info factories.Database
	}

	Redis struct {
		Db *redis.Client
	}
)

var (
	Rdb *redis.Client
)

func NewRedis() factories.IDatabase {
	return &DbRedis{Info: factories.Database{
		Psd:  os.Getenv("CACHE_PASSWORD"),
		Host: os.Getenv("CACHE_HOST"),
	}}
}

func (r *DbRedis) Make() (interface{}, error) {
	connect, err := r.Connect()
	if err != nil {
		return nil, err
	}
	return connect, nil
}

func (r *DbRedis) Connect() (interface{}, error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     r.Info.Host,
		Password: r.Info.Psd,
		DB:       0,
	})
	/*if err := Rdb.Ping(ctx); err != nil {
		return nil, err.Err()
	}*/
	return Rdb, nil
}