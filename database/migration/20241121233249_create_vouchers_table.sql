-- +goose Up
CREATE TABLE vouchers (
    id INT IDENTITY(1,1) PRIMARY KEY,
    brand_id INT NOT NULL,
    code VARCHAR(255) NOT NULL,
    cost_in_points INT NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

-- +goose Down
DROP TABLE vouchers;
