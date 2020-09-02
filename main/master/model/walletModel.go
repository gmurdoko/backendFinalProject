package model

//Wallets is a Tickect blueprint
type Wallets struct {
	ID        string `json:"id"`
	Saldo     string `json:"saldo"`
	Debit     string `json:"debit"`
	Kredit    string `json:"kredit"`
	CreatedAt string `json:"created_at"`
	EditedAt  string `json:"edited_at"`
	DeletedAt string `json:"deleted_at"`
	Status    string `json:"status"`
}
