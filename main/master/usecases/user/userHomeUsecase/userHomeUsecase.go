package userHomeUsecase

import (
	"finalproject/main/master/models"
	"mime/multipart"
)

type UserHome interface {
	GetSaldo(string) (*models.UserWallet, error)
	DeleteUserPhoto(string) error
	UpdateUserData(*models.UserModel, string) (*models.UserModel, error)
	UpdateUserSaldoTopUp(*models.Wallets, string) (string, error)
	GetUserPhoto(string) (*string, error)
	UpdateUserPhoto(photo multipart.File, handler *multipart.FileHeader, id string) error
}
