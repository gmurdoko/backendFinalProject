package models

type Assets struct {
	ID         string `json:"id"`
	AssetName  string `json:"asset_name"`
	AssetArea  string `json:"asset_area"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	CarCap     string `json:"car_capacity"`
	MotorCap   string `json:"motor_capacity"`
	BicycleCap string `json:"bicycle_capacity"`
	Photo      string `json:"photo"`
	Saldo      string `json:"saldo"`
	Status     string `json:"status"`
}

type AssetLocation struct {
	ID         string `json:"id"`
	AssetName  string `json:"asset_name"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	Photo      string `json:"photo"`
	CarCap     string `json:"car_capacity_available"`
	MotorCap   string `json:"motorcycle_capacity_available"`
	BicycleCap string `json:"bicycle_capacity_available"`
	Rating     string `json:"rating"`
	// Status    string `json:"status"`
}

type AssetCapacity struct {
	CarCapMax           string `json:"car_capacity"`
	CarCapAvailable     string `json:"car_capacity_available"`
	MotorCapMax         string `json:"motor_capacity"`
	MotorCapAvailable   string `json:"motor_capacity_available"`
	BicycleCapMax       string `json:"bicycle_capacity"`
	BicycleCapAvailable string `json:"bicycle_capacity_available"`
}
