package repo

import "errors"

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"pass"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (User, error)
	Find(email string) (*User, error)
}
type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	 return &userRepo{}

}

func (r *userRepo) Create(u User) (User, error) {
	if u.ID != 0 {
		return u, errors.New("user already exists")
	}
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return u, nil
}

func (r *userRepo) Find(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}