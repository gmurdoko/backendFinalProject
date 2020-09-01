package userRepo

import "finalproject/main/master/models"

type UserRepository interface {
	GetUser(*models.UserModel) (bool, error)
	CreateUser(*models.UserModel) (*models.UserModel, error)
	GetSaldo(string) (int, error)
	DeleteUserPhoto(string) error
	UpdateUserData(*models.UserModel, string) (*models.UserModel, error)
	UpdateUserSaldoTopUp(*models.WalletModel, string) (int, error)
	GetUserPhoto(string) (string, error)
	UpdateUserPhoto(*models.UserModel, string) (string, error)
}
