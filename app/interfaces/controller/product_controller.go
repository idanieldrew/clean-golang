package controller

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/interfaces/repository/product"
	product2 "clean-golang/app/interfaces/request/product"
	"clean-golang/app/usecase/interactor"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

type Type struct {
	Types string `json:"types"`
}

func getType(body []byte) string {
	var t Type
	err := json.Unmarshal(body, &t)
	if err != nil {
		log.Fatalln(err, 44)
	}
	//body, _ := io.ReadAll(r.ParseForm())
	return t.Types
}

func (p *ProductController) Store(w http.ResponseWriter, r *http.Request) {
	var req any
	body, _ := io.ReadAll(r.Body)
	t := getType(body)

	switch t {
	case "refrigerator":
		req = new(product2.Refrigerator)
		uErr := json.Unmarshal(body, req)
		if uErr != nil {
			logger.Error(uErr.Error())
			rp.Res(w, http.StatusInternalServerError, nil)
			return
		}
	case "vacuum":
		req = new(product2.VacuumCleaner)
		uErr := json.Unmarshal(body, req)
		if uErr != nil {
			logger.Error(uErr.Error())
			rp.Res(w, http.StatusInternalServerError, nil)
			return
		}
	}

	insert, err := mongo.Db.Collection(t).InsertOne(context.TODO(), req)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(insert)
}
