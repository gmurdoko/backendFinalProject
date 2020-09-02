package models

//Providers is a Tickect blueprint
type Providers struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Photo       string `json:"photo"`
	Borndate    string `json:"borndate"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_at"`
	EditedAt    string `json:"edited_at"`
	DeletedAt   string `json:"deleted_at"`
	Status      string `json:"status"`
}
