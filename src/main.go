package main

import (
	"log"
	connectdb "main/data-access"
	setup "main/setup"
)

const filePathAssets = "../arquivos-status/asset-files/"


func main() {
	connectdb.ConnectDB()

	// setup.CalculateAndDisplaySomestuff("2025-01-30")

	assetFiles, owners, err := setup.GetAllAssetFiles(filePathAssets)
	if err != nil {
		log.Fatalf("Error getting asset files: %v", err)
	}

	for i, assetFile := range assetFiles {
		_, err := setup.SetupAssetAllocations(filePathAssets+assetFile, owners[i])
		if err != nil {
			log.Printf("Error setting up allocations for file: %s: %v", assetFile, err)
		}
	}
}
