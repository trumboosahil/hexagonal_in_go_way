package application

import (
	"context"
	"hexagonal_in_go_way/internal/domain/entities"
)

type OrderRepository interface {
	Save(ctx context.Context, order *entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *entities.Order) error {
	return s.repo.Save(ctx, order)
}
