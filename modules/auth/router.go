package auth

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func AuthRouterBootstrap(subrouter *mux.Router, connection *sqlx.DB) {
	repo := newAuthRepo(connection)
	dataService := newAuthDataService(repo)
	templateService := newAuthTemplateService()
	mainService := newAuthService(dataService, templateService)
	controller := newAuthController(mainService)

	subrouter.HandleFunc("/login", controller.LoginPage).Methods("GET")
	subrouter.HandleFunc("/login", controller.LoginAction).Methods("POST")
	subrouter.HandleFunc("/logout", controller.LogoutAction).Methods("POST")
}
