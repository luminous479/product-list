package user

import "github.com/luminous479/product-list/domain"

type Service interface {
	Create(u domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
