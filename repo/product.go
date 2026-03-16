package repo

import "errors"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) error
	Delete(id int) (*Product, error)
}

type productRepo struct {
	products []Product
}

func NewProductRepo() ProductRepo {
	repo := productRepo{}
	generateInitialProducts(&repo)
	return &repo
}

func (r *productRepo) Create(p Product) (*Product, error) {

	p.ID = len(r.products) + 1
	r.products = append(r.products, p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {

	for i := range r.products {
		if r.products[i].ID == id {
			return &r.products[i], nil
		}
	}

	return nil, errors.New("product not found")
}

func (r *productRepo) List() ([]*Product, error) {

	var list []*Product

	for i := range r.products {
		list = append(list, &r.products[i])
	}

	return list, nil
}

func (r *productRepo) Update(p Product) error {

	for i := range r.products {
		if r.products[i].ID == p.ID {
			r.products[i] = p
			return nil
		}
	}

	return errors.New("product not found")
}

func (r *productRepo) Delete(id int) (*Product, error) {

	for i := range r.products {
		if r.products[i].ID == id {
			deleted := r.products[i]
			r.products = append(r.products[:i], r.products[i+1:]...)
			return &deleted, nil
		}
	}

	return nil, errors.New("product not found")
}

func generateInitialProducts(repo *productRepo) {
	repo.products = []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 5.49},
	}
}