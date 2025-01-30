package asset

import (
	connectdb "main/data-access"
)

type AssetAllocation struct {
	ID                  int64
	AssetAllocationDate string
	AssetOwner          string
	AssetIdentifier     string
	AssetType           string
	MedianPrice         float64
	ActualPrice         float64
	MedianReturn        float64
	Quantity            float64
	Balance             float64
	TodayReturn         float64
}

func AddAssetAllocation(assetAllocation AssetAllocation) (int64, error) {
	result, err := connectdb.Db.Exec("INSERT INTO asset_allocations (asset_allocation_date, asset_owner, asset_id, asset_type, median_price, actual_price, median_return, quantity, balance, today_return) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		assetAllocation.AssetAllocationDate, assetAllocation.AssetOwner, assetAllocation.AssetIdentifier, assetAllocation.AssetType, assetAllocation.MedianPrice, assetAllocation.ActualPrice, assetAllocation.MedianReturn, assetAllocation.Quantity, assetAllocation.Balance, assetAllocation.TodayReturn)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func AssetAllocationByAssetId(assetId string) ([]AssetAllocation, error) {
	var assetAllocations []AssetAllocation
	rows, err := connectdb.Db.Query("SELECT id, asset_allocation_date, asset_owner, asset_id, asset_type, median_price, actual_price, median_return, quantity, balance, today_return FROM asset_allocations WHERE asset_id = ?", assetId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assetAllocation AssetAllocation
		if err := rows.Scan(&assetAllocation.ID, &assetAllocation.AssetAllocationDate, &assetAllocation.AssetOwner, &assetAllocation.AssetIdentifier, &assetAllocation.AssetType, &assetAllocation.MedianPrice, &assetAllocation.ActualPrice, &assetAllocation.MedianReturn, &assetAllocation.Quantity, &assetAllocation.Balance, &assetAllocation.TodayReturn); err != nil {
			return nil, err
		}
		assetAllocations = append(assetAllocations, assetAllocation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return assetAllocations, nil
}

func AssetAllocationByAssetType(assetType string) ([]AssetAllocation, error) {
	var assetAllocations []AssetAllocation
	rows, err := connectdb.Db.Query("SELECT id, asset_allocation_date, asset_owner, asset_id, asset_type, median_price, actual_price, median_return, quantity, balance, today_return FROM asset_allocations WHERE asset_type = ?", assetType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assetAllocation AssetAllocation
		if err := rows.Scan(&assetAllocation.ID, &assetAllocation.AssetAllocationDate, &assetAllocation.AssetOwner, &assetAllocation.AssetIdentifier, &assetAllocation.AssetType, &assetAllocation.MedianPrice, &assetAllocation.ActualPrice, &assetAllocation.MedianReturn, &assetAllocation.Quantity, &assetAllocation.Balance, &assetAllocation.TodayReturn); err != nil {
			return nil, err
		}
		assetAllocations = append(assetAllocations, assetAllocation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return assetAllocations, nil
}

func AllAssetAllocations() ([]AssetAllocation, error) {
	var assetAllocations []AssetAllocation
	rows, err := connectdb.Db.Query("SELECT id, asset_allocation_date, asset_owner, asset_id, asset_type, median_price, actual_price, median_return, quantity, balance, today_return FROM asset_allocations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assetAllocation AssetAllocation
		if err := rows.Scan(&assetAllocation.ID, &assetAllocation.AssetAllocationDate, &assetAllocation.AssetOwner, &assetAllocation.AssetIdentifier, &assetAllocation.AssetType, &assetAllocation.MedianPrice, &assetAllocation.ActualPrice, &assetAllocation.MedianReturn, &assetAllocation.Quantity, &assetAllocation.Balance, &assetAllocation.TodayReturn); err != nil {
			return nil, err
		}
		assetAllocations = append(assetAllocations, assetAllocation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return assetAllocations, nil
}

func DeleteAssetAllocation(id int64) (int64, error) {
	result, err := connectdb.Db.Exec("DELETE FROM asset_allocations WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func TotalBalanceByAssetType(assetType string) (float64, error) {
	var totalBalance float64
	err := connectdb.Db.QueryRow("SELECT SUM(balance) FROM asset_allocations WHERE asset_type = ?", assetType).Scan(&totalBalance)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}

func TotalBalanceByAllAssetTypes() (map[string]float64, error) {
	balances := make(map[string]float64)
	rows, err := connectdb.Db.Query("SELECT asset_type, SUM(balance) FROM asset_allocations GROUP BY asset_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assetType string
		var totalBalance float64
		if err := rows.Scan(&assetType, &totalBalance); err != nil {
			return nil, err
		}
		balances[assetType] = totalBalance
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return balances, nil
}

func SumTotalBalance() (float64, error) {
	var totalBalance float64
	err := connectdb.Db.QueryRow("SELECT SUM(balance) FROM asset_allocations").Scan(&totalBalance)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}