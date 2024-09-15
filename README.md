# Hexagonal in Go Way

   This project demonstrates a hexagonal architecture implementation in Go.

## Project Structure

- `cmd/`: Contains the main application entry point.
- `internal/`: Contains the internal application code.
  - `adapters/`: Contains the adapter implementations.
    - `driven/`: Contains the driven adapters (e.g., PostgreSQL, Redis).
    - `driving/`: Contains the driving adapters (e.g., HTTP handlers).
  - `application/`: Contains the core application logic and interfaces.

## Prerequisites

- Docker
- Docker Compose
- Go (if running outside Docker)

## Getting Started

### Running with Docker Compose

1. Ensure Docker and Docker Compose are installed on your machine.
2. Clone the repository:
   ```sh
   git clone git@github.com:trumboosahil/hexagonal_in_go_way.git
   cd hexagonal_in_go_way
   ```
3. Build and start the services:
   ```sh
   docker compose up --build
   ```
4. The application should now be running and accessible at

```curl
curl  -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{
  "ID": "999234",
  "Amount": 150.50,
  "Status": "Pending"
}'
```

### Running Locally

1. Ensure Go is installed on your machine.
2. Clone the repository:
   ```sh
   git clone git@github.com:trumboosahil/hexagonal_in_go_way.git
   cd hexagonal_in_go_way
   ```
3. Set up the environment variables:
   ```sh
   export STORAGE_TYPE=postgres
   export POSTGRES_CONNECTION_STRING=postgres://user:password@localhost:5434/mydb?sslmode=disable
   ```
4. Run the application:
   ```sh
   go run cmd/main.go
   ```

## Database Initialization

   The PostgreSQL database will be initialized with the `init.sql` script located in the `db` directory. This script creates the necessary tables for the application.

# Change the repostry

update `docker-compose.yml` file
`STORAGE_TYPE=redis`

## Contributing

   Contributions are welcome! Please open an issue or submit a pull request.

## License

   This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
