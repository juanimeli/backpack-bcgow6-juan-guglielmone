package product

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"
)

type Respository interface {
	Store(name, productType string, count int, price float64) (domain.Product, error)
	GetOne(id int) domain.Product
	Update(product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepo() Respository {
	return &repository{}
}

func (r *repository) Store(name, productType string, count int, price float64) (domain.Product, error) {
	return domain.Product{}, nil
}
func (r *repository) GetOne(id int) domain.Product {
	return domain.Product{}
}
func (r *repository) Update(product domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}
func (r *repository) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}
func (r *repository) Delete(id int) error {
	return nil
}
