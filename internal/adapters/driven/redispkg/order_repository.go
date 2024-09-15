package redispkg

import (
	"context"
	"encoding/json"
	"hexagonal_in_go_way/internal/domain/entities"

	"github.com/go-redis/redis/v8"
)

type OrderRepository struct {
	client *redis.Client
}

func NewOrderRepository(client *redis.Client) *OrderRepository {
	return &OrderRepository{client: client}
}

func (r *OrderRepository) Save(ctx context.Context, order *entities.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "orders",
		Values: map[string]interface{}{"order": data},
	}).Err()
}
