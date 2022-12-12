package controller

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/interfaces/repository/product"
	"clean-golang/app/usecase/interactor"
	"io"
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
	body, _ := io.ReadAll(r.Body)

	p.ProductInteract.Store(body)
}
