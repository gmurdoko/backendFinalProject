package models

type Assets struct {
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
