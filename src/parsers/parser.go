package parsers

import (
	"log"
	"strconv"
	"strings"
	"time"

	transaction "main/data-access/transaction"

	"github.com/xuri/excelize/v2"
)

func ParseTransactionFile() [] transaction.Transaction {
	// Open the Excel file
	filePath := "../arquivos-statusinvest/StatusInvest-transactions-2025-01-22--23-43-32.xlsx"
	xlFile, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer func(xlFile *excelize.File) {
		err := xlFile.Close()
		if err != nil {

		}
	}(xlFile)

	// Read the "Carteira" sheet
	sheetName := "Carteira"
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to read sheet %q: %v", sheetName, err)
	}

	// Parse rows into transactions
	var transactions []transaction.Transaction
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

		// Parse the row into a Transaction
		var assetType string
		if row[1] == "Ações" {
			assetType = "Acao"
		} else if row[1] == "Fundos imobiliários" {
			assetType = "FII"
		} else {
			assetType = row[1]
		}

		quantity, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.Replace(row[4], ".", "", -1), ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse Quantity in row %d: %v", i+1, err)
			continue
		}

		price, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.Replace(row[5], ".", "", -1), ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse Price in row %d: %v", i+1, err)
			continue
		}

		tx := transaction.Transaction{
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
	// for _, tx := range transactions {
	// 	fmt.Printf("%+v\n", tx)
	// }

	return transactions
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
