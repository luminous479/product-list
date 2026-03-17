package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/luminous479/product-list/config"
)

func GetConnectionString(conf *config.DBConfig) string {

	conString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)

	if !conf.EnableSSLMODE {
		return conString + " sslmode=disable"
	}

	return conString
}

func NewConnection(conf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(conf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
