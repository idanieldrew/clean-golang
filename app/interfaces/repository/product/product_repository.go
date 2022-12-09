package product

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Connection *mongo.Database
	Cache      *redis.Client
}

func (p *ProductRepository) store() {

}
