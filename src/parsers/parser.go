package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
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


func parseTransactionFile() {
	// Open the Excel file
	filePath := "../../arquivos-statusinvest/StatusInvest-transactions-2025-01-22--23-43-32.xlsx"
	xlFile, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer xlFile.Close()

	// Read the "Carteira" sheet
	sheetName := "Carteira"
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to read sheet %q: %v", sheetName, err)
	}

	// Parse rows into transactions
	var transactions []Transaction
	var idCounter int64 = 1

	// Skip the header row
	for i, row := range rows {
		if i == 0 {
			continue
		}

		// Ensure the row has enough columns
		if len(row) < 7 {
			log.Printf("Skipping incomplete row: %v", row)
			continue
		}

		var assetType string
		if row[1] == "Ações" {
			assetType = "Acao"
		} else if row[1] == "Fundos imobiliários" {
			assetType = "FII"
		} else { assetType = row[1] }

		// Parse the row into a Transaction
		quantity, err := strconv.ParseFloat(
			strings.ReplaceAll(row[4], ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse Quantity in row %d: %v", i+1, err)
			continue
		}

		price, err := strconv.ParseFloat(
			strings.ReplaceAll(row[5], ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse Price in row %d: %v", i+1, err)
			continue
		}

		tx := Transaction{
			ID:            idCounter,
			OperationDate: parseDateTransactionFile(row[0]),
			AssetType:     assetType,
			AssetId:       row[2],
			Operation:     row[3],
			Quantity:      quantity,
			Price:         price,
			AssetManager:  row[6],
		}
		transactions = append(transactions, tx)
		idCounter++
	}

	// Print the parsed transactions
	for _, tx := range transactions {
		fmt.Printf("%+v\n", tx)
	}
}

// parseDateTransactionFile converts a date string from "dd/mm/yyyy" to "yyyy-mm-dd"
func parseDateTransactionFile(dateStr string) string {
	parsed, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		log.Printf("Failed to parse date %q: %v", dateStr, err)
		return ""
	}
	return parsed.Format("2006-01-02")
}


func main() {
    parseTransactionFile()
}