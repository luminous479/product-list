package product

import "github.com/luminous479/product-list/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) ProductRepo {
	return &service{
		productRepo: productRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	pro, err := svc.productRepo.Create(p)

	if err != nil {
		return nil, err
	}
	if pro == nil {
		return nil, nil
	}
	return pro, nil

}
func (svc *service) Get(id int) (*domain.Product, error) {

	pro, err := svc.productRepo.Get(id)

	if err != nil {
		return nil, err
	}
	if pro == nil {
		return nil, nil
	}
	return pro, nil

}
func (svc *service) List() ([]*domain.Product, error) {
	list, err := svc.productRepo.List()

	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, nil
	}
	return list, nil

}
func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	pro, err := svc.productRepo.Update(p)

	if err != nil {
		return nil, err
	}
	if pro == nil {
		return nil, nil
	}
	return pro, nil

}
func (svc *service) Delete(id int) error {

	err := svc.productRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
