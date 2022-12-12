package interactor

import (
	"clean-golang/app/interfaces/repository/product"
	"encoding/json"
	"log"
)

type ProductInteract struct {
	ProductRepository product.ProductRepository
}

type Type struct {
	Types string `json:"types"`
}

func getType(body []byte) string {
	var t Type
	err := json.Unmarshal(body, &t)
	if err != nil {
		log.Fatalln(err)
	}

	return t.Types
}

func (u ProductInteract) Store(b []byte) {
	t := getType(b)
	u.ProductRepository.Store(t, b)
}
