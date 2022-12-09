package interactor

import "clean-golang/app/interfaces/repository/product"

type ProductInteract struct {
	ProductRepository product.ProductRepository
}

func (u UserInteract) Store() {

}
