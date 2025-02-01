-- +goose Up
CREATE TABLE IF NOT EXISTS asset_allocations (
    id int AUTO_INCREMENT NOT NULL,
    asset_allocation_date DATE NOT NULL,
    asset_owner VARCHAR(255) NOT NULL,
    asset_id VARCHAR(255) NOT NULL,
    asset_type VARCHAR(255) NOT NULL,
    median_price DECIMAL(10, 6),
    actual_price DECIMAL(10, 2),
    median_return DECIMAL(10, 2) NOT NULL,
    quantity DECIMAL(10, 6),
    balance DECIMAL(10, 2) NOT NULL,
    today_return DECIMAL(10, 2),
    KEY(id),
    UNIQUE (asset_allocation_date, asset_owner, asset_id)
);

-- +goose Down

DROP TABLE IF EXISTS asset_allocations;