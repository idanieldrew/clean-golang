package controller

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/interfaces/repository/product"
	"clean-golang/app/usecase/interactor"
	"github.com/gorilla/mux"
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

	res, status := p.ProductInteract.Store(body)

	rp.Res(w, status, res)
}

func (p *ProductController) FindBySlug(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)

	res, status := p.ProductInteract.FindBySlug(q["slug"])

	rp.Res(w, status, res)
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)
	body, _ := io.ReadAll(r.Body)
	status := p.ProductInteract.Update(q["slug"], body)

	rp.Res(w, status, "successfully update")
}
