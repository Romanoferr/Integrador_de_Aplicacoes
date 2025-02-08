-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id int AUTO_INCREMENT PRIMARY KEY,
    operation_date DATE NOT NULL,
    asset_type VARCHAR(255) NOT NULL,
    asset_id VARCHAR(255) NOT NULL,
    operation VARCHAR(255) NOT NULL,
    quantity DECIMAL(10, 4) NOT NULL,
    price DECIMAL(10, 4) NOT NULL,
    asset_manager VARCHAR(255) NOT NULL,
    KEY(id)
);

-- +goose Down

DROP TABLE IF EXISTS transactions;