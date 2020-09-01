package model

//Tickets is a Tickect blueprint
type Tickets struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	AssetID      string `json:"asset_id"`
	FeeID        string `json:"fee_id"`
	VehicleID    string `json:"vehicle_id"`
	LicensePlate string `json:"license_plate"`
	StartAt      string `json:"start_at"`
	FinishedAt   string `json:"finished_at"`
	Status       string `json:"status"`
}
