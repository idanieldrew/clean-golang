package controller

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/interfaces/repository/product"
	"clean-golang/app/usecase/interactor"
	"context"
	"fmt"
	"net/http"
)

type ProductController struct {
	ProductInteract interactor.ProductInteract
}

func NewProduct() *ProductController {
	return &ProductController{
		ProductInteract: interactor.ProductInteract{
			ProductRepository: product.ProductRepository{
				Connection: mongo.Db,
				Cache:      redis.Rdb,
			}},
	}
}

func (p *ProductController) Store(w http.ResponseWriter, r *http.Request) {
	Vacuumcleaner := struct {
		Name    string  `bson:"name"`
		Price   float32 `bson:"price"`
		Color   string  `bson:"color"`
		Suction int     `bson:"Suction"`
		Type    string  `bson:"type"`
	}{
		Name:    "samsung-new-2022",
		Price:   500.99,
		Color:   "white",
		Suction: 3000,
		Type:    "vacuum_cleaner",
	}

	insert, err := mongo.Db.Collection("products").InsertOne(context.TODO(), Vacuumcleaner)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(insert)
}
