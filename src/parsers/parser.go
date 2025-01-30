package parsers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	transaction "main/data-access/transaction"

	asset "main/data-access/asset-allocations"

	"github.com/xuri/excelize/v2"
)

type AssetSheet struct {
	FileName string
	SheetName string
	Owner string
}

var assetSheetCollection = []string{"Ações", "FIIs", "Tesouro", "ETF", "ETF Exterior"}

func ParseTransactionFile(fullPathTransactionFile string) ([] transaction.Transaction, error) {
	// Open the Excel file
	xlFile, err := excelize.OpenFile(fullPathTransactionFile)
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
			OperationDate: ParseDateTransactionFile(row[0]),
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

	return transactions, nil
}


func ParseFixedIncome(filePath, sheetName, onwer string) ([]asset.AssetAllocation, error) {
	// Open the Excel file
	xlFile, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer xlFile.Close()

	// Read the specified sheet
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	// Parse rows into AssetAllocation structs
	var allocations []asset.AssetAllocation
	var idCounter int64 = 1

	for i, row := range rows {
		if i == 0 {
			continue
		}

		// Ensure the row has enough columns
		if len(row) < 6 {
			log.Printf("Skipping incomplete row %d: %v", i+1, row)
			continue
		}

		parseFloat := func(value string) float64 {
			val, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Printf("Failed to parse float value %q: %v", value, err)
				return 0.0
			}
			return val
		}

		medianReturn, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.ReplaceAll(row[6], "%", ""), ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse MedianReturn in row %d: %v", i+1, err)
			continue
		}

	allocation := asset.AssetAllocation{
		ID:              idCounter,
		AssetAllocationDate: ExtractDateFromFilePath(filePath),
		AssetOwner: onwer,
		AssetIdentifier: row[0],
		AssetType:       "Renda Fixa",
		MedianReturn:    medianReturn,
		Balance:         parseFloat(row[5]),
	}

	allocations = append(allocations, allocation)
	idCounter++
	}

	return allocations, nil
}


func ParseAllSheetsAssetAllocations(filePath string, onwer string) ([]asset.AssetAllocation, error) {
	var allocations []asset.AssetAllocation
	for _, sheet := range assetSheetCollection {
		allocationsB, err := ParseAssetAllocations(filePath, sheet, onwer)
		if err != nil {
			log.Printf("Error parsing asset allocations for owner %s: %v", onwer, err)
		}
		allocations = append(allocations, allocationsB...)
	}
	
	parseFixedIncome, err := ParseFixedIncome(filePath, "CDB LCI LCA LC RDB", onwer)
	if err != nil {
		log.Printf("Error parsing fixed income: %v", err)
	}
	allocations = append(allocations, parseFixedIncome...)
	return allocations, nil
}

func ParseAssetAllocations(filePath, sheetName string, onwer string) ([]asset.AssetAllocation, error) {
	// Open the Excel file
	xlFile, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer xlFile.Close()

	// Read the specified sheet
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	// Parse rows into AssetAllocation structs
	var allocations []asset.AssetAllocation
	var idCounter int64 = 1

	// Skip the header row
	for i, row := range rows {
		if i == 0 {
			continue
		}

		// Ensure the row has enough columns
		if len(row) < 7 {
			log.Printf("Skipping incomplete row %d: %v", i+1, row)
			continue
		}

		var assetType string
		if row[1] == "Ações" {
			assetType = "Acao"
		} else {
			assetType = row[1]
		}

		// Parse float values safely
		parseFloat := func(value string) float64 {
			val, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Printf("Failed to parse float value %q: %v", value, err)
				return 0.0
			}
			return val
		}

		medianReturn, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.ReplaceAll(row[4], "%", ""), ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse MedianReturn in row %d: %v", i+1, err)
			continue
		}

		todayReturn, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.ReplaceAll(row[7], "%", ""), ",", "."), 64,
		)
		if err != nil {
			log.Printf("Failed to parse MedianReturn in row %d: %v", i+1, err)
			continue
		}
		
		allocation := asset.AssetAllocation{
			ID:              idCounter,
			AssetAllocationDate: ExtractDateFromFilePath(filePath),
			AssetOwner: onwer,
			AssetIdentifier: row[0],
			AssetType:       assetType,
			MedianPrice:     parseFloat(row[2]),
			ActualPrice:     parseFloat(row[3]),
			MedianReturn:    medianReturn,
			Quantity:        parseFloat(row[5]),
			Balance:         parseFloat(row[6]), 
			TodayReturn:     todayReturn, 
		}

		allocations = append(allocations, allocation)
		idCounter++
	}

	return allocations, nil
}


// parseDateTransactionFile converts a date string from "dd/mm/yyyy" to "yyyy-mm-dd"
func ParseDateTransactionFile(dateStr string) string {
	parsed, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		log.Printf("Failed to parse date %q: %v", dateStr, err)
		return ""
	}
	return parsed.Format("2006-01-02")
}

func ExtractDateFromFilePath(filePath string) (string) {
	// Define a regular expression to match the date in the format yyyy-mm-dd
	datePattern := `\d{4}-\d{2}-\d{2}`
	re := regexp.MustCompile(datePattern)

	// Find the first match for the date pattern
	date := re.FindString(filePath)
	if date == "" {
		log.Printf("Failed to extract date from file path %q", filePath)
	}
	return date
}