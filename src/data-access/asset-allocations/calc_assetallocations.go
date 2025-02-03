package asset

import (
	"fmt"
	"sort"
)

type AssetBalance struct {
	AssetType string
	Balance   float64
	Percentage float64
}


func CalculateSumOfBalances(date string) (float64, error) {
    sumTotalBalance, err := SumTotalBalance(date)
    if err != nil {
        return 0.0, err
    }

    return sumTotalBalance, nil
}

func CalculateAssetTypeTotalBalance(assetType string) (float64, error) {
	totalBalance, err := TotalBalanceByAssetType(assetType)
	if err != nil {
        return 0.0, err
    }

	return totalBalance, nil
}

func CalculateAssetTypeTotalBalances(date string) (map[string]float64, error) {
    totalBalances, err := TotalBalanceByAllAssetTypes(date)
    if err != nil {
        return nil, err
    }

    totalBalances["ETF Exterior"] *= 5.86

    return totalBalances, nil
}

func CalculateAssetTypePercentages(date string) ([]AssetBalance, error) {
    totalBalances, err := TotalBalanceByAllAssetTypes(date)
    if err != nil {
        return nil, err
    }

    totalBalances["ETF Exterior"] *= 5.86

    totalSum := 0.0
    for _, balance := range totalBalances {
        totalSum += balance
    }

    if totalSum == 0 {
        return nil, fmt.Errorf("total sum of balances is zero")
    }

    var sortedBalances []AssetBalance
    for assetType, balance := range totalBalances {
        percentage := (balance / totalSum) * 100
        sortedBalances = append(sortedBalances, AssetBalance{assetType, balance, percentage})
    }

    // Sort the slice by balance in descending order
    sort.Slice(sortedBalances, func(i, j int) bool {
        return sortedBalances[i].Balance > sortedBalances[j].Balance
    })

    return sortedBalances, nil
}