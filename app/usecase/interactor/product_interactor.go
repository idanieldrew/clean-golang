package interactor

import (
	"clean-golang/app/interfaces/repository/product"
	"net/http"
)

type ProductInteract struct {
	ProductRepository product.ProductRepository
}

type Type struct {
	Types string `json:"types"`
}

func (u ProductInteract) Store(b []byte) (any, int) {
	res, err := u.ProductRepository.Store(b)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return res, http.StatusOK
}

func (p ProductInteract) FindBySlug(s string) (any, int) {
	res, err := p.ProductRepository.FindBySlug(s)
	if err != nil {
		return nil, http.StatusNotFound
	}
	return res, http.StatusOK
}

func (p ProductInteract) Update(s string, body []byte) int {
	err := p.ProductRepository.Update(s, body)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
