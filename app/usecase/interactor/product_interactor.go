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

func (p ProductInteract) Store(b []byte) (any, int) {
	res, err := p.ProductRepository.Store(b)
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

func (p ProductInteract) Update(s string, body []byte) (int, string) {
	err := p.ProductRepository.Update(s, body)
	if err != nil {
		return http.StatusInternalServerError, "unsuccessfully update"
	}
	return http.StatusOK, "successfully update"
}

func (p ProductInteract) DestroyBySlug(s string) (int, string) {
	err := p.ProductRepository.Destroy(s)
	if err != nil {
		return http.StatusInternalServerError, "successfully destroy"
	}
	return http.StatusOK, "successfully destroy"
}
