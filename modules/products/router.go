package products

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func ProductsRouterBootstrap(subrouter *mux.Router, connection *sqlx.DB) {
	repo := newProductsRepo(connection)
	dataService := newProductsDataService(repo)
	templateService := newProductsTemplateService()
	mainService := newProductsService(dataService, templateService)
	controller := newProductsController(mainService)

	subrouter.HandleFunc("/", controller.IndexPage)
	subrouter.HandleFunc("/getAll", controller.GetAllProducts)
}
