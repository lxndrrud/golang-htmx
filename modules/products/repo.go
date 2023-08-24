package products

import "github.com/jmoiron/sqlx"

type productsRepo struct {
	connection *sqlx.DB
}

func newProductsRepo(connection *sqlx.DB) *productsRepo {
	return &productsRepo{connection}
}

func (pr productsRepo) GetAllProducts() ([]product, error) {
	var products []product

	err := pr.connection.Select(
		&products,
		`
		SELECT 
			p.id as product_id,
			p.title as product_title
		FROM products as p
		`,
	)
	return products, err
}
