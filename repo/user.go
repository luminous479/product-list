package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"pass" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}
type UserRepo interface {
	Create(u User) (*User, error)
	Find(email,pass string) (*User, error)
}
type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}

}

func (r *userRepo) Create(u User) (*User, error) {

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

func (r *userRepo) Find(email, pass string) (*User, error) {

	query := `
	SELECT id, name, email, password, is_shop_owner
	FROM users
	WHERE email=$1 AND password=$2
	`

	var user User

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
