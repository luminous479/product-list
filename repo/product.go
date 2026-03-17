package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product,error)
	Delete(id int) ( error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {

	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(query, p.Name, p.Price).Scan(p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {

	query := `SELECT id, name, price FROM products WHERE id = $1`

	var product Product

	err := r.db.Get(&product, query, id)

	if err != nil {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {

	var list []*Product
	query := `SELECT id, name, price FROM products`

	err := r.db.Select(&list, query)

	if err != nil {
		return nil, errors.New("product not found")
	}

	return list, nil
}

func (r *productRepo) Update(p Product) (*Product,error) {

	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3;"

	row := r.db.QueryRow(query,p.Name,p.Price,p.ID)
    err :=row.Err()

	if err != nil {
		return nil, err
	}
	return &p, nil

}

func (r *productRepo) Delete(id int) ( error) {

	query := `DELETE FROM products WHERE id = $1`


_, err := r.db.Exec(query, id)
    
if err != nil {
    return  errors.New("product not found")
}

	return nil
}

