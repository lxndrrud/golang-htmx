package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		generatedTemplate, _ := template.ParseFiles(
			"templates/layouts/basic-layout.html",
			"templates/pages/index.html")
		err := generatedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
	})
	router.HandleFunc("/loadProducts", func(w http.ResponseWriter, r *http.Request) {
		product := &Product{
			Name: "Test",
		}
		products := make([]*Product, 0)
		products = append(products, product)
		parsedTemplate, err := template.ParseFiles(
			"templates/domain/products/products-list.html",
			"templates/domain/products/product.html",
		)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
		err = parsedTemplate.Execute(w, map[string]interface{}{
			"products": products,
		})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
	})
	http.ListenAndServe(":9000", router)
}
