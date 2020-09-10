package models

type Review struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	AssetID   string `json:"asset_id"`
	Rating    string `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"crated_at"`
	Status    string `json:"status"`
}

type ReviewManagement struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	AssetID   string `json:"asset_id"`
	Rating    string `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"crated_at"`
	Status    string `json:"status"`
}