package models

//Tickets is a Tickect blueprint
type Tickets struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	AssetID      string `json:"asset_id"`
	FeeID        string `json:"fee_id"`
	VehicleID    string `json:"vehicle_id"`
	LicensePlate string `json:"license_plate"`
	BookAt       string `json:"book_at"`
	StartAt      string `json:"start_at"`
	FinishedAt   string `json:"finished_at"`
	Status       string `json:"status"`
}

//TicketView is Tickets detail view
type TicketView struct {
	ID                  string `json:"id"`
	Username            string `json:"username"`
	AssetName           string `json:"asset_name"`
	LicensePlate        string `json:"license_plate"`
	BasedFee            string `json:"based_fee"`
	ParkingDurationHour string `json:"parking_duration_hour"`
	PayFee              string `json:"pay_fee"`
	BookAt              string `json:"book_at"`
	StartAt             string `json:"start_at"`
	FinishedAt          string `json:"finished_at"`
}
type TicketUser struct {
	UserID       string `json:"user_id"`
	AssetName    string `json:"asset_name"`
	Vehicle_type string `json:"vehicle_type"`
	LicensePlate string `json:"license_plate"`
	BookAt       string `json:"book_at"`
}
