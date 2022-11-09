package product

import "github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"

type Service interface {
	Store(name, productType string, count int, price float64) (domain.Product, error)
	GetOne(id int) domain.Product
	Update(product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}

type service struct {
	repo Respository
}

func NewService(repo Respository) Service {
	return &service{repo: repo}
}

func (s *service) Store(name, productType string, count int, price float64) (domain.Product, error) {
	return domain.Product{}, nil
}
func (s *service) GetOne(id int) domain.Product {
	return domain.Product{}
}
func (s *service) Update(product domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}
func (s *service) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}
func (s *service) Delete(id int) error {
	return nil
}
