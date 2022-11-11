package product

import (
	"context"
	"errors"

	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/domain"
)

type Service interface {
	GetOne(ctx context.Context, id int) (domain.Product, error)
	Store(ctx context.Context, p domain.Product) (int64, error)
	Update(ctx context.Context, p domain.Product, id int) error
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	GetFullData(ctx context.Context, id int) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Delete implements Service
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements Service
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

// GetOne implements Service
func (s *service) GetOne(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.GetOne(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// Store implements Service
func (s *service) Store(ctx context.Context, p domain.Product) (int64, error) {
	if s.repository.Exists(ctx, p.ID) {
		return 0, errors.New("error: producto ya existe")
	}
	id, err := s.repository.Store(ctx, p)
	if err != nil {
		return 0, err
	}
	p.ID = int(id) // es necesaria esta fila? funciona sin? PROBAR
	return id, nil
}

// Update implements Service
func (s *service) Update(ctx context.Context, p domain.Product, id int) error {
	err := s.repository.Update(ctx, p, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetFullData(ctx context.Context, id int) ([]domain.Product, error) {
	products, err := s.repository.GetFullData(ctx, id)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}
