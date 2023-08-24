package products

import "html/template"

type productsService struct {
	productsDataService     *productsDataService
	productsTemplateService *productsTemplateService
}

func newProductsService(productsDataService *productsDataService,
	productsTemplateService *productsTemplateService) *productsService {

	return &productsService{productsDataService, productsTemplateService}
}

func (ps productsService) IndexPage() (*template.Template, error) {
	return ps.productsTemplateService.IndexPageTemplate()
}

func (ps productsService) GetAllProducts() (*template.Template, []product, error) {
	productsChan := make(chan []product)
	templateChan := make(chan *template.Template)
	errChan := make(chan error)

	go func() {
		value, err := ps.productsDataService.GetAllProducts()
		productsChan <- value
		errChan <- err
	}()

	go func() {
		value, err := ps.productsTemplateService.GetAllProducts()
		templateChan <- value
		errChan <- err
	}()

	products := <-productsChan
	template := <-templateChan
	err := <-errChan

	return template, products, err
}
