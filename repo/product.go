package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/luminous479/product-list/domain"
)



type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {

	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(query, p.Name, p.Price).Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {

	query := `SELECT id, name, price FROM products WHERE id = $1`

	var product domain.Product

	err := r.db.Get(&product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &product, nil
}
func (r *productRepo) List(page int, limit int) ([]*domain.Product, error) {
	var list []*domain.Product

	// calculate offset
	offset := (page - 1) * limit

	query := `
		SELECT id, name, price 
		FROM products
		LIMIT $1 OFFSET $2
	`

	err := r.db.Select(&list, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {

	query := `UPDATE products SET name=$1, price=$2 WHERE id=$3`

	result, err := r.db.Exec(query, p.Name, p.Price, p.ID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &p, nil
}

func (r *productRepo) Delete(id int) error {

	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}