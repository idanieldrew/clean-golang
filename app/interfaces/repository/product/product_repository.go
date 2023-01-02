package product

import (
	"clean-golang/app/infrastructure/logger"
	request "clean-golang/app/interfaces/request/product"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Connection *mongo.Database
	Cache      *redis.Client
}

func (p *ProductRepository) Store(b []byte) (*mongo.InsertOneResult, error) {
	m := make(map[string]string)
	uErr := json.Unmarshal(b, &m)

	req := new(request.Product)
	req.Name = m["name"]
	req.Price = m["price"]
	req.Slug = m["slug"]
	delete(m, "name")
	delete(m, "price")
	delete(m, "slug")

	req.Fields = m

	if uErr != nil {
		logger.Error(uErr.Error())
		return nil, uErr
	}

	insert, err := p.Connection.Collection("products").InsertOne(context.TODO(), req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return insert, nil
}

func (p *ProductRepository) FindBySlug(s string) (*request.Product, error) {
	product := new(request.Product)
	err := p.Connection.Collection("products").FindOne(context.TODO(), bson.D{{"sluggi", s}}).Decode(product)

	if err != nil {
		text := fmt.Sprintf("%s not found", s)
		logger.Error(text)
		return nil, err
	}

	return product, nil
}
