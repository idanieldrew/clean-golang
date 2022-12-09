package mongo

import (
	"clean-golang/app/infrastructure/database/contracts/database"
	"clean-golang/app/infrastructure/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type (
	DbMongo struct {
		database.Database
	}
)

var (
	authentication *DbMongo
	Db             *mongo.Database
)

func NewMongo() database.Connection {
	authentication = &DbMongo{database.Database{
		User: os.Getenv("DB_USERNAME"),
		Psd:  os.Getenv("DB_PASSWORD"),
		Db:   os.Getenv("DB_DATABASE"),
	}}
	return authentication
}

func (m *DbMongo) Make() (interface{}, error) {
	connect, err := m.Connect()
	if err != nil {
		return nil, err
	}

	fmt.Println("success connection")
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

	auth := options.Credential{
		Username: authentication.User,
		Password: authentication.Psd,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(auth))
	Db = client.Database(authentication.Db)
	return client, ctx, cancelFunc, err
}
