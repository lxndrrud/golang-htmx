package products

import "fmt"

type productsDataService struct {
	productsRepo *productsRepo
}

func newProductsDataService(productsRepo *productsRepo) *productsDataService {
	return &productsDataService{productsRepo}
}

func (pds productsDataService) GetAllProducts() ([]product, error) {
	products, err := pds.productsRepo.GetAllProducts()
	if err != nil {
		err = fmt.Errorf("productsDataService - GetAllProducts: %v", err.Error())
		return products, err
	}
	return products, err
}
