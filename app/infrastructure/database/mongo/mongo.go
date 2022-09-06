package mongo

import (
	"clean-golang/app/infrastructure/database/factories"
	"clean-golang/app/infrastructure/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type (
	DbMongo struct {
		factories.Database
	}
)

func NewMongo() factories.IDatabase {
	return &DbMongo{factories.Database{
		User: os.Getenv("DB_USERNAME"),
		Psd:  os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Db:   os.Getenv("DB_DATABASE"),
	}}
}

func (m *DbMongo) Make() (interface{}, error) {
	connect, err := m.Connect()
	if err != nil {
		return nil, err
	}

	return connect, nil
}

func (m *DbMongo) Connect() (interface{}, error) {
	client, ctx, cancelFunc, err := connection()
	if err != nil {
		logger.Error("problem in connect to mongo")
		return nil, err
	}
	if pErr := client.Ping(ctx, readpref.Primary()); pErr != nil {
		logger.Error("problem in ping mongo")
		return nil, pErr
	}
	_ = cancelFunc

	return client, nil
}

func connection() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

	dm := new(DbMongo)
	auth := options.Credential{
		Username: dm.User,
		Password: dm.Psd,
	}

	client, err := mongo.Connect(ctx, options.Client().SetAuth(auth))
	return client, ctx, cancelFunc, err
}
