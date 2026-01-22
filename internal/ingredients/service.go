package ingredients

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidStock = errors.New("invalid stock")
	ErrInvalidDelta = errors.New("invalid delta")
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, stock int) (Ingredient, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Ingredient{}, ErrInvalidName
	}
	if stock < 0 {
		return Ingredient{}, ErrInvalidStock
	}
	return s.repo.Create(ctx, name, stock)
}

func (s *Service) Update(ctx context.Context, id int64, name string, stock int) (Ingredient, error) {
	name = strings.TrimSpace(name)
	if id <= 0 {
		return Ingredient{}, errors.New("invalid id")
	}
	if name == "" {
		return Ingredient{}, ErrInvalidName
	}
	if stock < 0 {
		return Ingredient{}, ErrInvalidStock
	}
	return s.repo.Update(ctx, id, name, stock)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]Ingredient, error) {
	return s.repo.List(ctx)
}

func (s *Service) Restock(ctx context.Context, id int64, delta int) (Ingredient, error) {
	if id <= 0 {
		return Ingredient{}, errors.New("invalid id")
	}
	if delta <= 0 {
		return Ingredient{}, ErrInvalidDelta
	}
	return s.repo.Restock(ctx, id, delta)
}
