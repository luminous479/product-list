package user

import "github.com/luminous479/product-list/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(u domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(u)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil

}

func (svc *service) Find(email, pass string) (*domain.User, error) {

	usr, err := svc.userRepo.Find(email, pass)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil

}
