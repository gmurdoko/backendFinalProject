package models

type Asset struct {
	AssetName  string `json:"asset_name"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	CarCap     string `json:"car_capacity"`
	MotorCap   string `json:"motor_capacity"`
	BicycleCap string `json:"bicycle_capacity"`
	Status     string `json:"status"`
}
