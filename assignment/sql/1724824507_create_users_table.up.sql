-- Create items table
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT,
    is_validated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

 
