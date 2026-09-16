-- Migration: 0001_create_products_table.sql
-- Description: Creates the products table for the e-commerce product management API
-- Created at: 2024

-- Enable foreign keys support (SQLite specific)
PRAGMA foreign_keys = ON;

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(50) NOT NULL,
    sku VARCHAR(50) UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index on name for unique validation queries
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- Create index on category for filtering queries
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);

-- Create index on is_active for active product queries
CREATE INDEX IF NOT EXISTS idx_products_is_active ON products(is_active);

-- Create index on created_at for ordering queries
CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);

-- Insert sample data for testing
INSERT INTO products (name, description, price, stock, category, sku, is_active) VALUES
    ('Laptop Pro 15', 'Laptop de alta gama con procesador i7 y 16GB RAM', 1299.99, 25, 'electronics', 'LAP-PRO-15', 1),
    ('Camiseta Algodón', 'Camiseta de algodón orgánica color negro', 19.99, 150, 'clothing', 'CAM-alg-001', 1),
    ('Café Premium 1kg', 'Café premium tostado medio de Colombia', 24.50, 80, 'food', 'CAF-pre-1kg', 1),
    ('Novela El Principito', 'Edición especial del clásico de Saint-Exupéry', 12.99, 200, 'books', 'NOV-PRI-001', 1),
    ('Lámpara LED Escritorio', 'Lámpara LED regulable para escritorio', 35.00, 60, 'home', 'LAM-LED-001', 1);

-- Create trigger to update updated_at timestamp on row update
CREATE TRIGGER IF NOT EXISTS update_products_timestamp
AFTER UPDATE ON products
BEGIN
    UPDATE products SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Create view for active products
CREATE VIEW IF NOT EXISTS active_products AS
SELECT 
    id,
    name,
    description,
    price,
    stock,
    category,
    sku,
    created_at,
    updated_at
FROM products
WHERE is_active = 1;

-- Create view for low stock products (stock < 10)
CREATE VIEW IF NOT EXISTS low_stock_products AS
SELECT 
    id,
    name,
    price,
    stock,
    category,
    sku
FROM products
WHERE is_active = 1 AND stock < 10
ORDER BY stock ASC;