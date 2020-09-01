package models

type ReportAssetDaily struct {
	AssetName    string `json:"asset_name"`
	Date         string `json:"date"`
	TotalParked  string `json:"total_parked"`
	TotalRevenue string `json:"total_revenue"`
}

type ReportAssetMonthly struct {
	AssetName    string `json:"asset_name"`
	Months       string `json:"months"`
	TotalParked  string `json:"total_parked"`
	TotalRevenue string `json:"total_revenue"`
}
