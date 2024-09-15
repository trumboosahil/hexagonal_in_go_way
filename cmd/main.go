package main

import (
	"context"
	"hexagonal_in_go_way/internal/adapters/driven/postgres"
	"hexagonal_in_go_way/internal/adapters/driven/redispkg"
	"hexagonal_in_go_way/internal/adapters/driving/http"
	"hexagonal_in_go_way/internal/application"

	"github.com/go-redis/redis/v8"

	"log"
	"os"
)

func main() {
	var orderRepo application.OrderRepository

	// Determine which repository to use based on an environment variable
	storageType := os.Getenv("STORAGE_TYPE") // "postgres" or "redis"
	if storageType == "" {
		storageType = "postgres"
	}

	switch storageType {
	case "postgres":
		connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
		if connStr == "" {
			connStr = "postgres://user:password@postgres:5432/mydb?sslmode=disable"
		}
		db, err := postgres.NewPostgresConnection(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}
		defer db.Close()
		orderRepo = postgres.NewOrderRepository(db)

	case "redis":
		var rdb *redis.Client
		connStr := os.Getenv("REDIS_CONNECTION_STRING")
		if connStr == "" {
			connStr = "redis://redis:6379"
		}
		opt, err := redis.ParseURL(connStr)
		if err != nil {
			log.Fatalf("Failed to parse Redis URL: %v", err)
		}
		rdb = redis.NewClient(opt)
		defer rdb.Close()

		// Verify the connection
		ctx := context.Background()
		_, err = rdb.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
		orderRepo = redispkg.NewOrderRepository(rdb)

	default:
		log.Fatalf("Unsupported storage type: %s", storageType)
	}

	orderService := application.NewOrderService(orderRepo)
	orderHandler := http.NewOrderHandler(orderService)
	server := http.NewServer(orderHandler)

	server.Start(":8080")
}
