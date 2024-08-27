-- Create variants table
CREATE TABLE variants (
    id SERIAL PRIMARY KEY,
    variant_name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
);