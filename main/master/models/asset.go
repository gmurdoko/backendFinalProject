package models

type AssetLocation struct {
	AssetName string `json:"asset_name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Status    string `json:"status"`
}

type AssetCapacity struct {
	CarCapMax           string `json:"car_capacity"`
	CarCapAvailable     string `json:"car_capacity_available"`
	MotorCapMax         string `json:"motor_capacity"`
	MotorCapAvailable   string `json:"motor_capacity_available"`
	BicycleCapMax       string `json:"bicycle_capacity"`
	BicycleCapAvailable string `json:"bicycle_capacity_available"`
}
