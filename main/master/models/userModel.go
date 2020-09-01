package models

type UserModel struct {
	ID          string      `json:"id"`
	IdWallet    WalletModel `json:"id_wallet"`
	Username    string      `json:"username"`
	Password    string      `json:"password"`
	Email       string      `json:"email"`
	Fullname    string      `json:"fullname"`
	Photo       string      `json:"photo"`
	BornDate    string      `json:"borndate"`
	PhoneNumber string      `json:"phone_number"`
	Address     string      `json:"address"`
	CreatedAt   string      `json:"created_at"`
	EditedAt    string      `json:"edited_at"`
	DeletedAt   string      `json:"deleted_at"`
	Status      string      `json:"status"`
}
