package products

type product struct {
	Id    uint64 `db:"product_id"`
	Title string `db:"product_title"`
}
