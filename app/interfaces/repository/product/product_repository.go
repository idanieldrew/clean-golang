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

func (p *ProductRepository) Store(b []byte) {
	m := make(map[string]string)
	uErr := json.Unmarshal(b, &m)

	req := new(request.Product)
	req.Name = m["name"]
	req.Price = m["price"]
	delete(m, "name")
	delete(m, "price")
	req.Fields = m

	if uErr != nil {
		logger.Error(uErr.Error())
		return
	}

	insert, err := p.Connection.Collection("products").InsertOne(context.TODO(), req)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(insert)
}
