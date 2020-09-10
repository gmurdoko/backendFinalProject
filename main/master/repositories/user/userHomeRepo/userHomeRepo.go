package userHomeRepo

import "finalproject/main/master/models"

type UserHome interface {
	GetSaldo(string) (*models.UserWallet, error)
	DeleteUserPhoto(string) error
	UpdateUserData(*models.UserModel, string) (*models.UserModel, error)
	UpdateUserSaldoTopUp(*models.Wallets, string) (string, error)
	GetUserPhoto(string) (*string, error)
	UpdateUserPhoto(string, string) error
<<<<<<< Updated upstream
	GetUserTicket(string) (*models.TicketUser, error)
=======
	GetUserTicket(*models.Ticket) (*models.Ticket, error)
>>>>>>> Stashed changes
}
