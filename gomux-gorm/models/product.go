package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) Count() int {
	var count int64
	DB.Model(&Product{}).Count(&count)
	return int(count)
}

func (p *Product) All(page, pageSize int) ([]Product, error) {
	var products []Product

	DB.Scopes(Paginate(page, pageSize)).Find(&products)
	if err := DB.Scopes(Paginate(page, pageSize)).Find(&products).Error; err != nil {
		return []Product{}, err
	}
	return products, nil
}
func (p *Product) Create() error {
	return DB.Create(p).Error
}
func (p *Product) Read(id int) error {
	return DB.First(p, id).Error
}
func (p *Product) Update(id int) error {
	return DB.Model(&Product{}).Where("id = ?", id).Updates(p).Error
}
func (p *Product) Delete(id int) error {
	return DB.Delete(&Product{}, id).Error
}
