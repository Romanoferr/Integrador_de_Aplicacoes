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

insert into transactions (operation_date, asset_type, asset_id, operation, quantity, price, asset_manager) 
values 
('2021-01-01', 'ACOES', 'SAPR4', 'C', 105.00, 5.05, 'INTER DTVM LTDA'),
('2025-01-02', 'Tesouro direto','tesouro-ipca-2035','C', 0.05000000, 2096.54,'INTER DTVM LTDA'),
('2024-01-01', 'ACOES', 'SAPR4', 'C', 65.00, 3.05, 'INTER DTVM LTDA');
