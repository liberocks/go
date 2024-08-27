-- Create index on product_id for better query performance
CREATE INDEX idx_variants_product_id ON variants(product_id);