package postgres

import (
	"context"
	"database/sql"
	"hexagonal_in_go_way/internal/domain/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(ctx context.Context, order *entities.Order) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO orders (id, amount, status) VALUES ($1, $2, $3)",
		order.ID, order.Amount, order.Status)
	return err
}
