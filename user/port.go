package user

import ("github.com/luminous479/product-list/domain"
 userHandler "github.com/luminous479/product-list/rest/handlers/user"
 )

type Service interface{
 userHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}


