package models

//Fees fee model
type Fees struct {
	ID        string `json:"id"`
	VehicleID string `json:"vehicle_id"`
	Fee       string `json:"fee"`
	CreatedAt string `json:"created_at"`
	EditedAt  string `json:"edited_at"`
	DeletedAt string `json:"deleted_at"`
	Status    string `json:"status"`
}
