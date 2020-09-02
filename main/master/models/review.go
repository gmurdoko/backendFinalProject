package models

type review struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	AssetID   string `json:"asset_id"`
	Rating    string `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"crated_at"`
	Status    string `json:"status"`
}
