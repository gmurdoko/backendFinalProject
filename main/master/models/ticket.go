package models

type Ticket struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	AssetID      string `json:"asset_id"`
	FeeID        string `json:"fee_id"`
	VehicleID    string `json:"vehicle_id"`
	LicensePlate string `json:"license_plate"`
	BookAt       string `json:"book_at"`
	Status       string `json:"status"`
}
