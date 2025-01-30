package main

import (
	"fmt"
	"log"
	asset "main/data-access/asset-allocations"
	"main/data-access/transaction"
	parsers "main/parsers"

	// transaction "main/data-access/transaction"
	connectdb "main/data-access"
)

const filePath = "../arquivos-status/"
const assetSheetFile1 = "carteira-patrimonio-export-2025-01-29_BC.xlsx"
const assetSheetOwner1 = "Bruna"
const assetSheetFile2 = "carteira-patrimonio-export-2025-01-29_R.xlsx"
const assetSheetOwner2 = "Romano"

const transactionSheetFile1="carteira-export-2025-01-29_BC.xlsx"


func setupAssetAllocations(assetSheetFile string, assetOwner string) (bool, error) {
	allocations, err := parsers.ParseAllSheetsAssetAllocations(filePath+assetSheetFile, assetOwner)
	if err != nil {
		log.Fatalf("Error parsing asset allocations: %v", err)
	}

	// // print allocations
	// for _, allocation := range allocations {
	// fmt.Printf("%+v\n", allocation)
	// }

	// adding allocations to database
	for _, alcs := range allocations {
		_, err := asset.AddAssetAllocation(alcs)
		if err != nil {
			fmt.Println(err)
	}
	}
	return true, err;
}

func setupTransactions(transactionSheetFile string) {
	transactions, err := parsers.ParseTransactionFile(filePath+transactionSheetFile)
	if err != nil{
		fmt.Println(err)
	}

	// // printing the parsed transactions
	// for _, transaction := range transactions {
	// 	fmt.Printf("%+v\n", transaction)
	// }

	// adding the parsed transactions to the database - schema: transactions
	fmt.Printf("adding all transactions to db\n")
	for _, trs := range transactions {
		_, err := transaction.AddTransaction(trs)
		if err != nil {
			fmt.Println(err)
		}
	}
}


func main() {

	connectdb.ConnectDB();

	percentages, err := asset.CalculateAssetTypePercentages()
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("------------------------------------")

	for _, per := range percentages {
		fmt.Printf("AssetType: %s, Percentage: %.2f%%, Balance: %.2f\n", per.AssetType, per.Percentage, per.Balance)
	}

	fmt.Println("------------------------------------")

	// st1, err := setupAssetAllocations(assetSheetFile1, assetSheetOwner1);
	// if err != nil {
	// 	log.Printf("Error setting up allocations for file: %s: ", assetSheetFile1)
	// }
	
	// st2, err := setupAssetAllocations(assetSheetFile2, assetSheetOwner2);
	// if err != nil {
	// 	log.Printf("Error setting up allocations for file: %s: ", assetSheetFile2)
	// }

	// if st1 && st2 {
	// 	setupTransactions(transactionSheetFile1);

	// }
	}
