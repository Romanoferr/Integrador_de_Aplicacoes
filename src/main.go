package main

import (
	"fmt"
	"main/parsers"
)

func main() {
	transactions := parsers.ParseTransactionFile()
	for _, transaction := range transactions {
		fmt.Printf("%+v\n", transaction)
	}
}
