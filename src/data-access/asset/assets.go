package asset

type Asset struct {
	ID              int64
	AssetIdentifier string
	AssetType       string
}

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