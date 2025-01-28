package transaction

import (
	"database/sql"
	"fmt"
	connectdb "main/data-access"
)

//var db *sql.DB

type Transaction struct {
	ID            int64
	OperationDate string
	AssetType     string
	AssetId       string
	Operation     string
	Quantity      float64
	Price         float64
	AssetManager  string
}

// transactionByAssetId queries for albums that have the specified artist name.
func TransactionByAssetId(AssetId string) ([]Transaction, error) {
	// An transactions slice to hold data from returned rows.
	var transactions []Transaction

	rows, err := connectdb.Db.Query("SELECT id, operation_date, asset_type, asset_id, quantity, price, asset_manager FROM transactions WHERE asset_id = ?", AssetId)
	if err != nil {
		return nil, fmt.Errorf("transactionByAssetId %q: %v", AssetId, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var trs Transaction
		if err := rows.Scan(&trs.ID, &trs.OperationDate, &trs.AssetType, &trs.AssetId, &trs.Quantity, &trs.Price, &trs.AssetManager); err != nil {
			return nil, fmt.Errorf("transactionByAssetId %q: %v", AssetId, err)
		}
		transactions = append(transactions, trs)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("transactionByAssetId %q: %v", AssetId, err)
	}
	return transactions, nil
}

// transactionByID queries for the transaction with the specified ID.
func TransactionByID(id int64) (Transaction, error) {
	var trs Transaction
	row := connectdb.Db.QueryRow("SELECT * FROM transactions WHERE id = ?", id)
	if err := row.Scan(&trs.ID, &trs.OperationDate, &trs.AssetType, &trs.AssetId, &trs.Operation, &trs.Quantity, &trs.Price, &trs.AssetManager); err != nil {
		if err == sql.ErrNoRows {
			return trs, fmt.Errorf("transactionById %d: no such transaction", id)
		}
		return trs, fmt.Errorf("transactionById %d: %v", id, err)
	}
	return trs, nil
}

// addTransaction adds a specific transaction to the database,
// returning the transaction ID of the new entry.
func AddTransaction(trs Transaction) (int64, error) {
	result, err := connectdb.Db.Exec("INSERT INTO transactions (operation_date, asset_type, asset_id, operation, quantity, price, asset_manager) VALUES (?, ?, ?, ?, ?, ?, ?)",
		trs.OperationDate, trs.AssetType, trs.AssetId, trs.Operation, trs.Quantity, trs.Price, trs.AssetManager)
	if err != nil {
		return 0, fmt.Errorf("addTransaction error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addTransaction error: %v", err)
	}
	return id, nil
}
