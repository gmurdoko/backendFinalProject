package models

type WalletModel struct {
	ID        string `json:"id"`
	Saldo     int    `json:"saldo"`
	Debit     int    `json:"debit"`
	Kredit    int    `json:"kredit"`
	CreatedAt string `json:"created_at"`
	EditedAt  string `json:"edited_at"`
	DeletedAt string `json:"deleted_at"`
	Status    string `json:"status"`
}
