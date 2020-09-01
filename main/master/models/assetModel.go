package models

type AssetModel struct {
	ID                 string  `json:"id"`
	ProviderId         string  `json:"provider_id"`
	AssetName          string  `json:"asset_name"`
	AssetArea          float64 `json:"asset_area"`
	Longitude          float64 `json:"longitude"`
	Latitude           float64 `json:"latitude"`
	CarCapacity        int     `json:"car_capacity"`
	MotorcycleCapacity int     `json:"motorcycle_capacity"`
	BicycleCapacity    int     `json:"bicycle_capacity"`
	CreatedAt          string  `json:"created_at"`
	EditedAt           string  `json:"edited_at"`
	DeletedAt          string  `json:"deleted_at"`
	Status             string  `json:"status"`
}
