package product

import (
	"context"
	"errors"

	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"
)

type Service interface {
	Store(ctx context.Context, p domain.Product) (domain.Product, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Respository
}

func NewService(repo Respository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Store(ctx context.Context, p domain.Product) (domain.Product, error) {
	if s.repo.Exists(ctx, p.ID) {

		return domain.Product{}, errors.New("error: product id alrready exists")
	}
	product_id, err := s.repo.Store(ctx, p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return domain.Product{}, err
	}
	p.ID = int(product_id)
	return p, nil
}
func (s *service) GetOne(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repo.GetOne(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
func (s *service) Update(ctx context.Context, product domain.Product, id int) (domain.Product, error) {

	err := s.repo.Update(ctx, product, id)
	if err != nil {
		return domain.Product{}, err
	}
	productUpdated, err := s.repo.GetOne(ctx, id)
	if err != nil {
		return product, err
	}
	return productUpdated, nil
}
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
