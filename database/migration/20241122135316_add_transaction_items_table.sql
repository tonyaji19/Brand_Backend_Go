-- +goose Up
CREATE TABLE transaction_items (
    id INT IDENTITY(1,1) PRIMARY KEY,
    transaction_id INT NOT NULL,
    voucher_id INT NOT NULL,
    quantity INT NOT NULL,
    total_points INT NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);

-- +goose Down
DROP TABLE transaction_items;
