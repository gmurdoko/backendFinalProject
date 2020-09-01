package userUsecase

import "finalproject/main/master/models"

type UserUsecase interface {
	CreateUser(*models.UserModel) (*models.UserModel, error)
	GetSaldo(string) (int, error)
	DeleteUserPhoto(string) error
	UpdateUserData(*models.UserModel, string) (*models.UserModel, error)
	UpdateUserSaldoTopUp(*models.WalletModel, string) (int, error)
	GetUserPhoto(string) (string, error)
	UpdateUserPhoto(*models.UserModel, string) (string, error)
}
