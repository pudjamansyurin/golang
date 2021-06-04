package models

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) create(db *sql.DB) error {
	return errors.New("implement me")
}
func (p *Product) read(db *sql.DB) error {
	return errors.New("implement me")
}
func (p *Product) update(db *sql.DB) error {
	return errors.New("implement me")
}
func (p *Product) delete(db *sql.DB) error {
	return errors.New("implement me")
}

func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	return nil, errors.New("implement me")
}
