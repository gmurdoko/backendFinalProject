package userAccountUsecase

import "finalproject/main/master/models"

type UserAccount interface {
	GetUser(*models.UserModel) (bool, error)
	CreateUser(*models.UserModel) (*models.UserModel, error)
}
