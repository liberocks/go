-- Create index on order_id for better query performance
CREATE INDEX idx_items_orders_id ON items(order_id);