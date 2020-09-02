package userAccountRepo

import "finalproject/main/master/models"

type UserAccount interface {
	GetUser(*models.UserModel) (*models.UserModel, bool, error)
	CreateUser(*models.UserModel) (*models.UserModel, error)
	GetUserById(string) (*models.UserModel, error)
}
