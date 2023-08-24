package products

import (
	"log"
	"net/http"
)

type productsController struct {
	productsService *productsService
}

func newProductsController(productsService *productsService) *productsController {
	return &productsController{productsService}
}

func (pc productsController) IndexPage(w http.ResponseWriter, r *http.Request) {
	login, errCookie := r.Cookie("login")
	template, err := pc.productsService.IndexPage()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	if errCookie == nil {
		err = template.Execute(w, map[string]any{
			"UserLogin": login.Value,
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
	} else {
		err = template.Execute(w, nil)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
	}

}

func (pc productsController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	template, products, err := pc.productsService.GetAllProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	err = template.Execute(w, map[string]any{
		"products": products,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
}
