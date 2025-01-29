package main

import (
	"fmt"
	"log"
	connectdb "main/data-access"
	transaction "main/data-access/transaction"
	parsers "main/parsers"
)

const filePath = "../arquivos-statusinvest/"
const assetSheetFile = "StatusInvest-assets-2025-01-28--12-43-30_BC.xlsx"
const assetSheetOwner = "Bruna"

func main() {
	allocations, err := parsers.ParseAllSheetsAssetAllocations(filePath+assetSheetFile, assetSheetOwner)
		if err != nil {
			log.Fatalf("Error parsing asset allocations: %v", err)
		}
	

	// print allocations BC
	for _, allocation := range allocations {
		fmt.Printf("%+v\n", allocation)
	}

	allocationsR, err := parsers.ParseAllSheetsAssetAllocations(filePath+"StatusInvest-assets-2025-01-22--23-43-37.xlsx", "Romano")
		if err != nil {
			log.Fatalf("Error parsing asset allocations: %v", err)
		}

	// print allocations R	
	for _, allocation := range allocationsR {
		fmt.Printf("%+v\n", allocation)
	}

	fmt.Printf("parsing transactions...\n")
	transactions := parsers.ParseTransactionFile("StatusInvest-transactions-2025-01-28--12-45-32_BC.xlsx")

	// printing the parsed transactions
	for _, transaction := range transactions {
		fmt.Printf("%+v\n", transaction)
	}

	fmt.Printf("connecting to do db\n")
	connectdb.ConnectDB();

	// adding the parsed transactions to the database - schema: transactions
	fmt.Printf("adding all transactions to db\n")
	for _, trs := range transactions {
		_, err := transaction.AddTransaction(trs)
		if err != nil {
			fmt.Println(err)
		}
	}

	trs1, err := transaction.TransactionByAssetId("BBAS3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", trs1)

}
