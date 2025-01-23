package transaction

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

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

var db *sql.DB

func connectDB() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "gomysql",
		Passwd: "gomysql",
		Addr:   "127.0.0.1:3306",
		DBName: "statusinvest",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	trs, err := transactionByAssetId("SAPR4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transactions found: %v\n", trs)

	trs2, err := transactionByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction found: %v\n", trs2)

	trs3 := Transaction{
		OperationDate: "2021-09-01",
		AssetType:     "ACOES",
		AssetId:       "ITSA4",
		Operation:     "C",
		Quantity:      10,
		Price:         13.00,
		AssetManager:  "INTER DTVM",
	}

	trs3IncludedID, err := addTransaction(trs3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction added with ID: %v\n", trs3IncludedID)
}

// transactionByAssetId queries for albums that have the specified artist name.
func transactionByAssetId(AssetId string) ([]Transaction, error) {
	// An transactions slice to hold data from returned rows.
	var transactions []Transaction

	rows, err := db.Query("SELECT id, operation_date, asset_type, asset_id, quantity, price, asset_manager FROM transactions WHERE asset_id = ?", AssetId)
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
func transactionByID(id int64) (Transaction, error) {
	var trs Transaction
	row := db.QueryRow("SELECT * FROM transactions WHERE id = ?", id)
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
func addTransaction(trs Transaction) (int64, error) {
	result, err := db.Exec("INSERT INTO transactions (operation_date, asset_type, asset_id, operation, quantity, price, asset_manager) VALUES (?, ?, ?, ?, ?, ?, ?)",
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
