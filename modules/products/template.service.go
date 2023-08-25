package products

import (
	"fmt"
	"html/template"
)

type productsTemplateService struct{}

func newProductsTemplateService() *productsTemplateService {
	return &productsTemplateService{}
}

func (pts productsTemplateService) IndexPageTemplate() (*template.Template, error) {
	template, err := template.ParseFiles(
		"shared/templates/layouts/basic-layout.html",
		"shared/templates/auth/auth-header.html",
		"modules/products/templates/pages/index.html",
	)
	if err != nil {
		err = fmt.Errorf("ProductsTemplateService - IndexPageTemplate: %v", err.Error())
		return template, err
	}
	return template, nil
}

func (pts productsTemplateService) GetAllProducts() (*template.Template, error) {
	template, err := template.ParseFiles(
		"modules/products/templates/products/products-list.html",
		"modules/products/templates/products/product.html",
	)
	if err != nil {
		err = fmt.Errorf("ProductsTemplateService - GetAllProducts: %v", err.Error())
		return template, err
	}
	return template, nil
}
