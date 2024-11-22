-- +goose Up
CREATE TABLE transactions (
    id INT IDENTITY(1,1) PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    total_points INT NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE()
);

-- +goose Down
DROP TABLE transactions;
