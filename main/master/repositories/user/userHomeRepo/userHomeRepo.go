package userHomeRepo

import "finalproject/main/master/models"

type UserHome interface {
	GetSaldo(string) (int, error)
	DeleteUserPhoto(string) error
	UpdateUserData(*models.UserModel, string) (*models.UserModel, error)
	UpdateUserSaldoTopUp(*models.WalletModel, string) (int, error)
	GetUserPhoto(string) (string, error)
	// UpdateUserPhoto(*models.UserModel, string) (*models.UserModel, error)
}
