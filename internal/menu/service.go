package menu

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidPrice = errors.New("invalid price")
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, price int, isActive bool) (Item, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Item{}, ErrInvalidName
	}
	if price < 0 {
		return Item{}, ErrInvalidPrice
	}
	return s.repo.Create(ctx, name, price, isActive)
}

func (s *Service) Update(ctx context.Context, id int64, name string, price int, isActive bool) (Item, error) {
	name = strings.TrimSpace(name)
	if id <= 0 {
		return Item{}, errors.New("invalid id")
	}
	if name == "" {
		return Item{}, ErrInvalidName
	}
	if price < 0 {
		return Item{}, ErrInvalidPrice
	}
	return s.repo.Update(ctx, id, name, price, isActive)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) ListAll(ctx context.Context) ([]Item, error) {
	return s.repo.List(ctx)
}

func (s *Service) ListPublic(ctx context.Context) ([]Item, error) {
	return s.repo.ListActive(ctx)
}
