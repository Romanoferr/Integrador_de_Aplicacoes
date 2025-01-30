CREATE DATABASE IF NOT EXISTS statusinvest;

DROP TABLE IF EXISTS transactions;
CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT NOT NULL,
    operation_date DATE NOT NULL,
    asset_type VARCHAR(255) NOT NULL,
    asset_id VARCHAR(255) NOT NULL,
    operation VARCHAR(255) NOT NULL,
    quantity DECIMAL(10, 4) NOT NULL,
    price DECIMAL(10, 4) NOT NULL,
    asset_manager VARCHAR(255) NOT NULL,
    KEY(id)
);

-- insert into transactions (operation_date, asset_type, asset_id, operation, quantity, price, asset_manager)
-- values
-- ('2021-01-01', 'ACOES', 'SAPR4', 'C', 105.00, 5.05, 'INTER DTVM LTDA'),
-- ('2025-01-02', 'Tesouro direto','tesouro-ipca-2035','C', 0.05000000, 2096.54,'INTER DTVM LTDA'),
-- ('2024-01-01', 'ACOES', 'SAPR4', 'C', 65.00, 3.05, 'INTER DTVM LTDA');


DROP TABLE IF EXISTS asset_allocations;
CREATE TABLE IF NOT EXISTS asset_allocations (
    id INT AUTO_INCREMENT NOT NULL,
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

-- insert into assets (asset_allocation_date, asset_owner, asset_id, asset_type, median_price, actual_price, median_return, quantity, balance, today_return)
-- values
-- ('2021-01-01', 'Romano', 'SAPR4', 'Acao', 5.05, 5.05, 0.0100, 105.00, 5.05, 0.0025),
-- ('2025-01-02', 'Romano', 'tesouro-ipca-2035', 'Tesouro direto', 2096.54, 2096.54, 0.02574, 0.05000000, 2096.54, 0.012),
-- ('2024-01-01', 'Bruna', 'SAPR4', 'Acao', 3.05, 3.05, 0.00, 65.00, 3.05, 0.03201);

