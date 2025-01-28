package main

import (
	"fmt"
	connectdb "main/data-access"
	transaction "main/data-access/transaction"
	parsers "main/parsers"
)

func main() {
	fmt.Printf("parsing transactions...\n")
	transactions := parsers.ParseTransactionFile()

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
