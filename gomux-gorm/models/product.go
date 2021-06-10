package models

type Product struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) List(page, pageSize int) ([]Product, error) {
	var products []Product
	if err := DB.Scopes(Paginate(page, pageSize)).Find(&products).Error; err != nil {
		return []Product{}, err
	}
	return products, nil
}
func (p *Product) Create() error {
	return DB.Create(p).Error
}
func (p *Product) Read() error {
	return DB.Where("id = ?", p.ID).First(p).Error
}
func (p *Product) Update() error {
	return DB.Model(p).Updates(p).Error
}
func (p *Product) Delete() error {
	return DB.Delete(p).Error
}
