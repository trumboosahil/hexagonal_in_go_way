   -- init.sql
   CREATE TABLE IF NOT EXISTS orders (
       id SERIAL PRIMARY KEY,
       amount DECIMAL(10, 2) NOT NULL,
       status VARCHAR(50) NOT NULL
   );