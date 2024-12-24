-- Migration for creating database tables

-- Create Customers table
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    signup_date DATE NOT NULL,
    location VARCHAR(255),
    lifetime_value NUMERIC(10, 2) NOT NULL
);

-- Create Products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- Create Orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    order_date DATE NOT NULL,
    status VARCHAR(10) NOT NULL CHECK (status IN ('PENDING', 'COMPLETED', 'CANCELED'))
);

-- Create Order_Items table
CREATE TABLE order_items (
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    price NUMERIC(10, 2) NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

-- Create Transactions table
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    payment_status VARCHAR(10) NOT NULL CHECK (payment_status IN ('SUCCESS', 'FAILED')),
    payment_date DATE NOT NULL,
    total_amount NUMERIC(10, 2) NOT NULL
);

-- Add indexes for optimization
CREATE INDEX idx_customers_signup_date ON customers(signup_date);
CREATE INDEX idx_orders_order_date ON orders(order_date);
CREATE INDEX idx_transactions_payment_date ON transactions(payment_date);

-- Add table partitioning (if needed, use declarative partitioning in PostgreSQL 11+)
-- Example: Partitioning orders by order_date
CREATE TABLE orders_2024 PARTITION OF orders FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
