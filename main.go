package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/golang-htmx/modules/auth"
	"github.com/lxndrrud/golang-htmx/modules/products"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connection, err := sqlx.Connect("sqlite3", "shop.db")
	if err != nil {
		log.Fatalln(err)
	}

	var schema = `
		DROP TABLE IF EXISTS users;
			
		CREATE TABLE users (
			id INTEGER PRIMARY KEY NOT NULL,
			login VARCHAR(30) NOT NULL,
			password VARCHAR(100) NOT NULL
		);

		INSERT INTO users(id, login, password) VALUES ($1, $2, $3);

		DROP TABLE IF EXISTS products;
		
		CREATE TABLE products (
			id INTEGER PRIMARY KEY NOT NULL,
			title VARCHAR(80) NOT NULL
		);

		INSERT INTO products(id, title) VALUES ($4, $5);
	`

	connection.MustExec(schema, 1, "lxndrrud", "testpassword", 1, "Тестовый продукт 1")

	router := mux.NewRouter()

	productsSubrouter := router.PathPrefix("/products").Subrouter()
	products.ProductsRouterBootstrap(productsSubrouter, connection)

	authSubrouter := router.PathPrefix("/auth").Subrouter()
	auth.AuthRouterBootstrap(authSubrouter, connection)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./shared/scripts/"))))
	http.ListenAndServe(":9000", router)
}
