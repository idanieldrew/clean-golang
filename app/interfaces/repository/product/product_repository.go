package product

import (
	"clean-golang/app/infrastructure/logger"
	request "clean-golang/app/interfaces/request/product"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Connection *mongo.Database
	Cache      *redis.Client
}

func (p *ProductRepository) Store(collection string, b []byte) {
	var req any
	switch collection {
	case "refrigerator":
		req = new(request.Refrigerator)
		uErr := json.Unmarshal(b, req)

		if uErr != nil {
			logger.Error(uErr.Error())
			return
		}
	case "vacuum":
		req = new(request.VacuumCleaner)
		uErr := json.Unmarshal(b, req)
		if uErr != nil {
			logger.Error(uErr.Error())
			return
		}
	}

	insert, err := p.Connection.Collection(collection).InsertOne(context.TODO(), req)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(insert)
}
