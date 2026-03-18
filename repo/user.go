package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/luminous479/product-list/domain"
)


type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Find(email,pass string) (*domain.User, error)
}
type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}

}

func (r *userRepo) Create(u domain.User) (*domain.User, error) {

	query := `
	INSERT INTO users (name, email, password, is_shop_owner)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		u.Name,
		u.Email,
		u.Password,
		u.IsShopOwner,
	).Scan(&u.ID)

	if err != nil {

		return nil, err

	}
	return &u, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {

	query := `
	SELECT id, name, email, password, is_shop_owner
	FROM users
	WHERE email=$1 AND password=$2
	`

	var user domain.User

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
