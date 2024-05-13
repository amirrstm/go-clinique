-- Set timezone
SET
    TIMEZONE = "Asia/Dhaka";

-- Create address table
CREATE TABLE
    "addresses" (
        id serial PRIMARY KEY,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT NOW (),
            updated_at TIMESTAMP
        WITH
            TIME ZONE NULL,
            name VARCHAR(100) NOT NULL,
            is_active BOOLEAN DEFAULT TRUE,
            is_deleted BOOLEAN DEFAULT FALSE
    );