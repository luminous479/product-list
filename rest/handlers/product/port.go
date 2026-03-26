package product

import "github.com/luminous479/product-list/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) error
	Count() (int64, error)
}