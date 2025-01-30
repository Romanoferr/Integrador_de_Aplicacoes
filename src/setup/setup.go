package setup

import (
	"fmt"
	"log"
	asset "main/data-access/asset-allocations"
	"main/data-access/transaction"
	parsers "main/parsers"
	"os"
	"strings"
	// transaction "main/data-access/transaction"
)

func GetAllAssetFiles(directory string) ([]string, []string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, nil, err
	}

	var assetFiles []string
	var owners []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			assetFiles = append(assetFiles, file.Name())
			if strings.Contains(file.Name(), "_R") {
				owners = append(owners, "Romano")
			} else if strings.Contains(file.Name(), "_BC") {
				owners = append(owners, "Bruna")
			} else {
				owners = append(owners, "undefined_Owner")
			}
		}
	}
	return assetFiles, owners, nil
}

func CalculateAndDisplaySomestuff() {
	percentages, err := asset.CalculateAssetTypePercentages()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------------------------------------------------------")

	for _, per := range percentages {
		fmt.Printf("Porcentagem: %.2f%%, Balance: R$%.2f -->> %s\n", per.Percentage, per.Balance, per.AssetType)
	}
	fmt.Println("-------------------------------------------------------")

	assetTypeTotalBalances, err := asset.CalculateAssetTypeTotalBalances()
	if err != nil {
		fmt.Println(err)
	}

	totalSum := 0.0
	for _, totalBalance := range assetTypeTotalBalances {
		totalSum += totalBalance
	}

	fmt.Printf("Total Asset: R$%.2f\n", totalSum)
	fmt.Println("-------------------------------------------------------")
}

func SetupAssetAllocations(assetSheetFile string, assetOwner string) (bool, error) {
	allocations, err := parsers.ParseAllSheetsAssetAllocations(assetSheetFile, assetOwner)
	if err != nil {
		log.Fatalf("Error parsing asset allocations: %v", err)
	}

	// adding allocations to database
	for _, alcs := range allocations {
		_, err := asset.AddAssetAllocation(alcs)
		if err != nil {
			fmt.Println(err)
		}
	}
	return true, err
}

func SetupTransactions(transactionSheetFile string) {
	transactions, err := parsers.ParseTransactionFile(transactionSheetFile)
	if err != nil {
		fmt.Println(err)
	}

	// adding the parsed transactions to the database - schema: transactions
	fmt.Printf("adding all transactions to db\n")
	for _, trs := range transactions {
		_, err := transaction.AddTransaction(trs)
		if err != nil {
			fmt.Println(err)
		}
	}
}
